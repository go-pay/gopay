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

	// 支出相关
	createBatchPayout         = "/v1/payments/payouts"                // 创建批量支出 POST
	showPayoutBatchDetail     = "/v1/payments/payouts/%s"             // payout_batch_id 获取批量支出详情 GET
	showPayoutItemDetail      = "/v1/payments/payouts-item/%s"        // payout_item_id 获取支出项目详情 GET
	cancelUnclaimedPayoutItem = "/v1/payments/payouts-item/%s/cancel" // payout_item_id 取消支出项目 POST

	// 订阅相关
	subscriptionCreate = "/v1/billing/plans" // 创建订阅 POST

	// 发票 Invoices 相关
	generateInvoiceNumber      = "/v2/invoicing/generate-next-invoice-number" // 生成发票号码 POST
	invoiceList                = "/v2/invoicing/invoices"                     // 获取发票列表 GET
	createDraftInvoice         = "/v2/invoicing/invoices"                     // 创建拟发票 POST
	deleteInvoice              = "/v2/invoicing/invoices/%s"                  // invoice_id 删除发票 DELETE
	fullyUpdateInvoice         = "/v2/invoicing/invoices/%s"                  // invoice_id 全量更新发票 PUT
	showInvoiceDetail          = "/v2/invoicing/invoices/%s"                  // invoice_id 获取发票详情 GET
	cancelSentInvoice          = "/v2/invoicing/invoices/%s/cancel"           // invoice_id 取消已发送发票 POST
	generateInvoiceQRCode      = "/v2/invoicing/invoices/%s/generate-qr-code" // invoice_id 生成发票二维码 POST
	recordPaymentForInvoice    = "/v2/invoicing/invoices/%s/payments"         // invoice_id 记录发票付款 POST
	deleteExternalPayment      = "/v2/invoicing/invoices/%s/payments/%s"      // invoice_id,transaction_id 删除额外支付 DELETE
	recordRefundForInvoice     = "/v2/invoicing/invoices/%s/refunds"          // invoice_id 记录发票退款 POST
	deleteExternalRefund       = "/v2/invoicing/invoices/%s/refunds/%s"       // invoice_id,transaction_id 删除额外退款 DELETE
	sendInvoiceReminder        = "/v2/invoicing/invoices/%s/remind"           // invoice_id 发送发票提醒 POST
	sendInvoice                = "/v2/invoicing/invoices/%s/send"             // invoice_id 发送发票 POST
	searchInvoice              = "/v2/invoicing/search-invoices"              // 搜索发票 POST
	invoiceTemplateList        = "/v2/invoicing/templates"                    // 获取发票模板列表 GET
	createInvoiceTemplate      = "/v2/invoicing/templates"                    // 创建发票模板 POST
	deleteInvoiceTemplate      = "/v2/invoicing/templates/%s"                 // template_id 删除发票模板 DELETE
	fullyUpdateInvoiceTemplate = "/v2/invoicing/templates/%s"                 // template_id 全量更新发票模板 PUT

	// 物流相关
	addTrackingNumber = "/v2/checkout/orders/%s/track" // order_id 授权物流信息 POST

	// webhook 相关
	createWebhook          = "/v1/notifications/webhooks"                 // 创建webhook POST
	listWebhook            = "/v1/notifications/webhooks"                 // 获取webhook列表 GET
	showWebhookDetail      = "/v1/notifications/webhooks/%s"              // webhook_id 获取webhook详情 GET
	updateWebhook          = "/v1/notifications/webhooks/%s"              // webhook_id 更新webhook PATCH
	deleteWebhook          = "/v1/notifications/webhooks/%s"              // webhook_id 删除webhook DELETE
	verifyWebhookSignature = "/v1/notifications/verify-webhook-signature" //webhook消息验签
	showWebhookEventDetail = "/v1/notifications/webhooks-events/%s"       // webhook_id 获取webhook-event详情 GET
)
