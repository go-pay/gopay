package alipay

type FaceVerificationInitializeRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CertifyId string `json:"certify_id"`
	WebUrl    string `json:"web_url"`
}

type FaceVerificationQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CertifyState string    `json:"certify_state"`
	Score        string    `json:"score"`
	Quality      string    `json:"quality"`
	AlivePhoto   string    `json:"alive_photo"`
	AttackFlag   string    `json:"attack_flag"`
	MetaInfo     *MetaInfo `json:"meta_info"`
}

type MetaInfo struct {
	DeviceType string `json:"device_type"`
}

type FaceCertifyInitializeRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CertifyId string `json:"certify_id"`
}

type FaceCertifyVerifyRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CertifyUrl string `json:"certify_url"`
}

type FaceCertifyQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	Passed string `json:"passed"`
}

type FaceSourceCertifyRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CertifyNo      string `json:"certify_no"`
	Passed         string `json:"passed"`
	Score          string `json:"score"`
	Quality        string `json:"quality"`
	MismatchReason string `json:"mismatch_reason"`
}

type FaceCheckInitializeRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CertifyId string `json:"certify_id"`
	WebUrl    string `json:"web_url"`
}

type FaceCheckQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CertifyState string    `json:"certify_state"`
	Quality      string    `json:"quality"`
	AlivePhoto   string    `json:"alive_photo"`
	MetaInfo     *MetaInfo `json:"meta_info"`
}

type IDCardTwoMetaCheckRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CertifyId string `json:"certify_id"`
	Match     string `json:"match"`
}

type BankCardCheckRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CertifyId string `json:"certify_id"`
	Match     string `json:"match"`
	Detail    string `json:"detail"`
}

type MobileThreeMetaSimpleCheckRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CertifyId string `json:"certify_id"`
	Match     string `json:"match"`
	Isp       string `json:"isp"`
	Detail    string `json:"detail"`
}

type MobileThreeMetaDetailCheckRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CertifyId string `json:"certify_id"`
	Match     string `json:"match"`
	Isp       string `json:"isp"`
	Detail    string `json:"detail"`
}

type OcrServerDetectRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	OcrData   string `json:"ocr_data"`
	CertifyId string `json:"certify_id"`
}

type OcrMobileInitializeRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CertifyId string `json:"certify_id"`
}
