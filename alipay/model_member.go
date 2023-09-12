package alipay

type SystemOauthTokenResponse struct {
	Response      *OauthTokenInfo `json:"alipay_system_oauth_token_response"`
	ErrorResponse *ErrorResponse  `json:"error_response,omitempty"`
	AlipayCertSn  string          `json:"alipay_cert_sn,omitempty"`
	SignData      string          `json:"-"`
	Sign          string          `json:"sign"`
}

type UserInfoShareResponse struct {
	Response      *UserInfoShare `json:"alipay_user_info_share_response"`
	ErrorResponse *ErrorResponse `json:"error_response,omitempty"`
	AlipayCertSn  string         `json:"alipay_cert_sn,omitempty"`
	SignData      string         `json:"-"`
	Sign          string         `json:"sign"`
}

type UserCertifyOpenInitResponse struct {
	Response     *UserCertifyOpenInit `json:"alipay_user_certify_open_initialize_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type UserCertifyOpenQueryResponse struct {
	Response     *UserCertifyOpenQuery `json:"alipay_user_certify_open_query_response"`
	AlipayCertSn string                `json:"alipay_cert_sn,omitempty"`
	SignData     string                `json:"-"`
	Sign         string                `json:"sign"`
}

type UserAgreementPageUnSignRsp struct {
	Response     *ErrorResponse `json:"alipay_user_agreement_unsign_response"`
	AlipayCertSn string         `json:"alipay_cert_sn,omitempty"`
	SignData     string         `json:"-"`
	Sign         string         `json:"sign"`
}

type UserAgreementQueryRsp struct {
	Response     *UserAgreementQuery `json:"alipay_user_agreement_query_response"`
	AlipayCertSn string              `json:"alipay_cert_sn,omitempty"`
	SignData     string              `json:"-"`
	Sign         string              `json:"sign"`
}

type UserAgreementExecutionplanModifyRsp struct {
	Response     *UserAgreementExecutionplanModify `json:"alipay_user_agreement_executionplan_modify_response"`
	AlipayCertSn string                            `json:"alipay_cert_sn,omitempty"`
	SignData     string                            `json:"-"`
	Sign         string                            `json:"sign"`
}

type UserAgreementTransferRsp struct {
	Response     *UserAgreementTransfer `json:"alipay_user_agreement_transfer_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type UserTwostageCommonUseRsp struct {
	Response     *UserTwostageCommonUse `json:"alipay_user_twostage_common_use_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type UserAuthZhimaorgIdentityApplyRsp struct {
	Response     *UserAuthZhimaorgIdentityApply `json:"alipay_user_auth_zhimaorg_identity_apply_response"`
	AlipayCertSn string                         `json:"alipay_cert_sn,omitempty"`
	SignData     string                         `json:"-"`
	Sign         string                         `json:"sign"`
}

type UserCharityRecordexistQueryRsp struct {
	Response     *UserCharityRecordexistQuery `json:"alipay_user_charity_recordexist_query_response"`
	AlipayCertSn string                       `json:"alipay_cert_sn,omitempty"`
	SignData     string                       `json:"-"`
	Sign         string                       `json:"sign"`
}

type UserAlipaypointSendRsp struct {
	Response     *UserAlipaypointSend `json:"alipay_user_alipaypoint_send_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type MemberDataIsvCreateRsp struct {
	Response     *MemberDataIsvCreate `json:"koubei_member_data_isv_create_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type UserFamilyArchiveQueryRsp struct {
	Response     *UserFamilyArchiveQuery `json:"alipay_user_family_archive_query_response"`
	AlipayCertSn string                  `json:"alipay_cert_sn,omitempty"`
	SignData     string                  `json:"-"`
	Sign         string                  `json:"sign"`
}

type UserFamilyArchiveInitializeRsp struct {
	Response     *UserFamilyArchiveInitialize `json:"alipay_user_family_archive_initialize_response"`
	AlipayCertSn string                       `json:"alipay_cert_sn,omitempty"`
	SignData     string                       `json:"-"`
	Sign         string                       `json:"sign"`
}

type UserCertdocCertverifyPreconsultRsp struct {
	Response     *UserCertdocCertverifyPreconsult `json:"alipay_user_certdoc_certverify_preconsult_response"`
	AlipayCertSn string                           `json:"alipay_cert_sn,omitempty"`
	SignData     string                           `json:"-"`
	Sign         string                           `json:"sign"`
}

type UserCertdocCertverifyConsultRsp struct {
	Response     *UserCertdocCertverifyConsult `json:"alipay_user_certdoc_certverify_consult_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type UserFamilyShareZmgoInitializeRsp struct {
	Response     *UserFamilyShareZmgoInitialize `json:"alipay_user_family_share_zmgo_initialize_response"`
	AlipayCertSn string                         `json:"alipay_cert_sn,omitempty"`
	SignData     string                         `json:"-"`
	Sign         string                         `json:"sign"`
}

