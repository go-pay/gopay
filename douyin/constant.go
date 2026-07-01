package douyin

const (
	Success = 0

	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"

	// 请求头
	HeaderAuthorization = "Authorization"
	HeaderRequestID     = "Request-Id"
	HeaderSerial        = "Douyinpay-Serial"

	// 响应时间戳与本地时间的最大允许差值（秒），超过则拒绝
	DefaultRespTimestampWindow int64 = 5 * 60

	// 应答/回调头
	HeaderTimestamp = "Douyinpay-Timestamp"
	HeaderNonce     = "Douyinpay-Nonce"
	HeaderSignature = "Douyinpay-Signature"

	// Authorization 认证类型（当前仅实现 RSA，SM2 后续按需扩展）
	Authorization = "DouyinPay-RSA"

	// 前端 JSAPI 调起签名类型
	SignTypeRSA = "DouyinPay-RSA"

	// App 端调起支付固定 package
	AppPackage = "Sign=DYPay"

	// 抖音支付开放平台域名
	baseUrlProd = "https://api.douyinpay.com"

	// 基础支付-下单（4 种方式 URL 不同）
	appOrder    = "/v1/trade/transactions/app"
	h5Order     = "/v1/trade/transactions/h5"
	jsapiOrder  = "/v1/trade/transactions/jsapi"
	nativeOrder = "/v1/trade/transactions/native"

	// 基础支付-交易共享（4 种方式 URL 一致）
	queryByTransactionId = "/v1/trade/transactions/id/%s"                 // GET，占位: transaction_id
	queryByOutTradeNo    = "/v1/trade/transactions/out-trade-no/%s"       // GET，占位: out_trade_no
	closeOrder           = "/v1/trade/transactions/out-trade-no/%s/close" // POST，占位: out_trade_no

	// 退款共享
	refund      = "/v1/trade/refund/domestic/refunds"    // POST
	refundQuery = "/v1/trade/refund/domestic/refunds/%s" // GET，占位: out_refund_no

	// 账单
	applyTradeBill  = "/v1/bill/billapply"    // GET
	applyFundBill   = "/v1/bill/fundflowbill" // GET
	applyProfitBill = "/v1/bill/splitbill"    // GET

	// 分账
	profitRequest       = "/v1/trade/profitsharing/orders"           // POST
	profitQuery         = "/v1/trade/profitsharing/orders/%s"        // GET，占位: out_order_no
	profitRollback      = "/v1/trade/profitsharing/return-orders"    // POST
	profitRollbackQuery = "/v1/trade/profitsharing/return-orders/%s" // GET，占位: out_return_no
	profitComplete      = "/v1/trade/profitsharing/finish-orders"    // POST
	profitBalanceQuery  = "/v1/trade/profitsharing/order/%s/amounts" // GET，占位: transaction_id
	receiverAdd         = "/v1/trade/profitsharing/receivers/add"    // POST
	receiverDelete      = "/v1/trade/profitsharing/receivers/delete" // POST

	// 转账
	transfer                      = "/v1/fund_trade/mch-transfer/transfer-bills"                     // POST
	transferQueryByOutBillNo      = "/v1/fund_trade/mch-transfer/transfer-bills/out-bill-no/%s"      // GET，占位: out_bill_no
	transferQueryByTransferBillNo = "/v1/fund_trade/mch-transfer/transfer-bills/transfer-bill-no/%s" // GET，占位: transfer_bill_no

	// 转账状态
	TransferStateAccepted    = "ACCEPTED"    // 已受理
	TransferStateTransfering = "TRANSFERING" // 转账中
	TransferStateSuccess     = "SUCCESS"     // 转账成功
	TransferStateFail        = "FAIL"        // 转账失败

	// 回调事件类型（event_type 字段值）
	EventTransactionSuccess = "TRANSACTION.SUCCESS" // 支付成功
	EventRefundSuccess      = "REFUND.SUCCESS"      // 退款成功
	EventRefundAbnormal     = "REFUND.ABNORMAL"     // 退款异常（暂未开放）
	EventRefundClosed       = "REFUND.CLOSED"       // 退款关闭（暂未开放）
	EventAsyncSplitFinish   = "ASYNC_SPLIT.FINISH"  // 分账结果通知
	EventSplitSuccess       = "SPLIT.SUCCESS"       // 分账动账通知
	EventTransferSuccess    = "TRANSFER.SUCCESS"    // 转账结果通知

	// 交易状态
	TradeStateSuccess = "SUCCESS"  // 支付成功
	TradeStateRefund  = "REFUND"   // 转入退款
	TradeStateNoPay   = "NOTPAY"   // 未支付
	TradeStateClosed  = "CLOSED"   // 已关闭
	TradeStatePayErr  = "PAYERROR" // 支付失败

	// 退款状态
	RefundStatusSuccess    = "SUCCESS"    // 退款成功
	RefundStatusClosed     = "CLOSED"     // 退款关闭
	RefundStatusProcessing = "PROCESSING" // 退款处理中
	RefundStatusAbnormal   = "ABNORMAL"   // 退款异常
)
