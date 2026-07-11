package douyin

import (
	"testing"

	"github.com/go-pay/xlog"
)

func TestOrderQueryByTransactionId(t *testing.T) {
	rsp, err := client.OrderQueryByTransactionId(ctx, TransactionId)
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

func TestOrderQueryByOutTradeNo(t *testing.T) {
	rsp, err := client.OrderQueryByOutTradeNo(ctx, OutTradeNo)
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

func TestCloseOrder(t *testing.T) {
	rsp, err := client.CloseOrder(ctx, OutTradeNo, nil)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debug("close order success")
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}
