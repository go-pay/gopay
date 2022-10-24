package unionpay

// 常量参数
const (
	Format      = "JSON"
	Charset     = "UTF-8"
	ContentType = "application/json"
)

// 通用应答参数，用于判断网关状态码
type CommonResponseParams struct {
	ErrCode string `json:"errCode,omitempty"`
	ErrInfo string `json:"errInfo,omitempty"`
}

// 商家扫用户请求
type MchScanRequest struct {
	MerchantCode            string `json:"merchantCode,omitempty"`            // 必填 商户号 服务商商户号
	TerminalCode            string `json:"terminalCode,omitempty"`            // 必填 终端号
	TransactionAmount       int64  `json:"transactionAmount,omitempty"`       // 必填 订单金额 单位：分
	TransactionCurrencyCode string `json:"transactionCurrencyCode,omitempty"` // 必填 交易币种 需填入156
	MerchantOrderId         string `json:"merchantOrderId,omitempty"`         // 必填 商户单号，全局唯一，不可重复，长度不超过50位
	MerchantRemark          string `json:"merchantRemark,omitempty"`          // 必填 商户备注
	PayMode                 string `json:"payMode,omitempty"`                 // 必填 支付方式
	PayCode                 string `json:"payCode,omitempty"`                 // 必填 支付码
}

// 商家扫用户应答
type MchScanResponse struct {
	ErrCode                 string `json:"errCode,omitempty"`
	ErrInfo                 string `json:"errInfo,omitempty"`
	Amount                  int64  `json:"amount,omitempty"`                  // 必填 用户实际支付金额，最大长度12 位
	OrderId                 string `json:"orderId,omitempty"`                 // 必填 银联商务的订单号
	ThirdPartyOrderId       string `json:"thirdPartyOrderId,omitempty"`       // 第三方订单号
	TotalDiscountAmount     int64  `json:"totalDiscountAmount,omitempty"`     // 优惠金额（合计）
	DiscountStatus          int64  `json:"discountStatus,omitempty"`          // 优惠状态 1：订单有优惠但未找到 2：订单有优惠且找到
	ActualTransactionAmount int64  `json:"actualTransactionAmount,omitempty"` // 营销联盟优惠后交易金额 实收金额
}

// 交易退款请求
type RefundRequest struct {
	MerchantCode      string `json:"merchantCode,omitempty"`      // 必填 商户号
	TerminalCode      string `json:"terminalCode,omitempty"`      // 必填 终端号
	MerchantOrderId   string `json:"merchantOrderId,omitempty"`   // 条件必填 商户单号，全局唯一，不可重复，长度不超过50位
	OriginalOrderId   string `json:"originalOrderId,omitempty"`   // 条件必填 银商订单号
	RefundRequestId   string `json:"refundRequestId,omitempty"`   // 必填 标识一次退款 请求，同一笔订单多次退款 需要保证唯一，长度不超过50位
	TransactionAmount int64  `json:"transactionAmount,omitempty"` // 必填 退货金额
}

// 交易退款应答 是同步返回结果的，errCode 为 00 则代表退款成功
type RefundResponse struct {
	ErrCode string `json:"errCode,omitempty"`
	ErrInfo string `json:"errInfo,omitempty"`
}

// 交易撤销请求
type CancelRequest struct {
	MerchantCode    string `json:"merchantCode,omitempty"`    // 必填 商户号
	TerminalCode    string `json:"terminalCode,omitempty"`    // 必填 终端号
	MerchantOrderId string `json:"merchantOrderId,omitempty"` // 条件必填 商户单号，全局唯一，不可重复，长度不超过50位
	OriginalOrderId string `json:"originalOrderId,omitempty"` // 条件必填 银商订单号
}

// 交易撤销应答 是同步返回结果的，errCode 为 00 则代表退款成功
type CancelResponse struct {
	ErrCode string `json:"errCode,omitempty"`
	ErrInfo string `json:"errInfo,omitempty"`
}

// 交易查询请求
type QueryRequest struct {
	MerchantCode    string `json:"merchantCode,omitempty"`    // 必填 商户号
	TerminalCode    string `json:"terminalCode,omitempty"`    // 必填 终端号
	MerchantOrderId string `json:"merchantOrderId,omitempty"` // 条件必填 商户单号，全局唯一，不可重复，长度不超过50位
	OriginalOrderId string `json:"originalOrderId,omitempty"` // 条件必填 银商订单号
}

