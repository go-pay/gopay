package saobei

import (
	"testing"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
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

	// 创建订单
	resp, err := client.MiniPay(ctx, bm)
	xlog.Debugf("saobeiRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
}

// 付款码支付(扫码支付)
func TestClient_BarcodePay(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("pay_type", "010").
		Set("terminal_ip", "127.0.0.1").
		Set("terminal_trace", "larry022").
		Set("terminal_time", time.Now().Format("20060102150405")).
		Set("total_fee", "1").
		Set("auth_no", "133159755289054986")

	// 创建订单
	resp, err := client.BarcodePay(ctx, bm)
	xlog.Debugf("saobeiRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
}
