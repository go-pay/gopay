package wechat

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/wechat"
)

func CloseOrder() {
	// client只需要初始化一个，此处为了演示，每个方法都做了初始化
	// 初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    ApiKey：Key值
	//    isProd：是否是正式环境
	client := wechat.NewClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "MfZC2segKxh0bnJSELbvKNeH3d9oWvvQ").
		Set("nonce_str", util.RandomString(32)).
		Set("sign_type", wechat.SignType_MD5)

	// 请求关闭订单，成功后得到结果
	wxRsp, err := client.CloseOrder(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("Response：", wxRsp)
}
