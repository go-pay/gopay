package douyin

import (
	"testing"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

func TestAppOrder(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("description", "抖音App支付测试").
		Set("out_trade_no", "OUT_"+util.RandomString(16)).
		Set("notify_url", NotifyUrl).
		Set("time_expire", time.Now().Add(15*time.Minute).Format(time.RFC3339)).
		SetBodyMap("amount", func(b gopay.BodyMap) {
			b.Set("total", 1).Set("currency", "CNY")
		})

	rsp, err := client.AppOrder(ctx, bm)
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

func TestJsapiOrder(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("description", "抖音JSAPI支付测试").
		Set("out_trade_no", "OUT_"+util.RandomString(16)).
		Set("notify_url", NotifyUrl).
		SetBodyMap("amount", func(b gopay.BodyMap) {
			b.Set("total", 1).Set("currency", "CNY")
		}).
		SetBodyMap("payer", func(b gopay.BodyMap) {
			b.Set("openid", Openid)
		})

	rsp, err := client.JsapiOrder(ctx, bm)
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

func TestH5Order(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("description", "抖音H5支付测试").
		Set("out_trade_no", "OUT_"+util.RandomString(16)).
		Set("notify_url", NotifyUrl).
		SetBodyMap("amount", func(b gopay.BodyMap) {
			b.Set("total", 1).Set("currency", "CNY")
		})

	rsp, err := client.H5Order(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("h5_url: %s", rsp.Response.H5Url)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestNativeOrder(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("description", "抖音Native支付测试").
		Set("out_trade_no", "OUT_"+util.RandomString(16)).
		Set("notify_url", NotifyUrl).
		SetBodyMap("amount", func(b gopay.BodyMap) {
			b.Set("total", 1).Set("currency", "CNY")
		})

	rsp, err := client.NativeOrder(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("code_url: %s", rsp.Response.CodeUrl)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}
