package alipay

type AdConversionUploadRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}
type AdReportdataQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	DataList []AdReportData `json:"data_list"`
	Total    int            `json:"total"`
}

type AdReportData struct {
	DataId                 string             `json:"data_id"`
	BizDate                string             `json:"biz_date"`
	Impression             int                `json:"impression"`
	Click                  int                `json:"click"`
	Cost                   int                `json:"cost"`
	PlanName               string             `json:"plan_name"`
	PlanId                 string             `json:"plan_id"`
	GroupName              string             `json:"group_name"`
	GroupId                string             `json:"group_id"`
	OrderName              string             `json:"order_name"`
	OrderId                string             `json:"order_id"`
	CreativeName           string             `json:"creative_name"`
	MarketTargetName       string             `json:"market_target_name"`
	SceneName              string             `json:"scene_name"`
	PrincipalAlipayAccount string             `json:"principal_alipay_account"`
	PrincipalName          string             `json:"principal_name"`
	ConversionDataList     []AdConversionData `json:"conversion_data_list"`
}

type AdConversionData struct {
	ConversionType     string `json:"conversion_type"`
	ConversionTypeName string `json:"conversion_type_name"`
	ConversionResult   string `json:"conversion_result"`
}

type AdPromotepageBatchqueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	PageNo   int             `json:"page_no"`
	PageSize int             `json:"page_size"`
	Total    int             `json:"total"`
	List     []AdPromotepage `json:"list"`
}

type AdPromotepage struct {
	Id           int          `json:"id"`
	Type         string       `json:"type"`
	Name         string       `json:"name"`
	GmtCreate    string       `json:"gmt_create"`
	PropertyList []AdProperty `json:"property_list"`
}

type AdProperty struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// 自建推广页留资数据查询响应
type AdPromotepageDownloadRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	PageNo   int                 `json:"page_no"`
	PageSize int                 `json:"page_size"`
	Total    int                 `json:"total"`
	List     []AdPromotePageData `json:"list"`
}

// 推广页留资数据
type AdPromotePageData struct {
	PageId     string `json:"page_id"`
	Name       string `json:"name"`
	Mobile     string `json:"mobile"`
	GmtCreate  string `json:"gmt_create"`
	ExtendInfo string `json:"extend_info"`
}

// 任务广告完成状态查询响应
type XlightTaskQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	TaskStatus string `json:"task_status"`
	TaskResult string `json:"task_result"`
}

// 消费明细查询响应
type AdConsumehistoryQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	PageNo   int                `json:"page_no"`
	PageSize int                `json:"page_size"`
	Total    int                `json:"total"`
	List     []AdConsumeHistory `json:"list"`
}

// 消费明细
type AdConsumeHistory struct {
	BizDate     string `json:"biz_date"`
	ChargeType  string `json:"charge_type"`
	Amount      string `json:"amount"`
	OuterCode   string `json:"outer_code"`
	ChargeDesc  string `json:"charge_desc"`
	GmtCreate   string `json:"gmt_create"`
	GmtModified string `json:"gmt_modified"`
}

// 商品落地页信息创建或更新响应
type ProductLandinginfoCreateOrModifyRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	Success bool `json:"success"`
}

// 商品落地页信息查询响应
type ProductLandinginfoQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}
