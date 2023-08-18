package lakala

const (
	baseUrlProd = "https://pay.lakala-japan.com" // 正式 URL

	//----qr_code----
	// https://pay.lakala-japan.com/api/v1.0/gateway/partners/{partner_code}/orders/{order_id}
	ordersUrl = "/api/v1.0/gateway/partners/%s/orders/%s"
	// https://pay.lakala-japan.com/api/v1.0/gateway/partners/{partner_code}/native_orders/{order_id}
	nativeOrdersUrl = "/api/v1.0/gateway/partners/%s/native_orders/%s"
	// https://pay.lakala-japan.com/api/v1.0/gateway/partners/{partner_code}/orders/{order_id}/pay
	orderspayUrl = "/api/v1.0/gateway/partners/%s/native_orders/%s"
	//----qr_code----

	//----js_api----
	// https://pay.lakala-japan.com/api/v1.0/jsapi_gateway/partners/{partner_code}/orders/{order_id}
	jsApiUrl = "/api/v1.0/jsapi_gateway/partners/%s/orders/%s"
	//https://pay.lakala-japan.com/api/v1.0/gateway/partners/{partner_code}/native_jsapi/{order_id}
	jsApiNativeUrl = "/api/v1.0/gateway/partners/%s/native_jsapi/%s"
	// https://pay.lakala-japan.com/api/v1.0/wechat_jsapi_gateway/partners/{partner_code}_order_{order_id}
	jsApiWechatUrl = "/api/v1.0/wechat_jsapi_gateway/partners/%s_order_%s"
	// https://pay.lakala-japan.com/api/v1.0/gateway/alipay/partners/{partner_code}/orders/{order_id}/app_pay
	jsApiAppPayUrl = "/api/v1.0/jsapi_gateway/partners/%s/orders/%s/app_pay"
	// https://pay.lakala-japan.com/api/v1.0/alipay_connect/partners/{partner_code}/orders/{order_id}/web_pay
	jsApiWebPayUrl = "/api/v1.0/jsapi_gateway/partners/%s/orders/%s/web_pay"
	// ----js_api----

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
