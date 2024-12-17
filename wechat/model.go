package wechat

// Notify
type NotifyRequest struct {
	ReturnCode         string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg          string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode         string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode            string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes         string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Appid              string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid           string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId              string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId           string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	DeviceInfo         string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr           string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign               string `xml:"sign,omitempty" json:"sign,omitempty"`
	SignType           string `xml:"sign_type,omitempty" json:"sign_type,omitempty"`
	Openid             string `xml:"openid,omitempty" json:"openid,omitempty"`
	IsSubscribe        string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"`
	SubOpenid          string `xml:"sub_openid,omitempty" json:"sub_openid,omitempty"`
	SubIsSubscribe     string `xml:"sub_is_subscribe,omitempty" json:"sub_is_subscribe,omitempty"`
	TradeType          string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	BankType           string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	TotalFee           string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee string `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	FeeType            string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	CashFee            string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	CashFeeType        string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty"`
	CouponFee          string `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`
	CouponCount        string `xml:"coupon_count,omitempty" json:"coupon_count,omitempty"`
	CouponType0        string `xml:"coupon_type_0,omitempty" json:"coupon_type_0,omitempty"`
	CouponType1        string `xml:"coupon_type_1,omitempty" json:"coupon_type_1,omitempty"`
	CouponType2        string `xml:"coupon_type_2,omitempty" json:"coupon_type_2,omitempty"`
	CouponId0          string `xml:"coupon_id_0,omitempty" json:"coupon_id_0,omitempty"`
	CouponId1          string `xml:"coupon_id_1,omitempty" json:"coupon_id_1,omitempty"`
	CouponId2          string `xml:"coupon_id_2,omitempty" json:"coupon_id_2,omitempty"`
	CouponFee0         string `xml:"coupon_fee_0,omitempty" json:"coupon_fee_0,omitempty"`
	CouponFee1         string `xml:"coupon_fee_1,omitempty" json:"coupon_fee_1,omitempty"`
	CouponFee2         string `xml:"coupon_fee_2,omitempty" json:"coupon_fee_2,omitempty"`
	TransactionId      string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo         string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach             string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd            string `xml:"time_end,omitempty" json:"time_end,omitempty"`
}
type UnifiedOrderResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid   string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId   string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	TradeType  string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	PrepayId   string `xml:"prepay_id,omitempty" json:"prepay_id,omitempty"`
	CodeUrl    string `xml:"code_url,omitempty" json:"code_url,omitempty"`
	MwebUrl    string `xml:"mweb_url,omitempty" json:"mweb_url,omitempty"`
}

type QueryOrderResponse struct {
	ReturnCode         string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg          string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid              string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid           string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId              string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId           string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	NonceStr           string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign               string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode         string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode            string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes         string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	DeviceInfo         string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	Openid             string `xml:"openid,omitempty" json:"openid,omitempty"`
	IsSubscribe        string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"`
	SubOpenid          string `xml:"sub_openid,omitempty" json:"sub_openid,omitempty"`
	SubIsSubscribe     string `xml:"sub_is_subscribe,omitempty" json:"sub_is_subscribe,omitempty"`
	TradeType          string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	TradeState         string `xml:"trade_state,omitempty" json:"trade_state,omitempty"`
	BankType           string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	TotalFee           string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee string `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	FeeType            string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	CashFee            string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	CashFeeType        string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty"`
	CouponFee          string `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`
	CouponCount        string `xml:"coupon_count,omitempty" json:"coupon_count,omitempty"`
	CouponType0        string `xml:"coupon_type_0,omitempty" json:"coupon_type_0,omitempty"`
	CouponType1        string `xml:"coupon_type_1,omitempty" json:"coupon_type_1,omitempty"`
	CouponType2        string `xml:"coupon_type_2,omitempty" json:"coupon_type_2,omitempty"`
	CouponId0          string `xml:"coupon_id_0,omitempty" json:"coupon_id_0,omitempty"`
	CouponId1          string `xml:"coupon_id_1,omitempty" json:"coupon_id_1,omitempty"`
	CouponId2          string `xml:"coupon_id_2,omitempty" json:"coupon_id_2,omitempty"`
	CouponFee0         string `xml:"coupon_fee_0,omitempty" json:"coupon_fee_0,omitempty"`
	CouponFee1         string `xml:"coupon_fee_1,omitempty" json:"coupon_fee_1,omitempty"`
	CouponFee2         string `xml:"coupon_fee_2,omitempty" json:"coupon_fee_2,omitempty"`
	TransactionId      string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo         string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach             string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd            string `xml:"time_end,omitempty" json:"time_end,omitempty"`
	TradeStateDesc     string `xml:"trade_state_desc,omitempty" json:"trade_state_desc,omitempty"`
}

type CloseOrderResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid   string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId   string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
}

type ReverseResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid   string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId   string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Recall     string `xml:"recall,omitempty" json:"recall,omitempty"`
}

