package alipay

const (
	// URL
	baseUrl            = "https://openapi.alipay.com/gateway.do"
	sandboxBaseUrl     = "https://openapi.alipaydev.com/gateway.do"
	baseUrlUtf8        = "https://openapi.alipay.com/gateway.do?charset=utf-8"
	sandboxBaseUrlUtf8 = "https://openapi.alipaydev.com/gateway.do?charset=utf-8"

	LocationShanghai          = "Asia/Shanghai"
	PKCS1            PKCSType = 1 // 非Java
	PKCS8            PKCSType = 2 // Java
	RSA                       = "RSA"
	RSA2                      = "RSA2"
)

type PKCSType uint8

// Deprecated
type NotifyRequest struct {
	NotifyTime        string              `json:"notify_time,omitempty"`
	NotifyType        string              `json:"notify_type,omitempty"`
	NotifyId          string              `json:"notify_id,omitempty"`
	AppId             string              `json:"app_id,omitempty"`
	Charset           string              `json:"charset,omitempty"`
	Version           string              `json:"version,omitempty"`
	SignType          string              `json:"sign_type,omitempty"`
	Sign              string              `json:"sign,omitempty"`
	AuthAppId         string              `json:"auth_app_id,omitempty"`
	TradeNo           string              `json:"trade_no,omitempty"`
	OutTradeNo        string              `json:"out_trade_no,omitempty"`
	OutBizNo          string              `json:"out_biz_no,omitempty"`
	BuyerId           string              `json:"buyer_id,omitempty"`
	BuyerLogonId      string              `json:"buyer_logon_id,omitempty"`
	SellerId          string              `json:"seller_id,omitempty"`
	SellerEmail       string              `json:"seller_email,omitempty"`
	TradeStatus       string              `json:"trade_status,omitempty"`
	TotalAmount       string              `json:"total_amount,omitempty"`
	ReceiptAmount     string              `json:"receipt_amount,omitempty"`
	InvoiceAmount     string              `json:"invoice_amount,omitempty"`
	BuyerPayAmount    string              `json:"buyer_pay_amount,omitempty"`
	PointAmount       string              `json:"point_amount,omitempty"`
	RefundFee         string              `json:"refund_fee,omitempty"`
	Subject           string              `json:"subject,omitempty"`
	Body              string              `json:"body,omitempty"`
	GmtCreate         string              `json:"gmt_create,omitempty"`
	GmtPayment        string              `json:"gmt_payment,omitempty"`
	GmtRefund         string              `json:"gmt_refund,omitempty"`
	GmtClose          string              `json:"gmt_close,omitempty"`
	FundBillList      []*FundBillListInfo `json:"fund_bill_list,omitempty"`
	PassbackParams    string              `json:"passback_params,omitempty"`
	VoucherDetailList []*VoucherDetail    `json:"voucher_detail_list,omitempty"`
	Method            string              `json:"method,omitempty"`    // 电脑网站支付 支付宝请求 return_url 同步返回参数
	Timestamp         string              `json:"timestamp,omitempty"` // 电脑网站支付 支付宝请求 return_url 同步返回参数
}

// Deprecated
type FundBillListInfo struct {
	Amount      string `json:"amount,omitempty"`
	FundChannel string `json:"fundChannel,omitempty"` // 异步通知里是 fundChannel
}