type UserDtbankQrcodedataQueryRsp struct {
	Response     *UserDtbankQrcodedataQuery `json:"alipay_user_dtbank_qrcodedata_query_response"`
	AlipayCertSn string                     `json:"alipay_cert_sn,omitempty"`
	SignData     string                     `json:"-"`
	Sign         string                     `json:"sign"`
}

type UserAlipaypointBudgetlibQueryRsp struct {
	Response     *UserAlipaypointBudgetlibQuery `json:"alipay_user_alipaypoint_budgetlib_query_response"`
	AlipayCertSn string                         `json:"alipay_cert_sn,omitempty"`
	SignData     string                         `json:"-"`
	Sign         string                         `json:"sign"`
}

// =========================================================分割=========================================================

type OauthTokenInfo struct {
	UserId        string `json:"user_id,omitempty"`
	OpenId        string `json:"open_id,omitempty"`
	AlipayUserId  string `json:"alipay_user_id,omitempty"`
	UnionId       string `json:"union_id,omitempty"`
	AccessToken   string `json:"access_token,omitempty"`
	ExpiresIn     int    `json:"expires_in,omitempty"`
	RefreshToken  string `json:"refresh_token,omitempty"`
	ReExpiresIn   int    `json:"re_expires_in,omitempty"`
	AuthStart     string `json:"auth_start,omitempty"`
	AuthTokenType string `json:"auth_token_type,omitempty"`
}

type UserInfoShare struct {
	UserId             string `json:"user_id,omitempty"`
	OpenId             string `json:"open_id,omitempty"`
	Avatar             string `json:"avatar,omitempty"`
	Province           string `json:"province,omitempty"`
	City               string `json:"city,omitempty"`
	NickName           string `json:"nick_name,omitempty"`
	IsStudentCertified string `json:"is_student_certified,omitempty"`
	UserType           string `json:"user_type,omitempty"`
	UserStatus         string `json:"user_status,omitempty"`
	IsCertified        string `json:"is_certified,omitempty"`
	Gender             string `json:"gender,omitempty"`
}

type UserCertifyOpenInit struct {
	ErrorResponse
	CertifyId string `json:"certify_id,omitempty"`
}

type UserCertifyOpenQuery struct {
	ErrorResponse
	Passed       string `json:"passed,omitempty"`
	IdentityInfo string `json:"identity_info,omitempty"`
	MaterialInfo string `json:"material_info,omitempty"`
}

type UserAgreementQuery struct {
	ErrorResponse
	PrincipalId         string `json:"principal_id"`
	PrincipalOpenId     string `json:"principal_open_id,omitempty"`
	ValidTime           string `json:"valid_time"`
	AlipayLogonId       string `json:"alipay_logon_id"`
	InvalidTime         string `json:"invalid_time"`
	PricipalType        string `json:"pricipal_type"`
	DeviceId            string `json:"device_id,omitempty"`
	SignScene           string `json:"sign_scene"`
	AgreementNo         string `json:"agreement_no"`
	ThirdPartyType      string `json:"third_party_type"`
	Status              string `json:"status"`
	SignTime            string `json:"sign_time"`
	PersonalProductCode string `json:"personal_product_code"`
	ExternalAgreementNo string `json:"external_agreement_no,omitempty"`
	ZmOpenId            string `json:"zm_open_id,omitempty"`
	ExternalLogonId     string `json:"external_logon_id,omitempty"`
	CreditAuthMode      string `json:"credit_auth_mode,omitempty"`
	SingleQuota         string `json:"single_quota,omitempty"`
	LastDeductTime      string `json:"last_deduct_time,omitempty"`
	NextDeductTime      string `json:"next_deduct_time,omitempty"`
}

type UserAgreementExecutionplanModify struct {
	ErrorResponse
	AgreementNo string `json:"agreement_no"`
	DeductTime  string `json:"deduct_time"`
}

