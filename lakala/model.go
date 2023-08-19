package lakala

type ErrorCode struct {
	ReturnCode string `json:"return_code,omitempty"` // 执行结果
	ResultCode string `json:"result_code,omitempty"` // SUCCESS表示成功
	ReturnMsg  string `json:"return_msg,omitempty"`  // 返回错误信息
}

type ExchangeRateRsp struct {
	ErrorCode
	WechatRate       float64 `json:"wechat_rate,omitempty"`        // 微信汇率
	AlipayRetailRate float64 `json:"alipay_retail_rate,omitempty"` // 支付宝线下渠道汇率
	AlipayOnlineRate float64 `json:"alipay_online_rate,omitempty"` // 支付宝线上渠道汇率
}

type EncryptRsp struct {
	ErrorCode
	PublicKey   string `json:"public_key,omitempty"`   // 公钥信息
	EncryptType string `json:"encrypt_type,omitempty"` // 加密类型(目前仅支持RSA)
	KeyId       string `json:"key_id,omitempty"`       // 密钥ID
	Expire      int64  `json:"expire,omitempty"`       // 过期时间(13位毫秒时间戳)
}

type OrdersRsp struct {
	ErrorCode
	OrderId             string  `json:"order_id,omitempty"`              // Lakala订单ID
	PartnerOrderId      string  `json:"partner_order_id,omitempty"`      // 商户订单ID
	ChannelOrderId      string  `json:"channel_order_id,omitempty"`      // 渠道方(微信、支付宝等)流水号ID，只有支付成功时返回
	TotalFee            int     `json:"total_fee,omitempty"`             // 订单金额，单位是货币最小面值单位
	RealFee             int     `json:"real_fee,omitempty"`              // 实际支付金额，单位是货币最小面值单位(目前等于订单金额，为卡券预留)
	Rate                float64 `json:"rate,omitempty"`                  // 交易时使用的汇率，1JPY=?CNY，只有支付成功时返回，若渠道未提供汇率，会返回1.0
	CustomerId          string  `json:"customer_id,omitempty"`           // 客户ID，若渠道未提供则不存在
	PayTime             string  `json:"pay_time,omitempty"`              // 支付时间（yyyy-MM-dd HH:mm:ss，GMT+9），只有支付成功时返回
	CreateTime          string  `json:"create_time,omitempty"`           // 订单创建时间（最新订单为准）（yyyy-MM-dd HH:mm:ss，GMT+9）
	Currency            string  `json:"currency,omitempty"`              // 币种，通常为JPY
	Channel             string  `json:"channel,omitempty"`               // 支付渠道 Alipay|支付宝、Wechat|微信、AlipayOnline|支付宝线上、Alipay+、日系QR*、UnionPay|银联，若订单未确定渠道则不返回
	PayType             string  `json:"pay_type,omitempty"`              // 支付钱包类型（日系QR* /Alipay+存在）
	OrderDescription    string  `json:"order_description,omitempty"`     // 备注
	ChannelErrorCode    string  `json:"channel_error_code,omitempty"`    // 渠道错误码，订单提交失败、支付失败或已关闭时存在
	ChannelErrorMessage string  `json:"channel_error_message,omitempty"` // 渠道错误描述，订单提交失败、支付失败或已关闭时存在
	MerchantAppId       string  `json:"merchant_app_id,omitempty"`       // 小程序/开放平台APPID，微信小程序或微信SDK订单且支付完成时存在
	MerchantCustomerId  string  `json:"merchant_customer_id,omitempty"`  // 小程序关联消费者openid，微信小程序或微信SDK订单且支付完成时存在
}

type RefundRsp struct {
	ErrorCode
	RefundId        string `json:"refund_id,omitempty"`         // Lakala退款单号
	PartnerRefundId string `json:"partner_refund_id,omitempty"` // 商户提交的退款单号
	Amount          int    `json:"amount,omitempty"`            // 退款金额，单位是货币最小单位
	Currency        string `json:"currency,omitempty"`          // 币种，通常为JPY
}

type OrderListRsp struct {
	ErrorCode
	Data       []*OrderData `json:"data,omitempty"`       // 订单列表
	Pagination Pagination   `json:"pagination,omitempty"` // 分页信息
	Analysis   Analysis     `json:"analysis,omitempty"`   // 分析信息
}

