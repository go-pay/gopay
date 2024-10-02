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
	// 交易
	v3TradePay                 = "/v3/alipay/trade/pay"                               // 统一收单交易支付接口
	v3TradeQuery               = "/v3/alipay/trade/query"                             // 统一收单交易查询接口
	v3TradeRefund              = "/v3/alipay/trade/refund"                            // 统一收单交易退款接口
	v3TradeFastPayRefundQuery  = "/v3/alipay/trade/fastpay/refund/query"              // 统一收单交易退款查询接口
	v3TradeCancel              = "/v3/alipay/trade/cancel"                            // 统一收单交易撤销接口
	v3TradeClose               = "/v3/alipay/trade/close"                             // 统一收单交易关闭接口
	v3DataBillDownloadUrlQuery = "/v3/alipay/data/dataservice/bill/downloadurl/query" // 查询对账单下载地址
	v3TradePrecreate           = "/v3/alipay/trade/precreate"                         // 统一收单线下交易预创建
	v3TradeCreate              = "/v3/alipay/trade/create"                            // 统一收单交易创建接口

	// 商家扣款
	v3UserAgreementQuery      = "/v3/alipay/user/agreement/query"  // 支付宝个人代扣协议查询接口
	v3UserAgreementPageUnSign = "/v3/alipay/user/agreement/unsign" // 支付宝个人代扣协议解约接口

	// 商家分账
	v3TradeRoyaltyRelationBind       = "/v3/alipay/trade/royalty/relation/bind"       // 分账关系绑定
	v3TradeRoyaltyRelationUnbind     = "/v3/alipay/trade/royalty/relation/unbind"     // 分账关系解绑
	v3TradeRoyaltyRelationBatchQuery = "/v3/alipay/trade/royalty/relation/batchquery" // 分账关系查询
	v3TradeRoyaltyRateQuery          = "/v3/alipay/trade/royalty/rate/query"          // 分账比例查询
	v3TradeOrderSettle               = "/v3/alipay/trade/order/settle"                // 统一收单交易结算接口
	v3TradeOrderSettleQuery          = "/v3/alipay/trade/order/settle/query"          // 交易分账查询接口
	v3TradeOrderOnSettleQuery        = "/v3/alipay/trade/order/onsettle/query"        // 分账剩余金额查询
	v3TradeOrderInfoSync             = "/v3/alipay/trade/orderinfo/sync"              // 支付宝订单信息同步接口

	// 刷脸付
	v3ZolozAuthenticationSmilepayInitialize  = "/v3/zoloz/authentication/smilepay/initialize"   // 刷脸支付初始化
	v3ZolozAuthenticationCustomerFtokenQuery = "/v3/zoloz/authentication/customer/ftoken/query" // 查询刷脸结果信息接口

	// 预授权支付
	v3FundAuthOperationDetailQuery = "/v3/alipay/fund/auth/operation/detail/query" // 资金授权操作查询接口
	v3FundAuthOrderFreeze          = "/v3/alipay/fund/auth/order/freeze"           // 资金授权冻结接口
	v3FundAuthOrderUnfreeze        = "/v3/alipay/fund/auth/order/unfreeze"         // 资金授权解冻接口
	v3FundAuthOrderVoucherCreate   = "/v3/alipay/fund/auth/order/voucher/create"   // 资金授权发码接口

)
