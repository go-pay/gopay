package paygateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	wechatv3 "github.com/go-pay/gopay/wechat/v3"
)

type Server struct {
	cfg       *Config
	store     *MerchantStore
	clients   *ClientManager
	publisher EventPublisher
	idem      IdempotencyStore
	dedup     CallbackDeduper
	nonce     NonceStore
	outbox    *OutboxWorker
	cfgSync   *MerchantSyncWorker
	mux       *http.ServeMux
	srv       *http.Server
}

func NewServer(cfg *Config) (*Server, error) {
	store, err := NewMerchantStore(cfg)
	if err != nil {
		return nil, err
	}
	clients, err := NewClientManager(cfg, store)
	if err != nil {
		return nil, err
	}

	redisClient, err := newRedisClient(cfg.Redis)
	if err != nil {
		return nil, err
	}

	var nonceStore NonceStore
	if cfg.SharedAuth.SharedSecret != "" || cfg.SharedAuth.SharedSecretPrev != "" {
		if redisClient != nil {
			nonceStore = NewRedisNonceStore(redisClient, cfg.Redis.KeyPrefix)
		} else {
			nonceStore = NewMemoryNonceStore()
		}
	}

	var outboxWorker *OutboxWorker
	var publisher EventPublisher
	if cfg.JavaWebhook.URL != "" {
		webhook := NewWebhookPublisher(cfg.JavaWebhook.URL, cfg.JavaWebhook.Token, cfg.SharedAuth.SharedSecret, clients.HTTPClient(), time.Duration(cfg.JavaWebhook.TimeoutMillis)*time.Millisecond)
		if cfg.JavaWebhook.Async {
			if redisClient == nil {
				return nil, errors.New("javaWebhook.async requires redis")
			}
			stream := cfg.Redis.KeyPrefix + "outbox"
			worker, err := NewOutboxWorker(redisClient, OutboxWorkerConfig{
				Stream:  stream,
				Group:   cfg.JavaWebhook.ConsumerGroup,
				MinIdle: 30 * time.Second,
				Block:   2 * time.Second,
				Count:   16,
			}, webhook)
			if err != nil {
				return nil, err
			}
			publisher = NewRedisOutboxPublisher(redisClient, stream)
			outboxWorker = worker
		} else {
			publisher = webhook
		}
	} else {
		publisher = EventPublisherFunc(func(ctx context.Context, event *Event) error {
			log.Printf("event: %+v", event)
			return nil
		})
	}

	s := &Server{
		cfg:       cfg,
		store:     store,
		clients:   clients,
		publisher: publisher,
		nonce:     nonceStore,
		outbox:    outboxWorker,
		mux:       http.NewServeMux(),
	}
	if cfg.MerchantSync.SnapshotURL != "" {
		s.cfgSync = NewMerchantSyncWorker(MerchantSyncWorkerOptions{
			SnapshotURL:     cfg.MerchantSync.SnapshotURL,
			PollInterval:    time.Duration(max(1, cfg.MerchantSync.PollIntervalSeconds)) * time.Second,
			ChangeStream:    cfg.MerchantSync.ChangeStream,
			ConsumerGroup:   cfg.MerchantSync.ChangeConsumerGroup,
			Redis:           redisClient,
			HTTPClient:      clients.HTTPClient(),
			SharedAuth:      cfg.SharedAuth,
			DefaultTenantID: cfg.DefaultTenantID,
			Store:           store,
			Clients:         clients,
		})
	}
	if redisClient != nil {
		s.idem = NewRedisIdempotencyStore(redisClient, cfg.Redis.KeyPrefix, time.Duration(cfg.Redis.IdempotencyTTLSeconds)*time.Second)
		s.dedup = NewRedisCallbackDeduper(redisClient, cfg.Redis.KeyPrefix, time.Duration(cfg.Redis.CallbackProcessingTTLSeconds)*time.Second, time.Duration(cfg.Redis.CallbackDedupTTLSeconds)*time.Second)
	} else {
		s.idem = NewMemoryIdempotencyStore(24 * time.Hour)
		s.dedup = NewMemoryCallbackDeduper(5*time.Minute, 7*24*time.Hour)
	}
	s.routes()

	s.srv = &http.Server{
		Addr:         cfg.Server.Addr,
		Handler:      s.mux,
		ReadTimeout:  time.Duration(max(1, cfg.Server.ReadTimeoutSeconds)) * time.Second,
		WriteTimeout: time.Duration(max(1, cfg.Server.WriteTimeoutSeconds)) * time.Second,
	}
	return s, nil
}

func (s *Server) ListenAndServe() error {
	if s.outbox != nil {
		if err := s.outbox.Start(); err != nil {
			return err
		}
		defer s.outbox.Stop()
	}
	if s.cfgSync != nil {
		if err := s.cfgSync.Start(); err != nil {
			return err
		}
		defer s.cfgSync.Stop()
	}
	return s.srv.ListenAndServe()
}

func (s *Server) routes() {
	s.mux.HandleFunc("GET /healthz", s.handleHealthz)

	s.mux.HandleFunc("POST /v1/payments", s.withInternalAuth(s.handleCreatePayment))
	s.mux.HandleFunc("/v1/payments/", s.withInternalAuth(s.handlePaymentsSubroutes))

	s.mux.HandleFunc("POST /v1/refunds", s.withInternalAuth(s.handleCreateRefund))
	s.mux.HandleFunc("/v1/refunds/", s.withInternalAuth(s.handleRefundsSubroutes))

	s.mux.HandleFunc("POST /v1/compensations/payments/query", s.withInternalAuth(s.handleCompensationQueryPayments))

	s.mux.HandleFunc("/callbacks/wechat/v3/", s.handleWechatV3Callback)
	s.mux.HandleFunc("/callbacks/alipay/", s.handleAlipayCallback)
}

