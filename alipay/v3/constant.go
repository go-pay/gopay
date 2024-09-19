package alipay

const (
	Success = 0

	MethodGet           = "GET"
	MethodPost          = "POST"
	MethodPut           = "PUT"
	MethodDelete        = "DELETE"
	MethodPATCH         = "PATCH"
	HeaderAuthorization = "Authorization"
	HeaderRequestID     = "alipay-request-id"
	HeaderSdkVersion    = "alipay-sdk-version"
	HeaderAppAuthToken  = "alipay-app-auth-token"

	SignTypeRSA = "ALIPAY-SHA256withRSA"

	v3BaseUrlCh      = "https://openapi.alipay.com"               // 正式环境
	v3SandboxBaseUrl = "https://openapi-sandbox.dl.alipaydev.com" // 沙箱环境

)
