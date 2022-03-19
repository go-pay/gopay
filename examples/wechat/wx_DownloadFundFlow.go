package wechat

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/wechat"
)

func DownloadFundFlow() {
	// client只需要初始化一个，此处为了演示，每个方法都做了初始化
	// 初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    ApiKey：Key值
	//    isProd：是否是正式环境
	//    好像不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
	client := wechat.NewClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("sign_type", wechat.SignType_HMAC_SHA256).
		Set("bill_date", "20190122").
		Set("account_type", "Basic")

	// 请求下载资金账单，成功后得到结果，沙箱环境下，证书路径参数可传空
	wxRsp, err := client.DownloadFundFlow(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("Response：", wxRsp)
}
