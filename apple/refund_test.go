package apple

import (
	"testing"

	"github.com/go-pay/xlog"
)

func TestGetRefundHistory(t *testing.T) {
	rsp, err := client.GetRefundHistory(ctx, "2000000184445477", "revision")
	if err != nil {
		if statusErr, ok := IsStatusCodeError(err); ok {
			xlog.Errorf("%+v", statusErr)
			// do something
			return
		}
		xlog.Errorf("client.GetRefundHistory(),err:%+v", err)
		return
	}
	for _, v := range rsp.SignedTransactions {
		transaction, _ := v.DecodeSignedTransaction()
		xlog.Debugf("refund transactions:%+v", transaction)
	}

}
