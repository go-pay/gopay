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
func (this *weChatClient) UnifiedOrder(body BodyMap) (wxRsp *weChatUnifiedOrderResponse, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		bytes, err = this.doWeChat(body, wxURL_UnifiedOrder)
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_UnifiedOrder)
	}

	wxRsp = new(weChatUnifiedOrderResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//查询订单
func (this *weChatClient) QueryOrder(body BodyMap) (wxRsp *weChatQueryOrderResponse, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		bytes, err = this.doWeChat(body, wxURL_OrderQuery)
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_OrderQuery)
	}

	wxRsp = new(weChatQueryOrderResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//关闭订单
func (this *weChatClient) CloseOrder(body BodyMap) (wxRsp *weChatCloseOrderResponse, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		bytes, err = this.doWeChat(body, wxURL_CloseOrder)
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_CloseOrder)
	}

	wxRsp = new(weChatCloseOrderResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//申请退款
func (this *weChatClient) Refund() {

}

//查询退款
func (this *weChatClient) QueryRefund() {

}

//下载对账单
func (this *weChatClient) DownloadBill() {

}

//下载资金账单
func (this *weChatClient) DownloadFundFlow() {

}

//拉取订单评价数据
func (this *weChatClient) BatchQueryComment() {

}

//向微信发送请求
func (this *weChatClient) doWeChat(body BodyMap, url string) (bytes []byte, err error) {
	var sign string
	body.Set("appid", this.AppId)
	body.Set("mch_id", this.MchId)
	//===============生成参数===================
	if !this.isProd {
		//沙箱环境
		body.Set("sign_type", SignType_MD5)
		//从微信接口获取SanBoxSignKey
		key, err := getSanBoxSign(this.MchId, body.Get("nonce_str"), this.secretKey, body.Get("sign_type"))
		if err != nil {
			return nil, err
		}
		sign = getLocalSign(key, body.Get("sign_type"), body)
	} else {
		//正式环境
		//本地计算Sign
		sign = getLocalSign(this.secretKey, body.Get("sign_type"), body)
	}
	body.Set("sign", sign)

	reqXML := generateXml(body)
	//fmt.Println("req:::", reqXML)
	//===============发起请求===================
	agent := gorequest.New()
	if this.isProd {
		agent.Post(url)
	} else {
		agent.Post(url)
	}
	agent.Type("xml")
	agent.SendString(reqXML)
	_, bytes, errs := agent.EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return bytes, nil
}
