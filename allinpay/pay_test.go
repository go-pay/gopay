package allinpay

import (
	"testing"
	"time"

	"github.com/go-pay/xlog"

	"github.com/go-pay/gopay"
)

func TestClient_ScanPay(t *testing.T) {
	// 扫码支付
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("trxamt", "1").
		Set("reqsn", "larry01").
		Set("body", "支付测试").
		SetBodyMap("terminfo", func(b gopay.BodyMap) {
			b.Set("devicetype", "10").
				Set("termno", "00000001")
		}).
		Set("authcode", "131104796948096102")
	// 创建订单
	resp, err := client.ScanPay(ctx, bm)
	xlog.Debugf("aliRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
}

func TestClient_Query(t *testing.T) {
	// 查询订单
	resp, err := client.Query(ctx, OrderTypeReqSN, "larry01")
	xlog.Debugf("aliRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
}

func TestClient_Refund(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("trxamt", "1").
		Set("reqsn", "relarry01").
		Set("remark", "支付测试退款").
		Set("oldreqsn", "larry01")

	// 退款
	resp, err := client.Refund(ctx, bm)
	xlog.Debugf("allRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
}

func TestClient_Cancel(t *testing.T) {
	// 订单退款
	bm := make(gopay.BodyMap)
	bm.Set("trxamt", "1").
		Set("reqsn", "cclarry01").
		Set("remark", "支付测试取消").
		Set("oldreqsn", "larry01")

	// 取消订单
	resp, err := client.Cancel(ctx, bm)
	xlog.Debugf("allRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
}

func TestClient_Close(t *testing.T) {
	// 订单关闭
	bm := make(gopay.BodyMap)
	bm.Set("oldreqsn", "larry01")
	// 创建订单
	resp, err := client.Close(ctx, bm)
	xlog.Debugf("allRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
}

func TestClient_NativePay(t *testing.T) {
	expire := time.Now().Add(120 * time.Minute).Format("20060102150405")
	bm := make(gopay.BodyMap)
	bm.Set("reqsn", "nativelarry01").Set("trxamt", "1").Set("body", "支付测试").Set("expiretime", expire)
	resp, err := client.NativePay(ctx, bm)
	xlog.Debugf("aliRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}

	// Output:
	// &{RspBase:{RetCode:SUCCESS RetMsg: Sign:*** Cusid:*** Appid:***} ReqSn:nativelarry01 TrxStatus:0000 ErrMsg:生成收款码成功 PayInfo:https://syb.allinpay.com/apiweb/h5unionpay/unionnative?token=** RandomStr:636549522285}
}

func TestClient_NativeClose(t *testing.T) {
	resp, err := client.NativeClose(ctx, OrderTypeReqSN, "nativelarry01")
	xlog.Debugf("aliRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}

	// Output:
	// &{RspBase:{RetCode:SUCCESS RetMsg: Sign:**** Cusid:***** Appid:*****} TrxStatus:0000 ErrMsg:交易关闭成功 RandomStr:476373212870}
}

func TestClient_QueryConfirm(t *testing.T) {
	resp, err := client.QueryConfirm(ctx, OrderTypeReqSN, "nativelarry01")
	xlog.Debugf("aliRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}

	// Output:
	// &{RspBase:{RetCode:SUCCESS RetMsg: Sign:*** Cusid:** Appid:***} Trxid:260626128079233486 ChnlTrxId: Reqsn:nativelarry01 TrxStatus:3045 Acct:opn0buM4VhGG17vZWkMLip9v1CNU TrxCode:VSP501 FinTime:20260626105619 ErrMsg:订单置关闭 RandomStr:045082197576 TrxAmt:1 Cmid:990361082416001 Chnlid: InitAmt:1 Fee: ChnlData: AcctType:99}
}
