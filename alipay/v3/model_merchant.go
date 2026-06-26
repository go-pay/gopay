package alipay

type UserAgreementQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	PrincipalId         string           `json:"principal_id"`
	PrincipalOpenId     string           `json:"principal_open_id"`
	ValidTime           string           `json:"valid_time"`
	AlipayLogonId       string           `json:"alipay_logon_id"`
	InvalidTime         string           `json:"invalid_time"`
	PricipalType        string           `json:"pricipal_type"`
	DeviceId            string           `json:"device_id"`
	SignScene           string           `json:"sign_scene"`
	AgreementNo         string           `json:"agreement_no"`
	ThirdPartyType      string           `json:"third_party_type"`
	Status              string           `json:"status"`
	SignTime            string           `json:"sign_time"`
	PersonalProductCode string           `json:"personal_product_code"`
	ExternalAgreementNo string           `json:"external_agreement_no"`
	ZmOpenId            string           `json:"zm_open_id"`
	ExternalLogonId     string           `json:"external_logon_id"`
	CreditAuthMode      string           `json:"credit_auth_mode"`
	SingleQuota         string           `json:"single_quota"`
	LastDeductTime      string           `json:"last_deduct_time"`
	NextDeductTime      string           `json:"next_deduct_time"`
	ExecutionPlans      []*ExecutionPlan `json:"execution_plans"`
}

type ExecutionPlan struct {
	SingleAmount      string `json:"single_amount"`
	PeriodId          string `json:"period_id"`
	ExecuteTime       string `json:"execute_time"`
	LatestExecuteTime string `json:"latest_execute_time"`
}

type UserAgreementPageUnSignRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

type TradeRelationBindRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ResultCode string `json:"result_code"`
}

type TradeRelationUnbindRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ResultCode string `json:"result_code"`
}

type TradeRelationBatchQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ResultCode      string      `json:"result_code"`
	ReceiverList    []*Receiver `json:"receiver_list"`
	TotalPageNum    int         `json:"total_page_num"`
	TotalRecordNum  int         `json:"total_record_num"`
	CurrentPageNum  int         `json:"current_page_num"`
	CurrentPageSize int         `json:"current_page_size"`
}

type Receiver struct {
	Type          string `json:"type"`
	Account       string `json:"account"`
	AccountOpenId string `json:"account_open_id"`
	Memo          string `json:"memo"`
	LoginName     string `json:"login_name"`
	BindLoginName string `json:"bind_login_name"`
}

type TradeRoyaltyRateQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	UserId   string `json:"user_id"`
	MaxRatio int    `json:"max_ratio"`
}

type TradeOrderSettleRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	TradeNo  string `json:"trade_no"`
	SettleNo string `json:"settle_no"`
}

type TradeOrderSettleQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	OutRequestNo      string           `json:"out_request_no"`
	OperationDt       string           `json:"operation_dt"`
	RoyaltyDetailList []*RoyaltyDetail `json:"royalty_detail_list"`
}

type RoyaltyDetail struct {
	OperationType  string `json:"operation_type"`
	ExecuteDt      string `json:"execute_dt"`
	TransOut       string `json:"trans_out"`
	TransOutType   string `json:"trans_out_type"`
	TransOutOpenId string `json:"trans_out_open_id"`
	TransIn        string `json:"trans_in"`
	TransInOpenId  string `json:"trans_in_open_id"`
	TransInType    string `json:"trans_in_type"`
	Amount         string `json:"amount"`
	State          string `json:"state"`
	DetailId       string `json:"detail_id"`
	ErrorCode      string `json:"error_code"`
	ErrorDesc      string `json:"error_desc"`
}

type TradeOrderOnSettleQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	UnsettledAmount string `json:"unsettled_amount"`
}