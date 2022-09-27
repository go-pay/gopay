package retry

import (
	"errors"
	"testing"
	"time"

	"github.com/go-pay/gopay/pkg/xlog"
)

func TestRetry(t *testing.T) {
	err := Retry(func() error {
		xlog.Warnf("retry func")
		return errors.New("please retry")
	}, 3, 2*time.Second)
	if err != nil {
		xlog.Error(err)
	}
}
