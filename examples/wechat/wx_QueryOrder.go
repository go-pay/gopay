package wechat

import (
	"context"

	"github.com/cedarwu/gopay"
	"github.com/cedarwu/gopay/pkg/util"
	"github.com/cedarwu/gopay/pkg/xlog"
	"github.com/cedarwu/gopay/wechat"
)

func QueryOrder() {
	// client只需要初始化一个，此处为了演示，每个方法都做了初始化
	// 初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := wechat.NewClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "GW201908091551421156").
		Set("nonce_str", util.GetRandomString(32)).
		Set("sign_type", wechat.SignType_MD5)

	// 请求订单查询，成功后得到结果
	wxRsp, resBm, err := client.QueryOrder(context.Background(), bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
	xlog.Debug("resBm:", resBm)
}
