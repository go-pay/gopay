package wechat

type OrderNoType uint8

type PlatformCertRsp struct {
	Code     int                 `json:"-"`
	SignInfo *SignInfo           `json:"-"`
	Certs    []*PlatformCertItem `json:"certs"`
	Error    string              `json:"-"`
}

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

// 退款 Rsp
type RefundRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *RefundOrderResponse `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 退款查询 Rsp
type RefundQueryRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *RefundQueryResponse `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 交易、资金账单 Rsp
type BillRsp struct {
	Code     int        `json:"-"`
	SignInfo *SignInfo  `json:"-"`
	Response *TradeBill `json:"response,omitempty"`
	Error    string     `json:"-"`
}

// 二级商户资金账单 Rsp
type Level2FundFlowBillRsp struct {
	Code     int           `json:"-"`
	SignInfo *SignInfo     `json:"-"`
	Response *DownloadBill `json:"response,omitempty"`
	Error    string        `json:"-"`
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

// 创建支付分订单 Rsp
type ScoreOrderCreateRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *ScoreOrderCreate `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 查询支付分订单 Rsp
type ScoreOrderQueryRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *ScoreOrderQuery `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// ==================================分割==================================

type JSAPIPayParams struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

type AppPayParams struct {
	Appid     string `json:"appid"`
	Partnerid string `json:"partnerid"`
	Prepayid  string `json:"prepayid"`
	Package   string `json:"package"`
	Noncestr  string `json:"noncestr"`
	Timestamp string `json:"timestamp"`
	PaySign   string `json:"paySign"`
}

type AppletParams struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

// ==================================分割==================================

type SignInfo struct {
	HeaderTimestamp string `json:"Wechatpay-Timestamp"`
	HeaderNonce     string `json:"Wechatpay-Nonce"`
	HeaderSignature string `json:"Wechatpay-Signature"`
	HeaderSerial    string `json:"Wechatpay-Serial"`
	SignBody        string `json:"sign_body"`
}

type PlatformCertItem struct {
	EffectiveTime string `json:"effective_time"`
	ExpireTime    string `json:"expire_time"`
	PublicKey     string `json:"public_key"`
	SerialNo      string `json:"serial_no"`
}

type PlatformCert struct {
	Data []*certData `json:"data"`
}

type certData struct {
	EffectiveTime      string       `json:"effective_time"`
	EncryptCertificate *encryptCert `json:"encrypt_certificate"`
	ExpireTime         string       `json:"expire_time"`
	SerialNo           string       `json:"serial_no"`
}

type encryptCert struct {
	Algorithm      string `json:"algorithm"`
	AssociatedData string `json:"associated_data"`
	Ciphertext     string `json:"ciphertext"`
	Nonce          string `json:"nonce"`
}

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
	Openid string `json:"openid"` // 用户在直连商户appid下的唯一标识
}

type Amount struct {
	Total         int    `json:"total,omitempty"`          // 订单总金额，单位为分
	PayerTotal    int    `json:"payer_total,omitempty"`    // 用户支付金额，单位为分
	Currency      string `json:"currency,omitempty"`       // CNY：人民币，境内商户号仅支持人民币
	PayerCurrency string `json:"payer_currency,omitempty"` // 用户支付币种
}

type RefundAmount struct {
	Total       int `json:"total,omitempty"`       // 订单总金额，单位为分，只能为整数
	Refund      int `json:"refund,omitempty"`      // 退款金额，币种的最小单位，只能为整数，不能超过原订单支付金额，如果有使用券，后台会按比例退
	PayerTotal  int `json:"payer_total,omitempty"` // 用户实际支付金额，单位为分，只能为整数
	PayerRefund int `json:"payer_refund"`          // 退款给用户的金额，不包含所有优惠券金额
}

type CombineAmount struct {
	TotalAmount   int    `json:"total_amount,omitempty"`   // 订单总金额，单位为分
	Currency      string `json:"currency,omitempty"`       // 标价币种：符合ISO 4217标准的三位字母代码，人民币：CNY
	PayerAmount   int    `json:"payer_amount"`             // 订单现金支付金额
	PayerCurrency string `json:"payer_currency,omitempty"` // 现金支付币种：符合ISO 4217标准的三位字母代码，默认人民币：CNY
}

type SceneInfo struct {
	DeviceId string `json:"device_id,omitempty"` // 商户端设备号（发起扣款请求的商户服务器设备号）
}

