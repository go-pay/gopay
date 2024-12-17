package alipay

type TradeRelationBindResponse struct {
	Response     *TradeRelationBind `json:"alipay_trade_royalty_relation_bind_response"`
	AlipayCertSn string             `json:"alipay_cert_sn,omitempty"`
	SignData     string             `json:"-"`
	Sign         string             `json:"sign"`
}

type TradeRelationUnbindResponse struct {
	Response     *TradeRelationBind `json:"alipay_trade_royalty_relation_unbind_response"`
	AlipayCertSn string             `json:"alipay_cert_sn,omitempty"`
	SignData     string             `json:"-"`
	Sign         string             `json:"sign"`
}

type TradeRelationBatchQueryResponse struct {
	Response     *TradeRelationBatchQuery `json:"alipay_trade_order_settle_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type TradeSettleConfirmResponse struct {
	Response     *TradeOrderSettle `json:"alipay_trade_settle_confirm_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type TradeOrderSettleResponse struct {
	Response     *TradeOrderSettle `json:"alipay_trade_order_settle_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type TradeOrderSettleQueryResponse struct {
	Response     *TradeOrderSettleQuery `json:"alipay_trade_order_settle_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type TradeOrderOnSettleQueryResponse struct {
	Response     *TradeOrderOnSettleQuery `json:"alipay_trade_order_onsettle_query_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type TradeRoyaltyRateQueryResponse struct {
	Response     *TradeRoyaltyRateQuery `json:"alipay_trade_royalty_rate_query_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type SecurityCustomerRiskSendRsp struct {
	Response     *SecurityCustomerRiskSend `json:"alipay_security_risk_customerrisk_send_response"`
	AlipayCertSn string                    `json:"alipay_cert_sn,omitempty"`
	SignData     string                    `json:"-"`
	Sign         string                    `json:"sign"`
}

type PayAppMarketingConsultRsp struct {
	Response     *PayAppMarketingConsult `json:"alipay_pay_app_marketing_consult_response"`
	AlipayCertSn string                  `json:"alipay_cert_sn,omitempty"`
	SignData     string                  `json:"-"`
	Sign         string                  `json:"sign"`
}

type MerchantQipanCrowdCreateRsp struct {
	Response     *MerchantQipanCrowdCreate `json:"alipay_merchant_qipan_crowd_create_response"`
	AlipayCertSn string                    `json:"alipay_cert_sn,omitempty"`
	SignData     string                    `json:"-"`
	Sign         string                    `json:"sign"`
}

type MerchantQipanCrowdUserAddRsp struct {
	Response     *MerchantQipanCrowdUserAdd `json:"alipay_merchant_qipan_crowduser_add_response"`
	AlipayCertSn string                     `json:"alipay_cert_sn,omitempty"`
	SignData     string                     `json:"-"`
	Sign         string                     `json:"sign"`
}

