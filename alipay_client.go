package gopay

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"log"
	"time"
)

type aliPayClient struct {
	AppId      string
	privateKey string
	ReturnUrl  string
	NotifyUrl  string
	Charset    string
	SignType   string
	isProd     bool
}

//初始化支付宝客户端
//    appId：应用ID
//    privateKey：应用私钥
//    isProd：是否是正式环境
func NewAliPayClient(appId, privateKey string, isProd bool) (client *aliPayClient) {
	client = new(aliPayClient)
	client.AppId = appId
	client.privateKey = privateKey
	client.isProd = isProd
	return client
}

//alipay.trade.fastpay.refund.query(统一收单交易退款查询)
func (this *aliPayClient) AliPayTradeFastPayRefundQuery() {

}

//alipay.trade.order.settle(统一收单交易结算接口)
func (this *aliPayClient) AliPayTradeOrderSettle() {

}

//alipay.trade.close(统一收单交易关闭接口)
func (this *aliPayClient) AliPayTradeClose() {

}

//alipay.trade.cancel(统一收单交易撤销接口)
func (this *aliPayClient) AliPayTradeCancel() {

}

//alipay.trade.refund(统一收单交易退款接口)
func (this *aliPayClient) AliPayTradeRefund() {

}

//alipay.trade.precreate(统一收单线下交易预创建)
func (this *aliPayClient) AliPayTradePrecreate() {

}

//alipay.trade.create(统一收单交易创建接口)
func (this *aliPayClient) AliPayTradeCreate() {

}

//alipay.trade.pay(统一收单交易支付接口)
func (this *aliPayClient) AliPayTradePay() {

}

//alipay.trade.query(统一收单线下交易查询)
func (this *aliPayClient) AliPayTradeQuery() {

}

//alipay.trade.app.pay(app支付接口2.0)
func (this *aliPayClient) AliPayTradeAppPay(body BodyMap) (payParam string, err error) {
	var bytes []byte
	bytes, err = this.doAliPay(body, "alipay.trade.app.pay")
	if err != nil {
		return null, err
	}
	payParam = string(bytes)
	return payParam, nil
}

//alipay.trade.wap.pay(手机网站支付接口2.0)
func (this *aliPayClient) AliPayTradeWapPay(body BodyMap) (payUrl string, err error) {
	var bytes []byte
	bytes, err = this.doAliPay(body, "alipay.trade.wap.pay")
	if err != nil {
		//log.Println("err::", err.Error())
		return null, err
	}
	payUrl = string(bytes)
	//fmt.Println("URL::", payUrl)
	return payUrl, nil
}

//alipay.trade.orderinfo.sync(支付宝订单信息同步接口)
func (this *aliPayClient) AliPayTradeOrderinfoSync() {

}

//alipay.trade.page.pay(统一收单下单并支付页面接口)
func (this *aliPayClient) AliPayTradePagePay() {

}

//zhima.credit.score.brief.get(芝麻分普惠版)
func (this *aliPayClient) ZhimaCreditScoreBriefGet() {

}

//zhima.credit.score.get(芝麻分)
func (this *aliPayClient) ZhimaCreditScoreGet() {

}

//向支付宝发送请求
func (this *aliPayClient) doAliPay(body BodyMap, method string) (bytes []byte, err error) {
	//===============转换body参数===================
	bodyStr, err := json.Marshal(body)
	if err != nil {
		log.Println("json.Marshal:", err)
		return nil, err
	}
	//fmt.Println(string(bodyStr))
	//===============生成参数===================
	reqBody := make(BodyMap)
	reqBody.Set("app_id", this.AppId)
	reqBody.Set("method", method)
	reqBody.Set("format", "JSON")
	if this.ReturnUrl != null {
		reqBody.Set("return_url", this.ReturnUrl)
	}
	if this.Charset == null {
		reqBody.Set("charset", "utf-8")
	} else {
		reqBody.Set("charset", this.Charset)
	}
	if this.SignType == null {
		reqBody.Set("sign_type", "RSA2")
	} else {
		reqBody.Set("sign_type", this.SignType)
	}
	reqBody.Set("timestamp", time.Now().Format(TimeLayout))
	reqBody.Set("version", "1.0")
	if this.NotifyUrl != null {
		reqBody.Set("notify_url", this.NotifyUrl)
	}
	reqBody.Set("biz_content", string(bodyStr))
	//===============获取签名===================
	pKey := FormatPrivateKey(this.privateKey)
	sign, err := getRsaSign(reqBody, pKey)
	if err != nil {
		return nil, err
	}
	reqBody.Set("sign", sign)
	//fmt.Println("rsaSign:", sign)
	//===============发起请求===================
	urlParam := FormatAliPayURLParam(reqBody)
	//fmt.Println("urlParam:", urlParam)
	if method == "alipay.trade.app.pay" {
		return []byte(urlParam), nil
	}
	var url string
	agent := gorequest.New()
	if !this.isProd {
		//沙箱环境
		url = zfb_sanbox_base_url
		//fmt.Println(url)
		agent.Post(url)
	} else {
		//正式环境
		url = zfb_base_url
		//fmt.Println(url)
		agent.Post(url)
	}
	rsp, b, errs := agent.
		Type("form-data").
		SendString(urlParam).
		EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if method == "alipay.trade.wap.pay" {
		//fmt.Println("rsp:::", rsp.Request.URL)
		return []byte(rsp.Request.URL.String()), nil
	}
	return b, nil
}
