package wechat

// 二级商户进件 Rsp
type EcommerceApplyRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *EcommerceApply `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 查询申请状态 Rsp
type EcommerceApplyStatusRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *EcommerceApplyStatus `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 电商平台预约提现 Rsp
type EcommerceWithdrawRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *EcommerceWithdraw `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 电商平台查询预约提现状态 Rsp
type EcommerceWithdrawStatusRsp struct {
	Code     int                      `json:"-"`
	SignInfo *SignInfo                `json:"-"`
	Response *EcommerceWithdrawStatus `json:"response,omitempty"`
	Error    string                   `json:"-"`
}

// 请求分账 Rsp
type EcommerceProfitShareRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *EcommerceProfitShare `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 查询分账结果 Rsp
type EcommerceProfitShareQueryRsp struct {
	Code     int                        `json:"-"`
	SignInfo *SignInfo                  `json:"-"`
	Response *EcommerceProfitShareQuery `json:"response,omitempty"`
	Error    string                     `json:"-"`
}

// 请求分账回退 Rsp
type EcommerceProfitShareReturnRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *EcommerceProfitShareReturn `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

// 查询分账回退结果 Rsp
type EcommerceProfitShareReturnResultRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *EcommerceProfitShareReturn `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

// 完结分账 Rsp
type EcommerceProfitShareFinishRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *EcommerceProfitShareFinish `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

// 查询订单剩余待分金额 Rsp
type EcommerceProfitShareUnsplitAmountRsp struct {
	Code     int                                `json:"-"`
	SignInfo *SignInfo                          `json:"-"`
	Response *EcommerceProfitShareUnsplitAmount `json:"response,omitempty"`
	Error    string                             `json:"-"`
}

// 添加分账接收方 Rsp
type EcommerceProfitShareAddReceiverRsp struct {
	Code     int                           `json:"-"`
	SignInfo *SignInfo                     `json:"-"`
	Response *EcommerceProfitShareReceiver `json:"response,omitempty"`
	Error    string                        `json:"-"`
}

// 删除分账接收方 Rsp
type EcommerceProfitShareDeleteReceiverRsp struct {
	Code     int                           `json:"-"`
	SignInfo *SignInfo                     `json:"-"`
	Response *EcommerceProfitShareReceiver `json:"response,omitempty"`
	Error    string                        `json:"-"`
}

// 请求补差 Rsp
type EcommerceSubsidiesRsp struct {
	Code     int                 `json:"-"`
	SignInfo *SignInfo           `json:"-"`
	Response *EcommerceSubsidies `json:"response,omitempty"`
	Error    string              `json:"-"`
}

// 请求补差回退 Rsp
type EcommerceSubsidiesReturnRsp struct {
	Code     int                       `json:"-"`
	SignInfo *SignInfo                 `json:"-"`
	Response *EcommerceSubsidiesReturn `json:"response,omitempty"`
	Error    string                    `json:"-"`
}

// 取消补差 Rsp
type EcommerceSubsidiesCancelRsp struct {
	Code     int                       `json:"-"`
	SignInfo *SignInfo                 `json:"-"`
	Response *EcommerceSubsidiesCancel `json:"response,omitempty"`
	Error    string                    `json:"-"`
}

// =========================================================分割=========================================================

type EcommerceApply struct {
	ApplymentId  int64  `json:"applyment_id"`
	OutRequestNo string `json:"out_request_no"`
}

type EcommerceApplyStatus struct {
	ApplymentState     string            `json:"applyment_state"`
	ApplymentStateDesc string            `json:"applyment_state_desc"`
	SignState          string            `json:"sign_state,omitempty"`
	SignUrl            string            `json:"sign_url,omitempty"`
	SubMchid           string            `json:"sub_mchid,omitempty"`
	AccountValidation  AccountValidation `json:"account_validation"`
	AuditDetail        []*AuditDetail    `json:"audit_detail,omitempty"`
	LegalValidationUrl string            `json:"legal_validation_url,omitempty"`
	OutRequestNo       string            `json:"out_request_no"`
	ApplymentId        int64             `json:"applyment_id"`
}

type AccountValidation struct {
	AccountName              string `json:"account_name"`
	AccountNo                string `json:"account_no,omitempty"`
	PayAmount                int    `json:"pay_amount"`
	DestinationAccountNumber string `json:"destination_account_number"`
	DestinationAccountName   string `json:"destination_account_name"`
	DestinationAccountBank   string `json:"destination_account_bank"`
	City                     string `json:"city"`
	Remark                   string `json:"remark"`
	Deadline                 string `json:"deadline"`
}

type AuditDetail struct {
	ParamName    string `json:"param_name"`
	RejectReason string `json:"reject_reason"`
}

type EcommerceWithdraw struct {
	WithdrawId   string `json:"withdraw_id"`    // 微信支付预约提现单号
	OutRequestNo string `json:"out_request_no"` // 商户预约提现单号
}

type EcommerceWithdrawStatus struct {
	Status        string `json:"status"`         // 提现单状态：CREATE_SUCCESS：受理成功，SUCCESS：提现成功，FAIL：提现失败，REFUND：提现退票，CLOSE：关单，INIT：业务单已创建
	WithdrawId    string `json:"withdraw_id"`    // 微信支付提现单号
	OutRequestNo  string `json:"out_request_no"` // 商户提现单号
	Amount        int    `json:"amount"`         // 提现金额
	CreateTime    string `json:"create_time"`    // 创建时间
	UpdateTime    string `json:"update_time"`    // 更新时间
	Reason        string `json:"reason"`         // 提现失败原因，仅在提现失败、退票、关单时有值
	Remark        string `json:"remark"`         // 商户对提现单的备注，若发起提现时未传入相应值或输入不合法，则该值为空
	BankMemo      string `json:"bank_memo"`      // 展示在收款银行系统中的附言，由数字、字母、汉字组成（能否成功展示依赖银行系统支持）。若发起提现时未传入相应值或输入不合法，则该值为空
	AccountType   string `json:"account_type"`   // 出款账户类型
	AccountNumber string `json:"account_number"` // 服务商提现入账的银行账号，仅显示后四位。
	AccountBank   string `json:"account_bank"`   // 服务商提现入账的开户银行
	BankName      string `json:"bank_name"`      // 服务商提现入账的开户银行全称（含支行）
	Solution      string `json:"solution"`       // 提现失败解决方案，仅在提现失败、退票、关单时有值
}

type EcommerceProfitShare struct {
	SubMchid      string               `json:"sub_mchid"`           // 二级商户号
	TransactionId string               `json:"transaction_id"`      // 微信订单号
	OutOrderNo    string               `json:"out_order_no"`        // 商户分账单号
	OrderId       string               `json:"order_id"`            // 微信分账单号
	Status        string               `json:"status"`              // 分账单状态:PROCESSING：处理中,FINISHED：分账完成
	Receivers     []*EcommerceReceiver `json:"receivers,omitempty"` // 分账接收方列表
}

type EcommerceReceiver struct {
	Amount          int    `json:"amount"`                // 分账金额
	Description     string `json:"description"`           // 分账描述
	Type            string `json:"type"`                  // 分账接收方类型
	ReceiverAccount string `json:"receiver_account"`      // 分账接收方帐号
	ReceiverMchid   string `json:"receiver_mchid"`        // 分账接收商户号
	Result          string `json:"result"`                // 分账结果,PENDING：待分账,SUCCESS：分账成功,CLOSED：已关闭
	DetailId        string `json:"detail_id"`             // 分账明细单号
	FailReason      string `json:"fail_reason,omitempty"` // 分账失败原因ACCOUNT_ABNORMAL : 分账接收账户异常、NO_RELATION : 分账关系已解除、RECEIVER_HIGH_RISK : 高风险接收方、RECEIVER_REAL_NAME_NOT_VERIFIED : 接收方未实名、NO_AUTH : 分账权限已解除
	FinishTime      string `json:"finish_time"`           // 分账完成时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE
}

type EcommerceProfitShareQuery struct {
	SubMchid          string               `json:"sub_mchid"`                    // 二级商户号
	TransactionId     string               `json:"transaction_id"`               // 微信订单号
	OutOrderNo        string               `json:"out_order_no"`                 // 商户分账单号
	OrderId           string               `json:"order_id"`                     // 微信分账单号
	Status            string               `json:"status"`                       // 分账单状态:PROCESSING：处理中,FINISHED：分账完成
	Receivers         []*EcommerceReceiver `json:"receivers,omitempty"`          // 分账接收方列表
	FinishAmount      int                  `json:"finish_amount,omitempty"`      // 分账完结的分账金额，单位为分，仅当查询分账完结的执行结果时，存在本字段。
	FinishDescription string               `json:"finish_description,omitempty"` // 分账完结的原因描述，仅当查询分账完结的执行结果时，存在本字段。
}

type EcommerceProfitShareReturn struct {
	SubMchid    string `json:"sub_mchid"`             // 二级商户号
	OrderId     string `json:"order_id"`              // 微信分账单号，微信系统返回的唯一标识
	OutOrderNo  string `json:"out_order_no"`          // 商户分账单号
	OutReturnNo string `json:"out_return_no"`         // 商户回退单号
	ReturnNo    string `json:"return_no"`             // 微信分账回退单号，微信支付系统返回的唯一标识。
	ReturnMchid string `json:"return_mchid"`          // 只能对原分账请求中成功分给商户接收方进行回退
	Amount      int    `json:"amount"`                // 需要从分账接收方回退的金额，单位为分，只能为整数
	Result      string `json:"result"`                // 回退结果: PROCESSING：处理中，SUCCESS：已成功，FAILED：已失败
	FailReason  string `json:"fail_reason,omitempty"` // 失败原因: ACCOUNT_ABNORMAL : 分账接收方账户异常，TIME_OUT_CLOSED : 超时关单
	FinishTime  string `json:"finish_time"`           // 分账回退完成时间
}

type EcommerceProfitShareFinish struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	TransactionId string `json:"transaction_id"` // 微信订单号
	OutOrderNo    string `json:"out_order_no"`   // 商户分账单号
	OrderId       string `json:"order_id"`       // 微信分账单号
}

type EcommerceProfitShareUnsplitAmount struct {
	TransactionId string `json:"transaction_id"` // 微信订单号
	UnsplitAmount int    `json:"unsplit_amount"` // 订单剩余待分金额
}

type EcommerceProfitShareReceiver struct {
	Type    string `json:"type"`    // 接收方类型:MERCHANT_ID：商户,PERSONAL_OPENID：个人
	Account string `json:"account"` // 分账接收方帐号
}

type EcommerceSubsidies struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	TransactionId string `json:"transaction_id"` // 微信订单号
	SubsidyId     string `json:"subsidy_id"`     // 微信补差单号
	Description   string `json:"description"`    // 补差描述
	Amount        int    `json:"amount"`         // 补差金额
	Result        string `json:"result"`         // 补差单结果
	SuccessTime   string `json:"success_time"`   // 补差完成时间
}

type EcommerceSubsidiesReturn struct {
	SubMchid        string `json:"sub_mchid"`         // 二级商户号
	TransactionId   string `json:"transaction_id"`    // 微信订单号
	SubsidyRefundId string `json:"subsidy_refund_id"` // 微信补差回退单号
	RefundId        string `json:"refund_id"`         // 微信退款单号
	OutOrderNo      string `json:"out_order_no"`      // 商户补差回退单号
	Amount          int    `json:"amount"`            // 补差回退金额
	Description     string `json:"description"`       // 补差回退描述
	Result          string `json:"result"`            // 补差回退结果
	SuccessTime     string `json:"success_time"`      // 补差回退完成时间
}

type EcommerceSubsidiesCancel struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	TransactionId string `json:"transaction_id"` // 微信订单号
	Result        string `json:"result"`         // 取消补差结果
	Description   string `json:"description"`    // 取消补差描述
}
