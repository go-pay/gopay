package douyin

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

func TestTransfer(t *testing.T) {
	// 大额转账（≥2000元）时 user_name 必填并先加密
	//encName, err := client.EncryptText("张三")
	//if err != nil { xlog.Error(err); return }

	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("out_bill_no", "OUT_"+util.RandomString(16)).
		Set("transfer_scene_id", "SCENE_001").
		Set("openid", Openid).
		//Set("user_name", encName).
		Set("transfer_amount", 100).
		Set("transfer_remark", "商户转账").
		Set("notify_url", NotifyUrl)

	rsp, err := client.Transfer(ctx, bm)
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

func TestTransferQueryByOutBillNo(t *testing.T) {
	rsp, err := client.TransferQueryByOutBillNo(ctx, OutBillNo)
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

func TestTransferQueryByTransferBillNo(t *testing.T) {
	rsp, err := client.TransferQueryByTransferBillNo(ctx, TransferBillNo)
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
