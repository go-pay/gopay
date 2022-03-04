package apple

import (
	"encoding/json"
	"github.com/go-pay/gopay/pkg/xlog"
	"testing"
)

func TestNotify(t *testing.T) {
	body := "{\"signedPayload\":\"eyJhbGciOiJFUzI1NiIsIng1YyI6WyJNSUlOW...mnpo2QrItvA\"}"
	var payload *NotificationV2SignedPayload
	err := json.Unmarshal([]byte(body), &payload)
	if err != nil {
		xlog.Error(err)
		return
	}
	rsp, err := payload.Decode()
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("notify data: %s", rsp)

}
