package paygateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type MerchantSyncWorkerOptions struct {
	SnapshotURL     string
	PollInterval    time.Duration
	ChangeStream    string
	ConsumerGroup   string
	Redis           *redis.Client
	HTTPClient      *http.Client
	SharedAuth      SharedAuthConfig
	DefaultTenantID string
	Store           *MerchantStore
	Clients         *ClientManager
}

type MerchantSyncWorker struct {
	snapshotURL  string
	pollInterval time.Duration
	changeStream string
	group        string

	rdb    *redis.Client
	hc     *http.Client
	shared SharedAuthConfig

	store   *MerchantStore
	clients *ClientManager

	cancel context.CancelFunc
	done   chan struct{}
}

func NewMerchantSyncWorker(opts MerchantSyncWorkerOptions) *MerchantSyncWorker {
	if opts.PollInterval <= 0 {
		opts.PollInterval = 60 * time.Second
	}
	if opts.ConsumerGroup == "" {
		opts.ConsumerGroup = "pay-gateway-config"
	}
	return &MerchantSyncWorker{
		snapshotURL:  opts.SnapshotURL,
		pollInterval: opts.PollInterval,
		changeStream: opts.ChangeStream,
		group:        opts.ConsumerGroup,
		rdb:          opts.Redis,
		hc:           opts.HTTPClient,
		shared:       opts.SharedAuth,
		store:        opts.Store,
		clients:      opts.Clients,
	}
}

func (w *MerchantSyncWorker) Start() error {
	if w == nil || w.done != nil || w.snapshotURL == "" {
		return nil
	}
	if w.hc == nil {
		return errors.New("merchant sync http client is nil")
	}
	if w.store == nil || w.clients == nil {
		return errors.New("merchant sync store/clients is nil")
	}

	ctx, cancel := context.WithCancel(context.Background())
	w.cancel = cancel
	w.done = make(chan struct{})
	go w.loop(ctx)
	return nil
}

func (w *MerchantSyncWorker) Stop() {
	if w == nil || w.cancel == nil || w.done == nil {
		return
	}
	w.cancel()
	<-w.done
}

func (w *MerchantSyncWorker) loop(ctx context.Context) {
	defer close(w.done)

	// Best-effort initial sync.
	_ = w.syncOnce(ctx)

	ticker := time.NewTicker(w.pollInterval)
	defer ticker.Stop()

	if w.rdb != nil && w.changeStream != "" {
		go w.consumeChangeStream(ctx)
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			_ = w.syncOnce(ctx)
		}
	}
}

func (w *MerchantSyncWorker) syncOnce(ctx context.Context) error {
	snap, err := w.fetchSnapshot(ctx)
	if err != nil {
		log.Printf("merchant sync fetch failed: %v", err)
		return err
	}
	if snap.Version != "" && snap.Version == w.store.Version() {
		return nil
	}
	keys, err := w.store.Replace(snap.Merchants, snap.Version)
	if err != nil {
		log.Printf("merchant sync apply failed: %v", err)
		return err
	}
	if len(keys) > 0 {
		w.clients.InvalidateKeys(keys)
	}
	log.Printf("merchant sync ok version=%s changed=%d", snap.Version, len(keys))
	return nil
}

type merchantSnapshot struct {
	Version   string           `json:"version"`
	Merchants []MerchantConfig `json:"merchants"`
}

func (w *MerchantSyncWorker) fetchSnapshot(ctx context.Context) (*merchantSnapshot, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, w.snapshotURL, nil)
	if err != nil {
		return nil, err
	}
	if w.shared.SharedSecret != "" {
		if err := signHTTPRequest(req, nil, w.shared.SharedSecret); err != nil {
			return nil, err
		}
	}
	res, err := w.hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode/100 != 2 {
		return nil, fmt.Errorf("snapshot http status=%d", res.StatusCode)
	}
	var snap merchantSnapshot
	dec := json.NewDecoder(res.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&snap); err != nil {
		return nil, err
	}
	return &snap, nil
}

func (w *MerchantSyncWorker) consumeChangeStream(ctx context.Context) {
	// Ensure consumer group exists.
	ctx2, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := w.rdb.XGroupCreateMkStream(ctx2, w.changeStream, w.group, "0").Err(); err != nil && !strings.Contains(err.Error(), "BUSYGROUP") {
		log.Printf("merchant sync create group failed: %v", err)
		return
	}
	consumer := "pay-gateway"
	if host, _ := os.Hostname(); host != "" {
		consumer = host
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		streams, err := w.rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    w.group,
			Consumer: consumer,
			Streams:  []string{w.changeStream, ">"},
			Count:    16,
			Block:    5 * time.Second,
		}).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				continue
			}
			continue
		}
		for _, s := range streams {
			for _, msg := range s.Messages {
				if err := w.syncOnce(ctx); err != nil {
					continue
				}
				_ = w.rdb.XAck(ctx, w.changeStream, w.group, msg.ID).Err()
			}
		}
	}
}
