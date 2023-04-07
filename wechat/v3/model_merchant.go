package wechat

// 查询特约商户账户实时余额 Rsp
type EcommerceBalanceRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *EcommerceBalance `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 查询账户实时余额 Rsp
type MerchantBalanceRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *MerchantBalance `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 商户银行来账查询 Rsp
type MerchantIncomeRecordRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *MerchantIncomeRecord `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 特约商户银行来账查询 Rsp
type PartnerIncomeRecordRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *PartnerIncomeRecord `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// =========================================================分割=========================================================

type EcommerceBalance struct {
	SubMchid        string `json:"sub_mchid"`                // 特约商户号
	AccountType     string `json:"account_type,omitempty"`   // 账户类型
	AvailableAmount int    `json:"available_amount"`         // 可用余额（单位：分），此余额可做提现操作
	PendingAmount   int    `json:"pending_amount,omitempty"` // 不可用余额（单位：分）
}

type MerchantBalance struct {
	AvailableAmount int `json:"available_amount"`         // 可用余额（单位：分），此余额可做提现操作
	PendingAmount   int `json:"pending_amount,omitempty"` // 不可用余额（单位：分）
}

type MerchantIncomeRecord struct {
	Data       []*IncomeData `json:"data,omitempty"` // 单次查询返回的银行来账记录列表结果数组，如果查询结果为空时，则为空数组
	Links      *Link         `json:"links"`          // 返回前后页和当前页面的访问链接
	Offset     int           `json:"offset"`         // 该次请求资源的起始位置，请求中包含偏移量时应答消息返回相同偏移量，否则返回默认值0
	Limit      int           `json:"limit"`          // 经过条件筛选，本次查询到的银行来账记录条数
	TotalCount int           `json:"total_count"`    // 经过条件筛选，查询到的银行来账记录总数
}

type IncomeData struct {
	Mchid             string `json:"mchid"`               // 微信支付分配的商户号
	AccountType       string `json:"account_type"`        // 需查询银行来账记录商户的账户类型：BASIC：基本账户，OPERATION：运营账户，FEES：手续费账户
	IncomeRecordType  string `json:"income_record_type"`  // 银行来账类型
	IncomeRecordId    string `json:"income_record_id"`    // 银行来账的微信单号
	Amount            int    `json:"amount"`              // 银行来账金额，单位为分，只能为整数
	SuccessTime       string `json:"success_time"`        // 银行来账完成时间
	BankName          string `json:"bank_name"`           // 银行来账的付款方银行名称，由于部分银行的数据获取限制，该字段有可能为空
	BankAccountName   string `json:"bank_account_name"`   // 银行来账的付款方银行账户信息，户名为全称、明文，由于部分银行的数据获取限制，该字段有可能为空
	BankAccountNumber string `json:"bank_account_number"` // 四位掩码+付款方银行卡尾号后四位
	RechargeRemark    string `json:"recharge_remark"`     // 银行备注
}

type Link struct {
	Next string `json:"next"` // 下一页链接
	Prev string `json:"prev"` // 上一页链接
	Self string `json:"self"` // 当前链接
}

type PartnerIncomeRecord struct {
	Data       []*PartnerIncomeData `json:"data,omitempty"` // 单次查询返回的银行来账记录列表结果数组，如果查询结果为空时，则为空数组
	Links      *Link                `json:"links"`          // 返回前后页和当前页面的访问链接
	Offset     int                  `json:"offset"`         // 该次请求资源的起始位置，请求中包含偏移量时应答消息返回相同偏移量，否则返回默认值0
	Limit      int                  `json:"limit"`          // 经过条件筛选，本次查询到的银行来账记录条数
	TotalCount int                  `json:"total_count"`    // 经过条件筛选，查询到的银行来账记录总数
}

type PartnerIncomeData struct {
	SubMchid          string `json:"sub_mchid"`           // 特约商户号
	AccountType       string `json:"account_type"`        // 需查询银行来账记录商户的账户类型：BASIC：基本账户，OPERATION：运营账户，FEES：手续费账户
	IncomeRecordType  string `json:"income_record_type"`  // 银行来账类型
	IncomeRecordId    string `json:"income_record_id"`    // 银行来账的微信单号
	Amount            int    `json:"amount"`              // 银行来账金额，单位为分，只能为整数
	SuccessTime       string `json:"success_time"`        // 银行来账完成时间
	BankName          string `json:"bank_name"`           // 银行来账的付款方银行名称，由于部分银行的数据获取限制，该字段有可能为空
	BankAccountName   string `json:"bank_account_name"`   // 银行来账的付款方银行账户信息，户名为全称、明文，由于部分银行的数据获取限制，该字段有可能为空
	BankAccountNumber string `json:"bank_account_number"` // 四位掩码+付款方银行卡尾号后四位
	RechargeRemark    string `json:"recharge_remark"`     // 银行备注
}