type MerchantQipanCrowdUserDeleteRsp struct {
	Response     *MerchantQipanCrowdUserDelete `json:"alipay_merchant_qipan_crowduser_delete_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type MarketingQipanTagBaseBatchQueryRsp struct {
	Response     *MarketingQipanTagBaseBatchQuery `json:"alipay_marketing_qipan_tagbase_batchquery_response"`
	AlipayCertSn string                           `json:"alipay_cert_sn,omitempty"`
	SignData     string                           `json:"-"`
	Sign         string                           `json:"sign"`
}

type MarketingQipanTagQueryRsp struct {
	Response     *MarketingQipanTagQuery `json:"alipay_marketing_qipan_tag_query_response"`
	AlipayCertSn string                  `json:"alipay_cert_sn,omitempty"`
	SignData     string                  `json:"-"`
	Sign         string                  `json:"sign"`
}

type MarketingQipanCrowdOperationCreateRsp struct {
	Response     *MarketingQipanCrowdOperationCreate `json:"alipay_marketing_qipan_crowdoperation_create_response"`
	AlipayCertSn string                              `json:"alipay_cert_sn,omitempty"`
	SignData     string                              `json:"-"`
	Sign         string                              `json:"sign"`
}

type MarketingQipanCrowdTagQueryRsp struct {
	Response     *MarketingQipanCrowdTagQuery `json:"alipay_marketing_qipan_crowdtag_query_response"`
	AlipayCertSn string                       `json:"alipay_cert_sn,omitempty"`
	SignData     string                       `json:"-"`
	Sign         string                       `json:"sign"`
}

type MarketingQipanCrowdWithTagCreateRsp struct {
	Response     *MarketingQipanCrowdWithTagCreate `json:"alipay_marketing_qipan_crowdwithtag_create_response"`
	AlipayCertSn string                            `json:"alipay_cert_sn,omitempty"`
	SignData     string                            `json:"-"`
	Sign         string                            `json:"sign"`
}

type MarketingQipanCrowdWithTagQueryRsp struct {
	Response     *MarketingQipanCrowdWithTagQuery `json:"alipay_marketing_qipan_crowdwithtag_query_response"`
	AlipayCertSn string                           `json:"alipay_cert_sn,omitempty"`
	SignData     string                           `json:"-"`
	Sign         string                           `json:"sign"`
}

type MarketingQipanCrowdBatchQueryRsp struct {
	Response     *MarketingQipanCrowdBatchQuery `json:"alipay_merchant_qipan_crowd_batchquery_response"`
	AlipayCertSn string                         `json:"alipay_cert_sn,omitempty"`
	SignData     string                         `json:"-"`
	Sign         string                         `json:"sign"`
}

type MarketingQipanCrowdQueryRsp struct {
	Response     *MarketingQipanCrowdQuery `json:"alipay_merchant_qipan_crowd_query_response"`
	AlipayCertSn string                    `json:"alipay_cert_sn,omitempty"`
	SignData     string                    `json:"-"`
	Sign         string                    `json:"sign"`
}

type MarketingQipanCrowdModifyRsp struct {
	Response     *MarketingQipanCrowdModify `json:"alipay_merchant_qipan_crowd_modify_response"`
	AlipayCertSn string                     `json:"alipay_cert_sn,omitempty"`
	SignData     string                     `json:"-"`
	Sign         string                     `json:"sign"`
}

type MarketingQipanBoardQueryRsp struct {
	Response     *MarketingQipanBoardQuery `json:"alipay_merchant_qipan_board_query_response"`
	AlipayCertSn string                    `json:"alipay_cert_sn,omitempty"`
	SignData     string                    `json:"-"`
	Sign         string                    `json:"sign"`
}

type MarketingQipanInsightQueryRsp struct {
	Response     *MarketingQipanInsightQuery `json:"alipay_merchant_qipan_insight_query_response"`
	AlipayCertSn string                      `json:"alipay_cert_sn,omitempty"`
	SignData     string                      `json:"-"`
	Sign         string                      `json:"sign"`
}

type MarketingQipanBehaviorQueryRsp struct {
	Response     *MarketingQipanBehaviorQuery `json:"alipay_merchant_qipan_behavior_query_response"`
	AlipayCertSn string                       `json:"alipay_cert_sn,omitempty"`
	SignData     string                       `json:"-"`
	Sign         string                       `json:"sign"`
}

type MarketingQipanTrendQueryRsp struct {
	Response     *MarketingQipanTrendQuery `json:"alipay_merchant_qipan_trend_query_response"`
	AlipayCertSn string                    `json:"alipay_cert_sn,omitempty"`
	SignData     string                    `json:"-"`
	Sign         string                    `json:"sign"`
}

type MarketingQipanInsightCityQueryRsp struct {
	Response     *MarketingQipanInsightCityQuery `json:"alipay_merchant_qipan_insightcity_query_response"`
	AlipayCertSn string                          `json:"alipay_cert_sn,omitempty"`
	SignData     string                          `json:"-"`
	Sign         string                          `json:"sign"`
}

// =========================================================分割=========================================================

type TradeRelationBind struct {
	ErrorResponse
	ResultCode string `json:"result_code"`
}

type TradeRelationBatchQuery struct {
	ErrorResponse
	ResultCode      string      `json:"result_code"`
	ReceiverList    []*Receiver `json:"receiver_list"`
	TotalPageNum    int         `json:"total_page_num"`
	TotalRecordNum  int         `json:"total_record_num"`
	CurrentPageNum  int         `json:"current_page_num"`
	CurrentPageSize int         `json:"current_page_size"`
}

type Receiver struct {
	Type          string `json:"type,omitempty"`
	Account       string `json:"account,omitempty"`
	AccountOpenId string `json:"account_open_id,omitempty"`
	Memo          string `json:"memo,omitempty"`
	LoginName     string `json:"login_name,omitempty"`
	BindLoginName string `json:"bind_login_name,omitempty"`
}

type TradeOrderSettle struct {
	ErrorResponse
	TradeNo  string `json:"trade_no,omitempty"`
	SettleNo string `json:"settle_no,omitempty"`
}

type TradeOrderSettleQuery struct {
	ErrorResponse
	OutTradeNo        string           `json:"out_request_no"`
	OperationDt       string           `json:"operation_dt"`
	RoyaltyDetailList []*RoyaltyDetail `json:"royalty_detail_list"`
}

type RoyaltyDetail struct {
	OperationType  string `json:"operation_type"`
	ExecuteDt      string `json:"execute_dt"`
	TransOut       string `json:"trans_out"`
	TransOutType   string `json:"trans_out_type"`
	TransOutOpenId string `json:"trans_out_open_id"`
	TransIn        string `json:"trans_in"`
	TransInType    string `json:"trans_in_type"`
	TransInOpenId  string `json:"trans_in_open_id"`
	Amount         string `json:"amount"`
	State          string `json:"state"`
	DetailId       string `json:"detail_id"`
	ErrorCode      string `json:"error_code"`
	ErrorDesc      string `json:"error_desc"`
}

type TradeOrderOnSettleQuery struct {
	ErrorResponse
	UnsettledAmount string `json:"unsettled_amount,omitempty"`
}

type TradeRoyaltyRateQuery struct {
	ErrorResponse
	UserId   string `json:"user_id,omitempty"`
	MaxRatio string `json:"max_ratio,omitempty"`
}

type SecurityCustomerRiskSend struct {
	ErrorResponse
}

type PayAppMarketingConsult struct {
	ErrorResponse
	PreConsultId       string         `json:"pre_consult_id"`
	ChannelInfoList    []*ChannelInfo `json:"channel_info_list"`
	ConfusedCipherList []string       `json:"confused_cipher_list"`
	BlindSignature     string         `json:"blind_signature"`
}

type ChannelInfo struct {
	ChannelIndex         string       `json:"channel_index"`
	ChannelName          string       `json:"channel_name"`
	ChannelEnable        bool         `json:"channel_enable"`
	ChannelCode          string       `json:"channel_code"`
	ChannelLogo          string       `json:"channel_logo"`
	ChannelOperationInfo string       `json:"channel_operation_info"`
	OperationList        []*Operation `json:"operation_list"`
}

type Operation struct {
	SceneCode string    `json:"scene_code"`
	ViewData  *ViewData `json:"view_data"`
}

type ViewData struct {
	OperationTip     string `json:"operation_tip"`
	OperationDesc    string `json:"operation_desc"`
	PromoType        string `json:"promo_type"`
	PromoPrice       string `json:"promo_price"`
	ThresholdAmount  string `json:"threshold_amount"`
	PayOperationInfo string `json:"pay_operation_info"`
}

type MerchantQipanCrowdCreate struct {
	ErrorResponse
	CrowdCode string `json:"crowd_code"`
}

type MerchantQipanCrowdUserAdd struct {
	ErrorResponse
}

type MerchantQipanCrowdUserDelete struct {
	ErrorResponse
}

type MarketingQipanTagBaseBatchQuery struct {
	ErrorResponse
	OperationTagList []*OperationTagItem `json:"operation_tag_list"`
}

type OperationTagItem struct {
	TagCode string `json:"tag_code"`
	TagName string `json:"tag_name"`
	TagDesc string `json:"tag_desc"`
}

type MarketingQipanTagQuery struct {
	ErrorResponse
	OperationTag *OperationTag `json:"operation_tag"`
}

type OperationTag struct {
	TagCode             string                 `json:"tag_code"`
	TagName             string                 `json:"tag_name"`
	TagDesc             string                 `json:"tag_desc"`
	OperationOptionList []*OperationOptionItem `json:"operation_option_list"`
}

type OperationOptionItem struct {
	OptionCode     string    `json:"option_code"`
	OptionName     string    `json:"option_name"`
	OptionDataType string    `json:"option_data_type"`
	OptionList     []*Option `json:"option_list"`
}

type Option struct {
	Id          int    `json:"id"`
	Value       string `json:"value"`
	Text        string `json:"text"`
	Desc        string `json:"desc"`
	ParentId    string `json:"parent_id"`
	ParentValue string `json:"parent_value"`
}

type MarketingQipanCrowdOperationCreate struct {
	ErrorResponse
	CrowdCode string `json:"crowd_code"`
}

type MarketingQipanCrowdTagQuery struct {
	ErrorResponse
	SelectTagList []*SelectTag `json:"select_tag_list"`
}

type SelectTag struct {
	TagId                 int                  `json:"tag_id"`
	TagName               string               `json:"tag_name"`
	CategoryTagOptionList []*CategoryTagOption `json:"category_tag_option_list"`
}

type CategoryTagOption struct {
	TagOptionCategoryId   int          `json:"tag_option_category_id"`
	TagOptionCategoryName string       `json:"tag_option_category_name"`
	TagOptionList         []*TagOption `json:"tag_option_list"`
}

type TagOption struct {
	Text     string      `json:"text"`
	Value    string      `json:"value"`
	Children []*Children `json:"children"`
}

type Children struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type MarketingQipanCrowdWithTagCreate struct {
	ErrorResponse
	CrowdId string `json:"crowd_id"`
}

type MarketingQipanCrowdWithTagQuery struct {
	ErrorResponse
	CountRange string `json:"count_range"`
}

type MarketingQipanCrowdBatchQuery struct {
	ErrorResponse
	TotalNumber string   `json:"total_number"`
	CrowdList   []*Crowd `json:"crowd_list"`
}

type Crowd struct {
	CrowdCode         string   `json:"crowd_code"`
	CrowdName         string   `json:"crowd_name"`
	CrowdDesc         string   `json:"crowd_desc"`
	ExternalCrowdCode string   `json:"external_crowd_code"`
	Status            string   `json:"status"`
	Processable       bool     `json:"processable"`
	ApplyChannelList  []string `json:"apply_channel_list"`
	CrowdSize         string   `json:"crowd_size"`
	Hidden            bool     `json:"hidden"`
}

type MarketingQipanCrowdQuery struct {
	ErrorResponse
	CrowdInfo *CrowdInfo `json:"crowd_info"`
}

type CrowdInfo struct {
	CrowdCode         string   `json:"crowd_code"`
	CrowdName         string   `json:"crowd_name"`
	CrowdDesc         string   `json:"crowd_desc"`
	ExternalCrowdCode string   `json:"external_crowd_code"`
	Status            string   `json:"status"`
	Processable       bool     `json:"processable"`
	ApplyChannelList  []string `json:"apply_channel_list"`
	CrowdSize         string   `json:"crowd_size"`
	Hidden            bool     `json:"hidden"`
}

type MarketingQipanCrowdModify struct {
	ErrorResponse
}

type MarketingQipanBoardQuery struct {
	ErrorResponse
	IndexList []*Index `json:"index_list"`
}

type Index struct {
	IndexKey   string `json:"index_key"`
	IndexName  string `json:"index_name"`
	IndexDesc  string `json:"index_desc"`
	IndexValue string `json:"index_value"`
	ReportDate string `json:"report_date"`
}

type MarketingQipanInsightQuery struct {
	ErrorResponse
	PortraitDataList []*PortraitData `json:"portrait_data_list"`
}

type PortraitData struct {
	PortraitKey  string      `json:"portrait_key"`
	PortraitName string      `json:"portrait_name"`
	PortraitDesc string      `json:"portrait_desc"`
	Coverage     string      `json:"coverage"`
	ReportDate   string      `json:"report_date"`
	DataList     []*DataItem `json:"data_list"`
}

type DataItem struct {
	Value string `json:"value"`
	Num   string `json:"num"`
}

type MarketingQipanBehaviorQuery struct {
	ErrorResponse
	MultiDataList []*MultiData `json:"multi_data_list"`
}

type MultiData struct {
	ObjectType string `json:"object_type"`
	ObjectId   string `json:"object_id"`
	ObjectName string `json:"object_name"`
	FeatureKey string `json:"feature_key"`
	UserCnt    string `json:"user_cnt"`
	UserRatio  string `json:"user_ratio"`
	AvgNum     string `json:"avg_num"`
}

type MarketingQipanTrendQuery struct {
	ErrorResponse
	IndexTrendResults []*IndexTrendResult `json:"index_trend_results"`
}

type IndexTrendResult struct {
	IndexKey   string `json:"index_key"`
	IndexName  string `json:"index_name"`
	IndexDesc  string `json:"index_desc"`
	IndexValue string `json:"index_value"`
	ReportDate string `json:"report_date"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}

type MarketingQipanInsightCityQuery struct {
	ErrorResponse
	PortraitKey  string  `json:"portrait_key"`
	PortraitName string  `json:"portrait_name"`
	PortraitDesc string  `json:"portrait_desc"`
	Coverage     string  `json:"coverage"`
	ReportDate   string  `json:"report_date"`
	DataList     []*Data `json:"data_list"`
}

type Data struct {
	PortraitValue string  `json:"portrait_value"`
	Num           int     `json:"num"`
	AreaCode      string  `json:"area_code"`
	CityList      []*City `json:"city_list"`
}

type City struct {
	PortraitValue string `json:"portrait_value"`
	AreaCode      string `json:"area_code"`
	Num           int    `json:"num"`
}
