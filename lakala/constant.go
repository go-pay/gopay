package lakala

const (
	baseUrlProd = "https://pay.lakala-japan.com" // 正式 URL

	// ----qr_code----
	newQrcode    = "/api/v1.0/gateway/partners/%s/orders/%s"        // partner_code、order_id 创建QRCode支付单 PUT
	nativeQrcode = "/api/v1.0/gateway/partners/%s/native_orders/%s" // partner_code、order_id 创建Native QRCode支付单 PUT
	qrcodePay    = "/api/v1.0/gateway/partners/%s/orders/%s/pay"    // partner_code、order_id QRCode支付跳转页 GET

	// ----js_api----
	newJSAPI           = "/api/v1.0/jsapi_gateway/partners/%s/orders/%s"          // partner_code、order_id 创建JSAPI订单 PUT
	newNativeJSAPI     = "/api/v1.0/gateway/partners/%s/native_jsapi/%s"          // partner_code、order_id 创建Native JSAPI订单(offline) PUT
	wechatJSAPIPay     = "/api/v1.0/wechat_jsapi_gateway/partners/%s_order_%s"    // partner_code、order_id 微信JSAPI支付跳转页 GET
	alipayJSAPIPay     = "/api/v1.0/gateway/alipay/partners/%s/orders/%s/app_pay" // partner_code、order_id 支付宝JSAPI支付跳转页 GET
	alipayPlusJSAPIPay = "/api/v1.0/alipay_connect/partners/%s/orders/%s/web_pay" // partner_code、order_id Alipay+JSAPI支付跳转页 GET

	// ----sdk_payment----
	// https://pay.lakala-japan.com/api/v1.0/gateway/partners/{partner_code}/app_orders/{order_id}
	appOrdersUrl = "/api/v1.0/gateway/partners/%s/app_orders/%s"
	// ----sdk_payment----

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
