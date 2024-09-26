package alipay

type TradePrecreateRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	// 商户的订单号
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// 当前预下单请求生成的二维码码串，有效时间2小时，可以用二维码生成工具根据该码串值生成对应的二维码
	QrCode string `json:"qr_code,omitempty"`
	// 当前预下单请求生成的吱口令码串，有效时间2小时，可以在支付宝app端访问对应内容
	ShareCode string `json:"share_code,omitempty"`
}