type OrderData struct {
	OrderId        string `json:"order_id,omitempty"`         // Lakala订单ID
	PartnerOrderId string `json:"partner_order_id,omitempty"` // 商户订单ID
	TotalFee       int    `json:"total_fee,omitempty"`        // 订单金额，单位是货币最小面值单位
	RealFee        int    `json:"real_fee,omitempty"`         // 实际支付金额，单位是货币最小面值单位(目前等于订单金额，为卡券预留)
	Channel        string `json:"channel,omitempty"`          // 支付渠道 Alipay|支付宝、Wechat|微信、AlipayOnline|支付宝线上、Alipay+、日系QR*、UnionPay|银联，若订单未确定渠道则不返回
	Currency       string `json:"currency,omitempty"`         // 币种，通常为JPY
	PayTime        string `json:"pay_time,omitempty"`         // 支付时间（yyyy-MM-dd HH:mm:ss，GMT+9），只有支付成功时返回
	CreateTime     string `json:"create_time,omitempty"`      // 订单创建时间（最新订单为准）（yyyy-MM-dd HH:mm:ss，GMT+9）
	Status         string `json:"status,omitempty"`           // 订单状态
	OrderBody      string `json:"order_body,omitempty"`       // 订单标题
	Gateway        string `json:"gateway,omitempty"`          // 订单模式
	PartnerCode    string `json:"partner_code,omitempty"`     // 商户编码
	PartnerName    string `json:"partner_name,omitempty"`     // 商户名称
	RefundFee      string `json:"refund_fee,omitempty"`       // 当前订单退款金额，货币最小单位
}

type Pagination struct {
	Page       int `json:"page,omitempty"`       // 当前页码
	Limit      int `json:"limit,omitempty"`      // 每页条数
	TotalCount int `json:"totalCount,omitempty"` // 总条数
	TotalPages int `json:"totalPages,omitempty"` // 总页数
}

type Analysis struct {
	OrderCount int `json:"order_count,omitempty"` // 支付成功订单数（包含有退款订单）
	TotalFee   int `json:"total_fee,omitempty"`   // 成交订单总额，货币最小单位
	RealFee    int `json:"real_fee,omitempty"`    // 支付总额，货币最小单位
}

type TransactionListRsp struct {
	ErrorCode
	TransactionCount int            `json:"transaction_count,omitempty"` // 流水总条目
	OrderCount       int            `json:"order_count,omitempty"`       // 付款单数
	RefundCount      int            `json:"refund_count,omitempty"`      // 退款单数
	Transactions     []*Transaction `json:"transactions,omitempty"`      // 流水列表
}

type Transaction struct {
	TransactionTime       string  `json:"transaction_time,omitempty"`        // 交易时间，格式yyyyMMddHHmmss，GMT+9
	OrderId               string  `json:"order_id,omitempty"`                // Lakala订单ID
	PartnerOrderId        string  `json:"partner_order_id,omitempty"`        // 商户订单ID
	ChannelOrderId        string  `json:"channel_order_id,omitempty"`        // 渠道方(微信、支付宝等)流水号ID，只有支付成功时返回
	RefundId              string  `json:"refund_id,omitempty"`               // Lakala退款单号
	PartnerRefundId       string  `json:"partner_refund_id,omitempty"`       // 商户提交的退款单号
	Gateway               string  `json:"gateway,omitempty"`                 // 下单接口
	Channel               string  `json:"channel,omitempty"`                 // 支付渠道 Alipay|支付宝、Wechat|微信、AlipayOnline|支付宝线上、Alipay+、日系QR*、UnionPay|银联，若订单未确定渠道则不返回
	Type                  string  `json:"type,omitempty"`                    // 流水类型
	Currency              string  `json:"currency,omitempty"`                // 币种，通常为JPY
	TotalAmount           int     `json:"total_amount,omitempty"`            // 订单总金额，单位是货币最小单位
	InputAmount           int     `json:"input_amount,omitempty"`            // 订单输入金额，单位是货币最小单位
	CustomerPaymentAmount int     `json:"customer_payment_amount,omitempty"` // 客户实际支付金额，单位是货币最小单位
	SettleAmount          int     `json:"settle_amount,omitempty"`           // 结算金额，币种为JPY，单位是货币最小单位
	SurchargeRate         string  `json:"surcharge_rate,omitempty"`          // 手续费费率(x%)
	Surcharge             int     `json:"surcharge,omitempty"`               // 手续费金额，单位是JPY分
	TransferAmount        int     `json:"transfer_amount,omitempty"`         // 打款金额，单位是JPY分
	ExchangeRate          float64 `json:"exchange_rate,omitempty"`           // 使用汇率
	Remark                string  `json:"remark,omitempty"`                  // 备注
}

