package paygateway

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
)

type MerchantStore struct {
	mu              sync.RWMutex
	byKey           map[string]*MerchantConfig
	hashByKey       map[string]string
	version         string
	defaultTenantID string
}

func NewMerchantStore(cfg *Config) (*MerchantStore, error) {
	ms := &MerchantStore{
		byKey:           make(map[string]*MerchantConfig),
		hashByKey:       make(map[string]string),
		defaultTenantID: cfg.DefaultTenantID,
	}
	if _, err := ms.Replace(cfg.Merchants, "bootstrap"); err != nil {
		return nil, err
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
	if tenantID == "" {
		tenantID = s.defaultTenantID
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	m, ok := s.byKey[merchantKey(tenantID, merchantID)]
	return m, ok
}

func (s *MerchantStore) Version() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.version
}

// Replace atomically replaces the merchant config set.
// Returns keys (tenant:merchant) that were added/removed/changed.
func (s *MerchantStore) Replace(merchants []MerchantConfig, version string) ([]string, error) {
	nextByKey := make(map[string]*MerchantConfig, len(merchants))
	nextHash := make(map[string]string, len(merchants))
	changed := make([]string, 0)

	for i := range merchants {
		mc := merchants[i] // copy
		if mc.TenantID == "" {
			mc.TenantID = s.defaultTenantID
		}
		if mc.TenantID == "" || mc.MerchantID == "" {
			return nil, fmt.Errorf("merchant missing tenantId/merchantId")
		}
		if err := resolveMerchantSecretFiles(&mc); err != nil {
			return nil, err
		}
		key := merchantKey(mc.TenantID, mc.MerchantID)
		if _, exists := nextByKey[key]; exists {
			return nil, fmt.Errorf("duplicate merchant: %s", key)
		}

		hash, err := hashMerchantConfig(&mc)
		if err != nil {
			return nil, err
		}
		nextByKey[key] = &mc
		nextHash[key] = hash
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for key, oldHash := range s.hashByKey {
		newHash, ok := nextHash[key]
		if !ok || newHash != oldHash {
			changed = append(changed, key)
		}
	}
	for key := range nextHash {
		if _, ok := s.hashByKey[key]; !ok {
			changed = append(changed, key)
		}
	}

	s.byKey = nextByKey
	s.hashByKey = nextHash
	if version != "" {
		s.version = version
	}
	return uniqueStrings(changed), nil
}

func hashMerchantConfig(mc *MerchantConfig) (string, error) {
	bs, err := json.Marshal(mc)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(bs)
	return hex.EncodeToString(sum[:]), nil
}

func uniqueStrings(in []string) []string {
	seen := make(map[string]struct{}, len(in))
	out := make([]string, 0, len(in))
	for _, v := range in {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
}

func merchantKey(tenantID, merchantID string) string {
	return tenantID + ":" + merchantID
}