type RefundResponse struct {
	ReturnCode          string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg           string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode          string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode             string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes          string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Appid               string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid            string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId               string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId            string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	NonceStr            string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign                string `xml:"sign,omitempty" json:"sign,omitempty"`
	TransactionId       string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo          string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	OutRefundNo         string `xml:"out_refund_no,omitempty" json:"out_refund_no,omitempty"`
	RefundId            string `xml:"refund_id,omitempty" json:"refund_id,omitempty"`
	RefundFee           string `xml:"refund_fee,omitempty" json:"refund_fee,omitempty"`
	SettlementRefundFee string `xml:"settlement_refund_fee,omitempty" json:"settlement_refund_fee,omitempty"`
	TotalFee            string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee  string `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	FeeType             string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	CashFee             string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	CashFeeType         string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty"`
	CashRefundFee       string `xml:"cash_refund_fee,omitempty" json:"cash_refund_fee,omitempty"`
	CouponType0         string `xml:"coupon_type_0,omitempty" json:"coupon_type_0,omitempty"`
	CouponType1         string `xml:"coupon_type_1,omitempty" json:"coupon_type_1,omitempty"`
	CouponType2         string `xml:"coupon_type_2,omitempty" json:"coupon_type_2,omitempty"`
	CouponRefundFee     string `xml:"coupon_refund_fee,omitempty" json:"coupon_refund_fee,omitempty"`
	CouponRefundFee0    string `xml:"coupon_refund_fee_0,omitempty" json:"coupon_refund_fee_0,omitempty"`
	CouponRefundFee1    string `xml:"coupon_refund_fee_1,omitempty" json:"coupon_refund_fee_1,omitempty"`
	CouponRefundFee2    string `xml:"coupon_refund_fee_2,omitempty" json:"coupon_refund_fee_2,omitempty"`
	CouponRefundCount   string `xml:"coupon_refund_count,omitempty" json:"coupon_refund_count,omitempty"`
	CouponRefundId0     string `xml:"coupon_refund_id_0,omitempty" json:"coupon_refund_id_0,omitempty"`
	CouponRefundId1     string `xml:"coupon_refund_id_1,omitempty" json:"coupon_refund_id_1,omitempty"`
	CouponRefundId2     string `xml:"coupon_refund_id_2,omitempty" json:"coupon_refund_id_2,omitempty"`
}

type QueryRefundResponse struct {
	ReturnCode           string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg            string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode           string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode              string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes           string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Appid                string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid             string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId                string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId             string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	NonceStr             string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign                 string `xml:"sign,omitempty" json:"sign,omitempty"`
	TotalRefundCount     string `xml:"total_refund_count,omitempty" json:"total_refund_count,omitempty"`
	TransactionId        string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo           string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	TotalFee             string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee   string `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	FeeType              string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	CashFee              string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	RefundCount          string `xml:"refund_count,omitempty" json:"refund_count,omitempty"`
	OutRefundNo0         string `xml:"out_refund_no_0,omitempty" json:"out_refund_no_0,omitempty"`
	OutRefundNo1         string `xml:"out_refund_no_1,omitempty" json:"out_refund_no_1,omitempty"`
	OutRefundNo2         string `xml:"out_refund_no_2,omitempty" json:"out_refund_no_2,omitempty"`
	RefundId0            string `xml:"refund_id_0,omitempty" json:"refund_id_0,omitempty"`
	RefundId1            string `xml:"refund_id_1,omitempty" json:"refund_id_1,omitempty"`
	RefundId2            string `xml:"refund_id_2,omitempty" json:"refund_id_2,omitempty"`
	RefundChannel0       string `xml:"refund_channel_0,omitempty" json:"refund_channel_0,omitempty"`
	RefundChannel1       string `xml:"refund_channel_1,omitempty" json:"refund_channel_1,omitempty"`
	RefundChannel2       string `xml:"refund_channel_2,omitempty" json:"refund_channel_2,omitempty"`
	RefundFee            string `xml:"refund_fee,omitempty" json:"refund_fee,omitempty"`
	RefundFee0           string `xml:"refund_fee_0,omitempty" json:"refund_fee_0,omitempty"`
	RefundFee1           string `xml:"refund_fee_1,omitempty" json:"refund_fee_1,omitempty"`
	RefundFee2           string `xml:"refund_fee_2,omitempty" json:"refund_fee_2,omitempty"`
	SettlementRefundFee0 string `xml:"settlement_refund_fee_0,omitempty" json:"settlement_refund_fee_0,omitempty"`
	SettlementRefundFee1 string `xml:"settlement_refund_fee_1,omitempty" json:"settlement_refund_fee_1,omitempty"`
	SettlementRefundFee2 string `xml:"settlement_refund_fee_2,omitempty" json:"settlement_refund_fee_2,omitempty"`
	CouponType00         string `xml:"coupon_type_0_0,omitempty" json:"coupon_type_0_0,omitempty"`
	CouponType01         string `xml:"coupon_type_0_1,omitempty" json:"coupon_type_0_1,omitempty"`
	CouponType10         string `xml:"coupon_type_1_0,omitempty" json:"coupon_type_1_0,omitempty"`
	CouponType11         string `xml:"coupon_type_1_1,omitempty" json:"coupon_type_1_1,omitempty"`
	CouponType20         string `xml:"coupon_type_2_0,omitempty" json:"coupon_type_2_0,omitempty"`
	CouponType21         string `xml:"coupon_type_2_1,omitempty" json:"coupon_type_2_1,omitempty"`
	CouponType22         string `xml:"coupon_type_2_2,omitempty" json:"coupon_type_2_2,omitempty"`
	CouponRefundFee0     string `xml:"coupon_refund_fee_0,omitempty" json:"coupon_refund_fee_0,omitempty"`
	CouponRefundFee1     string `xml:"coupon_refund_fee_1,omitempty" json:"coupon_refund_fee_1,omitempty"`
	CouponRefundFee2     string `xml:"coupon_refund_fee_2,omitempty" json:"coupon_refund_fee_2,omitempty"`
	CouponRefundCount0   string `xml:"coupon_refund_count_0,omitempty" json:"coupon_refund_count_0,omitempty"`
	CouponRefundCount1   string `xml:"coupon_refund_count_1,omitempty" json:"coupon_refund_count_1,omitempty"`
	CouponRefundCount2   string `xml:"coupon_refund_count_2,omitempty" json:"coupon_refund_count_2,omitempty"`
	CouponRefundId00     string `xml:"coupon_refund_id_0_0,omitempty" json:"coupon_refund_id_0_0,omitempty"`
	CouponRefundId01     string `xml:"coupon_refund_id_0_1,omitempty" json:"coupon_refund_id_0_1,omitempty"`
	CouponRefundId10     string `xml:"coupon_refund_id_1_0,omitempty" json:"coupon_refund_id_1_0,omitempty"`
	CouponRefundId11     string `xml:"coupon_refund_id_1_1,omitempty" json:"coupon_refund_id_1_1,omitempty"`
	CouponRefundId20     string `xml:"coupon_refund_id_2_0,omitempty" json:"coupon_refund_id_2_0,omitempty"`
	CouponRefundId21     string `xml:"coupon_refund_id_2_1,omitempty" json:"coupon_refund_id_2_1,omitempty"`
	CouponRefundId22     string `xml:"coupon_refund_id_2_2,omitempty" json:"coupon_refund_id_2_2,omitempty"`
	CouponRefundFee00    string `xml:"coupon_refund_fee_0_0,omitempty" json:"coupon_refund_fee_0_0,omitempty"`
	CouponRefundFee01    string `xml:"coupon_refund_fee_0_1,omitempty" json:"coupon_refund_fee_0_1,omitempty"`
	CouponRefundFee10    string `xml:"coupon_refund_fee_1_0,omitempty" json:"coupon_refund_fee_1_0,omitempty"`
	CouponRefundFee11    string `xml:"coupon_refund_fee_1_1,omitempty" json:"coupon_refund_fee_1_1,omitempty"`
	CouponRefundFee20    string `xml:"coupon_refund_fee_2_0,omitempty" json:"coupon_refund_fee_2_0,omitempty"`
	CouponRefundFee21    string `xml:"coupon_refund_fee_2_1,omitempty" json:"coupon_refund_fee_2_1,omitempty"`
	CouponRefundFee22    string `xml:"coupon_refund_fee_2_2,omitempty" json:"coupon_refund_fee_2_2,omitempty"`
	RefundStatus0        string `xml:"refund_status_0,omitempty" json:"refund_status_0,omitempty"`
	RefundStatus1        string `xml:"refund_status_1,omitempty" json:"refund_status_1,omitempty"`
	RefundStatus2        string `xml:"refund_status_2,omitempty" json:"refund_status_2,omitempty"`
	RefundAccount0       string `xml:"refund_account_0,omitempty" json:"refund_account_0,omitempty"`
	RefundAccount1       string `xml:"refund_account_1,omitempty" json:"refund_account_1,omitempty"`
	RefundAccount2       string `xml:"refund_account_2,omitempty" json:"refund_account_2,omitempty"`
	RefundRecvAccout0    string `xml:"refund_recv_accout_0,omitempty" json:"refund_recv_accout_0,omitempty"`
	RefundRecvAccout1    string `xml:"refund_recv_accout_1,omitempty" json:"refund_recv_accout_1,omitempty"`
	RefundRecvAccout2    string `xml:"refund_recv_accout_2,omitempty" json:"refund_recv_accout_2,omitempty"`
	RefundSuccessTime0   string `xml:"refund_success_time_0,omitempty" json:"refund_success_time_0,omitempty"`
	RefundSuccessTime1   string `xml:"refund_success_time_1,omitempty" json:"refund_success_time_1,omitempty"`
	RefundSuccessTime2   string `xml:"refund_success_time_2,omitempty" json:"refund_success_time_2,omitempty"`
}

type MicropayResponse struct {
	ReturnCode         string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg          string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid              string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid           string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId              string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId           string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	DeviceInfo         string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr           string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign               string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode         string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode            string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes         string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Openid             string `xml:"openid,omitempty" json:"openid,omitempty"`
	IsSubscribe        string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"`
	SubOpenid          string `xml:"sub_openid,omitempty" json:"sub_openid,omitempty"`
	SubIsSubscribe     string `xml:"sub_is_subscribe,omitempty" json:"sub_is_subscribe,omitempty"`
	TradeType          string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	BankType           string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	FeeType            string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	TotalFee           string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee string `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	CouponFee          string `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`
	CashFeeType        string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty"`
	CashFee            string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	TransactionId      string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo         string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach             string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd            string `xml:"time_end,omitempty" json:"time_end,omitempty"`
	PromotionDetail    string `xml:"promotion_detail,omitempty" json:"promotion_detail,omitempty"`
}

type AuthCodeToOpenIdResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid   string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId   string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	Openid     string `xml:"openid,omitempty" json:"openid,omitempty"`
}

type TransfersResponse struct {
	ReturnCode     string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg      string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	MchAppid       string `xml:"mch_appid,omitempty" json:"mch_appid,omitempty"`
	Mchid          string `xml:"mchid,omitempty" json:"mchid,omitempty"`
	DeviceInfo     string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr       string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	ResultCode     string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode        string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes     string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	PartnerTradeNo string `xml:"partner_trade_no,omitempty" json:"partner_trade_no,omitempty"`
	PaymentNo      string `xml:"payment_no,omitempty" json:"payment_no,omitempty"`
	PaymentTime    string `xml:"payment_time,omitempty" json:"payment_time,omitempty"`
}

type TransfersInfoResponse struct {
	ReturnCode     string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg      string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode     string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode        string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes     string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	PartnerTradeNo string `xml:"partner_trade_no,omitempty" json:"partner_trade_no,omitempty"`
	Appid          string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid       string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId          string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId       string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	DetailId       string `xml:"detail_id,omitempty" json:"detail_id,omitempty"`
	Status         string `xml:"status,omitempty" json:"status,omitempty"`
	Reason         string `xml:"reason,omitempty" json:"reason,omitempty"`
	Openid         string `xml:"openid,omitempty" json:"openid,omitempty"`
	TransferName   string `xml:"transfer_name,omitempty" json:"transfer_name,omitempty"`
	PaymentAmount  string `xml:"payment_amount,omitempty" json:"payment_amount,omitempty"`
	TransferTime   string `xml:"transfer_time,omitempty" json:"transfer_time,omitempty"`
	PaymentTime    string `xml:"payment_time,omitempty" json:"payment_time,omitempty"`
	Desc           string `xml:"desc,omitempty" json:"desc,omitempty"`
}

type ReportResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
}

type EntrustPublicResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ResultMsg  string `xml:"result_msg,omitempty" json:"result_msg,omitempty"`
}

type EntrustAppPreResponse struct {
	ReturnCode      string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg       string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode      string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode         string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes      string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Appid           string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid        string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId           string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId        string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	Sign            string `xml:"sign,omitempty" json:"sign,omitempty"`
	NonceStr        string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	PreEntrustwebId string `xml:"pre_entrustweb_id,omitempty" json:"pre_entrustweb_id,omitempty"`
}

type EntrustH5Response struct {
	ReturnCode  string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg   string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode  string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ResultMsg   string `xml:"result_msg,omitempty" json:"result_msg,omitempty"`
	RedirectUrl string `xml:"redirect_url,omitempty" json:"redirect_url,omitempty"`
}

type EntrustPayingResponse struct {
	ReturnCode             string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg              string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode             string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	Appid                  string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid               string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId                  string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId               string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	DeviceInfo             string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr               string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign                   string `xml:"sign,omitempty" json:"sign,omitempty"`
	ErrCode                string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes             string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	ContractResultCode     string `xml:"contract_result_code,omitempty" json:"contract_result_code,omitempty"`
	ContractErrCode        string `xml:"contract_err_code,omitempty" json:"contract_err_code,omitempty"`
	ContractErrCodeDes     string `xml:"contract_err_code_des,omitempty" json:"contract_err_code_des,omitempty"`
	PrepayId               string `xml:"prepay_id,omitempty" json:"prepay_id,omitempty"`
	TradeType              string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	CodeUrl                string `xml:"code_url,omitempty" json:"code_url,omitempty"`
	PlanId                 string `xml:"plan_id,omitempty" json:"plan_id,omitempty"`
	RequestSerial          string `xml:"request_serial,omitempty" json:"request_serial,omitempty"`
	ContractCode           string `xml:"contract_code,omitempty" json:"contract_code,omitempty"`
	ContractDisplayAccount string `xml:"contract_display_account,omitempty" json:"contract_display_account,omitempty"`
	MwebUrl                string `xml:"mweb_url,omitempty" json:"mweb_url,omitempty"`
	OutTradeNo             string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
}

type EntrustApplyPayResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
}

type EntrustDeleteResponse struct {
	ReturnCode   string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg    string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode   string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode      string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes   string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Appid        string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId        string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	Sign         string `xml:"sign,omitempty" json:"sign,omitempty"`
	PlanId       string `xml:"plan_id,omitempty" json:"plan_id,omitempty"`
	ContractId   string `xml:"contract_id,omitempty" json:"contract_id,omitempty"`
	ContractCode string `xml:"contract_code,omitempty" json:"contract_code,omitempty"`
}

type EntrustQueryResponse struct {
	ReturnCode                string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg                 string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode                string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode                   string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes                string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Appid                     string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId                     string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	Sign                      string `xml:"sign,omitempty" json:"sign,omitempty"`
	RequestSerial             string `xml:"request_serial,omitempty" json:"request_serial,omitempty"`
	PlanId                    string `xml:"plan_id,omitempty" json:"plan_id,omitempty"`
	ContractId                string `xml:"contract_id,omitempty" json:"contract_id,omitempty"`
	ContractCode              string `xml:"contract_code,omitempty" json:"contract_code,omitempty"`
	ContractDisplayAccount    string `xml:"contract_display_account,omitempty" json:"contract_display_account,omitempty"`
	ContractState             string `xml:"contract_state,omitempty" json:"contract_state,omitempty"`
	ContractSignedTime        string `xml:"contract_signed_time,omitempty" json:"contract_signed_time,omitempty"`
	ContractExpiredTime       string `xml:"contract_expired_time,omitempty" json:"contract_expired_time,omitempty"`
	ContractTerminatedTime    string `xml:"contract_terminated_time,omitempty" json:"contract_terminated_time,omitempty"`
	ContractTerminationMode   string `xml:"contract_termination_mode,omitempty" json:"contract_termination_mode,omitempty"`
	ContractTerminationRemark string `xml:"contract_termination_remark,omitempty" json:"contract_termination_remark,omitempty"`
	Openid                    string `xml:"openid,omitempty" json:"openid,omitempty"`
}

type getSignKeyResponse struct {
	ReturnCode     string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg      string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	MchId          string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SandboxSignkey string `xml:"sandbox_signkey,omitempty" json:"sandbox_signkey,omitempty"`
}

type RefundNotifyRequest struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid   string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId   string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	ReqInfo    string `xml:"req_info,omitempty" json:"req_info,omitempty"`
}

type RefundNotify struct {
	TransactionId       string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo          string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	RefundId            string `xml:"refund_id,omitempty" json:"refund_id,omitempty"`
	OutRefundNo         string `xml:"out_refund_no,omitempty" json:"out_refund_no,omitempty"`
	TotalFee            string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee  string `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	RefundFee           string `xml:"refund_fee,omitempty" json:"refund_fee,omitempty"`
	SettlementRefundFee string `xml:"settlement_refund_fee,omitempty" json:"settlement_refund_fee,omitempty"`
	RefundStatus        string `xml:"refund_status,omitempty" json:"refund_status,omitempty"`
	SuccessTime         string `xml:"success_time,omitempty" json:"success_time,omitempty"`
	RefundRecvAccout    string `xml:"refund_recv_accout,omitempty" json:"refund_recv_accout,omitempty"`
	RefundAccount       string `xml:"refund_account,omitempty" json:"refund_account,omitempty"`
	RefundRequestSource string `xml:"refund_request_source,omitempty" json:"refund_request_source,omitempty"`
}

type PaidUnionId struct {
	Unionid string `json:"unionid,omitempty"` // 用户在开放平台的唯一标识符
	Errcode int    `json:"errcode,omitempty"` // 错误码
	Errmsg  string `json:"errmsg,omitempty"`  // 错误信息
}

type AccessToken struct {
	AccessToken string `json:"access_token,omitempty"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in,omitempty"`   // SessionKey超时时间（秒）
	Errcode     int    `json:"errcode,omitempty"`      // 错误码
	Errmsg      string `json:"errmsg,omitempty"`       // 错误信息
}

