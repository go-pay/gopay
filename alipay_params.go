//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/13 14:42
//==================================
package gopay

type aliPayParams struct {
	NonceStr       string `xml:"nonce_str"`
	Body           string `xml:"body"`
	OutTradeNo     string `xml:"out_trade_no"`
	TotalFee       int    `xml:"total_fee"`
	SpbillCreateIp string `xml:"spbill_create_ip"`
	NotifyUrl      string `xml:"notify_url"`
	TradeType      string `xml:"trade_type"`

	DeviceInfo string `xml:"device_info"`
	SignType   string `xml:"sign_type"`
	Detail     string `xml:"detail"`
	Attach     string `xml:"attach"`
	FeeType    string `xml:"fee_type"`
	TimeStart  string `xml:"time_start"`
	TimeExpire string `xml:"time_expire"`
	GoodsTag   string `xml:"goods_tag"`
	ProductId  string `xml:"product_id"`
	LimitPay   string `xml:"limit_pay"`
	Openid     string `xml:"openid"`
	Receipt    string `xml:"receipt"`
}
