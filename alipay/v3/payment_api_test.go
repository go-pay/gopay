package alipay

import (
	"context"
	"encoding/json"
	"net/url"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
	"github.com/go-pay/util/js"
	"github.com/go-pay/xlog"
)

func TestTradePagePayAndWapPayProductCode(t *testing.T) {
	tests := []struct {
		name            string
		pay             func(ctx context.Context, bm gopay.BodyMap) (string, error)
		wantProductCode string
	}{
		{
			name:            "page pay",
			pay:             client.TradePagePay,
			wantProductCode: "FAST_INSTANT_TRADE_PAY",
		},
		{
			name:            "wap pay",
			pay:             client.TradeWapPay,
			wantProductCode: "QUICK_WAP_WAY",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bm := make(gopay.BodyMap)
			bm.Set("subject", "test order").
				Set("out_trade_no", "test-order-no").
				Set("total_amount", "0.01").
				Set("product_code", "incorrect-product-code")

			payURL, err := tt.pay(ctx, bm)
			if err != nil {
				t.Fatalf("create pay URL: %v", err)
			}

			parsedURL, err := url.Parse(payURL)
			if err != nil {
				t.Fatalf("parse pay URL: %v", err)
			}
			bizContent := make(map[string]any)
			if err = json.Unmarshal([]byte(parsedURL.Query().Get("biz_content")), &bizContent); err != nil {
				t.Fatalf("unmarshal biz_content: %v", err)
			}
			if got := bizContent["product_code"]; got != tt.wantProductCode {
				t.Fatalf("product_code = %v, want %s", got, tt.wantProductCode)
			}
		})
	}
}

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
		Set("op_app_id", "2021005143630063"). // 小程序的APPID
		Set("buyer_open_id", "018OacbttSLyJtNfdPbDOcaGoo-ncctDVT45IdYxaUsmIY8")

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
