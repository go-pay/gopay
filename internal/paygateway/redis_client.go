package paygateway

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

func newRedisClient(cfg RedisConfig) (*redis.Client, error) {
	if cfg.Addr == "" {
		return nil, nil
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		DialTimeout:  3 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		PoolSize:     64,
		MinIdleConns: 4,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		_ = rdb.Close()
		return nil, errors.New("redis ping failed: " + err.Error())
	}
	return rdb, nil
}
