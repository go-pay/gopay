package paygateway

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Event struct {
	EventID           string            `json:"eventId"`
	EventType         string            `json:"eventType"`
	EventVersion      int               `json:"eventVersion"`
	OccurredAt        string            `json:"occurredAt"`
	TenantID          string            `json:"tenantId"`
	MerchantID        string            `json:"merchantId"`
	Channel           string            `json:"channel"`
	OutTradeNo        string            `json:"outTradeNo,omitempty"`
	TransactionID     string            `json:"transactionId,omitempty"`
	OutRefundNo       string            `json:"outRefundNo,omitempty"`
	RefundID          string            `json:"refundId,omitempty"`
	Amount            int64             `json:"amount,omitempty"`
	Currency          string            `json:"currency,omitempty"`
	TradeState        string            `json:"tradeState,omitempty"`
	RefundState       string            `json:"refundState,omitempty"`
	SignatureVerified bool              `json:"signatureVerified"`
	IdempotencyKey    string            `json:"idempotencyKey,omitempty"`
	Ext               map[string]string `json:"ext,omitempty"`
}

type EventPublisher interface {
	Publish(ctx context.Context, event *Event) error
}

type WebhookPublisher struct {
	url     string
	token   string
	client  *http.Client
	timeout time.Duration
}

func NewWebhookPublisher(url, token string, client *http.Client, timeout time.Duration) *WebhookPublisher {
	return &WebhookPublisher{url: url, token: token, client: client, timeout: timeout}
}

func (p *WebhookPublisher) Publish(ctx context.Context, event *Event) error {
	if p.url == "" {
		return errors.New("webhook url is empty")
	}
	bs, err := json.Marshal(event)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.url, bytes.NewReader(bs))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if p.token != "" {
		req.Header.Set("X-Pay-Gateway-Token", p.token)
	}

	ctx2, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()
	req = req.WithContext(ctx2)

	res, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode/100 != 2 {
		return fmt.Errorf("webhook status=%d", res.StatusCode)
	}
	return nil
}
