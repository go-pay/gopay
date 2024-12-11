package paypal

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
	"testing"
)

var (
	webhookId  = ""
	replaceUrl = ""
)

func TestClient_CreateWebhook(t *testing.T) {
	url := "https://thahao-test.mynatapp.cc/pay_order/paypal"
	bm := make(gopay.BodyMap)

	var eventTypes []*WebhookEventType
	item := &WebhookEventType{
		Name: "PAYMENT.CAPTURE.REFUNDED",
	}
	eventTypes = append(eventTypes, item)
	bm.Set("url", url).
		Set("event_types", eventTypes)
	ppRsp, err := client.CreateWebhook(client.ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if ppRsp.Code != Success {
		xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
		xlog.Debugf("ppRsp.Error: %+v", ppRsp.Error)
		xlog.Debugf("ppRsp.ErrorResponse: %+v", ppRsp.ErrorResponse)
		return
	}
	xlog.Debugf("ppRsp.Response: %+v", ppRsp.Response)
}

func TestClient_ListWebhook(t *testing.T) {
	ppRsp, err := client.ListWebhook(client.ctx)
	if err != nil {
		xlog.Error(err)
		return
	}
	if ppRsp.Code != Success {
		xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
		xlog.Debugf("ppRsp.Error: %+v", ppRsp.Error)
		xlog.Debugf("ppRsp.ErrorResponse: %+v", ppRsp.ErrorResponse)
		return
	}
	xlog.Debugf("ppRsp.Response: %+v", ppRsp.Response)
}

func TestClient_ShowWebhookDetail(t *testing.T) {
	ppRsp, err := client.ShowWebhookDetail(client.ctx, webhookId)
	if err != nil {
		xlog.Error(err)
		return
	}
	if ppRsp.Code != Success {
		xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
		xlog.Debugf("ppRsp.Error: %+v", ppRsp.Error)
		xlog.Debugf("ppRsp.ErrorResponse: %+v", ppRsp.ErrorResponse)
		return
	}
	xlog.Debugf("ppRsp.Response: %+v", ppRsp.Response)
}

func TestClient_UpdateWebhook(t *testing.T) {
	var ps []*Patch
	item := &Patch{
		Op:    "replace",
		Path:  "/url", // reference_id is yourself set when create order
		Value: replaceUrl,
	}
	ps = append(ps, item)
	ppRsp, err := client.UpdateWebhook(client.ctx, webhookId, ps)
	if err != nil {
		xlog.Error(err)
		return
	}
	if ppRsp.Code != Success {
		xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
		xlog.Debugf("ppRsp.Error: %+v", ppRsp.Error)
		xlog.Debugf("ppRsp.ErrorResponse: %+v", ppRsp.ErrorResponse)
		return
	}
	xlog.Debugf("ppRsp.Response: %+v", ppRsp.Response)
}

func TestClient_DeleteWebhook(t *testing.T) {
	ppRsp, err := client.DeleteWebhook(client.ctx, webhookId)
	if err != nil {
		xlog.Error(err)
		return
	}
	if ppRsp.Code != Success {
		xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
		xlog.Debugf("ppRsp.Error: %+v", ppRsp.Error)
		xlog.Debugf("ppRsp.ErrorResponse: %+v", ppRsp.ErrorResponse)
		return
	}
	xlog.Debugf("ppRsp.Response: %+v", ppRsp.Response)
}

func TestClient_VerifyWebhookSignature(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("auth_algo", "SHA256withRSA").
		Set("cert_url", "https://api.paypal.com/v1/notifications/certs/CERT-360caa42-fca2a594-ad47cb8d").
		Set("transmission_id", "b9d46480-2162-11ee-a2ae-61fbe51a886c").
		Set("transmission_sig", "NcbK6Mxok1iu12VU2bEgXUiFhifdX9eYlJJLtfc0etlVPgbigCZiQq3+Z8z7uNnCMh9S9rKjGr5eTscIHvUmB3jnPqUeLlGI3d670lXUkATH+p6Q/HI33ZidDAFTsgc3kZizqlONsPvmu5fdSA9UmKsaDmBEbACZXH/P4hTY4/pdAmk9OOPdySAhXj7gDwSz4ChMM0H+nSwXdyQC5IrjFQdoGABNoEPtRDUI7n0RCphu/kaZmQl7BtDXhoJAKYKmUS0pw4DhVW8hGoxBNrwizSW9eFE5tDhYO5WdGuWraGPKS5X/FD5JVfA2Kxj83rFvxHgyfKuYiMtnvevZVDp3Xg==").
		Set("transmission_time", "2023-07-13T09:50:40Z").
		Set("webhook_id", "3WA07241VT312694T").
		SetBodyMap("webhook_event", func(b gopay.BodyMap) {
			b.Set("event_version", "1.0").
				Set("resource_version", "2.0")
		})

	xlog.Debug("bm：", bm.JsonBody())
	verifyRes, err := client.VerifyWebhookSignature(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("verifyRes: %+v", verifyRes)
}
