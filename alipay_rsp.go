//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/16 0:30
//==================================
package gopay

type AliPayNotifyRequest struct {
	NotifyTime        string                  `json:"notify_time,omitempty"`
	NotifyType        string                  `json:"notify_type,omitempty"`
	NotifyId          string                  `json:"notify_id,omitempty"`
	AppId             string                  `json:"app_id,omitempty"`
	Charset           string                  `json:"charset,omitempty"`
	Version           string                  `json:"version,omitempty"`
	SignType          string                  `json:"sign_type,omitempty"`
	Sign              string                  `json:"sign,omitempty"`
	AuthAppId         string                  `json:"auth_app_id,omitempty"`
	TradeNo           string                  `json:"trade_no,omitempty"`
	OutTradeNo        string                  `json:"out_trade_no,omitempty"`
	OutBizNo          string                  `json:"out_biz_no,omitempty"`
	BuyerId           string                  `json:"buyer_id,omitempty"`
	BuyerLogonId      string                  `json:"buyer_logon_id,omitempty"`
	SellerId          string                  `json:"seller_id,omitempty"`
	SellerEmail       string                  `json:"seller_email,omitempty"`
	TradeStatus       string                  `json:"trade_status,omitempty"`
	TotalAmount       string                  `json:"total_amount,omitempty"`
	ReceiptAmount     string                  `json:"receipt_amount,omitempty"`
	InvoiceAmount     string                  `json:"invoice_amount,omitempty"`
	BuyerPayAmount    string                  `json:"buyer_pay_amount,omitempty"`
	PointAmount       string                  `json:"point_amount,omitempty"`
	RefundFee         string                  `json:"refund_fee,omitempty"`
	Subject           string                  `json:"subject,omitempty"`
	Body              string                  `json:"body,omitempty"`
	GmtCreate         string                  `json:"gmt_create,omitempty"`
	GmtPayment        string                  `json:"gmt_payment,omitempty"`
	GmtRefund         string                  `json:"gmt_refund,omitempty"`
	GmtClose          string                  `json:"gmt_close,omitempty"`
	FundBillList      []FundBillListInfo      `json:"fund_bill_list,omitempty"`
	PassbackParams    string                  `json:"passback_params,omitempty"`
	VoucherDetailList []VoucherDetailListInfo `json:"voucher_detail_list,omitempty"`
}

type FundBillListInfo struct {
	Amount      string `json:"amount,omitempty"`
	FundChannel string `json:"fundChannel,omitempty"`
	BankCode    string `json:"bank_code,omitempty"`
	RealAmount  string `json:"real_amount,omitempty"`
}

