package wechat

import (
	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/pkg/util"
	"github.com/iGoogle-ink/gopay/pkg/xlog"
	"github.com/iGoogle-ink/gopay/wechat"
)

func Refund() {
	// client只需要初始化一个，此处为了演示，每个方法都做了初始化
	// 初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    ApiKey：Key值
	//    isProd：是否是正式环境
	client := wechat.NewClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	s := util.GetRandomString(64)
	xlog.Debug("out_refund_no:", s)
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "SdZBAqJHBQGKVwb7aMR2mUwC588NG2Sd").
		Set("nonce_str", util.GetRandomString(32)).
		Set("sign_type", wechat.SignType_MD5).
		Set("out_refund_no", s).
		Set("total_fee", 1).
		Set("refund_fee", 1).
		Set("notify_url", "https://www.fumm.cc")

	//请求申请退款（沙箱环境下，证书路径参数可传空）
	//    body：参数Body
	wxRsp, resBm, err := client.Refund(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
	xlog.Debug("resBm:", resBm)
}