func (s *Server) handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("ok"))
}

func (s *Server) handlePaymentsSubroutes(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/v1/payments/")
	if path == "" {
		http.NotFound(w, r)
		return
	}
	parts := strings.Split(path, "/")
	if len(parts) == 1 && r.Method == http.MethodGet {
		s.handleQueryPayment(w, r, parts[0])
		return
	}
	if len(parts) == 2 && parts[1] == "close" && r.Method == http.MethodPost {
		s.handleClosePayment(w, r, parts[0])
		return
	}
	http.NotFound(w, r)
}

func (s *Server) handleRefundsSubroutes(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/v1/refunds/")
	if path == "" {
		http.NotFound(w, r)
		return
	}
	parts := strings.Split(path, "/")
	if len(parts) == 1 && r.Method == http.MethodGet {
		s.handleQueryRefund(w, r, parts[0])
		return
	}
	http.NotFound(w, r)
}

// =================================================
// Helpers

func decodeJSON[T any](r *http.Request, out *T) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(out)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeJSONBytes(w http.ResponseWriter, status int, v any) []byte {
	bs, err := json.Marshal(v)
	if err != nil {
		http.Error(w, "marshal json failed", http.StatusInternalServerError)
		return nil
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(bs)
	_, _ = w.Write([]byte("\n"))
	return bs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (s *Server) withInternalAuth(next http.HandlerFunc) http.HandlerFunc {
	expectedToken := strings.TrimSpace(s.cfg.APIAuth.Token)
	shared := s.cfg.SharedAuth.SharedSecret != "" || s.cfg.SharedAuth.SharedSecretPrev != ""
	if expectedToken == "" && !shared {
		return next
	}
	return func(w http.ResponseWriter, r *http.Request) {
		// Prefer shared HMAC auth when configured.
		if shared {
			if !s.cfg.SharedAuth.Required && expectedToken != "" {
				if token := firstNonEmpty(r.Header.Get(HeaderPayToken), r.Header.Get(HeaderPayTokenLegacy)); token != "" && token == expectedToken {
					next(w, r)
					return
				}
			}
			if _, err := verifyHMACRequest(r, s.cfg.SharedAuth, s.nonce); err == nil {
				next(w, r)
				return
			} else if !s.cfg.SharedAuth.Required && expectedToken != "" {
				if token := firstNonEmpty(r.Header.Get(HeaderPayToken), r.Header.Get(HeaderPayTokenLegacy)); token != "" && token == expectedToken {
					next(w, r)
					return
				}
			}
			writeJSON(w, http.StatusUnauthorized, map[string]string{
				"code":    "UNAUTHORIZED",
				"message": "missing or invalid shared auth",
			})
			return
		}

		// Legacy token-only auth.
		token := firstNonEmpty(r.Header.Get(HeaderPayToken), r.Header.Get(HeaderPayTokenLegacy))
		if token == "" || token != expectedToken {
			writeJSON(w, http.StatusUnauthorized, map[string]string{
				"code":    "UNAUTHORIZED",
				"message": "missing or invalid X-Pay-Token",
			})
			return
		}
		next(w, r)
	}
}

func (s *Server) callbackURL(channel string, tenantID, merchantID string) (string, error) {
	if s.cfg.PublicBaseURL == "" {
		return "", errors.New("publicBaseUrl is required to generate callback urls")
	}
	base := strings.TrimRight(s.cfg.PublicBaseURL, "/")
	switch channel {
	case string(ChannelWechatV3):
		return fmt.Sprintf("%s/callbacks/wechat/v3/%s/%s", base, tenantID, merchantID), nil
	case string(ChannelAlipay):
		return fmt.Sprintf("%s/callbacks/alipay/%s/%s", base, tenantID, merchantID), nil
	default:
		return "", fmt.Errorf("unsupported channel: %s", channel)
	}
}

// =================================================
// Payment APIs

func (s *Server) handleCreatePayment(w http.ResponseWriter, r *http.Request) {
	var req CreatePaymentRequest
	if err := decodeJSON(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, CreatePaymentResponse{Code: "BAD_REQUEST", Message: err.Error()})
		return
	}
	if req.TenantID == "" {
		req.TenantID = s.cfg.DefaultTenantID
	}
	if req.TenantID == "" || req.MerchantID == "" || req.OutTradeNo == "" || req.Currency == "" || req.Amount <= 0 || req.Subject == "" {
		writeJSON(w, http.StatusBadRequest, CreatePaymentResponse{Code: "BAD_REQUEST", Message: "missing required fields"})
		return
	}

	idemKey := r.Header.Get("X-Idempotency-Key")
	if idemKey == "" {
		idemKey = fmt.Sprintf("payment:%s:%s:%s:%s", req.TenantID, req.MerchantID, req.Channel, req.OutTradeNo)
	}
	if status, body, ok, err := s.idem.Get(r.Context(), idemKey); err != nil {
		writeJSON(w, http.StatusInternalServerError, CreatePaymentResponse{Code: "ERROR", Message: "idempotency store error"})
		return
	} else if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write(body)
		return
	}

	switch req.Channel {
	case ChannelWechatV3:
		c, wc, err := s.clients.WechatV3(req.TenantID, req.MerchantID)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, CreatePaymentResponse{Code: "BAD_REQUEST", Message: err.Error()})
			return
		}
		notifyURL, err := s.callbackURL(string(ChannelWechatV3), req.TenantID, req.MerchantID)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, CreatePaymentResponse{Code: "ERROR", Message: err.Error()})
			return
		}

		bm := make(gopay.BodyMap)
		bm.Set("appid", wc.AppID).
			Set("description", firstNonEmpty(req.Description, req.Subject)).
			Set("out_trade_no", req.OutTradeNo).
			Set("notify_url", notifyURL).
			SetBodyMap("amount", func(b gopay.BodyMap) {
				b.Set("total", req.Amount).Set("currency", req.Currency)
			})
		if req.ExpireAt != "" {
			bm.Set("time_expire", req.ExpireAt)
		}

		switch strings.ToUpper(req.Scene) {
		case "JSAPI", "MINIAPP":
			if req.OpenID == "" {
				writeJSON(w, http.StatusBadRequest, CreatePaymentResponse{Code: "BAD_REQUEST", Message: "missing openid for JSAPI/MINIAPP"})
				return
			}
			bm.SetBodyMap("payer", func(b gopay.BodyMap) { b.Set("openid", req.OpenID) })
			wxRsp, err := c.V3TransactionJsapi(r.Context(), bm)
			if err != nil {
				writeJSON(w, http.StatusBadGateway, CreatePaymentResponse{Code: "UPSTREAM_ERROR", Message: err.Error()})
				return
			}
			if wxRsp.Code != wechatv3.Success {
				writeJSON(w, http.StatusBadGateway, CreatePaymentResponse{Code: "UPSTREAM_ERROR", Message: wxRsp.Error, OutTradeNo: req.OutTradeNo})
				return
			}
			prepayID := wxRsp.Response.PrepayId
			var payData any
			if strings.ToUpper(req.Scene) == "MINIAPP" {
				payData, err = c.PaySignOfApplet(wc.AppID, prepayID)
			} else {
				payData, err = c.PaySignOfJSAPI(wc.AppID, prepayID)
			}
			if err != nil {
				writeJSON(w, http.StatusInternalServerError, CreatePaymentResponse{Code: "ERROR", Message: err.Error()})
				return
			}
			resp := CreatePaymentResponse{Code: "OK", OutTradeNo: req.OutTradeNo, Status: "PAYING", PayData: payData}
			bs := writeJSONBytes(w, http.StatusOK, resp)
			if bs != nil {
				if err := s.idem.Put(r.Context(), idemKey, http.StatusOK, bs); err != nil {
					log.Printf("idempotency put failed: %v", err)
				}
			}
			return

		case "APP":
			wxRsp, err := c.V3TransactionApp(r.Context(), bm)
			if err != nil {
				writeJSON(w, http.StatusBadGateway, CreatePaymentResponse{Code: "UPSTREAM_ERROR", Message: err.Error()})
				return
			}
			if wxRsp.Code != wechatv3.Success {
				writeJSON(w, http.StatusBadGateway, CreatePaymentResponse{Code: "UPSTREAM_ERROR", Message: wxRsp.Error, OutTradeNo: req.OutTradeNo})
				return
			}
			payData, err := c.PaySignOfApp(wc.AppID, wxRsp.Response.PrepayId)
			if err != nil {
				writeJSON(w, http.StatusInternalServerError, CreatePaymentResponse{Code: "ERROR", Message: err.Error()})
				return
			}
			resp := CreatePaymentResponse{Code: "OK", OutTradeNo: req.OutTradeNo, Status: "PAYING", PayData: payData}
			bs := writeJSONBytes(w, http.StatusOK, resp)
			if bs != nil {
				if err := s.idem.Put(r.Context(), idemKey, http.StatusOK, bs); err != nil {
					log.Printf("idempotency put failed: %v", err)
				}
			}
			return

		case "H5":
			wxRsp, err := c.V3TransactionH5(r.Context(), bm)
			if err != nil {
				writeJSON(w, http.StatusBadGateway, CreatePaymentResponse{Code: "UPSTREAM_ERROR", Message: err.Error()})
				return
			}
			if wxRsp.Code != wechatv3.Success {
				writeJSON(w, http.StatusBadGateway, CreatePaymentResponse{Code: "UPSTREAM_ERROR", Message: wxRsp.Error, OutTradeNo: req.OutTradeNo})
				return
			}
			resp := CreatePaymentResponse{Code: "OK", OutTradeNo: req.OutTradeNo, Status: "PAYING", PayData: wxRsp.Response}
			bs := writeJSONBytes(w, http.StatusOK, resp)
			if bs != nil {
				if err := s.idem.Put(r.Context(), idemKey, http.StatusOK, bs); err != nil {
					log.Printf("idempotency put failed: %v", err)
				}
			}
			return

		case "NATIVE":
			wxRsp, err := c.V3TransactionNative(r.Context(), bm)
			if err != nil {
				writeJSON(w, http.StatusBadGateway, CreatePaymentResponse{Code: "UPSTREAM_ERROR", Message: err.Error()})
				return
			}
			if wxRsp.Code != wechatv3.Success {
				writeJSON(w, http.StatusBadGateway, CreatePaymentResponse{Code: "UPSTREAM_ERROR", Message: wxRsp.Error, OutTradeNo: req.OutTradeNo})
				return
			}
			resp := CreatePaymentResponse{Code: "OK", OutTradeNo: req.OutTradeNo, Status: "PAYING", PayData: wxRsp.Response}
			bs := writeJSONBytes(w, http.StatusOK, resp)
			if bs != nil {
				if err := s.idem.Put(r.Context(), idemKey, http.StatusOK, bs); err != nil {
					log.Printf("idempotency put failed: %v", err)
				}
			}
			return

		default:
			writeJSON(w, http.StatusBadRequest, CreatePaymentResponse{Code: "BAD_REQUEST", Message: "unsupported scene for WECHAT_V3"})
			return
		}

	case ChannelAlipay:
		if strings.ToUpper(req.Currency) != "CNY" {
			writeJSON(w, http.StatusBadRequest, CreatePaymentResponse{Code: "BAD_REQUEST", Message: "ALIPAY only supports CNY in this gateway"})
			return
		}
		c, ac, err := s.clients.Alipay(req.TenantID, req.MerchantID)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, CreatePaymentResponse{Code: "BAD_REQUEST", Message: err.Error()})
			return
		}
		notifyURL, err := s.callbackURL(string(ChannelAlipay), req.TenantID, req.MerchantID)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, CreatePaymentResponse{Code: "ERROR", Message: err.Error()})
			return
		}
		totalAmount, err := formatCNYFenToYuanString(req.Amount)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, CreatePaymentResponse{Code: "BAD_REQUEST", Message: err.Error()})
			return
		}

		bm := make(gopay.BodyMap)
		bm.Set("subject", req.Subject).
			Set("out_trade_no", req.OutTradeNo).
			Set("total_amount", totalAmount).
			Set("notify_url", notifyURL)
		if req.Description != "" {
			bm.Set("body", req.Description)
		}

		switch strings.ToUpper(req.Scene) {
		case "APP":
			orderStr, err := c.TradeAppPay(r.Context(), bm)
			if err != nil {
				writeJSON(w, http.StatusBadGateway, CreatePaymentResponse{Code: "UPSTREAM_ERROR", Message: err.Error()})
				return
			}
			_ = ac // currently unused but kept for symmetry
			resp := CreatePaymentResponse{Code: "OK", OutTradeNo: req.OutTradeNo, Status: "PAYING", PayData: map[string]string{"orderStr": orderStr}}
			bs := writeJSONBytes(w, http.StatusOK, resp)
			if bs != nil {
				if err := s.idem.Put(r.Context(), idemKey, http.StatusOK, bs); err != nil {
					log.Printf("idempotency put failed: %v", err)
				}
			}
			return
		case "WAP":
			payURL, err := c.TradeWapPay(r.Context(), bm)
			if err != nil {
				writeJSON(w, http.StatusBadGateway, CreatePaymentResponse{Code: "UPSTREAM_ERROR", Message: err.Error()})
				return
			}
			resp := CreatePaymentResponse{Code: "OK", OutTradeNo: req.OutTradeNo, Status: "PAYING", PayData: map[string]string{"payUrl": payURL}}
			bs := writeJSONBytes(w, http.StatusOK, resp)
			if bs != nil {
				if err := s.idem.Put(r.Context(), idemKey, http.StatusOK, bs); err != nil {
					log.Printf("idempotency put failed: %v", err)
				}
			}
			return
		case "PAGE":
			payURL, err := c.TradePagePay(r.Context(), bm)
			if err != nil {
				writeJSON(w, http.StatusBadGateway, CreatePaymentResponse{Code: "UPSTREAM_ERROR", Message: err.Error()})
				return
			}
			resp := CreatePaymentResponse{Code: "OK", OutTradeNo: req.OutTradeNo, Status: "PAYING", PayData: map[string]string{"payUrl": payURL}}
			bs := writeJSONBytes(w, http.StatusOK, resp)
			if bs != nil {
				if err := s.idem.Put(r.Context(), idemKey, http.StatusOK, bs); err != nil {
					log.Printf("idempotency put failed: %v", err)
				}
			}
			return
		case "PRECREATE":
			aliRsp, err := c.TradePrecreate(r.Context(), bm)
			if err != nil {
				writeJSON(w, http.StatusBadGateway, CreatePaymentResponse{Code: "UPSTREAM_ERROR", Message: err.Error()})
				return
			}
			resp := CreatePaymentResponse{Code: "OK", OutTradeNo: req.OutTradeNo, Status: "PAYING", PayData: aliRsp.Response}
			bs := writeJSONBytes(w, http.StatusOK, resp)
			if bs != nil {
				if err := s.idem.Put(r.Context(), idemKey, http.StatusOK, bs); err != nil {
					log.Printf("idempotency put failed: %v", err)
				}
			}
			return
		default:
			writeJSON(w, http.StatusBadRequest, CreatePaymentResponse{Code: "BAD_REQUEST", Message: "unsupported scene for ALIPAY"})
			return
		}
	default:
		writeJSON(w, http.StatusBadRequest, CreatePaymentResponse{Code: "BAD_REQUEST", Message: "unsupported channel"})
		return
	}
}

