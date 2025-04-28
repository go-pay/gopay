package wechat

// 商户上传(营销专用)反馈图片 Rsp
type MarketMediaUploadRsp struct {
	Code        int                `json:"-"`
	SignInfo    *SignInfo          `json:"-"`
	Response    *MarketMediaUpload `json:"response,omitempty"`
	ErrResponse ErrResponse        `json:"err_response,omitempty"`
	Error       string             `json:"-"`
}

// 商户上传反馈图片 Rsp
type MediaUploadRsp struct {
	Code        int          `json:"-"`
	SignInfo    *SignInfo    `json:"-"`
	Response    *MediaUpload `json:"response,omitempty"`
	ErrResponse ErrResponse  `json:"err_response,omitempty"`
	Error       string       `json:"-"`
}

// ComplaintImageRsp 是原返回结构
type ComplaintImageRsp struct {
	Code        int                // 返回码
	Error       string             // 错误信息
	ErrResponse *gopay.ErrResponse // 解析后的错误
	Response    *ComplaintImage    // 原来 JSON 结构体
	ImageData   []byte             // 🔥 新增：如果是图片流，放到这里
	SignInfo    *SignInfo          // 签名信息
}

// =========================================================分割=========================================================

type MarketMediaUpload struct {
	MediaUrl string `json:"media_url"` // 微信返回的媒体文件URL地址
}

type MediaUpload struct {
	MediaId string `json:"media_id"` // 微信返回的媒体文件标识Id。
}

type ComplaintImage struct {
	MediaData *MediaData `json:"media_data"`
}

type MediaData struct {
	Filename              string `json:"filename"`
	ContentType           string `json:"content_type"`
	TotalSize             int    `json:"total_size"`
	FirstPos              int    `json:"first_pos"`
	Chunk                 string `json:"chunk"`
	Etag                  string `json:"etag"`
	SupportPartialContent bool   `json:"support_partial_content"`
}
