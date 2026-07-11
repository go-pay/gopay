package cmbpay

// 服务器地址（详见接口文档 2.1 服务器地址）。
const (
	// HostProd 生产环境主机地址。
	HostProd = "https://api.cmbchina.com"
	// HostUAT 测试（联调）环境主机地址。
	HostUAT = "https://api.cmburl.cn:8065"
)

// 接口路径常量（详见接口文档 2.1 服务器地址与第 4 章各接口）。
// 所有路径均以 /polypay/v1.0 为前缀，直接与 Host 拼接即为完整 URL。
const (
	// PathQrCodeApply 收款码申请（4.1）。
	PathQrCodeApply = "/polypay/v1.0/mchorders/qrcodeapply"
	// PathOrderQuery 支付结果查询（4.2）。
	PathOrderQuery = "/polypay/v1.0/mchorders/orderquery"
	// PathRefund 退款申请（4.4）。
	PathRefund = "/polypay/v1.0/mchorders/refund"
	// PathRefundQuery 退款结果查询（4.5）。
	PathRefundQuery = "/polypay/v1.0/mchorders/refundquery"
	// PathPay 付款码收款（4.8）。
	PathPay = "/polypay/v1.0/mchorders/pay"
	// PathOnlinePay 微信统一下单（4.10）。
	PathOnlinePay = "/polypay/v1.0/mchorders/onlinepay"
	// PathCancel 付款码支付撤销（4.9）。
	PathCancel = "/polypay/v1.0/mchorders/cancel"
	// PathClose 关闭订单（4.7）。
	PathClose = "/polypay/v1.0/mchorders/close"
	// PathServPay 服务窗支付（4.11）。
	PathServPay = "/polypay/v1.0/mchorders/servpay"
	// PathZfbQrCode 支付宝 native 码支付（4.12）。
	PathZfbQrCode = "/polypay/v1.0/mchorders/zfbqrcode"
	// PathStatementURL 对账单下载地址获取（4.13）。
	PathStatementURL = "/polypay/v1.0/mchorders/statementurl"
	// PathKeySet 商户密钥设置（2.4）。
	PathKeySet = "/polypay/v1.0/mchkey/keyset"
	// PathOrderQrCodeApply 订单二维码申请（4.14）。
	PathOrderQrCodeApply = "/polypay/v1.0/mchorders/orderqrcodeapply"
	// PathMiniAppOrderApply 微信小程序下单（4.15）。
	PathMiniAppOrderApply = "/polypay/v1.0/mchorders/MiniAppOrderApply"
	// PathCloudPay 银联云闪付（4.16）。
	PathCloudPay = "/polypay/v1.0/mchorders/cloudpay"
	// PathEcnyUnifiedOrder 数字人民币统一下单（4.17）。
	PathEcnyUnifiedOrder = "/polypay/v1.0/mchorders/ecny/unifiedOrder"
	// PathEcnyUnifiedPayment 数字人民币统一支付（4.18）。
	PathEcnyUnifiedPayment = "/polypay/v1.0/mchorders/ecny/unifiedPayment"
	// PathEcnySubwalletPay 数字人民币子钱包支付（4.19）。
	PathEcnySubwalletPay = "/polypay/v1.0/mchorders/ecny/subwalletpay"
	// PathEcnyContractSubwalletPay 数字人民币子钱包支付-带合约（4.20）。
	PathEcnyContractSubwalletPay = "/polypay/v1.0/mchorders/ecny/contractsubwalletpay"
	// PathEcnyContractUnifiedOrder 数字人民币统一下单-带合约（4.21）。
	PathEcnyContractUnifiedOrder = "/polypay/v1.0/mchorders/ecny/contractUnifiedOrder"
	// PathPap 微信委托代扣（4.22）。
	PathPap = "/polypay/v1.0/mchorders/pap"
	// PathPapOrderQuery 微信委托代扣结果查询（4.23）。
	PathPapOrderQuery = "/polypay/v1.0/mchorders/paporderquery"
	// PathPap2 微信委托代扣-支付分（4.22）。
	PathPap2 = "/polypay/v1.0/mchorders/pap2"
	// PathZfbApp 支付宝 APP 支付（4.24）。
	PathZfbApp = "/polypay/v1.0/mchorders/zfbapp"
	// PathZfbWap 支付宝手机网站支付（4.25）。
	PathZfbWap = "/polypay/v1.0/mchorders/zfbwap"
	// PathPayScorePermissions 微信支付分预授权（4.26）。
	PathPayScorePermissions = "/polypay/v1.0/mchorders/payscore/permissions"
	// PathPayScoreQueryPermissions 微信支付分预授权查询（4.27）。
	PathPayScoreQueryPermissions = "/polypay/v1.0/mchorders/payscore/querypermissions"
	// PathPayScoreTerminatePermissions 微信支付分解除授权（4.28）。
	PathPayScoreTerminatePermissions = "/polypay/v1.0/mchorders/payscore/terminatepermissions"
	// PathPayScoreOrder 微信支付分创建订单（4.30）。
	PathPayScoreOrder = "/polypay/v1.0/mchorders/payscore/order"
	// PathPayScoreCompleteOrder 微信支付分完结订单（4.31）。
	PathPayScoreCompleteOrder = "/polypay/v1.0/mchorders/payscore/completeorder"
	// PathPayScoreQueryOrder 微信支付分查询订单（4.32）。
	PathPayScoreQueryOrder = "/polypay/v1.0/mchorders/payscore/queryorder"
	// PathPayScoreCancelOrder 微信支付分撤销订单（4.33）。
	PathPayScoreCancelOrder = "/polypay/v1.0/mchorders/payscore/cancelorder"
	// PathPayScoreModifyOrder 微信支付分修改订单金额（4.34）。
	PathPayScoreModifyOrder = "/polypay/v1.0/mchorders/payscore/modifyorder"
	// PathEcnyContractBenefit 智能合约-分润（4.36）。
	PathEcnyContractBenefit = "/polypay/v1.0/mchorders/ecny/contractBenefit"
	// PathOpenIDQryByAC 微信授权码查询 openid（4.37）。
	PathOpenIDQryByAC = "/polypay/v1.0/mchorders/openidqrybyac"
	// PathPayAfterUserPay 支付宝先享后付-统一收单交易支付（4.38）。
	PathPayAfterUserPay = "/polypay/v1.0/mchorders/payafteruser/pay"
	// PathAlipayMarketingConsult 支付宝-商户前置内容咨询。
	PathAlipayMarketingConsult = "/polypay/v1.0/mchorders/alipay/marketing/consult"
	// PathAlipayShareTokenCreate 支付宝-吱口令获取。
	PathAlipayShareTokenCreate = "/polypay/v1.0/mchorders/alipay/sharetoken/create"
)
