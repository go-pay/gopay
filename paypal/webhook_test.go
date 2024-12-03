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
