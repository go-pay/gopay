package gopay

const (
	//URL
	wx_base_url        = "https://api.mch.weixin.qq.com/"
	wx_sanbox_base_url = "https://api.mch.weixin.qq.com/sandboxnew/"

	wxURL_UnifiedOrder = wx_base_url + "pay/unifiedorder"
	wxURL_OrderQuery   = wx_base_url + "pay/orderquery"
	wxURL_CloseOrder   = wx_base_url + "pay/closeorder"

	wxURL_SanBox_GetSignKey   = wx_sanbox_base_url + "pay/getsignkey"
	wxURL_SanBox_UnifiedOrder = wx_sanbox_base_url + "pay/unifiedorder"
	wxURL_SanBox_OrderQuery   = wx_sanbox_base_url + "pay/orderquery"
	wxURL_SanBox_CloseOrder   = wx_sanbox_base_url + "pay/closeorder"

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
	secretKey = "GFDS8j98rewnmgl45wHTt980jg543wmg"
	appID     = "wxdaa2ab9ef87b5497"
	mchID     = "1368139502"
	openID    = "o0Df70H2Q0fY8JXh1aFPIRyOBgu8"
)
