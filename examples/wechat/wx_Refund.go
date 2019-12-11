package wechat

import (
	"fmt"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/wechat"
)

func Refund() {
	//初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    ApiKey：Key值
	//    isProd：是否是正式环境
	client := wechat.NewClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	// 初始化参数结构体
	body := make(gopay.BodyMap)
	body.Set("out_trade_no", "SdZBAqJHBQGKVwb7aMR2mUwC588NG2Sd")
	body.Set("nonce_str", gopay.GetRandomString(32))
	body.Set("sign_type", wechat.SignType_MD5)
	s := gopay.GetRandomString(64)
	fmt.Println("out_refund_no:", s)
	body.Set("out_refund_no", s)
	body.Set("total_fee", 1)
	body.Set("refund_fee", 1)
	body.Set("notify_url", "https://www.gopay.ink")

	//请求申请退款（沙箱环境下，证书路径参数可传空）
	//    body：参数Body
	//    certFilePath：cert证书路径
	//    keyFilePath：Key证书路径
	//    pkcs12FilePath：p12证书路径
	wxRsp, err := client.Refund(body, "iguiyu_cert/apiclient_cert.pem", "iguiyu_cert/apiclient_key.pem", "iguiyu_cert/apiclient_cert.p12")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("wxRsp：", *wxRsp)
}
