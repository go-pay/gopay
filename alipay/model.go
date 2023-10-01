package alipay

const (
	// URL
	baseUrl = "https://openapi.alipay.com/gateway.do"
	//sandboxBaseUrl     = "https://openapi.alipaydev.com/gateway.do"
	sandboxBaseUrl = "https://openapi-sandbox.dl.alipaydev.com/gateway.do"
	baseUrlUtf8    = "https://openapi.alipay.com/gateway.do?charset=utf-8"
	//sandboxBaseUrlUtf8 = "https://openapi.alipaydev.com/gateway.do?charset=utf-8"
	sandboxBaseUrlUtf8 = "https://openapi-sandbox.dl.alipaydev.com/gateway.do?charset=utf-8"

	LocationShanghai          = "Asia/Shanghai"
	PKCS1            PKCSType = 1 // 非Java
	PKCS8            PKCSType = 2 // Java
	RSA                       = "RSA"
	RSA2                      = "RSA2"
	UTF8                      = "utf-8"
)

var (
	//不需要处理AppAuthToken的方法
	appAuthTokenInBizContent = map[string]bool{
		"alipay.open.auth.token.app.query": true,
	}
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
	ExpiresIn       int      `json:"expires_in,omitempty"`
	AppRefreshToken string   `json:"app_refresh_token,omitempty"`
	ReExpiresIn     int      `json:"re_expires_in,omitempty"`
	Tokens          []*Token `json:"tokens,omitempty"`
}

type Token struct {
	AuthAppId       string `json:"auth_app_id,omitempty"`
	AppAuthToken    string `json:"app_auth_token,omitempty"`
	ExpiresIn       int    `json:"expires_in,omitempty"`
	AppRefreshToken string `json:"app_refresh_token,omitempty"`
	ReExpiresIn     int    `json:"re_expires_in,omitempty"`
	UserId          string `json:"user_id,omitempty"`
}

// ===================================================
type OpenAuthTokenAppInviteCreateResponse struct {
	Response     *OpenAuthTokenAppInviteCreate `json:"alipay_open_auth_appauth_invite_create_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type OpenAuthTokenAppInviteCreate struct {
	ErrorResponse
	TaskPageUrl string `json:"task_page_url,omitempty"`
}

// ===================================================
type OpenAuthTokenAppQueryResponse struct {
	Response     *AuthTokenAppQuery `json:"alipay_open_auth_token_app_query_response"`
	AlipayCertSn string             `json:"alipay_cert_sn,omitempty"`
	SignData     string             `json:"-"`
	Sign         string             `json:"sign"`
}

type AuthTokenAppQuery struct {
	ErrorResponse
	UserId      string   `json:"user_id"`      //授权商户的user_id
	AuthAppId   string   `json:"auth_app_id"`  //授权商户的appid
	ExpiresIn   int      `json:"expires_in"`   //应用授权令牌失效时间，单位到秒
	AuthMethods []string `json:"auth_methods"` //当前app_auth_token的授权接口列表
	AuthStart   string   `json:"auth_start"`   //授权生效时间
	AuthEnd     string   `json:"auth_end"`     //授权失效时间
	Status      string   `json:"status"`       //valid：有效状态；invalid：无效状态
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
	Pid string `json:"pid"`
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
	SettleAmount    string `json:"settle_amount,omitempty"`
}

// ===================================================
type DataBillAccountLogQueryResponse struct {
	Response     *DataBillAccountLogQuery `json:"alipay_data_bill_accountlog_query_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type DataBillAccountLogQuery struct {
	ErrorResponse
	PageNo     string                 `json:"page_no,omitempty"`
	PageSize   string                 `json:"page_size,omitempty"`
	TotalSize  string                 `json:"total_size,omitempty"`
	DetailList []AccountLogItemResult `json:"detail_list,omitempty"`
}

type AccountLogItemResult struct {
	TransDt             string `json:"trans_dt,omitempty"`
	AccountLogId        string `json:"account_log_id,omitempty"`
	AlipayOrderNo       string `json:"alipay_order_no,omitempty"`
	MerchantOrderNo     string `json:"merchant_order_no,omitempty"`
	TransAmount         string `json:"trans_amount,omitempty"`
	Balance             string `json:"balance,omitempty"`
	Type                string `json:"type,omitempty"`
	OtherAccount        string `json:"other_account,omitempty"`
	TransMemo           string `json:"trans_memo,omitempty"`
	Direction           string `json:"direction,omitempty"`
	BillSource          string `json:"bill_source,omitempty"`
	BizNos              string `json:"biz_nos,omitempty"`
	BizOrigNo           string `json:"biz_orig_no,omitempty"`
	BizDesc             string `json:"biz_desc,omitempty"`
	MerchantOutRefundNo string `json:"merchant_out_refund_no,omitempty"`
	ComplementInfo      string `json:"complement_info,omitempty"`
	StoreName           string `json:"store_name,omitempty"`
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

// ===================================================
type DataDataserviceAdDataQueryRsp struct {
	Response     *DataDataserviceAdDataQuery `json:"alipay_data_dataservice_ad_data_query_response"`
	AlipayCertSn string                      `json:"alipay_cert_sn,omitempty"`
	SignData     string                      `json:"-"`
	Sign         string                      `json:"sign"`
}

type DataDataserviceAdDataQuery struct {
	ErrorResponse
	DataList []*DataDetail `json:"data_list,omitempty"`
}

type DataDetail struct {
	OuterId            string                  `json:"outer_id,omitempty"`
	Impression         int64                   `json:"impression,omitempty"`
	Click              int64                   `json:"click,omitempty"`
	Cost               int64                   `json:"cost,omitempty"`
	ConversionDataList []*ConversionDataDetail `json:"conversion_data_list,omitempty"`
	BizDate            string                  `json:"biz_date,omitempty"`
}

type ConversionDataDetail struct {
	ConversionId     string `json:"conversion_id,omitempty"`
	ConversionResult string `json:"conversion_result,omitempty"`
}

// ===================================================
type OpenAppApiQueryResponse struct {
	Response     *OpenAppApiQuery `json:"alipay_open_app_api_query_response"`
	AlipayCertSn string           `json:"alipay_cert_sn,omitempty"`
	SignData     string           `json:"-"`
	Sign         string           `json:"sign"`
}

type OpenAppApiQuery struct {
	ErrorResponse
	Apis []*Apis `json:"apis"`
}

type Apis struct {
	ApiName     string `json:"api_name,omitempty"`
	FieldName   string `json:"field_name,omitempty"`
	PackageCode string `json:"package_code,omitempty"`
}
