package gopay

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
	"time"
)

type aliPayClient struct {
	AppId        string
	privateKey   string
	ReturnUrl    string
	NotifyUrl    string
	Charset      string
	SignType     string
	AppAuthToken string
	isProd       bool
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
func (this *aliPayClient) AliPayTradeFastPayRefundQuery(body BodyMap) {

}

//alipay.trade.order.settle(统一收单交易结算接口)
func (this *aliPayClient) AliPayTradeOrderSettle(body BodyMap) {

}

//alipay.trade.close(统一收单交易关闭接口)
func (this *aliPayClient) AliPayTradeClose(body BodyMap) {

}

//alipay.trade.cancel(统一收单交易撤销接口)
func (this *aliPayClient) AliPayTradeCancel(body BodyMap) {

}

//alipay.trade.refund(统一收单交易退款接口)
func (this *aliPayClient) AliPayTradeRefund(body BodyMap) {

}

//alipay.trade.precreate(统一收单线下交易预创建)
func (this *aliPayClient) AliPayTradePrecreate(body BodyMap) {

}

//alipay.trade.create(统一收单交易创建接口)
func (this *aliPayClient) AliPayTradeCreate(body BodyMap) {

}

//alipay.trade.pay(统一收单交易支付接口)
func (this *aliPayClient) AliPayTradePay(body BodyMap) (aliRsp *AliPayTradePayResponse, err error) {
	var bytes []byte
	//===============product_code值===================
	//body.Set("product_code", "FACE_TO_FACE_PAYMENT")
	bytes, err = this.doAliPay(body, "alipay.trade.pay")
	if err != nil {
		return nil, err
	}

	convertBytes, _ := simplifiedchinese.GBK.NewDecoder().Bytes(bytes)
	//log.Println("convertBytes::::", string(convertBytes))
	aliRsp = new(AliPayTradePayResponse)
	err = json.Unmarshal(convertBytes, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AlipayTradePayResponse.Code != "10000" {
		info := aliRsp.AlipayTradePayResponse
		return nil, fmt.Errorf("code:%v,msg:%v,sub_code:%v,sub_msg:%v.", info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	return aliRsp, nil
}

//alipay.trade.query(统一收单线下交易查询)
func (this *aliPayClient) AliPayTradeQuery(body BodyMap) (aliRsp *AliPayTradeQueryResponse, err error) {
	var bytes []byte
	trade1 := body.Get("out_trade_no")
	trade2 := body.Get("trade_no")
	if trade1 == null && trade2 == null {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	bytes, err = this.doAliPay(body, "alipay.trade.query")
	if err != nil {
		return nil, err
	}
	convertBytes, _ := simplifiedchinese.GBK.NewDecoder().Bytes(bytes)
	//log.Println("convertBytes::::", string(convertBytes))
	aliRsp = new(AliPayTradeQueryResponse)
	err = json.Unmarshal(convertBytes, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AlipayTradePayResponse.Code != "10000" {
		info := aliRsp.AlipayTradePayResponse
		log.Println("aliRsp:", aliRsp)
		return nil, fmt.Errorf("code:%v,msg:%v,sub_code:%v,sub_msg:%v.", info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	return aliRsp, nil
}

//alipay.trade.app.pay(app支付接口2.0)
func (this *aliPayClient) AliPayTradeAppPay(body BodyMap) (payParam string, err error) {
	var bytes []byte
	//===============product_code值===================
	//body.Set("product_code", "QUICK_MSECURITY_PAY")
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
	//===============product_code值===================
	body.Set("product_code", "QUICK_WAP_WAY")
	bytes, err = this.doAliPay(body, "alipay.trade.wap.pay")
	if err != nil {
		//log.Println("err::", err.Error())
		return null, err
	}
	payUrl = string(bytes)
	//fmt.Println("URL::", payUrl)
	if payUrl == zfb_sanbox_base_url || payUrl == zfb_base_url {
		return null, errors.New("请求失败，请查看文档并检查参数")
	}
	return payUrl, nil
}

//alipay.trade.page.pay(统一收单下单并支付页面接口)
func (this *aliPayClient) AliPayTradePagePay(body BodyMap) (payUrl string, err error) {
	var bytes []byte
	//===============product_code值===================
	body.Set("product_code", "FAST_INSTANT_TRADE_PAY")
	bytes, err = this.doAliPay(body, "alipay.trade.page.pay")
	if err != nil {
		//log.Println("err::", err.Error())
		return null, err
	}
	payUrl = string(bytes)
	if payUrl == zfb_sanbox_base_url || payUrl == zfb_base_url {
		return null, errors.New("请求失败，请查看文档并检查参数")
	}
	return payUrl, nil
}

//alipay.trade.orderinfo.sync(支付宝订单信息同步接口)
func (this *aliPayClient) AliPayTradeOrderinfoSync(body BodyMap) {

}

//zhima.credit.score.brief.get(芝麻分普惠版)
func (this *aliPayClient) ZhimaCreditScoreBriefGet(body BodyMap) {

}

//zhima.credit.score.get(芝麻分)
func (this *aliPayClient) ZhimaCreditScoreGet(body BodyMap) {

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
	pubBody := make(BodyMap)
	pubBody.Set("app_id", this.AppId)
	pubBody.Set("method", method)
	pubBody.Set("format", "JSON")
	if this.ReturnUrl != null {
		pubBody.Set("return_url", this.ReturnUrl)
	}
	if this.Charset == null {
		pubBody.Set("charset", "UTF-8")
	} else {
		pubBody.Set("charset", this.Charset)
	}
	if this.SignType == null {
		pubBody.Set("sign_type", "RSA2")
	} else {
		pubBody.Set("sign_type", this.SignType)
	}
	pubBody.Set("timestamp", time.Now().Format(TimeLayout))
	pubBody.Set("version", "1.0")
	if this.NotifyUrl != null {
		pubBody.Set("notify_url", this.NotifyUrl)
	}
	if this.AppAuthToken != null {
		pubBody.Set("app_auth_token", this.AppAuthToken)
	}
	pubBody.Set("biz_content", string(bodyStr))
	//===============获取签名===================
	pKey := FormatPrivateKey(this.privateKey)
	sign, err := getRsaSign(pubBody, pubBody.Get("sign_type"), pKey)
	if err != nil {
		return nil, err
	}
	pubBody.Set("sign", sign)
	//fmt.Println("rsaSign:", sign)
	//===============发起请求===================
	urlParam := FormatAliPayURLParam(pubBody)
	//fmt.Println("urlParam:", urlParam)
	if method == "alipay.trade.app.pay" {
		return []byte(urlParam), nil
	}
	if method == "alipay.trade.page.pay" {
		if !this.isProd {
			//沙箱环境
			return []byte(zfb_sanbox_base_url + "?" + urlParam), nil
		} else {
			//正式环境
			return []byte(zfb_base_url + "?" + urlParam), nil
		}
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
	rsp, bs, errs := agent.
		Type("form-data").
		SendString(urlParam).
		EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if method == "alipay.trade.wap.pay" {
		//fmt.Println("rsp:::", rsp.Body)
		if rsp.Request.URL.String() == zfb_base_url || rsp.Request.URL.String() == zfb_sanbox_base_url {
			return nil, errors.New("请求手机网站支付出错，请检查各个参数或秘钥是否正确")
		}
		return []byte(rsp.Request.URL.String()), nil
	}
	return bs, nil
}
