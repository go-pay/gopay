package alipay

type AntMerchantShopCreateRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	OrderId     string `json:"order_id"`
	ExistShopId string `json:"exist_shop_id"`
}

type AntMerchantShopQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	BusinessAddress struct {
		CityCode     string `json:"city_code"`
		DistrictCode string `json:"district_code"`
		Address      string `json:"address"`
		ProvinceCode string `json:"province_code"`
		Poiid        string `json:"poiid"`
		Longitude    string `json:"longitude"`
		Latitude     string `json:"latitude"`
		Type         string `json:"type"`
	} `json:"business_address"`
	ShopCategory           string   `json:"shop_category"`
	StoreId                string   `json:"store_id"`
	ShopType               string   `json:"shop_type"`
	IpRoleId               string   `json:"ip_role_id"`
	ShopName               string   `json:"shop_name"`
	ShopId                 string   `json:"shop_id"`
	ContactPhone           string   `json:"contact_phone"`
	ContactMobile          string   `json:"contact_mobile"`
	CertNo                 string   `json:"cert_no"`
	OutDoorImages          []string `json:"out_door_images"`
	CertType               string   `json:"cert_type"`
	CertName               string   `json:"cert_name"`
	CertImage              string   `json:"cert_image"`
	LegalName              string   `json:"legal_name"`
	LegalCertNo            string   `json:"legal_cert_no"`
	LicenseAuthLetterImage string   `json:"license_auth_letter_image"`
	SettleAlipayLogonId    string   `json:"settle_alipay_logon_id"`
	ExtInfos               []struct {
		KeyName string `json:"key_name"`
		Value   string `json:"value"`
	} `json:"ext_infos"`
	BusinessTime []struct {
		WeekDay   int    `json:"week_day"`
		OpenTime  string `json:"open_time"`
		CloseTime string `json:"close_time"`
	} `json:"business_time"`
	ContactInfos []struct {
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Mobile   string `json:"mobile"`
		Email    string `json:"email"`
		IdCardNo string `json:"id_card_no"`
	} `json:"contact_infos"`
	Memo              string `json:"memo"`
	BrandId           string `json:"brand_id"`
	Scene             string `json:"scene"`
	AlipayPoiid       string `json:"alipay_poiid"`
	NewShopCategory   string `json:"new_shop_category"`
	ShopInfoStatus    string `json:"shop_info_status"`
	ShopRecommendInfo struct {
		UnconfidenceReason  string `json:"unconfidence_reason"`
		Recommend           string `json:"recommend"`
		RecommendName       string `json:"recommend_name"`
		RecommendLongtitude string `json:"recommend_longtitude"`
		RecommendLatitude   string `json:"recommend_latitude"`
		RecommendAddress    string `json:"recommend_address"`
	} `json:"shop_recommend_info"`
}

type AntMerchantShopModifyRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	OrderId string `json:"order_id"`
}

type AntMerchantShopCloseRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

type AntMerchantOrderQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	IpRoleId     []string `json:"ip_role_id"`
	MerchantName string   `json:"merchant_name"`
	Status       string   `json:"status"`
	ApplyTime    string   `json:"apply_time"`
	ExtInfo      string   `json:"ext_info"`
}

type AntMerchantShopPageQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ShopInfos []struct {
		ShopId          string `json:"shop_id"`
		BusinessAddress struct {
			CityCode     string `json:"city_code"`
			DistrictCode string `json:"district_code"`
			Address      string `json:"address"`
			ProvinceCode string `json:"province_code"`
			Poiid        string `json:"poiid"`
			Longitude    string `json:"longitude"`
			Latitude     string `json:"latitude"`
			Type         string `json:"type"`
		} `json:"business_address"`
		ShopCategory    string `json:"shop_category"`
		NewShopCategory string `json:"new_shop_category"`
		StoreId         string `json:"store_id"`
		ShopType        string `json:"shop_type"`
		ShopName        string `json:"shop_name"`
		ContactPhone    string `json:"contact_phone"`
		ContactMobile   string `json:"contact_mobile"`
		BusinessTime    []struct {
			WeekDay   int    `json:"week_day"`
			OpenTime  string `json:"open_time"`
			CloseTime string `json:"close_time"`
		} `json:"business_time"`
		ShopStatus     string `json:"shop_status"`
		ShopInfoStatus string `json:"shop_info_status"`
	} `json:"shop_infos"`
	TotalPages  int    `json:"total_pages"`
	AlipayPoiid string `json:"alipay_poiid"`
}

type AntMerchantExpandIndirectImageUploadRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ImageId string `json:"image_id"`
}

type AntMerchantExpandMccQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	MccInfoList []struct {
		MccLevel1           string `json:"mcc_level_1"`
		MccLevel1Name       string `json:"mcc_level_1_name"`
		MccLevel2           string `json:"mcc_level_2"`
		MccLevel2Name       string `json:"mcc_level_2_name"`
		IsSpecial           bool   `json:"is_special"`
		SpecialQualRequired bool   `json:"special_qual_required"`
		MccRequirements     string `json:"mcc_requirements"`
	} `json:"mcc_info_list"`
}

type AntMerchantExpandShopReceiptAccountSaveRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}