type UserPhone struct {
	ErrorResponse
	Mobile string `json:"mobile,omitempty"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
}

// ===================================================
type TradePayResponse struct {
	Response     *TradePay `json:"alipay_trade_pay_response"`
	AlipayCertSn string    `json:"alipay_cert_sn,omitempty"`
	SignData     string    `json:"-"`
	Sign         string    `json:"sign"`
}

type TradePay struct {
	ErrorResponse
	TradeNo             string           `json:"trade_no,omitempty"`
	OutTradeNo          string           `json:"out_trade_no,omitempty"`
	BuyerLogonId        string           `json:"buyer_logon_id,omitempty"`
	SettleAmount        string           `json:"settle_amount,omitempty"`
	PayCurrency         string           `json:"pay_currency,omitempty"`
	PayAmount           string           `json:"pay_amount,omitempty"`
	SettleTransRate     string           `json:"settle_trans_rate,omitempty"`
	TransPayRate        string           `json:"trans_pay_rate,omitempty"`
	TotalAmount         string           `json:"total_amount,omitempty"`
	TransCurrency       string           `json:"trans_currency,omitempty"`
	SettleCurrency      string           `json:"settle_currency,omitempty"`
	ReceiptAmount       string           `json:"receipt_amount,omitempty"`
	BuyerPayAmount      string           `json:"buyer_pay_amount,omitempty"`
	PointAmount         string           `json:"point_amount,omitempty"`
	InvoiceAmount       string           `json:"invoice_amount,omitempty"`
	GmtPayment          string           `json:"gmt_payment,omitempty"`
	FundBillList        []*FundBill      `json:"fund_bill_list"`
	CardBalance         string           `json:"card_balance,omitempty"`
	StoreName           string           `json:"store_name,omitempty"`
	BuyerUserId         string           `json:"buyer_user_id,omitempty"`
	DiscountGoodsDetail string           `json:"discount_goods_detail,omitempty"`
	VoucherDetailList   []*VoucherDetail `json:"voucher_detail_list"`
	AdvanceAmount       string           `json:"advance_amount,omitempty"`
	AuthTradePayMode    string           `json:"auth_trade_pay_mode,omitempty"`
	ChargeAmount        string           `json:"charge_amount,omitempty"`
	ChargeFlags         string           `json:"charge_flags,omitempty"`
	SettlementId        string           `json:"settlement_id,omitempty"`
	BusinessParams      string           `json:"business_params,omitempty"`
	BuyerUserType       string           `json:"buyer_user_type,omitempty"`
	MdiscountAmount     string           `json:"mdiscount_amount,omitempty"`
	DiscountAmount      string           `json:"discount_amount,omitempty"`
	BuyerUserName       string           `json:"buyer_user_name,omitempty"`
}

type FundBill struct {
	FundChannel string `json:"fund_channel,omitempty"`
	BankCode    string `json:"bank_code,omitempty"`
	Amount      string `json:"amount,omitempty"`
	RealAmount  string `json:"real_amount,omitempty"`
}

type VoucherDetail struct {
	Id                         string `json:"id,omitempty"`
	Name                       string `json:"name,omitempty"`
	Type                       string `json:"type,omitempty"`
	Amount                     string `json:"amount,omitempty"`
	MerchantContribute         string `json:"merchant_contribute,omitempty"`
	OtherContribute            string `json:"other_contribute,omitempty"`
	Memo                       string `json:"memo,omitempty"`
	TemplateId                 string `json:"template_id,omitempty"`
	PurchaseBuyerContribute    string `json:"purchase_buyer_contribute,omitempty"`
	PurchaseMerchantContribute string `json:"purchase_merchant_contribute,omitempty"`
	PurchaseAntContribute      string `json:"purchase_ant_contribute,omitempty"`
}

// ===================================================
type TradeQueryResponse struct {
	Response     *TradeQuery `json:"alipay_trade_query_response"`
	AlipayCertSn string      `json:"alipay_cert_sn,omitempty"`
	SignData     string      `json:"-"`
	Sign         string      `json:"sign"`
}

type TradeQuery struct {
	ErrorResponse
	TradeNo         string      `json:"trade_no,omitempty"`
	OutTradeNo      string      `json:"out_trade_no,omitempty"`
	BuyerLogonId    string      `json:"buyer_logon_id,omitempty"`
	TradeStatus     string      `json:"trade_status,omitempty"`
	TotalAmount     string      `json:"total_amount,omitempty"`
	TransCurrency   string      `json:"trans_currency,omitempty"`
	SettleCurrency  string      `json:"settle_currency,omitempty"`
	SettleAmount    string      `json:"settle_amount,omitempty"`
	PayCurrency     string      `json:"pay_currency,omitempty"`
	PayAmount       string      `json:"pay_amount,omitempty"`
	SettleTransRate string      `json:"settle_trans_rate,omitempty"`
	TransPayRate    string      `json:"trans_pay_rate,omitempty"`
	BuyerPayAmount  string      `json:"buyer_pay_amount,omitempty"`
	PointAmount     string      `json:"point_amount,omitempty"`
	InvoiceAmount   string      `json:"invoice_amount,omitempty"`
	SendPayDate     string      `json:"send_pay_date,omitempty"`
	ReceiptAmount   string      `json:"receipt_amount,omitempty"`
	StoreId         string      `json:"store_id,omitempty"`
	TerminalId      string      `json:"terminal_id,omitempty"`
	FundBillList    []*FundBill `json:"fund_bill_list"`
	StoreName       string      `json:"store_name,omitempty"`
	BuyerUserId     string      `json:"buyer_user_id,omitempty"`
	ChargeAmount    string      `json:"charge_amount,omitempty"`
	ChargeFlags     string      `json:"charge_flags,omitempty"`
	SettlementId    string      `json:"settlement_id,omitempty"`
	TradeSettleInfo *struct {
		TradeSettleDetailList []*struct {
			OperationType     string `json:"operation_type,omitempty"`
			OperationSerialNo string `json:"operation_serial_no,omitempty"`
			OperationDt       string `json:"operation_dt,omitempty"`
			TransOut          string `json:"trans_out,omitempty"`
			TransIn           string `json:"trans_in,omitempty"`
			Amount            string `json:"amount,omitempty"`
		} `json:"trade_settle_detail_list,omitempty"`
	} `json:"trade_settle_info,omitempty"`
	AuthTradePayMode    string `json:"auth_trade_pay_mode,omitempty"`
	BuyerUserType       string `json:"buyer_user_type,omitempty"`
	MdiscountAmount     string `json:"mdiscount_amount,omitempty"`
	DiscountAmount      string `json:"discount_amount,omitempty"`
	BuyerUserName       string `json:"buyer_user_name,omitempty"`
	Subject             string `json:"subject,omitempty"`
	Body                string `json:"body,omitempty"`
	AlipaySubMerchantId string `json:"alipay_sub_merchant_id,omitempty"`
	ExtInfos            string `json:"ext_infos,omitempty"`
}

// ===================================================
type TradeCreateResponse struct {
	Response     *TradeCreate `json:"alipay_trade_create_response"`
	AlipayCertSn string       `json:"alipay_cert_sn,omitempty"`
	SignData     string       `json:"-"`
	Sign         string       `json:"sign"`
}

type TradeCreate struct {
	ErrorResponse
	TradeNo    string `json:"trade_no,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
}

// ===================================================
type TradeCloseResponse struct {
	Response     *TradeClose `json:"alipay_trade_close_response"`
	AlipayCertSn string      `json:"alipay_cert_sn,omitempty"`
	SignData     string      `json:"-"`
	Sign         string      `json:"sign"`
}

type TradeClose struct {
	ErrorResponse
	TradeNo    string `json:"trade_no,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
}

// ===================================================
type TradeCancelResponse struct {
	Response     *TradeCancel `json:"alipay_trade_cancel_response"`
	AlipayCertSn string       `json:"alipay_cert_sn,omitempty"`
	SignData     string       `json:"-"`
	Sign         string       `json:"sign"`
}

type TradeCancel struct {
	ErrorResponse
	TradeNo            string `json:"trade_no,omitempty"`
	OutTradeNo         string `json:"out_trade_no,omitempty"`
	RetryFlag          string `json:"retry_flag,omitempty"`
	Action             string `json:"action,omitempty"`
	GmtRefundPay       string `json:"gmt_refund_pay,omitempty"`
	RefundSettlementId string `json:"refund_settlement_id,omitempty"`
}

// ===================================================
type SystemOauthTokenResponse struct {
	Response      *OauthTokenInfo `json:"alipay_system_oauth_token_response"`
	ErrorResponse *ErrorResponse  `json:"error_response,omitempty"`
	AlipayCertSn  string          `json:"alipay_cert_sn,omitempty"`
	SignData      string          `json:"-"`
	Sign          string          `json:"sign"`
}

type OauthTokenInfo struct {
	AccessToken  string `json:"access_token,omitempty"`
	AlipayUserId string `json:"alipay_user_id,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	ReExpiresIn  int64  `json:"re_expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	UserId       string `json:"user_id,omitempty"`
}

