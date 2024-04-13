package alipay

type ZolozAuthenticationSmilepayInitializeRsp struct {
	Response     *ZolozAuthenticationSmilepayInitialize `json:"zoloz_authentication_smilepay_initialize_response"`
	AlipayCertSn string                                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                                 `json:"-"`
	Sign         string                                 `json:"sign"`
}

type ZolozAuthenticationCustomerFtokenQueryRsp struct {
	Response     *ZolozAuthenticationCustomerFtokenQuery `json:"zoloz_authentication_customer_ftoken_query_response"`
	AlipayCertSn string                                  `json:"alipay_cert_sn,omitempty"`
	SignData     string                                  `json:"-"`
	Sign         string                                  `json:"sign"`
}

// =========================================================分割=========================================================

type ZolozAuthenticationSmilepayInitialize struct {
	ErrorResponse
	RetCodeSub        string `json:"ret_code_sub"`
	RetMessageSub     string `json:"ret_message_sub"`
	ZimId             string `json:"zim_id"`
	ZimInitClientData string `json:"zim_init_client_data"`
}

type ZolozAuthenticationCustomerFtokenQuery struct {
	ErrorResponse
	Uid            string        `json:"uid"`
	OpenId         string        `json:"open_id"`
	UidTelPairList []*UidTelPair `json:"uid_tel_pair_list"`
	AgeCheckResult string        `json:"age_check_result"`
	CertNo         string        `json:"cert_no"`
	CertName       string        `json:"cert_name"`
	FaceId         string        `json:"face_id"`
}

type UidTelPair struct {
	UserId string `json:"user_id"`
	OpenId string `json:"open_id"`
}
