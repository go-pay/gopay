package alipay

type MarketingActivityDeliveryCreateRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	DeliveryId              string `json:"delivery_id"`
	DeliveryGuidePreviewUrl string `json:"delivery_guide_preview_url"`
}

type MarketingActivityDeliveryQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	DeliveryId        string `json:"delivery_id"`
	DeliveryStatus    string `json:"delivery_status"`
	DeliveryBoothCode string `json:"delivery_booth_code"`
	DeliveryErrorMsg  string `json:"delivery_error_msg"`
	DeliveryBaseInfo  struct {
		DeliveryName      string `json:"delivery_name"`
		DeliveryBeginTime string `json:"delivery_begin_time"`
		DeliveryEndTime   string `json:"delivery_end_time"`
		DeliveryMaterial  struct {
			DeliverySingleMaterial struct {
				DeliveryImage string `json:"delivery_image"`
			} `json:"delivery_single_material"`
		} `json:"delivery_material"`
	} `json:"delivery_base_info"`
	DeliveryPlayConfig struct {
		DeliverySingleSendConfig struct {
			DeliveryContentInfo struct {
				DeliveryContentType     string `json:"delivery_content_type"`
				DeliveryActivityContent struct {
					ActivityId string `json:"activity_id"`
				} `json:"delivery_activity_content"`
			} `json:"delivery_content_info"`
			DeliveryContentConfig struct {
				DeliverySendGuide struct {
					DeliveryGuideUrl string `json:"delivery_guide_url"`
				} `json:"delivery_send_guide"`
			} `json:"delivery_content_config"`
		} `json:"delivery_single_send_config"`
	} `json:"delivery_play_config"`
	DeliveryTargetRule struct {
		DeliveryCityCodeRule struct {
			CityCodes []string `json:"city_codes"`
			AllCity   bool     `json:"all_city"`
		} `json:"delivery_city_code_rule"`
		DeliveryRecallMode string `json:"delivery_recall_mode"`
	} `json:"delivery_target_rule"`
}

type MarketingActivityDeliveryStopRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	DeliveryId string `json:"delivery_id"`
}

type MarketingMaterialImageUploadRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ResourceId      string `json:"resource_id"`
	ResourceEnhance bool   `json:"resource_enhance"`
	ResourceUrl     string `json:"resource_url"`
}

type MarketingCampaignCashCreateRsp struct {
	StatusCode  int
	ErrResponse ErrResponse `json:"-"`

	CrowdNo       string `json:"crowd_no"`
	PayUrl        string `json:"pay_url"`
	OriginCrowdNo string `json:"origin_crowd_no"`
}

type MarketingCampaignCashTriggerRsp struct {
	StatusCode  int
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
	StatusCode  int
	ErrResponse ErrResponse `json:"-"`
}

type MarketingCampaignCashListQueryRsp struct {
	StatusCode  int
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
	StatusCode  int
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
