package wechat

// 查询投诉单列表 Rsp
type ComplaintListRsp struct {
	Code     int            `json:"-"`
	SignInfo *SignInfo      `json:"-"`
	Response *ComplaintList `json:"response,omitempty"`
	Error    string         `json:"-"`
}

// 查询投诉单详情 Rsp
type ComplaintDetailRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *ComplaintDetail `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 查询投诉单详情 Rsp
type ComplaintNegotiationHistoryRsp struct {
	Code     int                          `json:"-"`
	SignInfo *SignInfo                    `json:"-"`
	Response *ComplaintNegotiationHistory `json:"response,omitempty"`
	Error    string                       `json:"-"`
}

// 创建、查询、更新投诉通知回调地址 Rsp
type ComplaintNotifyUrlRsp struct {
	Code     int                 `json:"-"`
	SignInfo *SignInfo           `json:"-"`
	Response *ComplaintNotifyUrl `json:"response,omitempty"`
	Error    string              `json:"-"`
}

// =========================================================分割=========================================================

type ComplaintList struct {
	Data       []*ComplaintListItem `json:"data,omitempty"`        // 用户投诉信息详情
	Limit      int                  `json:"limit"`                 // 设置该次请求返回的最大投诉条数，范围【1,50】
	Offset     int                  `json:"offset"`                // 该次请求的分页开始位置，从0开始计数，例如offset=10，表示从第11条记录开始返回。
	TotalCount int                  `json:"total_count,omitempty"` // 投诉总条数，当offset=0时返回
}

type ComplaintListItem struct {
	ComplaintId           string                `json:"complaint_id"`                   // 投诉单对应的投诉单号
	ComplaintTime         string                `json:"complaint_time"`                 // 投诉时间, 例如：2015-05-20T13:29:35.120+08:00表示北京时间2015年05月20日13点29分35秒
	ComplaintDetail       string                `json:"complaint_detail"`               // 投诉的具体描述
	ComplaintState        string                `json:"complaint_state"`                // 投诉单状态, PENDING：待处理, PROCESSING：处理中, PROCESSED：已处理完成
	PayerPhone            string                `json:"payer_phone,omitempty"`          // 投诉人联系方式。该字段已做加密处理
	ComplaintOrderInfo    []*ComplaintOrderInfo `json:"complaint_order_info,omitempty"` // 投诉单关联订单信息
	ServiceOrderInfo      []*ServiceOrderInfo   `json:"service_order_info,omitempty"`   // 投诉单关联服务订单信息
	ComplaintFullRefunded bool                  `json:"complaint_full_refunded"`        // 投诉单下所有订单是否已全部全额退款
	IncomingUserResponse  bool                  `json:"incoming_user_response"`         // 投诉单是否有待回复的用户留言
	UserComplaintTimes    int                   `json:"user_complaint_times"`           // 用户投诉次数
	ComplaintMediaList    []*ComplaintMediaList `json:"complaint_media_list,omitempty"` // 投诉资料列表
	ProblemDescription    string                `json:"problem_description"`
	ProblemType           string                `json:"problem_type"`
	ApplyRefundAmount     int                   `json:"apply_refund_amount"`
	UserTagList           []string              `json:"user_tag_list"`
	AdditionalInfo        *AdditionalInfo       `json:"additional_info"`
}

type ComplaintOrderInfo struct {
	TransactionId string `json:"transaction_id"` // 投诉单关联的微信订单号
	OutTradeNo    string `json:"out_trade_no"`   // 投诉单关联的商户订单号
	Amount        int    `json:"amount"`         // 订单金额，单位（分）
}

type ServiceOrderInfo struct {
	OrderId    string `json:"order_id"`
	OutOrderNo string `json:"out_order_no"`
	State      string `json:"state"`
}

type ComplaintMediaList struct {
	MediaType string   `json:"media_type"` // 投诉单对应的投诉媒体信息
	MediaUrl  []string `json:"media_url"`
}

type AdditionalInfo struct {
	Type           string          `json:"type"`
	SharePowerInfo *SharePowerInfo `json:"share_power_info"`
}

type SharePowerInfo struct {
	ReturnTime string `json:"return_time"`
}

type ComplaintDetail struct {
	ComplaintId           string                `json:"complaint_id"`                   // 投诉单对应的投诉单号
	ComplaintTime         string                `json:"complaint_time"`                 // 投诉时间, 例如：2015-05-20T13:29:35.120+08:00表示北京时间2015年05月20日13点29分35秒
	ComplaintDetail       string                `json:"complaint_detail"`               // 投诉的具体描述
	ComplaintedMchid      string                `json:"complainted_mchid,omitempty"`    // 投诉单对应的被诉商户号。
	ComplaintState        string                `json:"complaint_state"`                // 投诉单状态, PENDING：待处理, PROCESSING：处理中, PROCESSED：已处理完成
	PayerPhone            string                `json:"payer_phone,omitempty"`          // 投诉人联系方式。该字段已做加密处理
	PayerOpenid           string                `json:"payer_openid"`                   // 投诉人在商户appid下的唯一标识
	ComplaintOrderInfo    []*ComplaintOrderInfo `json:"complaint_order_info,omitempty"` // 投诉单关联订单信息
	ComplaintMediaList    []*ComplaintMediaList `json:"complaint_media_list,omitempty"` // 投诉资料列表
	ServiceOrderInfo      []*ServiceOrderInfo   `json:"service_order_info,omitempty"`   // 投诉单关联服务订单信息
	ComplaintFullRefunded bool                  `json:"complaint_full_refunded"`        // 投诉单下所有订单是否已全部全额退款
	IncomingUserResponse  bool                  `json:"incoming_user_response"`         // 投诉单是否有待回复的用户留言
	UserComplaintTimes    int                   `json:"user_complaint_times"`           // 用户投诉次数
	ProblemDescription    string                `json:"problem_description"`
	ProblemType           string                `json:"problem_type"`
	ApplyRefundAmount     int                   `json:"apply_refund_amount"`
	UserTagList           []string              `json:"user_tag_list"`
	AdditionalInfo        *AdditionalInfo       `json:"additional_info"`
}

type ComplaintNegotiationHistory struct {
	Data       []*ComplaintNegotiationHistoryItem `json:"data,omitempty"`        // 投诉协商历史
	Limit      int                                `json:"limit"`                 // 设置该次请求返回的最大投诉条数，范围【1,50】
	Offset     int                                `json:"offset"`                // 该次请求的分页开始位置，从0开始计数，例如offset=10，表示从第11条记录开始返回。
	TotalCount int                                `json:"total_count,omitempty"` // 投诉总条数，当offset=0时返回
}

type ComplaintNegotiationHistoryItem struct {
	ComplaintMediaList *ComplaintMediaList `json:"complaint_media_list,omitempty"` // 投诉资料列表
	LogId              string              `json:"log_id"`                         // 操作流水号
	Operator           string              `json:"operator"`                       // 当前投诉协商记录的操作人
	OperateTime        string              `json:"operate_time"`                   // 当前投诉协商记录的操作时间
	OperateType        string              `json:"operate_type"`                   // 当前投诉协商记录的操作类型
	OperateDetails     string              `json:"operate_details"`                // 当前投诉协商记录的具体内容
	ImageList          []string            `json:"image_list"`                     // 当前投诉协商记录提交的图片凭证（url格式），最多返回4张图片，url有效时间为1小时。如未查询到协商历史图片凭证，则返回空数组。
}

type ComplaintNotifyUrl struct {
	Mchid string `json:"mchid"` // 返回创建回调地址的商户号，由微信支付生成并下发。
	Url   string `json:"url"`   // 通知地址，仅支持https。
}
