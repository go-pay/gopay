package alipay

import "net/http"

const (
	Success = http.StatusOK

	MethodGet           = "GET"
	MethodPost          = "POST"
	MethodPut           = "PUT"
	MethodDelete        = "DELETE"
	MethodPatch         = "PATCH"
	HeaderAuthorization = "Authorization"
	HeaderRequestID     = "alipay-request-id"
	HeaderSdkVersion    = "alipay-sdk-version"
	HeaderAppAuthToken  = "alipay-app-auth-token"
	HeaderTimestamp     = "alipay-timestamp"
	HeaderNonce         = "alipay-nonce"
	HeaderSignature     = "alipay-signature"

	SignTypeRSA = "ALIPAY-SHA256withRSA"

	v3BaseUrlCh      = "https://openapi.alipay.com"               // 正式环境
	v3SandboxBaseUrl = "https://openapi-sandbox.dl.alipaydev.com" // 沙箱环境

)

// 支付
const (
	// 支付-交易
	v3TradePay                 = "/v3/alipay/trade/pay"                               // 统一收单交易支付接口
	v3TradeQuery               = "/v3/alipay/trade/query"                             // 统一收单交易查询接口
	v3TradeRefund              = "/v3/alipay/trade/refund"                            // 统一收单交易退款接口
	v3TradeFastPayRefundQuery  = "/v3/alipay/trade/fastpay/refund/query"              // 统一收单交易退款查询接口
	v3TradeCancel              = "/v3/alipay/trade/cancel"                            // 统一收单交易撤销接口
	v3TradeClose               = "/v3/alipay/trade/close"                             // 统一收单交易关闭接口
	v3DataBillDownloadUrlQuery = "/v3/alipay/data/dataservice/bill/downloadurl/query" // 查询对账单下载地址
	v3TradePrecreate           = "/v3/alipay/trade/precreate"                         // 统一收单线下交易预创建
	v3TradeCreate              = "/v3/alipay/trade/create"                            // 统一收单交易创建接口
	v3TradeOrderInfoSync       = "/v3/alipay/trade/orderinfo/sync"                    // 支付宝订单信息同步接口

	// 支付-商家扣款
	v3UserAgreementQuery      = "/v3/alipay/user/agreement/query"  // 支付宝个人代扣协议查询接口
	v3UserAgreementPageUnSign = "/v3/alipay/user/agreement/unsign" // 支付宝个人代扣协议解约接口

	// 资金-商家分账
	v3TradeRoyaltyRelationBind       = "/v3/alipay/trade/royalty/relation/bind"       // 分账关系绑定
	v3TradeRoyaltyRelationUnbind     = "/v3/alipay/trade/royalty/relation/unbind"     // 分账关系解绑
	v3TradeRoyaltyRelationBatchQuery = "/v3/alipay/trade/royalty/relation/batchquery" // 分账关系查询
	v3TradeRoyaltyRateQuery          = "/v3/alipay/trade/royalty/rate/query"          // 分账比例查询
	v3TradeOrderSettle               = "/v3/alipay/trade/order/settle"                // 统一收单交易结算接口
	v3TradeOrderSettleQuery          = "/v3/alipay/trade/order/settle/query"          // 交易分账查询接口
	v3TradeOrderOnSettleQuery        = "/v3/alipay/trade/order/onsettle/query"        // 分账剩余金额查询

	// 支付-刷脸付
	v3ZolozAuthenticationSmilepayInitialize  = "/v3/zoloz/authentication/smilepay/initialize"   // 刷脸支付初始化
	v3ZolozAuthenticationCustomerFtokenQuery = "/v3/zoloz/authentication/customer/ftoken/query" // 查询刷脸结果信息接口

	// 支付-预授权支付
	v3FundAuthOperationDetailQuery = "/v3/alipay/fund/auth/operation/detail/query" // 资金授权操作查询接口
	v3FundAuthOrderFreeze          = "/v3/alipay/fund/auth/order/freeze"           // 资金授权冻结接口
	v3FundAuthOrderUnfreeze        = "/v3/alipay/fund/auth/order/unfreeze"         // 资金授权解冻接口
	v3FundAuthOrderVoucherCreate   = "/v3/alipay/fund/auth/order/voucher/create"   // 资金授权发码接口

	// 会员
	v3UserCertifyOpenQuery       = "/v3/alipay/user/certify/open/query"                                // 身份认证记录查询
	v3UserCertifyOpenInitialize  = "/v3/alipay/user/certify/open/initialize"                           // 身份认证初始化服务
	v3SystemOauthToken           = "/v3/alipay/system/oauth/token"                                     // 换取授权访问令牌
	v3UserInfoShare              = "/v3/alipay/user/info/share"                                        // 支付宝会员授权信息查询接口
	v3UserAuthRelationshipQuery  = "/v3/alipay/open/auth/userauth/relationship/query"                  // 用户授权关系查询
	v3UserDelOauthDetailQuery    = "/v3/alipay/user/deloauth/detail/query"                             // 查询解除授权明细
	v3FaceVerificationInitialize = "/v3/datadigital/fincloud/generalsaas/face/verification/initialize" // 人脸核身初始化
	v3FaceVerificationQuery      = "/v3/datadigital/fincloud/generalsaas/face/verification/query"      // 人脸核身结果查询
	v3FaceCertifyInitialize      = "/v3/datadigital/fincloud/generalsaas/face/certify/initialize"      // 跳转支付宝人脸核身初始化
	v3FaceCertifyVerify          = "/v3/datadigital/fincloud/generalsaas/face/certify/verify"          // 跳转支付宝人脸核身开始认证
	v3FaceCertifyQuery           = "/v3/datadigital/fincloud/generalsaas/face/certify/query"           // 跳转支付宝人脸核身查询记录
	v3FaceSourceCertify          = "/v3/datadigital/fincloud/generalsaas/face/source/certify"          // 纯服务端人脸核身
	v3FaceCheckInitialize        = "/v3/datadigital/fincloud/generalsaas/face/check/initialize"        // 活体检测初始化
	v3FaceCheckQuery             = "/v3/datadigital/fincloud/generalsaas/face/check/query"             // 活体检测结果查询
	v3IDCardTwoMetaCheck         = "/v3/datadigital/fincloud/generalsaas/twometa/check"                // 身份证二要素核验
	v3BankCardCheck              = "/v3/datadigital/fincloud/generalsaas/bankcard/check"               // 银行卡核验
	v3MobileThreeMetaSimpleCheck = "/v3/datadigital/fincloud/generalsaas/mobilethreemeta/simple/check" // 手机号三要素核验简版
	v3MobileThreeMetaDetailCheck = "/v3/datadigital/fincloud/generalsaas/mobilethreemeta/detail/check" // 手机号三要素核验详版
	v3OcrServerDetect            = "/v3/datadigital/fincloud/generalsaas/ocr/server/detect"            // 服务端OCR
	v3OcrMobileInitialize        = "/v3/datadigital/fincloud/generalsaas/ocr/mobile/initialize"        // App端OCR初始化

	// 资金-商家转账
	v3FundAccountQuery           = "/v3/alipay/fund/account/query"            // GET 支付宝资金账户资产查询接口
	v3FundQuotaQuery             = "/v3/alipay/fund/quota/query"              // GET 转账额度查询接口
	v3DataBillEreceiptApply      = "/v3/alipay/data/bill/ereceipt/apply"      // POST 申请电子回单(incubating)
	v3DataBillEreceiptQuery      = "/v3/alipay/data/bill/ereceipt/query"      // GET 查询电子回单状态(incubating)
	v3FundTransUniTransfer       = "/v3/alipay/fund/trans/uni/transfer"       // POST 单笔转账接口
	v3FundTransCommonQuery       = "/v3/alipay/fund/trans/common/query"       // GET 转账业务单据查询接口
	v3FundTransMultistepTransfer = "/v3/alipay/fund/trans/multistep/transfer" // POST 多步转账创建并支付
	v3FundTransMultistepQuery    = "/v3/alipay/fund/trans/multistep/query"    // POST 多步转账查询接口

	// 公域-经营推广-推广计划
	v3MarketingActivityDeliveryCreate = "/v3/alipay/marketing/delivery"          // POST 创建推广计划
	v3MarketingActivityDeliveryQuery  = "/v3/alipay/marketing/delivery/%s/query" // POST delivery_id 查询推广计划
	v3MarketingActivityDeliveryStop   = "/v3/alipay/marketing/delivery/%s/stop"  // PATCH delivery_id 停止推广计划

	// 公域-经营推广-推广计划
	v3MarketingMaterialImageUpload = "/v3/alipay/marketing/material/image" // POST 营销图片资源上传接口

	// 营销-营销活动送红包
	v3MarketingCampaignCashCreate       = "/v3/alipay/marketing/campaign/cash/create"        // POST 创建现金活动
	v3MarketingCampaignCashTrigger      = "/v3/alipay/marketing/campaign/cash/trigger"       // POST 触发现金红包活动
	v3MarketingCampaignCashStatusModify = "/v3/alipay/marketing/campaign/cash/status/modify" // POST 更改现金活动状态
	v3MarketingCampaignCashListQuery    = "/v3/alipay/marketing/campaign/cash/list/query"    // GET 现金活动列表查询
	v3MarketingCampaignCashDetailQuery  = "/v3/alipay/marketing/campaign/cash/detail/query"  // GET	现金活动详情查询

	// 营销-商家券-活动
	v3MarketingActivityOrderVoucherCreate         = "/v3/alipay/marketing/ordervoucher/activity"                        // POST 创建商家券活动
	v3MarketingActivityOrderVoucherCodeDeposit    = "/v3/alipay/marketing/ordervoucher/activity/%s/voucher/codedeposit" // POST activity_id 同步商家券券码
	v3MarketingActivityOrderVoucherModify         = "/v3/alipay/marketing/ordervoucher/activity/%s"                     // PATCH activity_id 修改商家券活动基本信息
	v3MarketingActivityOrderVoucherStop           = "/v3/alipay/marketing/ordervoucher/activity/%s/stop"                // PATCH activity_id 停止商家券活动
	v3MarketingActivityOrderVoucherAppend         = "/v3/alipay/marketing/ordervoucher/activity/%s/append"              // PATCH activity_id 修改商家券活动发券数量上限
	v3MarketingActivityOrderVoucherUse            = "/v3/alipay/marketing/ordervoucher/activity/%s/voucher/%s/use"      // POST activity_id, voucher_code 同步券核销状态
	v3MarketingActivityOrderVoucherRefund         = "/v3/alipay/marketing/ordervoucher/activity/%s/voucher/%s/refund"   // POST activity_id, voucher_code 取消券核销状态
	v3MarketingActivityConsult                    = "/v3/alipay/marketing/activity/consult"                             // POST 活动领取咨询接口
	v3MarketingActivityOrderVoucherQuery          = "/v3/alipay/marketing/activity/ordervoucher/query"                  // GET 查询商家券活动
	v3MarketingActivityQuery                      = "/v3/alipay/marketing/activity/%s"                                  // GET activity_id 查询活动详情
	v3MarketingActivityOrderVoucherCodeCount      = "/v3/alipay/marketing/ordervoucher/activity/%s/voucher/codecount"   // GET activity_id 统计商家券券码数量
	v3MarketingActivityBatchQuery                 = "/v3/alipay/marketing/activity/batchquery"                          // POST 条件查询活动列表
	v3MarketingActivityQueryUserBatchQueryVoucher = "/v3/alipay/marketing/activity/ordervoucher/user/batchqueryvoucher" // GET 条件查询用户券
	v3MarketingActivityQueryUserQueryVoucher      = "/v3/alipay/marketing/activity/ordervoucher/user/voucher"           // GET 查询用户券详情
	v3MarketingActivityQueryAppBatchQuery         = "/v3/alipay/marketing/activity/%s/app/batchquery"                   // GET activity_id 查询活动可用小程序
	v3MarketingActivityQueryShopBatchQuery        = "/v3/alipay/marketing/activity/%s/shop/batchquery"                  // GET activity_id 查询活动可用门店
	v3MarketingActivityQueryGoodsBatchQuery       = "/v3/alipay/marketing/activity/%s/goods/batchquery"                 // GET activity_id 查询活动适用商品

	// 营销-商家券-蚂蚁店铺
	v3AntMerchantShopCreate                   = "/v3/ant/merchant/shop"                            // POST 蚂蚁店铺创建
	v3AntMerchantShopQuery                    = "/v3/ant/merchant/shop"                            // GET 店铺查询接口
	v3AntMerchantShopModify                   = "/v3/ant/merchant/shop"                            // PATCH 修改蚂蚁店铺
	v3AntMerchantShopClose                    = "/v3/ant/merchant/shop/close"                      // PATCH 蚂蚁店铺关闭
	v3AntMerchantOrderQuery                   = "/v3/ant/merchant/order/{order_id}"                // GET order_id 商户申请单查询
	v3AntMerchantShopPageQuery                = "/v3/ant/merchant/shop/pagequery"                  // GET 店铺分页查询接口
	v3AntMerchantExpandIndirectImageUpload    = "/v3/ant/merchant/indirect/image"                  // POST 图片上传
	v3AntMerchantExpandMccQuery               = "/v3/ant/merchant/mcc/query"                       // GET 商户mcc信息查询
	v3AntMerchantExpandShopReceiptAccountSave = "/v3/ant/merchant/expand/shop/receiptaccount/save" // POST 店铺增加收单账号

	// 营销-商家会员卡
	v3MarketingCardTemplateCreate  = "/v3/alipay/marketing/card/template/create"  // POST 会员卡模板创建
	v3MarketingCardTemplateQuery   = "/v3/alipay/marketing/card/template/query"   // GET 会员卡模板查询接口
	v3MarketingCardTemplateModify  = "/v3/alipay/marketing/card/template/modify"  // POST 会员卡模板修改
	v3MarketingCardFormTemplateSet = "/v3/alipay/marketing/card/formtemplate/set" // POST 会员卡开卡表单模板配置
	v3MarketingCardQuery           = "/v3/alipay/marketing/card/query"            // POST 会员卡查询
	v3MarketingCardUpdate          = "/v3/alipay/marketing/card/update"           // POST 会员卡更新
	v3MarketingCardDelete          = "/v3/alipay/marketing/card/delete"           // DELETE 会员卡删卡
	v3MarketingCardMessageNotify   = "/v3/alipay/marketing/card/message/notify"   // POST 会员卡消息通知
	v3OfflineMaterialImageUpload   = "/v3/alipay/offline/material/image/upload"   // POST 上传门店照片和视频

	// 营销-红包
	v3FundTransRefund = "/v3/alipay/fund/trans/refund" // POST 资金退回

	// 营销-棋盘密云
	v3MerchantQipanCrowdCreate              = "/v3/alipay/merchant/qipan/crowd/create"              // POST 上传创建人群
	v3MerchantQipanCrowdUserAdd             = "/v3/alipay/merchant/qipan/crowduser/add"             // POST 人群中追加用户
	v3MerchantQipanCrowdUserDelete          = "/v3/alipay/merchant/qipan/crowduser/delete"          // POST 人群中删除用户
	v3MarketingQipanTagBaseBatchQuery       = "/v3/alipay/marketing/qipan/tagbase/query"            // GET 棋盘人群圈选标签基本信息查询
	v3MarketingQipanTagQuery                = "/v3/alipay/marketing/qipan/operationnode/query"      // GET 棋盘标签圈选值查询
	v3MarketingQipanCrowdOperationCreate    = "/v3/alipay/marketing/qipan/crowdoperation/create"    // POST 棋盘人群创建
	v3MarketingQipanCrowdBatchQuery         = "/v3/alipay/merchant/qipan/crowd/batchquery"          // POST 查询人群列表
	v3MarketingQipanCrowdQuery              = "/v3/alipay/merchant/qipan/crowd/query"               // GET 查询人群详情
	v3MarketingQipanCrowdModify             = "/v3/alipay/merchant/qipan/crowd/modify"              // POST 修改人群
	v3MarketingQipanBoardQuery              = "/v3/alipay/merchant/qipan/board/query"               // POST 看板分析
	v3MarketingQipanInsightQuery            = "/v3/alipay/merchant/qipan/insight/query"             // POST 画像分析
	v3MarketingQipanBehaviorQuery           = "/v3/alipay/merchant/qipan/behavior/query"            // POST 行为分析
	v3MarketingQipanTrendQuery              = "/v3/alipay/merchant/qipan/trend/query"               // POST 趋势分析
	v3MarketingQipanInsightCityQuery        = "/v3/alipay/merchant/qipan/insightcity/query"         // POST 常住省市查询
	v3MerchantQipanCrowdPoolCreate          = "/v3/alipay/merchant/qipan/crowdpool/create"          // POST 人群池创建
	v3MerchantQipanCrowdSpread              = "/v3/alipay/merchant/qipan/crowd/spread"              // POST 人群扩展接口
	v3MerchantQipanGreyBlackCrowdCreate     = "/v3/alipay/merchant/qipan/greyblackcrowd/create"     // POST 上传创建灰黑产人群
	v3MerchantQipanGreyBlackCrowdUserAdd    = "/v3/alipay/merchant/qipan/greyblackcrowduser/add"    // POST 灰黑产人群中追加用户
	v3MerchantQipanGreyBlackCrowdUserDelete = "/v3/alipay/merchant/qipan/greyblackcrowduser/delete" // POST 灰黑产人群中删除用户

	// 营销-支付券（支付券中有很多v2接口，可能是支付宝还未重构完）
	v3MarketingCampaignOrderVoucherConsult = "/v3/alipay/marketing/campaign/order/voucher/consult" // POST 订单优惠前置咨询

	// 信用-芝麻go
	v3ZmGoPreorderCreate       = "/v3/zhima/credit/pe/zmgo/preorder/create"  // PUT 芝麻GO签约预创单接口
	v3ZmGoCumulateSync         = "/v3/zhima/merchant/zmgo/cumulate/sync"     // POST 商家芝麻GO累计数据回传接口
	v3ZmGoCumulateQuery        = "/v3/zhima/merchant/zmgo/cumulate/query"    // GET 商家芝麻GO累计数据查询接口
	v3ZmGoSettleApply          = "/v3/zhima/credit/pe/zmgo/settle/apply"     // POST 芝麻GO结算申请
	v3ZmGoSettleApplyRefund    = "/v3/zhima/credit/pe/zmgo/settle/refund"    // POST 芝麻GO结算退款
	v3ZmGoAgreementQuery       = "/v3/zhima/credit/pe/zmgo/agreement/query"  // GET 芝麻GO协议查询
	v3ZmGoAgreementQueryUnsign = "/v3/zhima/credit/pe/zmgo/agreement/unsign" // POST 芝麻GO协议解约
	v3ZmGoTemplateCreate       = "/v3/zhima/merchant/zmgo/template/create"   // POST 商家芝麻GO模板创建
	v3ZmGoTemplateQuery        = "/v3/zhima/merchant/zmgo/template/query"    // GET 商家芝麻GO模板查询

	// 广告-支付宝广告投放
	v3DataDataserviceAdConversionUpload               = "/v3/alipay/data/dataservice/ad/conversion/upload"               // POST 转化数据回传
	v3DataDataserviceAdReportdataQuery                = "/v3/alipay/data/dataservice/ad/reportdata/query"                // POST 广告投放数据通用查询
	v3DataDataserviceAdPromotepageBatchquery          = "/v3/alipay/data/dataservice/ad/promotepage/batchquery"          // GET 自建推广页列表批量查询
	v3DataDataserviceAdPromotepageDownload            = "/v3/alipay/data/dataservice/ad/promotepage/download"            // GET 自建推广页留资数据查询
	v3DataDataserviceXlightTaskQuery                  = "/v3/alipay/data/dataservice/xlight/task/query"                  // POST 任务广告完成状态查询接口
	v3DataDataserviceAdConsumehistoryQuery            = "/v3/alipay/data/dataservice/ad/consumehistory/query"            // POST 消费明细查询接口
	v3DataDataserviceProductLandinginfoCreateOrModify = "/v3/alipay/data/dataservice/product/landinginfo/createormodify" // POST 商品落地页信息创建或更新
	v3DataDataserviceProductLandinginfoQuery          = "/v3/alipay/data/dataservice/product/landinginfo/query"          // POST 商品落地页信息查询
	v3DataDataserviceAdAgentreportdataQuery           = "/v3/alipay/data/dataservice/ad/agentreportdata/query"           // POST 广告投放数据代理商查询
)