type SettlementsRsp struct {
	ErrorCode
	SettleFrom       string         `json:"settle_from,omitempty"`
	SettleTo         string         `json:"settle_to,omitempty"`
	SettleDays       string         `json:"settle_days,omitempty"`
	TransactionCount int            `json:"transaction_count,omitempty"` // 流水总条目
	OrderCount       int            `json:"order_count,omitempty"`       // 付款单数
	RefundCount      int            `json:"refund_count,omitempty"`      // 退款单数
	TotalCredit      int            `json:"total_credit,omitempty"`
	TotalDebit       int            `json:"total_debit,omitempty"`
	TotalSurcharge   int            `json:"total_surcharge,omitempty"`
	TotalTransfer    int            `json:"total_transfer,omitempty"`
	Transactions     []*Transaction `json:"transactions,omitempty"` // 流水列表
}

type ConsultPaymentRsp struct {
	ErrorCode
	SubChannels []*SubChannel `json:"sub_channels,omitempty"`
}

type SubChannel struct {
	SubChannel     string `json:"sub_channel,omitempty"`
	SubChannelName string `json:"sub_channel_name,omitempty"`
	Logo           string `json:"logo,omitempty"`
}

type GetCouponRsp struct {
	Res           string `json:"res,omitempty"`
	AvailableTime string `json:"available_time,omitempty"`
	VoucherNum    string `json:"voucher_num,omitempty"`
	EndTime       string `json:"end_time,omitempty"`
	Discount      string `json:"discount,omitempty"`
	Title         string `json:"title,omitempty"`
	State         string `json:"state,omitempty"`
	Type          string `json:"type,omitempty"`
	IsValidNow    string `json:"is_valid_now,omitempty"`
	AvailableWeek string `json:"available_week,omitempty"`
	BeginAmount   string `json:"begin_amount,omitempty"`
}

type PaymentRsp struct {
	ErrorCode
	Channel        string `json:"channel,omitempty"`          // 支付渠道 允许值: Alipay, Alipay+, Wechat, UnionPay
	PartnerCode    string `json:"partner_code,omitempty"`     // 商户编码
	FullName       string `json:"full_name,omitempty"`        // 商户注册全名
	PartnerName    string `json:"partner_name,omitempty"`     // 商户名称
	OrderId        string `json:"order_id,omitempty"`         // Lakala订单ID
	PartnerOrderId string `json:"partner_order_id,omitempty"` // 商户订单ID
	CodeUrl        string `json:"code_url,omitempty"`         // 支付码链接，商户可以据此自行生成二维码
	QrcodeImg      string `json:"qrcode_img,omitempty"`       // Base64封装的二维码图片，可直接作为img的src属性
	PayUrl         string `json:"pay_url,omitempty"`          // 跳转URL
	SdkParams      string `json:"sdk_params,omitempty"`       // NativeJsapi、小程序支付所需参数(Json字符串)
}

type RetailPayRsp struct {
	ErrorCode
	OrderId             string `json:"order_id,omitempty"`              // Lakala订单ID
	PartnerOrderId      string `json:"partner_order_id,omitempty"`      // 商户订单ID
	TotalFee            int    `json:"total_fee,omitempty"`             // 订单金额，单位是货币最小面值单位
	RealFee             int    `json:"real_fee,omitempty"`              // 实际支付金额，单位是货币最小面值单位(目前等于订单金额，为卡券预留)
	PayTime             string `json:"pay_time,omitempty"`              // 支付时间（yyyy-MM-dd HH:mm:ss，GMT+9），只有支付成功时返回
	CreateTime          string `json:"create_time,omitempty"`           // 订单创建时间（最新订单为准）（yyyy-MM-dd HH:mm:ss，GMT+9）
	Currency            string `json:"currency,omitempty"`              // 币种 (JPY/CNY)
	Channel             string `json:"channel,omitempty"`               // 支付渠道 Alipay|支付宝、Wechat|微信、AlipayOnline|支付宝线上、Alipay+、日系QR*、UnionPay|银联，若订单未确定渠道则不返回
	OrderDescription    string `json:"order_description,omitempty"`     // 备注
	ChannelErrorCode    string `json:"channel_error_code,omitempty"`    // 渠道错误码，订单提交失败、支付失败或已关闭时存在
	ChannelErrorMessage string `json:"channel_error_message,omitempty"` // 渠道错误描述，订单提交失败、支付失败或已关闭时存在
}

