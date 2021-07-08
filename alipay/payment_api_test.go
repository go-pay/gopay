package alipay

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay/cert"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xlog"
)

func TestClient_TradePrecreate(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "预创建创建订单").
		Set("out_trade_no", util.GetRandomString(32)).
		Set("total_amount", "0.01")

	// 创建订单
	aliRsp, err := client.TradePrecreate(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("aliRsp:%+v", aliRsp.Response)
	xlog.Debug("aliRsp.QrCode:", aliRsp.Response.QrCode)
	xlog.Debug("aliRsp.OutTradeNo:", aliRsp.Response.OutTradeNo)

	// 同步返回验签
	ok, err := VerifySyncSignWithCert(cert.AlipayPublicContentRSA2, aliRsp.SignData, aliRsp.Sign)
	if err != nil {
		xlog.Error(err)
	}
	xlog.Debug("同步返回验签：", ok)
}

func TestClient_TradeCreate(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "创建订单").
		Set("buyer_id", "2088802095984694").
		Set("out_trade_no", util.GetRandomString(32)).
		Set("total_amount", "0.01")

	// 创建订单
	aliRsp, err := client.TradeCreate(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
	xlog.Debug("aliRsp.TradeNo:", aliRsp.Response.TradeNo)
}

func TestClient_TradeAppPay(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "测试APP支付").
		Set("out_trade_no", "GZ201901301040355706100469").
		Set("total_amount", "1.00")

	// 手机APP支付参数请求
	payParam, err := client.TradeAppPay(bm)
	if err != nil {
		xlog.Errorf("client.TradeAppPay(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("payParam:", payParam)
}

func TestClient_TradeCancel(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "GZ201909081743431443")

	// 撤销支付订单
	aliRsp, err := client.TradeCancel(bm)
	if err != nil {
		xlog.Errorf("client.TradeCancel(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_TradeClose(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "GZ201909081743431443")

	// 条码支付
	aliRsp, err := client.TradeClose(bm)
	if err != nil {
		xlog.Errorf("client.TradeClose(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_TradePay(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "条码支付").
		Set("scene", "bar_code").
		Set("auth_code", "286248566432274952").
		Set("out_trade_no", "GZ201909081743431443").
		Set("total_amount", "0.01").
		Set("timeout_express", "2m")

	// 条码支付
	aliRsp, err := client.TradePay(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

	// 同步返回验签
	ok, err := VerifySyncSignWithCert(cert.AlipayPublicContentRSA2, aliRsp.SignData, aliRsp.Sign)
	if err != nil {
		xlog.Error(err)
	}
	xlog.Debug("同步返回验签：", ok)
}

func TestClient_TradeQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "Xdhxpe4bI5hhXAldhkMiGTZ03Jm9V6V0")

	// 查询订单
	aliRsp, err := client.TradeQuery(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("aliRsp:%+v", aliRsp.Response)

	// 同步返回验签
	ok, err := VerifySyncSignWithCert(cert.AlipayPublicContentRSA2, aliRsp.SignData, aliRsp.Sign)
	if err != nil {
		xlog.Error(err)
	}
	xlog.Debug("同步返回验签：", ok)
}

func TestClient_TradeWapPay(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "手机网站测试支付").
		Set("out_trade_no", "GZ201909081743431443").
		Set("quit_url", "https://www.fmm.ink").
		Set("total_amount", "100.00").
		Set("product_code", "QUICK_WAP_WAY")

	// 手机网站支付请求
	payUrl, err := client.TradeWapPay(bm)
	if err != nil {
		xlog.Errorf("client.TradeWapPay(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("payUrl:", payUrl)
}

func TestClient_TradePagePay(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "网站测试支付").
		Set("out_trade_no", "GZ201909081743431443").
		Set("total_amount", "88.88").
		Set("product_code", "FAST_INSTANT_TRADE_PAY")

	// 电脑网站支付请求
	payUrl, err := client.TradePagePay(bm)
	if err != nil {
		xlog.Errorf("client.TradePagePay(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("payUrl:", payUrl)
}

func TestClient_TradeRefund(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "GZ201909081743431443").
		Set("refund_amount", "5").
		Set("refund_reason", "测试退款")

	// 发起退款请求
	aliRsp, err := client.TradeRefund(bm)
	if err != nil {
		xlog.Errorf("client.TradeRefund(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_TradePageRefund(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "GZ201909081743431443").
		Set("refund_amount", "5").
		Set("out_request_no", util.GetRandomString(32))

	// 发起退款请求
	aliRsp, err := client.TradePageRefund(bm)
	if err != nil {
		xlog.Errorf("client.TradePageRefund(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_TradeFastPayRefundQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "GZ201909081743431443").
		Set("out_request_no", "GZ201909081743431443")

	// 发起退款查询请求
	aliRsp, err := client.TradeFastPayRefundQuery(bm)
	if err != nil {
		xlog.Errorf("client.TradeFastPayRefundQuery(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_TradeOrderSettle(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_request_no", "201907301518083384").
		Set("trade_no", "2019072522001484690549776067")

	var listParams []OpenApiRoyaltyDetailInfoPojo
	listParams = append(listParams, OpenApiRoyaltyDetailInfoPojo{"transfer", "2088802095984694", "userId", "userId", "2088102363632794", "0.01", "分账给2088102363632794"})

	bm.Set("royalty_parameters", listParams)
	// xlog.Debug("listParams:", bm.GetString("royalty_parameters"))

	// 发起交易结算接口
	aliRsp, err := client.TradeOrderSettle(bm)
	if err != nil {
		xlog.Errorf("client.TradeOrderSettle(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

// 订单咨询服务测试
func TestTradeAdvanceConsult(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("alipay_user_id", "2088302483540171").
		Set("consult_scene", "ORDER_RISK_EVALUATION")

	aliRsp, err := client.TradeAdvanceConsult(bm)
	if err != nil {
		xlog.Errorf("client.TradeAdvanceConsult(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}
