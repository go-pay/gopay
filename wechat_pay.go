package go_pay

import (
	"encoding/xml"
	"github.com/parnurzeal/gorequest"
)

type WechatPayResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	Appid      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	DeviceInfo string `xml:"device_info"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	PrepayId   string `xml:"prepay_id"`
	TradeType  string `xml:"trade_type"`
}

type WechatPayClient struct {
	Appid    string
	MchId    string
	Params   *WechatParams
	ReqParam WXReq
	WXRsp    *WechatPayResponse
	isDebug  bool
}

//    Appid: 微信支付分配的公众账号ID（企业号corpid即为此appId）
//    MchId: 微信支付分配的商户号
func NewWechatPayClient(appid, mchId string, isDebug bool) *WechatPayClient {
	client := new(WechatPayClient)
	client.Appid = appid
	client.MchId = mchId
	client.isDebug = isDebug
	client.WXRsp = new(WechatPayResponse)
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
func (this *WechatPayClient) GoWechatPay() (err error) {
	//fmt.Println("reqs:", this.ReqParam)
	reqXML := this.ReqParam.generateXml()
	agent := gorequest.New()
	agent.Post(WX_PayUrl)
	agent.Type("xml")
	agent.SendString(reqXML)
	response, bytes, errs := agent.EndBytes()
	defer response.Body.Close()
	if len(errs) > 0 {
		return errs[0]
	}
	//fmt.Println("bytes:", string(bytes))
	err = xml.Unmarshal(bytes, this.WXRsp)
	if err != nil {
		return err
	}
	return nil
}
