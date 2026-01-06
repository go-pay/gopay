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

	Address                   string `json:"address"`
	Age                       string `json:"age"`
	Area                      string `json:"area"`
	BusinessScope             string `json:"business_scope"`
	CertNo                    string `json:"cert_no"`
	CertType                  string `json:"cert_type"`
	CollegeName               string `json:"college_name"`
	CountryCode               string `json:"country_code"`
	Degree                    string `json:"degree"`
	DisplayName               string `json:"display_name"`
	Email                     string `json:"email"`
	InstOrCorp                string `json:"inst_or_corp"`
	IsAdult                   string `json:"is_adult"`
	IsBalanceFrozen           string `json:"is_balance_frozen"`
	IsBlocked                 string `json:"is_blocked"`
	IsCertified               string `json:"is_certified"`
	MemberGrade               string `json:"member_grade"`
	Mobile                    string `json:"mobile"`
	OrganizationCode          string `json:"organization_code"`
	PersonBirthday            string `json:"person_birthday"`
	PersonBirthdayWithoutYear string `json:"person_birthday_without_year"`
	Phone                     string `json:"phone"`
	UserName                  string `json:"user_name"`
	UserStatus                string `json:"user_status"`
	UserType                  string `json:"user_type"`
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
