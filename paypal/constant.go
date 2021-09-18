package paypal

const (
	Success = 0

	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
	MethodPATCH  = "PATCH"

	HeaderAuthorization = "Authorization" // 请求头Auth

	BaseUrl       = "https://api-m.paypal.com"         // 正式 URL
	SanBoxBaseUrl = "https://api-m.sandbox.paypal.com" // 沙箱 URL
)
