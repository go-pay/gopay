package alipay

type TradeCustomsDeclareRsp struct {
	Response     *TradeCustomsDeclare `json:"alipay_trade_customs_declare_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

// =========================================================分割=========================================================

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
