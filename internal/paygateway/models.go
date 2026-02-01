package paygateway

type Channel string

const (
	ChannelWechatV3 Channel = "WECHAT_V3"
	ChannelAlipay   Channel = "ALIPAY"
)

type CreatePaymentRequest struct {
	TenantID   string  `json:"tenantId"`
	MerchantID string  `json:"merchantId"`
	Channel    Channel `json:"channel"`
	Scene      string  `json:"scene"`

	OutTradeNo string `json:"outTradeNo"`
	BizOrderNo string `json:"bizOrderNo,omitempty"`

	Currency string `json:"currency"`
	Amount   int64  `json:"amount"`

	Subject     string `json:"subject"`
	Description string `json:"description,omitempty"`
	ExpireAt    string `json:"expireAt,omitempty"`

	// WeChat JSAPI/MINIAPP
	OpenID string `json:"openid,omitempty"`

	Ext map[string]string `json:"ext,omitempty"`
}

type CreatePaymentResponse struct {
	Code       string      `json:"code"`
	Message    string      `json:"message,omitempty"`
	OutTradeNo string      `json:"outTradeNo,omitempty"`
	Status     string      `json:"status,omitempty"`
	PayData    interface{} `json:"payData,omitempty"`
}

type ClosePaymentResponse struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
}

type QueryPaymentResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type CreateRefundRequest struct {
	TenantID   string  `json:"tenantId"`
	MerchantID string  `json:"merchantId"`
	Channel    Channel `json:"channel"`

	OutTradeNo  string `json:"outTradeNo"`
	OutRefundNo string `json:"outRefundNo"`

	Currency     string `json:"currency"`
	TotalAmount  int64  `json:"totalAmount,omitempty"`
	RefundAmount int64  `json:"refundAmount"`
	Reason       string `json:"reason,omitempty"`
}

type CreateRefundResponse struct {
	Code        string      `json:"code"`
	Message     string      `json:"message,omitempty"`
	OutRefundNo string      `json:"outRefundNo,omitempty"`
	Status      string      `json:"status,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

type QueryRefundResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type CompensationQueryPaymentsRequest struct {
	TenantID    string   `json:"tenantId"`
	MerchantID  string   `json:"merchantId"`
	Channel     Channel  `json:"channel"`
	OutTradeNos []string `json:"outTradeNos"`
}

type CompensationQueryPaymentsResponse struct {
	Code    string                    `json:"code"`
	Message string                    `json:"message,omitempty"`
	Items   []CompensationPaymentItem `json:"items,omitempty"`
}

type CompensationPaymentItem struct {
	OutTradeNo string      `json:"outTradeNo"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
}
