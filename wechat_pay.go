package gopay

import (
	"encoding/xml"
	"errors"
	"github.com/parnurzeal/gorequest"
)

type weChatClient struct {
	AppId     string
	MchId     string
	secretKey string
	Params    *WeChatPayParams
	isProd    bool
}

//初始化微信客户端
//    appId：应用ID
//    mchID：商户ID
//    isProd：是否是正式环境
//    secretKey：key，（当isProd为true时，此参数必传；false时，此参数为空）
func NewWeChatClient(appId, mchId string, isProd bool, secretKey ...string) *weChatClient {
	client := new(weChatClient)
	client.AppId = appId
	client.MchId = mchId
	client.isProd = isProd
	if isProd && len(secretKey) > 0 {
		client.secretKey = secretKey[0]
	}
	return client
}

//统一下单
func (this weChatClient) UnifiedOrder(param *WeChatPayParams) (wxRsp *weChatPayResponse, err error) {
	this.Params = param
	//fmt.Println("reqs:", this.ReqParam)
	var sign string
	var reqs requestBody
	if this.isProd {
		sign, reqs = getSignAndRequestBody(this.AppId, this.MchId, this.secretKey, this.Params)
	} else {
		return nil, errors.New("暂不支持沙箱测试")
		//getSanBoxSignKey(this.Appid, this.MchId, this.Params)
	}
	reqs.Set("sign", sign)

	reqXML := generateXml(reqs)
	agent := gorequest.New()
	if this.isProd {
		agent.Post(wX_PayUrl)
	} else {
		agent.Post(wX_PayUrl_SanBox)
	}
	agent.Type("xml")
	agent.SendString(reqXML)
	response, bytes, errs := agent.EndBytes()
	defer response.Body.Close()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	//fmt.Println("bytes:", string(bytes))
	wxRsp = new(weChatPayResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//查询订单
func (this weChatClient) QueryOrder() {

}

//关闭订单
func (this weChatClient) CloseOrder() {

}

//申请退款
func (this weChatClient) Refund() {

}

//查询退款
func (this weChatClient) QueryRefund() {

}

//下载对账单
func (this weChatClient) DownloadBill() {

}

//下载资金账单
func (this weChatClient) DownloadFundFlow() {

}

//拉取订单评价数据
func (this weChatClient) BatchQueryComment() {

}
