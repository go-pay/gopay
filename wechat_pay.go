package go_pay

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type WechatPayClient struct {
	Appid    string
	MchId    string
	Params   *WechatParams
	ReqParam WXReq
	isDebug  bool
}

//    Appid: 微信支付分配的公众账号ID（企业号corpid即为此appId）
//    MchId: 微信支付分配的商户号
func NewWechatPayClient(appid, mchId string, isDebug bool) *WechatPayClient {
	client := new(WechatPayClient)
	client.Appid = appid
	client.MchId = mchId
	client.isDebug = isDebug
	return client
}

//设置参数
func (this *WechatPayClient) SetParams(param *WechatParams) {
	this.Params = param
}

func (this *WechatPayClient) GetSignAndSetReqParam(secretKey string) string {
	sign, reqs := getSignAndRequestParam(this.Appid, this.MchId, secretKey, this.Params)
	this.ReqParam = reqs
	this.ReqParam.Set("sign", sign)
	return sign
}

//APP支付
func (this *WechatPayClient) GoWechatPay() error {
	//fmt.Println("reqs:", this.ReqParam)
	reqXML := this.ReqParam.generateXml()
	agent := gorequest.New()
	agent.Post(WX_PayUrl)
	agent.Type("xml")
	agent.SendString(reqXML)
	response, bytes, _ := agent.EndBytes()
	defer response.Body.Close()

	//fmt.Println("response:", response.Body)
	fmt.Println("bytes:", string(bytes))
	//fmt.Println("errors:", errors)
	return nil
}
