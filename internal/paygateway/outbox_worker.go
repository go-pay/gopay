package paygateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type OutboxWorker struct {
	rdb      *redis.Client
	stream   string
	group    string
	consumer string

	minIdle time.Duration
	block   time.Duration
	count   int64

	publisher EventPublisher

	cancel context.CancelFunc
	done   chan struct{}
}

type OutboxWorkerConfig struct {
	Stream   string
	Group    string
	MinIdle  time.Duration
	Block    time.Duration
	Count    int64
	Consumer string
}

func NewOutboxWorker(rdb *redis.Client, cfg OutboxWorkerConfig, publisher EventPublisher) (*OutboxWorker, error) {
	if rdb == nil {
		return nil, errors.New("outbox requires redis")
	}
	if cfg.Stream == "" {
		return nil, errors.New("outbox stream is empty")
	}
	if cfg.Group == "" {
		return nil, errors.New("outbox consumer group is empty")
	}
	if cfg.MinIdle <= 0 {
		cfg.MinIdle = 30 * time.Second
	}
	if cfg.Block <= 0 {
		cfg.Block = 2 * time.Second
	}
	if cfg.Count <= 0 {
		cfg.Count = 16
	}
	if cfg.Consumer == "" {
		host, _ := os.Hostname()
		cfg.Consumer = fmt.Sprintf("%s-%d", host, os.Getpid())
	}
	if publisher == nil {
		return nil, errors.New("outbox publisher is nil")
	}
	return &OutboxWorker{
		rdb:       rdb,
		stream:    cfg.Stream,
		group:     cfg.Group,
		consumer:  cfg.Consumer,
		minIdle:   cfg.MinIdle,
		block:     cfg.Block,
		count:     cfg.Count,
		publisher: publisher,
	}, nil
}

func (w *OutboxWorker) Start() error {
	if w == nil {
		return nil
	}
	if w.done != nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := w.ensureConsumerGroup(ctx); err != nil {
		return err
	}

	ctx2, cancel2 := context.WithCancel(context.Background())
	w.cancel = cancel2
	w.done = make(chan struct{})
	go w.loop(ctx2)
	return nil
}

func (w *OutboxWorker) Stop() {
	if w == nil || w.cancel == nil || w.done == nil {
		return
	}
	w.cancel()
	<-w.done
}

func (w *OutboxWorker) ensureConsumerGroup(ctx context.Context) error {
	if w == nil {
		return nil
	}
	err := w.rdb.XGroupCreateMkStream(ctx, w.stream, w.group, "0").Err()
	if err == nil {
		return nil
	}
	// BUSYGROUP: Consumer Group name already exists
	if strings.Contains(err.Error(), "BUSYGROUP") {
		return nil
	}
	return err
}

func (w *OutboxWorker) loop(ctx context.Context) {
	defer close(w.done)

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		// 1) Reclaim and process pending messages (from crashed consumers).
		w.claimAndProcess(ctx)

		// 2) Read new messages.
		// This blocks for up to w.block, so it also acts as a throttle when idle.
		w.readAndProcess(ctx)
	}
}

func (w *OutboxWorker) claimAndProcess(ctx context.Context) {
	start := "0-0"
	for {
		msgs, nextStart, err := w.rdb.XAutoClaim(ctx, &redis.XAutoClaimArgs{
			Stream:   w.stream,
			Group:    w.group,
			Consumer: w.consumer,
			MinIdle:  w.minIdle,
			Start:    start,
			Count:    w.count,
		}).Result()
		if err != nil {
			return
		}
		if len(msgs) == 0 {
			return
		}
		w.processMessages(ctx, msgs)
		start = nextStart
	}
}

func (w *OutboxWorker) readAndProcess(ctx context.Context) {
	streams, err := w.rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    w.group,
		Consumer: w.consumer,
		Streams:  []string{w.stream, ">"},
		Count:    w.count,
		Block:    w.block,
	}).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return
		}
		return
	}
	for _, s := range streams {
		w.processMessages(ctx, s.Messages)
	}
}

func (w *OutboxWorker) processMessages(ctx context.Context, msgs []redis.XMessage) {
	for _, msg := range msgs {
		eventBytes, ok := msg.Values["event"]
		if !ok {
			_ = w.rdb.XAck(ctx, w.stream, w.group, msg.ID).Err()
			continue
		}
		bs, err := asBytes(eventBytes)
		if err != nil {
			_ = w.rdb.XAck(ctx, w.stream, w.group, msg.ID).Err()
			continue
		}
		var ev Event
		if err := json.Unmarshal(bs, &ev); err != nil {
			_ = w.rdb.XAck(ctx, w.stream, w.group, msg.ID).Err()
			continue
		}

		if err := w.publisher.Publish(ctx, &ev); err != nil {
			log.Printf("outbox publish failed id=%s eventId=%s: %v", msg.ID, ev.EventID, err)
			continue
		}
		if err := w.rdb.XAck(ctx, w.stream, w.group, msg.ID).Err(); err != nil {
			log.Printf("outbox ack failed id=%s: %v", msg.ID, err)
		}
	}
}

func asBytes(v any) ([]byte, error) {
	switch t := v.(type) {
	case []byte:
		return t, nil
	case string:
		return []byte(t), nil
	default:
		return nil, fmt.Errorf("unsupported value type: %T", v)
	}
}
