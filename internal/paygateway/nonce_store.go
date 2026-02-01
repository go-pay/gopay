package paygateway

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type NonceStore interface {
	UseOnce(ctx context.Context, nonce string, ttl time.Duration) (bool, error)
}

type MemoryNonceStore struct {
	mu      sync.Mutex
	entries map[string]time.Time
}

func NewMemoryNonceStore() *MemoryNonceStore {
	return &MemoryNonceStore{entries: make(map[string]time.Time)}
}

func (s *MemoryNonceStore) UseOnce(_ context.Context, nonce string, ttl time.Duration) (bool, error) {
	if nonce == "" {
		return false, errors.New("nonce is empty")
	}
	if ttl <= 0 {
		ttl = 5 * time.Minute
	}
	now := time.Now()
	s.mu.Lock()
	defer s.mu.Unlock()
	for k, exp := range s.entries {
		if now.After(exp) {
			delete(s.entries, k)
		}
	}
	if exp, ok := s.entries[nonce]; ok && now.Before(exp) {
		return false, nil
	}
	s.entries[nonce] = now.Add(ttl)
	return true, nil
}

type RedisNonceStore struct {
	rdb       *redis.Client
	keyPrefix string
}

func NewRedisNonceStore(rdb *redis.Client, keyPrefix string) *RedisNonceStore {
	return &RedisNonceStore{rdb: rdb, keyPrefix: keyPrefix}
}

func (s *RedisNonceStore) UseOnce(ctx context.Context, nonce string, ttl time.Duration) (bool, error) {
	if s == nil || s.rdb == nil {
		return false, errors.New("redis nonce store not initialized")
	}
	if nonce == "" {
		return false, errors.New("nonce is empty")
	}
	if ttl <= 0 {
		ttl = 5 * time.Minute
	}
	key := s.keyPrefix + "auth:nonce:" + nonce
	return s.rdb.SetNX(ctx, key, "1", ttl).Result()
}
