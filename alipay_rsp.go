package gopay

type AliPayNotifyRequest struct {
	NotifyTime        string                   `json:"notify_time,omitempty"`
	NotifyType        string                   `json:"notify_type,omitempty"`
	NotifyId          string                   `json:"notify_id,omitempty"`
	AppId             string                   `json:"app_id,omitempty"`
	Charset           string                   `json:"charset,omitempty"`
	Version           string                   `json:"version,omitempty"`
	SignType          string                   `json:"sign_type,omitempty"`
	Sign              string                   `json:"sign,omitempty"`
	AuthAppId         string                   `json:"auth_app_id,omitempty"`
	TradeNo           string                   `json:"trade_no,omitempty"`
	OutTradeNo        string                   `json:"out_trade_no,omitempty"`
	OutBizNo          string                   `json:"out_biz_no,omitempty"`
	BuyerId           string                   `json:"buyer_id,omitempty"`
	BuyerLogonId      string                   `json:"buyer_logon_id,omitempty"`
	SellerId          string                   `json:"seller_id,omitempty"`
	SellerEmail       string                   `json:"seller_email,omitempty"`
	TradeStatus       string                   `json:"trade_status,omitempty"`
	TotalAmount       string                   `json:"total_amount,omitempty"`
	ReceiptAmount     string                   `json:"receipt_amount,omitempty"`
	InvoiceAmount     string                   `json:"invoice_amount,omitempty"`
	BuyerPayAmount    string                   `json:"buyer_pay_amount,omitempty"`
	PointAmount       string                   `json:"point_amount,omitempty"`
	RefundFee         string                   `json:"refund_fee,omitempty"`
	Subject           string                   `json:"subject,omitempty"`
	Body              string                   `json:"body,omitempty"`
	GmtCreate         string                   `json:"gmt_create,omitempty"`
	GmtPayment        string                   `json:"gmt_payment,omitempty"`
	GmtRefund         string                   `json:"gmt_refund,omitempty"`
	GmtClose          string                   `json:"gmt_close,omitempty"`
	FundBillList      []*fundBillListInfo      `json:"fund_bill_list,omitempty"`
	PassbackParams    string                   `json:"passback_params,omitempty"`
	VoucherDetailList []*voucherDetailListInfo `json:"voucher_detail_list,omitempty"`
	Method            string                   `json:"method,omitempty"`    //电脑网站支付 支付宝请求 return_url 同步返回参数
	Timestamp         string                   `json:"timestamp,omitempty"` //电脑网站支付 支付宝请求 return_url 同步返回参数
}

type fundBillListInfo struct {
	Amount      string `json:"amount,omitempty"`
	FundChannel string `json:"fundChannel,omitempty"` //异步通知里是 fundChannel
	BankCode    string `json:"bank_code,omitempty"`
	RealAmount  string `json:"real_amount,omitempty"`
}

