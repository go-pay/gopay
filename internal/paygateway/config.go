package paygateway

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Server        ServerConfig      `json:"server"`
	PublicBaseURL string            `json:"publicBaseUrl"`
	APIAuth       APIAuthConfig     `json:"apiAuth"`
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

type APIAuthConfig struct {
	Token string `json:"token"`
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
	if err := applyEnvOverrides(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func applyEnvOverrides(cfg *Config) error {
	if cfg == nil {
		return nil
	}
	if v, ok := envString("PAY_GATEWAY_ADDR"); ok {
		cfg.Server.Addr = v
	}
	if v, ok := envString("PAY_GATEWAY_PUBLIC_BASE_URL"); ok {
		cfg.PublicBaseURL = v
	}
	if v, ok := envString("PAY_GATEWAY_API_AUTH_TOKEN"); ok {
		cfg.APIAuth.Token = v
	}
	if v, ok := envString("PAY_GATEWAY_JAVA_WEBHOOK_URL"); ok {
		cfg.JavaWebhook.URL = v
	}
	if v, ok := envString("PAY_GATEWAY_JAVA_WEBHOOK_TOKEN"); ok {
		cfg.JavaWebhook.Token = v
	}
	if v, ok := envInt("PAY_GATEWAY_JAVA_WEBHOOK_TIMEOUT_MILLIS"); ok {
		cfg.JavaWebhook.TimeoutMillis = v
	}
	if v, ok := envString("PAY_GATEWAY_TLS_CA_FILE"); ok {
		cfg.TLS.CAFile = v
	}
	return nil
}

func envString(key string) (string, bool) {
	v, ok := os.LookupEnv(key)
	if !ok {
		return "", false
	}
	v = strings.TrimSpace(v)
	if v == "" {
		return "", false
	}
	return v, true
}

func envInt(key string) (int, bool) {
	v, ok := envString(key)
	if !ok {
		return 0, false
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, false
	}
	return n, true
}
