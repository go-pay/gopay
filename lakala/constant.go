package lakala

const (
	baseUrlProd = "https://pay.lakala-japan.com" // 正式 URL

	// ----QRCode----
	newQrcodeOrder       = "/api/v1.0/gateway/partners/%s/orders/%s"        // partner_code、order_id 创建QRCode支付单 PUT
	newNativeQrcodeOrder = "/api/v1.0/gateway/partners/%s/native_orders/%s" // partner_code、order_id 创建Native QRCode支付单 PUT
	qrcodePay            = "/api/v1.0/gateway/partners/%s/orders/%s/pay"    // partner_code、order_id QRCode支付跳转页 GET

	// ----JSAPI----
	newJSAPIOrder       = "/api/v1.0/jsapi_gateway/partners/%s/orders/%s"          // partner_code、order_id 创建JSAPI订单 PUT
	newNativeJSAPIOrder = "/api/v1.0/gateway/partners/%s/native_jsapi/%s"          // partner_code、order_id 创建Native JSAPI订单(offline) PUT
	wechatJSAPIPay      = "/api/v1.0/wechat_jsapi_gateway/partners/%s_order_%s"    // partner_code、order_id 微信JSAPI支付跳转页 GET
	alipayJSAPIPay      = "/api/v1.0/gateway/alipay/partners/%s/orders/%s/app_pay" // partner_code、order_id 支付宝JSAPI支付跳转页 GET
	alipayPlusJSAPIPay  = "/api/v1.0/alipay_connect/partners/%s/orders/%s/web_pay" // partner_code、order_id Alipay+JSAPI支付跳转页 GET

	// ----H5----
	newH5Order      = "/api/v1.0/h5_payment/partners/%s/orders/%s"             // partner_code、order_id 创建H5支付单 PUT
	h5Pay           = "/api/v1.0/h5_payment/partners/%s/orders/%s/pay"         // partner_code、order_id H5支付跳转页 GET
	alipayPlusH5Pay = "/api/v1.0/alipay_connect/partners/%s/orders/%s/web_pay" // partner_code、order_id H5支付跳转页(Alipay+) GET

	// ----小程序----
	newMiniProgramOrder = "/api/v1.0/gateway/partners/%s/microapp_orders/%s" // partner_code、order_id 创建小程序订单 PUT

	// ----RetailPay----
	newRetailOrder       = "/api/v1.0/micropay/partners/%s/orders/%s"      // partner_code、order_id 创建线下支付订单 PUT
	newRetailQrcodeOrder = "/api/v1.0/retail_qrcode/partners/%s/orders/%s" // partner_code、order_id 创建线下QRCode支付单 PUT

	// ----WebGateway----
	newWebGatewayOrder = "/api/v1.0/web_gateway/partners/%s/orders/%s" // partner_code、order_id 创建渠道Web网关订单 PUT

	// ----SDKPayment----
	newSDKPaymentOrder = "/api/v1.0/gateway/partners/%s/app_orders/%s" // partner_code、order_id 创建SDK订单(Online) PUT

	// ----Custom----
	createReportSingle   = "/api/v1.0/customs/partners/%s/customs_declare/reports/%s"                       // partner_code、partner_report_id 创建报关单（非拆单） PUT
	createReportSeparate = "/api/v1.0/customs/partners/%s/customs_declare/reports/%s/sub_reports/%s"        // partner_code、partner_report_id、partner_sub_report_id 创建报关单（拆单） PUT
	queryReportStatus    = "/api/v1.0/customs/partners/%s/customs_declare/reports/%s"                       // partner_code、partner_report_id 报关状态查询 GET
	queryReportSubStatus = "/api/v1.0/customs/partners/%s/customs_declare/reports/%s/sub_reports/%s"        // partner_code、partner_report_id、partner_sub_report_id 报关子单状态查询 GET
	modifyReportSingle   = "/api/v1.0/customs/partners/%s/customs_declare_modify/reports/%s"                // partner_code、partner_report_id 修改报关信息（非拆单） PUT
	modifyReportSeparate = "/api/v1.0/customs/partners/%s/customs_declare_modify/reports/%s/sub_reports/%s" // partner_code、partner_report_id、partner_sub_report_id 修改报关信息（拆单） PUT
	reSendReportSingle   = "/api/v1.0/customs/partners/%s/customs_redeclare/reports/%s"                     // partner_code、partner_report_id 重推报关（非拆单） PUT
	modifyReportSub      = "/api/v1.0/customs/partners/%s/customs_redeclare/reports/%s/sub_reports/%s"      // partner_code、partner_report_id、partner_sub_report_id 报关单子单重推 PUT

	// ----CommonApi ----
	getExchangeRate      = "/api/v1.0/gateway/partners/%s/channel_exchange_rate" // partner_code 获取当前汇率 GET
	getEncrypt           = "/api/v1.0/gateway/partners/%s/encrypt"               // partner_code 获取加密密钥 GET
	closeOrder           = "/api/v1.0/gateway/partners/%s/orders/%s/cancel"      // partner_code、order_id 关闭订单 PUT
	getOrderStatus       = "/api/v1.0/gateway/partners/%s/orders/%s"             // partner_code、order_id 查询订单状态 GET
	applyRefund          = "/api/v1.0/gateway/partners/%s/orders/%s/refunds/%s"  // partner_code、order_id、refund_id 申请退款 PUT
	getRefundStatus      = "/api/v1.0/gateway/partners/%s/orders/%s/refunds/%s"  // partner_code、order_id、refund_id 查询退款状态 GET
	queryOrderList       = "/api/v1.0/gateway/partners/%s/orders"                // partner_code 查看账单 GET
	queryTransactionList = "/api/v1.0/gateway/partners/%s/transactions"          // partner_code 查看账单流水 GET
	querySettlements     = "/api/v1.0/gateway/partners/%s/settlements"           // partner_code 查看清算详情 GET
	queryConsultPayment  = "/api/v1.0/gateway/partners/%s/consult_payment"       // partner_code 查询可用钱包 POST
	getCoupon            = "/api/v1.0/%s/coupon/%s"                              // partner_code、coupon_id 获取优惠券信息 GET

)

// 配置结构
type Config struct {
	PartnerCode    string `toml:"PartnerCode"`    //partner_code:商户编码，由4~6位大写字母或数字构成
	CredentialCode string `toml:"credentialCode"` //credential_code:系统为商户分配的开发校验码，请妥善保管，不要在公开场合泄露
	AppId          string `toml:"appId"`          //微信appid，微信通道要求必填
	IsProd         bool   `toml:"isProd"`         //是否正式环境,沙盒
	NotifyUrl      string `toml:"notifyUrl"`      //支付回调地址
	Redirect       string `toml:"redirect"`       //可选参数。支付成功后跳转回商户APP的页面地址。
	Version        string `toml:"version"`        //客户端版本号，可选参数
}
