package qq

const (
	// URL
	unifiedOrder  = "https://qpay.qq.com/cgi-bin/pay/qpay_unified_order.cgi"              // 统一下单
	microPay      = "https://qpay.qq.com/cgi-bin/pay/qpay_micro_pay.cgi"                  // 提交付款码支付
	orderQuery    = "https://qpay.qq.com/cgi-bin/pay/qpay_order_query.cgi"                // 订单查询
	orderClose    = "https://qpay.qq.com/cgi-bin/pay/qpay_close_order.cgi"                // 关闭订单
	refundQuery   = "https://qpay.qq.com/cgi-bin/pay/qpay_refund_query.cgi"               // 退款查询
	statementDown = "https://qpay.qq.com/cgi-bin/sp_download/qpay_mch_statement_down.cgi" // 交易账单
	accRoll       = "https://qpay.qq.com/cgi-bin/sp_download/qpay_mch_acc_roll.cgi"       // 资金账单
	reverse       = "https://api.qpay.qq.com/cgi-bin/pay/qpay_reverse.cgi"                // 撤销订单
	refund        = "https://api.qpay.qq.com/cgi-bin/pay/qpay_refund.cgi"                 // 申请退款

	createCashRed = "https://api.qpay.qq.com/cgi-bin/hongbao/qpay_hb_mch_send.cgi"           // 创建现金红包
	redFileDown   = "https://api.qpay.qq.com/cgi-bin/hongbao/qpay_hb_mch_down_list_file.cgi" // 红包对账单下载
	queryRedInfo  = "https://qpay.qq.com/cgi-bin/mch_query/qpay_hb_mch_list_query.cgi"       // 红包详情查询

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

type NotifyRequest struct {
	Appid         string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId         string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	NonceStr      string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign          string `xml:"sign,omitempty" json:"sign,omitempty"`
	DeviceInfo    string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	TradeType     string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	TradeState    string `xml:"trade_state,omitempty" json:"trade_state,omitempty"`
	BankType      string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	FeeType       string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	TotalFee      string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	CashFee       string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	CouponFee     string `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach        string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd       string `xml:"time_end,omitempty" json:"time_end,omitempty"`
	Openid        string `xml:"openid,omitempty" json:"openid,omitempty"`
}

type MicroPayResponse struct {
	ReturnCode     string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg      string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	RetCode        string `xml:"retcode,omitempty" json:"retcode,omitempty"`
	RetMsg         string `xml:"retmsg,omitempty" json:"retmsg,omitempty"`
	Appid          string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId          string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	Sign           string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode     string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode        string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes     string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	NonceStr       string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	DeviceInfo     string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	TradeType      string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	TradeState     string `xml:"trade_state,omitempty" json:"trade_state,omitempty"`
	BankType       string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	FeeType        string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	TotalFee       string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	CashFee        string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	CouponFee      string `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`
	CouponFee0     string `xml:"coupon_fee_0,omitempty" json:"coupon_fee_0,omitempty"`
	CouponFee1     string `xml:"coupon_fee_1,omitempty" json:"coupon_fee_1,omitempty"`
	TransactionId  string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo     string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach         string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd        string `xml:"time_end,omitempty" json:"time_end,omitempty"`
	TradeStateDesc string `xml:"trade_state_desc,omitempty" json:"trade_state_desc,omitempty"`
	Openid         string `xml:"openid,omitempty" json:"openid,omitempty"`
}

type ReverseResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	RetCode    string `xml:"retcode,omitempty" json:"retcode,omitempty"`
	RetMsg     string `xml:"retmsg,omitempty" json:"retmsg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid   string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId   string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Recall     string `json:"recall,omitempty"`
}

type UnifiedOrderResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	RetCode    string `xml:"retcode,omitempty" json:"retcode,omitempty"`
	RetMsg     string `xml:"retmsg,omitempty" json:"retmsg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	TradeType  string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	PrepayId   string `xml:"prepay_id,omitempty" json:"prepay_id,omitempty"`
	CodeUrl    string `xml:"code_url,omitempty" json:"code_url,omitempty"`
}

type OrderQueryResponse struct {
	ReturnCode     string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg      string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	RetCode        string `xml:"retcode,omitempty" json:"retcode,omitempty"`
	RetMsg         string `xml:"retmsg,omitempty" json:"retmsg,omitempty"`
	Appid          string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId          string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	Sign           string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode     string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode        string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes     string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	NonceStr       string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	DeviceInfo     string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	TradeType      string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	TradeState     string `xml:"trade_state,omitempty" json:"trade_state,omitempty"`
	BankType       string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	FeeType        string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	TotalFee       string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	CashFee        string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	CouponFee      string `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`
	TransactionId  string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo     string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach         string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd        string `xml:"time_end,omitempty" json:"time_end,omitempty"`
	TradeStateDesc string `xml:"trade_state_desc,omitempty" json:"trade_state_desc,omitempty"`
	Openid         string `xml:"openid,omitempty" json:"openid,omitempty"`
}

type CloseOrderResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	RetCode    string `xml:"retcode,omitempty" json:"retcode,omitempty"`
	RetMsg     string `xml:"retmsg,omitempty" json:"retmsg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
}

