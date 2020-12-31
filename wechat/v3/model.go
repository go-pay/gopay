package wechat

type OrderNoType uint8

const (
	Success     = 0
	SignTypeRSA = "RSA"

	MethodPost          = "POST"
	MethodGet           = "GET"
	HeaderAuthorization = "Authorization"

	HeaderTimestamp = "Wechatpay-Timestamp"
	HeaderNonce     = "Wechatpay-Nonce"
	HeaderSignature = "Wechatpay-Signature"
	HeaderSerial    = "Wechatpay-Serial"

	Authorization = "WECHATPAY2-SHA256-RSA2048"

	v3BaseUrlCh = "https://api.mch.weixin.qq.com" // 中国国内

	// 基础H支付
	v3GetCerts                   = "/v3/certificates"
	v3ApiPayApp                  = "/v3/pay/transactions/app"                   // APP 下单
	v3ApiJsapi                   = "/v3/pay/transactions/jsapi"                 // JSAPI 下单
	v3ApiNative                  = "/v3/pay/transactions/native"                // Native 下单
	v3ApiH5                      = "/v3/pay/transactions/h5"                    // H5 下单
	v3ApiQueryOrderTransactionId = "/v3/pay/transactions/id/%s"                 // transaction_id 查询订单
	v3ApiQueryOrderOutTradeNo    = "/v3/pay/transactions/out-trade-no/%s"       // out_trade_no 查询订单
	v3ApiCloseOrder              = "/v3/pay/transactions/out-trade-no/%s/close" // out_trade_no 关闭订单
	v3ApiTradeBill               = "/v3/bill/tradebill"                         // 申请交易账单
	v3ApiFundFlowBill            = "/v3/bill/fundflowbill"                      // 申请资金账单
	v3ApiLevel2FundFlowBill      = "/v3/ecommerce/bill/fundflowbill"            // 申请二级商户资金账单

	// 合单支付
	v3CombinePayApp   = "/v3/combine-transactions/app"
	v3CombinePayH5    = "/v3/combine-transactions/h5"
	v3CombinePayJsapi = "/v3/combine-transactions/jsapi"
	v3CombineNative   = "/v3/combine-transactions/native"
	v3CombineQuery    = "/v3/combine-transactions/out-trade-no/%s"
	v3CombineClose    = "/v3/combine-transactions/out-trade-no/%s/close"

	// 订单号类型，1-微信订单号，2-商户订单号
	TransactionId OrderNoType = 1
	OutTradeNo    OrderNoType = 2

	// v3 异步通知订单状态
	TradeStateSuccess  = "SUCCESS"    // 支付成功
	TradeStateRefund   = "REFUND"     // 转入退款
	TradeStateNoPay    = "NOTPAY"     // 未支付
	TradeStateClosed   = "CLOSED"     // 已关闭
	TradeStateRevoked  = "REVOKED"    // 已撤销（付款码支付）
	TradeStatePaying   = "USERPAYING" // 用户支付中（付款码支付）
	TradeStatePayError = "PAYERROR"   // 支付失败(其他原因，如银行返回失败)
)

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
	GoodsId        string `json:"goods_id"`               // 商品编码
	Quantity       int    `json:"quantity"`               // 用户购买的数量
	UnitPrice      int    `json:"unit_price"`             // 商品单价，单位为分
	DiscountAmount int    `json:"discount_amount"`        // 商品优惠金额
	GoodsRemark    string `json:"goods_remark,omitempty"` // 商品备注信息
}

type QueryOrder struct {
	Appid           string             `json:"appid"`                      // 直连商户申请的公众号或移动应用appid
	Mchid           string             `json:"mchid"`                      // 直连商户的商户号，由微信支付生成并下发
	OutTradeNo      string             `json:"out_trade_no"`               // 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
	TransactionId   string             `json:"transaction_id,omitempty"`   // 微信支付系统生成的订单号
	TradeType       string             `json:"trade_type,omitempty"`       // 交易类型，枚举值：JSAPI：公众号支付, NATIVE：扫码支付, APP：APP支付, MICROPAY：付款码支付, MWEB：H5支付, FACEPAY：刷脸支付
	TradeState      string             `json:"trade_state"`                // 交易状态，枚举值：SUCCESS：支付成功, REFUND：转入退款, NOTPAY：未支付, CLOSED：已关闭, REVOKED：已撤销（付款码支付）, USERPAYING：用户支付中（付款码支付）, PAYERROR：支付失败(其他原因，如银行返回失败)
	TradeStateDesc  string             `json:"trade_state_desc"`           // 交易状态描述
	BankType        string             `json:"bank_type,omitempty"`        // 银行类型，采用字符串类型的银行标识
	Attach          string             `json:"attach,omitempty"`           // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用
	SuccessTime     string             `json:"success_time,omitempty"`     // 支付完成时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒
	Payer           *Payer             `json:"payer"`                      // 支付者信息
	Amount          *Amount            `json:"amount,omitempty"`           // 订单金额信息，当支付成功时返回该字段
	SceneInfo       *SceneInfo         `json:"scene_info,omitempty"`       // 支付场景描述
	PromotionDetail []*PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能，享受优惠时返回该字段
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
