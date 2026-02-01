package paygateway

import (
	"fmt"
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
		key := merchantKey(m.TenantID, m.MerchantID)
		if _, exists := ms.byKey[key]; exists {
			return nil, fmt.Errorf("duplicate merchant: %s", key)
		}
		ms.byKey[key] = m
	}
	return ms, nil
}

func (s *MerchantStore) Get(tenantID, merchantID string) (*MerchantConfig, bool) {
	m, ok := s.byKey[merchantKey(tenantID, merchantID)]
	return m, ok
}

func merchantKey(tenantID, merchantID string) string {
	return tenantID + ":" + merchantID
}
