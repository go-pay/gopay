package alipay

// HTTP 状态码
// 常见的 HTTP 状态码见下表。
// HTTP 状态码				错误类型						一般的解决方案
// 200 - OK					处理成功     					-
// 202 - Accepted 			服务器已接受请求，但尚未处理 	请使用原参数重复请求一遍。
// 204 - No Content			处理成功，无返回Body		        -
// 400 - Bad Request		协议或者参数非法				请检查请求参数是否符合要求。
// 401 - Unauthorized		调用未授权					请检查签名参数和方法是否都符合签名算法要求。
// 403 - Forbidden			无权限调用					请检查产品权限开通情况，可联系产品或商务申请。
// 404 - Not Found			请求的资源不存在				请检查需要查询的 id 或者请求 URL 是否正确。
// 429 - Too Many Requests	请求超过频率限制				请求未受理，请降低频率后重试。
// 500 - Server Error		系统错误						按具体接口的错误指引进行重试。

type ErrResponse struct {
	Code    string    `json:"code"`    // 详细错误码，参考接口描述及公共错误码，商家需要对该错误码处理。
	Message string    `json:"message"` // 错误描述，具体错误原因的文字描述，开发者可参考该描述判断错误原因。
	Details []*Detail `json:"details,omitempty"`
	Links   []*Link   `json:"links,omitempty"`
}

type Detail struct {
	Field       string `json:"field"`
	Value       string `json:"value"`
	Location    string `json:"location"`
	Issue       string `json:"issue"`
	Description string `json:"description"`
}

type Link struct {
	Link string `json:"link"`
	Desc string `json:"desc"`
	Rel  string `json:"rel"`
}