// ===================================================
type UserInfoShareResponse struct {
	Response     *UserInfoShare `json:"alipay_user_info_share_response"`
	AlipayCertSn string         `json:"alipay_cert_sn,omitempty"`
	SignData     string         `json:"-"`
	Sign         string         `json:"sign"`
}

type UserInfoShare struct {
	ErrorResponse
	UserId             string `json:"user_id,omitempty"`
	Avatar             string `json:"avatar,omitempty"`
	Province           string `json:"province,omitempty"`
	City               string `json:"city,omitempty"`
	NickName           string `json:"nick_name,omitempty"`
	IsStudentCertified string `json:"is_student_certified,omitempty"`
	UserType           string `json:"user_type,omitempty"`
	UserStatus         string `json:"user_status,omitempty"`
	IsCertified        string `json:"is_certified,omitempty"`
	Gender             string `json:"gender,omitempty"`
}

// ===================================================
type TradeRefundResponse struct {
	Response     *TradeRefund `json:"alipay_trade_refund_response"`
	AlipayCertSn string       `json:"alipay_cert_sn,omitempty"`
	SignData     string       `json:"-"`
	Sign         string       `json:"sign"`
}

type TradeRefund struct {
	ErrorResponse
	TradeNo                      string                 `json:"trade_no,omitempty"`
	OutTradeNo                   string                 `json:"out_trade_no,omitempty"`
	BuyerLogonId                 string                 `json:"buyer_logon_id,omitempty"`
	FundChange                   string                 `json:"fund_change,omitempty"`
	RefundFee                    string                 `json:"refund_fee,omitempty"`
	RefundCurrency               string                 `json:"refund_currency,omitempty"`
	GmtRefundPay                 string                 `json:"gmt_refund_pay,omitempty"`
	RefundDetailItemList         []*TradeFundBill       `json:"refund_detail_item_list,omitempty"`
	StoreName                    string                 `json:"store_name,omitempty"`
	BuyerUserId                  string                 `json:"buyer_user_id,omitempty"`
	RefundPresetPaytoolList      []*RefundPresetPaytool `json:"refund_preset_paytool_list,omitempty"`
	RefundSettlementId           string                 `json:"refund_settlement_id,omitempty"`
	PresentRefundBuyerAmount     string                 `json:"present_refund_buyer_amount,omitempty"`
	PresentRefundDiscountAmount  string                 `json:"present_refund_discount_amount,omitempty"`
	PresentRefundMdiscountAmount string                 `json:"present_refund_mdiscount_amount,omitempty"`
}

type TradeFundBill struct {
	FundChannel string `json:"fund_channel,omitempty"` //同步通知里是 fund_channel
	BankCode    string `json:"bank_code,omitempty"`
	Amount      string `json:"amount,omitempty"`
	RealAmount  string `json:"real_amount,omitempty"`
	FundType    string `json:"fund_type,omitempty"`
}

type RefundPresetPaytool struct {
	Amount         []string `json:"amount,omitempty"`
	AssertTypeCode string   `json:"assert_type_code,omitempty"`
}

// ===================================================
type TradeFastpayRefundQueryResponse struct {
	Response     *TradeRefundQuery `json:"alipay_trade_fastpay_refund_query_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type TradeRefundQuery struct {
	ErrorResponse
	TradeNo                      string           `json:"trade_no,omitempty"`
	OutTradeNo                   string           `json:"out_trade_no,omitempty"`
	OutRequestNo                 string           `json:"out_request_no,omitempty"`
	RefundReason                 string           `json:"refund_reason,omitempty"`
	TotalAmount                  string           `json:"total_amount,omitempty"`
	RefundAmount                 string           `json:"refund_amount,omitempty"`
	RefundRoyaltys               []*RefundRoyalty `json:"refund_royaltys,omitempty"`
	GmtRefundPay                 string           `json:"gmt_refund_pay,omitempty"`
	RefundDetailItemList         []*TradeFundBill `json:"refund_detail_item_list,omitempty"`
	SendBackFee                  string           `json:"send_back_fee,omitempty"`
	RefundSettlementId           string           `json:"refund_settlement_id,omitempty"`
	PresentRefundBuyerAmount     string           `json:"present_refund_buyer_amount,omitempty"`
	PresentRefundDiscountAmount  string           `json:"present_refund_discount_amount,omitempty"`
	PresentRefundMdiscountAmount string           `json:"present_refund_mdiscount_amount,omitempty"`
}

type RefundRoyalty struct {
	RefundAmount  string `json:"refund_amount,omitempty"`
	RoyaltyType   string `json:"royalty_type,omitempty"`
	ResultCode    string `json:"result_code,omitempty"`
	TransOut      string `json:"trans_out,omitempty"`
	TransOutEmail string `json:"trans_out_email,omitempty"`
	TransIn       string `json:"trans_in,omitempty"`
	TransInEmail  string `json:"trans_in_email,omitempty"`
}

// ===================================================
type TradeOrderSettleResponse struct {
	Response     *TradeOrderSettle `json:"alipay_trade_order_settle_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type TradeOrderSettle struct {
	ErrorResponse
	TradeNo string `json:"trade_no,omitempty"`
}

