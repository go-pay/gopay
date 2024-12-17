package saobei

const (
	baseUrl        = "https://pay.lcsw.cn/lcsw"
	sandboxBaseUrl = "http://test2.lcsw.cn:8117/lcsw"

	//PayTypeWX 支付方式:微信
	PayTypeWX = "010"
	//PayTypeAli 支付方式:支付宝
	PayTypeAli = "020"
	//PayTypeQQ 支付方式:QQ钱包
	PayTypeQQ = "060"
	//PayTypeYi 支付方式:翼支付
	PayTypeYi = "100"
	//PayTypeYL 支付方式:银联二维码
	PayTypeYL = "110"

	//ResultCodeSuccess 业务结果:01 成功
	ResultCodeSuccess = "01"
	//ResultCodeFail 业务结果:02 失败
	ResultCodeFail = "02"
	//ResultCodePaying 业务结果:03 支付中
	ResultCodePaying = "03"

	//TradeStatusSuccess 交易订单状态:支付成功
	TradeStatusSuccess = "SUCCESS"
	//TradeStatusRefund 交易订单状态:转入退款
	TradeStatusRefund = "REFUND"
	//TradeStatusNotPay 交易订单状态:未支付
	TradeStatusNotPay = "NOTPAY"
	//TradeStatusClosed 交易订单状态:已关闭
	TradeStatusClosed = "CLOSED"
	//TradeStatusUserPaying 交易订单状态:用户支付中
	TradeStatusUserPaying = "USERPAYING"
	//TradeStatusRevoked 交易订单状态:已撤销
	TradeStatusRevoked = "REVOKED"
	//TradeStatusNoPay 交易订单状态:未支付支付超时
	TradeStatusNoPay = "NOPAY"
	//TradeStatusPayError 交易订单状态:支付失败
	TradeStatusPayError = "PAYERROR"
)
