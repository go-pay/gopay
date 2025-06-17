package alipay

type SystemOauthTokenRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UnionId      string `json:"union_id"`
	OpenId       string `json:"open_id"`
	ReExpiresIn  int    `json:"re_expires_in"`
	AuthStart    string `json:"auth_start"`
	ExpiresIn    int    `json:"expires_in"`
}

type UserCertifyOpenQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	Passed       string `json:"passed"`
	IdentityInfo string `json:"identity_info"`
	MaterialInfo string `json:"material_info"`
	FailReason   string `json:"fail_reason"`
}

type UserCertifyOpenInitializeRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CertifyId string `json:"certify_id"`
}

type UserInfoShareRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	UserId   string `json:"user_id"`
	OpenId   string `json:"open_id"`
	Avatar   string `json:"avatar"`
	City     string `json:"city"`
	NickName string `json:"nick_name"`
	Province string `json:"province"`
	Gender   string `json:"gender"`
}

type UserAuthRelationshipQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	QueryDetail string `json:"query_detail"`
}

type UserDelOauthDetailQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	Details []*UserDelOauthDetail `json:"details"`
}

type UserDelOauthDetail struct {
	DelAuthTime string `json:"del_auth_time"`
	UserId      string `json:"user_id"`
	OpenId      string `json:"open_id"`
}

type MobilePhoneNumberDecryptionResp struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"subCode"`
	SubMsg  string `json:"subMsg"`
	Mobile  string `json:"mobile"`
}
