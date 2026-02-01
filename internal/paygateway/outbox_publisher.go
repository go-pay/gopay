package paygateway

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/redis/go-redis/v9"
)

type RedisOutboxPublisher struct {
	rdb    *redis.Client
	stream string
}

func NewRedisOutboxPublisher(rdb *redis.Client, stream string) *RedisOutboxPublisher {
	return &RedisOutboxPublisher{rdb: rdb, stream: stream}
}

func (p *RedisOutboxPublisher) Publish(ctx context.Context, event *Event) error {
	if p == nil || p.rdb == nil {
		return errors.New("redis outbox publisher not initialized")
	}
	if p.stream == "" {
		return errors.New("redis outbox stream is empty")
	}
	bs, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return p.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: p.stream,
		Values: map[string]interface{}{"event": bs},
	}).Err()
}