// ===================================================
type TradePrecreateResponse struct {
	Response     *TradePrecreate `json:"alipay_trade_precreate_response"`
	NullResponse *ErrorResponse  `json:"null_response,omitempty"`
	AlipayCertSn string          `json:"alipay_cert_sn,omitempty"`
	SignData     string          `json:"-"`
	Sign         string          `json:"sign"`
}

type TradePrecreate struct {
	ErrorResponse
	OutTradeNo string `json:"out_trade_no,omitempty"`
	QrCode     string `json:"qr_code,omitempty"`
}

// ===================================================
type TradePageRefundResponse struct {
	Response     *TradePageRefund `json:"alipay_trade_page_refund_response"`
	AlipayCertSn string           `json:"alipay_cert_sn,omitempty"`
	SignData     string           `json:"-"`
	Sign         string           `json:"sign"`
}

type TradePageRefund struct {
	ErrorResponse
	TradeNo      string `json:"trade_no,omitempty"`
	OutTradeNo   string `json:"out_trade_no,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
	RefundAmount string `json:"refund_amount,omitempty"`
}

// ===================================================
type FundTransToaccountTransferResponse struct {
	Response     *TransToaccountTransfer `json:"alipay_fund_trans_toaccount_transfer_response"`
	AlipayCertSn string                  `json:"alipay_cert_sn,omitempty"`
	SignData     string                  `json:"-"`
	Sign         string                  `json:"sign"`
}

type TransToaccountTransfer struct {
	ErrorResponse
	OutBizNo string `json:"out_biz_no,omitempty"`
	OrderId  string `json:"order_id,omitempty"`
	PayDate  string `json:"pay_date,omitempty"`
}

type FundTransUniTransferResponse struct {
	Response     *TransUniTransfer `json:"alipay_fund_trans_uni_transfer_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type TransUniTransfer struct {
	ErrorResponse
	OutBizNo       string `json:"out_biz_no,omitempty"`
	OrderId        string `json:"order_id,omitempty"`
	PayFundOrderId string `json:"pay_fund_order_id,omitempty"`
	Status         string `json:"status,omitempty"`
	TransDate      string `json:"trans_date,omitempty"`
}

// ===================================================
type FundTransCommonQueryResponse struct {
	Response     *FundTransCommonQuery `json:"alipay_fund_trans_common_query_response"`
	AlipayCertSn string                `json:"alipay_cert_sn,omitempty"`
	SignData     string                `json:"-"`
	Sign         string                `json:"sign"`
}

type FundTransCommonQuery struct {
	ErrorResponse
	OrderId        string `json:"order_id,omitempty"`
	PayFundOrderId string `json:"pay_fund_order_id,omitempty"`
	OutBizNo       string `json:"out_biz_no,omitempty"`
	TransAmount    string `json:"trans_amount,omitempty"`
	Status         string `json:"status,omitempty"`
	PayDate        string `json:"pay_date,omitempty"`
	ArrivalTimeEnd string `json:"arrival_time_end,omitempty"`
	OrderFee       string `json:"order_fee,omitempty"`
	ErrorCode      string `json:"error_code,omitempty"`
	FailReason     string `json:"fail_reason,omitempty"`
}

// ===================================================

type FundTransOrderQueryResponse struct {
	Response     *FundTransOrderQuery `json:"alipay_fund_trans_order_query_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type FundTransOrderQuery struct {
	ErrorResponse
	OrderId        string `json:"order_id,omitempty"`
	Status         string `json:"status,omitempty"`
	PayDate        string `json:"pay_date,omitempty"`
	ArrivalTimeEnd string `json:"arrival_time_end,omitempty"`
	OrderFee       string `json:"order_fee,omitempty"`
	FailReason     string `json:"fail_reason,omitempty"`
	OutBizNo       string `json:"out_biz_no,omitempty"`
	ErrorCode      string `json:"error_code,omitempty"`
}

// ===================================================
type FundTransRefundResponse struct {
	Response     *FundTransRefund `json:"alipay_fund_trans_refund_response"`
	AlipayCertSn string           `json:"alipay_cert_sn,omitempty"`
	SignData     string           `json:"-"`
	Sign         string           `json:"sign"`
}

type FundTransRefund struct {
	ErrorResponse
	RefundOrderId string `json:"refund_order_id"`
	OrderId       string `json:"order_id"`
	OutRequestNo  string `json:"out_request_no"`
	Status        string `json:"status"`
	RefundAmount  string `json:"refund_amount"`
	RefundDate    string `json:"refund_date"`
}

// ===================================================
type FundAuthOrderFreezeResponse struct {
	Response     *FundAuthOrderFreeze `json:"alipay_fund_auth_order_freeze_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type FundAuthOrderFreeze struct {
	ErrorResponse
	AuthNo       string `json:"auth_no,omitempty"`
	OutOrderNo   string `json:"out_order_no,omitempty"`
	OperationId  string `json:"operation_id,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
	Amount       string `json:"amount,omitempty"`
	Status       string `json:"status,omitempty"`
	PayerUserId  string `json:"payer_user_id,omitempty"`
	PayerLogonId string `json:"payer_logon_id,omitempty"`
	GmtTrans     string `json:"gmt_trans,omitempty"`
}

// ===================================================
type FundAuthOrderVoucherCreateResponse struct {
	Response     *FundAuthOrderVoucherCreate `json:"alipay_fund_auth_order_voucher_create_response"`
	AlipayCertSn string                      `json:"alipay_cert_sn,omitempty"`
	SignData     string                      `json:"-"`
	Sign         string                      `json:"sign"`
}

type FundAuthOrderVoucherCreate struct {
	ErrorResponse
	OutOrderNo   string `json:"out_order_no,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
	CodeType     string `json:"code_type,omitempty"`
	CodeValue    string `json:"code_value,omitempty"`
	CodeUrl      string `json:"code_url,omitempty"`
}

