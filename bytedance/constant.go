package bytedance

// https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/ecpay/APIlist/in
const (
	// 进件及提现
	appAddSubMerchant = "https://developer.toutiao.com/api/apps/ecpay/saas/app_add_sub_merchant" // POST 开发者为小程序收款商户/合作方进件
	addMerchant       = "https://developer.toutiao.com/api/apps/ecpay/saas/add_merchant"         // POST 服务商为自己进件
	getAppMerchant    = "https://developer.toutiao.com/api/apps/ecpay/saas/get_app_merchant"     // POST 服务商为小程序收款商户/合作方进件
	addSubMerchant    = "https://developer.toutiao.com/api/apps/ecpay/saas/add_sub_merchant"     // POST 服务商为第三方进件

	// 支付
	createOrder = "https://developer.toutiao.com/api/apps/ecpay/v1/create_order" // POST 预下单接口
	queryOrder  = "https://developer.toutiao.com/api/apps/ecpay/v1/query_order"  // POST 支付结果查询

	// 结算及分账
	createSettle = "https://developer.toutiao.com/api/apps/ecpay/v1/settle"       // POST 发起结算及分账
	querySettle  = "https://developer.toutiao.com/api/apps/ecpay/v1/query_settle" // POST 结算及分账结果查询

	// 退款
	createRefund = "https://developer.toutiao.com/api/apps/ecpay/v1/create_refund" // POST 发起退款
	queryRefund  = "https://developer.toutiao.com/api/apps/ecpay/v1/query_refund"  // POST 退款结果查询

	// 获取对账单
	getBill = "https://developer.toutiao.com/api/apps/bill" // GET 获取对账单
)
