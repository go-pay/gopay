/*
@Author: wzy
@Time: 2022/6/8
*/
package paypal

import (
	"encoding/json"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

func TestCreateBatchPayout(t *testing.T) {
	receiver := "test-email@testemail.com"
	bm := make(gopay.BodyMap)
	bm.SetBodyMap("sender_batch_header", func(bm gopay.BodyMap) {
		bm.Set("sender_batch_id", "2022060811140003").
			Set("email_subject", "You have a payout!").
			Set("email_message", "You have received a payout! Thanks for using our service!")
	}).Set("items", []map[string]any{
		{
			"recipient_type": "EMAIL",
			"amount": map[string]string{
				"value":    "5",
				"currency": "USD",
			},
			"note":           "Thanks for your verify",
			"sender_item_id": "20220608111304",
			"receiver":       receiver,
		},
	})
	xlog.Debug("bm", bm.JsonBody())
	ppRsp, err := client.CreateBatchPayout(ctx, bm)
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
	_rspJson, _ := json.MarshalIndent(ppRsp.Response, "", "\t")
	xlog.Debugf("ppRsp.Response: %s", _rspJson)
	for _, v := range ppRsp.Response.Links {
		xlog.Debugf("ppRsp.Response.Links: %+v", v)
	}
}

func TestShowPayoutBatchDetails(t *testing.T) {
	ppRsp, err := client.ShowPayoutBatchDetails(ctx, "YSESATZEDPRY6", nil)
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
	for _, v := range ppRsp.Response.Items {
		_item, _ := json.MarshalIndent(v, "", "\t")
		xlog.Debugf("ppRsp.Response.PayoutItemDetail: \n%s", _item)
	}
	for _, v := range ppRsp.Response.Links {
		xlog.Debugf("ppRsp.Response.Links: %+v", v)
	}
}

func TestShowPayoutItemDetails(t *testing.T) {
	ppRsp, err := client.ShowPayoutItemDetails(ctx, "HGYYMWW7PRJKW")
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
	_rspJson, _ := json.MarshalIndent(ppRsp.Response, "", "\t")
	xlog.Debugf("ppRsp.Response: %s", _rspJson)
}

func TestCancelUnclaimedPayoutItem(t *testing.T) {
	ppRsp, err := client.CancelUnclaimedPayoutItem(ctx, "HGYYMWW7PRJKW")
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
	_rspJson, _ := json.MarshalIndent(ppRsp.Response, "", "\t")
	xlog.Debugf("ppRsp.Response: %s", _rspJson)
}
