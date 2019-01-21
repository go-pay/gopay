package gopay

import (
	"fmt"
	"log"
	"testing"
)

func TestWeChatClient_UnifiedOrder(t *testing.T) {

	//初始化微信客户端
	//    appId：应用ID
	//    mchID：商户ID
	//    secretKey：Key值
	//    isProd：是否是正式环境
	client := NewWeChatClient(appID, mchID, secretKey, true)

	//初始化参数Map
	body := make(BodyMap)
	body.Set("nonce_str", GetRandomString(32))
	body.Set("body", "测试支付")
	number := GetRandomString(32)
	log.Println("Number:", number)
	body.Set("out_trade_no", number)
	body.Set("total_fee", 10)
	body.Set("spbill_create_ip", "180.171.101.212")
	body.Set("notify_url", "http://www.igoogle.ink")
	body.Set("trade_type", TradeType_JsApi)
	//body.Set("device_info", "WEB")
	body.Set("sign_type", SignType_MD5)
	//body.Set("scene_info", `{"h5_info": {"type":"Wap","wap_url": "http://www.igoogle.ink","wap_name": "测试支付"}}`)
	body.Set("openid", openID)

	//请求支付下单，成功后得到结果
	wxRsp, err := client.UnifiedOrder(body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Response:", wxRsp)
}

func TestWeChatClient_QueryOrder(t *testing.T) {
	//初始化微信客户端
	//    appId：应用ID
	//    mchID：商户ID
	//    secretKey：Key值
	//    isProd：是否是正式环境
	client := NewWeChatClient(appID, mchID, secretKey, true)

	//初始化参数结构体
	body := make(BodyMap)
	body.Set("out_trade_no", "Osgn3y181hYfFoGvn31MM61hk0mCCpYS")
	body.Set("nonce_str", GetRandomString(32))
	body.Set("sign_type", SignType_MD5)

	//请求订单查询，成功后得到结果
	wxRsp, err := client.QueryOrder(body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Response：", wxRsp)
}

func TestWeChatClient_CloseOrder(t *testing.T) {
	//初始化微信客户端
	//    appId：应用ID
	//    mchID：商户ID
	//    secretKey：Key值
	//    isProd：是否是正式环境
	client := NewWeChatClient(appID, mchID, secretKey, true)

	//初始化参数结构体
	body := make(BodyMap)
	body.Set("out_trade_no", "Osgn3y181hYfFoGvn31MM61hk0mCCpYS")
	body.Set("nonce_str", GetRandomString(32))
	body.Set("sign_type", SignType_MD5)

	//请求订单查询，成功后得到结果
	wxRsp, err := client.CloseOrder(body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Response：", wxRsp)
}
