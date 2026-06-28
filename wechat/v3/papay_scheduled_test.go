package wechat

import (
	"encoding/json"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

func TestV3ScheduledDeductPreSignMiniProgram(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", "wxd678efh567hg6787").
		Set("openid", "oYobu0MVnQfWpSMOYJz2AHPG_gQw").
		Set("plan_id", 12535).
		Set("out_contract_code", "wxwtdk20200910100000").
		Set("contract_display_account", "微信代扣用户A").
		Set("contract_notify_url", "https://yourapp.com/notify")

	wxRsp, err := client.V3ScheduledDeductPreSignMiniProgram(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Debugf("wxRsp: %+v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

func TestV3ScheduledDeductPreSignApp(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", "wxd678efh567hg6787").
		Set("plan_id", 12535).
		Set("out_contract_code", "wxwtdk20200910100001").
		Set("contract_display_account", "微信代扣用户A").
		Set("contract_notify_url", "https://yourapp.com/notify")

	wxRsp, err := client.V3ScheduledDeductPreSignApp(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Debugf("wxRsp: %+v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

func TestV3ScheduledDeductPreSignH5(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", "wxd678efh567hg6787").
		Set("plan_id", 12535).
		Set("out_contract_code", "wxwtdk20200910100002").
		Set("contract_display_account", "微信代扣用户A").
		Set("contract_notify_url", "https://yourapp.com/notify")

	wxRsp, err := client.V3ScheduledDeductPreSignH5(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Debugf("wxRsp: %+v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

func TestV3ScheduledDeductPreSignJsapi(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", "wxd678efh567hg6787").
		Set("openid", "oYobu0MVnQfWpSMOYJz2AHPG_gQw").
		Set("plan_id", 12535).
		Set("out_contract_code", "wxwtdk20200910100003").
		Set("contract_display_account", "微信代扣用户A").
		Set("contract_notify_url", "https://yourapp.com/notify")

	wxRsp, err := client.V3ScheduledDeductPreSignJsapi(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Debugf("wxRsp: %+v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

func TestV3ScheduledDeductContractQuery(t *testing.T) {
	wxRsp, err := client.V3ScheduledDeductContractQuery(ctx, "12535", "wxwtdk20200910100000")
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Debugf("wxRsp: %+v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

func TestV3ScheduledDeductContractTerminate(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("contract_termination_remark", "用户主动解约")

	wxRsp, err := client.V3ScheduledDeductContractTerminate(ctx, "12535", "wxwtdk20200910100000", bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Debugf("wxRsp: %+v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

func TestV3ScheduledDeductSchedule(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", "wxd678efh567hg6787").
		SetBodyMap("schedule_amount", func(bm gopay.BodyMap) {
			bm.Set("total", 1).Set("currency", "CNY")
		})

	wxRsp, err := client.V3ScheduledDeductSchedule(ctx, "123124412412423431", bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Debugf("wxRsp: %+v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

func TestV3ScheduledDeductScheduleQuery(t *testing.T) {
	wxRsp, err := client.V3ScheduledDeductScheduleQuery(ctx, "123124412412423431")
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Debugf("wxRsp: %+v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

func TestV3ScheduledDeductApply(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", "wxd678efh567hg6787").
		Set("out_trade_no", "1217752501201407033233368018").
		Set("description", "测试受理扣款").
		Set("transaction_notify_url", "https://yourapp.com/pay-notify").
		Set("contract_id", "123124412412423431").
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", 1).Set("currency", "CNY")
		})

	wxRsp, err := client.V3ScheduledDeductApply(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Debugf("wxRsp: %+v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

// 验证回调 resource 解密后的 JSON 能正确反序列化到 model 上。
// 测试 JSON 来源：
//   - 签约通知：https://pay.weixin.qq.com/doc/v3/merchant/4012286323
//   - 支付通知：https://pay.weixin.qq.com/doc/v3/merchant/4012286313
func TestPapayScheduledNotifyDecode(t *testing.T) {
	signPlain := []byte(`{
		"appid": "wxd678efh567hg6787",
		"contract_display_account": "微信代扣用户A",
		"contract_expired_time": "2021-09-10T13:29:35+08:00",
		"contract_id": "123124412412423431",
		"contract_signed_time": "2020-09-10T13:29:35+08:00",
		"contract_state": "SIGNED",
		"deduct_schedule": {
			"deduct_amount": {"currency":"CNY","total":1},
			"deduct_date": "2019-11-22",
			"estimated_deduct_amount": {"currency":"CNY","total":1},
			"estimated_deduct_date": "2019-11-22",
			"schedule_state": "PAID",
			"scheduled_amount": {"currency":"CNY","total":1}
		},
		"mchid": "1900000109",
		"openid": "o-MYE42l80oelYMDE34nYD456Xoy",
		"out_contract_code": "wxwtdk20200910100000",
		"out_user_code": "用户A",
		"plan_id": 12535
	}`)
	sign := new(PapayScheduledSignNotifyResource)
	if err := json.Unmarshal(signPlain, sign); err != nil {
		t.Fatalf("sign notify unmarshal: %v", err)
	}
	if sign.ContractState != PapayContractStateSigned {
		t.Fatalf("sign notify contract_state mismatch: %q", sign.ContractState)
	}
	if sign.PlanId != 12535 || sign.OutContractCode != "wxwtdk20200910100000" {
		t.Fatalf("sign notify field mismatch: %+v", sign)
	}
	if sign.DeductSchedule == nil || sign.DeductSchedule.ScheduleState != PapayScheduleStatePaid {
		t.Fatalf("sign notify deduct_schedule mismatch: %+v", sign.DeductSchedule)
	}

	payPlain := []byte(`{
		"transaction_id":"1217752501201407033233368018",
		"amount":{"payer_total":100,"total":100,"currency":"CNY","payer_currency":"CNY"},
		"mchid":"1230000109",
		"trade_state":"SUCCESS",
		"trade_state_desc":"支付成功",
		"trade_type":"PAP",
		"bank_type":"CMC",
		"out_trade_no":"1217752501201407033233368018",
		"success_time":"2020-09-10T13:29:35+08:00",
		"payer":{"openid":"o-MYE42l80oelYMDE34nYD456Xoy"},
		"appid":"wxd678efh567hg6787"
	}`)
	pay := new(PapayScheduledPayNotifyResource)
	if err := json.Unmarshal(payPlain, pay); err != nil {
		t.Fatalf("pay notify unmarshal: %v", err)
	}
	if pay.TradeState != "SUCCESS" || pay.TradeType != "PAP" {
		t.Fatalf("pay notify state mismatch: %+v", pay)
	}
	if pay.Amount == nil || pay.Amount.Total != 100 || pay.Amount.PayerTotal != 100 {
		t.Fatalf("pay notify amount mismatch: %+v", pay.Amount)
	}
	if pay.Payer == nil || pay.Payer.Openid == "" {
		t.Fatalf("pay notify payer mismatch: %+v", pay.Payer)
	}
}