type PromotionDetail struct {
	Amount              int            `json:"amount"`                         // 优惠券面额
	CouponId            string         `json:"coupon_id"`                      // 券ID
	Name                string         `json:"name,omitempty"`                 // 优惠名称
	Scope               string         `json:"scope,omitempty"`                // 优惠范围：GLOBAL：全场代金券, SINGLE：单品优惠
	Type                string         `json:"type,omitempty"`                 // 优惠类型：CASH：充值, NOCASH：预充值
	StockId             string         `json:"stock_id,omitempty"`             // 活动ID
	WechatpayContribute int            `json:"wechatpay_contribute,omitempty"` // 微信出资，单位为分
	MerchantContribute  int            `json:"merchant_contribute,omitempty"`  // 商户出资，单位为分
	OtherContribute     int            `json:"other_contribute,omitempty"`     // 其他出资，单位为分
	Currency            string         `json:"currency,omitempty"`             // CNY：人民币，境内商户号仅支持人民币
	GoodsDetail         []*GoodsDetail `json:"goods_detail,omitempty"`         // 单品列表信息
}

type GoodsDetail struct {
	GoodsId         string `json:"goods_id"`                    // 商品编码
	Quantity        int    `json:"quantity"`                    // 用户购买的数量
	UnitPrice       int    `json:"unit_price"`                  // 商品单价，单位为分
	DiscountAmount  int    `json:"discount_amount"`             // 商品优惠金额
	GoodsRemark     string `json:"goods_remark,omitempty"`      // 商品备注信息
	MerchantGoodsID string `json:"merchant_goods_id,omitempty"` // 商户侧商品编码，服务商模式下无此字段
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

type CombineQueryOrder struct {
	CombineAppid      string       `json:"combine_appid"`        // 合单发起方的appid
	CombineMchid      string       `json:"combine_mchid"`        // 合单发起方商户号
	CombineOutTradeNo string       `json:"combine_out_trade_no"` // 合单支付总订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一
	SceneInfo         *SceneInfo   `json:"scene_info,omitempty"` // 支付场景描述
	SubOrders         []*SubOrders `json:"sub_orders"`           // 最多支持子单条数：50
	CombinePayerInfo  *Payer       `json:"combine_payer_info"`   // 支付者信息
}

type TradeBill struct {
	HashType    string `json:"hash_type"`
	HashValue   string `json:"hash_value"`
	DownloadUrl string `json:"download_url"`
}

type BillDetail struct {
	BillSequence int    `json:"bill_sequence"` // 商户将多个文件按账单文件序号的顺序合并为完整的资金账单文件，起始值为1
	HashType     string `json:"hash_type"`
	HashValue    string `json:"hash_value"`
	DownloadUrl  string `json:"download_url"`
	EncryptKey   string `json:"encrypt_key"` // 加密账单文件使用的加密密钥。密钥用商户证书的公钥进行加密，然后进行Base64编码
	Nonce        string `json:"nonce"`       // 加密账单文件使用的随机字符串
}

type DownloadBill struct {
	DownloadBillCount int           `json:"download_bill_count"`
	DownloadBillList  []*BillDetail `json:"download_bill_list"`
}

type RefundOrderResponse struct {
	RefundID            string                        `json:"refund_id"`             // 微信支付退款号
	OutRefundNo         string                        `json:"out_refund_no"`         // 商户退款单号
	TransactionID       string                        `json:"transaction_id"`        // 微信支付系统生成的订单号
	OutTradeNo          string                        `json:"out_trade_no"`          // 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
	Channel             string                        `json:"channel"`               // 退款渠道
	UserReceivedAccount string                        `json:"user_received_account"` // 退款入账账户
	SuccessTime         string                        `json:"success_time"`          // 退款成功时间
	CreateTime          string                        `json:"create_time"`           // 退款创建时间
	Status              string                        `json:"status"`                // 退款状态
	FundsAccount        string                        `json:"funds_account"`         // 资金账户
	Amount              *RefundQueryAmount            `json:"amount"`                // 金额信息
	PromotionDetail     []*RefundQueryPromotionDetail `json:"promotion_detail"`      // 优惠退款信息
}

type RefundQueryResponse struct {
	RefundID            string                        `json:"refund_id"`             // 微信支付退款号
	OutRefundNo         string                        `json:"out_refund_no"`         // 商户退款单号
	TransactionID       string                        `json:"transaction_id"`        // 微信支付系统生成的订单号
	OutTradeNo          string                        `json:"out_trade_no"`          // 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
	Channel             string                        `json:"channel"`               // 退款渠道
	UserReceivedAccount string                        `json:"user_received_account"` // 退款入账账户
	SuccessTime         string                        `json:"success_time"`          // 退款成功时间
	CreateTime          string                        `json:"create_time"`           // 退款创建时间
	Status              string                        `json:"status"`                // 退款状态
	FundsAccount        string                        `json:"funds_account"`         // 资金账户
	Amount              *RefundQueryAmount            `json:"amount"`                // 金额信息
	PromotionDetail     []*RefundQueryPromotionDetail `json:"promotion_detail"`      // 优惠退款信息
}

type RefundQueryAmount struct {
	Total            int    `json:"total"`             // 订单总金额，单位为分
	Refund           int    `json:"refund"`            // 退款金额，币种的最小单位，只能为整数，不能超过原订单支付金额。
	PayerTotal       int    `json:"payer_total"`       // 用户支付金额，单位为分
	PayerRefund      int    `json:"payer_refund"`      // 用户退款金额，不包含所有优惠券金额
	SettlementRefund int    `json:"settlement_refund"` // 应结退款金额，去掉非充值代金券退款金额后的退款金额，单位为分
	DiscountRefund   int    `json:"discount_refund"`   // 优惠退款金额
	Currency         string `json:"currency"`          // CNY：人民币，境内商户号仅支持人民币
}

type RefundQueryPromotionDetail struct {
	PromotionID  string                    `json:"promotion_id"`           // 券ID，券或立减金额
	Scope        string                    `json:"scope"`                  // 优惠范围，GLOBAL：全场代金券，SINGLE：单品优惠
	Type         string                    `json:"type"`                   // 优惠类型，COUPON：代金券，DISCOUNT：优惠券
	Amount       int                       `json:"amount"`                 // 优惠券面额，用户享受优惠的金额（优惠券面额=微信出资金额+商家出资金额+其他出资方金额），单位为分
	RefundAmount int                       `json:"refund_amount"`          // 优惠退款金额，单位为分
	GoodsDetail  []*RefundQueryGoodsDetail `json:"goods_detail,omitempty"` // 商品列表，优惠商品发送退款时返回商品信息
}

type RefundQueryGoodsDetail struct {
	MerchantGoodsID  string `json:"merchant_goods_id"`            // 商户侧商品编码
	WechatpayGoodsID string `json:"wechatpay_goods_id,omitempty"` // 微信侧商品编码
	GoodsName        string `json:"goods_name,omitempty"`         // 商品名称
	UnitPrice        int    `json:"unit_price"`                   // 商品单价金额
	RefundAmount     int    `json:"refund_amount"`                // 商品退款金额
	RefundQuantity   int    `json:"refund_quantity"`              // 商品退货数量
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

type ScoreOrderCreate struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号ID。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	OutTradeNo          string           `json:"out_trade_no"`                // 调用接口提交的商户服务订单号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务ID。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	PostPayments        []*PostPayments  `json:"post_payments,omitempty"`     // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	RiskFund            *RiskFund        `json:"risk_fund"`                   // 订单风险金信息
	TimeRange           *TimeRange       `json:"time_range"`                  // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	Attach              string           `json:"attach,omitempty"`            // 商户数据包,可存放本订单所需信息，需要先urlencode后传入，总长度不大于256字符,超出报错处理。
	NotifyUrl           string           `json:"notify_url,omitempty"`        // 商户接收用户确认订单或扣款成功回调通知的地址。
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	Package             string           `json:"package"`                     // 用户跳转到微信侧小程序订单数据，需确认模式特有API中调起支付分-确认订单传入。该数据一小时内有效。
}

type PostPayments struct {
	Name        string `json:"name,omitempty"`        // 付费项目名称
	Description string `json:"description,omitempty"` // 描述计费规则，不超过30个字符，超出报错处理。
	Amount      int    `json:"amount"`                // 此付费项目总金额，大于等于0，单位为分，等于0时代表不需要扣费，只能为整数
	Count       int    `json:"count,omitempty"`       // 付费项目的数量。
}

type PostDiscounts struct {
	Name        string `json:"name,omitempty"`        // 优惠名称说明。
	Description string `json:"description,omitempty"` // 优惠使用条件说明。
	Amount      int    `json:"amount"`                // 优惠金额，只能为整数
	Count       int    `json:"count,omitempty"`       // 优惠的数量。
}

type RiskFund struct {
	Name        string `json:"name"`                  // 风险金名称。DEPOSIT：押金，ADVANCE：预付款，CASH_DEPOSIT：保证金，ESTIMATE_ORDER_COST：预估订单费用
	Description string `json:"description,omitempty"` // 风险说明
	Amount      int    `json:"amount"`                // 风险金额
}

type TimeRange struct {
	StartTime       string `json:"start_time"`                  // 服务开始时间，20091225091010
	StartTimeRemark string `json:"start_time_remark,omitempty"` // 服务开始时间备注
	EndTime         string `json:"end_time,omitempty"`          // 预计服务结束时间，20091225121010
	EndTimeRemark   string `json:"end_time_remark,omitempty"`   // 预计服务结束时间备注
}

type Location struct {
	StartLocation string `json:"start_location,omitempty"` // 服务开始地点
	EndLocation   string `json:"end_location,omitempty"`   // 服务结束位置
}

type ScoreOrderQuery struct {
	// todo: finish
}
