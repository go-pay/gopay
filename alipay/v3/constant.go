package alipay

import "net/http"

const (
	Success = http.StatusOK

	MethodGet           = "GET"
	MethodPost          = "POST"
	MethodPut           = "PUT"
	MethodDelete        = "DELETE"
	MethodPATCH         = "PATCH"
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
)
