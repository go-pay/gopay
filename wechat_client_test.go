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
	client := NewWeChatClient(appID, mchID, secretKey, false)

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
	client := NewWeChatClient(appID, mchID, secretKey, false)

	//初始化参数结构体
	body := make(BodyMap)
	body.Set("out_trade_no", "MfZC2segKxh0bnJSELbvKNeH3d9oWvvQ")
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
	client := NewWeChatClient(appID, mchID, secretKey, false)

	//初始化参数结构体
	body := make(BodyMap)
	body.Set("out_trade_no", "MfZC2segKxh0bnJSELbvKNeH3d9oWvvQ")
	body.Set("nonce_str", GetRandomString(32))
	body.Set("sign_type", SignType_MD5)

	//请求订单查询，成功后得到结果
	wxRsp, err := client.CloseOrder(body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Response：", wxRsp)
}

func TestWeChatClient_Refund(t *testing.T) {
	//初始化微信客户端
	//    appId：应用ID
	//    mchID：商户ID
	//    secretKey：Key值
	//    isProd：是否是正式环境
	client := NewWeChatClient(appID, mchID, secretKey, false)

	//初始化参数结构体
	body := make(BodyMap)
	body.Set("out_trade_no", "MfZC2segKxh0bnJSELbvKNeH3d9oWvvQ")
	body.Set("nonce_str", GetRandomString(32))
	body.Set("sign_type", SignType_MD5)
	s := GetRandomString(64)
	fmt.Println("s:", s)
	body.Set("out_refund_no", s)
	body.Set("total_fee", 101)
	body.Set("refund_fee", 101)

	//请求申请退款
	wxRsp, err := client.Refund(body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Response：", wxRsp)
}

func TestWeChatClient_QueryRefund(t *testing.T) {
	//初始化微信客户端
	//    appId：应用ID
	//    mchID：商户ID
	//    secretKey：Key值
	//    isProd：是否是正式环境
	client := NewWeChatClient(appID, mchID, secretKey, false)

	//初始化参数结构体
	body := make(BodyMap)
	body.Set("out_refund_no", "vk4264I1UQ3Hm3E4AKsavK8npylGSgQA092f9ckUxp8A2gXmnsLEdsupURVTcaC7")
	body.Set("nonce_str", GetRandomString(32))
	body.Set("sign_type", SignType_MD5)

	//请求申请退款
	wxRsp, err := client.QueryRefund(body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Response：", wxRsp)
}

func TestWeChatClient_DownloadBill(t *testing.T) {
	//初始化微信客户端
	//    appId：应用ID
	//    mchID：商户ID
	//    secretKey：Key值
	//    isProd：是否是正式环境
	client := NewWeChatClient(appID, mchID, secretKey, false)

	//初始化参数结构体
	body := make(BodyMap)
	body.Set("nonce_str", GetRandomString(32))
	body.Set("sign_type", SignType_MD5)
	body.Set("bill_date", "20190122")
	body.Set("bill_type", "ALL")

	//请求订单查询，成功后得到结果
	wxRsp, err := client.DownloadBill(body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Response：", wxRsp)
}

func TestWeChatClient_DownloadFundFlow(t *testing.T) {
	//初始化微信客户端
	//    appId：应用ID
	//    mchID：商户ID
	//    secretKey：Key值
	//    isProd：是否是正式环境
	client := NewWeChatClient(appID, mchID, secretKey, false)

	//初始化参数结构体
	body := make(BodyMap)
	body.Set("nonce_str", GetRandomString(32))
	body.Set("sign_type", SignType_HMAC_SHA256)
	body.Set("bill_date", "20190122")
	body.Set("account_type", "Basic")

	//请求订单查询，成功后得到结果
	wxRsp, err := client.DownloadFundFlow(body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Response：", wxRsp)
}

func TestWeChatClient_BatchQueryComment(t *testing.T) {
	//初始化微信客户端
	//    appId：应用ID
	//    mchID：商户ID
	//    secretKey：Key值
	//    isProd：是否是正式环境
	client := NewWeChatClient(appID, mchID, secretKey, false)

	//初始化参数结构体
	body := make(BodyMap)
	body.Set("nonce_str", GetRandomString(32))
	body.Set("sign_type", SignType_HMAC_SHA256)
	body.Set("begin_time", "20190120000000")
	body.Set("end_time", "20190122174000")
	body.Set("offset", "0")

	//请求订单查询，成功后得到结果
	wxRsp, err := client.BatchQueryComment(body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Response：", wxRsp)
}
