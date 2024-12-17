package wechat

import (
	"testing"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

func TestClient_EntrustPublic(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("plan_id", "12535").
		Set("contract_code", "100000").
		Set("request_serial", "1000").
		Set("contract_display_account", "微信代扣").
		Set("notify_url", "https://www.igoogle.ink").
		Set("version", "1.0").
		Set("timestamp", time.Now().Unix())

	// 公众号纯签约
	wxRsp, err := client.EntrustPublic(ctx, bm)
	if err != nil {
		xlog.Errorf("client.EntrustPublic(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
}

func TestClient_EntrustAppPre(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("plan_id", "12535").
		Set("contract_code", "100000").
		Set("request_serial", "1000").
		Set("contract_display_account", "微信代扣").
		Set("notify_url", "https://www.igoogle.ink").
		Set("version", "1.0").
		Set("timestamp", time.Now().Unix())

	// APP纯签约
	wxRsp, err := client.EntrustAppPre(ctx, bm)
	if err != nil {
		xlog.Errorf("client.EntrustAppPre(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
}

func TestClient_EntrustH5(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("plan_id", "12535").
		Set("contract_code", "100000").
		Set("request_serial", "1000").
		Set("contract_display_account", "微信代扣").
		Set("notify_url", "https://www.igoogle.ink").
		Set("version", "1.0").
		Set("timestamp", time.Now().Unix()).
		Set("clientip", "127.0.0.1")

	// H5纯签约
	wxRsp, err := client.EntrustH5(ctx, bm)
	if err != nil {
		xlog.Errorf("client.EntrustH5(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
}

func TestClient_EntrustPaying(t *testing.T) {
	number := util.RandomString(32)
	xlog.Info("out_trade_no:", number)
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("contract_mchid", mchId).
		Set("contract_appid", appId).
		Set("out_trade_no", number).
		Set("nonce_str", util.RandomString(32)).
		Set("body", "测试签约").
		Set("total_fee", 1).
		Set("spbill_create_ip", "127.0.0.1").
		Set("trade_type", TradeType_App).
		Set("plan_id", "12535").
		Set("contract_code", "100000").
		Set("request_serial", "1000").
		Set("contract_display_account", "微信代扣").
		Set("notify_url", "https://www.igoogle.ink").
		Set("contract_notify_url", "https://www.igoogle.ink")

	//bm.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")

	// 支付中签约
	wxRsp, err := client.EntrustPaying(ctx, bm)
	if err != nil {
		xlog.Errorf("client.EntrustPaying(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
}
