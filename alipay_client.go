package gopay

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"time"
)

type aliPayClient struct {
	AppId           string
	aliPayPublicKey string
	privateKey      string
	ReturnUrl       string
	NotifyUrl       string
	Charset         string
	SignType        string
	isProd          bool
}

//初始化支付宝客户端
//    appId：应用ID
//    aliPayPublicKey：支付宝公钥
//    privateKey：应用私钥
//    isProd：是否是正式环境
func NewAliPayClient(appId, aliPayPublicKey, privateKey string, isProd bool) (client *aliPayClient) {
	client = new(aliPayClient)
	client.AppId = appId
	client.aliPayPublicKey = aliPayPublicKey
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
func (this *aliPayClient) AliPayTradeAppPay() {

}

//alipay.trade.wap.pay(手机网站支付接口2.0)
func (this *aliPayClient) AliPayTradeWapPay(body BodyMap) (err error) {
	var bytes []byte

	bytes, err = this.doAliPay(body, "alipay.trade.wap.pay", nil)
	if err != nil {
		//log.Println("err::", err.Error())
		return err
	}
	log.Println("Result:", string(bytes))
	return nil
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

/*
https://openapi.alipay.com/gateway.do?timestamp=2013-01-01 08:08:08&method=alipay.trade.wap.pay&app_id=1990&sign_type=RSA2&sign=ERITJKEIJKJHKKKKKKKHJEREEEEEEEEEEE&version=1.0&biz_content=
  {
    "body":"对一笔交易的具体描述信息。如果是多种商品，请将商品描述字符串累加传给body。",
    "subject":"大乐透",
    "out_trade_no":"70501111111S001111119",
    "timeout_express":"90m",
    "total_amount":9.00,
    "product_code":"QUICK_WAP_WAY"
  }
*/
//向支付宝发送请求
func (this *aliPayClient) doAliPay(body BodyMap, method string, tlsConfig ...*tls.Config) (bytes []byte, err error) {
	//===============转换body参数===================
	bodyStr, err := json.Marshal(body)
	if err != nil {
		log.Println("json.Marshal:", err)
		return nil, err
	}
	//log.Println("bodyStr:", string(bodyStr))
	//===============生成参数===================
	timeStamp := time.Now().Format(TimeLayout)
	b := new(aliPayPublicBody)
	b.AppId = this.AppId
	b.Method = method
	b.Format = "JSON"
	b.ReturnUrl = this.ReturnUrl
	b.Charset = this.Charset
	b.SignType = this.SignType
	b.Timestamp = timeStamp
	b.Version = "1.0"
	b.NotifyUrl = this.NotifyUrl
	b.BizContent = string(bodyStr)
	//===============获取签名===================

	pKey := FormatPrivateKey(this.privateKey)
	sign, err := getRsaSign(b, pKey)
	if err != nil {
		return nil, err
	}
	log.Println("rsaSign:", sign)
	b.Sign = sign
	//===============发起请求===================
	agent := HttpAgent()
	if !this.isProd {
		//沙箱环境
		agent.Post(zfb_sanbox_base_url)
	} else {
		//正式环境
		agent.Post(zfb_base_url)
	}
	//log.Println("HttpBody:", *b)
	_, bytes, errs := agent.
		Set("app_id", b.AppId).
		Set("", b.Method).
		Set("", b.AppId).
		Set("", b.AppId).
		Set("biz_content", b.BizContent).
		EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return bytes, nil
}
