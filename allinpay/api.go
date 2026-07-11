package allinpay

const (
	// 统一支付接口
	payPath = "/unitorder/pay"
	// 统一扫码接口
	scanQrPath = "/unitorder/scanqrpay"
	// 统一主扫接口
	nativePayPath = "/unitorder/nativepay"
	// 统一主扫关闭接口
	nativeClosePath = "/unitorder/closenative"
	// 交易确认查询接口
	queryConfirmPath = "/tranx/queryconfirm"
	// 交易查询接口
	queryPath = "/tranx/query"
	// 交易退款接口
	refundPath = "/tranx/refund"
	// 交易取消接口
	cancelPath = "/tranx/cancel"
	// 订单关闭
	closePath = "/unitorder/close"
)
