package lakala

type QRCodeRsp struct {
	ReturnMsg  string `json:"return_msg,omitempty"`  //返回错误信息
	ReturnCode string `json:"return_code,omitempty"` //执行结果
	ResultCode string `json:"result_code,omitempty"` //SUCCESS表示创建订单成功，EXISTS表示订单已存在

	Channel string `json:"channel,omitempty"` //支付渠道 允许值: Alipay, Alipay+, Wechat, UnionPay

	PartnerCode string `json:"partner_code,omitempty"` //商户编码
	FullName    string `json:"full_name,omitempty"`    //商户注册全名
	PartnerName string `json:"partner_name,omitempty"` //商户名称

	OrderId        string `json:"order_id,omitempty"`         //Lakala订单ID
	PartnerOrderId string `json:"partner_order_id,omitempty"` //商户订单ID

	CodeUrl   string `json:"code_url,omitempty"`   //支付码链接，商户可以据此自行生成二维码
	QrcodeImg string `json:"qrcode_img,omitempty"` //Base64封装的二维码图片，可直接作为img的src属性
	PayUrl    string `json:"pay_url,omitempty"`    //跳转URL

}

type SdkParams struct { //调用SDK的参数字符串
	//----支付宝+----
	SchemeUrl  string `json:"scheme_url,omitempty"`
	ApplinkUrl string `json:"applink_url,omitempty"`
	NormalUrl  string `json:"normal_url,omitempty"`
	//----支付宝+----
	//----微信----
	Package   string `json:"package,omitempty"`
	Appid     string `json:"appid,omitempty"`
	Sign      string `json:"sign,omitempty"`
	Partnerid string `json:"partnerid,omitempty"`
	Prepayid  string `json:"prepayid,omitempty"`
	Noncestr  string `json:"noncestr,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	//----微信----
}
type SDKPaymentRsp struct {
	ReturnMsg string `json:"return_msg,omitempty"` //返回错误信息

	PartnerOrderId string    `json:"partner_order_id,omitempty"` //商户订单ID
	FullName       string    `json:"full_name,omitempty"`        //商户注册全名
	PartnerName    string    `json:"partner_name,omitempty"`     //商户名称
	Channel        string    `json:"channel,omitempty"`          //支付渠道   Alipay, Wechat, Alipay+
	SdkParams      SdkParams `json:"sdk_params,omitempty"`       //调用SDK的参数字符串

	ResultCode  string `json:"result_code,omitempty"`  //SUCCESS表示创建订单成功，EXISTS表示订单已存在
	PartnerCode string `json:"partner_code,omitempty"` //商户编码
	OrderId     string `json:"order_id,omitempty"`     //Lakala订单ID
	ReturnCode  string `json:"return_code,omitempty"`  //执行结果
}

// CommonApi - 申请退款1.2.0
type OrdersRefundsRsp struct {
	ReturnMsg  string `json:"return_msg,omitempty"`  //平台错误信息，订单提交失败、支付失败或已关闭时存在
	ReturnCode string `json:"return_code,omitempty"` //执行结果
	ResultCode string `json:"result_code,omitempty"` //SUCCESS表示创建订单成功，EXISTS表示订单已存在

	RefundId        string `json:"refund_id,omitempty"`         //Lakala退款单号
	PartnerRefundId string `json:"partner_refund_id,omitempty"` //商户提交的退款单号

	Amount   int64  `json:"amount,omitempty"`   //退款金额，单位是货币最小单位
	Currency string `json:"currency,omitempty"` //币种，通常为JPY

}

// CommonApi - 查询订单状态1.2.0
type OrdersRsp struct {
	ReturnMsg  string `json:"return_msg,omitempty"`  //平台错误信息，订单提交失败、支付失败或已关闭时存在
	ReturnCode string `json:"return_code,omitempty"` //执行结果
	ResultCode string `json:"result_code,omitempty"` //SUCCESS表示创建订单成功，EXISTS表示订单已存在

	OrderId        string `json:"order_id,omitempty"`         //Lakala订单ID
	PartnerOrderId string `json:"partner_order_id,omitempty"` //商户订单ID
	ChannelOrderId string `json:"channel_order_id,omitempty"` //渠道方(微信、支付宝等)流水号ID，只有支付成功时返回

	TotalFee int64 `json:"total_fee,omitempty"` //订单金额，单位是货币最小面值单位
	RealFee  int64 `json:"real_fee,omitempty"`  //实际支付金额，单位是货币最小面值单位(目前等于订单金额，为卡券预留)

	Rate float64 `json:"rate,omitempty"` //交易时使用的汇率，1JPY=?CNY，只有支付成功时返回，若渠道未提供汇率，会返回1.0

	CustomerId string `json:"customer_id,omitempty"` //客户ID，若渠道未提供则不存在

	PayTime string `json:"pay_time,omitempty"` //支付时间（yyyy-MM-dd HH:mm:ss，GMT+9），只有支付成功时返回

	CreateTime string `json:"create_time,omitempty"` //订单创建时间（最新订单为准）（yyyy-MM-dd HH:mm:ss，GMT+9）
	Currency   string `json:"currency,omitempty"`    //币种，通常为JPY
	Channel    string `json:"channel,omitempty"`     //支付渠道 Alipay|支付宝、Wechat|微信、AlipayOnline|支付宝线上、Alipay+、日系QR*、UnionPay|银联，若订单未确定渠道则不返回

	PayType          string `json:"pay_type,omitempty"`          //支付钱包类型（日系QR* /Alipay+存在）
	OrderDescription string `json:"order_description,omitempty"` //备注

	ChannelErrorCode    string `json:"channel_error_code,omitempty"`    //渠道错误码，订单提交失败、支付失败或已关闭时存在
	ChannelErrorMessage string `json:"channel_error_message,omitempty"` //渠道错误描述，订单提交失败、支付失败或已关闭时存在
	MerchantAppId       string `json:"merchant_app_id,omitempty"`       //小程序/开放平台APPID，微信小程序或微信SDK订单且支付完成时存在
	MerchantCustomerId  string `json:"merchant_customer_id,omitempty"`  //小程序关联消费者openid，微信小程序或微信SDK订单且支付完成时存在

}

// CommonApi - 付款通知
type NotifyRequest struct {
	Time           string `json:"time,omitempty"`             //UTC时间戳
	NonceStr       string `json:"nonce_str,omitempty"`        //随机字符串
	Sign           string `json:"sign,omitempty"`             //签名
	PartnerOrderId string `json:"partner_order_id,omitempty"` //商户订单ID
	OrderId        string `json:"order_id,omitempty"`         //Lakala订单ID
	ChannelOrderId string `json:"channel_order_id,omitempty"` //渠道方(微信、支付宝等)流水号ID

	TotalFee int64   `json:"total_fee,omitempty"` //订单金额，单位是货币最小面值单位
	RealFee  int64   `json:"real_fee,omitempty"`  //实际支付金额，单位是货币最小面值单位(目前等于订单金额，为卡券预留)
	Rate     float64 `json:"rate,omitempty"`      //交易时使用的汇率，1JPY=?CNY，只有支付成功时返回，若渠道未提供汇率，会返回1.0

	CustomerId string `json:"customer_id,omitempty"` //客户ID

	Currency string `json:"currency,omitempty"` //币种 (JPY/CNY)
	Channel  string `json:"channel,omitempty"`  //支付渠道 Alipay|支付宝、Wechat|微信、AlipayOnline|支付宝线上、Alipay+、日系QR*、UnionPay|银联，若订单未确定渠道则不返回

	CreateTime string `json:"create_time,omitempty"` //订单创建时间（最新订单为准）（yyyy-MM-dd HH:mm:ss，GMT+9）

	PayTime string `json:"pay_time,omitempty"` //支付时间（yyyy-MM-dd HH:mm:ss，GMT+9），只有支付成功时返回

	System    string `json:"system,omitempty"`
	PaymentId string `json:"payment_id,omitempty"`
	PayType   string `json:"pay_type,omitempty"` //支付钱包类型（日系QR* /Alipay+存在）
}
