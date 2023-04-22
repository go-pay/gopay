package alipay

type CommerceTransportNfccardSendRsp struct {
	Response     *CommerceTransportNfccardSend `json:"alipay_commerce_transport_nfccard_send_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type CommerceAirCallcenterTradeApplyRsp struct {
	Response     *CommerceAirCallcenterTradeApply `json:"alipay_commerce_air_callcenter_trade_apply_response"`
	AlipayCertSn string                           `json:"alipay_cert_sn,omitempty"`
	SignData     string                           `json:"-"`
	Sign         string                           `json:"sign"`
}

type CommerceBenefitApplyRsp struct {
	Response     *CommerceBenefitApply `json:"alipay_commerce_operation_gamemarketing_benefit_apply_response"`
	AlipayCertSn string                `json:"alipay_cert_sn,omitempty"`
	SignData     string                `json:"-"`
	Sign         string                `json:"sign"`
}

type CommerceBenefitVerifyRsp struct {
	Response     *CommerceBenefitVerify `json:"alipay_commerce_operation_gamemarketing_benefit_verify_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

// =========================================================分割=========================================================

type CommerceTransportNfccardSend struct {
	ErrorResponse
}

type CommerceAirCallcenterTradeApply struct {
	ErrorResponse
}

type CommerceBenefitApply struct {
	ErrorResponse
	ApplyVoucherCodeList string `json:"apply_voucher_code_list,omitempty"`
}

type CommerceBenefitVerify struct {
	ErrorResponse
	VoucherVerifyStatus bool `json:"voucher_verify_status,omitempty"`
}