// 交易查询应答
type QueryResponse struct {
	ErrCode                              string `json:"errCode,omitempty"`
	ErrInfo                              string `json:"errInfo,omitempty"`
	QueryResCode                         string `json:"queryResCode,omitempty"`                         // 0：成功 1：超时 2：已撤销 3：已退货 4：已冲正 5：失败（失败情况，后面追加 失败描述) FF：交易状态未知
	QueryResDesc                         string `json:"queryResDesc,omitempty"`                         // 查询结果描述
	OriginalSystemTraceNum               string `json:"originalSystemTraceNum,omitempty"`               // 原交易流水号
	OriginalTransactionAmount            int64  `json:"originalTransactionAmount,omitempty"`            // 原交易金额
	OrderId                              string `json:"orderId,omitempty"`                              // 必填 订单号
	RefundedAmount                       int64  `json:"refundedAmount,omitempty"`                       // 已退货金额
	ActualTransactionAmount              int64  `json:"actualTransactionAmount,omitempty"`              // 营销联盟优惠后交易金额
	Amount                               int64  `json:"amount,omitempty"`                               // 必填 用户实际支付金额，最大长度12位
	MarketingAllianceDiscountInstruction string `json:"marketingAllianceDiscountInstruction,omitempty"` // 交易营销联盟优惠说明
	ThirdPartyOrderId                    string `json:"thirdPartyOrderId,omitempty"`                    // 第三方订单号
	TotalDiscountAmount                  int64  `json:"totalDiscountAmount,omitempty"`                  // 优惠金额（合计）
	DiscountStatus                       int64  `json:"discountStatus,omitempty"`                       // 优惠状态 1：订单有优惠但未找到 2：订单有优惠且找到
}

// 交易退款查询请求
type QueryRefundRequest struct {
	MerchantCode    string `json:"merchantCode,omitempty"`    // 必填 商户号
	TerminalCode    string `json:"terminalCode,omitempty"`    // 必填 终端号
	MerchantOrderId string `json:"merchantOrderId,omitempty"` // 条件必填 原支付商户单号，全局唯一，不可重复，长度不超过50位
	OriginalOrderId string `json:"originalOrderId,omitempty"` // 条件必填 银商订单号
	RefundRequestId string `json:"refundRequestId,omitempty"` // 必填 退款请求标识 标识一次退款请求，同一笔订单多次退款需要保证唯一，长度不超过50位
}

// 交易退款查询应答
type QueryRefundResponse struct {
	ErrCode      string `json:"errCode,omitempty"`
	ErrInfo      string `json:"errInfo,omitempty"`
	QueryResCode string `json:"queryResCode,omitempty"` // 00：成功，其余 均为失败
	QueryResInfo string `json:"queryResInfo,omitempty"` // 当queryResCode 不为00时存在
}

// URL 接口地址
const (
	// 测试环境
	MchScanApiBeta     = "http://0.0.0.0:12345/v2/poslink/transaction/pay"          // 聚合反扫 商家扫用户
	QueryApiBeta       = "http://0.0.0.0:12345/v2/poslink/transaction/query"        // 支付查询
	RefundApiBeta      = "http://0.0.0.0:12345/v2/poslink/transaction/refund"       // 退款
	QueryRefundApiBeta = "http://0.0.0.0:12345/v2/poslink/transaction/query-refund" // 退款查询
	CancelApiBeta      = "http://0.0.0.0:12345/v2/poslink/transaction/voidpayment"  // 交易撤销

	// 正式环境
	MchScanApi     = "https://api-mop.chinaums.com/v2/poslink/transaction/pay"          // 聚合反扫 商家扫用户
	QueryApi       = "https://api-mop.chinaums.com/v2/poslink/transaction/query"        // 支付查询
	RefundApi      = "https://api-mop.chinaums.com/v2/poslink/transaction/refund"       // 退款
	QueryRefundApi = "https://api-mop.chinaums.com/v2/poslink/transaction/query-refund" // 退款查询
	CancelApi      = "https://api-mop.chinaums.com/v2/poslink/transaction/voidpayment"  // 交易撤销
)

const (
	TransactionCurrencyCode = "156" // 交易币种,固定值

	PayMode_E_CASH    = "E_CASH"    // 电子现金
	PayMode_SOUNDWAVE = "SOUNDWAVE" // 声波
	PayMode_NFC       = "NFC"       // NFC
	PayMode_CODE_SCAN = "CODE_SCAN" // 扫码
	PayMode_MANUAL    = "MANUAL"    // 手输
)

const (
	GateWaySuccess = "00" // 网关成功状态码
	// BizStatusSuccess = "00" // 业务成功状态码
)

// 交易查询状态码
const (
	QueryResCode_Success = "0"  // 成功
	QueryResCode_Timeout = "1"  // 超时
	QueryResCode_Cancel  = "2"  // 已撤销
	QueryResCode_Return  = "3"  // 已退货
	QueryResCode_Rush    = "4"  // 已冲正
	QueryResCode_Failed  = "5"  // 失败（失败情况，后面追加 失败描述)
	QueryResCode_Unkonwn = "FF" // 交易状态未知
)
