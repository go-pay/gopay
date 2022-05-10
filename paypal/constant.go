package paypal

const (
	Success = 0

	HeaderAuthorization       = "Authorization" // 请求头Auth
	AuthorizationPrefixBasic  = "Basic "
	AuthorizationPrefixBearer = "Bearer "

	baseUrlProd    = "https://api-m.paypal.com"         // 正式 URL
	baseUrlSandbox = "https://api-m.sandbox.paypal.com" // 沙箱 URL

	// 获取AccessToken
	getAccessToken = "/v1/oauth2/token" // 获取AccessToken POST

	// 订单相关
	orderCreate    = "/v2/checkout/orders"                           // 创建订单 POST
	orderUpdate    = "/v2/checkout/orders/%s"                        // order_id 更新订单 PATCH
	orderDetail    = "/v2/checkout/orders/%s"                        // order_id 订单详情 GET
	orderAuthorize = "/v2/checkout/orders/%s/authorize"              // order_id 订单支付授权 POST
	orderCapture   = "/v2/checkout/orders/%s/capture"                // order_id 订单支付捕获 POST
	orderConfirm   = "/v2/checkout/orders/%s/confirm-payment-source" // order_id 订单支付确认 POST

	// 支付相关
	paymentAuthorizeDetail  = "/v2/payments/authorizations/%s"             // authorization_id 支付授权详情 GET
	paymentAuthorizeCapture = "/v2/payments/authorizations/%s/capture"     // authorization_id 支付授权捕获 POST
	paymentReauthorize      = "/v2/payments/authorizations/%s/reauthorize" // authorization_id 重新授权支付授权 POST
	paymentAuthorizeVoid    = "/v2/payments/authorizations/%s/void"        // authorization_id 作废支付授权 POST
	paymentCaptureDetail    = "/v2/payments/captures/%s"                   // capture_id 支付捕获详情 GET
	paymentCaptureRefund    = "/v2/payments/captures/%s/refund"            // capture_id 支付捕获退款 POST
	paymentRefundDetail     = "/v2/payments/refunds/%s"                    // refund_id 支付退款详情 GET

	// 订阅相关
	subscriptionCreate = "v1/billing/plans" // 创建订阅 POST
)
