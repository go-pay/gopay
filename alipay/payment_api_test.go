package alipay

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay/cert"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

func TestClient_TradePrecreate(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "预创建创建订单").
		Set("out_trade_no", util.RandomString(32)).
		Set("total_amount", "0.01")

	// 创建订单
	aliRsp, err := client.TradePrecreate(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
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
		Set("buyer_id", "2088722003236450").
		Set("out_trade_no", util.RandomString(32)).
		Set("total_amount", "100.10")

	// 创建订单
	aliRsp, err := client.TradeCreate(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%s, %s", bizErr.Code, bizErr.Msg)
			// do something
			return
		}
		xlog.Errorf("%s", err)
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
	payParam, err := client.TradeAppPay(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("payParam:", payParam)
}

func TestClient_TradeCancel(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "GZ201909081743431443")

	// 撤销支付订单
	aliRsp, err := client.TradeCancel(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_TradeClose(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "GZ201909081743431443")

	// 条码支付
	aliRsp, err := client.TradeClose(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
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
	aliRsp, err := client.TradePay(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
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
	const outTradeNo = "GZ201909081743431443"
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", outTradeNo)

	// 查询订单
	aliRsp, err := client.TradeQuery(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
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
	payUrl, err := client.TradeWapPay(ctx, bm)
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
	payUrl, err := client.TradePagePay(ctx, bm)
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
	aliRsp, err := client.TradeRefund(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_TradePageRefund(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "GZ201909081743431443").
		Set("refund_amount", "5").
		Set("out_request_no", util.RandomString(32))

	// 发起退款请求
	aliRsp, err := client.TradePageRefund(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
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
	aliRsp, err := client.TradeFastPayRefundQuery(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_TradeOrderSettle(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_request_no", "201907301518083384").
		Set("trade_no", "2019072522001484690549776067")

	var listParams []RoyaltyDetailInfoPojo
	listParams = append(listParams, RoyaltyDetailInfoPojo{"transfer", "2088802095984694", "userId", "userId", "2088102363632794", "0.01", "分账给2088102363632794"})

	bm.Set("royalty_parameters", listParams)
	// xlog.Debug("listParams:", bm.GetString("royalty_parameters"))

	// 发起交易结算接口
	aliRsp, err := client.TradeOrderSettle(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

// 订单咨询服务测试
func TestClient_TradeAdvanceConsult(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("alipay_user_id", "2088302483540171").
		Set("consult_scene", "ORDER_RISK_EVALUATION")

	aliRsp, err := client.TradeAdvanceConsult(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_PcreditHuabeiAuthSettleApply(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("agreement_no", "20170502000610755993")
	bm.Set("pay_amount", "3.00")
	bm.Set("out_request_no", "8077735255938032")
	bm.Set("alipay_user_id", "2088101117955611")

	aliRsp, err := client.PcreditHuabeiAuthSettleApply(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_CommerceTransportNfccardSend(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("issue_org_no", "12345678")
	bm.Set("card_no", "12345678")
	bm.Set("card_status", "CANCEL")

	aliRsp, err := client.CommerceTransportNfccardSend(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_DataDataserviceAdDataQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("query_type", "ALL_SUM")
	bm.Set("biz_token", "e09d869511189c24ce13fe3294f2bd6e")
	bm.Set("ad_level", "CREATIVE")
	bm.Set("start_date", "20180820")
	bm.Set("end_date", "20180820")
	bm.Set("outer_id_list", "10760000471-2")

	aliRsp, err := client.DataDataserviceAdDataQuery(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_CommerceAirCallcenterTradeApply(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("scene_code", "flight_ticket_order")
	bm.Set("op_code", "order_acquire_proxy")
	bm.Set("channel", "hangsiair")
	bm.Set("target_id", "2088102122001010")
	bm.Set("target_id_type", "ALIPAY_USER_ID")
	bm.SetBodyMap("trade_apply_params", func(bm gopay.BodyMap) {
		bm.Set("buyer_name", "张三")
		bm.Set("subject", "北京---上海 单程")
		bm.Set("expire_time", "2017-03-30 18:30:00")
		bm.Set("out_trade_no", "2387238273827387")
		bm.Set("total_amount", "1000.00")
		bm.Set("currency", "CNY")

	})

	aliRsp, err := client.CommerceAirCallcenterTradeApply(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_PaymentTradeOrderCreate(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("partner_id", "202210000000000001278")
	bm.Set("out_trade_no", "20150320010101001")
	bm.Set("recon_related_no", "20150320010101001")
	bm.Set("pd_code", "01050200000000000009")
	bm.Set("ev_code", "12050001")
	bm.Set("total_amount", "10000")
	bm.Set("currency_code", "156")
	bm.Set("seller_id", "2088102146225135")
	bm.Set("pay_type", "pay")
	bm.Set("pay_date", "2014-07-24 03:07:50")
	bm.SetBodyMap("goods_info", func(bm gopay.BodyMap) {
		bm.Set("goods_name", "ipad")
		bm.Set("goods_price", "2000.00")
	})

	aliRsp, err := client.PaymentTradeOrderCreate(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		xlog.Errorf("client.PaymentTradeOrderCreate(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_CommerceBenefitApply(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("activity_code", "alipay_game_marketing")
	bm.Set("trade_no", "2020081210122512120003")
	bm.Set("user_account", "342812199010013210")
	bm.Set("platform", "ios")

	aliRsp, err := client.CommerceBenefitApply(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		xlog.Errorf("client.CommerceBenefitApply(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_CommerceBenefitVerify(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("activity_code", "alipay_1212_marketing")
	bm.Set("voucher_code", "JHK51ITYK")
	bm.Set("user_account", "342812199010013210")
	bm.Set("trade_no", "2020081210122512120003")

	aliRsp, err := client.CommerceBenefitVerify(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		xlog.Errorf("client.CommerceBenefitVerify(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_TradeRepaybillQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	aliRsp, err := client.TradeRepaybillQuery(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}