// 微信开放平台用户信息
type Oauth2UserInfo struct {
	Openid     string   `json:"openid,omitempty"`     // 普通用户的标识，对当前开发者帐号唯一
	Nickname   string   `json:"nickname,omitempty"`   // 普通用户昵称
	Sex        int      `json:"sex,omitempty"`        // 普通用户性别，1为男性，2为女性
	City       string   `json:"city,omitempty"`       // 普通用户个人资料填写的城市
	Province   string   `json:"province,omitempty"`   // 普通用户个人资料填写的省份
	Country    string   `json:"country,omitempty"`    // 国家，如中国为CN
	Headimgurl string   `json:"headimgurl,omitempty"` // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	Privilege  []string `json:"privilege,omitempty"`  // 用户特权信息，json数组，如微信沃卡用户为（chinaunicom）
	Unionid    string   `json:"unionid,omitempty"`    // 用户统一标识。针对一个微信开放平台帐号下的应用，同一用户的unionid是唯一的。
}

// 微信公众号用户信息
type PublicUserInfo struct {
	Subscribe      int    `json:"subscribe,omitempty"`       // 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	Openid         string `json:"openid,omitempty"`          // 用户唯一标识
	Nickname       string `json:"nickname,omitempty"`        // 用户的昵称
	Sex            int    `json:"sex,omitempty"`             // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	City           string `json:"city,omitempty"`            // 用户所在城市
	Province       string `json:"province,omitempty"`        // 用户所在省份
	Country        string `json:"country,omitempty"`         // 用户所在国家
	Language       string `json:"language,omitempty"`        // 用户的语言，简体中文为zh_CN
	Headimgurl     string `json:"headimgurl,omitempty"`      // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	SubscribeTime  int    `json:"subscribe_time,omitempty"`  // 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	Unionid        string `json:"unionid,omitempty"`         // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	Remark         string `json:"remark,omitempty"`          // 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	Groupid        int    `json:"groupid,omitempty"`         // 用户所在的分组ID（兼容旧的用户分组接口）
	TagidList      []int  `json:"tagid_list,omitempty"`      // 用户被打上的标签ID列表
	SubscribeScene string `json:"subscribe_scene,omitempty"` // 返回用户关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENEPROFILE LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_OTHERS 其他
	QrScene        int    `json:"qr_scene,omitempty"`        // 二维码扫码场景（开发者自定义）
	QrSceneStr     string `json:"qr_scene_str,omitempty"`    // 二维码扫码场景描述（开发者自定义）
	Errcode        int    `json:"errcode,omitempty"`         // 错误码
	Errmsg         string `json:"errmsg,omitempty"`          // 错误信息
}

type PublicOpenids struct {
	UserList []*struct {
		Openid string `json:"openid"`
		Lang   string `json:"lang"`
	} `json:"user_list"`
}

type PublicUserInfoBatch struct {
	UserInfoList []*PublicUserInfo `json:"user_info_list"`
}

// 微信小程序解密后 用户手机号结构体
type UserPhone struct {
	PhoneNumber     string         `json:"phoneNumber,omitempty"`
	PurePhoneNumber string         `json:"purePhoneNumber,omitempty"`
	CountryCode     string         `json:"countryCode,omitempty"`
	Watermark       *watermarkInfo `json:"watermark,omitempty"`
}

// 微信小程序解密后 用户信息结构体
type AppletUserInfo struct {
	OpenId    string         `json:"openId,omitempty"`
	NickName  string         `json:"nickName,omitempty"`
	Gender    int            `json:"gender,omitempty"`
	City      string         `json:"city,omitempty"`
	Province  string         `json:"province,omitempty"`
	Country   string         `json:"country,omitempty"`
	AvatarUrl string         `json:"avatarUrl,omitempty"`
	UnionId   string         `json:"unionId,omitempty"`
	Watermark *watermarkInfo `json:"watermark,omitempty"`
}

type watermarkInfo struct {
	Appid     string `json:"appid,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
}

// 授权码查询openid 返回
type OpenIdByAuthCodeRsp struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid   string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId   string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	Openid     string `xml:"openid,omitempty" json:"openid,omitempty"` // 用户唯一标识
}

// 获取开放平台，access_token 返回结构体
type Oauth2AccessToken struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Openid       string `json:"openid,omitempty"`
	Scope        string `json:"scope,omitempty"`
	Unionid      string `json:"unionid,omitempty"`
	Errcode      int    `json:"errcode,omitempty"` // 错误码
	Errmsg       string `json:"errmsg,omitempty"`  // 错误信息
}

type CheckAccessTokenRsp struct {
	Errcode int    `json:"errcode,omitempty"` // 错误码
	Errmsg  string `json:"errmsg,omitempty"`  // 错误信息
}

// ProfitSharingResponse 请求分账返回结果
type ProfitSharingResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"` // 返回状态码 SUCCESS/FAIL 此字段是通信标识，非交易标识
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`   // 返回信息，如非空，为错误原因
	//以下字段在return_code为SUCCESS的时候有返回
	ResultCode    string `xml:"result_code,omitempty" json:"result_code,omitempty"`   // 业务结果 SUCCESS：分账申请接收成功，结果通过分账查询接口查询 FAIL ：提交业务失败
	ErrCode       string `xml:"err_code,omitempty" json:"err_code,omitempty"`         // 错误代码
	ErrCodeDes    string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"` // 错误代码描述
	Appid         string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid      string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId         string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId      string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	NonceStr      string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`           // 随机字符串
	Sign          string `xml:"sign,omitempty" json:"sign,omitempty"`                     // 签名
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"` // 微信订单号
	OutOrderNo    string `xml:"out_order_no,omitempty" json:"out_order_no,omitempty"`     // 商户分账单号
	OrderId       string `xml:"order_id,omitempty" json:"order_id,omitempty"`             // 微信分账单号
}

