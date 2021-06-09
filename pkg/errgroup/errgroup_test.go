package errgroup

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-pay/gopay/pkg/xlog"
)

func TestErrgroup(t *testing.T) {
	var count int64 = 1
	countBackup := count
	eg := WithContext(context.Background())

	// go 协程
	eg.Go(func(ctx context.Context) error {
		atomic.AddInt64(&count, 1)
		return nil
	})
	// go 协程
	eg.Go(func(ctx context.Context) error {
		atomic.AddInt64(&count, 1)
		return nil
	})
	// go 协程
	eg.Go(func(ctx context.Context) error {
		atomic.AddInt64(&count, 1)
		return errors.New("error ,reset count")
	})
	// wait 协程 Done
	if err := eg.Wait(); err != nil {
		// do some thing
		count = countBackup
		xlog.Error(err)
		//return
	}
	xlog.Debug(count)
}

func TestErrgroup1(t *testing.T) {
	var (
		count int64 = 1
		eg    Group
		goNum = 3 // every times run goNum goroutine
	)
	for i := 0; i < 10; i++ {
		eg.Go(func(ctx context.Context) error {
			atomic.AddInt64(&count, 1)
			xlog.Debug("count:", count)
			return nil
		})
		if eg.WorkNum() == goNum {
			if err := eg.Wait(); err != nil {
				xlog.Error(err)
				// to do something you need
			}
			xlog.Info("wat")
			time.Sleep(time.Second)
		}
	}
}