type ReportRsp struct {
	ErrorCode
	Reports []*ReportSub `json:"reports,omitempty"`
}

type ReportSub struct {
	ReportId                string `json:"report_id,omitempty"` // Lakala海关单号
	PartnerReportId         string `json:"partner_report_id,omitempty"`
	Status                  string `json:"status,omitempty"`
	Channel                 string `json:"channel,omitempty"`
	ChannelReportId         string `json:"channel_report_id,omitempty"`
	PartnerSubReportId      string `json:"partner_sub_report_id,omitempty"`
	ChannelSubReportId      string `json:"channel_sub_report_id,omitempty"`
	Customs                 string `json:"customs,omitempty"`
	MchCustomsNo            string `json:"mch_customs_no,omitempty"`
	MchCustomsName          string `json:"mch_customs_name,omitempty"`
	OrderId                 string `json:"order_id,omitempty"`
	TransactionId           string `json:"transaction_id,omitempty"`
	OrderCurrency           string `json:"order_currency,omitempty"`
	OrderAmount             int    `json:"order_amount,omitempty"`
	SubOrderFee             int    `json:"sub_order_fee,omitempty"`
	SubProductFee           int    `json:"sub_product_fee,omitempty"`
	SubTransportFee         int    `json:"sub_transport_fee,omitempty"`
	CreationDate            string `json:"creation_date,omitempty"`
	LastUpdateDate          string `json:"last_update_date,omitempty"`
	VerifyDepartment        string `json:"verify_department,omitempty"`
	VerifyDepartmentTradeId string `json:"verify_department_trade_id,omitempty"`
	ErrorCode               string `json:"error_code,omitempty"`
	ErrorMsg                string `json:"error_msg,omitempty"`
}

// CommonApi - 付款通知
type NotifyRequest struct {
	Time           string  `json:"time,omitempty"`             // UTC时间戳
	NonceStr       string  `json:"nonce_str,omitempty"`        // 随机字符串
	Sign           string  `json:"sign,omitempty"`             // 签名
	PartnerOrderId string  `json:"partner_order_id,omitempty"` // 商户订单ID
	OrderId        string  `json:"order_id,omitempty"`         // Lakala订单ID
	ChannelOrderId string  `json:"channel_order_id,omitempty"` // 渠道方(微信、支付宝等)流水号ID
	TotalFee       int     `json:"total_fee,omitempty"`        // 订单金额，单位是货币最小面值单位
	RealFee        int     `json:"real_fee,omitempty"`         // 实际支付金额，单位是货币最小面值单位(目前等于订单金额，为卡券预留)
	Rate           float64 `json:"rate,omitempty"`             // 交易时使用的汇率，1JPY=?CNY，只有支付成功时返回，若渠道未提供汇率，会返回1.0
	CustomerId     string  `json:"customer_id,omitempty"`      // 客户ID
	Currency       string  `json:"currency,omitempty"`         // 币种 (JPY/CNY)
	Channel        string  `json:"channel,omitempty"`          // 支付渠道 Alipay|支付宝、Wechat|微信、AlipayOnline|支付宝线上、Alipay+、日系QR*、UnionPay|银联，若订单未确定渠道则不返回
	CreateTime     string  `json:"create_time,omitempty"`      // 订单创建时间（最新订单为准）（yyyy-MM-dd HH:mm:ss，GMT+9）
	PayTime        string  `json:"pay_time,omitempty"`         // 支付时间（yyyy-MM-dd HH:mm:ss，GMT+9），只有支付成功时返回
	System         string  `json:"system,omitempty"`
	PaymentId      string  `json:"payment_id,omitempty"`
	PayType        string  `json:"pay_type,omitempty"` // 支付钱包类型（日系QR* /Alipay+存在）
}
