package wechat

// 建立合作关系 Rsp
type PartnershipsBuildRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *PartnershipsBuild `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 终止合作关系 Rsp
type PartnershipsTerminateRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *PartnershipsTerminate `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

// 查询合作关系列表 Rsp
type PartnershipsListRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *PartnershipsList `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// =========================================================分割=========================================================

type PartnershipsBuild struct {
	Partner        *Partner        `json:"partner"`         // 合作方相关的信息
	AuthorizedData *AuthorizedData `json:"authorized_data"` // 被授权的数据
	State          string          `json:"state"`           // 合作状态，ESTABLISHED：已建立，TERMINATED：已终止
	BuildTime      string          `json:"build_time"`      // 建立合作关系时间
	CreateTime     string          `json:"create_time"`     // 创建时间
	UpdateTime     string          `json:"update_time"`     // 更新时间
}

type Partner struct {
	Appid      string `json:"appid"`       // 合作方APPID
	Type       string `json:"type"`        // 合作方类别
	MerchantId string `json:"merchant_id"` // 合作方商户Id
}

type AuthorizedData struct {
	BusinessType string   `json:"business_type"` // 授权业务类别
	Scenarios    []string `json:"scenarios"`     // 授权场景
	StockId      string   `json:"stock_id"`      // 授权批次Id
}

type PartnershipsTerminate struct {
	TerminateTime string `json:"terminate_time"` // 终止合作关系时间
}

type PartnershipsList struct {
	Data       []*Partnerships `json:"data,omitempty"` // 合作关系结果集
	TotalCount int             `json:"total_count"`    // 批次总数
	Offset     int             `json:"offset"`         // 分页页码
	Limit      int             `json:"limit"`          // 分页大小
}

type Partnerships struct {
	Partner        *Partner        `json:"partner"`         // 合作方相关的信息
	AuthorizedData *AuthorizedData `json:"authorized_data"` // 被授权的数据
	BuildTime      string          `json:"build_time"`      // 建立合作关系时间
	TerminateTime  string          `json:"terminate_time"`  // 终止合作关系时间
	CreateTime     string          `json:"create_time"`     // 创建时间
	UpdateTime     string          `json:"update_time"`     // 更新时间
}