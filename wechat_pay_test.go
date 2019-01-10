package go_pay

import (
	"fmt"
	"testing"
)

func TestWXPay(t *testing.T) {
	//New一个微信支付客户端，目前isDebug参数只支持false
	client := NewWechatPayClient("wxdaa2ab9ef87b54971", "13681395021", false)

	//初始化参数结构体
	params := new(WechatParams)
	params.NonceStr = "dyUNIkNS29hvDUC1CmoF0alSdfCQGg9I"
	params.Body = "测试充值"
	params.OutTradeNo = "GYsadfjk4dhg3fkhffgnlsdkf"
	params.TotalFee = 10 //单位为分
	params.SpbillCreateIp = "127.0.0.1"
	params.NotifyUrl = "http://www.igoogle.ink"
	params.TradeType = WX_PayType_JsApi //目前只支持JSAPI有效
	params.DeviceInfo = "WEB"
	params.SignType = WX_SignType_HMAC_SHA256 //如不设置此参数，默认为 MD5
	params.Openid = "o0Df70H2Q0fY8JXh1aFPIRyOBgu81"

	//客户端设置参数
	client.SetParams(params)

	//传入secretKey获取Sign并重新设置参数
	client.GetSignAndSetReqParam("GFDS8j98rewnmgl45wHTt980jg543wmg1")

	//请求支付，成功后得到结构
	err := client.GoWechatPay()
	if err != nil {
		fmt.Println(err)
	}
	//err为空，请求支付成功后，输出请求结果
	fmt.Println(client.WXRsp)

	fmt.Println("ReturnCode：", client.WXRsp.ReturnCode)
	fmt.Println("ReturnMsg：", client.WXRsp.ReturnMsg)
	fmt.Println("Appid：", client.WXRsp.Appid)
	fmt.Println("MchId：", client.WXRsp.MchId)
	fmt.Println("DeviceInfo：", client.WXRsp.DeviceInfo)
	fmt.Println("NonceStr：", client.WXRsp.NonceStr)
	fmt.Println("Sign：", client.WXRsp.Sign)
	fmt.Println("ResultCode：", client.WXRsp.ResultCode)
	fmt.Println("PrepayId：", client.WXRsp.PrepayId)
	fmt.Println("TradeType：", client.WXRsp.TradeType)
}