func (s *Server) handleQueryPayment(w http.ResponseWriter, r *http.Request, outTradeNo string) {
	tenantID := r.URL.Query().Get("tenantId")
	merchantID := r.URL.Query().Get("merchantId")
	channel := Channel(r.URL.Query().Get("channel"))
	if tenantID == "" {
		tenantID = s.cfg.DefaultTenantID
	}
	if tenantID == "" || merchantID == "" || channel == "" {
		writeJSON(w, http.StatusBadRequest, QueryPaymentResponse{Code: "BAD_REQUEST", Message: "missing tenantId/merchantId/channel"})
		return
	}

	switch channel {
	case ChannelWechatV3:
		c, _, err := s.clients.WechatV3(tenantID, merchantID)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, QueryPaymentResponse{Code: "BAD_REQUEST", Message: err.Error()})
			return
		}
		wxRsp, err := c.V3TransactionQueryOrder(r.Context(), wechatv3.OutTradeNo, outTradeNo)
		if err != nil {
			writeJSON(w, http.StatusBadGateway, QueryPaymentResponse{Code: "UPSTREAM_ERROR", Message: err.Error()})
			return
		}
		if wxRsp.Code != wechatv3.Success {
			writeJSON(w, http.StatusBadGateway, QueryPaymentResponse{Code: "UPSTREAM_ERROR", Message: wxRsp.Error, Data: wxRsp.ErrResponse})
			return
		}
		writeJSON(w, http.StatusOK, QueryPaymentResponse{Code: "OK", Data: wxRsp.Response})
		return
	case ChannelAlipay:
		c, _, err := s.clients.Alipay(tenantID, merchantID)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, QueryPaymentResponse{Code: "BAD_REQUEST", Message: err.Error()})
			return
		}
		bm := make(gopay.BodyMap)
		bm.Set("out_trade_no", outTradeNo)
		aliRsp, err := c.TradeQuery(r.Context(), bm)
		if err != nil {
			writeJSON(w, http.StatusBadGateway, QueryPaymentResponse{Code: "UPSTREAM_ERROR", Message: err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, QueryPaymentResponse{Code: "OK", Data: aliRsp.Response})
		return
	default:
		writeJSON(w, http.StatusBadRequest, QueryPaymentResponse{Code: "BAD_REQUEST", Message: "unsupported channel"})
		return
	}
}