// ProfitSharingQueryResponse 查询分账结果
type ProfitSharingQueryResponse struct {
	ReturnCode    string `xml:"return_code,omitempty" json:"return_code,omitempty"`       // 返回状态码 SUCCESS/FAIL 此字段是通信标识，非交易标识
	ReturnMsg     string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`         // 返回信息，如非空，为错误原因
	ResultCode    string `xml:"result_code,omitempty" json:"result_code,omitempty"`       // 业务结果 SUCCESS：分账申请接收成功，结果通过分账查询接口查询 FAIL ：提交业务失败
	ErrCode       string `xml:"err_code,omitempty" json:"err_code,omitempty"`             // 错误代码
	ErrCodeDes    string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`     // 错误代码描述
	MchId         string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`                 // 商户号
	NonceStr      string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`           // 随机字符串
	Sign          string `xml:"sign,omitempty" json:"sign,omitempty"`                     // 签名
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"` // 微信订单号
	OutOrderNo    string `xml:"out_order_no,omitempty" json:"out_order_no,omitempty"`     // 商户分账单号
	OrderId       string `xml:"order_id,omitempty" json:"order_id,omitempty"`             // 微信分账单号
	Status        string `xml:"status,omitempty" json:"status,omitempty"`                 // 分账单状态 ACCEPTED—受理成功 PROCESSING—处理中 FINISHED—处理完成 CLOSED—处理失败，已关单
	CloseReason   string `xml:"close_reason,omitempty" json:"close_reason,omitempty"`     // 关单原因 NO_AUTH:分账授权已解除
	Receivers     string `xml:"receivers,omitempty" json:"receivers,omitempty"`
}

// ProfitSharingAddReceiverResponse 添加分账接收者结果
type ProfitSharingAddReceiverResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"` // 返回状态码 SUCCESS/FAIL 此字段是通信标识，非交易标识
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`   // 返回信息，如非空，为错误原因
	//以下字段在return_code为SUCCESS的时候有返回
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`   // 业务结果 SUCCESS：分账申请接收成功，结果通过分账查询接口查询 FAIL ：提交业务失败
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`         // 错误代码
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"` // 错误代码描述
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid   string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId   string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"` // 随机字符串
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`           // 签名
	Receiver   string `xml:"receiver,omitempty" json:"receiver,omitempty"`   // 接收方
}

// ProfitSharingReturnResponse 分账退回响应结果
type ProfitSharingReturnResponse struct {
	ReturnCode        string `xml:"return_code,omitempty" json:"return_code,omitempty"`   // 返回状态码 SUCCESS/FAIL 此字段是通信标识，非交易标识
	ErrCode           string `xml:"err_code,omitempty" json:"err_code,omitempty"`         // 错误代码
	ErrorMsg          string `xml:"error_msg,omitempty" json:"error_msg,omitempty"`       // 返回信息 如果返回状态码为FAIL，则本字段存在，且为失败的错误信息
	ErrCodeDes        string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"` // 错误代码描述
	Appid             string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid          string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId             string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId          string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`                   // 调用接口提供的公众账号ID
	NonceStr          string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`                     // 随机字符串
	Sign              string `xml:"sign,omitempty" json:"sign,omitempty"`                               // 签名
	OrderId           string `xml:"order_id,omitempty" json:"order_id,omitempty"`                       // 微信分账单号
	OutOrderNo        string `xml:"out_order_no,omitempty" json:"out_order_no,omitempty"`               // 商户分账单号
	OutReturnNo       string `xml:"out_return_no,omitempty" json:"out_return_no,omitempty"`             // 商户回退单号 调用接口提供的商户系统内部的回退单号
	ReturnNo          string `xml:"return_no,omitempty" json:"return_no,omitempty"`                     // 微信回退单号 微信分账回退单号，微信系统返回的唯一标识
	ReturnAccountType string `xml:"return_account_type,omitempty" json:"return_account_type,omitempty"` // 回退方类型
	ReturnAccount     string `xml:"return_account,omitempty" json:"return_account,omitempty"`           // 回退方账号
	ReturnAmount      string `xml:"return_amount,omitempty" json:"return_amount,omitempty"`             // 回退金额
	Description       string `xml:"description,omitempty" json:"description,omitempty"`                 // 退回描述
	Result            string `xml:"result,omitempty" json:"result,omitempty"`                           // 退回结果
	FailReason        string `xml:"fail_reason,omitempty" json:"fail_reason,omitempty"`                 // 失败原因
	FinishTime        string `xml:"finish_time,omitempty" json:"finish_time,omitempty"`                 // 完成时间
}

