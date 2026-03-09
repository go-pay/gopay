package alipay

// 上传创建人群响应
type MerchantQipanCrowdCreateRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CrowdCode string `json:"crowd_code"`
}

// 人群中追加用户响应
type MerchantQipanCrowdUserAddRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

// 人群中删除用户响应
type MerchantQipanCrowdUserDeleteRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

// 棋盘人群圈选标签基本信息查询响应
type MarketingQipanTagBaseBatchQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	OperationTagList []struct {
		TagCode string `json:"tag_code"`
		TagName string `json:"tag_name"`
		TagDesc string `json:"tag_desc"`
	} `json:"operation_tag_list"`
}

// 棋盘标签圈选值查询响应
type MarketingQipanTagQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	OperationTag *OpenCrowdOperationTag `json:"operation_tag"`
}

// OpenCrowdOperationTag 圈选标签详情
type OpenCrowdOperationTag struct {
	TagCode             string                      `json:"tag_code"`
	TagName             string                      `json:"tag_name"`
	TagDesc             string                      `json:"tag_desc"`
	OperationOptionList []*OpenCrowdOperationOption `json:"operation_option_list"`
}

// OpenCrowdOperationOption 标签圈选项
type OpenCrowdOperationOption struct {
	OptionCode     string                 `json:"option_code"`
	OptionName     string                 `json:"option_name"`
	OptionDataType string                 `json:"option_data_type"`
	OptionList     []*OpenOperationOption `json:"option_list"`
}

// OpenOperationOption 标签圈选值
type OpenOperationOption struct {
	Id          int    `json:"id"`
	Value       string `json:"value"`
	Text        string `json:"text"`
	Desc        string `json:"desc"`
	ParentId    string `json:"parent_id"`
	ParentValue string `json:"parent_value"`
}

// 查询人群列表响应
type MarketingQipanCrowdBatchQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	TotalNumber string                `json:"total_number"`
	CrowdList   []*QipanMerchantCrowd `json:"crowd_list"`
}

// QipanMerchantCrowd 商家自定义人群
type QipanMerchantCrowd struct {
	CrowdCode         string   `json:"crowd_code"`
	CrowdName         string   `json:"crowd_name"`
	ExternalCrowdCode string   `json:"external_crowd_code"`
	Status            string   `json:"status"`
	CrowdDesc         string   `json:"crowd_desc"`
	Processable       bool     `json:"processable"`
	ApplyChannelList  []string `json:"apply_channel_list"`
	CrowdSize         string   `json:"crowd_size"`
	Hidden            bool     `json:"hidden"`
}

// 查询人群详情响应
type MarketingQipanCrowdQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CrowdInfo *QipanMerchantCrowd `json:"crowd_info"`
}

// 修改人群响应
type MarketingQipanCrowdModifyRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

// 看板分析响应
type MarketingQipanBoardQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	BoardData interface{} `json:"board_data"`
}

// 画像分析响应
type MarketingQipanInsightQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	InsightData interface{} `json:"insight_data"`
}

// 行为分析响应
type MarketingQipanBehaviorQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	BehaviorData interface{} `json:"behavior_data"`
}

// 趋势分析响应
type MarketingQipanTrendQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	TrendData interface{} `json:"trend_data"`
}

// 常住省市查询响应
type MarketingQipanInsightCityQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CityData interface{} `json:"city_data"`
}

// 人群池创建响应
type MerchantQipanCrowdPoolCreateRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CrowdCode string `json:"crowd_code"`
}

// 人群扩展接口响应
type MerchantQipanCrowdSpreadRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CrowdCode string `json:"crowd_code"`
}

// 上传创建灰黑产人群响应
type MerchantQipanGreyBlackCrowdCreateRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	CrowdCode string `json:"crowd_code"`
}

// 灰黑产人群中追加用户响应
type MerchantQipanGreyBlackCrowdUserAddRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

// 灰黑产人群中删除用户响应
type MerchantQipanGreyBlackCrowdUserDeleteRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}
