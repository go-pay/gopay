package paygateway

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	HeaderPayGatewayTimestamp  = "X-Pay-Timestamp"
	HeaderPayGatewayNonce      = "X-Pay-Nonce"
	HeaderPayGatewayBodySHA256 = "X-Pay-Body-SHA256"
	HeaderPayGatewaySignature  = "X-Pay-Signature"

	HeaderPayGatewayTimestampLegacy  = "X-Pay-Gateway-Timestamp"
	HeaderPayGatewayNonceLegacy      = "X-Pay-Gateway-Nonce"
	HeaderPayGatewayBodySHA256Legacy = "X-Pay-Gateway-Body-SHA256"
	HeaderPayGatewaySignatureLegacy  = "X-Pay-Gateway-Signature"

	HeaderPayToken       = "X-Pay-Token"
	HeaderPayTokenLegacy = "X-Pay-Gateway-Token"
)

func signHeaders(method, requestURI, ts, nonce, bodySHA string, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write([]byte(method))
	_, _ = mac.Write([]byte("\n"))
	_, _ = mac.Write([]byte(requestURI))
	_, _ = mac.Write([]byte("\n"))
	_, _ = mac.Write([]byte(ts))
	_, _ = mac.Write([]byte("\n"))
	_, _ = mac.Write([]byte(nonce))
	_, _ = mac.Write([]byte("\n"))
	_, _ = mac.Write([]byte(bodySHA))
	_, _ = mac.Write([]byte("\n"))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func computeBodySHA256(bs []byte) string {
	sum := sha256.Sum256(bs)
	return base64.StdEncoding.EncodeToString(sum[:])
}

func newNonce() (string, error) {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b[:]), nil
}

func parseUnixSeconds(ts string) (time.Time, error) {
	sec, err := strconv.ParseInt(strings.TrimSpace(ts), 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(sec, 0).UTC(), nil
}

func verifyHMACRequest(r *http.Request, cfg SharedAuthConfig, nonceStore NonceStore) (body []byte, err error) {
	if cfg.SharedSecret == "" && cfg.SharedSecretPrev == "" {
		return nil, errors.New("shared secret not configured")
	}
	ts := firstNonEmpty(r.Header.Get(HeaderPayGatewayTimestamp), r.Header.Get(HeaderPayGatewayTimestampLegacy))
	nonce := firstNonEmpty(r.Header.Get(HeaderPayGatewayNonce), r.Header.Get(HeaderPayGatewayNonceLegacy))
	bodySHA := firstNonEmpty(r.Header.Get(HeaderPayGatewayBodySHA256), r.Header.Get(HeaderPayGatewayBodySHA256Legacy))
	sig := firstNonEmpty(r.Header.Get(HeaderPayGatewaySignature), r.Header.Get(HeaderPayGatewaySignatureLegacy))
	if ts == "" || nonce == "" || bodySHA == "" || sig == "" {
		return nil, errors.New("missing hmac headers")
	}
	tm, err := parseUnixSeconds(ts)
	if err != nil {
		return nil, errors.New("invalid timestamp")
	}
	skew := time.Duration(max(1, cfg.ClockSkewSeconds)) * time.Second
	if time.Since(tm) > skew || tm.Sub(time.Now().UTC()) > skew {
		return nil, errors.New("timestamp skew too large")
	}

	body, err = readAndReplaceBody(r, 5<<20)
	if err != nil {
		return nil, err
	}
	if computeBodySHA256(body) != bodySHA {
		return nil, errors.New("body sha mismatch")
	}

	if nonceStore != nil {
		ok, err := nonceStore.UseOnce(r.Context(), nonce, time.Duration(max(1, cfg.NonceTTLSeconds))*time.Second)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("replayed nonce")
		}
	}

	method := strings.ToUpper(r.Method)
	requestURI := r.URL.RequestURI()
	if requestURI == "" {
		requestURI = r.URL.Path
	}

	expected1 := ""
	if cfg.SharedSecret != "" {
		expected1 = signHeaders(method, requestURI, ts, nonce, bodySHA, cfg.SharedSecret)
	}
	expected2 := ""
	if cfg.SharedSecretPrev != "" {
		expected2 = signHeaders(method, requestURI, ts, nonce, bodySHA, cfg.SharedSecretPrev)
	}
	if hmac.Equal([]byte(sig), []byte(expected1)) || (expected2 != "" && hmac.Equal([]byte(sig), []byte(expected2))) {
		return body, nil
	}
	return nil, errors.New("invalid signature")
}

func signHTTPRequest(r *http.Request, body []byte, secret string) error {
	if secret == "" {
		return nil
	}
	ts := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	nonce, err := newNonce()
	if err != nil {
		return err
	}
	bodySHA := computeBodySHA256(body)
	sig := signHeaders(strings.ToUpper(r.Method), r.URL.RequestURI(), ts, nonce, bodySHA, secret)

	r.Header.Set(HeaderPayGatewayTimestamp, ts)
	r.Header.Set(HeaderPayGatewayNonce, nonce)
	r.Header.Set(HeaderPayGatewayBodySHA256, bodySHA)
	r.Header.Set(HeaderPayGatewaySignature, sig)
	return nil
}

func readAndReplaceBody(r *http.Request, limit int64) ([]byte, error) {
	if r.Body == nil {
		return nil, nil
	}
	bs, err := io.ReadAll(io.LimitReader(r.Body, limit))
	_ = r.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}
	r.Body = io.NopCloser(bytes.NewReader(bs))
	return bs, nil
}