type VoucherDetailListInfo struct {
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

type AliPayTradePayResponse struct {
	AlipayTradePayResponse AlipayTradeResponseInfo `json:"alipay_trade_pay_response"`
	Sign                   string                  `json:"sign"`
}

type AliPayTradeQueryResponse struct {
	AlipayTradeQueryResponse AlipayTradeResponseInfo `json:"alipay_trade_query_response"`
	Sign                     string                  `json:"sign"`
}

type AliPayTradeCreateResponse struct {
	AliPayTradeCreateResponse AlipayTradeResponseInfo `json:"alipay_trade_create_response"`
	Sign                      string                  `json:"sign"`
}

type AliPayTradeCloseResponse struct {
	AlipayTradeCloseResponse AlipayTradeResponseInfo `json:"alipay_trade_close_response"`
	Sign                     string                  `json:"sign"`
}

type AliPayTradeCancelResponse struct {
	AliPayTradeCancelResponse AlipayTradeResponseInfo `json:"alipay_trade_cancel_response"`
	Sign                      string                  `json:"sign"`
}

type AlipayTradeResponseInfo struct {
	Code                string                  `json:"code,omitempty"`
	Msg                 string                  `json:"msg,omitempty"`
	SubCode             string                  `json:"sub_code,omitempty"`
	SubMsg              string                  `json:"sub_msg,omitempty"`
	TradeNo             string                  `json:"trade_no,omitempty"`
	OutTradeNo          string                  `json:"out_trade_no,omitempty"`
	BuyerLogonId        string                  `json:"buyer_logon_id,omitempty"`
	TradeStatus         string                  `json:"trade_status,omitempty"`
	SettleAmount        string                  `json:"settle_amount,omitempty"`
	PayCurrency         string                  `json:"pay_currency,omitempty"`
	PayAmount           string                  `json:"pay_amount,omitempty"`
	SettleTransRate     string                  `json:"settle_trans_rate,omitempty"`
	TransPayRate        string                  `json:"trans_pay_rate,omitempty"`
	TotalAmount         string                  `json:"total_amount,omitempty"`
	TransCurrency       string                  `json:"trans_currency,omitempty"`
	SettleCurrency      string                  `json:"settle_currency,omitempty"`
	ReceiptAmount       string                  `json:"receipt_amount,omitempty"`
	BuyerPayAmount      string                  `json:"buyer_pay_amount,omitempty"`
	PointAmount         string                  `json:"point_amount,omitempty"`
	InvoiceAmount       string                  `json:"invoice_amount,omitempty"`
	SendPayDate         string                  `json:"send_pay_date,omitempty"`
	GmtPayment          string                  `json:"gmt_payment,omitempty"`
	FundBillList        []FundBillListInfo      `json:"fund_bill_list,omitempty"`
	CardBalance         string                  `json:"card_balance,omitempty"`
	TerminalId          string                  `json:"terminal_id,omitempty"`
	StoreId             string                  `json:"store_id,omitempty"`
	StoreName           string                  `json:"store_name,omitempty"`
	BuyerUserId         string                  `json:"buyer_user_id,omitempty"`
	DiscountGoodsDetail string                  `json:"discount_goods_detail,omitempty"`
	VoucherDetailList   []VoucherDetailListInfo `json:"voucher_detail_list,omitempty"`
	AdvanceAmount       string                  `json:"advance_amount,omitempty"`
	AuthTradePayMode    string                  `json:"auth_trade_pay_mode,omitempty"`
	ChargeAmount        string                  `json:"charge_amount,omitempty"`
	ChargeFlags         string                  `json:"charge_flags,omitempty"`
	SettlementId        string                  `json:"settlement_id,omitempty"`
	BusinessParams      string                  `json:"business_params,omitempty"`
	BuyerUserType       string                  `json:"buyer_user_type,omitempty"`
	MdiscountAmount     string                  `json:"mdiscount_amount,omitempty"`
	DiscountAmount      string                  `json:"discount_amount,omitempty"`
	BuyerUserName       string                  `json:"buyer_user_name,omitempty"`
	Subject             string                  `json:"subject,omitempty"`
	Body                string                  `json:"body,omitempty"`
	AlipaySubMerchantId string                  `json:"alipay_sub_merchant_id,omitempty"`
	ExtInfos            string                  `json:"ext_infos,omitempty"`
	RetryFlag           string                  `json:"retry_flag,omitempty"`
	Action              string                  `json:"action,omitempty"`
	GmtRefundPay        string                  `json:"gmt_refund_pay,omitempty"`
	RefundSettlementId  string                  `json:"refund_settlement_id,omitempty"`
}

type AlipaySystemOauthTokenResponse struct {
	AlipaySystemOauthTokenResponse OauthTokenInfo      `json:"alipay_system_oauth_token_response,omitempty"`
	ErrorResponse                  AlipayErrorResponse `json:"error_response,omitempty"`
	Sign                           string              `json:"sign"`
}

type AlipayErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Msg     string `json:"msg,omitempty"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
}

type OauthTokenInfo struct {
	//Code         string `json:"code,omitempty"`
	//Msg          string `json:"msg,omitempty"`
	//SubCode      string `json:"sub_code,omitempty"`
	//SubMsg       string `json:"sub_msg,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	AlipayUserId string `json:"alipay_user_id,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
	ReExpiresIn  int    `json:"re_expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	UserId       string `json:"user_id,omitempty"`
}

type PhoneNumberResponse struct {
	Code    string `json:"code,omitempty"`
	Msg     string `json:"msg,omitempty"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
	Mobile  string `json:"mobile,omitempty"`
}
