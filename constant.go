package gopay

const (
	null       string = ""
	TimeLayout string = "2006-01-02 15:04:05"
	DateLayout string = "2006-01-02"
	Version    string = "1.3.7"
	//微信
	//===========================================================================================

	//境外国家地区
	China         Country = 1 //中国国内
	China2        Country = 2 //中国国内（冗灾方案）
	SoutheastAsia Country = 3 //东南亚
	Other         Country = 4 //其他国家

	//URL
	wx_base_url_ch  = "https://api.mch.weixin.qq.com/"   //中国国内
	wx_base_url_ch2 = "https://api2.mch.weixin.qq.com/"  //中国国内
	wx_base_url_hk  = "https://apihk.mch.weixin.qq.com/" //东南亚
	wx_base_url_us  = "https://apius.mch.weixin.qq.com/" //其他

	//正式
	wx_Micropay          = "pay/micropay"                          //提交付款码支付
	wx_UnifiedOrder      = "pay/unifiedorder"                      //统一下单
	wx_OrderQuery        = "pay/orderquery"                        //查询订单
	wx_CloseOrder        = "pay/closeorder"                        //关闭订单
	wx_Refund            = "secapi/pay/refund"                     //申请退款
	wx_Reverse           = "secapi/pay/reverse"                    //撤销订单
	wx_RefundQuery       = "pay/refundquery"                       //查询退款
	wx_DownloadBill      = "pay/downloadbill"                      //下载对账单
	wx_DownloadFundFlow  = "pay/downloadfundflow"                  //下载资金账单
	wx_BatchQueryComment = "billcommentsp/batchquerycomment"       //拉取订单评价数据
	wx_Transfers         = "mmpaymkttransfers/promotion/transfers" //企业向微信用户个人付款

	//SanBox
	wx_SanBox_GetSignKey        = "https://api.mch.weixin.qq.com/sandboxnew/pay/getsignkey"
	wx_SanBox_Micropay          = "sandboxnew/pay/micropay"
	wx_SanBox_UnifiedOrder      = "sandboxnew/pay/unifiedorder"
	wx_SanBox_OrderQuery        = "sandboxnew/pay/orderquery"
	wx_SanBox_CloseOrder        = "sandboxnew/pay/closeorder"
	wx_SanBox_Refund            = "sandboxnew/pay/refund"
	wx_SanBox_Reverse           = "sandboxnew/pay/reverse"
	wx_SanBox_RefundQuery       = "sandboxnew/pay/refundquery"
	wx_SanBox_DownloadBill      = "sandboxnew/pay/downloadbill"
	wx_SanBox_DownloadFundFlow  = "sandboxnew/pay/downloadfundflow"
	wx_SanBox_BatchQueryComment = "sandboxnew/billcommentsp/batchquerycomment"
	wx_SanBox_Transfers         = "sandboxnew/mmpaymkttransfers/promotion/transfers"

	//支付类型
	TradeType_Mini   = "JSAPI"
	TradeType_JsApi  = "JSAPI"
	TradeType_App    = "APP"
	TradeType_H5     = "MWEB"
	TradeType_Native = "NATIVE"

	//签名方式
	SignType_MD5         = "MD5"
	SignType_HMAC_SHA256 = "HMAC-SHA256"

	//支付宝
	//===========================================================================================

	//URL
	zfb_base_url             = "https://openapi.alipay.com/gateway.do"
	zfb_sanbox_base_url      = "https://openapi.alipaydev.com/gateway.do"
	zfb_base_url_utf8        = "https://openapi.alipay.com/gateway.do?charset=utf-8"
	zfb_sanbox_base_url_utf8 = "https://openapi.alipaydev.com/gateway.do?charset=utf-8"
)
