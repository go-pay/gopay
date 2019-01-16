package gopay

import (
	"fmt"
	"testing"
)

func TestWXPay(t *testing.T) {

	//初始化微信客户端
	//    appId：应用ID
	//    mchID：商户ID
	//    secretKey：Key值
	//    isProd：是否是正式环境
	client := NewWeChatClient(appID, mchID, secretKey, true)

	//初始化参数结构体
	params := new(WeChatPayParams)
	params.NonceStr = "dyUNIkNS29hvDUC1CmoF0alSdfCQGg9I"
	params.Body = "测试充值"
	params.OutTradeNo = "GYsadfjk4dhg3fk13ffgnlsdkf"
	params.TotalFee = 10 //单位为分
	params.SpbillCreateIp = "127.0.0.1"
	params.NotifyUrl = "http://www.igoogle.ink"
	params.TradeType = WX_PayType_JsApi //目前只支持JSAPI有效
	params.DeviceInfo = "WEB"
	params.SignType = WX_SignType_HMAC_SHA256 //如不设置此参数，默认为 MD5
	params.Openid = openID

	//请求支付下单，成功后得到结果
	wxRsp, err := client.UnifiedOrder(params)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("ReturnCode：", wxRsp.ReturnCode)
		fmt.Println("ReturnMsg：", wxRsp.ReturnMsg)
		fmt.Println("Appid：", wxRsp.Appid)
		fmt.Println("MchId：", wxRsp.MchId)
		fmt.Println("DeviceInfo：", wxRsp.DeviceInfo)
		fmt.Println("NonceStr：", wxRsp.NonceStr)
		fmt.Println("Sign：", wxRsp.Sign)
		fmt.Println("ResultCode：", wxRsp.ResultCode)
		fmt.Println("ErrCode：", wxRsp.ErrCode)
		fmt.Println("ErrCodeDes：", wxRsp.ErrCodeDes)
		fmt.Println("PrepayId：", wxRsp.PrepayId)
		fmt.Println("TradeType：", wxRsp.TradeType)
		fmt.Println("CodeUrl：", wxRsp.CodeUrl)
		fmt.Println("MwebUrl：", wxRsp.MwebUrl)
	}
}
