package allinpay

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
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
	xlog.Debugf("allRsp:%+v", resp)
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
