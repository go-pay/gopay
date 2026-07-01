package douyin

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

func TestProfitRequest(t *testing.T) {
	// 若接收方需要 name（type=MERCHANT_ID 时必传），请先加密
	//encName, err := client.EncryptText("接收方商户名称")
	//if err != nil { xlog.Error(err); return }

	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("transaction_id", TransactionId).
		Set("out_order_no", "SPLIT_"+util.RandomString(16)).
		Set("unfreeze_unsplit", false).
		Set("notify_url", NotifyUrl).
		Set("receivers", []ProfitReceiverReq{
			{Type: "MERCHANT_ID", Account: "6020230307605084", /* Name: encName, */ Amount: 100, Description: "分给合作方"},
		})

	rsp, err := client.ProfitRequest(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestProfitQuery(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("transaction_id", TransactionId)

	rsp, err := client.ProfitQuery(ctx, OutOrderNo, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestProfitRollback(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("order_id", "抖音支付分账单号").
		Set("out_order_no", OutOrderNo).
		Set("out_return_no", "OUT_"+util.RandomString(16)).
		Set("return_mchid", Mchid).
		Set("amount", 100).
		Set("description", "退分账")

	rsp, err := client.ProfitRollback(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestProfitRollbackQuery(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("out_order_no", OutOrderNo)

	rsp, err := client.ProfitRollbackQuery(ctx, OutReturnNo, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestProfitComplete(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("transaction_id", TransactionId).
		Set("out_order_no", "FIN_"+util.RandomString(16)).
		Set("description", "完结分账").
		Set("notify_url", NotifyUrl)

	rsp, err := client.ProfitComplete(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestProfitBalanceQuery(t *testing.T) {
	rsp, err := client.ProfitBalanceQuery(ctx, TransactionId, "")
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestProfitReceiverAdd(t *testing.T) {
	// name 需先加密
	//encName, err := client.EncryptText("接收方名称")
	//if err != nil { xlog.Error(err); return }

	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("type", "MERCHANT_ID").
		Set("account", "6020230307605084").
		//Set("name", encName).
		Set("relation_type", "STORE")

	rsp, err := client.ProfitReceiverAdd(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestProfitReceiverDelete(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("type", "MERCHANT_ID").
		Set("account", "6020230307605084")

	rsp, err := client.ProfitReceiverDelete(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}
