package wechat

// Prepay 支付Rsp
type PrepayRsp struct {
	Code     int       `json:"-"`
	SignInfo *SignInfo `json:"-"`
	Response *Prepay   `json:"response,omitempty"`
	Error    string    `json:"-"`
}

// H5 支付Rsp
type H5Rsp struct {
	Code     int       `json:"-"`
	SignInfo *SignInfo `json:"-"`
	Response *H5Url    `json:"response,omitempty"`
	Error    string    `json:"-"`
}

// Native 支付Rsp
type NativeRsp struct {
	Code     int       `json:"-"`
	SignInfo *SignInfo `json:"-"`
	Response *Native   `json:"response,omitempty"`
	Error    string    `json:"-"`
}

// 查询订单 Rsp
type QueryOrderRsp struct {
	Code     int         `json:"-"`
	SignInfo *SignInfo   `json:"-"`
	Response *QueryOrder `json:"response,omitempty"`
	Error    string      `json:"-"`
}

// 关闭订单 Rsp
type CloseOrderRsp struct {
	Code     int       `json:"-"`
	SignInfo *SignInfo `json:"-"`
	Error    string    `json:"-"`
}

// 合单查询订单 Rsp
type CombineQueryOrderRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *CombineQueryOrder `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 服务商查询订单 Rsp
type PartnerQueryOrderRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *PartnerQueryOrder `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// =========================================================分割=========================================================

type Prepay struct {
	PrepayId string `json:"prepay_id"` // 预支付交易会话标识。用于后续接口调用中使用，该值有效期为2小时
}

type Native struct {
	CodeUrl string `json:"code_url"` // 此URL用于生成支付二维码，然后提供给用户扫码支付
}

type H5Url struct {
	H5Url string `json:"h5_url"` // h5_url为拉起微信支付收银台的中间页面，可通过访问该url来拉起微信客户端，完成支付，h5_url的有效期为5分钟
}

type Payer struct {
	Openid    string `json:"openid"`               // 用户在直连商户appid下的唯一标识
	SubOpenid string `json:"sub_openid,omitempty"` // 用户在sub_appid下的标识
}

type QueryOrder struct {
	Appid           string             `json:"appid"`                      // 直连商户申请的公众号或移动应用appid
	Mchid           string             `json:"mchid"`                      // 直连商户的商户号，由微信支付生成并下发
	OutTradeNo      string             `json:"out_trade_no"`               // 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
	TransactionId   string             `json:"transaction_id"`             // 微信支付系统生成的订单号
	TradeType       string             `json:"trade_type"`                 // 交易类型，枚举值：JSAPI：公众号支付, NATIVE：扫码支付, APP：APP支付, MICROPAY：付款码支付, MWEB：H5支付, FACEPAY：刷脸支付
	TradeState      string             `json:"trade_state"`                // 交易状态，枚举值：SUCCESS：支付成功, REFUND：转入退款, NOTPAY：未支付, CLOSED：已关闭, REVOKED：已撤销（付款码支付）, USERPAYING：用户支付中（付款码支付）, PAYERROR：支付失败(其他原因，如银行返回失败)
	TradeStateDesc  string             `json:"trade_state_desc"`           // 交易状态描述
	BankType        string             `json:"bank_type,omitempty"`        // 银行类型，采用字符串类型的银行标识
	Attach          string             `json:"attach"`                     // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用
	SuccessTime     string             `json:"success_time,omitempty"`     // 支付完成时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒
	Payer           *Payer             `json:"payer"`                      // 支付者信息
	Amount          *Amount            `json:"amount,omitempty"`           // 订单金额信息，当支付成功时返回该字段
	SceneInfo       *SceneInfo         `json:"scene_info,omitempty"`       // 支付场景描述
	PromotionDetail []*PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能，享受优惠时返回该字段
}

