package wechat

// 特约商户进件提交申请单 Rsp
type Apply4SubSubmitRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *Apply4SubSubmit `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 特约商户进件申请单查询 Rsp
type Apply4SubQueryRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *Apply4SubQuery `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 特约商户查询结算账号 Rsp
type Apply4SubQuerySettlementRsp struct {
	Code     int                       `json:"-"`
	SignInfo *SignInfo                 `json:"-"`
	Response *Apply4SubQuerySettlement `json:"response,omitempty"`
	Error    string                    `json:"-"`
}

// =========================================================分割=========================================================

type Apply4SubSubmit struct {
	ApplymentId int64 `json:"applyment_id"` // 微信支付申请单号
}

type Apply4SubQuery struct {
	BusinessCode      string                      `json:"business_code"`       // 业务申请编号
	ApplymentId       int64                       `json:"applyment_id"`        // 微信支付申请单号
	SubMchid          string                      `json:"sub_mchid"`           // 特约商户号
	SignUrl           string                      `json:"sign_url"`            // 超级管理员签约链接
	ApplymentState    string                      `json:"applyment_state"`     // 申请单状态
	ApplymentStateMsg string                      `json:"applyment_state_msg"` // 申请状态描述
	AuditDetail       []*Applyment4SubAuditDetail `json:"audit_detail"`        // 驳回原因详情
}

type Applyment4SubAuditDetail struct {
	Field        string `json:"field"`         // 字段名
	FieldName    string `json:"field_name"`    // 字段名称
	RejectReason string `json:"reject_reason"` // 驳回原因
}

type Apply4SubQuerySettlement struct {
	AccountType      string `json:"account_type"`       // 账户类型
	AccountBank      string `json:"account_bank"`       // 开户银行
	BankName         string `json:"bank_name"`          // 开户银行全称（含支行）
	BankBranchId     string `json:"bank_branch_id"`     // 开户银行联行号
	AccountNumber    string `json:"account_number"`     // 银行账号
	VerifyResult     string `json:"verify_result"`      // 汇款验证结果
	VerifyFailReason string `json:"verify_fail_reason"` // 汇款验证失败原因
}
