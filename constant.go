package gopay

const (
	null       string = ""
	TimeLayout string = "2006-01-02 15:04:05"
	DateLayout string = "2006-01-02"

	//URL
	wx_base_url        = "https://api.mch.weixin.qq.com/"
	wx_sanbox_base_url = "https://api.mch.weixin.qq.com/sandboxnew/"

	//正式
	wxURL_Micropay          = wx_base_url + "pay/micropay"                    //提交付款码支付
	wxURL_UnifiedOrder      = wx_base_url + "pay/unifiedorder"                //统一下单
	wxURL_OrderQuery        = wx_base_url + "pay/orderquery"                  //查询订单
	wxURL_CloseOrder        = wx_base_url + "pay/closeorder"                  //关闭订单
	wxURL_Refund            = wx_base_url + "secapi/pay/refund"               //申请退款
	wxURL_Reverse           = wx_base_url + "secapi/pay/reverse"              //撤销订单
	wxURL_RefundQuery       = wx_base_url + "pay/refundquery"                 //查询退款
	wxURL_DownloadBill      = wx_base_url + "pay/downloadbill"                //下载对账单
	wxURL_DownloadFundFlow  = wx_base_url + "pay/downloadfundflow"            //下载资金账单
	wxURL_BatchQueryComment = wx_base_url + "billcommentsp/batchquerycomment" //拉取订单评价数据

	//SanBox
	wxURL_SanBox_GetSignKey        = wx_sanbox_base_url + "pay/getsignkey"
	wxURL_SanBox_Micropay          = wx_sanbox_base_url + "pay/micropay"
	wxURL_SanBox_UnifiedOrder      = wx_sanbox_base_url + "pay/unifiedorder"
	wxURL_SanBox_OrderQuery        = wx_sanbox_base_url + "pay/orderquery"
	wxURL_SanBox_CloseOrder        = wx_sanbox_base_url + "pay/closeorder"
	wxURL_SanBox_Refund            = wx_sanbox_base_url + "pay/refund"
	wxURL_SanBox_Reverse           = wx_sanbox_base_url + "pay/reverse"
	wxURL_SanBox_RefundQuery       = wx_sanbox_base_url + "pay/refundquery"
	wxURL_SanBox_DownloadBill      = wx_sanbox_base_url + "pay/downloadbill"
	wxURL_SanBox_DownloadFundFlow  = wx_sanbox_base_url + "pay/downloadfundflow"
	wxURL_SanBox_BatchQueryComment = wx_sanbox_base_url + "billcommentsp/batchquerycomment"

	//支付类型
	TradeType_Mini   = "JSAPI"
	TradeType_JsApi  = "JSAPI"
	TradeType_App    = "APP"
	TradeType_H5     = "MWEB"
	TradeType_Native = "NATIVE"

	//签名方式
	SignType_MD5         = "MD5"
	SignType_HMAC_SHA256 = "HMAC-SHA256"

	//Debug数据
	//
	//===========================================================================================
	//
	zfb_base_url_2        = "https://openapi.alipay.com/gateway.do"
	zfb_sanbox_base_url_2 = "https://openapi.alipaydev.com/gateway.do"
	zfb_base_url          = "https://openapi.alipay.com/gateway.do?charset=utf-8"
	zfb_sanbox_base_url   = "https://openapi.alipaydev.com/gateway.do?charset=utf-8"
)
