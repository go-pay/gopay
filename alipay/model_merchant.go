package alipay

type TradeRelationBindResponse struct {
	Response     *TradeRelationBind `json:"alipay_trade_royalty_relation_bind_response"`
	AlipayCertSn string             `json:"alipay_cert_sn,omitempty"`
	SignData     string             `json:"-"`
	Sign         string             `json:"sign"`
}

type TradeRelationUnbindResponse struct {
	Response     *TradeRelationBind `json:"alipay_trade_royalty_relation_unbind_response"`
	AlipayCertSn string             `json:"alipay_cert_sn,omitempty"`
	SignData     string             `json:"-"`
	Sign         string             `json:"sign"`
}

type TradeRelationBatchQueryResponse struct {
	Response     *TradeRelationBatchQuery `json:"alipay_trade_order_settle_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type TradeSettleConfirmResponse struct {
	Response     *TradeOrderSettle `json:"alipay_trade_settle_confirm_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type TradeOrderSettleResponse struct {
	Response     *TradeOrderSettle `json:"alipay_trade_order_settle_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type TradeOrderSettleQueryResponse struct {
	Response     *TradeOrderSettleQuery `json:"alipay_trade_order_settle_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type TradeOrderOnSettleQueryResponse struct {
	Response     *TradeOrderOnSettleQuery `json:"alipay_trade_order_onsettle_query_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type TradeRoyaltyRateQueryResponse struct {
	Response     *TradeRoyaltyRateQuery `json:"alipay_trade_royalty_rate_query_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

// =========================================================分割=========================================================

type TradeRelationBind struct {
	ErrorResponse
	ResultCode string `json:"result_code"`
}

type TradeRelationBatchQuery struct {
	ErrorResponse
	ResultCode      string      `json:"result_code"`
	ReceiverList    []*Receiver `json:"receiver_list"`
	TotalPageNum    int         `json:"total_page_num"`
	TotalRecordNum  int         `json:"total_record_num"`
	CurrentPageNum  int         `json:"current_page_num"`
	CurrentPageSize int         `json:"current_page_size"`
}

type Receiver struct {
	Type          string `json:"type,omitempty"`
	Account       string `json:"account,omitempty"`
	AccountOpenId string `json:"account_open_id,omitempty"`
	Memo          string `json:"memo,omitempty"`
	LoginName     string `json:"login_name,omitempty"`
	BindLoginName string `json:"bind_login_name,omitempty"`
}

type TradeOrderSettle struct {
	ErrorResponse
	TradeNo string `json:"trade_no,omitempty"`
}

type TradeOrderSettleQuery struct {
	ErrorResponse
	OutTradeNo        string           `json:"out_request_no"`
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
	TransInType    string `json:"trans_in_type"`
	TransInOpenId  string `json:"trans_in_open_id"`
	Amount         string `json:"amount"`
	State          string `json:"state"`
	DetailId       string `json:"detail_id"`
	ErrorCode      string `json:"error_code"`
	ErrorDesc      string `json:"error_desc"`
}

type TradeOrderOnSettleQuery struct {
	ErrorResponse
	UnsettledAmount string `json:"unsettled_amount,omitempty"`
}

type TradeRoyaltyRateQuery struct {
	ErrorResponse
	UserId   string `json:"user_id,omitempty"`
	MaxRatio string `json:"max_ratio,omitempty"`
}