type RefundResponse struct {
	ReturnCode    string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg     string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	RetCode       string `xml:"retcode,omitempty" json:"retcode,omitempty"`
	RetMsg        string `xml:"retmsg,omitempty" json:"retmsg,omitempty"`
	Appid         string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId         string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	Sign          string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode    string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode       string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes    string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	NonceStr      string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	TotalFee      string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	OutRefundNo   string `xml:"out_refund_no,omitempty" json:"out_refund_no,omitempty"`
	RefundId      string `xml:"refund_id,omitempty" json:"refund_id,omitempty"`
	RefundChannel string `xml:"refund_channel,omitempty" json:"refund_channel,omitempty"`
	RefundFee     string `xml:"refund_fee,omitempty" json:"refund_fee,omitempty"`
}

type RefundQueryResponse struct {
	ReturnCode        string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg         string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	RetCode           string `xml:"retcode,omitempty" json:"retcode,omitempty"`
	RetMsg            string `xml:"retmsg,omitempty" json:"retmsg,omitempty"`
	Appid             string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId             string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	Sign              string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode        string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode           string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes        string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	NonceStr          string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	TransactionId     string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo        string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	TotalFee          string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	CashFee           string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	FeeType           string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	OutRefundNo0      string `xml:"out_refund_no_0,omitempty" json:"out_refund_no_0,omitempty"`
	OutRefundNo1      string `xml:"out_refund_no_1,omitempty" json:"out_refund_no_1,omitempty"`
	RefundId0         string `xml:"refund_id_0,omitempty" json:"refund_id_0,omitempty"`
	RefundId1         string `xml:"refund_id_1,omitempty" json:"refund_id_1,omitempty"`
	RefundChannel0    string `xml:"refund_channel_0,omitempty" json:"refund_channel_0,omitempty"`
	RefundChannel1    string `xml:"refund_channel_1,omitempty" json:"refund_channel_1,omitempty"`
	RefundFee0        string `xml:"refund_fee_0,omitempty" json:"refund_fee_0,omitempty"`
	RefundFee1        string `xml:"refund_fee_1,omitempty" json:"refund_fee_1,omitempty"`
	CouponRefundFee0  string `xml:"coupon_refund_fee_0,omitempty" json:"coupon_refund_fee_0,omitempty"`
	CouponRefundFee1  string `xml:"coupon_refund_fee_1,omitempty" json:"coupon_refund_fee_1,omitempty"`
	CashRefundFee0    string `xml:"cash_refund_fee_0,omitempty" json:"cash_refund_fee_0,omitempty"`
	CashRefundFee1    string `xml:"cash_refund_fee_1,omitempty" json:"cash_refund_fee_1,omitempty"`
	RefundStatus0     string `xml:"refund_status_0,omitempty" json:"refund_status_0,omitempty"`
	RefundStatus1     string `xml:"refund_status_1,omitempty" json:"refund_status_1,omitempty"`
	RefundRecvAccout0 string `xml:"refund_recv_accout_0,omitempty" json:"refund_recv_accout_0,omitempty"`
	RefundRecvAccout1 string `xml:"refund_recv_accout_1,omitempty" json:"refund_recv_accout_1,omitempty"`
}

type SendCashRedResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	RetCode    string `xml:"retcode,omitempty" json:"retcode,omitempty"`
	RetMsg     string `xml:"retmsg,omitempty" json:"retmsg,omitempty"`
	Listid     string `xml:"listid,omitempty" json:"listid,omitempty"`
}

type QueryRedInfoResponse struct {
	Result      string  `xml:"result,omitempty" json:"result,omitempty"`
	ResInfo     string  `xml:"res_info,omitempty" json:"res_info,omitempty"`
	Listid      string  `xml:"listid,omitempty" json:"listid,omitempty"`
	State       string  `xml:"state,omitempty" json:"state,omitempty"`
	TotalNum    string  `xml:"total_num,omitempty" json:"total_num,omitempty"`
	RecvNum     string  `xml:"recv_num,omitempty" json:"recv_num,omitempty"`
	TotalAmount string  `xml:"total_amount,omitempty" json:"total_amount,omitempty"`
	RecvAmount  string  `xml:"recv_amount,omitempty" json:"recv_amount,omitempty"`
	RecvDetails *Detail `xml:"recv_details,omitempty" json:"recv_details,omitempty"`
}

type Detail struct {
	Uin []string `xml:"uin,omitempty" json:"uin,omitempty"`
}
