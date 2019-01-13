package gopay

import (
	"encoding/xml"
	"github.com/parnurzeal/gorequest"
)

type weChatClient struct {
	AppId     string
	MchId     string
	secretKey string
	isProd    bool
}

//初始化微信客户端
//    appId：应用ID
//    mchID：商户ID
//    secretKey：Key值
//    isProd：是否是正式环境
func NewWeChatClient(appId, mchId, secretKey string, isProd bool) *weChatClient {
	client := new(weChatClient)
	client.AppId = appId
	client.MchId = mchId
	client.secretKey = secretKey
	client.isProd = isProd
	return client
}

//统一下单
func (this weChatClient) UnifiedOrder(param *WeChatPayParams) (wxRsp *weChatPayResponse, err error) {
	var reqs requestBody
	var sign string
	//生成下单请求参数
	if !this.isProd {
		//沙箱环境
		param.TotalFee = 101
		param.SignType = WX_SignType_MD5
		reqs = param.getRequestBody(this.AppId, this.MchId, param)
		key, err := param.getSanBoxSignKey(this.MchId, param.NonceStr, this.secretKey, param.SignType)
		if err != nil {
			return nil, err
		}
		sign = getSign(key, param.SignType, reqs)
	} else {
		reqs = param.getRequestBody(this.AppId, this.MchId, param)
		//计算Sign
		sign = getSign(this.secretKey, param.SignType, reqs)
	}

	reqs.Set("sign", sign)

	reqXML := generateXml(reqs)
	//fmt.Println("req:::", reqXML)
	agent := gorequest.New()
	if this.isProd {
		agent.Post(wxURL_unifiedOrder)
	} else {
		agent.Post(wxURL_sanbox_unifiedOrder)
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