func (s *Server) handleClosePayment(w http.ResponseWriter, r *http.Request, outTradeNo string) {
	var body struct {
		TenantID   string  `json:"tenantId"`
		MerchantID string  `json:"merchantId"`
		Channel    Channel `json:"channel"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeJSON(w, http.StatusBadRequest, ClosePaymentResponse{Code: "BAD_REQUEST", Message: err.Error()})
		return
	}
	if body.TenantID == "" {
		body.TenantID = s.cfg.DefaultTenantID
	}
	if body.TenantID == "" || body.MerchantID == "" || body.Channel == "" {
		writeJSON(w, http.StatusBadRequest, ClosePaymentResponse{Code: "BAD_REQUEST", Message: "missing tenantId/merchantId/channel"})
		return
	}
	switch body.Channel {
	case ChannelWechatV3:
		c, _, err := s.clients.WechatV3(body.TenantID, body.MerchantID)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, ClosePaymentResponse{Code: "BAD_REQUEST", Message: err.Error()})
			return
		}
		wxRsp, err := c.V3TransactionCloseOrder(r.Context(), outTradeNo)
		if err != nil {
			writeJSON(w, http.StatusBadGateway, ClosePaymentResponse{Code: "UPSTREAM_ERROR", Message: err.Error()})
			return
		}
		if wxRsp.Code != wechatv3.Success {
			writeJSON(w, http.StatusBadGateway, ClosePaymentResponse{Code: "UPSTREAM_ERROR", Message: wxRsp.Error})
			return
		}
		writeJSON(w, http.StatusOK, ClosePaymentResponse{Code: "OK"})
		return
	case ChannelAlipay:
		c, _, err := s.clients.Alipay(body.TenantID, body.MerchantID)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, ClosePaymentResponse{Code: "BAD_REQUEST", Message: err.Error()})
			return
		}
		bm := make(gopay.BodyMap)
		bm.Set("out_trade_no", outTradeNo)
		aliRsp, err := c.TradeClose(r.Context(), bm)
		if err != nil {
			writeJSON(w, http.StatusBadGateway, ClosePaymentResponse{Code: "UPSTREAM_ERROR", Message: err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, ClosePaymentResponse{Code: "OK", Message: aliRsp.Response.Msg})
		return
	default:
		writeJSON(w, http.StatusBadRequest, ClosePaymentResponse{Code: "BAD_REQUEST", Message: "unsupported channel"})
		return
	}
}

// =================================================
// Refund APIs

func (s *Server) handleCreateRefund(w http.ResponseWriter, r *http.Request) {
	var req CreateRefundRequest
	if err := decodeJSON(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, CreateRefundResponse{Code: "BAD_REQUEST", Message: err.Error()})
		return
	}
	if req.TenantID == "" {
		req.TenantID = s.cfg.DefaultTenantID
	}
	if req.TenantID == "" || req.MerchantID == "" || req.OutTradeNo == "" || req.OutRefundNo == "" || req.Currency == "" || req.RefundAmount <= 0 {
		writeJSON(w, http.StatusBadRequest, CreateRefundResponse{Code: "BAD_REQUEST", Message: "missing required fields"})
		return
	}

	idemKey := r.Header.Get("X-Idempotency-Key")
	if idemKey == "" {
		idemKey = fmt.Sprintf("refund:%s:%s:%s:%s", req.TenantID, req.MerchantID, req.Channel, req.OutRefundNo)
	}
	if status, body, ok, err := s.idem.Get(r.Context(), idemKey); err != nil {
		writeJSON(w, http.StatusInternalServerError, CreateRefundResponse{Code: "ERROR", Message: "idempotency store error"})
		return
	} else if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write(body)
		return
	}

	switch req.Channel {
	case ChannelWechatV3:
		if req.TotalAmount <= 0 {
			writeJSON(w, http.StatusBadRequest, CreateRefundResponse{Code: "BAD_REQUEST", Message: "missing totalAmount for WECHAT_V3 refund"})
			return
		}
		if req.TotalAmount < req.RefundAmount {
			writeJSON(w, http.StatusBadRequest, CreateRefundResponse{Code: "BAD_REQUEST", Message: "totalAmount must be >= refundAmount"})
			return
		}
		c, _, err := s.clients.WechatV3(req.TenantID, req.MerchantID)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, CreateRefundResponse{Code: "BAD_REQUEST", Message: err.Error()})
			return
		}
		notifyURL, err := s.callbackURL(string(ChannelWechatV3), req.TenantID, req.MerchantID)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, CreateRefundResponse{Code: "ERROR", Message: err.Error()})
			return
		}
		bm := make(gopay.BodyMap)
		bm.Set("out_trade_no", req.OutTradeNo).
			Set("out_refund_no", req.OutRefundNo).
			Set("notify_url", notifyURL).
			SetBodyMap("amount", func(b gopay.BodyMap) {
				b.Set("refund", req.RefundAmount).
					Set("total", req.TotalAmount).
					Set("currency", req.Currency)
			})
		if req.Reason != "" {
			bm.Set("reason", req.Reason)
		}
		wxRsp, err := c.V3Refund(r.Context(), bm)
		if err != nil {
			writeJSON(w, http.StatusBadGateway, CreateRefundResponse{Code: "UPSTREAM_ERROR", Message: err.Error()})
			return
		}
		if wxRsp.Code != wechatv3.Success {
			writeJSON(w, http.StatusBadGateway, CreateRefundResponse{Code: "UPSTREAM_ERROR", Message: wxRsp.Error, OutRefundNo: req.OutRefundNo, Data: wxRsp.ErrResponse})
			return
		}
		resp := CreateRefundResponse{Code: "OK", OutRefundNo: req.OutRefundNo, Status: "REFUNDING", Data: wxRsp.Response}
		bs := writeJSONBytes(w, http.StatusOK, resp)
		if bs != nil {
			if err := s.idem.Put(r.Context(), idemKey, http.StatusOK, bs); err != nil {
				log.Printf("idempotency put failed: %v", err)
			}
		}
		return

	case ChannelAlipay:
		if strings.ToUpper(req.Currency) != "CNY" {
			writeJSON(w, http.StatusBadRequest, CreateRefundResponse{Code: "BAD_REQUEST", Message: "ALIPAY only supports CNY in this gateway"})
			return
		}
		c, _, err := s.clients.Alipay(req.TenantID, req.MerchantID)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, CreateRefundResponse{Code: "BAD_REQUEST", Message: err.Error()})
			return
		}
		refundAmount, err := formatCNYFenToYuanString(req.RefundAmount)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, CreateRefundResponse{Code: "BAD_REQUEST", Message: err.Error()})
			return
		}
		bm := make(gopay.BodyMap)
		bm.Set("out_trade_no", req.OutTradeNo).
			Set("refund_amount", refundAmount).
			Set("out_request_no", req.OutRefundNo)
		if req.Reason != "" {
			bm.Set("refund_reason", req.Reason)
		}
		aliRsp, err := c.TradeRefund(r.Context(), bm)
		if err != nil {
			writeJSON(w, http.StatusBadGateway, CreateRefundResponse{Code: "UPSTREAM_ERROR", Message: err.Error()})
			return
		}
		resp := CreateRefundResponse{Code: "OK", OutRefundNo: req.OutRefundNo, Status: "REFUNDING", Data: aliRsp.Response}
		bs := writeJSONBytes(w, http.StatusOK, resp)
		if bs != nil {
			if err := s.idem.Put(r.Context(), idemKey, http.StatusOK, bs); err != nil {
				log.Printf("idempotency put failed: %v", err)
			}
		}
		return

	default:
		writeJSON(w, http.StatusBadRequest, CreateRefundResponse{Code: "BAD_REQUEST", Message: "unsupported channel"})
		return
	}
}

func (s *Server) handleQueryRefund(w http.ResponseWriter, r *http.Request, outRefundNo string) {
	tenantID := r.URL.Query().Get("tenantId")
	merchantID := r.URL.Query().Get("merchantId")
	channel := Channel(r.URL.Query().Get("channel"))
	if tenantID == "" {
		tenantID = s.cfg.DefaultTenantID
	}
	if tenantID == "" || merchantID == "" || channel == "" {
		writeJSON(w, http.StatusBadRequest, QueryRefundResponse{Code: "BAD_REQUEST", Message: "missing tenantId/merchantId/channel"})
		return
	}

	switch channel {
	case ChannelWechatV3:
		c, _, err := s.clients.WechatV3(tenantID, merchantID)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, QueryRefundResponse{Code: "BAD_REQUEST", Message: err.Error()})
			return
		}
		wxRsp, err := c.V3RefundQuery(r.Context(), outRefundNo, nil)
		if err != nil {
			writeJSON(w, http.StatusBadGateway, QueryRefundResponse{Code: "UPSTREAM_ERROR", Message: err.Error()})
			return
		}
		if wxRsp.Code != wechatv3.Success {
			writeJSON(w, http.StatusBadGateway, QueryRefundResponse{Code: "UPSTREAM_ERROR", Message: wxRsp.Error, Data: wxRsp.ErrResponse})
			return
		}
		writeJSON(w, http.StatusOK, QueryRefundResponse{Code: "OK", Data: wxRsp.Response})
		return

	case ChannelAlipay:
		outTradeNo := r.URL.Query().Get("outTradeNo")
		tradeNo := r.URL.Query().Get("tradeNo")
		if outTradeNo == "" && tradeNo == "" {
			writeJSON(w, http.StatusBadRequest, QueryRefundResponse{Code: "BAD_REQUEST", Message: "missing outTradeNo or tradeNo for ALIPAY refund query"})
			return
		}
		c, _, err := s.clients.Alipay(tenantID, merchantID)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, QueryRefundResponse{Code: "BAD_REQUEST", Message: err.Error()})
			return
		}
		bm := make(gopay.BodyMap)
		if outTradeNo != "" {
			bm.Set("out_trade_no", outTradeNo)
		}
		if tradeNo != "" {
			bm.Set("trade_no", tradeNo)
		}
		bm.Set("out_request_no", outRefundNo)
		aliRsp, err := c.TradeFastPayRefundQuery(r.Context(), bm)
		if err != nil {
			writeJSON(w, http.StatusBadGateway, QueryRefundResponse{Code: "UPSTREAM_ERROR", Message: err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, QueryRefundResponse{Code: "OK", Data: aliRsp.Response})
		return

	default:
		writeJSON(w, http.StatusBadRequest, QueryRefundResponse{Code: "BAD_REQUEST", Message: "unsupported channel"})
		return
	}
}

func (s *Server) handleCompensationQueryPayments(w http.ResponseWriter, r *http.Request) {
	var req CompensationQueryPaymentsRequest
	if err := decodeJSON(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, CompensationQueryPaymentsResponse{Code: "BAD_REQUEST", Message: err.Error()})
		return
	}
	if req.TenantID == "" {
		req.TenantID = s.cfg.DefaultTenantID
	}
	if req.TenantID == "" || req.MerchantID == "" || req.Channel == "" || len(req.OutTradeNos) == 0 {
		writeJSON(w, http.StatusBadRequest, CompensationQueryPaymentsResponse{Code: "BAD_REQUEST", Message: "missing required fields"})
		return
	}
	if len(req.OutTradeNos) > 50 {
		writeJSON(w, http.StatusBadRequest, CompensationQueryPaymentsResponse{Code: "BAD_REQUEST", Message: "outTradeNos too large (max 50)"})
		return
	}

	items := make([]CompensationPaymentItem, 0, len(req.OutTradeNos))
	for _, outTradeNo := range req.OutTradeNos {
		item := CompensationPaymentItem{OutTradeNo: outTradeNo}
		switch req.Channel {
		case ChannelWechatV3:
			c, _, err := s.clients.WechatV3(req.TenantID, req.MerchantID)
			if err != nil {
				item.Error = err.Error()
				break
			}
			wxRsp, err := c.V3TransactionQueryOrder(r.Context(), wechatv3.OutTradeNo, outTradeNo)
			if err != nil {
				item.Error = err.Error()
				break
			}
			if wxRsp.Code != wechatv3.Success {
				item.Error = wxRsp.Error
				item.Data = wxRsp.ErrResponse
				break
			}
			item.Data = wxRsp.Response
		case ChannelAlipay:
			c, _, err := s.clients.Alipay(req.TenantID, req.MerchantID)
			if err != nil {
				item.Error = err.Error()
				break
			}
			bm := make(gopay.BodyMap)
			bm.Set("out_trade_no", outTradeNo)
			aliRsp, err := c.TradeQuery(r.Context(), bm)
			if err != nil {
				item.Error = err.Error()
				break
			}
			item.Data = aliRsp.Response
		default:
			item.Error = "unsupported channel"
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, CompensationQueryPaymentsResponse{Code: "OK", Items: items})
}

// =================================================
// Callback handlers

func (s *Server) handleWechatV3Callback(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	tenantID, merchantID, ok := parse2Segments(strings.TrimPrefix(r.URL.Path, "/callbacks/wechat/v3/"))
	if !ok {
		http.NotFound(w, r)
		return
	}
	c, wc, err := s.clients.WechatV3(tenantID, merchantID)
	if err != nil {
		writeWechatV3NotifyResp(w, http.StatusBadRequest, "FAIL", err.Error())
		return
	}
	notifyReq, err := wechatv3.V3ParseNotify(r)
	if err != nil {
		writeWechatV3NotifyResp(w, http.StatusBadRequest, "FAIL", err.Error())
		return
	}
	if err := notifyReq.VerifySignByPKMap(c.WxPublicKeyMap()); err != nil {
		writeWechatV3NotifyResp(w, http.StatusBadRequest, "FAIL", "verify sign failed")
		return
	}

	var ev *Event
	if isWechatRefundNotify(notifyReq) {
		var decrypt wechatv3.V3DecryptRefundResult
		if err := notifyReq.DecryptCipherTextToStruct(wc.ApiV3Key, &decrypt); err != nil {
			writeWechatV3NotifyResp(w, http.StatusBadRequest, "FAIL", "decrypt failed")
			return
		}
		if decrypt.Amount == nil {
			writeWechatV3NotifyResp(w, http.StatusBadRequest, "FAIL", "missing amount")
			return
		}
		ev = &Event{
			EventID:           "WECHAT_V3:" + notifyReq.Id,
			EventType:         wechatV3RefundEventType(notifyReq.EventType, decrypt.RefundStatus),
			EventVersion:      1,
			OccurredAt:        time.Now().UTC().Format(time.RFC3339),
			TenantID:          tenantID,
			MerchantID:        merchantID,
			Channel:           string(ChannelWechatV3),
			OutTradeNo:        decrypt.OutTradeNo,
			TransactionID:     decrypt.TransactionId,
			OutRefundNo:       decrypt.OutRefundNo,
			RefundID:          decrypt.RefundId,
			Amount:            int64(decrypt.Amount.Refund),
			Currency:          "CNY",
			RefundState:       decrypt.RefundStatus,
			SignatureVerified: true,
			IdempotencyKey:    merchantKey(tenantID, merchantID) + ":" + decrypt.OutRefundNo,
		}
	} else {
		var decrypt wechatv3.V3DecryptPayResult
		if err := notifyReq.DecryptCipherTextToStruct(wc.ApiV3Key, &decrypt); err != nil {
			writeWechatV3NotifyResp(w, http.StatusBadRequest, "FAIL", "decrypt failed")
			return
		}
		if decrypt.Amount == nil {
			writeWechatV3NotifyResp(w, http.StatusBadRequest, "FAIL", "missing amount")
			return
		}
		ev = &Event{
			EventID:           "WECHAT_V3:" + notifyReq.Id,
			EventType:         wechatV3EventType(notifyReq.EventType, decrypt.TradeState),
			EventVersion:      1,
			OccurredAt:        time.Now().UTC().Format(time.RFC3339),
			TenantID:          tenantID,
			MerchantID:        merchantID,
			Channel:           string(ChannelWechatV3),
			OutTradeNo:        decrypt.OutTradeNo,
			TransactionID:     decrypt.TransactionId,
			Amount:            int64(decrypt.Amount.Total),
			Currency:          decrypt.Amount.Currency,
			TradeState:        decrypt.TradeState,
			SignatureVerified: true,
			IdempotencyKey:    merchantKey(tenantID, merchantID) + ":" + decrypt.OutTradeNo,
		}
	}
	dedupKey := fmt.Sprintf("%s:%s:%s:%s", ChannelWechatV3, tenantID, merchantID, notifyReq.Id)
	locked, state, err := s.dedup.TryLock(r.Context(), dedupKey)
	if err != nil {
		writeWechatV3NotifyResp(w, http.StatusInternalServerError, "FAIL", "dedup error")
		return
	}
	if !locked {
		if state == DedupStateDone {
			writeWechatV3NotifyResp(w, http.StatusOK, "SUCCESS", "ok")
			return
		}
		writeWechatV3NotifyResp(w, http.StatusInternalServerError, "FAIL", "processing")
		return
	}
	if err := s.publisher.Publish(r.Context(), ev); err != nil {
		_ = s.dedup.Release(r.Context(), dedupKey)
		writeWechatV3NotifyResp(w, http.StatusInternalServerError, "FAIL", "publish failed")
		return
	}
	if err := s.dedup.MarkDone(r.Context(), dedupKey); err != nil {
		log.Printf("dedup mark done failed: %v", err)
	}
	writeWechatV3NotifyResp(w, http.StatusOK, "SUCCESS", "ok")
}

func (s *Server) handleAlipayCallback(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	tenantID, merchantID, ok := parse2Segments(strings.TrimPrefix(r.URL.Path, "/callbacks/alipay/"))
	if !ok {
		http.NotFound(w, r)
		return
	}
	_, ac, err := s.clients.Alipay(tenantID, merchantID)
	if err != nil {
		http.Error(w, "failure", http.StatusBadRequest)
		return
	}
	bm, err := alipay.ParseNotifyToBodyMap(r)
	if err != nil {
		http.Error(w, "failure", http.StatusBadRequest)
		return
	}
	okSign, err := alipay.VerifySign(ac.AlipayPublicKey, bm)
	if err != nil || !okSign {
		http.Error(w, "failure", http.StatusBadRequest)
		return
	}

	tradeStatus := bm.GetString("trade_status")
	outTradeNo := bm.GetString("out_trade_no")
	tradeNo := bm.GetString("trade_no")
	totalAmount := bm.GetString("total_amount")
	notifyID := bm.GetString("notify_id")
	eventID := "ALIPAY:" + outTradeNo
	if notifyID != "" {
		eventID = "ALIPAY:" + notifyID
	}

	ev := &Event{
		EventID:           eventID,
		EventType:         alipayEventType(tradeStatus),
		EventVersion:      1,
		OccurredAt:        time.Now().UTC().Format(time.RFC3339),
		TenantID:          tenantID,
		MerchantID:        merchantID,
		Channel:           string(ChannelAlipay),
		OutTradeNo:        outTradeNo,
		TransactionID:     tradeNo,
		Currency:          "CNY",
		TradeState:        tradeStatus,
		SignatureVerified: true,
		IdempotencyKey:    merchantKey(tenantID, merchantID) + ":" + outTradeNo,
		Ext:               map[string]string{"total_amount": totalAmount},
	}
	dedupKey := fmt.Sprintf("%s:%s:%s:%s", ChannelAlipay, tenantID, merchantID, eventID)
	locked, state, err := s.dedup.TryLock(r.Context(), dedupKey)
	if err != nil {
		http.Error(w, "failure", http.StatusInternalServerError)
		return
	}
	if !locked {
		if state == DedupStateDone {
			_, _ = w.Write([]byte("success"))
			return
		}
		http.Error(w, "failure", http.StatusInternalServerError)
		return
	}
	if err := s.publisher.Publish(r.Context(), ev); err != nil {
		_ = s.dedup.Release(r.Context(), dedupKey)
		http.Error(w, "failure", http.StatusInternalServerError)
		return
	}
	if err := s.dedup.MarkDone(r.Context(), dedupKey); err != nil {
		log.Printf("dedup mark done failed: %v", err)
	}
	_, _ = w.Write([]byte("success"))
}

func writeWechatV3NotifyResp(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, wechatv3.V3NotifyRsp{Code: code, Message: message})
}

func parse2Segments(rest string) (a, b string, ok bool) {
	parts := strings.Split(strings.Trim(rest, "/"), "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", false
	}
	return parts[0], parts[1], true
}

func firstNonEmpty(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

func wechatV3EventType(eventType, tradeState string) string {
	switch tradeState {
	case "SUCCESS":
		return "payment.succeeded"
	case "CLOSED":
		return "payment.closed"
	default:
		if eventType != "" {
			return "wechat." + strings.ToLower(eventType)
		}
		return "payment.updated"
	}
}

func wechatV3RefundEventType(eventType, refundStatus string) string {
	switch refundStatus {
	case "SUCCESS":
		return "refund.succeeded"
	case "CLOSED":
		return "refund.closed"
	case "ABNORMAL":
		return "refund.failed"
	default:
		if eventType != "" {
			return "wechat." + strings.ToLower(eventType)
		}
		return "refund.updated"
	}
}

func isWechatRefundNotify(req *wechatv3.V3NotifyReq) bool {
	if req == nil {
		return false
	}
	if strings.HasPrefix(strings.ToUpper(req.EventType), "REFUND.") {
		return true
	}
	if req.Resource != nil && strings.EqualFold(req.Resource.OriginalType, "refund") {
		return true
	}
	return false
}

func alipayEventType(tradeStatus string) string {
	switch tradeStatus {
	case "TRADE_SUCCESS", "TRADE_FINISHED":
		return "payment.succeeded"
	case "TRADE_CLOSED":
		return "payment.closed"
	default:
		return "payment.updated"
	}
}

// EventPublisherFunc adapts a function to EventPublisher.
type EventPublisherFunc func(ctx context.Context, event *Event) error

func (f EventPublisherFunc) Publish(ctx context.Context, event *Event) error { return f(ctx, event) }
