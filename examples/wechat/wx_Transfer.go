//==================================
//  * Name：Jerry
//  * DateTime：2019/8/21 21:02
//  * Desc：
//==================================
package wechat

import (
	"fmt"

	"github.com/iGoogle-ink/gopay"
)

func Transfer() {
	//初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    ApiKey：Key值
	//    isProd：是否是正式环境（企业转账到个人账户，默认正式环境）
	client := gopay.NewWeChatClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", true)

	nonceStr := gopay.GetRandomString(32)
	partnerTradeNo := gopay.GetRandomString(32)

	fmt.Println("partnerTradeNo:", partnerTradeNo)
	//初始化参数结构体
	body := make(gopay.BodyMap)
	body.Set("nonce_str", nonceStr)
	body.Set("partner_trade_no", partnerTradeNo)
	body.Set("openid", "oMlss5F06l97UpwtB-8jvZd6Yabc")
	body.Set("check_name", "FORCE_CHECK") // NO_CHECK：不校验真实姓名 , FORCE_CHECK：强校验真实姓名
	body.Set("re_user_name", "付明明")       //收款用户真实姓名。 如果check_name设置为FORCE_CHECK，则必填用户真实姓名
	body.Set("amount", 30)                //企业付款金额，单位为分
	body.Set("desc", "测试转账")              //企业付款备注，必填。注意：备注中的敏感词会被转成字符*
	body.Set("spbill_create_ip", "127.0.0.1")

	//请求申请退款（沙箱环境下，证书路径参数可传空）
	//    body：参数Body
	//    certFilePath：cert证书路径
	//    keyFilePath：Key证书路径
	//    pkcs12FilePath：p12证书路径
	wxRsp, err := client.Transfer(body, "iguiyu_cert/apiclient_cert.pem", "iguiyu_cert/apiclient_key.pem", "iguiyu_cert/apiclient_cert.p12")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("wxRsp：", *wxRsp)
}
