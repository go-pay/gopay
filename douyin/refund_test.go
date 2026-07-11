package douyin

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

func TestRefund(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("out_trade_no", OutTradeNo).
		Set("out_refund_no", "REF_"+util.RandomString(16)).
		Set("reason", "用户申请退款").
		Set("notify_url", NotifyUrl).
		SetBodyMap("amount", func(b gopay.BodyMap) {
			b.Set("refund", 1).Set("total", 1).Set("currency", "CNY")
		})

	rsp, err := client.Refund(ctx, bm)
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

func TestRefundQuery(t *testing.T) {
	rsp, err := client.RefundQuery(ctx, OutRefundNo, "", Appid)
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
