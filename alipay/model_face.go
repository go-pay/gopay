package alipay

type FaceVerificationInitializeRsp struct {
	Response     *FaceVerificationInitialize `json:"datadigital_fincloud_generalsaas_face_verification_initialize_response"`
	AlipayCertSn string                      `json:"alipay_cert_sn,omitempty"`
	SignData     string                      `json:"-"`
	Sign         string                      `json:"sign"`
}

type FaceVerificationQueryRsp struct {
	Response     *FaceVerificationQuery `json:"datadigital_fincloud_generalsaas_face_verification_query_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type FaceCertifyInitializeRsp struct {
	Response     *FaceCertifyInitialize `json:"datadigital_fincloud_generalsaas_face_certify_initialize_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type FaceCertifyVerifyRsp struct {
	Response     *FaceCertifyVerify `json:"datadigital_fincloud_generalsaas_face_certify_verify_response"`
	AlipayCertSn string             `json:"alipay_cert_sn,omitempty"`
	SignData     string             `json:"-"`
	Sign         string             `json:"sign"`
}

type FaceCertifyQueryRsp struct {
	Response     *FaceCertifyQuery `json:"datadigital_fincloud_generalsaas_face_certify_query_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type FaceSourceCertifyRsp struct {
	Response     *FaceSourceCertify `json:"datadigital_fincloud_generalsaas_face_source_certify_response"`
	AlipayCertSn string             `json:"alipay_cert_sn,omitempty"`
	SignData     string             `json:"-"`
	Sign         string             `json:"sign"`
}

type FaceCheckInitializeRsp struct {
	Response     *FaceCheckInitialize `json:"datadigital_fincloud_generalsaas_face_check_initialize_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type FaceCheckQueryRsp struct {
	Response     *FaceCheckQuery `json:"datadigital_fincloud_generalsaas_face_check_query_response"`
	AlipayCertSn string          `json:"alipay_cert_sn,omitempty"`
	SignData     string          `json:"-"`
	Sign         string          `json:"sign"`
}

// =========================================================分割=========================================================

type FaceVerificationInitialize struct {
	ErrorResponse
	CertifyId string `json:"certify_id"`
	WebUrl    string `json:"web_url"`
}

type FaceVerificationQuery struct {
	ErrorResponse
	CertifyState string `json:"certify_state"`
	Score        string `json:"score"`
	Quality      string `json:"quality"`
	AlivePhoto   string `json:"alive_photo"`
	AttackFlag   string `json:"attack_flag"`
}

type FaceCertifyInitialize struct {
	ErrorResponse
	CertifyId string `json:"certify_id"`
}

type FaceCertifyVerify struct {
	ErrorResponse
	CertifyUrl string `json:"certify_url"`
}

type FaceCertifyQuery struct {
	ErrorResponse
	Passed string `json:"passed"`
}

type FaceSourceCertify struct {
	ErrorResponse
	CertifyNo      string `json:"certify_no"`
	Passed         string `json:"passed"`
	Score          string `json:"score"`
	Quality        string `json:"quality"`
	MismatchReason string `json:"mismatch_reason"`
}

type FaceCheckInitialize struct {
	ErrorResponse
	CertifyId string `json:"certify_id"`
	WebUrl    string `json:"web_url"`
}

type FaceCheckQuery struct {
	ErrorResponse
	CertifyState string `json:"certify_state"`
	Quality      string `json:"quality"`
	AlivePhoto   string `json:"alive_photo"`
}
