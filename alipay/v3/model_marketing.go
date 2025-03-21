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
