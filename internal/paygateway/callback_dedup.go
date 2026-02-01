package paygateway

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type DedupState string

const (
	DedupStateNone       DedupState = ""
	DedupStateProcessing DedupState = "PROCESSING"
	DedupStateDone       DedupState = "DONE"
)

type CallbackDeduper interface {
	// TryLock attempts to lock processing for the given key.
	// If locked=false, state indicates the current state (PROCESSING or DONE).
	TryLock(ctx context.Context, key string) (locked bool, state DedupState, err error)
	MarkDone(ctx context.Context, key string) error
	Release(ctx context.Context, key string) error
}

// ============================================
// Memory deduper (dev-only / single instance)

type MemoryCallbackDeduper struct {
	mu            sync.Mutex
	entries       map[string]dedupEntry
	processingTTL time.Duration
	doneTTL       time.Duration
}

type dedupEntry struct {
	state     DedupState
	expiresAt time.Time
}

func NewMemoryCallbackDeduper(processingTTL, doneTTL time.Duration) *MemoryCallbackDeduper {
	if processingTTL <= 0 {
		processingTTL = 5 * time.Minute
	}
	if doneTTL <= 0 {
		doneTTL = 7 * 24 * time.Hour
	}
	return &MemoryCallbackDeduper{
		entries:       make(map[string]dedupEntry),
		processingTTL: processingTTL,
		doneTTL:       doneTTL,
	}
}

func (d *MemoryCallbackDeduper) TryLock(_ context.Context, key string) (locked bool, state DedupState, err error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if e, ok := d.entries[key]; ok {
		if time.Now().After(e.expiresAt) {
			delete(d.entries, key)
		} else {
			return false, e.state, nil
		}
	}
	d.entries[key] = dedupEntry{state: DedupStateProcessing, expiresAt: time.Now().Add(d.processingTTL)}
	return true, DedupStateNone, nil
}

func (d *MemoryCallbackDeduper) MarkDone(_ context.Context, key string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.entries[key] = dedupEntry{state: DedupStateDone, expiresAt: time.Now().Add(d.doneTTL)}
	return nil
}

func (d *MemoryCallbackDeduper) Release(_ context.Context, key string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.entries, key)
	return nil
}

// ============================================
// Redis deduper (prod recommended)

type RedisCallbackDeduper struct {
	rdb           *redis.Client
	keyPrefix     string
	processingTTL time.Duration
	doneTTL       time.Duration
}

func NewRedisCallbackDeduper(rdb *redis.Client, keyPrefix string, processingTTL, doneTTL time.Duration) *RedisCallbackDeduper {
	if processingTTL <= 0 {
		processingTTL = 5 * time.Minute
	}
	if doneTTL <= 0 {
		doneTTL = 7 * 24 * time.Hour
	}
	return &RedisCallbackDeduper{
		rdb:           rdb,
		keyPrefix:     keyPrefix,
		processingTTL: processingTTL,
		doneTTL:       doneTTL,
	}
}

func (d *RedisCallbackDeduper) TryLock(ctx context.Context, key string) (locked bool, state DedupState, err error) {
	if d == nil || d.rdb == nil {
		return false, DedupStateNone, errors.New("redis deduper not initialized")
	}
	rk := d.redisKey(key)
	ok, err := d.rdb.SetNX(ctx, rk, string(DedupStateProcessing), d.processingTTL).Result()
	if err != nil {
		return false, DedupStateNone, err
	}
	if ok {
		return true, DedupStateNone, nil
	}
	val, err := d.rdb.Get(ctx, rk).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, DedupStateNone, nil
		}
		return false, DedupStateNone, err
	}
	return false, DedupState(val), nil
}

func (d *RedisCallbackDeduper) MarkDone(ctx context.Context, key string) error {
	if d == nil || d.rdb == nil {
		return errors.New("redis deduper not initialized")
	}
	return d.rdb.Set(ctx, d.redisKey(key), string(DedupStateDone), d.doneTTL).Err()
}

func (d *RedisCallbackDeduper) Release(ctx context.Context, key string) error {
	if d == nil || d.rdb == nil {
		return errors.New("redis deduper not initialized")
	}
	return d.rdb.Del(ctx, d.redisKey(key)).Err()
}

func (d *RedisCallbackDeduper) redisKey(key string) string {
	return d.keyPrefix + "cb:" + key
}
