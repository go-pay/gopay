package wechat

// 预受理领卡请求 Rsp
type DiscountCardApplyRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *DiscountCardApply `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 查询先享卡订单 Rsp
type DiscountCardQueryRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *DiscountCardQuery `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// =========================================================分割=========================================================

type DiscountCardApply struct {
	PrepareCardToken string `json:"prepare_card_token"` // 预领卡请求token，在引导用户进入先享卡领卡时，需要传入prepare_card_token
}

type DiscountCardQuery struct {
	CardId         string `json:"card_id"`          // 先享卡Id，唯一标识一个先享卡。
	CardTemplateId string `json:"card_template_id"` // 先享卡卡模板Id，唯一定义此资源的标识。
	Openid         string `json:"openid"`           // 微信用户在商户对应appid下的唯一标识
	OutCardCode    string `json:"out_card_code"`    // 商户领卡号
	Appid          string `json:"appid"`            // 公众账号Id
	Mchid          string `json:"mchid"`            // 商户号
	TimeRange      *struct {
		BeginTime string `json:"begin_time"` // 约定开始时间
		EndTime   string `json:"end_time"`   // 约定结束时间
	} `json:"time_range"` // 用户领取先享卡之后，约定的生效时间和到期时间。
	State            string          `json:"state"`                       // 先享卡的守约状态：ONGOING：约定进行中，SETTLING：约定到期核对中，FINISHED：已完成约定，UNFINISHED：未完成约定
	UnfinishedReason string          `json:"unfinished_reason,omitempty"` // 用户未完成约定的原因
	TotalAmount      int             `json:"total_amount,omitempty"`      // 享受优惠总金额
	PayInformation   *PayInformation `json:"pay_information,omitempty"`   // 用户退回优惠的付款信息
	CreateTime       string          `json:"create_time"`                 // 创卡时间
	Objectives       []*Objective    `json:"objectives"`                  // 用户先享卡目标列表
	Rewards          []*Reward       `json:"rewards"`                     // 用户先享卡优惠列表
	SharerOpenid     string          `json:"sharer_openid,omitempty"`     // 邀请者用户标识
}
