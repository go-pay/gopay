package go_pay

import (
	"encoding/xml"
	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
)

type WeChatPayResponse struct {
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

type WeChatClient struct {
	AppId     string
	MchId     string
	secretKey string
	Params    *WeChatPayParams
	ReqParam  WeChatRequestBody
	isProd    bool
}

//初始化微信客户端
//    appId：应用ID
//    mchID：商户ID
//    isProd：是否是正式环境
//    secretKey：key，（当isProd为true时，此参数必传；false时，此参数为空）
func NewWeChatClient(appId, mchId string, isProd bool, secretKey ...string) *WeChatClient {
	client := new(WeChatClient)
	client.AppId = appId
	client.MchId = mchId
	client.isProd = isProd
	if isProd && len(secretKey) > 0 {
		client.secretKey = secretKey[0]
	}
	return client
}

//支付
func (this *WeChatClient) GoWeChatPay(param *WeChatPayParams) (wxRsp *WeChatPayResponse, err error) {
	this.Params = param
	//fmt.Println("reqs:", this.ReqParam)
	var sign string
	var reqs WeChatRequestBody
	if this.isProd {
		sign, reqs = getSignAndRequestParam(this.AppId, this.MchId, this.secretKey, this.Params)
	} else {
		return nil, errors.New("暂不支持沙箱测试")
		//getSanBoxSignString(this.Appid, this.MchId, this.Params)
	}
	this.ReqParam = reqs
	this.ReqParam.Set("sign", sign)

	reqXML := this.ReqParam.generateXml()
	agent := gorequest.New()
	if this.isProd {
		agent.Post(WX_PayUrl)
	} else {
		agent.Post(WX_PayUrl_SanBox)
	}
	agent.Type("xml")
	agent.SendString(reqXML)
	response, bytes, errs := agent.EndBytes()
	defer response.Body.Close()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	//fmt.Println("bytes:", string(bytes))
	wxRsp = new(WeChatPayResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}
