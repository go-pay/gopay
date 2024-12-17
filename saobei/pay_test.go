package saobei

import (
	"testing"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

// 小程序支付接口
func TestClient_MiniPay(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("pay_type", "010").
		Set("terminal_ip", "127.0.0.1").
		Set("terminal_trace", "larry01").
		Set("terminal_time", time.Now().Format("20060102150405")).
		Set("total_fee", "1").
		Set("sub_appid", "wx91b9fee6ce0135c9").
		Set("open_id", "oXJQK5paQaKRhgrXm_ZzF_8azJj0")

	resp, err := client.MiniPay(ctx, bm)
	xlog.Debugf("saobeiRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
}

// 付款码支付(扫码支付)
func TestClient_BarcodePay(t *testing.T) {

	terminalTrace := "larry02456"                       // 终端流水号，填写商户系统的支付订单号，不可重复
	terminalTime := time.Now().Format("20060102150405") // 终端交易时间，yyyyMMddHHmmss，全局统一时间格式

	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("pay_type", "010").
		Set("terminal_ip", "127.0.0.1").
		Set("terminal_trace", terminalTrace).
		Set("terminal_time", terminalTime).
		Set("total_fee", "1").
		Set("auth_no", "132038911197761804")

	resp, err := client.BarcodePay(ctx, bm)
	xlog.Debugf("saobeiRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}

	xlog.Debugf("terminal_trace:%s", terminalTrace)
	xlog.Debugf("terminal_time:%s", terminalTime)
}

// 支付查询
func TestClient_Query(t *testing.T) {
	// out_trade_no 和 pay_trace|pay_time 两种方式二选一

	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("pay_type", "010").
		Set("terminal_trace", "larry02456").
		Set("terminal_time", "20231008133303").
		Set("out_trade_no", "443505910021123100813330300001")
	//Set("pay_trace", "larry02456").   // terminal_trace
	//Set("pay_time", "20231008133303") // terminal_time

	resp, err := client.Query(ctx, bm)
	xlog.Debugf("saobeiRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
}

// 退款申请
func TestClient_Refund(t *testing.T) {
	// out_trade_no 和 pay_trace|pay_time 两种方式二选一

	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("pay_type", "010").
		Set("terminal_trace", "larry02456").
		Set("terminal_time", "20231008133303").
		Set("refund_fee", "1").
		Set("out_trade_no", "443505910021123100813330300001")
	//Set("pay_trace", "larry02456").   // terminal_trace
	//Set("pay_time", "20231008133303") // terminal_time

	resp, err := client.Refund(ctx, bm)
	xlog.Debugf("saobeiRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
}

// 退款订单查询
func TestClient_QueryRefund(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("pay_type", "010").
		Set("terminal_trace", "larry02456").
		Set("terminal_time", "20231008133303").
		Set("out_refund_no", "1111111")

	resp, err := client.QueryRefund(ctx, bm)
	xlog.Debugf("saobeiRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
}
