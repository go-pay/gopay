package alipay

// =========================================================分割=========================================================

// 转化数据回传 Response
type DataServiceAdConversionUploadResponse struct {
	Response     *DataServiceAdConversionUpload `json:"alipay_data_dataservice_ad_conversion_upload_response"`
	AlipayCertSn string                         `json:"alipay_cert_sn,omitempty"`
	SignData     string                         `json:"-"`
	Sign         string                         `json:"sign"`
}

// 广告投放数据通用查询 Response
type DataServiceAdReportdataQueryResponse struct {
	Response     *DataServiceAdReportdataQuery `json:"alipay_data_dataservice_ad_reportdata_query_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

// 自建推广页列表批量查询 Response
type DataServiceAdPromotepageBatchqueryResponse struct {
	Response     *DataServiceAdPromotepageBatchquery `json:"alipay_data_dataservice_ad_promotepage_batchquery_response"`
	AlipayCertSn string                              `json:"alipay_cert_sn,omitempty"`
	SignData     string                              `json:"-"`
	Sign         string                              `json:"sign"`
}

// 自建推广页留资数据查询 Response
type DataServiceAdPromotepageDownloadResponse struct {
	Response     *DataServiceAdPromotepageDownload `json:"alipay_data_dataservice_ad_promotepage_download_response"`
	AlipayCertSn string                            `json:"alipay_cert_sn,omitempty"`
	SignData     string                            `json:"-"`
	Sign         string                            `json:"sign"`
}

// 任务广告完成状态查询 Response
type DataServiceXlightTaskQueryResponse struct {
	Response     *DataServiceXlightTaskQuery `json:"alipay_data_dataservice_xlight_task_query_response"`
	AlipayCertSn string                      `json:"alipay_cert_sn,omitempty"`
	SignData     string                      `json:"-"`
	Sign         string                      `json:"sign"`
}

// =========================================================分割=========================================================

type DataServiceAdConversionUpload struct {
	ErrorResponse
}

type DataServiceAdReportdataQuery struct {
	ErrorResponse
	DataList []*AdReportDataDetail `json:"data_list,omitempty"` // 汇总结果数据
	Total    int                   `json:"total,omitempty"`     // 总条数
}

type DataServiceAdPromotepageBatchquery struct {
	ErrorResponse
	PageNo   int                  `json:"page_no,omitempty"`   // 页码
	PageSize int                  `json:"page_size,omitempty"` // 每页大小
	Total    int                  `json:"total,omitempty"`     // 总数据量
	List     []*PromotePageDetail `json:"list,omitempty"`      // 推广页详情列表
}

type DataServiceAdPromotepageDownload struct {
	ErrorResponse
	PageNo       int                `json:"page_no,omitempty"`       // 页码
	PageSize     int                `json:"page_size,omitempty"`     // 每页大小
	Total        int                `json:"total,omitempty"`         // 总数据量
	ConversionId string             `json:"conversion_id,omitempty"` // 转化事件ID
	List         []*PromotePageData `json:"list,omitempty"`          // 留资数据列表
}

type DataServiceXlightTaskQuery struct {
	ErrorResponse
	Status           string `json:"status,omitempty"`             // 任务完成状态：FINISHED/UNFINISHED
	FinishTime       string `json:"finish_time,omitempty"`        // 任务完成时间
	TaskRewardAmount string `json:"task_reward_amount,omitempty"` // 建议给用户的最大发奖金额(分)
}

// =========================================================分割=========================================================

// AdReportDataDetail 广告投放数据明细
type AdReportDataDetail struct {
	DataId                 string              `json:"data_id,omitempty"`                  // 数据ID
	BizDate                string              `json:"biz_date,omitempty"`                 // 业务日期
	Impression             int                 `json:"impression,omitempty"`               // 曝光数
	Click                  int                 `json:"click,omitempty"`                    // 点击数
	Cost                   int                 `json:"cost,omitempty"`                     // 花费(分)
	PlanName               string              `json:"plan_name,omitempty"`                // 计划名称
	PlanId                 string              `json:"plan_id,omitempty"`                  // 计划ID
	GroupId                string              `json:"group_id,omitempty"`                 // 单元ID
	GroupName              string              `json:"group_name,omitempty"`               // 单元名称
	OrderId                string              `json:"order_id,omitempty"`                 // 订单ID
	OrderName              string              `json:"order_name,omitempty"`               // 订单名称
	CreativeName           string              `json:"creative_name,omitempty"`            // 创意名称
	CreativeId             string              `json:"creative_id,omitempty"`              // 创意ID
	MarketTargetName       string              `json:"market_target_name,omitempty"`       // 营销目标名称
	SceneName              string              `json:"scene_name,omitempty"`               // 场景名称
	PrincipalAlipayAccount string              `json:"principal_alipay_account,omitempty"` // 商家支付宝账号
	PrincipalName          string              `json:"principal_name,omitempty"`           // 商家名称
	PrincipalPid           string              `json:"principal_pid,omitempty"`            // 商家pid
	CostFormat             string              `json:"cost_format,omitempty"`              // 花费格式化(元)
	Cpm                    string              `json:"cpm,omitempty"`                      // 千次曝光均价
	ClickRate              string              `json:"click_rate,omitempty"`               // 点击率
	Cpc                    string              `json:"cpc,omitempty"`                      // 点击均价
	ConvResult             string              `json:"conv_result,omitempty"`              // 转化数
	Cvr                    string              `json:"cvr,omitempty"`                      // 转化率
	AvgConvCost            string              `json:"avg_conv_cost,omitempty"`            // 转化均价
	AgentName              string              `json:"agent_name,omitempty"`               // 代理商名称
	AgentAlipayAccount     string              `json:"agent_alipay_account,omitempty"`     // 代理商支付宝账号
	MainPictureId          string              `json:"main_picture_id,omitempty"`          // 主图ID
	MainPictureName        string              `json:"main_picture_name,omitempty"`        // 主图名称
	MainPictureUrl         string              `json:"main_picture_url,omitempty"`         // 主图URL
	MainPictureWidth       string              `json:"main_picture_width,omitempty"`       // 主图宽度
	MainPictureHeight      string              `json:"main_picture_height,omitempty"`      // 主图高度
	MainVideoId            string              `json:"main_video_id,omitempty"`            // 主视频ID
	MainVideoName          string              `json:"main_video_name,omitempty"`          // 主视频名称
	MainVideoUrl           string              `json:"main_video_url,omitempty"`           // 主视频URL
	MainVideoWidth         int                 `json:"main_video_width,omitempty"`         // 主视频宽度
	MainVideoHeight        int                 `json:"main_video_height,omitempty"`        // 主视频高度
	MainTitleId            string              `json:"main_title_id,omitempty"`            // 主标题ID
	MainTitle              string              `json:"main_title,omitempty"`               // 主标题
	SubTitleId             string              `json:"sub_title_id,omitempty"`             // 子标题ID
	SubTitle               string              `json:"sub_title,omitempty"`                // 子标题
	ActionPointId          string              `json:"action_point_id,omitempty"`          // 行动点ID
	ActionPoint            string              `json:"action_point,omitempty"`             // 行动点
	ConversionDataList     []*AdConversionData `json:"conversion_data_list,omitempty"`     // 转化数据列表
}

// AdConversionData 广告转化数据
type AdConversionData struct {
	ConversionType     string `json:"conversion_type,omitempty"`      // 转化类型
	ConversionTypeName string `json:"conversion_type_name,omitempty"` // 转化类型名称
	ConversionResult   string `json:"conversion_result,omitempty"`    // 转化数
}

// PromotePageDetail 推广页详情
type PromotePageDetail struct {
	Id           int                    `json:"id,omitempty"`            // 推广页ID
	Type         string                 `json:"type,omitempty"`          // 推广页类型
	Name         string                 `json:"name,omitempty"`          // 推广页名称
	GmtCreate    string                 `json:"gmt_create,omitempty"`    // 创建时间
	PropertyList []*PromotePageProperty `json:"property_list,omitempty"` // 属性列表
}

// PromotePageProperty 推广页属性
type PromotePageProperty struct {
	Key   string `json:"key,omitempty"`   // 属性key
	Name  string `json:"name,omitempty"`  // 属性名称
	Type  string `json:"type,omitempty"`  // 属性类型
	Value string `json:"value,omitempty"` // 属性值
}

// PromotePageData 推广页留资数据
type PromotePageData struct {
	BizNo        string                 `json:"biz_no,omitempty"`        // 业务号
	PropertyList []*PromotePageProperty `json:"property_list,omitempty"` // 属性列表
	EncryptUid   string                 `json:"encrypt_uid,omitempty"`   // 加密uid
}
