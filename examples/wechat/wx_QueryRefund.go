package wechat

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

func QueryRefund() {
	// client只需要初始化一个，此处为了演示，每个方法都做了初始化
	// 初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    ApiKey：Key值
	//    isProd：是否是正式环境
	client := wechat.NewClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "97HiM5j6kGmM2fk7fYMc8MgKhPnEQ5Rk").
		//Set("out_refund_no", "vk4264I1UQ3Hm3E4AKsavK8npylGSgQA092f9ckUxp8A2gXmnsLEdsupURVTcaC7").
		//Set("transaction_id", "97HiM5j6kGmM2fk7fYMc8MgKhPnEQ5Rk").
		//Set("refund_id", "97HiM5j6kGmM2fk7fYMc8MgKhPnEQ5Rk").
		Set("nonce_str", util.RandomString(32)).
		Set("sign_type", wechat.SignType_MD5)

	//请求申请退款
	wxRsp, resBm, err := client.QueryRefund(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
	xlog.Debug("resBm:", resBm)

}