// ===================================================
type FundAuthOrderAppFreezeResponse struct {
	Response     *FundAuthOrderAppFreeze `json:"alipay_fund_auth_order_app_freeze_response"`
	AlipayCertSn string                  `json:"alipay_cert_sn,omitempty"`
	SignData     string                  `json:"-"`
	Sign         string                  `json:"sign"`
}

type FundAuthOrderAppFreeze struct {
	ErrorResponse
	AuthNo        string `json:"auth_no,omitempty"`
	OutOrderNo    string `json:"out_order_no,omitempty"`
	OperationId   string `json:"operation_id,omitempty"`
	OutRequestNo  string `json:"out_request_no,omitempty"`
	Amount        string `json:"amount,omitempty"`
	Status        string `json:"status,omitempty"`
	PayerUserId   string `json:"payer_user_id,omitempty"`
	GmtTrans      string `json:"gmt_trans,omitempty"`
	PreAuthType   string `json:"pre_auth_type,omitempty"`
	CreditAmount  string `json:"credit_amount,omitempty"`
	FundAmount    string `json:"fund_amount,omitempty"`
	TransCurrency string `json:"trans_currency,omitempty"`
}

// ===================================================
type FundAuthOrderUnfreezeResponse struct {
	Response     *FundAuthOrderUnfreeze `json:"alipay_fund_auth_order_unfreeze_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type FundAuthOrderUnfreeze struct {
	ErrorResponse
	AuthNo       string `json:"auth_no,omitempty"`
	OutOrderNo   string `json:"out_order_no,omitempty"`
	OperationId  string `json:"operation_id,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
	Amount       string `json:"amount,omitempty"`
	Status       string `json:"status,omitempty"`
	GmtTrans     string `json:"gmt_trans,omitempty"`
	CreditAmount string `json:"credit_amount,omitempty"`
	FundAmount   string `json:"fund_amount,omitempty"`
}

