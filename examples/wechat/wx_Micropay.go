package wechat

import (
	"fmt"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/wechat"
	"github.com/iGoogle-ink/goutil"
)

func Micropay() {
	//初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    ApiKey：Key值
	//    isProd：是否是正式环境
	client := wechat.NewClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	// 初始化参数Map
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", goutil.GetRandomString(32))
	bm.Set("body", "扫用户付款码支付")
	number := goutil.GetRandomString(32)
	fmt.Println("out_trade_no:", number)
	bm.Set("out_trade_no", number)
	bm.Set("total_fee", 1)
	bm.Set("spbill_create_ip", "127.0.0.1")
	bm.Set("auth_code", "134595229789828537")
	bm.Set("sign_type", wechat.SignType_MD5)

	sign := wechat.GetParamSign("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", bm)
	//sign, _ := gopay.GetSanBoxParamSign("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", body)

	bm.Set("sign", sign)
	//请求支付，成功后得到结果
	wxRsp, err := client.Micropay(bm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", *wxRsp)

	ok, err := wechat.VerifySign("GFDS8j98rewnmgl45wHTt980jg543abc", wechat.SignType_MD5, wxRsp)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("同步验签结果：", ok)
}
