package alipay

type MarketingCardTemplateCreateRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	TemplateId string `json:"template_id"`
}

type MarketingCardTemplateQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	AccessVersion     string `json:"access_version"`
	TemplateStyleInfo struct {
		CardShowName string `json:"card_show_name"`
		LogoId       string `json:"logo_id"`
		BackgroundId string `json:"background_id"`
		BrandName    string `json:"brand_name"`
	} `json:"template_style_info"`
	CardLevelConfs []struct {
		Level         string `json:"level"`
		LevelShowName string `json:"level_show_name"`
		LevelIcon     string `json:"level_icon"`
		LevelDesc     string `json:"level_desc"`
	} `json:"card_level_confs"`
	TemplateFormConfig struct {
		Fields struct {
			Required []string `json:"required"`
			Optional []string `json:"optional"`
		} `json:"fields"`
		OpenCardMiniAppId string `json:"open_card_mini_app_id"`
	} `json:"template_form_config"`
	SpiAppId string `json:"spi_app_id"`
}

type MarketingCardTemplateModifyRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	TemplateId string `json:"template_id"`
}

type MarketingCardFormTemplateSetRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

type MarketingCardQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CardInfo struct {
		BizCardNo      string `json:"biz_card_no"`
		ExternalCardNo string `json:"external_card_no"`
		OpenDate       string `json:"open_date"`
		ValidDate      string `json:"valid_date"`
		Level          string `json:"level"`
		Point          string `json:"point"`
		Balance        string `json:"balance"`
		TemplateId     string `json:"template_id"`
		MdcodeInfo     struct {
			CodeStatus string `json:"code_status"`
			CodeValue  string `json:"code_value"`
			ExpireTime string `json:"expire_time"`
			TimeStamp  int    `json:"time_stamp"`
		} `json:"mdcode_info"`
		FrontTextList []struct {
			Label string `json:"label"`
			Value string `json:"value"`
		} `json:"front_text_list"`
		FrontImageId string `json:"front_image_id"`
	} `json:"card_info"`
	SchemaUrl         string `json:"schema_url"`
	PassId            string `json:"pass_id"`
	PaidOuterCardInfo struct {
		Action       string `json:"action"`
		PurchaseInfo struct {
			Source        string `json:"source"`
			Price         string `json:"price"`
			ActionDate    string `json:"action_date"`
			OutTradeNo    string `json:"out_trade_no"`
			AlipayTradeNo string `json:"alipay_trade_no"`
		} `json:"purchase_info"`
		CycleInfo struct {
			OpenStatus              string `json:"open_status"`
			CloseReason             string `json:"close_reason"`
			CycleType               string `json:"cycle_type"`
			AlipayDeductScene       string `json:"alipay_deduct_scene"`
			AlipayDeductProductCode string `json:"alipay_deduct_product_code"`
			AlipayDeductAgreement   string `json:"alipay_deduct_agreement"`
		} `json:"cycle_info"`
	} `json:"paid_outer_card_info"`
}

type MarketingCardUpdateRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ResultCode string `json:"result_code"`
}

type MarketingCardDeleteRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	BizSerialNo string `json:"biz_serial_no"`
}

type MarketingCardMessageNotifyRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ResultCode string `json:"result_code"`
}

type OfflineMaterialImageUploadRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ImageId  string `json:"image_id"`
	ImageUrl string `json:"image_url"`
}
