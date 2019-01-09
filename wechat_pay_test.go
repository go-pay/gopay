package go_pay

import (
	"testing"
)

func TestWXPay(t *testing.T) {
	client := NewWechatPayClient("fdgdfg", "23466", false)
	params := new(WechatParams)
	params.NonceStr = "dyUNIkNS29hvDUC1CmoF0alSdfCQGg9I"
	params.Body = "测试充值"
	params.OutTradeNo = "GYsadfjksdhgflkhfgnlsdkf"
	params.TotalFee = 15
	params.SpbillCreateIp = "127.0.0.1"
	params.NotifyUrl = "http://www.igoogle.ink"
	params.TradeType = WX_PayType_Mini
	params.DeviceInfo = "WEB"
	params.SignType = WX_SignType_HMAC_SHA256
	params.Openid = "o0Df70H2Q0fY8JXh1aFPIRyOBgu8"

	client.SetParams(params)

	client.GetSignAndSetReqParam("bfvnbhnmt5435")
	//fmt.Println("sign:", sign)

	client.GoWechatPay()
}
