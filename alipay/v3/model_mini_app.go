package alipay

type OpenMiniVersionAuditedCancelRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

type OpenMiniVersionGrayOnlineRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

type OpenMiniVersionGrayCancelRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

type OpenMiniVersionOnlineRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

type OpenMiniVersionOfflineRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

type OpenMiniVersionRollbackRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

type OpenMiniVersionDeleteRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

type OpenMiniVersionAuditApplyRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	SpeedUp     string `json:"speed_up"`
	SpeedUpMemo string `json:"speed_up_memo"`
}

type OpenMiniVersionUploadRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	BuildStatus    string `json:"build_status"`
	NeedRotation   string `json:"need_rotation"`
	CreateStatus   string `json:"create_status"`
	VersionCreated string `json:"version_created"`
}

type OpenMiniTemplateUsageQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	TemplateUsageInfoList []struct {
		MiniAppId  string `json:"mini_app_id"`
		AppVersion string `json:"app_version"`
	} `json:"template_usage_info_list"`
}

type OpenMiniVersionBuildQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	NeedRotation string `json:"need_rotation"`
	CreateStatus string `json:"create_status"`
}

type OpenMiniVersionDetailQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	AppVersion        string `json:"app_version"`
	AppName           string `json:"app_name"`
	AppEnglishName    string `json:"app_english_name"`
	AppLogo           string `json:"app_logo"`
	VersionDesc       string `json:"version_desc"`
	GrayStrategy      string `json:"gray_strategy"`
	Status            string `json:"status"`
	RejectReason      string `json:"reject_reason"`
	ScanResult        string `json:"scan_result"`
	GmtCreate         string `json:"gmt_create"`
	GmtApplyAudit     string `json:"gmt_apply_audit"`
	GmtOnline         string `json:"gmt_online"`
	GmtOffline        string `json:"gmt_offline"`
	AppDesc           string `json:"app_desc"`
	GmtAuditEnd       string `json:"gmt_audit_end"`
	ServiceRegionType string `json:"service_region_type"`
	ServiceRegionInfo []struct {
		ProvinceCode string `json:"province_code"`
		ProvinceName string `json:"province_name"`
		CityCode     string `json:"city_code"`
		CityName     string `json:"city_name"`
		AreaCode     string `json:"area_code"`
		AreaName     string `json:"area_name"`
	} `json:"service_region_info"`
	ScreenShotList          []string `json:"screen_shot_list"`
	AppSlogan               string   `json:"app_slogan"`
	Memo                    string   `json:"memo"`
	ServicePhone            string   `json:"service_phone"`
	ServiceEmail            string   `json:"service_email"`
	MiniAppCategoryInfoList []struct {
		FirstCategoryId    string `json:"first_category_id"`
		FirstCategoryName  string `json:"first_category_name"`
		SecondCategoryId   string `json:"second_category_id"`
		SecondCategoryName string `json:"second_category_name"`
		ThirdCategoryId    string `json:"third_category_id"`
		ThirdCategoryName  string `json:"third_category_name"`
	} `json:"mini_app_category_info_list"`
	PackageInfoList []struct {
		PackageName     string `json:"package_name"`
		PackageDesc     string `json:"package_desc"`
		DocUrl          string `json:"doc_url"`
		Status          string `json:"status"`
		PackageOpenType string `json:"package_open_type"`
	} `json:"package_info_list"`
	MiniCategoryInfoList []struct {
		FirstCategoryId    string `json:"first_category_id"`
		FirstCategoryName  string `json:"first_category_name"`
		SecondCategoryId   string `json:"second_category_id"`
		SecondCategoryName string `json:"second_category_name"`
		ThirdCategoryId    string `json:"third_category_id"`
		ThirdCategoryName  string `json:"third_category_name"`
	} `json:"mini_category_info_list"`
	BaseAudit       string `json:"base_audit"`
	PromoteAudit    string `json:"promote_audit"`
	CanRelease      string `json:"can_release"`
	BaseAuditRecord struct {
		AuditImages []string `json:"audit_images"`
		Memos       []struct {
			Memo          string   `json:"memo"`
			MemoImageList []string `json:"memo_image_list"`
		} `json:"memos"`
	} `json:"base_audit_record"`
	PromoteAuditRecord struct {
		AuditImages []string `json:"audit_images"`
		Memos       []struct {
			Memo          string   `json:"memo"`
			MemoImageList []string `json:"memo_image_list"`
		} `json:"memos"`
	} `json:"promote_audit_record"`
}

type OpenMiniVersionListQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	AppVersionInfos []struct {
		BundleId           string `json:"bundle_id"`
		AppVersion         string `json:"app_version"`
		VersionDescription string `json:"version_description"`
		VersionStatus      string `json:"version_status"`
		CreateTime         string `json:"create_time"`
		BaseAudit          string `json:"base_audit"`
		PromoteAudit       string `json:"promote_audit"`
		CanRelease         string `json:"can_release"`
	} `json:"app_version_infos"`
	AppVersions []string `json:"app_versions"`
}

type OpenMiniExperienceCreateRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

type OpenMiniExperienceQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	Status       string `json:"status"`
	ExpQrCodeUrl string `json:"exp_qr_code_url"`
	ExpSchemaUrl string `json:"exp_schema_url"`
}

type OpenMiniExperienceCancelRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}
