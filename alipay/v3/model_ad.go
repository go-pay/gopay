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

	ConversionId string              `json:"conversion_id"`
	PageNo       int                 `json:"page_no"`
	PageSize     int                 `json:"page_size"`
	Total        int                 `json:"total"`
	List         []AdPromotepageData `json:"list"`
}

type AdPromotepageData struct {
	BizNo        string           `json:"biz_no"`
	PropertyList []AdPropertyData `json:"property_list"`
	EncryptUid   string           `json:"encrypt_uid"`
}

type AdPropertyData struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

// 任务广告完成状态查询响应
type XlightTaskQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	TaskStatus       string `json:"task_status"`
	TaskResult       string `json:"task_result"`
	TaskRewardAmount string `json:"task_reward_amount"`
}

// 消费明细查询响应
type AdConsumehistoryQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	DataList []AdConsumehistoryData `json:"data_list"`
	Total    int                    `json:"total"`
}

type AdConsumehistoryData struct {
	BizDate                     string `json:"biz_date"`
	PrincipalName               string `json:"principal_name"`
	AlipayAccount               string `json:"alipay_account"`
	SellModeName                string `json:"sell_mode_name"`
	SceneTypeName               string `json:"scene_type_name"`
	CutAmountFormat             string `json:"cut_amount_format"`
	CashAmountFormat            string `json:"cash_amount_format"`
	CreditAmountFormat          string `json:"credit_amount_format"`
	RedPacketAmountFormat       string `json:"red_packet_amount_format"`
	RebateRedPacketAmountFormat string `json:"rebate_red_packet_amount_format"`
}

// 商品落地页信息创建或更新响应
type ProductLandinginfoCreateOrModifyRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ItemId    string `json:"item_id"`
	OutItemId string `json:"out_item_id"`
}

// 商品落地页信息查询响应
type ProductLandinginfoQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ItemId    string `json:"item_id"`
	OutItemId string `json:"out_item_id"`
	Landing   struct {
		LandingId   string   `json:"landing_id"`
		LandingName string   `json:"landing_name"`
		LandingUrl  string   `json:"landing_url"`
		LandingType string   `json:"landing_type"`
		PicInfoList []string `json:"pic_info_list"`
		LandingAct  []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"landing_act"`
		ProductVideos []struct {
			Url       string `json:"url"`
			OriginUrl string `json:"origin_url"`
			OssUrl    string `json:"oss_url"`
			PosterUrl string `json:"poster_url"`
			Width     int    `json:"width"`
			Height    int    `json:"height"`
			Size      int    `json:"size"`
			Duration  int    `json:"duration"`
			Signature string `json:"signature"`
		} `json:"product_videos"`
	} `json:"landing"`
}

type AdAgentreportdataQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	DataList []struct {
		DataId                 string `json:"data_id"`
		BizDate                string `json:"biz_date"`
		Impression             int    `json:"impression"`
		Click                  int    `json:"click"`
		Cost                   int    `json:"cost"`
		PlanName               string `json:"plan_name"`
		PlanId                 string `json:"plan_id"`
		GroupName              string `json:"group_name"`
		GroupId                string `json:"group_id"`
		OrderName              string `json:"order_name"`
		OrderId                string `json:"order_id"`
		CreativeName           string `json:"creative_name"`
		MarketTargetName       string `json:"market_target_name"`
		SceneName              string `json:"scene_name"`
		PrincipalAlipayAccount string `json:"principal_alipay_account"`
		PrincipalName          string `json:"principal_name"`
		PrincipalPid           string `json:"principal_pid"`
		CostFormat             string `json:"cost_format"`
		Cpm                    string `json:"cpm"`
		ClickRate              string `json:"click_rate"`
		Cpc                    string `json:"cpc"`
		ConvResult             string `json:"conv_result"`
		Cvr                    string `json:"cvr"`
		AvgConvCost            string `json:"avg_conv_cost"`
		AgentName              string `json:"agent_name"`
		AgentAlipayAccount     string `json:"agent_alipay_account"`
		ConversionDataList     []struct {
			ConversionType     string `json:"conversion_type"`
			ConversionTypeName string `json:"conversion_type_name"`
			ConversionResult   string `json:"conversion_result"`
		} `json:"conversion_data_list"`
	} `json:"data_list"`
	Total int `json:"total"`
}
