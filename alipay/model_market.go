package alipay

type MarketingCampaignCashCreateRsp struct {
	Response     *MarketingCampaignCashCreate `json:"alipay_marketing_campaign_cash_create_response"`
	AlipayCertSn string                       `json:"alipay_cert_sn,omitempty"`
	SignData     string                       `json:"-"`
	Sign         string                       `json:"sign"`
}

type MarketingCampaignCashTriggerRsp struct {
	Response     *MarketingCampaignCashTrigger `json:"alipay_marketing_campaign_cash_trigger_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type MarketingCampaignCashStatusModifyRsp struct {
	Response     *MarketingCampaignCashStatusModify `json:"alipay_marketing_campaign_cash_status_modify_response"`
	AlipayCertSn string                             `json:"alipay_cert_sn,omitempty"`
	SignData     string                             `json:"-"`
	Sign         string                             `json:"sign"`
}

type MarketingCampaignCashListQueryRsp struct {
	Response     *MarketingCampaignCashListQuery `json:"alipay_marketing_campaign_cash_list_query_response"`
	AlipayCertSn string                          `json:"alipay_cert_sn,omitempty"`
	SignData     string                          `json:"-"`
	Sign         string                          `json:"sign"`
}

type MarketingCampaignCashDetailQueryRsp struct {
	Response     *MarketingCampaignCashDetailQuery `json:"alipay_marketing_campaign_cash_detail_query_response"`
	AlipayCertSn string                            `json:"alipay_cert_sn,omitempty"`
	SignData     string                            `json:"-"`
	Sign         string                            `json:"sign"`
}

type MarketingActivityDeliveryStopRsp struct {
	Response     *MarketingActivityDeliveryStop `json:"alipay_marketing_activity_delivery_stop_response"`
	AlipayCertSn string                         `json:"alipay_cert_sn,omitempty"`
	SignData     string                         `json:"-"`
	Sign         string                         `json:"sign"`
}

type MarketingActivityDeliveryQueryRsp struct {
	Response     *MarketingActivityDeliveryQuery `json:"alipay_marketing_activity_delivery_query_response"`
	AlipayCertSn string                          `json:"alipay_cert_sn,omitempty"`
	SignData     string                          `json:"-"`
	Sign         string                          `json:"sign"`
}

type MarketingActivityDeliveryCreateRsp struct {
	Response     *MarketingActivityDeliveryCreate `json:"alipay_marketing_activity_delivery_create_response"`
	AlipayCertSn string                           `json:"alipay_cert_sn,omitempty"`
	SignData     string                           `json:"-"`
	Sign         string                           `json:"sign"`
}

type MarketingMaterialImageUploadRsp struct {
	Response     *MarketingMaterialImageUpload `json:"alipay_marketing_material_image_upload_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

// =========================================================分割=========================================================

type MarketingCampaignCashCreate struct {
	ErrorResponse
	CrowdNo       string `json:"crowd_no"`
	PayURL        string `json:"pay_url"`
	OriginCrowdNo string `json:"origin_crowd_no"`
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

type MarketingCampaignCashStatusModify struct {
	ErrorResponse
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

type MarketingActivityDeliveryStop struct {
	ErrorResponse
	DeliveryId string `json:"delivery_id"`
}

type MarketingActivityDeliveryQuery struct {
	ErrorResponse
	DeliveryId         string              `json:"delivery_id"`
	DeliveryStatus     string              `json:"delivery_status"`
	DeliveryBoothCode  string              `json:"delivery_booth_code"`
	DeliveryErrorMsg   string              `json:"delivery_error_msg"`
	DeliveryBaseInfo   *DeliveryBaseInfo   `json:"delivery_base_info"`
	DeliveryPlayConfig *DeliveryPlayConfig `json:"delivery_play_config"`
	DeliveryTargetRule *DeliveryTargetRule `json:"delivery_target_rule"`
}

type DeliveryBaseInfo struct {
	DeliveryName      string `json:"delivery_name"`
	DeliveryBeginTime string `json:"delivery_begin_time"`
	DeliveryEndTime   string `json:"delivery_end_time"`
}

type DeliveryPlayConfig struct {
	DeliverySingleSendConfig *DeliverySingleSendConfig `json:"delivery_single_send_config"`
	DeliveryFullSendConfig   *DeliveryFullSendConfig   `json:"delivery_full_send_config"`
}

type DeliverySingleSendConfig struct {
	DeliveryContentInfo *DeliveryContentInfo `json:"delivery_content_info"`
}

type DeliveryFullSendConfig struct {
	DeliveryFloorAmount string               `json:"delivery_floor_amount"`
	DeliveryContentInfo *DeliveryContentInfo `json:"delivery_content_info"`
}

type DeliveryContentInfo struct {
	DeliveryContentType     string                   `json:"delivery_content_type"`
	DeliveryActivityContent *DeliveryActivityContent `json:"delivery_activity_content"`
	DeliveryAppContent      *DeliveryAppContent      `json:"delivery_app_content"`
	DeliveryDisplayInfo     *DeliveryDisplayInfo     `json:"delivery_display_info"`
}

type DeliveryActivityContent struct {
	ActivityId string `json:"activity_id"`
}

type DeliveryAppContent struct {
	MiniAppId           string   `json:"mini_app_id"`
	ServiceCodeList     []string `json:"service_code_list"`
	MiniAppDeliveryType string   `json:"mini_app_delivery_type"`
}

type DeliveryDisplayInfo struct {
	MainTitle string `json:"main_title"`
	SubTitle  string `json:"sub_title"`
}

type DeliveryTargetRule struct {
	DeliveryMerchantRule *DeliveryMerchantRule `json:"delivery_merchant_rule"`
	DeliveryCityCodeRule *DeliveryCityCodeRule `json:"delivery_city_code_rule"`
	DeliveryPromoTags    string                `json:"delivery_promo_tags"`
}

type DeliveryMerchantRule struct {
	DeliveryMerchantMode  string                  `json:"delivery_merchant_mode"`
	DeliveryMerchantInfos []*DeliveryMerchantInfo `json:"delivery_merchant_infos"`
}

type DeliveryMerchantInfo struct {
	MerchantIdType string `json:"merchant_id_type"`
	MerchantId     string `json:"merchant_id"`
}

type DeliveryCityCodeRule struct {
	CityCodes []string `json:"city_codes"`
	AllCity   bool     `json:"all_city"`
}

type MarketingActivityDeliveryCreate struct {
	ErrorResponse
	DeliveryId string `json:"delivery_id"`
}

type MarketingMaterialImageUpload struct {
	ErrorResponse
	ResourceId      string `json:"resource_id"`
	ResourceEnhance bool   `json:"resource_enhance"`
	ResourceUrl     string `json:"resource_url"`
}
