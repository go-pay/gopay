package paygateway

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Server        ServerConfig      `json:"server"`
	PublicBaseURL string            `json:"publicBaseUrl"`
	JavaWebhook   JavaWebhookConfig `json:"javaWebhook"`
	TLS           TLSConfig         `json:"tls"`
	Merchants     []MerchantConfig  `json:"merchants"`
}

type ServerConfig struct {
	Addr                string `json:"addr"`
	ReadTimeoutSeconds  int    `json:"readTimeoutSeconds"`
	WriteTimeoutSeconds int    `json:"writeTimeoutSeconds"`
}

type TLSConfig struct {
	CAFile string `json:"caFile"`
}

type JavaWebhookConfig struct {
	URL           string `json:"url"`
	Token         string `json:"token"`
	TimeoutMillis int    `json:"timeoutMillis"`
}

type MerchantConfig struct {
	TenantID   string          `json:"tenantId"`
	MerchantID string          `json:"merchantId"`
	WechatV3   *WechatV3Config `json:"wechatV3,omitempty"`
	Alipay     *AlipayConfig   `json:"alipay,omitempty"`
}

type WechatV3Config struct {
	AppID      string `json:"appId"`
	MchID      string `json:"mchId"`
	SerialNo   string `json:"serialNo"`
	ApiV3Key   string `json:"apiV3Key"`
	PrivateKey string `json:"privateKey"`

	// Optional: avoid network fetching platform certs (verify-only mode).
	WechatPayPublicKey   string `json:"wechatPayPublicKey,omitempty"`
	WechatPayPublicKeyID string `json:"wechatPayPublicKeyId,omitempty"`
}

type AlipayConfig struct {
	IsProd          bool   `json:"isProd"`
	AppID           string `json:"appId"`
	PrivateKey      string `json:"privateKey"`
	AlipayPublicKey string `json:"alipayPublicKey"`
}

func LoadConfig(path string) (*Config, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	dec := json.NewDecoder(bytesReader(bs))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("decode %s: %w", path, err)
	}
	if cfg.Server.Addr == "" {
		cfg.Server.Addr = ":8088"
	}
	if cfg.JavaWebhook.TimeoutMillis == 0 {
		cfg.JavaWebhook.TimeoutMillis = 1500
	}
	return &cfg, nil
}