// ===================================================
type FundAuthOperationDetailQueryResponse struct {
	Response     *FundAuthOperationDetailQuery `json:"alipay_fund_auth_operation_detail_query_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type FundAuthOperationDetailQuery struct {
	ErrorResponse
	AuthNo                  string `json:"auth_no,omitempty"`
	OutOrderNo              string `json:"out_order_no,omitempty"`
	TotalFreezeAmount       string `json:"total_freeze_amount,omitempty"`
	RestAmount              string `json:"rest_amount,omitempty"`
	TotalPayAmount          string `json:"total_pay_amount,omitempty"`
	OrderTitle              string `json:"order_title,omitempty"`
	PayerLogonId            string `json:"payer_logon_id,omitempty"`
	PayerUserId             string `json:"payer_user_id,omitempty"`
	ExtraParam              string `json:"extra_param,omitempty"`
	OperationId             string `json:"operation_id,omitempty"`
	OutRequestNo            string `json:"out_request_no,omitempty"`
	Amount                  string `json:"amount,omitempty"`
	OperationType           string `json:"operation_type,omitempty"`
	Status                  string `json:"status,omitempty"`
	Remark                  string `json:"remark,omitempty"`
	GmtCreate               string `json:"gmt_create,omitempty"`
	GmtTrans                string `json:"gmt_trans,omitempty"`
	PreAuthType             string `json:"pre_auth_type,omitempty"`
	TransCurrency           string `json:"trans_currency,omitempty"`
	TotalFreezeCreditAmount string `json:"total_freeze_credit_amount,omitempty"`
	TotalFreezeFundAmount   string `json:"total_freeze_fund_amount,omitempty"`
	TotalPayCreditAmount    string `json:"total_pay_credit_amount,omitempty"`
	TotalPayFundAmount      string `json:"total_pay_fund_amount,omitempty"`
	RestCreditAmount        string `json:"rest_credit_amount,omitempty"`
	RestFundAmount          string `json:"rest_fund_amount,omitempty"`
	CreditAmount            string `json:"credit_amount,omitempty"`
	FundAmount              string `json:"fund_amount,omitempty"`
}

// ===================================================
type FundAuthOperationCancelResponse struct {
	Response     *FundAuthOperationCancel `json:"alipay_fund_auth_operation_cancel_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type FundAuthOperationCancel struct {
	ErrorResponse
	AuthNo       string `json:"auth_no,omitempty"`
	OutOrderNo   string `json:"out_order_no,omitempty"`
	OperationId  string `json:"operation_id,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
	Action       string `json:"action,omitempty"`
}

// ===================================================
type FundBatchCreateResponse struct {
	Response     *FundBatchCreate `json:"alipay_fund_batch_create_response"`
	AlipayCertSn string           `json:"alipay_cert_sn,omitempty"`
	SignData     string           `json:"-"`
	Sign         string           `json:"sign"`
}

type FundBatchCreate struct {
	ErrorResponse
	OutBatchNo   string `json:"out_batch_no,omitempty"`
	BatchTransId string `json:"batch_trans_id,omitempty"`
	Status       string `json:"status,omitempty"`
}

// ===================================================
type FundBatchCloseResponse struct {
	Response     *FundBatchClose `json:"alipay_fund_batch_close_response"`
	AlipayCertSn string          `json:"alipay_cert_sn,omitempty"`
	SignData     string          `json:"-"`
	Sign         string          `json:"sign"`
}

type FundBatchClose struct {
	ErrorResponse
	BatchTransId string `json:"batch_trans_id,omitempty"`
	Status       string `json:"status,omitempty"`
}

// ===================================================
type FundBatchDetailQueryResponse struct {
	Response     *FundBatchDetailQuery `json:"alipay_fund_batch_detail_query_response"`
	AlipayCertSn string                `json:"alipay_cert_sn,omitempty"`
	SignData     string                `json:"-"`
	Sign         string                `json:"sign"`
}

type FundBatchDetailQuery struct {
	ErrorResponse
	BatchTransId    string `json:"batch_trans_id,omitempty"`
	BatchNo         string `json:"batch_no,omitempty"`
	BizCode         string `json:"biz_code,omitempty"`
	BizScene        string `json:"biz_scene,omitempty"`
	BatchStatus     string `json:"batch_status,omitempty"`
	ApprovalStatus  string `json:"approval_status,omitempty"`
	ErrorCode       string `json:"error_code,omitempty"`
	FailReason      string `json:"fail_reason,omitempty"`
	SignPrincipal   string `json:"sign_principal,omitempty"`
	PaymentAmount   string `json:"payment_amount,omitempty"`
	PaymentCurrency string `json:"payment_currency,omitempty"`
	PageSize        string `json:"page_size,omitempty"`
	PageNum         string `json:"page_num,omitempty"`
	ProductCode     string `json:"product_code,omitempty"`
	TotalPageCount  string `json:"total_page_count,omitempty"`
	OutBatchNo      string `json:"out_batch_no,omitempty"`
	GmtFinish       string `json:"gmt_finish,omitempty"`
	TotalAmount     string `json:"total_amount,omitempty"`
	GmtPayFinish    string `json:"gmt_pay_finish,omitempty"`
	PayerId         string `json:"payer_id,omitempty"`
	SuccessAmount   string `json:"success_amount,omitempty"`
	FailAmount      string `json:"fail_amount,omitempty"`
	FailCount       string `json:"fail_count,omitempty"`
	SuccessCount    string `json:"success_count,omitempty"`
	TotalItemCount  string `json:"total_item_count,omitempty"`
	AccDetailList   []*struct {
		DetailNo           string `json:"detail_no,omitempty"`
		PaymentAmount      string `json:"payment_amount,omitempty"`
		PaymentCurrency    string `json:"payment_currency,omitempty"`
		TransAmount        string `json:"trans_amount,omitempty"`
		TransCurrency      string `json:"trans_currency,omitempty"`
		SettlementAmount   string `json:"settlement_amount,omitempty"`
		SettlementCurrency string `json:"settlement_currency,omitempty"`
		PayeeInfo          *struct {
			PayeeAccount string `json:"payee_account,omitempty"`
			PayeeType    string `json:"payee_type,omitempty"`
			PayeeName    string `json:"payee_name,omitempty"`
		} `json:"payee_info,omitempty"`
		CertInfo *struct {
			CertNo   string `json:"cert_no,omitempty"`
			CertType string `json:"cert_type,omitempty"`
		} `json:"cert_info,omitempty"`
		Remark       string `json:"remark,omitempty"`
		Status       string `json:"status,omitempty"`
		ExchangeRate *struct {
			Rate             string `json:"rate,omitempty"`
			BaseCurrency     string `json:"base_currency,omitempty"`
			ExchangeCurrency string `json:"exchange_currency,omitempty"`
		} `json:"exchange_rate,omitempty"`
		NeedRetry     string `json:"need_retry,omitempty"`
		AlipayOrderNo string `json:"alipay_order_no,omitempty"`
		OutBizNo      string `json:"out_biz_no,omitempty"`
		DetailId      string `json:"detail_id,omitempty"`
		ErrorCode     string `json:"error_code,omitempty"`
		ErrorMsg      string `json:"error_msg,omitempty"`
		GmtCreate     string `json:"gmt_create,omitempty"`
		GmtFinish     string `json:"gmt_finish,omitempty"`
		SubStatus     string `json:"sub_status,omitempty"`
	} `json:"acc_detail_list,omitempty"`
}

// ===================================================
type FundTransAppPayResponse struct {
	Response     *FundTransAppPay `json:"alipay_fund_trans_app_pay_response"`
	AlipayCertSn string           `json:"alipay_cert_sn,omitempty"`
	SignData     string           `json:"-"`
	Sign         string           `json:"sign"`
}

type FundTransAppPay struct {
	ErrorResponse
	OutBizNo string `json:"out_biz_no,omitempty"`
	OrderId  string `json:"order_id,omitempty"`
	Status   string `json:"status,omitempty"`
}

// ===================================================
type FundTransPayeeBindQueryRsp struct {
	Response     *FundTransPayeeBindQuery `json:"alipay_fund_trans_payee_bind_query_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type FundTransPayeeBindQuery struct {
	ErrorResponse
	Bind string `json:"bind"` // 是否绑定收款账号。true：已绑定；false：未绑定
}

