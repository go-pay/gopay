package wecaht

const (
	MethodPost          = "POST"
	MethodGet           = "GET"
	HeaderAuthorization = "Authorization"

	Authorization = "WECHATPAY2-SHA256-RSA2048"

	v3BaseUrlCh = "https://api.mch.weixin.qq.com" // 中国国内

	v3ApiPayApp     = "/v3/pay/transactions/app"
	v3ApiJsapi      = "/v3/pay/transactions/jsapi"
	v3ApiNative     = "/v3/pay/transactions/native"
	v3ApiH5         = "/v3/pay/transactions/h5"
	v3ApiQueryOrder = "/v3/pay/transactions/id/%s"                 // transaction_id
	v3ApiCloseOrder = "/v3/pay/transactions/out-trade-no/%s/close" // out_trade_no

)

type PrepayRsp struct {
	PrepayId string `json:"prepay_id"`
}
