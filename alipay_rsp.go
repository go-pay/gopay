//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/16 0:30
//==================================
package gopay

type AliPayTradePayAppResponse struct {
	AlipayTradeWapPayResponse AliPayInfo `json:"alipay_trade_wap_pay_response"`
	Sign                      string     `json:"sign"`
}

type AliPayInfo struct {
	Code            string `json:"code"`
	Msg             string `json:"msg"`
	SubCode         string `json:"sub_code"`
	SubMsg          string `json:"sub_msg"`
	OutTradeNo      string `json:"out_trade_no"`
	TradeNo         string `json:"trade_no"`
	TotalAmount     string `json:"total_amount"`
	SellerId        string `json:"seller_id"`
	MerchantOrderNo string `json:"merchant_order_no"`
}
