package alipay

type MarketingCampaignCashCreateRsp struct {
	Response     *MarketingCampaignCashCreate `json:"alipay_marketing_campaign_cash_create_response"`
	AlipayCertSn string                       `json:"alipay_cert_sn,omitempty"`
	SignData     string                       `json:"-"`
	Sign         string                       `json:"sign"`
}

type MarketingCampaignCashCreate struct {
	ErrorResponse
	CrowdNo       string `json:"crowd_no"`
	PayURL        string `json:"pay_url"`
	OriginCrowdNo string `json:"origin_crowd_no"`
}

type MarketingCampaignCashTriggerRsp struct {
	Response     *MarketingCampaignCashTrigger `json:"alipay_marketing_campaign_cash_trigger_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type MarketingCampaignCashTrigger struct {
	ErrorResponse
	TriggerResult     string `json:"trigger_result"`
	PrizeAmount       string `json:"prize_amount"`
	RepeatTriggerFlag string `json:"repeat_trigger_flag"`
	PartnerId         string `json:"partner_id"`
	ErrorMsg          string `json:"error_msg"`
	CouponName        string `json:"coupon_name"`
	PrizeMsg          string `json:"prize_msg"`
	MerchantLogo      string `json:"merchant_logo"`
	BizNo             string `json:"biz_no"`
	OutBizNo          string `json:"out_biz_no"`
}

type MarketingCampaignCashStatusModifyRsp struct {
	Response     *MarketingCampaignCashStatusModify `json:"alipay_marketing_campaign_cash_status_modify_response"`
	AlipayCertSn string                             `json:"alipay_cert_sn,omitempty"`
	SignData     string                             `json:"-"`
	Sign         string                             `json:"sign"`
}

type MarketingCampaignCashStatusModify struct {
	ErrorResponse
}

type MarketingCampaignCashListQueryRsp struct {
	Response     *MarketingCampaignCashListQuery `json:"alipay_marketing_campaign_cash_list_query_response"`
	AlipayCertSn string                          `json:"alipay_cert_sn,omitempty"`
	SignData     string                          `json:"-"`
	Sign         string                          `json:"sign"`
}

type MarketingCampaignCashListQuery struct {
	ErrorResponse
	PageSize  string             `json:"page_size"`
	CampList  []CashCampaignInfo `json:"camp_list"`
	PageIndex string             `json:"page_index"`
	TotalSize string             `json:"total_size"`
}

type CashCampaignInfo struct {
	CrowdNo       string `json:"crowd_no"`
	OriginCrowdNo string `json:"origin_crowd_no"`
	CampStatus    string `json:"camp_status"`
	CouponName    string `json:"coupon_name"`
}

type MarketingCampaignCashDetailQueryRsp struct {
	Response     *MarketingCampaignCashDetailQuery `json:"alipay_marketing_campaign_cash_detail_query_response"`
	AlipayCertSn string                            `json:"alipay_cert_sn,omitempty"`
	SignData     string                            `json:"-"`
	Sign         string                            `json:"sign"`
}

type MarketingCampaignCashDetailQuery struct {
	ErrorResponse
	CrowdNo       string  `json:"crowd_no"`
	CouponName    string  `json:"coupon_name"`
	PrizeMsg      string  `json:"prize_msg"`
	PrizeType     string  `json:"prize_type"`
	StartTime     string  `json:"start_time"`
	EndTime       string  `json:"end_time"`
	TotalAmount   float64 `json:"total_amount"`
	SendAmount    float64 `json:"send_amount"`
	TotalNum      int     `json:"total_num"`
	TotalCount    int     `json:"total_count"`
	OriginCrowdNo string  `json:"origin_crowd_no"`
	CampStatus    string  `json:"camp_status"`
}
