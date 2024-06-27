package alipay

type OcrServerDetectRsp struct {
	Response     *OcrServerDetect `json:"datadigital_fincloud_generalsaas_ocr_server_detect_response"`
	AlipayCertSn string           `json:"alipay_cert_sn,omitempty"`
	SignData     string           `json:"-"`
	Sign         string           `json:"sign"`
}

type OcrMobileInitializeRsp struct {
	Response     *OcrMobileInitialize `json:"datadigital_fincloud_generalsaas_ocr_mobile_initialize_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type OcrCommonDetectRsp struct {
	Response     *OcrCommonDetect `json:"datadigital_fincloud_generalsaas_ocr_common_detect_response"`
	AlipayCertSn string           `json:"alipay_cert_sn,omitempty"`
	SignData     string           `json:"-"`
	Sign         string           `json:"sign"`
}

// =========================================================分割=========================================================

type OcrServerDetect struct {
	ErrorResponse
	OcrData   string `json:"ocr_data"`
	CertifyId string `json:"certify_id"`
}

type OcrMobileInitialize struct {
	ErrorResponse
	CertifyId string `json:"certify_id"`
}

type OcrCommonDetect struct {
	ErrorResponse
	CertifyId string `json:"certify_id"`
	OcrData   string `json:"ocr_data"`
}
