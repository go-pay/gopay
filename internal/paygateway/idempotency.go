package paygateway

import (
	"sync"
	"time"
)

type IdempotencyStore struct {
	mu      sync.Mutex
	entries map[string]idemEntry
	ttl     time.Duration
}

type idemEntry struct {
	status    int
	body      []byte
	createdAt time.Time
}

func NewIdempotencyStore(ttl time.Duration) *IdempotencyStore {
	if ttl <= 0 {
		ttl = time.Hour
	}
	return &IdempotencyStore{
		entries: make(map[string]idemEntry),
		ttl:     ttl,
	}
}

func (s *IdempotencyStore) Get(key string) (status int, body []byte, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.entries[key]
	if !ok {
		return 0, nil, false
	}
	if time.Since(e.createdAt) > s.ttl {
		delete(s.entries, key)
		return 0, nil, false
	}
	return e.status, e.body, true
}

func (s *IdempotencyStore) Put(key string, status int, body []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.entries[key] = idemEntry{status: status, body: body, createdAt: time.Now()}
}
