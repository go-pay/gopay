package wechat

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

func Micropay() {
	// client只需要初始化一个，此处为了演示，每个方法都做了初始化
	// 初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    ApiKey：Key值
	//    isProd：是否是正式环境
	client := wechat.NewClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	number := util.RandomString(32)
	xlog.Debug("out_trade_no:", number)
	// 初始化参数Map
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("body", "扫用户付款码支付").
		Set("out_trade_no", number).
		Set("total_fee", 1).
		Set("spbill_create_ip", "127.0.0.1").
		Set("auth_code", "134595229789828537").
		Set("sign_type", wechat.SignType_MD5)

	//请求支付，成功后得到结果
	wxRsp, err := client.Micropay(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("Response：", wxRsp)

	ok, err := wechat.VerifySign("GFDS8j98rewnmgl45wHTt980jg543abc", wechat.SignType_MD5, wxRsp)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("SignOk?：", ok)
}
