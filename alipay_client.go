package gopay

import (
	"crypto/tls"
)

type aliPayClient struct {
	AppId     string
	MchId     string
	secretKey string
	isProd    bool
}

//初始化支付宝客户端
//    appId：应用ID
//    MchID：商户ID
//    isProd：是否是正式环境
//    secretKey：key，（当isProd为true时，此参数必传；false时，此参数为空）
func NewAlipayClient(appId, mchId string, isProd bool, secretKey ...string) *aliPayClient {
	client := new(aliPayClient)
	client.AppId = appId
	client.MchId = mchId
	client.isProd = isProd
	if isProd && len(secretKey) > 0 {
		client.secretKey = secretKey[0]
	}
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
func (this *aliPayClient) AliPayTradeWapPay() {

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
func (this *aliPayClient) doAliPay(body BodyMap, method string, tlsConfig ...*tls.Config) (bytes []byte, err error) {

	return bytes, nil
}
