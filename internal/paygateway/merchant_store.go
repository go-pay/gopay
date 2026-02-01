package paygateway

import (
	"fmt"
	"os"
	"strings"
)

type MerchantStore struct {
	byKey map[string]*MerchantConfig
}

func NewMerchantStore(cfg *Config) (*MerchantStore, error) {
	ms := &MerchantStore{byKey: make(map[string]*MerchantConfig)}
	for i := range cfg.Merchants {
		m := &cfg.Merchants[i]
		if m.TenantID == "" || m.MerchantID == "" {
			return nil, fmt.Errorf("merchant missing tenantId/merchantId")
		}
		if err := resolveMerchantSecretFiles(m); err != nil {
			return nil, err
		}
		key := merchantKey(m.TenantID, m.MerchantID)
		if _, exists := ms.byKey[key]; exists {
			return nil, fmt.Errorf("duplicate merchant: %s", key)
		}
		ms.byKey[key] = m
	}
	return ms, nil
}

func resolveMerchantSecretFiles(m *MerchantConfig) error {
	if m == nil {
		return nil
	}
	if m.WechatV3 != nil {
		if m.WechatV3.PrivateKey == "" && m.WechatV3.PrivateKeyFile != "" {
			bs, err := os.ReadFile(m.WechatV3.PrivateKeyFile)
			if err != nil {
				return fmt.Errorf("read wechatV3.privateKeyFile: %w", err)
			}
			m.WechatV3.PrivateKey = strings.TrimSpace(string(bs))
		}
		if m.WechatV3.WechatPayPublicKey == "" && m.WechatV3.WechatPayPublicKeyFile != "" {
			bs, err := os.ReadFile(m.WechatV3.WechatPayPublicKeyFile)
			if err != nil {
				return fmt.Errorf("read wechatV3.wechatPayPublicKeyFile: %w", err)
			}
			m.WechatV3.WechatPayPublicKey = strings.TrimSpace(string(bs))
		}
	}
	if m.Alipay != nil {
		if m.Alipay.PrivateKey == "" && m.Alipay.PrivateKeyFile != "" {
			bs, err := os.ReadFile(m.Alipay.PrivateKeyFile)
			if err != nil {
				return fmt.Errorf("read alipay.privateKeyFile: %w", err)
			}
			m.Alipay.PrivateKey = strings.TrimSpace(string(bs))
		}
		if m.Alipay.AlipayPublicKey == "" && m.Alipay.AlipayPublicKeyFile != "" {
			bs, err := os.ReadFile(m.Alipay.AlipayPublicKeyFile)
			if err != nil {
				return fmt.Errorf("read alipay.alipayPublicKeyFile: %w", err)
			}
			m.Alipay.AlipayPublicKey = strings.TrimSpace(string(bs))
		}
	}
	return nil
}

func (s *MerchantStore) Get(tenantID, merchantID string) (*MerchantConfig, bool) {
	m, ok := s.byKey[merchantKey(tenantID, merchantID)]
	return m, ok
}

func merchantKey(tenantID, merchantID string) string {
	return tenantID + ":" + merchantID
}
