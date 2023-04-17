package alipay

type AntMerchantShopModifyRsp struct {
	Response     *AntMerchantShopModify `json:"ant_merchant_expand_shop_modify_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type AntMerchantShopCreateRsp struct {
	Response     *AntMerchantShopCreate `json:"ant_merchant_expand_shop_create_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type AntMerchantShopConsultRsp struct {
	Response     *AntMerchantShopConsult `json:"ant_merchant_expand_shop_consult_response"`
	AlipayCertSn string                  `json:"alipay_cert_sn,omitempty"`
	SignData     string                  `json:"-"`
	Sign         string                  `json:"sign"`
}

type AntMerchantOrderQueryRsp struct {
	Response     *AntMerchantOrderQuery `json:"ant_merchant_expand_order_query_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type AntMerchantShopQueryRsp struct {
	Response     *AntMerchantShopQuery `json:"ant_merchant_expand_shop_query_response"`
	AlipayCertSn string                `json:"alipay_cert_sn,omitempty"`
	SignData     string                `json:"-"`
	Sign         string                `json:"sign"`
}

type AntMerchantShopCloseRsp struct {
	Response     *AntMerchantShopClose `json:"ant_merchant_expand_shop_close_response"`
	AlipayCertSn string                `json:"alipay_cert_sn,omitempty"`
	SignData     string                `json:"-"`
	Sign         string                `json:"sign"`
}

// =========================================================分割=========================================================

type AntMerchantShopModify struct {
	ErrorResponse
	OrderId string `json:"order_id"`
}

type AntMerchantShopCreate struct {
	ErrorResponse
	OrderId string `json:"order_id"`
}

type AntMerchantShopConsult struct {
	ErrorResponse
	AccountAudit bool   `json:"account_audit"`
	RiskAudit    bool   `json:"risk_audit"`
	OrderId      string `json:"order_id"`
}

type AntMerchantOrderQuery struct {
	ErrorResponse
	IPRoleID     []string `json:"ip_role_id,omitempty"`
	MerchantName string   `json:"merchant_name"`
	Status       string   `json:"status"`
	ApplyTime    string   `json:"apply_time"`
	ExtInfo      string   `json:"ext_info"`
}

type AntMerchantShopQuery struct {
	ErrorResponse
	ShopID          string `json:"shop_id"`
	BusinessAddress struct {
		CityCode     string `json:"city_code"`
		DistrictCode string `json:"district_code"`
		Address      string `json:"address"`
		ProvinceCode string `json:"province_code"`
		Poiid        string `json:"poiid,omitempty"`
		Longitude    string `json:"longitude,omitempty"`
		Latitude     string `json:"latitude,omitempty"`
		Type         string `json:"type,omitempty"`
	} `json:"business_address"`
	ShopCategory   string   `json:"shop_category"`
	StoreID        string   `json:"store_id"`
	ShopType       string   `json:"shop_type"`
	IPRoleID       string   `json:"ip_role_id"`
	ShopName       string   `json:"shop_name"`
	ContactPhone   string   `json:"contact_phone"`
	ContactMobile  string   `json:"contact_mobile"`
	CertNo         string   `json:"cert_no"`
	OutDoorImages  []string `json:"out_door_images,omitempty"`
	Qualifications []struct {
		IndustryQualificationType  string `json:"industry_qualification_type,omitempty"`
		IndustryQualificationImage string `json:"industry_qualification_image,omitempty"`
	} `json:"qualifications,omitempty"`
	CertType               string `json:"cert_type,omitempty"`
	CertName               string `json:"cert_name,omitempty"`
	CertImage              string `json:"cert_image,omitempty"`
	LegalName              string `json:"legal_name,omitempty"`
	LegalCertNo            string `json:"legal_cert_no,omitempty"`
	LicenseAuthLetterImage string `json:"license_auth_letter_image,omitempty"`
	SettleAlipayLogonID    string `json:"settle_alipay_logon_id,omitempty"`
	ExtInfos               []struct {
		KeyName string `json:"key_name"`
		Value   string `json:"value"`
	} `json:"ext_infos,omitempty"`
	BusinessTime []struct {
		WeekDay   int    `json:"week_day"`
		OpenTime  string `json:"open_time"`
		CloseTime string `json:"close_time"`
	} `json:"business_time,omitempty"`
	ContactInfos []struct {
		Name     string   `json:"name"`
		Phone    string   `json:"phone,omitempty"`
		Mobile   string   `json:"mobile,omitempty"`
		Email    string   `json:"email,omitempty"`
		Tag      []string `json:"tag"`
		Type     string   `json:"type"`
		IDCardNo string   `json:"id_card_no,omitempty"`
	} `json:"contact_infos,omitempty"`
	Memo            string `json:"memo,omitempty"`
	BrandID         string `json:"brand_id,omitempty"`
	Scene           string `json:"scene,omitempty"`
	NewShopCategory string `json:"new_shop_category"`
}

type AntMerchantShopClose struct {
	ErrorResponse
}