// ProfitSharingOrderAmountQueryResponse 查询订单待分账金额响应结果
type ProfitSharingOrderAmountQueryResponse struct {
	ReturnCode    string `xml:"return_code,omitempty" json:"return_code,omitempty"`       // 返回状态码 SUCCESS/FAIL 此字段是通信标识，非交易标识
	ErrCode       string `xml:"err_code,omitempty" json:"err_code,omitempty"`             // 错误代码
	ErrorMsg      string `xml:"error_msg,omitempty" json:"error_msg,omitempty"`           // 返回信息 如果返回状态码为FAIL，则本字段存在，且为失败的错误信息
	ErrCodeDes    string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`     // 错误代码描述
	MchId         string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`                 //调用接口时提供的服务商户号
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"` //微信支付订单号
	UnsplitAmount string `xml:"unsplit_amount,omitempty" json:"unsplit_amount,omitempty"` //订单剩余待分金额，整数，单位为分
	NonceStr      string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`           //微信返回的随机字符串
	Sign          string `xml:"sign,omitempty" json:"sign,omitempty"`                     //微信返回的签名
}

// ProfitSharingMerchantRatioQuery 分账退回响应结果
type ProfitSharingMerchantRatioQuery struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`   // 返回状态码 SUCCESS/FAIL 此字段是通信标识，非交易标识
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`         // 错误代码
	ErrorMsg   string `xml:"error_msg,omitempty" json:"error_msg,omitempty"`       // 返回信息 如果返回状态码为FAIL，则本字段存在，且为失败的错误信息
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"` // 错误代码描述
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`             //调用接口时提供的服务商户号
	SubMchId   string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`     //微信支付分配的子商户号，即分账的出资商户号。查询子商户号的设置的最大分账比例（普通分账）时返回此字段
	BrandMchId string `xml:"brand_mch_id,omitempty" json:"brand_mch_id,omitempty"` //调用接口时提供的品牌主商户号。查询品牌主商户设置的全局分账比例（品牌分账）时返回此字段。
	MaxRatio   string `xml:"max_ratio,omitempty" json:"max_ratio,omitempty"`       //子商户允许服务商分账的最大比例，单位万分比，比如2000表示20%
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`       //微信返回的随机字符串
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`                 //微信返回的签名
}

type PayBankResponse struct {
	ReturnCode     string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg      string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode     string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode        string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes     string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	MchId          string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	PartnerTradeNo string `xml:"partner_trade_no,omitempty" json:"partner_trade_no,omitempty"`
	Amount         string `xml:"amount,omitempty" json:"amount,omitempty"`
	NonceStr       string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign           string `xml:"sign,omitempty" json:"sign,omitempty"`
	PaymentNo      string `xml:"payment_no,omitempty" json:"payment_no,omitempty"`
	CmmsAmt        string `xml:"cmms_amt,omitempty" json:"cmms_amt,omitempty"`
}

type QueryBankResponse struct {
	ReturnCode     string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg      string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode     string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode        string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes     string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	MchId          string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	PartnerTradeNo string `xml:"partner_trade_no,omitempty" json:"partner_trade_no,omitempty"`
	PaymentNo      string `xml:"payment_no,omitempty" json:"payment_no,omitempty"`
	BankNoMd5      string `xml:"bank_no_md5,omitempty" json:"bank_no_md5,omitempty"`
	TrueNameMd5    string `xml:"true_name_md5,omitempty" json:"true_name_md5,omitempty"`
	Amount         string `xml:"amount,omitempty" json:"amount,omitempty"`
	Status         string `xml:"status,omitempty" json:"status,omitempty"`
	CmmsAmt        string `xml:"cmms_amt,omitempty" json:"cmms_amt,omitempty"`
	CreateTime     string `xml:"create_time,omitempty" json:"create_time,omitempty"`
	PaySuccTime    string `xml:"pay_succ_time,omitempty" json:"pay_succ_time,omitempty"`
	Reason         string `xml:"reason,omitempty" json:"reason,omitempty"`
}

type RSAPublicKeyResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	PubKey     string `xml:"pub_key,omitempty" json:"pub_key,omitempty"`
}

type SendCashRedResponse struct {
	ReturnCode  string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg   string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode  string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode     string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes  string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	MchBillno   string `xml:"mch_billno,omitempty" json:"mch_billno,omitempty"`
	MchId       string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	Wxappid     string `xml:"wxappid,omitempty" json:"wxappid,omitempty"`
	ReOpenid    string `xml:"re_openid,omitempty" json:"re_openid,omitempty"`
	TotalAmount string `xml:"total_amount,omitempty" json:"total_amount,omitempty"`
	SendListid  string `xml:"send_listid,omitempty" json:"send_listid,omitempty"`
}

type SendAppletRedResponse struct {
	ReturnCode  string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg   string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode  string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode     string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes  string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	MchBillno   string `xml:"mch_billno,omitempty" json:"mch_billno,omitempty"`
	MchId       string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	Wxappid     string `xml:"wxappid,omitempty" json:"wxappid,omitempty"`
	ReOpenid    string `xml:"re_openid,omitempty" json:"re_openid,omitempty"`
	TotalAmount string `xml:"total_amount,omitempty" json:"total_amount,omitempty"`
	SendListid  string `xml:"send_listid,omitempty" json:"send_listid,omitempty"`
	Packages    string `xml:"package,omitempty" json:"package,omitempty"`
}

type QueryRedRecordResponse struct {
	ReturnCode   string  `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg    string  `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode   string  `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode      string  `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes   string  `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	MchBillno    string  `xml:"mch_billno,omitempty" json:"mch_billno,omitempty"`
	MchId        string  `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	DetailId     string  `xml:"detail_id,omitempty" json:"detail_id,omitempty"`
	Status       string  `xml:"status,omitempty" json:"status,omitempty"`
	SendType     string  `xml:"send_type,omitempty" json:"send_type,omitempty"`
	HbType       string  `xml:"hb_type,omitempty" json:"hb_type,omitempty"`
	TotalNum     string  `xml:"total_num,omitempty" json:"total_num,omitempty"`
	TotalAmount  string  `xml:"total_amount,omitempty" json:"total_amount,omitempty"`
	Reason       string  `xml:"reason,omitempty" json:"reason,omitempty"`
	SendTime     string  `xml:"send_time,omitempty" json:"send_time,omitempty"`
	RefundTime   string  `xml:"refund_time,omitempty" json:"refund_time,omitempty"`
	RefundAmount string  `xml:"refund_amount,omitempty" json:"refund_amount,omitempty"`
	Wishing      string  `xml:"wishing,omitempty" json:"wishing,omitempty"`
	Remark       string  `xml:"remark,omitempty" json:"remark,omitempty"`
	ActName      string  `xml:"act_name,omitempty" json:"act_name,omitempty"`
	Hblist       *hbList `xml:"hblist,omitempty" json:"hblist,omitempty"`
}