type voucherDetailListInfo struct {
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

type AliPayUserPhone struct {
	Code    string `json:"code,omitempty"`
	Msg     string `json:"msg,omitempty"`
	SubCode string `json:"subCode,omitempty"`
	SubMsg  string `json:"subMsg,omitempty"`
	Mobile  string `json:"mobile,omitempty"`
}

//===================================================
type AliPayTradePayResponse struct {
	Response *payResponse `json:"alipay_trade_pay_response,omitempty"`
	SignData string       `json:"-"`
	Sign     string       `json:"sign"`
}

type payResponse struct {
	Code            string `json:"code,omitempty"`
	Msg             string `json:"msg,omitempty"`
	SubCode         string `json:"sub_code,omitempty"`
	SubMsg          string `json:"sub_msg,omitempty"`
	TradeNo         string `json:"trade_no,omitempty"`
	OutTradeNo      string `json:"out_trade_no,omitempty"`
	BuyerLogonId    string `json:"buyer_logon_id,omitempty"`
	SettleAmount    string `json:"settle_amount,omitempty"`
	PayCurrency     string `json:"pay_currency,omitempty"`
	PayAmount       string `json:"pay_amount,omitempty"`
	SettleTransRate string `json:"settle_trans_rate,omitempty"`
	TransPayRate    string `json:"trans_pay_rate,omitempty"`
	TotalAmount     string `json:"total_amount,omitempty"`
	TransCurrency   string `json:"trans_currency,omitempty"`
	SettleCurrency  string `json:"settle_currency,omitempty"`
	ReceiptAmount   string `json:"receipt_amount,omitempty"`
	BuyerPayAmount  string `json:"buyer_pay_amount,omitempty"`
	PointAmount     string `json:"point_amount,omitempty"`
	InvoiceAmount   string `json:"invoice_amount,omitempty"`
	GmtPayment      string `json:"gmt_payment,omitempty"`
	FundBillList    []*struct {
		FundChannel string `json:"fund_channel,omitempty"`
		BankCode    string `json:"bank_code,omitempty"`
		Amount      string `json:"amount,omitempty"`
		RealAmount  string `json:"real_amount,omitempty"`
	} `json:"fund_bill_list"`
	CardBalance         string `json:"card_balance,omitempty"`
	StoreName           string `json:"store_name,omitempty"`
	BuyerUserId         string `json:"buyer_user_id,omitempty"`
	DiscountGoodsDetail string `json:"discount_goods_detail,omitempty"`
	VoucherDetailList   []*struct {
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
	} `json:"voucher_detail_list"`
	AdvanceAmount    string `json:"advance_amount,omitempty"`
	AuthTradePayMode string `json:"auth_trade_pay_mode,omitempty"`
	ChargeAmount     string `json:"charge_amount,omitempty"`
	ChargeFlags      string `json:"charge_flags,omitempty"`
	SettlementId     string `json:"settlement_id,omitempty"`
	BusinessParams   string `json:"business_params,omitempty"`
	BuyerUserType    string `json:"buyer_user_type,omitempty"`
	MdiscountAmount  string `json:"mdiscount_amount,omitempty"`
	DiscountAmount   string `json:"discount_amount,omitempty"`
	BuyerUserName    string `json:"buyer_user_name,omitempty"`
}

//===================================================
type AliPayTradeQueryResponse struct {
	Response *queryResponse `json:"alipay_trade_query_response,omitempty"`
	SignData string         `json:"-"`
	Sign     string         `json:"sign"`
}

type queryResponse struct {
	Code            string `json:"code,omitempty"`
	Msg             string `json:"msg,omitempty"`
	SubCode         string `json:"sub_code,omitempty"`
	SubMsg          string `json:"sub_msg,omitempty"`
	TradeNo         string `json:"trade_no,omitempty"`
	OutTradeNo      string `json:"out_trade_no,omitempty"`
	BuyerLogonId    string `json:"buyer_logon_id,omitempty"`
	TradeStatus     string `json:"trade_status,omitempty"`
	TotalAmount     string `json:"total_amount,omitempty"`
	TransCurrency   string `json:"trans_currency,omitempty"`
	SettleCurrency  string `json:"settle_currency,omitempty"`
	SettleAmount    string `json:"settle_amount,omitempty"`
	PayCurrency     string `json:"pay_currency,omitempty"`
	PayAmount       string `json:"pay_amount,omitempty"`
	SettleTransRate string `json:"settle_trans_rate,omitempty"`
	TransPayRate    string `json:"trans_pay_rate,omitempty"`
	BuyerPayAmount  string `json:"buyer_pay_amount,omitempty"`
	PointAmount     string `json:"point_amount,omitempty"`
	InvoiceAmount   string `json:"invoice_amount,omitempty"`
	SendPayDate     string `json:"send_pay_date,omitempty"`
	ReceiptAmount   string `json:"receipt_amount,omitempty"`
	StoreId         string `json:"store_id,omitempty"`
	TerminalId      string `json:"terminal_id,omitempty"`
	FundBillList    []*struct {
		FundChannel string `json:"fund_channel,omitempty"`
		BankCode    string `json:"bank_code,omitempty"`
		Amount      string `json:"amount,omitempty"`
		RealAmount  string `json:"real_amount,omitempty"`
	} `json:"fund_bill_list"`
	StoreName       string `json:"store_name,omitempty"`
	BuyerUserId     string `json:"buyer_user_id,omitempty"`
	ChargeAmount    string `json:"charge_amount,omitempty"`
	ChargeFlags     string `json:"charge_flags,omitempty"`
	SettlementId    string `json:"settlement_id,omitempty"`
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

//===================================================
type AliPayTradeCreateResponse struct {
	Response *createResponse `json:"alipay_trade_create_response,omitempty"`
	SignData string          `json:"-"`
	Sign     string          `json:"sign"`
}

type createResponse struct {
	Code       string `json:"code,omitempty"`
	Msg        string `json:"msg,omitempty"`
	SubCode    string `json:"sub_code,omitempty"`
	SubMsg     string `json:"sub_msg,omitempty"`
	TradeNo    string `json:"trade_no,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
}

//===================================================
type AliPayTradeCloseResponse struct {
	Response *closeResponse `json:"alipay_trade_close_response,omitempty"`
	SignData string         `json:"-"`
	Sign     string         `json:"sign"`
}

type closeResponse struct {
	Code       string `json:"code,omitempty"`
	Msg        string `json:"msg,omitempty"`
	SubCode    string `json:"sub_code,omitempty"`
	SubMsg     string `json:"sub_msg,omitempty"`
	TradeNo    string `json:"trade_no,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
}

//===================================================
type AliPayTradeCancelResponse struct {
	Response     *cancelResponse `json:"alipay_trade_cancel_response,omitempty"`
	AliPayCertSn string          `json:"alipay_cert_sn,omitempty"`
	SignData     string          `json:"-"`
	Sign         string          `json:"sign"`
}

type cancelResponse struct {
	Code               string `json:"code,omitempty"`
	Msg                string `json:"msg,omitempty"`
	SubCode            string `json:"sub_code,omitempty"`
	SubMsg             string `json:"sub_msg,omitempty"`
	TradeNo            string `json:"trade_no,omitempty"`
	OutTradeNo         string `json:"out_trade_no,omitempty"`
	RetryFlag          string `json:"retry_flag,omitempty"`
	Action             string `json:"action,omitempty"`
	GmtRefundPay       string `json:"gmt_refund_pay,omitempty"`
	RefundSettlementId string `json:"refund_settlement_id,omitempty"`
}

//===================================================
type AliPaySystemOauthTokenResponse struct {
	Response      *oauthTokenInfo `json:"alipay_system_oauth_token_response,omitempty"`
	ErrorResponse *struct {
		Code    string `json:"code,omitempty"`
		Msg     string `json:"msg,omitempty"`
		SubCode string `json:"sub_code,omitempty"`
		SubMsg  string `json:"sub_msg,omitempty"`
	} `json:"error_response,omitempty"`
	AliPayCertSn string `json:"alipay_cert_sn,omitempty"`
	SignData     string `json:"-"`
	Sign         string `json:"sign"`
}
type oauthTokenInfo struct {
	AccessToken  string `json:"access_token,omitempty"`
	AlipayUserId string `json:"alipay_user_id,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
	ReExpiresIn  int    `json:"re_expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	UserId       string `json:"user_id,omitempty"`
}

//===================================================
type AliPayUserInfoShareResponse struct {
	Response     *userInfoShare `json:"alipay_user_info_share_response,omitempty"`
	AliPayCertSn string         `json:"alipay_cert_sn,omitempty"`
	SignData     string         `json:"-"`
	Sign         string         `json:"sign"`
}

type userInfoShare struct {
	Code               string `json:"code,omitempty"`
	Msg                string `json:"msg,omitempty"`
	SubCode            string `json:"sub_code,omitempty"`
	SubMsg             string `json:"sub_msg,omitempty"`
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

//===================================================
type AliPayTradeRefundResponse struct {
	Response     *refundResponse `json:"alipay_trade_refund_response,omitempty"`
	AliPayCertSn string          `json:"alipay_cert_sn,omitempty"`
	SignData     string          `json:"-"`
	Sign         string          `json:"sign"`
}

type refundResponse struct {
	Code                    string           `json:"code,omitempty"`
	Msg                     string           `json:"msg,omitempty"`
	SubCode                 string           `json:"sub_code,omitempty"`
	SubMsg                  string           `json:"sub_msg,omitempty"`
	TradeNo                 string           `json:"trade_no,omitempty"`
	OutTradeNo              string           `json:"out_trade_no,omitempty"`
	BuyerLogonId            string           `json:"buyer_logon_id,omitempty"`
	FundChange              string           `json:"fund_change,omitempty"`
	RefundFee               string           `json:"refund_fee,omitempty"`
	RefundCurrency          string           `json:"refund_currency,omitempty"`
	GmtRefundPay            string           `json:"gmt_refund_pay,omitempty"`
	RefundDetailItemList    []*tradeFundBill `json:"refund_detail_item_list,omitempty"`
	StoreName               string           `json:"store_name,omitempty"`
	BuyerUserId             string           `json:"buyer_user_id,omitempty"`
	RefundPresetPaytoolList []*struct {
		Amount         []string `json:"amount,omitempty"`
		AssertTypeCode string   `json:"assert_type_code,omitempty"`
	} `json:"refund_preset_paytool_list,omitempty"`
	RefundSettlementId           string `json:"refund_settlement_id,omitempty"`
	PresentRefundBuyerAmount     string `json:"present_refund_buyer_amount,omitempty"`
	PresentRefundDiscountAmount  string `json:"present_refund_discount_amount,omitempty"`
	PresentRefundMdiscountAmount string `json:"present_refund_mdiscount_amount,omitempty"`
}

type tradeFundBill struct {
	FundChannel string `json:"fund_channel,omitempty"` //同步通知里是 fund_channel
	BankCode    string `json:"bank_code,omitempty"`
	Amount      string `json:"amount,omitempty"`
	RealAmount  string `json:"real_amount,omitempty"`
	FundType    string `json:"fund_type,omitempty"`
}

//===================================================
type AliPayTradeFastpayRefundQueryResponse struct {
	Response     *refundQueryResponse `json:"alipay_trade_fastpay_refund_query_response,omitempty"`
	AliPayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type refundQueryResponse struct {
	Code           string `json:"code,omitempty"`
	Msg            string `json:"msg,omitempty"`
	SubCode        string `json:"sub_code,omitempty"`
	SubMsg         string `json:"sub_msg,omitempty"`
	TradeNo        string `json:"trade_no,omitempty"`
	OutTradeNo     string `json:"out_trade_no,omitempty"`
	OutRequestNo   string `json:"out_request_no,omitempty"`
	RefundReason   string `json:"refund_reason,omitempty"`
	TotalAmount    string `json:"total_amount,omitempty"`
	RefundAmount   string `json:"refund_amount,omitempty"`
	RefundRoyaltys []*struct {
		RefundAmount  string `json:"refund_amount,omitempty"`
		RoyaltyType   string `json:"royalty_type,omitempty"`
		ResultCode    string `json:"result_code,omitempty"`
		TransOut      string `json:"trans_out,omitempty"`
		TransOutEmail string `json:"trans_out_email,omitempty"`
		TransIn       string `json:"trans_in,omitempty"`
		TransInEmail  string `json:"trans_in_email,omitempty"`
	} `json:"refund_royaltys,omitempty"`
	GmtRefundPay                 string           `json:"gmt_refund_pay,omitempty"`
	RefundDetailItemList         []*tradeFundBill `json:"refund_detail_item_list,omitempty"`
	SendBackFee                  string           `json:"send_back_fee,omitempty"`
	RefundSettlementId           string           `json:"refund_settlement_id,omitempty"`
	PresentRefundBuyerAmount     string           `json:"present_refund_buyer_amount,omitempty"`
	PresentRefundDiscountAmount  string           `json:"present_refund_discount_amount,omitempty"`
	PresentRefundMdiscountAmount string           `json:"present_refund_mdiscount_amount,omitempty"`
}

//===================================================
type AliPayTradeOrderSettleResponse struct {
	Response     *orderSettleResponse `json:"alipay_trade_order_settle_response,omitempty"`
	AliPayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}
type orderSettleResponse struct {
	Code    string `json:"code,omitempty"`
	Msg     string `json:"msg,omitempty"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
	TradeNo string `json:"trade_no,omitempty"`
}

//===================================================
type AliPayTradePrecreateResponse struct {
	Response     *precreateResponse `json:"alipay_trade_precreate_response,omitempty"`
	AliPayCertSn string             `json:"alipay_cert_sn,omitempty"`
	SignData     string             `json:"-"`
	Sign         string             `json:"sign"`
}

type precreateResponse struct {
	Code       string `json:"code,omitempty"`
	Msg        string `json:"msg,omitempty"`
	SubCode    string `json:"sub_code,omitempty"`
	SubMsg     string `json:"sub_msg,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
	QrCode     string `json:"qr_code,omitempty"`
}

//===================================================
type AliPayTradePageRefundResponse struct {
	Response     *pageRefundResponse `json:"alipay_trade_page_refund_response,omitempty"`
	AliPayCertSn string              `json:"alipay_cert_sn,omitempty"`
	SignData     string              `json:"-"`
	Sign         string              `json:"sign"`
}

type pageRefundResponse struct {
	Code         string `json:"code,omitempty"`
	Msg          string `json:"msg,omitempty"`
	SubCode      string `json:"sub_code,omitempty"`
	SubMsg       string `json:"sub_msg,omitempty"`
	TradeNo      string `json:"trade_no,omitempty"`
	OutTradeNo   string `json:"out_trade_no,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
	RefundAmount string `json:"refund_amount,omitempty"`
}

//===================================================
type AliPayFundTransToaccountTransferResponse struct {
	Response     *transToaccountTransferResponse `json:"alipay_fund_trans_toaccount_transfer_response,omitempty"`
	AliPayCertSn string                          `json:"alipay_cert_sn,omitempty"`
	SignData     string                          `json:"-"`
	Sign         string                          `json:"sign"`
}

type transToaccountTransferResponse struct {
	Code     string `json:"code,omitempty"`
	Msg      string `json:"msg,omitempty"`
	SubCode  string `json:"sub_code,omitempty"`
	SubMsg   string `json:"sub_msg,omitempty"`
	OutBizNo string `json:"out_biz_no,omitempty"`
	OrderId  string `json:"order_id,omitempty"`
	PayDate  string `json:"pay_date,omitempty"`
}

//===================================================
type ZhimaCreditScoreGetResponse struct {
	Response     *scoreGetResponse `json:"zhima_credit_score_get_response,omitempty"`
	AliPayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type scoreGetResponse struct {
	Code    string `json:"code,omitempty"`
	Msg     string `json:"msg,omitempty"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
	BizNo   string `json:"biz_no,omitempty"`
	ZmScore string `json:"zm_score,omitempty"`
}

//===================================================
type AliPayOpenAuthTokenAppResponse struct {
	Response     *authTokenAppResponse `json:"alipay_open_auth_token_app_response,omitempty"`
	AliPayCertSn string                `json:"alipay_cert_sn,omitempty"`
	SignData     string                `json:"-"`
	Sign         string                `json:"sign"`
}

type authTokenAppResponse struct {
	Code            string `json:"code,omitempty"`
	Msg             string `json:"msg,omitempty"`
	SubCode         string `json:"sub_code,omitempty"`
	SubMsg          string `json:"sub_msg,omitempty"`
	UserId          string `json:"user_id,omitempty"`
	AuthAppId       string `json:"auth_app_id,omitempty"`
	AppAuthToken    string `json:"app_auth_token,omitempty"`
	AppRefreshToken string `json:"app_refresh_token,omitempty"`
	ExpiresIn       int    `json:"expires_in,omitempty"`
	ReExpiresIn     int    `json:"re_expires_in,omitempty"`
	Tokens          []*struct {
		AppAuthToken    string `json:"app_auth_token,omitempty"`
		AppRefreshToken string `json:"app_refresh_token,omitempty"`
		AuthAppId       string `json:"auth_app_id,omitempty"`
		ExpiresIn       int    `json:"expires_in,omitempty"`
		ReExpiresIn     int    `json:"re_expires_in,omitempty"`
		UserId          string `json:"user_id,omitempty"`
	} `json:"tokens,omitempty"`
}

//===================================================
type AliPayUserCertifyOpenInitResponse struct {
	Response     *userCertifyOpenInitResponse `json:"alipay_user_certify_open_initialize_response,omitempty"`
	AliPayCertSn string                       `json:"alipay_cert_sn,omitempty"`
	SignData     string                       `json:"-"`
	Sign         string                       `json:"sign"`
}

type userCertifyOpenInitResponse struct {
	Code      string `json:"code,omitempty"`
	Msg       string `json:"msg,omitempty"`
	SubCode   string `json:"sub_code,omitempty"`
	SubMsg    string `json:"sub_msg,omitempty"`
	CertifyId string `json:"certify_id,omitempty"`
}

//===================================================
type AliPayUserCertifyOpenCertifyResponse struct {
	Response     *userCertifyOpenCertifyResponse `json:"alipay_user_certify_open_certify_response,omitempty"`
	AliPayCertSn string                          `json:"alipay_cert_sn,omitempty"`
	SignData     string                          `json:"-"`
	Sign         string                          `json:"sign"`
}

type userCertifyOpenCertifyResponse struct {
	Code    string `json:"code,omitempty"`
	Msg     string `json:"msg,omitempty"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
}

//===================================================
type AliPayUserCertifyOpenQueryResponse struct {
	Response     *userCertifyOpenQueryResponse `json:"alipay_user_certify_open_query_response,omitempty"`
	AliPayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type userCertifyOpenQueryResponse struct {
	Code         string   `json:"code,omitempty"`
	Msg          string   `json:"msg,omitempty"`
	SubCode      string   `json:"sub_code,omitempty"`
	SubMsg       string   `json:"sub_msg,omitempty"`
	Passed       []string `json:"passed,omitempty"`
	IdentityInfo string   `json:"identity_info,omitempty"`
	MaterialInfo string   `json:"material_info,omitempty"`
}
