package gopay

const (
	null       string = ""
	TimeLayout string = "2006-01-02 15:04:05"
	DateLayout string = "2006-01-02"
	Version    string = "1.4.8"
	// 微信
	// ===========================================================================================

	// 境外国家地区
	China         Country = 1 // 中国国内
	China2        Country = 2 // 中国国内（冗灾方案）
	SoutheastAsia Country = 3 // 东南亚
	Other         Country = 4 // 其他国家

	// URL
	wxBaseUrlCh  = "https://api.mch.weixin.qq.com/"   //中国国内
	wxBaseUrlCh2 = "https://api2.mch.weixin.qq.com/"  //中国国内
	wxBaseUrlHk  = "https://apihk.mch.weixin.qq.com/" //东南亚
	wxBaseUrlUs  = "https://apius.mch.weixin.qq.com/" //其他

	// 正式
	wxMicropay          = "pay/micropay"                          //提交付款码支付
	wxUnifiedorder      = "pay/unifiedorder"                      //统一下单
	wxOrderquery        = "pay/orderquery"                        //查询订单
	wxCloseorder        = "pay/closeorder"                        //关闭订单
	wxRefund            = "secapi/pay/refund"                     //申请退款
	wxReverse           = "secapi/pay/reverse"                    //撤销订单
	wxRefundquery       = "pay/refundquery"                       //查询退款
	wxDownloadbill      = "pay/downloadbill"                      //下载对账单
	wxDownloadfundflow  = "pay/downloadfundflow"                  //下载资金账单
	wxBatchquerycomment = "billcommentsp/batchquerycomment"       //拉取订单评价数据
	wxTransfers         = "mmpaymkttransfers/promotion/transfers" //企业向微信用户个人付款
	wxEntrustPublic     = "papay/entrustweb"                      //公众号纯签约
	wxEntrustApp        = "papay/preentrustweb"                   //APP纯签约
	wxEntrustH5         = "papay/h5entrustweb"                    //H5纯签约
	wxEntrustQuery      = "papay/querycontract"                   //查询签约关系
	wxEntrustDelete     = "papay/deletecontract"                  //申请解约
	wxEntrustApplyPay   = "pay/pappayapply"                       //申请扣款
	wxEntrustQueryOrder = "pay/paporderquery"                     //查询扣款订单

	// SanBox
	wxSandboxGetsignkey        = "https://api.mch.weixin.qq.com/sandboxnew/pay/getsignkey"
	wxSandboxMicropay          = "sandboxnew/pay/micropay"
	wxSandboxUnifiedorder      = "sandboxnew/pay/unifiedorder"
	wxSandboxOrderquery        = "sandboxnew/pay/orderquery"
	wxSandboxCloseorder        = "sandboxnew/pay/closeorder"
	wxSandboxRefund            = "sandboxnew/pay/refund"
	wxSandboxReverse           = "sandboxnew/pay/reverse"
	wxSandboxRefundquery       = "sandboxnew/pay/refundquery"
	wxSandboxDownloadbill      = "sandboxnew/pay/downloadbill"
	wxSandboxDownloadfundflow  = "sandboxnew/pay/downloadfundflow"
	wxSandboxBatchquerycomment = "sandboxnew/billcommentsp/batchquerycomment"

	// 支付类型
	TradeType_Mini   = "JSAPI"
	TradeType_JsApi  = "JSAPI"
	TradeType_App    = "APP"
	TradeType_H5     = "MWEB"
	TradeType_Native = "NATIVE"

	// 签名方式
	SignType_MD5         = "MD5"
	SignType_HMAC_SHA256 = "HMAC-SHA256"

	// 支付宝
	// ==========================================================================================

	// URL
	zfbBaseUrl            = "https://openapi.alipay.com/gateway.do"
	zfbSandboxBaseUrl     = "https://openapi.alipaydev.com/gateway.do"
	zfbBaseUrlUtf8        = "https://openapi.alipay.com/gateway.do?charset=utf-8"
	zfbSandboxBaseUrlUtf8 = "https://openapi.alipaydev.com/gateway.do?charset=utf-8"
)
