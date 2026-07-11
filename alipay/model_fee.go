package alipay

// =========================================================分割=========================================================

// 特殊费率申请 Response
type OpenFeeAdjustApplyResponse struct {
	Response     *OpenFeeAdjustApply `json:"alipay_open_fee_adjust_apply_response"`
	AlipayCertSn string              `json:"alipay_cert_sn,omitempty"`
	SignData     string              `json:"-"`
	Sign         string              `json:"sign"`
}

// =========================================================分割=========================================================

type OpenFeeAdjustApply struct {
	ErrorResponse
}
