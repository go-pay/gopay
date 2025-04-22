package alipay

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay/cert"
	"github.com/go-pay/util"
	"github.com/go-pay/util/js"
	"github.com/go-pay/xlog"
)

func TestTradePrecreate(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "预创建创建订单").
		Set("out_trade_no", util.RandomString(32)).
		Set("total_amount", "0.01")

	// 创建订单
	aliRsp, err := client.TradePrecreate(ctx, bm)
	if err != nil {
		xlog.Errorf("client.TradePrecreate(), err:%v", err)
		return
	}
	xlog.Debugf("aliRsp:%s", js.MarshalString(aliRsp))

	if aliRsp.StatusCode != Success {
		xlog.Errorf("aliRsp.StatusCode:%d", aliRsp.StatusCode)
		return
	}
	xlog.Debug("aliRsp.QrCode:", aliRsp.QrCode)
	xlog.Debug("aliRsp.OutTradeNo:", aliRsp.OutTradeNo)
}

func TestTradeCreate(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "创建订单").
		Set("out_trade_no", util.RandomString(32)).
		Set("total_amount", "0.01").
		Set("product_code", "JSAPI_PAY").
		Set("op_app_id", cert.Appid).
		Set("buyer_open_id", "xxxxxx")

	// 创建订单
	aliRsp, err := client.TradeCreate(ctx, bm)
	if err != nil {
		xlog.Errorf("client.TradeCreate(), err:%v", err)
		return
	}
	xlog.Debugf("aliRsp:%s", js.MarshalString(aliRsp))

	if aliRsp.StatusCode != Success {
		xlog.Errorf("aliRsp.StatusCode:%d", aliRsp.StatusCode)
		return
	}
	xlog.Debug("aliRsp.TradeNo:", aliRsp.TradeNo)
	xlog.Debug("aliRsp.OutTradeNo:", aliRsp.OutTradeNo)
}
