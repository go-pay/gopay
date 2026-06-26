package alipay

// =========================================================分割=========================================================

// 处理消费者投诉 Response
type SecurityRiskComplaintProcessFinishResponse struct {
	Response     *SecurityRiskComplaintProcessFinish `json:"alipay_security_risk_complaint_process_finish_response"`
	AlipayCertSn string                              `json:"alipay_cert_sn,omitempty"`
	SignData     string                              `json:"-"`
	Sign         string                              `json:"sign"`
}

// 投诉处理附件图片上传 Response
type SecurityRiskComplaintFileUploadResponse struct {
	Response     *SecurityRiskComplaintFileUpload `json:"alipay_security_risk_complaint_file_upload_response"`
	AlipayCertSn string                           `json:"alipay_cert_sn,omitempty"`
	SignData     string                           `json:"-"`
	Sign         string                           `json:"sign"`
}

// 查询消费者投诉详情 Response
type SecurityRiskComplaintInfoQueryResponse struct {
	Response     *SecurityRiskComplaintInfoQuery `json:"alipay_security_risk_complaint_info_query_response"`
	AlipayCertSn string                          `json:"alipay_cert_sn,omitempty"`
	SignData     string                          `json:"-"`
	Sign         string                          `json:"sign"`
}

// 查询消费者投诉列表 Response
type SecurityRiskComplaintInfoBatchqueryResponse struct {
	Response     *SecurityRiskComplaintInfoBatchquery `json:"alipay_security_risk_complaint_info_batchquery_response"`
	AlipayCertSn string                               `json:"alipay_cert_sn,omitempty"`
	SignData     string                               `json:"-"`
	Sign         string                               `json:"sign"`
}

// 营销风险识别发奖 Response
type SecurityRiskMarketingAwardingQueryResponse struct {
	Response     *SecurityRiskMarketingAwardingQuery `json:"alipay_security_risk_marketing_awarding_query_response"`
	AlipayCertSn string                              `json:"alipay_cert_sn,omitempty"`
	SignData     string                              `json:"-"`
	Sign         string                              `json:"sign"`
}

// 营销风险识别抢购 Response
type SecurityRiskMarketingPurchaseQueryResponse struct {
	Response     *SecurityRiskMarketingPurchaseQuery `json:"alipay_security_risk_marketing_purchase_query_response"`
	AlipayCertSn string                              `json:"alipay_cert_sn,omitempty"`
	SignData     string                              `json:"-"`
	Sign         string                              `json:"sign"`
}

// 行业风险识别黄牛 Response
type SecurityRiskIndustryScalperQueryResponse struct {
	Response     *SecurityRiskIndustryScalperQuery `json:"alipay_security_risk_industry_scalper_query_response"`
	AlipayCertSn string                            `json:"alipay_cert_sn,omitempty"`
	SignData     string                            `json:"-"`
	Sign         string                            `json:"sign"`
}

// 行业风险识别刷单 Response
type SecurityRiskIndustryFarmingQueryResponse struct {
	Response     *SecurityRiskIndustryFarmingQuery `json:"alipay_security_risk_industry_farming_query_response"`
	AlipayCertSn string                            `json:"alipay_cert_sn,omitempty"`
	SignData     string                            `json:"-"`
	Sign         string                            `json:"sign"`
}

