package wechat

// 商户上传(营销专用)反馈图片 Rsp
type MarketMediaUploadRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *MarketMediaUpload `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 商户上传反馈图片 Rsp
type MediaUploadRsp struct {
	Code     int          `json:"-"`
	SignInfo *SignInfo    `json:"-"`
	Response *MediaUpload `json:"response,omitempty"`
	Error    string       `json:"-"`
}

// =========================================================分割=========================================================

type MarketMediaUpload struct {
	MediaUrl string `json:"media_url"` // 微信返回的媒体文件URL地址
}

type MediaUpload struct {
	MediaId string `json:"media_id"` // 微信返回的媒体文件标识Id。
}