type UserAgreementTransfer struct {
	ErrorResponse
	ExecuteTime   string `json:"execute_time,omitempty"`
	PeriodType    string `json:"period_type,omitempty"`
	Amount        string `json:"amount,omitempty"`
	TotalAmount   string `json:"total_amount,omitempty"`
	TotalPayments string `json:"total_payments,omitempty"`
	Period        string `json:"period,omitempty"`
}

type UserTwostageCommonUse struct {
	ErrorResponse
	UserId           string              `json:"user_id,omitempty"`
	OpenId           string              `json:"open_id,omitempty"`
	UserIdentityInfo []*UserIdentityInfo `json:"user_identity_info,omitempty"`
}

type UserIdentityInfo struct {
	HSchoolInfo []*HSchoolInfo `json:"h_school_info,omitempty"`
}

type HSchoolInfo struct {
	SchoolStdCode string `json:"school_std_code"`
	CampusNo      string `json:"campus_no"`
}

type UserAuthZhimaorgIdentityApply struct {
	ErrorResponse
	AccessToken   string `json:"access_token"`
	AuthTokenType string `json:"auth_token_type,omitempty"`
	RefreshToken  string `json:"refresh_token"`
}

type UserCharityRecordexistQuery struct {
	ErrorResponse
	DonationTag string `json:"donation_tag"`
}

type UserAlipaypointSend struct {
	ErrorResponse
	RecordId string `json:"record_id"`
}

type MemberDataIsvCreate struct {
	ErrorResponse
}

type UserFamilyArchiveQuery struct {
	ErrorResponse
	ArchiveList []*FamilyArchiveDetail `json:"archive_list"`
}

type FamilyArchiveDetail struct {
	ArchiveId            string `json:"archive_id"`
	RealName             string `json:"real_name,omitempty"`
	CertNo               string `json:"cert_no,omitempty"`
	CertType             string `json:"cert_type,omitempty"`
	Mobile               string `json:"mobile,omitempty"`
	Email                string `json:"email,omitempty"`
	Role                 string `json:"role,omitempty"`
	Province             string `json:"province,omitempty"`
	City                 string `json:"city,omitempty"`
	DesensitizedLogonId  string `json:"desensitized_logon_id,omitempty"`
	Area                 string `json:"area,omitempty"`
	DesensitizedRealName string `json:"desensitized_real_name,omitempty"`
	Address              string `json:"address,omitempty"`
	Zip                  string `json:"zip,omitempty"`
	Birthday             string `json:"birthday,omitempty"`
	Gender               string `json:"gender,omitempty"`
	Profession           string `json:"profession,omitempty"`
}

type UserFamilyArchiveInitialize struct {
	ErrorResponse
	ArchivePluginUrl string `json:"archive_plugin_url"`
}

type UserCertdocCertverifyPreconsult struct {
	ErrorResponse
	VerifyId string `json:"verify_id"`
}

type UserCertdocCertverifyConsult struct {
	ErrorResponse
	Passed     string `json:"passed"`
	FailReason string `json:"fail_reason,omitempty"`
	FailParams string `json:"fail_params,omitempty"`
}

type UserFamilyShareZmgoInitialize struct {
	ErrorResponse
	Shareable         bool   `json:"shareable"`
	HasSharing        bool   `json:"has_sharing"`
	FamilySharingLink string `json:"family_sharing_link"`
}

type UserDtbankQrcodedataQuery struct {
	ErrorResponse
	DataDate           string `json:"data_date,omitempty"`
	QrcodeId           string `json:"qrcode_id,omitempty"`
	QrcodeOutId        string `json:"qrcode_out_id,omitempty"`
	BindCard           string `json:"bind_card,omitempty"`
	SendVoucherCnt     string `json:"send_voucher_cnt,omitempty"`
	SendVoucherAmt     string `json:"send_voucher_amt,omitempty"`
	WriteOffVoucherCnt string `json:"write_off_voucher_cnt,omitempty"`
	WriteOffVoucherAmt string `json:"write_off_voucher_amt,omitempty"`
	LeadToFollow       string `json:"lead_to_follow,omitempty"`
}

type UserAlipaypointBudgetlibQuery struct {
	ErrorResponse
	BudgetCode       string `json:"budget_code"`
	BudgetDesc       string `json:"budget_desc"`
	Enabled          bool   `json:"enabled"`
	CumulativeAmount int64  `json:"cumulative_amount"`
	RemainAmount     int64  `json:"remain_amount"`
	StartTime        string `json:"start_time"`
	EndTime          string `json:"end_time"`
}