type hbList struct {
	HbinfoList []*hbinfo `xml:"hbinfo,omitempty" json:"hbinfo,omitempty"`
}

type hbinfo struct {
	Openid  string `xml:"openid,omitempty" json:"openid,omitempty"`
	Amount  string `xml:"amount,omitempty" json:"amount,omitempty"`
	RcvTime string `xml:"rcv_time,omitempty" json:"rcv_time,omitempty"`
}

type CustomsDeclareOrderResponse struct {
	ReturnCode              string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg               string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode              string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode                 string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes              string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	SignType                string `xml:"sign_type,omitempty" json:"sign_type,omitempty"`
	Sign                    string `xml:"sign,omitempty" json:"sign,omitempty"`
	Appid                   string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId                   string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	State                   string `xml:"state,omitempty" json:"state,omitempty"`
	TransactionId           string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo              string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	SubOrderNo              string `xml:"sub_order_no,omitempty" json:"sub_order_no,omitempty"`
	SubOrderId              string `xml:"sub_order_id,omitempty" json:"sub_order_id,omitempty"`
	ModifyTime              string `xml:"modify_time,omitempty" json:"modify_time,omitempty"`
	CertCheckResult         string `xml:"cert_check_result,omitempty" json:"cert_check_result,omitempty"`
	VerifyDepartment        string `xml:"verify_department,omitempty" json:"verify_department,omitempty"`
	VerifyDepartmentTradeId string `xml:"verify_department_trade_id,omitempty" json:"verify_department_trade_id,omitempty"`
}

type CustomsDeclareQueryResponse struct {
	ReturnCode              string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg               string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode              string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode                 string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes              string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Sign                    string `xml:"sign,omitempty" json:"sign,omitempty"`
	Appid                   string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId                   string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	TransactionId           string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	Count                   string `xml:"count,omitempty" json:"count,omitempty"`
	SubOrderNo0             string `xml:"sub_order_no_0,omitempty"`
	SubOrderNo1             string `xml:"sub_order_no_1,omitempty"`
	SubOrderNo2             string `xml:"sub_order_no_2,omitempty"`
	SubOrderId0             string `xml:"sub_order_id_0,omitempty"`
	SubOrderId1             string `xml:"sub_order_id_1,omitempty"`
	SubOrderId2             string `xml:"sub_order_id_2,omitempty"`
	MchCustomsNo0           string `xml:"mch_customs_no_0,omitempty"`
	MchCustomsNo1           string `xml:"mch_customs_no_1,omitempty"`
	MchCustomsNo2           string `xml:"mch_customs_no_2,omitempty"`
	Customs0                string `xml:"customs_0,omitempty"`
	Customs1                string `xml:"customs_1,omitempty"`
	Customs2                string `xml:"customs_2,omitempty"`
	FeeType0                string `xml:"fee_type_0,omitempty"`
	FeeType1                string `xml:"fee_type_1,omitempty"`
	FeeType2                string `xml:"fee_type_2,omitempty"`
	OrderFee0               string `xml:"order_fee_0,omitempty"`
	OrderFee1               string `xml:"order_fee_1,omitempty"`
	OrderFee2               string `xml:"order_fee_2,omitempty"`
	Duty0                   string `xml:"duty_0,omitempty"`
	Duty1                   string `xml:"duty_1,omitempty"`
	Duty2                   string `xml:"duty_2,omitempty"`
	TransportFee0           string `xml:"transport_fee_0,omitempty"`
	TransportFee1           string `xml:"transport_fee_1,omitempty"`
	TransportFee2           string `xml:"transport_fee_2,omitempty"`
	ProductFee0             string `xml:"product_fee_0,omitempty"`
	ProductFee1             string `xml:"product_fee_1,omitempty"`
	ProductFee2             string `xml:"product_fee_2,omitempty"`
	State0                  string `xml:"state_0,omitempty"`
	State1                  string `xml:"state_1,omitempty"`
	State2                  string `xml:"state_2,omitempty"`
	Explanation0            string `xml:"explanation_0,omitempty"`
	Explanation1            string `xml:"explanation_1,omitempty"`
	Explanation2            string `xml:"explanation_2,omitempty"`
	ModifyTime0             string `xml:"modify_time_0,omitempty"`
	ModifyTime1             string `xml:"modify_time_1,omitempty"`
	ModifyTime2             string `xml:"modify_time_2,omitempty"`
	CertCheckResult0        string `xml:"cert_check_result_0,omitempty"`
	CertCheckResult1        string `xml:"cert_check_result_1,omitempty"`
	CertCheckResult2        string `xml:"cert_check_result_2,omitempty"`
	VerifyDepartment        string `xml:"verify_department,omitempty" json:"verify_department,omitempty"`
	VerifyDepartmentTradeId string `xml:"verify_department_trade_id,omitempty" json:"verify_department_trade_id,omitempty"`
}

type CustomsReDeclareOrderResponse struct {
	ReturnCode    string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg     string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode    string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode       string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes    string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	SignType      string `xml:"sign_type,omitempty" json:"sign_type,omitempty"`
	Sign          string `xml:"sign,omitempty" json:"sign,omitempty"`
	Appid         string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId         string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	State         string `xml:"state,omitempty" json:"state,omitempty"`
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	SubOrderNo    string `xml:"sub_order_no,omitempty" json:"sub_order_no,omitempty"`
	SubOrderId    string `xml:"sub_order_id,omitempty" json:"sub_order_id,omitempty"`
	ModifyTime    string `xml:"modify_time,omitempty" json:"modify_time,omitempty"`
	Explanation   string `xml:"explanation,omitempty" json:"explanation,omitempty"`
}
