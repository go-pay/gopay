package wechat

const (
	Success     = 0
	SignTypeRSA = "RSA"
	SignTypeSM2 = "SM2"

	CertTypeALL CertType = "ALL"
	CertTypeRSA CertType = "RSA"
	CertTypeSM2 CertType = "SM2"

	MethodGet           = "GET"
	MethodPost          = "POST"
	MethodPut           = "PUT"
	MethodDelete        = "DELETE"
	MethodPATCH         = "PATCH"
	HeaderAuthorization = "Authorization"
	HeaderRequestID     = "Request-ID"

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
	v3DomesticRefundQuery = "/v3/refund/domestic/refunds/%s" // out_refund_no 查询单笔退款

	// 账单
	v3TradeBill             = "/v3/bill/tradebill"                 // 申请交易账单 GET
	v3FundFlowBill          = "/v3/bill/fundflowbill"              // 申请资金账单 GET
	v3EcommerceFundFlowBill = "/v3/ecommerce/bill/fundflowbill"    // 申请特约商户资金账单 GET
	v3SubFundFlowBill       = "/v3/bill/sub-merchant-fundflowbill" // 申请单个子商户资金账单 GET

	// 提现
	v3Withdraw                    = "/v3/ecommerce/fund/withdraw"                   // 特约商户余额提现 POST
	v3WithdrawStatusById          = "/v3/ecommerce/fund/withdraw/%s"                // withdraw_id 查询特约商户提现状态 GET
	v3WithdrawStatusByNo          = "/v3/ecommerce/fund/withdraw/out-request-no/%s" // out_request_no 查询特约商户提现状态 GET
	v3EcommerceWithdraw           = "/v3/merchant/fund/withdraw"                    // 电商平台预约提现 POST
	v3EcommerceWithdrawStatusById = "/v3/merchant/fund/withdraw/withdraw-id/%s"     // withdraw_id 电商平台查询预约提现状态 POST
	v3EcommerceWithdrawStatusByNo = "/v3/merchant/fund/withdraw/out-request-no/%s"  // out_request_no 电商平台查询预约提现状态 POST
	v3WithdrawDownloadErrBill     = "/v3/merchant/fund/withdraw/bill-type/%s"       // bill_type 按日下载提现异常文件 GET

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
	v3GuideReg    = "/v3/smartguide/guides"           // 服务人员注册 POST
	v3GuideAssign = "/v3/smartguide/guides/%s/assign" // guide_id 服务人员分配 POST
	v3GuideQuery  = "/v3/smartguide/guides"           // 服务人员查询 GET
	v3GuideUpdate = "/v3/smartguide/guides/%s"        // guide_id 服务人员信息更新 PATCH

	// 智慧商圈
	v3BusinessPointsSync        = "/v3/businesscircle/points/notify"                 // 商圈积分同步 POST
	v3BusinessAuthPointsQuery   = "/v3/businesscircle/user-authorizations/%s"        // openid 商圈积分授权查询 GET
	v3BusinessPointsStatusQuery = "/v3/businesscircle/users/%s/points/commit_status" // openid 商圈会员待积分状态查询 GET
	v3BusinessParkingSync       = "/v3/businesscircle/parkings"                      // 商圈会员停车状态同步 POST

	// 微信支付分停车服务
	v3VehicleParkingQuery = "/v3/vehicle/parking/services/find"        // 查询车牌服务开通信息 GET
	v3VehicleParkingIn    = "/v3/vehicle/parking/parkings"             // 创建停车入场 POST
	v3VehicleParkingFee   = "/v3/vehicle/transactions/parking"         // 扣费受理 POST
	v3VehicleParkingOrder = "/v3/vehicle/transactions/out-trade-no/%s" // out_trade_no 查询订单 GET

	// 代金券
	v3FavorBatchCreate        = "/v3/marketing/favor/coupon-stocks"         // 创建代金券批次 POST
	v3FavorBatchStart         = "/v3/marketing/favor/stocks/%s/start"       // stock_id 激活代金券批次 POST
	v3FavorBatchGrant         = "/v3/marketing/favor/users/%s/coupons"      // openid 发放代金券批次 POST
	v3FavorBatchPause         = "/v3/marketing/favor/stocks/%s/pause"       // stock_id 暂停代金券批次 POST
	v3FavorBatchRestart       = "/v3/marketing/favor/stocks/%s/restart"     // stock_id 重启代金券批次 POST
	v3FavorBatchList          = "/v3/marketing/favor/stocks"                // 条件查询批次列表 GET
	v3FavorBatchDetail        = "/v3/marketing/favor/stocks/%s"             // stock_id 查询批次详情 GET
	v3FavorDetail             = "/v3/marketing/favor/users/%s/coupons/%s"   // openid、coupon_id 查询代金券详情 GET
	v3FavorMerchant           = "/v3/marketing/favor/stocks/%s/merchants"   // stock_id 查询代金券可用商户 GET
	v3FavorItems              = "/v3/marketing/favor/stocks/%s/items"       // stock_id 查询代金券可用单品 GET
	v3FavorUserCoupons        = "/v3/marketing/favor/users/%s/coupons"      // openid 根据商户号查用户的券 GET
	v3FavorUseFlowDownload    = "/v3/marketing/favor/stocks/%s/use-flow"    // stock_id 下载批次核销明细 GET
	v3FavorRefundFlowDownload = "/v3/marketing/favor/stocks/%s/refund-flow" // stock_id 下载批次退款明细 GET
	v3FavorCallbackUrlSet     = "/v3/marketing/favor/callbacks"             // 设置消息通知地址 POST
	v3FavorMediaUploadImage   = "/v3/marketing/favor/media/image-upload"    // 图片上传(营销专用) POST

	// 商家券
	v3BusiFavorBatchCreate      = "/v3/marketing/busifavor/stocks"                        // 创建商家券 POST
	v3BusiFavorBatchDetail      = "/v3/marketing/busifavor/stocks/%s"                     // stock_id 查询商家券详情 GET
	v3BusiFavorUse              = "/v3/marketing/busifavor/coupons/use"                   // 核销用户券 POST
	v3BusiFavorUserCoupons      = "/v3/marketing/busifavor/users/%s/coupons"              // openid 根据过滤条件查询用户券 GET
	v3BusiFavorUserCouponDetail = "/v3/marketing/busifavor/users/%s/coupons/%s/appids/%s" // openid、coupon_code、appid 查询用户单张券详情 GET
	v3BusiFavorCodeUpload       = "/v3/marketing/busifavor/stocks/%s/couponcodes"         // stock_id 上传预存code POST
	v3BusiFavorCallbackUrlSet   = "/v3/marketing/busifavor/callbacks"                     // 设置商家券事件通知地址 POST
	v3BusiFavorCallbackUrl      = "/v3/marketing/busifavor/callbacks"                     // 查询商家券事件通知地址 GET
	v3BusiFavorAssociate        = "/v3/marketing/busifavor/coupons/associate"             // 关联订单信息 POST
	v3BusiFavorDisassociate     = "/v3/marketing/busifavor/coupons/disassociate"          // 取消关联订单信息 POST
	v3BusiFavorBatchUpdate      = "/v3/marketing/busifavor/stocks/%s/budget"              // stock_id 修改批次预算 PATCH
	v3BusiFavorInfoUpdate       = "/v3/marketing/busifavor/stocks/%s"                     // stock_id 修改商家券基本信息 PATCH
	v3BusiFavorSend             = "/v3/marketing/busifavor/coupons/%s/send"               // card_id 发放消费卡 POST
	v3BusiFavorReturn           = "/v3/marketing/busifavor/coupons/return"                // 申请退券 POST
	v3BusiFavorDeactivate       = "/v3/marketing/busifavor/coupons/deactivate"            // 使券失效 POST
	v3BusiFavorSubsidyPay       = "/v3/marketing/busifavor/subsidy/pay-receipts"          // 营销补差付款 POST
	v3BusiFavorSubsidyPayDetail = "/v3/marketing/busifavor/subsidy/pay-receipts/%s"       // subsidy_receipt_id 查询营销补差付款单详情 GET

	// 委托营销（合作伙伴）
	v3PartnershipsBuild     = "/v3/marketing/partnerships/build"     // 建立合作关系 POST
	v3PartnershipsTerminate = "/v3/marketing/partnerships/terminate" // 终止合作关系 POST
	v3PartnershipsList      = "/v3/marketing/partnerships"           // 查询合作关系列表 GET

	// 点金计划（服务商）
	v3GoldPlanManage       = "/v3/goldplan/merchants/changegoldplanstatus"            // 点金计划管理 POST
	v3GoldPlanBillManage   = "/v3/goldplan/merchants/changecustompagestatus"          // 商家小票管理 POST
	v3GoldPlanFilterManage = "/v3/goldplan/merchants/set-advertising-industry-filter" // 同业过滤标签管理 POST
	v3GoldPlanOpenAdShow   = "/v3/goldplan/merchants/open-advertising-show"           // 开通广告展示 PATCH
	v3GoldPlanCloseAdShow  = "/v3/goldplan/merchants/close-advertising-show"          // 关闭广告展示 POST

	// 消费者投诉2.0
	v3ComplaintList                 = "/v3/merchant-service/complaints-v2"                           // 查询投诉单列表 GET
	v3ComplaintDetail               = "/v3/merchant-service/complaints-v2/%s"                        // complaint_id 查询投诉单详情 GET
	v3ComplaintNegotiationHistory   = "/v3/merchant-service/complaints-v2/%s/negotiation-historys"   // complaint_id 查询投诉协商历史 GET
	v3ComplaintNotifyUrlCreate      = "/v3/merchant-service/complaint-notifications"                 // 创建投诉通知回调地址 POST
	v3ComplaintNotifyUrlQuery       = "/v3/merchant-service/complaint-notifications"                 // 查询投诉通知回调地址 GET
	v3ComplaintNotifyUrlUpdate      = "/v3/merchant-service/complaint-notifications"                 // 查询投诉通知回调地址 PUT
	v3ComplaintNotifyUrlDelete      = "/v3/merchant-service/complaint-notifications"                 // 删除投诉通知回调地址 DELETE
	v3ComplaintResponse             = "/v3/merchant-service/complaints-v2/%s/response"               // complaint_id 回复用户 POST
	v3ComplaintComplete             = "/v3/merchant-service/complaints-v2/%s/complete"               // complaint_id 反馈处理完成 POST
	v3ComplaintUploadImage          = "/v3/merchant-service/images/upload"                           // 商户上传反馈图片 POST
	v3ComplaintUpdateRefundProgress = "/v3/merchant-service/complaints-v2/%s/update-refund-progress" // complaint_id 更新退款审批结果 POST

	// 商户平台处置通知
	v3ViolationNotifyUrlCreate = "/v3/merchant-risk-manage/violation-notifications" // 创建商户违规通知回调地址 POST
	v3ViolationNotifyUrlQuery  = "/v3/merchant-risk-manage/violation-notifications" // 查询商户违规通知回调地址 GET
	v3ViolationNotifyUrlUpdate = "/v3/merchant-risk-manage/violation-notifications" // 查询商户违规通知回调地址 PUT
	v3ViolationNotifyUrlDelete = "/v3/merchant-risk-manage/violation-notifications" // 删除商户违规通知回调地址 DELETE

	// 分账（服务商）
	v3ProfitShareOrder           = "/v3/profitsharing/orders"                  // 请求分账 POST
	v3ProfitShareQuery           = "/v3/profitsharing/orders/%s"               // 查询分账结果 GET
	v3ProfitShareReturn          = "/v3/profitsharing/return-orders"           // 请求分账回退 POST
	v3ProfitShareReturnResult    = "/v3/profitsharing/return-orders/%s"        // 查询分账回退结果 GET
	v3ProfitShareUnfreeze        = "/v3/profitsharing/orders/unfreeze"         // 解冻剩余资金 POST
	v3ProfitShareUnsplitAmount   = "/v3/profitsharing/transactions/%s/amounts" // 查询剩余待分金额 GET
	v3ProfitShareAddReceiver     = "/v3/profitsharing/receivers/add"           // 添加分账接收方 POST
	v3ProfitShareDeleteReceiver  = "/v3/profitsharing/receivers/delete"        // 删除分账接收方 POST
	v3ProfitShareMerchantConfigs = "/v3/profitsharing/merchant-configs/%s"     // 查询最大分账比例API GET
	v3ProfitShareBills           = "/v3/profitsharing/bills"                   // 申请分账账单 GET

	// 其他能力
	v3MediaUploadImage = "/v3/merchant/media/upload"       // 图片上传 POST
	v3MediaUploadVideo = "/v3/merchant/media/video_upload" // 视频上传 POST

	// 转账
	v3Transfer                   = "/v3/transfer/batches"                                          // 发起商家转账 POST
	v3TransferQuery              = "/v3/transfer/batches/batch-id/%s"                              // batch_id 微信批次单号查询批次单 GET
	v3TransferDetail             = "/v3/transfer/batches/batch-id/%s/details/detail-id/%s"         // batch_id、detail_id 微信明细单号查询明细单 GET
	v3TransferMerchantQuery      = "/v3/transfer/batches/out-batch-no/%s"                          // out_batch_no 商家批次单号查询批次单 GET
	v3TransferMerchantDetail     = "/v3/transfer/batches/out-batch-no/%s/details/out-detail-no/%s" // out_batch_no、out_detail_no 商家明细单号查询明细单 GET
	v3TransferReceipt            = "/v3/transfer/bill-receipt"                                     // 转账电子回单申请受理 POST
	v3TransferReceiptQuery       = "/v3/transfer/bill-receipt/%s"                                  // out_batch_no 查询转账电子回单 GET
	v3TransferDetailReceipt      = "/v3/transfer-detail/electronic-receipts"                       // 转账明细电子回单受理 POST
	v3TransferDetailReceiptQuery = "/v3/transfer-detail/electronic-receipts"                       // 查询转账明细电子回单受理结果 GET

	// 转账（服务商）
	v3PartnerTransfer               = "/v3/partner-transfer/batches"                                          // 发起批量转账 POST
	v3PartnerTransferQuery          = "/v3/partner-transfer/batches/batch-id/%s"                              // batch_id 微信批次单号查询批次单 GET
	v3PartnerTransferDetail         = "/v3/partner-transfer/batches/batch-id/%s/details/detail-id/%s"         // batch_id、detail_id 微信明细单号查询明细单 GET
	v3PartnerTransferMerchantQuery  = "/v3/partner-transfer/batches/out-batch-no/%s"                          // out_batch_no 商家批次单号查询批次单 GET
	v3PartnerTransferMerchantDetail = "/v3/partner-transfer/batches/out-batch-no/%s/details/out-detail-no/%s" // out_batch_no、out_detail_no 商家明细单号查询明细单 GET

	// 余额
	v3MerchantBalance     = "/v3/merchant/fund/balance/%s"        // account_type 查询账户实时余额 GET
	v3MerchantDayBalance  = "/v3/merchant/fund/dayendbalance/%s"  // account_type 查询账户日终余额 GET
	v3EcommerceBalance    = "/v3/ecommerce/fund/balance/%s"       // sub_mchid 查询特约商户账户实时余额 GET
	v3EcommerceDayBalance = "/v3/ecommerce/fund/enddaybalance/%s" // sub_mchid 查询二级商户账户日终余额 GET

	// 来账识别API
	v3MerchantIncomeRecord  = "/v3/merchantfund/merchant/income-records" // 商户银行来账查询 GET
	v3EcommerceIncomeRecord = "/v3/merchantfund/partner/income-records"  // 特约商户银行来账查询 GET

	// 服务商-特约商户进件
	v3Apply4SubSubmit               = "/v3/applyment4sub/applyment/"                     // 提交申请单 POST
	v3Apply4SubQueryByBusinessCode  = "/v3/applyment4sub/applyment/business_code/%s"     // business_code 通过业务申请编号查询申请状态 GET
	v3Apply4SubQueryByApplyId       = "/v3/applyment4sub/applyment/applyment_id/%s"      // applyment_id 通过申请单号查询申请状态 GET
	v3Apply4SubModifySettlement     = "/v3/apply4sub/sub_merchants/%s/modify-settlement" // sub_mchid 修改结算账号 POST
	v3Apply4SubQuerySettlement      = "/v3/apply4sub/sub_merchants/%s/settlement"        // sub_mchid 查询结算账户 GET
	v3Apply4SubMerchantsApplication = "/v3/apply4sub/sub_merchants/%s/application/%s"    // sub_mchid、application_no 查询结算账户修改申请状态

	// 电商收付通（商户进件）
	v3EcommerceApply          = "/v3/ecommerce/applyments/"                  // 二级商户进件 POST
	v3EcommerceApplyQueryById = "/v3/ecommerce/applyments/%d"                // applyment_id 通过申请单ID查询申请状态 GET
	v3EcommerceApplyQueryByNo = "/v3/ecommerce/applyments/out-request-no/%s" // out_request_no 通过业务申请编号查询申请状态 GET

	// 电商收付通（分账）
	v3EcommerceProfitShare               = "/v3/ecommerce/profitsharing/orders"            // 请求分账 POST
	v3EcommerceProfitShareQuery          = "/v3/ecommerce/profitsharing/orders"            // 查询分账结果 GET
	v3EcommerceProfitShareReturn         = "/v3/ecommerce/profitsharing/returnorders"      // 请求分账回退 POST
	v3EcommerceProfitShareReturnResult   = "/v3/ecommerce/profitsharing/returnorders"      // 查询分账回退结果 GET
	v3EcommerceProfitShareFinish         = "/v3/ecommerce/profitsharing/finish-order"      // 完结分账 POST
	v3EcommerceProfitShareUnsplitAmount  = "/v3/ecommerce/profitsharing/orders/%s/amounts" // transaction_id 查询订单剩余待分金额 GET
	v3EcommerceProfitShareAddReceiver    = "/v3/ecommerce/profitsharing/receivers/add"     // 添加分账接收方 POST
	v3EcommerceProfitShareDeleteReceiver = "/v3/ecommerce/profitsharing/receivers/delete"  // 删除分账接收方 POST

	// 电商收付通（补差）
	v3EcommerceSubsidies       = "/v3/ecommerce/subsidies/create" // 请求补差 POST
	v3EcommerceSubsidiesReturn = "/v3/ecommerce/subsidies/return" // 请求补差回退 POST
	v3EcommerceSubsidiesCancel = "/v3/ecommerce/subsidies/cancel" // 取消补差 POST

	// 电商收付通（退款）
	v3CommerceRefund              = "/v3/ecommerce/refunds/apply"             // 申请退款 POST
	v3CommerceRefundQueryById     = "/v3/ecommerce/refunds/id/%s"             // refund_id 通过微信支付退款单号查询退款 GET
	v3CommerceRefundQueryByNo     = "/v3/ecommerce/refunds/out-refund-no/%s"  // out_refund_no 通过商户退款单号查询退款 GET
	v3CommerceRefundAdvance       = "/v3/ecommerce/refunds/%s/return-advance" // refund_id 垫付退款回补 POST
	v3CommerceRefundAdvanceResult = "/v3/ecommerce/refunds/%s/return-advance" // refund_id 查询垫付回补结果 GET

	// 银行组件（服务商）
	v3BankSearchBank          = "/v3/capital/capitallhh/banks/search-banks-by-bank-account" // 获取对私银行卡号开户银行 GET
	v3BankSearchPersonalList  = "/v3/capital/capitallhh/banks/personal-banking"             // 查询支持个人业务的银行列表 GET
	v3BankSearchCorporateList = "/v3/capital/capitallhh/banks/corporate-banking"            // 查询支持对公业务的银行列表 GET
	v3BankSearchProvinceList  = "/v3/capital/capitallhh/areas/provinces"                    // 查询省份列表 GET
	v3BankSearchCityList      = "/v3/capital/capitallhh/areas/provinces/%d/cities"          // province_code 查询城市列表 GET
	v3BankSearchBranchList    = "/v3/capital/capitallhh/banks/%s/branches"                  // bank_alias_code 查询支行列表 GET

	// 扣款服务-直连模式（其他相关接口在v2接口中）
	v3EntrustPayNotify = "/v3/papay/contracts/%s/notify" // contract_id 预扣费通知 POST

	// 特约商户进件申请单状态
	ApplyStateEditing       = "APPLYMENT_STATE_EDITTING"        // 编辑中
	ApplyStateAuditing      = "APPLYMENT_STATE_AUDITING"        // 审核中
	ApplyStateRejected      = "APPLYMENT_STATE_REJECTED"        // 已驳回
	ApplyStateToBeConfirmed = "APPLYMENT_STATE_TO_BE_CONFIRMED" // 待账户验证
	ApplyStateSigning       = "APPLYMENT_STATE_SIGNING"         // 开通权限中
	ApplyStateFinished      = "APPLYMENT_STATE_FINISHED"        // 已完成
	ApplyStateCanceled      = "APPLYMENT_STATE_CANCELED"        // 已作废

	// 特约商户结算账号类型
	ApplySettlementAccountTypeBusiness = "ACCOUNT_TYPE_BUSINESS" // 对公银行账户
	ApplySettlementAccountTypePrivate  = "ACCOUNT_TYPE_PRIVATE"  // 经营者个人银行卡

	//结算账号修改模式
	ApplySettlementModifyModeAsync = "MODIFY_MODE_ASYNC" //受理模式

	// 特约商户结算账号汇款验证结果
	ApplySettlementVerifying     = "VERIFYING"      // 系统汇款验证中，商户可发起提现尝试
	ApplySettlementVerifySuccess = "VERIFY_SUCCESS" // 系统成功汇款，该账户可正常发起提现
	ApplySettlementVerifyFail    = "VERIFY_FAIL"    // 系统汇款失败，该账户无法发起提现，请检查修改

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
