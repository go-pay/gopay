package alipay

type MarketingCampaignCashCreateRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CrowdNo       string `json:"crowd_no"`
	PayUrl        string `json:"pay_url"`
	OriginCrowdNo string `json:"origin_crowd_no"`
}

type MarketingCampaignCashTriggerRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

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
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

type MarketingCampaignCashListQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	PageSize string `json:"page_size"`
	CampList []struct {
		CrowdNo       string `json:"crowd_no"`
		OriginCrowdNo string `json:"origin_crowd_no"`
		CampStatus    string `json:"camp_status"`
		CouponName    string `json:"coupon_name"`
	} `json:"camp_list"`
	PageIndex string `json:"page_index"`
	TotalSize string `json:"total_size"`
}

type MarketingCampaignCashDetailQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

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

type MarketingCampaignOrderVoucherConsultRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	OptimalTotalPromoAmount string                `json:"optimal_total_promo_amount"`
	VoucherConsultList      []*VoucherConsultInfo `json:"voucher_consult_list"`
}

// VoucherConsultInfo 优惠券咨询信息
type VoucherConsultInfo struct {
	VoucherId       string         `json:"voucher_id"`
	VoucherName     string         `json:"voucher_name"`
	VoucherType     string         `json:"voucher_type"`
	PromoType       string         `json:"promo_type"`
	PromoAmount     string         `json:"promo_amount"`
	Optimal         bool           `json:"optimal"`
	ThresholdAmount string         `json:"threshold_amount"`
	ReductionAmount string         `json:"reduction_amount"`
	SpecifiedAmount string         `json:"specified_amount"`
	ReductionRatio  string         `json:"reduction_ratio"`
	CeilingAmount   string         `json:"ceiling_amount"`
	PromoText       string         `json:"promo_text"`
	ItemPromoInfo   *ItemPromoInfo `json:"item_promo_info"`
}

// ItemPromoInfo 单品券商品优惠信息
type ItemPromoInfo struct {
	ItemName          string             `json:"item_name"`
	ItemConsultList   []*ItemConsultInfo `json:"item_consult_list"`
	ItemDesc          string             `json:"item_desc"`
	ItemCoverPic      string             `json:"item_cover_pic"`
	ItemDetailPicList []string           `json:"item_detail_pic_list"`
}

// ItemConsultInfo 商品咨询信息
type ItemConsultInfo struct {
	ItemId      string `json:"item_id"`
	PromoAmount string `json:"promo_amount"`
	PromoCount  string `json:"promo_count"`
}
