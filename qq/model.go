package qq

const (
	// URL
	qqUnifiedOrder  = "https://qpay.qq.com/cgi-bin/pay/qpay_unified_order.cgi"              // 统一下单
	qqMicroPay      = "https://qpay.qq.com/cgi-bin/pay/qpay_micro_pay.cgi"                  // 提交付款码支付
	qqOrderQuery    = "https://qpay.qq.com/cgi-bin/pay/qpay_order_query.cgi"                // 订单查询
	qqOrderClose    = "https://qpay.qq.com/cgi-bin/pay/qpay_close_order.cgi"                // 关闭订单
	qqRefundQuery   = "https://qpay.qq.com/cgi-bin/pay/qpay_refund_query.cgi"               // 退款查询
	qqStatementDown = "https://qpay.qq.com/cgi-bin/sp_download/qpay_mch_statement_down.cgi" // 交易账单
	qqAccRoll       = "https://qpay.qq.com/cgi-bin/sp_download/qpay_mch_acc_roll.cgi"       // 资金账单
	qqReverse       = "https://api.qpay.qq.com/cgi-bin/pay/qpay_reverse.cgi"                // 撤销订单
	qqRefund        = "https://api.qpay.qq.com/cgi-bin/pay/qpay_refund.cgi"                 // 申请退款

	// 支付类型
	TradeType_MicroPay = "MICROPAY" // 提交付款码支付
	TradeType_JsApi    = "JSAPI"    // 公众号支付
	TradeType_Native   = "NATIVE"   // 原生扫码支付
	TradeType_App      = "APP"      // APP支付
	TradeType_Mini     = "MINIAPP"  // QQ小程序支付

	// 签名方式
	SignType_MD5         = "MD5"
	SignType_HMAC_SHA256 = "HMAC-SHA256"
)

type MicroPayResponse struct {
	ReturnCode     string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg      string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	RetCode        string `xml:"retcode,omitempty" json:"retcode,omitempty"`
	RetMsg         string `xml:"retmsg,omitempty" json:"retmsg,omitempty"`
	ResultCode     string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode        string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes     string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Sign           string `xml:"sign,omitempty" json:"sign,omitempty"`
	Appid          string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid       string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId          string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId       string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	NonceStr       string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	DeviceInfo     string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	TradeType      string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	TradeState     string `xml:"trade_state,omitempty" json:"trade_state,omitempty"`
	BankType       string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	FeeType        string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	TotalFee       int    `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	CashFee        int    `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	CouponFee0     int    `xml:"coupon_fee_0,omitempty" json:"coupon_fee_0,omitempty"`
	CouponFee1     int    `xml:"coupon_fee_1,omitempty" json:"coupon_fee_1,omitempty"`
	TransactionId  string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo     string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach         string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd        string `xml:"time_end,omitempty" json:"time_end,omitempty"`
	TradeStateDesc string `xml:"trade_state_desc,omitempty" json:"trade_state_desc,omitempty"`
	Openid         string `xml:"openid,omitempty" json:"openid,omitempty"`
	SubOpenid      string `xml:"sub_openid,omitempty" json:"sub_openid,omitempty"`
}
