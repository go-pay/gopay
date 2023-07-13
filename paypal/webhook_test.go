package paypal

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
	"testing"
)

func TestVerifyWebhookSignature(t *testing.T) {

	bm := make(gopay.BodyMap)
	bm.Set("auth_algo", "").
		Set("cert_url", "").
		Set("transmission_id", "").
		Set("transmission_sig", "").
		Set("transmission_time", "").
		Set("webhook_id", "").
		SetBodyMap("webhook_event", func(b gopay.BodyMap) {
			b.Set("event_version", "").
				Set("resource_version", "")
		})

	xlog.Debug("bm：", bm.JsonBody())
	verifyRes, err := client.VerifyWebhookSignature(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("verifyRes: %+v", verifyRes)
}