type SubOrders struct {
	Mchid         string         `json:"mchid"`               // 子单发起方商户号，必须与发起方Appid有绑定关系
	TradeType     string         `json:"trade_type"`          // 交易类型，枚举值：NATIVE：扫码支付，JSAPI：公众号支付，APP：APP支付，MWEB：H5支付
	TradeState    string         `json:"trade_state"`         // 交易状态，枚举值：SUCCESS：支付成功, REFUND：转入退款, NOTPAY：未支付, CLOSED：已关闭, REVOKED：已撤销（付款码支付）, USERPAYING：用户支付中（付款码支付）, PAYERROR：支付失败(其他原因，如银行返回失败)
	BankType      string         `json:"bank_type,omitempty"` // 银行类型，采用字符串类型的银行标识
	Attach        string         `json:"attach"`              // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用
	SuccessTime   string         `json:"success_time"`        // 支付完成时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒
	TransactionId string         `json:"transaction_id"`      // 微信支付系统生成的订单号
	OutTradeNo    string         `json:"out_trade_no"`        // 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
	SubMchid      string         `json:"sub_mchid"`           // 二级商户商户号，由微信支付生成并下发。服务商子商户的商户号，被合单方。直连商户不用传二级商户号
	Amount        *CombineAmount `json:"amount"`              // 订单金额信息，当支付成功时返回该字段
}

type CombineAmount struct {
	TotalAmount   int    `json:"total_amount,omitempty"`   // 订单总金额，单位为分
	Currency      string `json:"currency,omitempty"`       // 标价币种：符合ISO 4217标准的三位字母代码，人民币：CNY
	PayerAmount   int    `json:"payer_amount"`             // 订单现金支付金额
	PayerCurrency string `json:"payer_currency,omitempty"` // 现金支付币种：符合ISO 4217标准的三位字母代码，默认人民币：CNY
}

type CombineQueryOrder struct {
	CombineAppid      string       `json:"combine_appid"`        // 合单发起方的appid
	CombineMchid      string       `json:"combine_mchid"`        // 合单发起方商户号
	CombineOutTradeNo string       `json:"combine_out_trade_no"` // 合单支付总订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一
	SceneInfo         *SceneInfo   `json:"scene_info,omitempty"` // 支付场景描述
	SubOrders         []*SubOrders `json:"sub_orders"`           // 最多支持子单条数：50
	CombinePayerInfo  *Payer       `json:"combine_payer_info"`   // 支付者信息
}

type PartnerQueryOrder struct {
	SpAppid         string             `json:"sp_appid"`                   // 服务商申请的公众号或移动应用appid。
	SpMchid         string             `json:"sp_mchid"`                   // 服务商户号，由微信支付生成并下发
	SubAppid        string             `json:"sub_appid"`                  // 子商户申请的公众号或移动应用appid。如果返回sub_appid，那么sub_openid一定会返回。
	SubMchid        string             `json:"sub_mchid"`                  // 子商户的商户号，有微信支付生成并下发。
	OutTradeNo      string             `json:"out_trade_no"`               // 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
	TransactionId   string             `json:"transaction_id"`             // 微信支付系统生成的订单号
	TradeType       string             `json:"trade_type"`                 // 交易类型，枚举值：JSAPI：公众号支付, NATIVE：扫码支付, APP：APP支付, MICROPAY：付款码支付, MWEB：H5支付, FACEPAY：刷脸支付
	TradeState      string             `json:"trade_state"`                // 交易状态，枚举值：SUCCESS：支付成功, REFUND：转入退款, NOTPAY：未支付, CLOSED：已关闭, REVOKED：已撤销（付款码支付）, USERPAYING：用户支付中（付款码支付）, PAYERROR：支付失败(其他原因，如银行返回失败)
	TradeStateDesc  string             `json:"trade_state_desc"`           // 交易状态描述
	BankType        string             `json:"bank_type,omitempty"`        // 银行类型，采用字符串类型的银行标识
	Attach          string             `json:"attach"`                     // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用
	SuccessTime     string             `json:"success_time,omitempty"`     // 支付完成时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒
	Payer           *PartnerPayer      `json:"payer"`                      // 支付者信息
	Amount          *Amount            `json:"amount,omitempty"`           // 订单金额信息，当支付成功时返回该字段
	SceneInfo       *SceneInfo         `json:"scene_info,omitempty"`       // 支付场景描述
	PromotionDetail []*PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能，享受优惠时返回该字段
}

type PartnerPayer struct {
	SpOpenid  string `json:"sp_openid"`  // 用户在服务商appid下的唯一标识。
	SubOpenid string `json:"sub_openid"` // 用户在子商户appid下的唯一标识。 如果返回sub_appid，那么sub_openid一定会返回。
}
