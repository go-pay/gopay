package wechat

import (
	"fmt"

	"github.com/iGoogle-ink/gopay"
)

func BatchQueryComment() {
	//初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    ApiKey：Key值
	//    isProd：是否是正式环境
	//    好像不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
	client := gopay.NewWeChatClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	//初始化参数结构体
	body := make(gopay.BodyMap)
	body.Set("nonce_str", gopay.GetRandomString(32))
	body.Set("sign_type", gopay.SignType_HMAC_SHA256)
	body.Set("begin_time", "20190120000000")
	body.Set("end_time", "20190122174000")
	body.Set("offset", "0")

	//请求拉取订单评价数据，成功后得到结果，沙箱环境下，证书路径参数可传空
	wxRsp, err := client.BatchQueryComment(body, "", "", "")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response：", wxRsp)

}
