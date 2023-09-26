package wechat

// 创建、查询、更新投诉通知回调地址 Rsp
type ViolationNotifyUrlRsp struct {
	Code     int                 `json:"-"`
	SignInfo *SignInfo           `json:"-"`
	Response *ViolationNotifyUrl `json:"response,omitempty"`
	Error    string              `json:"-"`
}

type ViolationNotifyUrl struct {
	NotifyUrl string `json:"notify_url"` // 通知地址，仅支持https。
}
