package wechat

// 获取对私银行卡号开户银行 Rsp
type BankSearchBankRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *BankSearchBank `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 查询支持个人业务的银行列表 Rsp
type BankSearchPersonalListRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *BankSearchList `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 查询支持对公业务的银行列表 Rsp
type BankSearchCorporateListRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *BankSearchList `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 查询省份列表 Rsp
type BankSearchProvinceListRsp struct {
	Code     int                 `json:"-"`
	SignInfo *SignInfo           `json:"-"`
	Response *BankSearchProvince `json:"response,omitempty"`
	Error    string              `json:"-"`
}

// 查询城市列表 Rsp
type BankSearchCityListRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *BankSearchCity `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 查询支行列表 Rsp
type BankSearchBranchListRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *BankSearchBranch `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// =========================================================分割=========================================================

type BankSearchBank struct {
	TotalCount int         `json:"total_count"`    // 查询数据总条数
	Data       []*BankInfo `json:"data,omitempty"` // 银行列表
}

type BankInfo struct {
	BankAlias       string `json:"bank_alias"`        // 银行别名
	BankAliasCode   string `json:"bank_alias_code"`   // 银行别名编码
	AccountBank     string `json:"account_bank"`      // 开户银行
	AccountBankCode int    `json:"account_bank_code"` // 开户银行编码
	NeedBankBranch  bool   `json:"need_bank_branch"`  // 是否需要填写支行
}

type BankSearchList struct {
	TotalCount int         `json:"total_count"`    // 查询数据总条数
	Count      int         `json:"count"`          // 本次查询数据条数
	Offset     int         `json:"offset"`         // 本次查询偏移量
	Data       []*BankInfo `json:"data,omitempty"` // 银行列表
	Links      *Link       `json:"links"`          // 分页链接
}

type BankSearchProvince struct {
	TotalCount int             `json:"total_count"`    // 查询数据总条数
	Data       []*ProvinceInfo `json:"data,omitempty"` // 省份列表
}

type ProvinceInfo struct {
	ProvinceName string `json:"province_name"` // 省份名称
	ProvinceCode int    `json:"province_code"` // 省份编码
}

type BankSearchCity struct {
	TotalCount int         `json:"total_count"`    // 查询数据总条数
	Data       []*CityInfo `json:"data,omitempty"` // 城市列表
}

type CityInfo struct {
	CityName string `json:"city_name"` // 城市名称
	CityCode int    `json:"city_code"` // 城市编码
}

type BankSearchBranch struct {
	TotalCount      int               `json:"total_count"`       // 查询数据总条数
	Count           int               `json:"count"`             // 本次查询数据条数
	Offset          int               `json:"offset"`            // 本次查询偏移量
	BankAlias       string            `json:"bank_alias"`        // 银行别名
	BankAliasCode   string            `json:"bank_alias_code"`   // 银行别名编码
	AccountBank     string            `json:"account_bank"`      // 开户银行
	AccountBankCode int               `json:"account_bank_code"` // 开户银行编码
	Data            []*BankBranchInfo `json:"data,omitempty"`    // 支行列表
	Links           *Link             `json:"links"`             // 分页链接
}

type BankBranchInfo struct {
	BankBranchName string `json:"bank_branch_name"` // 开户银行支行名称
	BankBranchId   string `json:"bank_branch_id"`   // 开户银行支行联行号
}
