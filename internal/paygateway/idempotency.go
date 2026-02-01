package paygateway

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type IdempotencyStore interface {
	Get(ctx context.Context, key string) (status int, body []byte, ok bool, err error)
	Put(ctx context.Context, key string, status int, body []byte) error
}

// ============================================
// Memory store (dev-only / single instance)

type MemoryIdempotencyStore struct {
	mu      sync.Mutex
	entries map[string]idemEntry
	ttl     time.Duration
}

type idemEntry struct {
	status    int
	body      []byte
	createdAt time.Time
}

func NewMemoryIdempotencyStore(ttl time.Duration) *MemoryIdempotencyStore {
	if ttl <= 0 {
		ttl = time.Hour
	}
	return &MemoryIdempotencyStore{
		entries: make(map[string]idemEntry),
		ttl:     ttl,
	}
}

func (s *MemoryIdempotencyStore) Get(_ context.Context, key string) (status int, body []byte, ok bool, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.entries[key]
	if !ok {
		return 0, nil, false, nil
	}
	if time.Since(e.createdAt) > s.ttl {
		delete(s.entries, key)
		return 0, nil, false, nil
	}
	return e.status, e.body, true, nil
}

func (s *MemoryIdempotencyStore) Put(_ context.Context, key string, status int, body []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.entries[key] = idemEntry{status: status, body: body, createdAt: time.Now()}
	return nil
}

// ============================================
// Redis store (prod recommended)

type RedisIdempotencyStore struct {
	rdb       *redis.Client
	keyPrefix string
	ttl       time.Duration
}

type redisIdemValue struct {
	Status  int    `json:"status"`
	BodyB64 string `json:"bodyB64"`
}

func NewRedisIdempotencyStore(rdb *redis.Client, keyPrefix string, ttl time.Duration) *RedisIdempotencyStore {
	if ttl <= 0 {
		ttl = 24 * time.Hour
	}
	return &RedisIdempotencyStore{
		rdb:       rdb,
		keyPrefix: keyPrefix,
		ttl:       ttl,
	}
}

func (s *RedisIdempotencyStore) Get(ctx context.Context, key string) (status int, body []byte, ok bool, err error) {
	if s == nil || s.rdb == nil {
		return 0, nil, false, errors.New("redis idempotency store not initialized")
	}
	val, err := s.rdb.Get(ctx, s.redisKey(key)).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil, false, nil
		}
		return 0, nil, false, err
	}
	var v redisIdemValue
	if err := json.Unmarshal(val, &v); err != nil {
		return 0, nil, false, err
	}
	body, err = base64.StdEncoding.DecodeString(v.BodyB64)
	if err != nil {
		return 0, nil, false, err
	}
	return v.Status, body, true, nil
}

func (s *RedisIdempotencyStore) Put(ctx context.Context, key string, status int, body []byte) error {
	if s == nil || s.rdb == nil {
		return errors.New("redis idempotency store not initialized")
	}
	v := redisIdemValue{
		Status:  status,
		BodyB64: base64.StdEncoding.EncodeToString(body),
	}
	bs, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return s.rdb.Set(ctx, s.redisKey(key), bs, s.ttl).Err()
}

func (s *RedisIdempotencyStore) redisKey(key string) string {
	return s.keyPrefix + "idem:" + key
}
