package paygateway

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type MerchantStore struct {
	mu              sync.RWMutex
	byKey           map[string]*MerchantConfig
	hashByKey       map[string]string
	version         string
	defaultTenantID string
	secretsBaseDir  string
}

func NewMerchantStore(cfg *Config) (*MerchantStore, error) {
	ms := &MerchantStore{
		byKey:           make(map[string]*MerchantConfig),
		hashByKey:       make(map[string]string),
		defaultTenantID: cfg.DefaultTenantID,
		secretsBaseDir:  cfg.SecretsBaseDir,
	}
	if ms.secretsBaseDir == "" {
		return nil, fmt.Errorf("secretsBaseDir is empty")
	}
	if !filepath.IsAbs(ms.secretsBaseDir) {
		return nil, fmt.Errorf("secretsBaseDir must be absolute: %s", ms.secretsBaseDir)
	}
	if _, err := ms.Replace(cfg.Merchants, "bootstrap"); err != nil {
		return nil, err
	}
	return ms, nil
}

func safeJoinUnderBaseDir(baseDir, ref string) (string, error) {
	ref = strings.TrimSpace(ref)
	if ref == "" {
		return "", fmt.Errorf("secret ref is empty")
	}
	if strings.ContainsRune(ref, '\x00') {
		return "", fmt.Errorf("invalid secret ref")
	}

	base := filepath.Clean(baseDir)
	if base == "." || base == "" {
		return "", fmt.Errorf("invalid secretsBaseDir")
	}

	var target string
	if filepath.IsAbs(ref) {
		target = filepath.Clean(ref)
	} else {
		target = filepath.Clean(filepath.Join(base, ref))
	}

	rel, err := filepath.Rel(base, target)
	if err != nil {
		return "", fmt.Errorf("resolve secret ref: %w", err)
	}
	if rel == "." || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("secret ref escapes base dir")
	}
	return target, nil
}

func resolveMerchantSecretFiles(m *MerchantConfig, secretsBaseDir string) error {
	if m == nil {
		return nil
	}
	if m.WechatV3 != nil {
		if m.WechatV3.PrivateKey == "" && (m.WechatV3.PrivateKeyRef != "" || m.WechatV3.PrivateKeyFile != "") {
			ref := firstNonEmpty(m.WechatV3.PrivateKeyRef, m.WechatV3.PrivateKeyFile)
			path, err := safeJoinUnderBaseDir(secretsBaseDir, ref)
			if err != nil {
				return fmt.Errorf("resolve wechatV3 private key: %w", err)
			}
			bs, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("read wechatV3 private key: %w", err)
			}
			m.WechatV3.PrivateKey = strings.TrimSpace(string(bs))
		}
		if m.WechatV3.WechatPayPublicKey == "" && (m.WechatV3.WechatPayPublicKeyRef != "" || m.WechatV3.WechatPayPublicKeyFile != "") {
			ref := firstNonEmpty(m.WechatV3.WechatPayPublicKeyRef, m.WechatV3.WechatPayPublicKeyFile)
			path, err := safeJoinUnderBaseDir(secretsBaseDir, ref)
			if err != nil {
				return fmt.Errorf("resolve wechatV3 wechatPayPublicKey: %w", err)
			}
			bs, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("read wechatV3 wechatPayPublicKey: %w", err)
			}
			m.WechatV3.WechatPayPublicKey = strings.TrimSpace(string(bs))
		}
	}
	if m.Alipay != nil {
		if m.Alipay.PrivateKey == "" && (m.Alipay.PrivateKeyRef != "" || m.Alipay.PrivateKeyFile != "") {
			ref := firstNonEmpty(m.Alipay.PrivateKeyRef, m.Alipay.PrivateKeyFile)
			path, err := safeJoinUnderBaseDir(secretsBaseDir, ref)
			if err != nil {
				return fmt.Errorf("resolve alipay private key: %w", err)
			}
			bs, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("read alipay private key: %w", err)
			}
			m.Alipay.PrivateKey = strings.TrimSpace(string(bs))
		}
		if m.Alipay.AlipayPublicKey == "" && (m.Alipay.AlipayPublicKeyRef != "" || m.Alipay.AlipayPublicKeyFile != "") {
			ref := firstNonEmpty(m.Alipay.AlipayPublicKeyRef, m.Alipay.AlipayPublicKeyFile)
			path, err := safeJoinUnderBaseDir(secretsBaseDir, ref)
			if err != nil {
				return fmt.Errorf("resolve alipay public key: %w", err)
			}
			bs, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("read alipay public key: %w", err)
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
		if err := resolveMerchantSecretFiles(&mc, s.secretsBaseDir); err != nil {
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
