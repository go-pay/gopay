package douyin

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

func TestApplyTradeBill(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("bill_date", "2026-06-30")

	rsp, err := client.ApplyTradeBill(ctx, bm)
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

func TestApplyFundBill(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("bill_date", "2026-06-30").Set("account_type", "BaseAccount")

	rsp, err := client.ApplyFundBill(ctx, bm)
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

func TestApplyProfitBill(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("bill_date", "2026-06-30")

	rsp, err := client.ApplyProfitBill(ctx, bm)
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

func TestDownloadBillFile(t *testing.T) {
	// 从 ApplyXxxBill 的返回中取 download_url，填入下面
	url := ""
	fileBytes, err := client.DownloadBillFile(ctx, url)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("fileBytes len:%d", len(fileBytes))

	// 若为 GZIP 压缩包，可解压
	//raw, err := UngzipBill(fileBytes)
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//xlog.Debugf("raw len:%d", len(raw))

	// 完整性校验（hashType 与 hashValue 来自 ApplyXxxBill 响应）
	//if err := VerifyBillHash(raw, "SHA1", "hash_value_from_apply_response"); err != nil {
	//	xlog.Error(err)
	//}
}
