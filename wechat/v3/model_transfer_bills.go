package wechat

// 发起转账 Rsp
type TransferBillsRsp struct {
	Code     int            `json:"-"`
	SignInfo *SignInfo      `json:"-"`
	Response *TransferBills `json:"response,omitempty"`
	Error    string         `json:"-"`
}

// 发起撤销转账 Rsp
type TransferBillsCancelRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *TransferBillsCancel `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 商户单号查询转账单 Rsp
type TransferBillsQueryRsp struct {
	Code     int                 `json:"-"`
	SignInfo *SignInfo           `json:"-"`
	Response *TransferBillsQuery `json:"response,omitempty"`
	Error    string              `json:"-"`
}

// 微信单号查询转账单 Rsp
type TransferBillsMerchantQueryRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *TransferBillsMerchantQuery `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

// 商户单号申请电子回单 Rsp
type TransferElecsignMerchantRsp struct {
	Code     int                       `json:"-"`
	SignInfo *SignInfo                 `json:"-"`
	Response *TransferElecsignMerchant `json:"response,omitempty"`
	Error    string                    `json:"-"`
}

// 商户单号查询电子回单 Rsp
type TransferElecsignMerchantQueryRsp struct {
	Code     int                            `json:"-"`
	SignInfo *SignInfo                      `json:"-"`
	Response *TransferElecsignMerchantQuery `json:"response,omitempty"`
	Error    string                         `json:"-"`
}

// 微信单号申请电子回单 Rsp
type TransferElecsignRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *TransferElecsign `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 微信单号申请电子回单 Rsp
type TransferElecsignQueryRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *TransferElecsignQuery `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

// =========================================================分割=========================================================

type TransferBills struct {
	OutBillNo      string `json:"out_bill_no"`
	TransferBillNo string `json:"transfer_bill_no"`
	CreateTime     string `json:"create_time"`
	State          string `json:"state"`
	FailReason     string `json:"fail_reason"`
	PackageInfo    string `json:"package_info"`
}

type TransferBillsCancel struct {
	OutBillNo      string `json:"out_bill_no"`
	TransferBillNo string `json:"transfer_bill_no"`
	State          string `json:"state"`
	UpdateTime     string `json:"update_time"`
}

type TransferBillsMerchantQuery struct {
	Mchid          string `json:"mch_id"`
	OutBillNo      string `json:"out_bill_no"`
	TransferBillNo string `json:"transfer_bill_no"`
	Appid          string `json:"appid"`
	State          string `json:"state"`
	TransferAmount int    `json:"transfer_amount"`
	TransferRemark string `json:"transfer_remark"`
	FailReason     string `json:"fail_reason"`
	Openid         string `json:"openid"`
	UserName       string `json:"user_name"`
	CreateTime     string `json:"create_time"`
	UpdateTime     string `json:"update_time"`
}

type TransferBillsQuery struct {
	Mchid          string `json:"mch_id"`
	OutBillNo      string `json:"out_bill_no"`
	TransferBillNo string `json:"transfer_bill_no"`
	Appid          string `json:"appid"`
	State          string `json:"state"`
	TransferAmount int    `json:"transfer_amount"`
	TransferRemark string `json:"transfer_remark"`
	FailReason     string `json:"fail_reason"`
	Openid         string `json:"openid"`
	UserName       string `json:"user_name"`
	CreateTime     string `json:"create_time"`
	UpdateTime     string `json:"update_time"`
}

type TransferElecsignMerchant struct {
	State      string `json:"state"`
	CreateTime string `json:"create_time"`
}

type TransferElecsignMerchantQuery struct {
	State       string `json:"state"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
	HashType    string `json:"hash_type"`
	HashValue   string `json:"hash_value"`
	DownloadURL string `json:"download_url"`
	FailReason  string `json:"fail_reason"`
}

type TransferElecsign struct {
	State      string `json:"state"`
	CreateTime string `json:"create_time"`
}

type TransferElecsignQuery struct {
	State       string `json:"state"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
	HashType    string `json:"hash_type"`
	HashValue   string `json:"hash_value"`
	DownloadURL string `json:"download_url"`
	FailReason  string `json:"fail_reason"`
}
