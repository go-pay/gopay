package alipay

type MarketingCampaignCashCreateRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CrowdNo       string `json:"crowd_no"`
	PayUrl        string `json:"pay_url"`
	OriginCrowdNo string `json:"origin_crowd_no"`
}
