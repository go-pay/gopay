package wechat

// 交易、资金账单 Rsp
type BillRsp struct {
	Code     int        `json:"-"`
	SignInfo *SignInfo  `json:"-"`
	Response *TradeBill `json:"response,omitempty"`
	Error    string     `json:"-"`
}

// 二级商户资金账单 Rsp
type EcommerceFundFlowBillRsp struct {
	Code     int           `json:"-"`
	SignInfo *SignInfo     `json:"-"`
	Response *DownloadBill `json:"response,omitempty"`
	Error    string        `json:"-"`
}

// =========================================================分割=========================================================

type TradeBill struct {
	HashType    string `json:"hash_type"`
	HashValue   string `json:"hash_value"`
	DownloadUrl string `json:"download_url"`
}

type BillDetail struct {
	BillSequence int    `json:"bill_sequence"` // 商户将多个文件按账单文件序号的顺序合并为完整的资金账单文件，起始值为1
	HashType     string `json:"hash_type"`
	HashValue    string `json:"hash_value"`
	DownloadUrl  string `json:"download_url"`
	EncryptKey   string `json:"encrypt_key"` // 加密账单文件使用的加密密钥。密钥用商户证书的公钥进行加密，然后进行Base64编码
	Nonce        string `json:"nonce"`       // 加密账单文件使用的随机字符串
}

type DownloadBill struct {
	DownloadBillCount int           `json:"download_bill_count"`
	DownloadBillList  []*BillDetail `json:"download_bill_list"`
}
