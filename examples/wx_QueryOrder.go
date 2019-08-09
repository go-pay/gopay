//==================================
//  * Name：Jerry
//  * DateTime：2019/8/9 16:08
//  * Desc：
//==================================
package main

import (
	"fmt"
	"github.com/iGoogle-ink/gopay"
)

func main() {
	//初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := gopay.NewWeChatClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	//初始化参数结构体
	body := make(gopay.BodyMap)
	body.Set("out_trade_no", "GW201908091551421156")
	body.Set("nonce_str", gopay.GetRandomString(32))
	body.Set("sign_type", gopay.SignType_MD5)

	//请求订单查询，成功后得到结果
	wxRsp, err := client.QueryOrder(body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("wxRsp：", *wxRsp)
}
