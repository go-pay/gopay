package gopay

const (
	//URL
	wx_base_url        = "https://api.mch.weixin.qq.com/"
	wx_sanbox_base_url = "https://api.mch.weixin.qq.com/sandboxnew/"

	wxURL_unifiedOrder = wx_base_url + "pay/unifiedorder"
	wxURL_orderquery   = wx_base_url + "pay/orderquery"
	wxURL_closeorder   = wx_base_url + "pay/closeorder"

	wxURL_sanbox_getsignkey   = wx_sanbox_base_url + "pay/getsignkey"
	wxURL_sanbox_unifiedOrder = wx_sanbox_base_url + "pay/unifiedorder"
	wxURL_sanbox_orderquery   = wx_sanbox_base_url + "pay/orderquery"
	wxURL_sanbox_closeorder   = wx_sanbox_base_url + "pay/closeorder"

	//支付类型
	WX_PayType_Mini   = "JSAPI"
	WX_PayType_JsApi  = "JSAPI"
	WX_PayType_App    = "APP"
	WX_PayType_H5     = "MWEB"
	WX_PayType_Native = "NATIVE"

	//签名方式
	WX_SignType_MD5         = "MD5"
	WX_SignType_HMAC_SHA256 = "HMAC-SHA256"

	//Debug数据
	secretKey = "GFDS8j98rewnmgl45wHTt980jg543wmg"
	appID     = "wxdaa2ab9ef87b5497"
	mchID     = "1368139502"
	openID    = "o0Df70H2Q0fY8JXh1aFPIRyOBgu8"
)
