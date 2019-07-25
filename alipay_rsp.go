//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/16 0:30
//==================================
package gopay

type AliPayNotifyRequest struct {
	NotifyTime        string                  `json:"notify_time"`
	NotifyType        string                  `json:"notify_type"`
	NotifyId          string                  `json:"notify_id"`
	AppId             string                  `json:"app_id"`
	Charset           string                  `json:"charset"`
	Version           string                  `json:"version"`
	SignType          string                  `json:"sign_type"`
	Sign              string                  `json:"sign"`
	AuthAppId         string                  `json:"auth_app_id"`
	TradeNo           string                  `json:"trade_no"`
	OutTradeNo        string                  `json:"out_trade_no"`
	OutBizNo          string                  `json:"out_biz_no"`
	BuyerId           string                  `json:"buyer_id"`
	BuyerLogonId      string                  `json:"buyer_logon_id"`
	SellerId          string                  `json:"seller_id"`
	SellerEmail       string                  `json:"seller_email"`
	TradeStatus       string                  `json:"trade_status"`
	TotalAmount       string                  `json:"total_amount"`
	ReceiptAmount     string                  `json:"receipt_amount"`
	InvoiceAmount     string                  `json:"invoice_amount"`
	BuyerPayAmount    string                  `json:"buyer_pay_amount"`
	PointAmount       string                  `json:"point_amount"`
	RefundFee         string                  `json:"refund_fee"`
	Subject           string                  `json:"subject"`
	Body              string                  `json:"body"`
	GmtCreate         string                  `json:"gmt_create"`
	GmtPayment        string                  `json:"gmt_payment"`
	GmtRefund         string                  `json:"gmt_refund"`
	GmtClose          string                  `json:"gmt_close"`
	FundBillList      []FundBillListInfo      `json:"fund_bill_list"`
	PassbackParams    string                  `json:"passback_params"`
	VoucherDetailList []VoucherDetailListInfo `json:"voucher_detail_list"`
}

type FundBillListInfo struct {
	Amount      string `json:"amount"`
	FundChannel string `json:"fundChannel"`
	BankCode    string `json:"bank_code,omitempty"`
	RealAmount  string `json:"real_amount,omitempty"`
}

type VoucherDetailListInfo struct {
	Id                         string `json:"id"`
	Name                       string `json:"name"`
	Type                       string `json:"type"`
	Amount                     string `json:"amount"`
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