// ===================================================
type FundAccountQueryResponse struct {
	Response     *FundAccountQuery `json:"alipay_fund_account_query_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type FundAccountQuery struct {
	ErrorResponse
	AvailableAmount string       `json:"available_amount,omitempty"`
	ExtCardInfo     *ExtCardInfo `json:"ext_card_info,omitempty"`
}

type ExtCardInfo struct {
	CardNo       string `json:"card_no,omitempty"`
	BankAccName  string `json:"bank_acc_name,omitempty"`
	CardBranch   string `json:"card_branch,omitempty"`
	CardBank     string `json:"card_bank,omitempty"`
	CardLocation string `json:"card_location,omitempty"`
	CardDeposit  string `json:"card_deposit,omitempty"`
	Status       string `json:"status,omitempty"`
}

// ===================================================
type ZhimaCreditScoreGetResponse struct {
	Response     *ScoreGet `json:"zhima_credit_score_get_response"`
	AlipayCertSn string    `json:"alipay_cert_sn,omitempty"`
	SignData     string    `json:"-"`
	Sign         string    `json:"sign"`
}

type ScoreGet struct {
	ErrorResponse
	BizNo   string `json:"biz_no,omitempty"`
	ZmScore string `json:"zm_score,omitempty"`
}

// ===================================================
type OpenAuthTokenAppResponse struct {
	Response     *AuthTokenApp `json:"alipay_open_auth_token_app_response"`
	AlipayCertSn string        `json:"alipay_cert_sn,omitempty"`
	SignData     string        `json:"-"`
	Sign         string        `json:"sign"`
}

type AuthTokenApp struct {
	ErrorResponse
	UserId          string   `json:"user_id,omitempty"`
	AuthAppId       string   `json:"auth_app_id,omitempty"`
	AppAuthToken    string   `json:"app_auth_token,omitempty"`
	AppRefreshToken string   `json:"app_refresh_token,omitempty"`
	ExpiresIn       int64    `json:"expires_in,omitempty"`
	ReExpiresIn     int64    `json:"re_expires_in,omitempty"`
	Tokens          []*Token `json:"tokens,omitempty"`
}

type Token struct {
	AppAuthToken    string `json:"app_auth_token,omitempty"`
	AppRefreshToken string `json:"app_refresh_token,omitempty"`
	AuthAppId       string `json:"auth_app_id,omitempty"`
	ExpiresIn       int64  `json:"expires_in,omitempty"`
	ReExpiresIn     int64  `json:"re_expires_in,omitempty"`
	UserId          string `json:"user_id,omitempty"`
}

// ===================================================
type UserCertifyOpenInitResponse struct {
	Response     *UserCertifyOpenInit `json:"alipay_user_certify_open_initialize_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type UserCertifyOpenInit struct {
	ErrorResponse
	CertifyId string `json:"certify_id,omitempty"`
}

// ===================================================
type UserCertifyOpenQueryResponse struct {
	Response     *UserCertifyOpenQuery `json:"alipay_user_certify_open_query_response"`
	AlipayCertSn string                `json:"alipay_cert_sn,omitempty"`
	SignData     string                `json:"-"`
	Sign         string                `json:"sign"`
}

type UserCertifyOpenQuery struct {
	ErrorResponse
	Passed       string `json:"passed,omitempty"`
	IdentityInfo string `json:"identity_info,omitempty"`
	MaterialInfo string `json:"material_info,omitempty"`
}

// ===================================================
type UserInfoAuthResponse struct {
	Response     *ErrorResponse `json:"alipay_user_info_auth_response"`
	AlipayCertSn string         `json:"alipay_cert_sn,omitempty"`
	SignData     string         `json:"-"`
	Sign         string         `json:"sign"`
}

// ===================================================
type MonitorHeartbeatSynResponse struct {
	Response     *MonitorHeartbeatSynRes `json:"monitor_heartbeat_syn_response"`
	AlipayCertSn string                  `json:"alipay_cert_sn,omitempty"`
	SignData     string                  `json:"-"`
	Sign         string                  `json:"sign"`
}

type MonitorHeartbeatSynRes struct {
	ErrorResponse
	Pid string `json:"pid,omitempty"`
}

// ===================================================
type DataBillBalanceQueryResponse struct {
	Response     *DataBillBalanceQuery `json:"alipay_data_bill_balance_query_response"`
	AlipayCertSn string                `json:"alipay_cert_sn,omitempty"`
	SignData     string                `json:"-"`
	Sign         string                `json:"sign"`
}

type DataBillBalanceQuery struct {
	ErrorResponse
	TotalAmount     string `json:"total_amount,omitempty"`
	AvailableAmount string `json:"available_amount,omitempty"`
	FreezeAmount    string `json:"freeze_amount,omitempty"`
}

// ===================================================
type DataBillDownloadUrlQueryResponse struct {
	Response     *DataBillDownloadUrlQuery `json:"alipay_data_dataservice_bill_downloadurl_query_response"`
	AlipayCertSn string                    `json:"alipay_cert_sn,omitempty"`
	SignData     string                    `json:"-"`
	Sign         string                    `json:"sign"`
}

type DataBillDownloadUrlQuery struct {
	ErrorResponse
	BillDownloadUrl string `json:"bill_download_url,omitempty"`
}

// ===================================================
type PublicCertDownloadRsp struct {
	Response *PublicCertDownload `json:"alipay_open_app_alipaycert_download_response"`
}

type PublicCertDownload struct {
	ErrorResponse
	AlipayCertContent string `json:"alipay_cert_content"`
}

// ===================================================
type TradeOrderInfoSyncRsp struct {
	Response     *TradeOrderInfoSync `json:"alipay_trade_orderinfo_sync_response"`
	AlipayCertSn string              `json:"alipay_cert_sn,omitempty"`
	SignData     string              `json:"-"`
	Sign         string              `json:"sign"`
}

type TradeOrderInfoSync struct {
	ErrorResponse
	TradeNo     string `json:"trade_no"`
	OutTradeNo  string `json:"out_trade_no"`
	BuyerUserId string `json:"buyer_user_id"`
}

// ===================================================
type TradeAdvanceConsultRsp struct {
	Response     *TradeAdvanceConsult `json:"alipay_trade_advance_consult_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type TradeAdvanceConsult struct {
	ErrorResponse
	ReferResult             bool                      `json:"refer_result"`
	WaitRepaymentOrderInfos []*WaitRepaymentOrderInfo `json:"wait_repayment_order_infos,omitempty"`
	WaitRepaymentAmount     string                    `json:"wait_repayment_amount,omitempty"`
	WaitRepaymentOrderCount string                    `json:"wait_repayment_order_count,omitempty"`
	RiskLevel               string                    `json:"risk_level,omitempty"`
	ResultMessage           string                    `json:"result_message"`
	ResultCode              string                    `json:"result_code"`
}

type WaitRepaymentOrderInfo struct {
	AdvanceOrderId      string `json:"advance_order_id"`
	AlipayUserId        string `json:"alipay_user_id"`
	OrigBizOrderId      string `json:"orig_biz_order_id"`
	BizProduct          string `json:"biz_product"`
	WaitRepaymentAmount string `json:"wait_repayment_amount"`
}

// ===================================================
type UserAgreementPageSignRsp struct {
	Response     *UserAgreementPageSign `json:"alipay_user_agreement_page_sign_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type UserAgreementPageSign struct {
	ErrorResponse
	ExternalAgreementNo string `json:"external_agreement_no,omitempty"`
	PersonalProductCode string `json:"personal_product_code"`
	ValidTime           string `json:"valid_time"`
	SignScene           string `json:"sign_scene"`
	AgreementNo         string `json:"agreement_no"`
	ZmOpenId            string `json:"zm_open_id,omitempty"`
	InvalidTime         string `json:"invalid_time"`
	SignTime            string `json:"sign_time"`
	AlipayUserId        string `json:"alipay_user_id"`
	Status              string `json:"status"`
	ForexEligible       string `json:"forex_eligible,omitempty"`
	ExternalLogonId     string `json:"external_logon_id,omitempty"`
	AlipayLogonId       string `json:"alipay_logon_id"`
	CreditAuthMode      string `json:"credit_auth_mode,omitempty"`
}

