package wechat

import (
	"fmt"

	"github.com/iGoogle-ink/gopay"
)

func Micropay() {
	//初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    ApiKey：Key值
	//    isProd：是否是正式环境
	client := gopay.NewWeChatClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	//初始化参数Map
	body := make(gopay.BodyMap)
	body.Set("nonce_str", gopay.GetRandomString(32))
	body.Set("body", "扫用户付款码支付")
	number := gopay.GetRandomString(32)
	fmt.Println("out_trade_no:", number)
	body.Set("out_trade_no", number)
	body.Set("total_fee", 1)
	body.Set("spbill_create_ip", "127.0.0.1")
	body.Set("auth_code", "134595229789828537")
	body.Set("sign_type", gopay.SignType_MD5)

	sign := gopay.GetWeChatParamSign("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", body)
	//sign, _ := gopay.GetWeChatSanBoxParamSign("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", body)

	body.Set("sign", sign)
	//请求支付，成功后得到结果
	wxRsp, err := client.Micropay(body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", *wxRsp)

	ok, err := gopay.VerifyWeChatSign("GFDS8j98rewnmgl45wHTt980jg543abc", gopay.SignType_MD5, wxRsp)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("同步验签结果：", ok)
}
