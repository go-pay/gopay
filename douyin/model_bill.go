package douyin

// Bill 账单申请返回体
type Bill struct {
	DownloadUrl string `json:"download_url"`         // 账单下载地址，5min 内有效
	HashType    string `json:"hash_type"`            // 哈希类型，目前固定 SHA1
	HashValue   string `json:"hash_value"`           // 原始账单摘要值（gzip 需先解压再校验）
}

// BillRsp 账单申请响应
type BillRsp struct {
	Code        int          `json:"-"`
	SignInfo    *SignInfo    `json:"-"`
	Response    *Bill        `json:"response,omitempty"`
	ErrResponse ErrResponse  `json:"-"`
	Error       string       `json:"-"`
}