// ===================================================
type UserAgreementPageUnSignRsp struct {
	Response     *ErrorResponse `json:"alipay_user_agreement_unsign_response"`
	AlipayCertSn string         `json:"alipay_cert_sn,omitempty"`
	SignData     string         `json:"-"`
	Sign         string         `json:"sign"`
}

// ===================================================
type UserAgreementQueryRsp struct {
	Response     *UserAgreementQuery `json:"alipay_user_agreement_query_response"`
	AlipayCertSn string              `json:"alipay_cert_sn,omitempty"`
	SignData     string              `json:"-"`
	Sign         string              `json:"sign"`
}

type UserAgreementQuery struct {
	ErrorResponse
	ValidTime           string `json:"valid_time"`
	AlipayLogonId       string `json:"alipay_logon_id"`
	InvalidTime         string `json:"invalid_time"`
	PricipalType        string `json:"pricipal_type"`
	DeviceId            string `json:"device_id,omitempty"`
	PrincipalId         string `json:"principal_id"`
	SignScene           string `json:"sign_scene"`
	AgreementNo         string `json:"agreement_no"`
	ThirdPartyType      string `json:"third_party_type"`
	Status              string `json:"status"`
	SignTime            string `json:"sign_time"`
	PersonalProductCode string `json:"personal_product_code"`
	ExternalAgreementNo string `json:"external_agreement_no,omitempty"`
	ZmOpenId            string `json:"zm_open_id,omitempty"`
	ExternalLogonId     string `json:"external_logon_id,omitempty"`
	CreditAuthMode      string `json:"credit_auth_mode,omitempty"`
}

// ===================================================
type OpenAppQrcodeCreateRsp struct {
	Response     *OpenAppQrcodeCreate `json:"alipay_open_app_qrcode_create_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type OpenAppQrcodeCreate struct {
	ErrorResponse
	QrCodeUrl string `json:"qr_code_url"`
}

// ===================================================
type MerchantItemFileUploadRsp struct {
	Response     *MerchantItemFileUpload `json:"alipay_merchant_item_file_upload_response"`
	AlipayCertSn string                  `json:"alipay_cert_sn,omitempty"`
	SignData     string                  `json:"-"`
	Sign         string                  `json:"sign"`
}

type MerchantItemFileUpload struct {
	ErrorResponse
	MaterialId  string `json:"material_id"`  // 文件在商品中心的素材标识（素材ID长期有效）
	MaterialKey string `json:"material_key"` // 文件在商品中心的素材标示，创建/更新商品时使用
}

// ===============================================================
type TradeCustomsDeclareRsp struct {
	Response     *TradeCustomsDeclare `json:"alipay_trade_customs_declare_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type TradeCustomsDeclare struct {
	ErrorResponse
	TradeNo          string `json:"trade_no,omitempty"`
	AlipayDeclareNo  string `json:"alipay_declare_no"`
	PayCode          string `json:"pay_code,omitempty"`
	PayTransactionId string `json:"pay_transaction_id,omitempty"`
	TotalAmount      string `json:"total_amount,omitempty"`
	Currency         string `json:"currency,omitempty"`
	VerDept          string `json:"ver_dept,omitempty"`
	IdentityCheck    string `json:"identity_check,omitempty"`

}

// ===================================================
type TradeOrderAggregateConsultRsp struct {
	Response     *TradeOrderAggregateConsult `json:"koubei_trade_order_aggregate_consult_response"`
	AlipayCertSn string                      `json:"alipay_cert_sn,omitempty"`
	SignData     string                      `json:"-"`
	Sign         string                      `json:"sign"`
}

type TradeOrderAggregateConsult struct {
	ErrorResponse
	OutOrderNo             string                `json:"out_order_no,omitempty"`
	OrderNo                string                `json:"order_no,omitempty"`
	TradeNo                string                `json:"trade_no,omitempty"`
	BuyerId                string                `json:"buyer_id,omitempty"`
	BuyerIdType            string                `json:"buyer_id_type,omitempty"`
	TotalAmount            string                `json:"total_amount,omitempty"`
	ReceiptAmount          string                `json:"receipt_amount,omitempty"`
	BuyerPayAmount         string                `json:"buyer_pay_amount,omitempty"`
	MerchantDiscountAmount string                `json:"merchant_discount_amount,omitempty"`
	PlatformDiscountAmount string                `json:"platform_discount_amount,omitempty"`
	DiscountDetailList     []*DiscountDetailInfo `json:"discount_detail_list,omitempty"`
	OrderStatus            string                `json:"order_status,omitempty"`
	PayChannel             string                `json:"pay_channel,omitempty"`
	CreateTime             string                `json:"create_time"`
	GmtPaymentTime         string                `json:"gmt_payment_time,omitempty"`
}

type DiscountDetailInfo struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Type   string `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`

  
