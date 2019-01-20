package gopay

import (
	"encoding/xml"
	"fmt"
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
	var sign string
	body.Set("appid", this.AppId)
	body.Set("mch_id", this.MchId)
	//===============生成参数===================
	if !this.isProd {
		//沙箱环境
		body.Set("total_fee", 101)
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
	fmt.Println("req:::", reqXML)
	//===============发起请求===================
	agent := gorequest.New()
	if this.isProd {
		agent.Post(wxURL_UnifiedOrder)
	} else {
		agent.Post(wxURL_SanBox_UnifiedOrder)
	}
	agent.Type("xml")
	agent.SendString(reqXML)
	response, bytes, errs := agent.EndBytes()
	defer response.Body.Close()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	//fmt.Println("bytes:", string(bytes))
	wxRsp = new(weChatUnifiedOrderResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//查询订单
func (this *weChatClient) QueryOrder(body BodyMap) (wxRsp *weChatQueryOrderResponse, err error) {
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
		agent.Post(wxURL_OrderQuery)
	} else {
		agent.Post(wxURL_SanBox_OrderQuery)
	}
	agent.Type("xml")
	agent.SendString(reqXML)
	response, bytes, errs := agent.EndBytes()
	defer response.Body.Close()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	//fmt.Println("bytes:", string(bytes))
	wxRsp = new(weChatQueryOrderResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//关闭订单
func (this *weChatClient) CloseOrder() {

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
