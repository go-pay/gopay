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

const (
	v3TradePrecreate = "/v3/alipay/trade/precreate"
)