// 行业风险识别先享后付违约 Response
type SecurityRiskIndustryNsfQueryResponse struct {
	Response     *SecurityRiskIndustryNsfQuery `json:"alipay_security_risk_industry_nsf_query_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

// 内容风险同步识别 Response
type SecurityRiskContentSyncDetectResponse struct {
	Response     *SecurityRiskContentSyncDetect `json:"alipay_security_risk_content_sync_detect_response"`
	AlipayCertSn string                         `json:"alipay_cert_sn,omitempty"`
	SignData     string                         `json:"-"`
	Sign         string                         `json:"sign"`
}

// =========================================================分割=========================================================

type SecurityRiskComplaintProcessFinish struct {
	ErrorResponse
	ComplaintProcessSuccess bool `json:"complaint_process_success,omitempty"` // 投诉处理是否成功
}

type SecurityRiskComplaintFileUpload struct {
	ErrorResponse
	FileName string `json:"file_name,omitempty"` // 文件名
	FileKey  string `json:"file_key,omitempty"`  // 文件key
	FileUrl  string `json:"file_url,omitempty"`  // 文件url
}

type SecurityRiskComplaintInfoQuery struct {
	ErrorResponse
	Id                     int64                 `json:"id,omitempty"`                        // 投诉id
	OppositePid            string                `json:"opposite_pid,omitempty"`              // 被投诉方pid
	OppositeName           string                `json:"opposite_name,omitempty"`             // 被投诉方名称
	ComplainAmount         string                `json:"complain_amount,omitempty"`           // 投诉金额
	Contact                string                `json:"contact,omitempty"`                   // 联系方式
	GmtComplain            string                `json:"gmt_complain,omitempty"`              // 投诉时间
	GmtProcess             string                `json:"gmt_process,omitempty"`               // 处理时间
	GmtOverdue             string                `json:"gmt_overdue,omitempty"`               // 到期时间
	ComplainContent        string                `json:"complain_content,omitempty"`          // 投诉内容
	TradeNo                string                `json:"trade_no,omitempty"`                  // 支付宝交易号
	Status                 string                `json:"status,omitempty"`                    // 投诉状态
	StatusDescription      string                `json:"status_description,omitempty"`        // 状态描述
	ProcessCode            string                `json:"process_code,omitempty"`              // 投诉处理结果码
	ProcessMessage         string                `json:"process_message,omitempty"`           // 处理结果描述
	ProcessRemark          string                `json:"process_remark,omitempty"`            // 处理备注
	ProcessImgUrlList      []string              `json:"process_img_url_list,omitempty"`      // 处理图片url列表
	GmtRiskFinishTime      string                `json:"gmt_risk_finish_time,omitempty"`      // 风险处理完结时间
	ComplaintTradeInfoList []*ComplaintTradeInfo `json:"complaint_trade_info_list,omitempty"` // 投诉交易信息列表
	TaskId                 string                `json:"task_id,omitempty"`                   // 任务id
}

type SecurityRiskComplaintInfoBatchquery struct {
	ErrorResponse
	TotalSize     int                           `json:"total_size,omitempty"`     // 总条数
	CurrentPage   int                           `json:"current_page,omitempty"`   // 当前页码
	PageSize      int                           `json:"page_size,omitempty"`      // 每页大小
	ComplaintList []*ComplaintInfoQueryResponse `json:"complaint_list,omitempty"` // 投诉列表
}

type SecurityRiskMarketingAwardingQuery struct {
	ErrorResponse
	RiskResult []*RiskQueryResult `json:"risk_result,omitempty"` // 风险分析结果
}

type SecurityRiskMarketingPurchaseQuery struct {
	ErrorResponse
	RiskResult []*RiskQueryResult `json:"risk_result,omitempty"` // 风险分析结果
}

type SecurityRiskIndustryScalperQuery struct {
	ErrorResponse
	RiskResult []*RiskQueryResult `json:"risk_result,omitempty"` // 风险分析结果
}

type SecurityRiskIndustryFarmingQuery struct {
	ErrorResponse
	RiskResult []*RiskQueryResult `json:"risk_result,omitempty"` // 风险分析结果
}

type SecurityRiskIndustryNsfQuery struct {
	ErrorResponse
	RiskResult []*RiskQueryResult `json:"risk_result,omitempty"` // 风险分析结果
}

type SecurityRiskContentSyncDetect struct {
	ErrorResponse
	RequestId         string            `json:"request_id,omitempty"`          // 请求ID
	ResultCode        string            `json:"result_code,omitempty"`         // 业务返回结果码
	ResultMsg         string            `json:"result_msg,omitempty"`          // 返回结果信息
	Suggestion        string            `json:"suggestion,omitempty"`          // 处置建议：pass/block/review
	IsMeter           bool              `json:"is_meter,omitempty"`            // 是否计费
	MeterProducts     string            `json:"meter_products,omitempty"`      // 计量产品
	IsSync            bool              `json:"is_sync,omitempty"`             // 是否同步返回
	DetectCheckLabels *DetectCheckLabel `json:"detect_check_labels,omitempty"` // 风险识别标签内容
}

// DetectCheckLabel 风险识别标签
type DetectCheckLabel struct {
	Label          string           `json:"label,omitempty"`            // 标签
	Rate           string           `json:"rate,omitempty"`             // 置信度
	SubCheckLabels []*SubCheckLabel `json:"sub_check_labels,omitempty"` // 子标签
}

// SubCheckLabel 子标签
type SubCheckLabel struct {
	HitStrategy int    `json:"hit_strategy,omitempty"` // 命中策略
	SubLabel    string `json:"sub_label,omitempty"`    // 子标签
	Rate        string `json:"rate,omitempty"`         // 置信度
}

// =========================================================分割=========================================================

// ComplaintTradeInfo 投诉交易信息
type ComplaintTradeInfo struct {
	Id                string `json:"id,omitempty"`                  // id
	ComplaintRecordId string `json:"complaint_record_id,omitempty"` // 投诉记录id
	TradeNo           string `json:"trade_no,omitempty"`            // 支付宝交易号
	OutNo             string `json:"out_no,omitempty"`              // 商户订单号
	GmtTrade          string `json:"gmt_trade,omitempty"`           // 交易时间
	GmtRefund         string `json:"gmt_refund,omitempty"`          // 退款时间
	Status            string `json:"status,omitempty"`              // 状态
	StatusDescription string `json:"status_description,omitempty"`  // 状态描述
	Amount            string `json:"amount,omitempty"`              // 金额
}

// ComplaintInfoQueryResponse 投诉详情信息
type ComplaintInfoQueryResponse struct {
	Id                     int64                 `json:"id,omitempty"`                        // 投诉id
	OppositePid            string                `json:"opposite_pid,omitempty"`              // 被投诉方pid
	OppositeName           string                `json:"opposite_name,omitempty"`             // 被投诉方名称
	ComplainAmount         string                `json:"complain_amount,omitempty"`           // 投诉金额
	Contact                string                `json:"contact,omitempty"`                   // 联系方式
	GmtComplain            string                `json:"gmt_complain,omitempty"`              // 投诉时间
	GmtProcess             string                `json:"gmt_process,omitempty"`               // 处理时间
	GmtOverdue             string                `json:"gmt_overdue,omitempty"`               // 到期时间
	ComplainContent        string                `json:"complain_content,omitempty"`          // 投诉内容
	TradeNo                string                `json:"trade_no,omitempty"`                  // 支付宝交易号
	Status                 string                `json:"status,omitempty"`                    // 投诉状态
	StatusDescription      string                `json:"status_description,omitempty"`        // 状态描述
	ProcessCode            string                `json:"process_code,omitempty"`              // 投诉处理结果码
	ProcessMessage         string                `json:"process_message,omitempty"`           // 处理结果描述
	ProcessRemark          string                `json:"process_remark,omitempty"`            // 处理备注
	ProcessImgUrlList      []string              `json:"process_img_url_list,omitempty"`      // 处理图片url列表
	GmtRiskFinishTime      string                `json:"gmt_risk_finish_time,omitempty"`      // 风险处理完结时间
	ComplaintTradeInfoList []*ComplaintTradeInfo `json:"complaint_trade_info_list,omitempty"` // 投诉交易信息列表
	TaskId                 string                `json:"task_id,omitempty"`                   // 任务id
}

// RiskQueryResult 风险分析结果
type RiskQueryResult struct {
	RiskType  string `json:"risk_type,omitempty"`  // 风险类型
	RiskValue string `json:"risk_value,omitempty"` // 风险值
	RiskDesc  string `json:"risk_desc,omitempty"`  // 风险描述
	InfoCode  string `json:"info_code,omitempty"`  // 信息码
}
