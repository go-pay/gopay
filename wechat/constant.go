package wechat

const (
	// 境外国家地区
	China         Country = 1 // 中国国内
	China2        Country = 2 // 中国国内（冗灾方案）
	SoutheastAsia Country = 3 // 东南亚
	Other         Country = 4 // 其他国家

	// URL
	baseUrlCh  = "https://api.mch.weixin.qq.com"   // 中国国内
	baseUrlCh2 = "https://api2.mch.weixin.qq.com"  // 中国国内
	baseUrlHk  = "https://apihk.mch.weixin.qq.com" // 东南亚
	baseUrlUs  = "https://apius.mch.weixin.qq.com" // 其他

	// 正式
	microPay                        = "/pay/micropay"                                     // 提交付款码支付
	unifiedOrder                    = "/pay/unifiedorder"                                 // 统一下单
	orderQuery                      = "/pay/orderquery"                                   // 查询订单
	closeOrder                      = "/pay/closeorder"                                   // 关闭订单
	refund                          = "/secapi/pay/refund"                                // 申请退款
	reverse                         = "/secapi/pay/reverse"                               // 撤销订单
	refundQuery                     = "/pay/refundquery"                                  // 查询退款
	downloadBill                    = "/pay/downloadbill"                                 // 下载对账单
	downloadFundFlow                = "/pay/downloadfundflow"                             // 下载资金账单
	report                          = "/payitil/report"                                   // 交易保障
	batchQueryComment               = "/billcommentsp/batchquerycomment"                  // 拉取订单评价数据
	transfers                       = "/mmpaymkttransfers/promotion/transfers"            // 企业付款（企业向微信用户个人付款）
	getTransferInfo                 = "/mmpaymkttransfers/gettransferinfo"                // 查询企业付款
	sendCashRed                     = "/mmpaymkttransfers/sendredpack"                    // 发放现金红包
	sendAppletRed                   = "/mmpaymkttransfers/sendminiprogramhb"              // 发放小程序红包
	sendGroupCashRed                = "/mmpaymkttransfers/sendgroupredpack"               // 发放裂变红包
	getRedRecord                    = "/mmpaymkttransfers/gethbinfo"                      // 查询红包记录
	authCodeToOpenid                = "/tools/authcodetoopenid"                           // 授权码查询openid
	entrustPublic                   = "/papay/entrustweb"                                 // 公众号纯签约
	entrustApp                      = "/papay/preentrustweb"                              // APP纯签约
	entrustH5                       = "/papay/h5entrustweb"                               // H5纯签约
	entrustPaying                   = "/pay/contractorder"                                // 支付中签约
	entrustQuery                    = "/papay/querycontract"                              // 查询签约关系
	entrustApplyPay                 = "/pay/pappayapply"                                  // 申请扣款
	entrustDelete                   = "/papay/deletecontract"                             // 申请解约
	profitSharing                   = "/secapi/pay/profitsharing"                         // 请求单次分账
	multiProfitSharing              = "/secapi/pay/multiprofitsharing"                    // 请求多次分账
	profitSharingQuery              = "/pay/profitsharingquery"                           // 查询分账结果
	profitSharingAddReceiver        = "/pay/profitsharingaddreceiver"                     // 添加分账接收方
	profitSharingRemoveReceiver     = "/pay/profitsharingremovereceiver"                  // 删除分账接收方
	profitSharingFinish             = "/secapi/pay/profitsharingfinish"                   // 完结分账
	profitSharingOrderAmountQuery   = "/pay/profitsharingorderamountquery"                // 查询订单待分账金额
	profitSharingMerchantRatioQuery = "/pay/profitsharingmerchantratioquery"              // 查询最大分账比例
	profitSharingReturn             = "/secapi/pay/profitsharingreturn"                   // 分账退回
	profitSharingReturnQuery        = "/pay/profitsharingreturnquery"                     // 分账回退结果查询
	payBank                         = "/mmpaysptrans/pay_bank"                            // 企业付款到银行卡API
	queryBank                       = "/mmpaysptrans/query_bank"                          // 查询企业付款到银行卡API
	getPublicKey                    = "https://fraud.mch.weixin.qq.com/risk/getpublickey" // 获取RSA加密公钥API

	// 海关自助清关
	customsDeclareOrder   = "/cgi-bin/mch/customs/customdeclareorder"        // 订单附加信息提交
	customsDeclareQuery   = "/cgi-bin/mch/customs/customdeclarequery"        // 订单附加信息查询
	customsReDeclareOrder = "/cgi-bin/mch/newcustoms/customdeclareredeclare" // 订单附加信息重推

	// SanBox
	sandboxGetSignKey   = "https://api.mch.weixin.qq.com/sandboxnew/pay/getsignkey"
	sandboxMicroPay     = "/sandboxnew/pay/micropay"
	sandboxUnifiedOrder = "/sandboxnew/pay/unifiedorder"
	sandboxOrderQuery   = "/sandboxnew/pay/orderquery"
	sandboxCloseOrder   = "/sandboxnew/pay/closeorder"
	sandboxRefund       = "/sandboxnew/pay/refund"
	sandboxReverse      = "/sandboxnew/pay/reverse"
	sandboxRefundQuery  = "/sandboxnew/pay/refundquery"
	sandboxDownloadBill = "/sandboxnew/pay/downloadbill"
	sandboxReport       = "/sandboxnew/payitil/report"

	// 支付类型
	TradeType_Mini   = "JSAPI"  // 小程序支付
	TradeType_JsApi  = "JSAPI"  // JSAPI支付
	TradeType_App    = "APP"    // app支付
	TradeType_H5     = "MWEB"   // H5支付
	TradeType_Native = "NATIVE" // Native支付

	// 签名方式
	SignType_MD5         = "MD5"
	SignType_HMAC_SHA256 = "HMAC-SHA256"
)
