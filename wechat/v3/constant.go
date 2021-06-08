package wechat

const (
	Success     = 0
	SignTypeRSA = "RSA"

	MethodGet           = "GET"
	MethodPost          = "POST"
	MethodPut           = "PUT"
	MethodDelete        = "DELETE"
	HeaderAuthorization = "Authorization"

	HeaderTimestamp = "Wechatpay-Timestamp"
	HeaderNonce     = "Wechatpay-Nonce"
	HeaderSignature = "Wechatpay-Signature"
	HeaderSerial    = "Wechatpay-Serial"

	Authorization = "WECHATPAY2-SHA256-RSA2048"

	v3BaseUrlCh = "https://api.mch.weixin.qq.com" // 中国国内

	v3GetCerts = "/v3/certificates"
	// 基础支付（直连模式）
	v3ApiApp                     = "/v3/pay/transactions/app"                   // APP 下单
	v3ApiJsapi                   = "/v3/pay/transactions/jsapi"                 // JSAPI 下单
	v3ApiNative                  = "/v3/pay/transactions/native"                // Native 下单
	v3ApiH5                      = "/v3/pay/transactions/h5"                    // H5 下单
	v3ApiQueryOrderTransactionId = "/v3/pay/transactions/id/%s"                 // transaction_id 查询订单
	v3ApiQueryOrderOutTradeNo    = "/v3/pay/transactions/out-trade-no/%s"       // out_trade_no 查询订单
	v3ApiCloseOrder              = "/v3/pay/transactions/out-trade-no/%s/close" // out_trade_no 关闭订单

	// 基础支付（服务商模式）
	v3ApiPartnerPayApp                  = "/v3/pay/partner/transactions/app"                   // partner APP 下单
	v3ApiPartnerJsapi                   = "/v3/pay/partner/transactions/jsapi"                 // partner JSAPI 下单
	v3ApiPartnerNative                  = "/v3/pay/partner/transactions/native"                // partner Native 下单
	v3ApiPartnerH5                      = "/v3/pay/partner/transactions/h5"                    // partner H5 下单
	v3ApiPartnerQueryOrderTransactionId = "/v3/pay/partner/transactions/id/%s"                 // partner transaction_id 查询订单
	v3ApiPartnerQueryOrderOutTradeNo    = "/v3/pay/partner/transactions/out-trade-no/%s"       // partner out_trade_no 查询订单
	v3ApiPartnerCloseOrder              = "/v3/pay/partner/transactions/out-trade-no/%s/close" // partner out_trade_no 关闭订单

	// 基础支付（合单支付）
	v3CombinePayApp   = "/v3/combine-transactions/app"
	v3CombinePayH5    = "/v3/combine-transactions/h5"
	v3CombinePayJsapi = "/v3/combine-transactions/jsapi"
	v3CombineNative   = "/v3/combine-transactions/native"
	v3CombineQuery    = "/v3/combine-transactions/out-trade-no/%s"
	v3CombineClose    = "/v3/combine-transactions/out-trade-no/%s/close"

	// 退款
	v3DomesticRefund      = "/v3/refund/domestic/refunds"    // 申请退款
	v3DomesticRefundQuery = "/v3/refund/domestic/refunds/%s" // 查询单笔退款

	// 退款（电商收付通）
	v3CommerceRefund      = "/v3/ecommerce/refunds/apply"
	v3CommerceRefundQuery = "/v3/ecommerce/refunds/id/%s"

	// 基础支付（账单）
	v3ApiTradeBill          = "/v3/bill/tradebill"              // 申请交易账单
	v3ApiFundFlowBill       = "/v3/bill/fundflowbill"           // 申请资金账单
	v3ApiLevel2FundFlowBill = "/v3/ecommerce/bill/fundflowbill" // 申请二级商户资金账单

	// 微信支付分（免确认模式）
	v3ScoreDirectComplete = "/payscore/serviceorder/direct-complete" // 创单结单合并 POST

	// 微信支付分（免确认预授权模式）
	v3ScorePermission                = "/v3/payscore/permissions"                                 // 商户预授权 POST
	v3ScorePermissionQuery           = "/v3/payscore/permissions/authorization-code/%s"           // authorization_code 查询用户授权记录（授权协议号） GET
	v3ScorePermissionTerminate       = "/v3/payscore/permissions/authorization-code/%s/terminate" // authorization_code 解除用户授权关系（授权协议号） POST
	v3ScorePermissionOpenidQuery     = "/v3/payscore/permissions/openid/%s"                       // openid 查询用户授权记录（openid） GET
	v3ScorePermissionOpenidTerminate = "/v3/payscore/permissions/openid/%s/terminate"             // openid 解除用户授权记录（openid） POST

	// 微信支付分（公共API）
	v3ScoreOrderCreate   = "/v3/payscore/serviceorder"             // 创建支付分订单 POST
	v3ScoreOrderQuery    = "/v3/payscore/serviceorder"             // 查询支付分订单 GET
	v3ScoreOrderCancel   = "/v3/payscore/serviceorder/%s/cancel"   // out_trade_no 取消支付分订单 POST
	v3ScoreOrderModify   = "/v3/payscore/serviceorder/%s/modify"   // out_trade_no 修改订单金额 POST
	v3ScoreOrderComplete = "/v3/payscore/serviceorder/%s/complete" // out_trade_no 完结支付分订单 POST
	v3ScoreOrderPay      = "/v3/payscore/serviceorder/%s/pay"      // out_trade_no 商户发起催收扣款 POST
	v3ScoreOrderSync     = "/v3/payscore/serviceorder/%s/sync"     // out_trade_no 同步服务订单信息 POST

	// 微信先享卡
	v3CardPre     = "/v3/discount-card/cards"                     // 预受理领卡请求 POST
	v3CardAddUser = "/v3/discount-card/cards/%s/add-user-records" // out_card_code 增加用户记录 POST
	v3CardQuery   = "/v3/discount-card/cards/%s"                  // out_card_code 查询先享卡订单 GET

	// 支付即服务
	v3GuideRegister = "/v3/smartguide/guides"           // 服务人员注册 POST
	v3GuideAssign   = "/v3/smartguide/guides/%s/assign" // guide_id 服务人员分配 POST
	v3GuideQuery    = "/v3/smartguide/guides"           // 服务人员查询 GET
	v3GuideUpdate   = "/v3/smartguide/guides/%s"        // guide_id 服务人员信息更新 PATCH

	// 点金计划
	v3GoldPlanManage       = "/v3/goldplan/merchants/changegoldplanstatus"            // 点金计划管理 POST
	v3GoldPlanBillManage   = "/v3/goldplan/merchants/changecustompagestatus"          // 商家小票管理 POST
	v3GoldPlanFilterManage = "/v3/goldplan/merchants/set-advertising-industry-filter" // 同业过滤标签管理 POST
	v3GoldPlanOpenAdShow   = "/v3/goldplan/merchants/open-advertising-show"           // 开通广告展示 PATCH
	v3GoldPlanCloseAdShow  = "/v3/goldplan/merchants/close-advertising-show"          // 关闭广告展示 PATCH

	// 消费者投诉2.0
	v3ComplaintList               = "/v3/merchant-service/complaints-v2"                         // 查询投诉单列表 GET
	v3ComplaintDetail             = "/v3/merchant-service/complaints-v2/%s"                      // 查询投诉单详情 GET
	v3ComplaintNegotiationHistory = "/v3/merchant-service/complaints-v2/%s/negotiation-historys" // 查询投诉协商历史 GET
	v3ComplaintNotifyUrlCreate    = "/v3/merchant-service/complaint-notifications"               // 创建投诉通知回调地址 POST
	v3ComplaintNotifyUrlQuery     = "/v3/merchant-service/complaint-notifications"               // 查询投诉通知回调地址 GET
	v3ComplaintNotifyUrlUpdate    = "/v3/merchant-service/complaint-notifications"               // 查询投诉通知回调地址 PUT
	v3ComplaintNotifyUrlDelete    = "/v3/merchant-service/complaint-notifications"               // 删除投诉通知回调地址 DELETE
	v3ComplaintResponse           = "/v3/merchant-service/complaints-v2/%s/response"             // 提交回复 POST
	v3ComplaintComplete           = "/v3/merchant-service/complaints-v2/%s/complete"             // 反馈处理完成 POST
	v3ComplaintUploadImage        = "/v3/merchant-service/images/upload"                         // 商户上传反馈图片 POST

	v3ProfitSharingOrders         = "/v3/profitsharing/orders"                  // 请求分账 POST
	v3ProfitSharingQuery          = "/v3/profitsharing/orders"                  // 查询分账结果 GET
	v3ProfitSharingUnfreeze       = "/v3/profitsharing/orders/unfreeze"         // 解冻剩余资金 POST
	v3ProfitSharingUnsplitAmount  = "/v3/profitsharing/transactions/%s/amounts" // 查询剩余待分金额 GET
	v3ProfitSharingAddReceiver    = "/v3/profitsharing/receivers/add"           // 添加分账接收方 POST
	v3ProfitSharingDeleteReceiver = "/v3/profitsharing/receivers/delete"        // 删除分账接收方 POST

	// 订单号类型，1-微信订单号，2-商户订单号，3-微信侧回跳到商户前端时用于查单的单据查询id（查询支付分订单中会使用）
	TransactionId OrderNoType = 1
	OutTradeNo    OrderNoType = 2
	QueryId       OrderNoType = 3

	// v3 异步通知订单状态
	TradeStateSuccess  = "SUCCESS"    // 支付成功
	TradeStateRefund   = "REFUND"     // 转入退款
	TradeStateNoPay    = "NOTPAY"     // 未支付
	TradeStateClosed   = "CLOSED"     // 已关闭
	TradeStateRevoked  = "REVOKED"    // 已撤销（付款码支付）
	TradeStatePaying   = "USERPAYING" // 用户支付中（付款码支付）
	TradeStatePayError = "PAYERROR"   // 支付失败(其他原因，如银行返回失败)
)
