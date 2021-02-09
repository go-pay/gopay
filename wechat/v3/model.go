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
	v3Refund          = "/v3/refund/domestic/refunds"

	// 微信支付分
	v3ScorePermission                       = "/v3/payscore/permissions"                                 // 商户预授权 POST
	v3ScorePermissionAuthorizationQuery     = "/v3/payscore/permissions/authorization-code/%s"           // authorization_code 查询用户授权记录（授权协议号） GET
	v3ScorePermissionAuthorizationTerminate = "/v3/payscore/permissions/authorization-code/%s/terminate" // authorization_code 解除用户授权关系（授权协议号） POST
	v3ScorePermissionOpenidQuery            = "/v3/payscore/permissions/openid/%s"                       // openid 查询用户授权记录（openid） GET
	v3ScorePermissionOpenidTerminate        = "/v3/payscore/permissions/openid/%s/terminate"             // openid 解除用户授权记录（openid） POST
	v3ScoreOrderDirectComplete              = "/payscore/serviceorder/direct-complete"                   // 创单结单合并 POST
	v3ScoreOrder                            = "/v3/payscore/serviceorder"                                // 创建支付分订单 POST
	v3ScoreOrderQuery                       = "/v3/payscore/serviceorder"                                // 查询支付分订单 GET
	v3ScoreOrderCancel                      = "/v3/payscore/serviceorder/%s/cancel"                      // out_trade_no 取消支付分订单 POST
	v3ScoreOrderModify                      = "/v3/payscore/serviceorder/%s/modify"                      // out_trade_no 修改订单金额 POST
	v3ScoreOrderComplete                    = "/v3/payscore/serviceorder/%s/complete"                    // out_trade_no 完结支付分订单 POST
	v3ScoreOrderPay                         = "/v3/payscore/serviceorder/%s/pay"                         // out_trade_no 商户发起催收扣款 POST
	v3ScoreOrderSync                        = "/v3/payscore/serviceorder/%s/sync"                        // out_trade_no 同步服务订单信息 POST

	// 微信先享卡
	v3CardPre     = "/v3/discount-card/cards"                     // 预受理领卡请求 POST
	v3CardAddUser = "/v3/discount-card/cards/%s/add-user-records" // out_card_code 增加用户记录 POST
	v3CardQuery   = "/v3/discount-card/cards/%s"                  // out_card_code 查询先享卡订单 GET

	// 支付即服务
	v3GuideRegister = "/v3/smartguide/guides"           // 服务人员注册 POST
	v3GuideAssign   = "/v3/smartguide/guides/%s/assign" // guide_id 服务人员分配 POST
	v3GuideQuery    = "/v3/smartguide/guides"           // 服务人员查询 GET
	v3GuideUpdate   = "/v3/smartguide/guides/%s"        // guide_id 服务人员信息更新 PATCH

	// 点金计划
	v3GoldPlanManage       = "/v3/goldplan/merchants/changegoldplanstatus"            // 点金计划管理 POST
	v3GoldPlanBillManage   = "/v3/goldplan/merchants/changecustompagestatus"          // 商家小票管理 POST
	v3GoldPlanFilterManage = "/v3/goldplan/merchants/set-advertising-industry-filter" // 同业过滤标签管理 POST
	v3GoldPlanOpenAdShow   = "/v3/goldplan/merchants/open-advertising-show"           // 开通广告展示 PATCH
	v3GoldPlanCloseAdShow  = "/v3/goldplan/merchants/close-advertising-show"          // 关闭广告展示 PATCH

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

// 退款 Rsp
type RefundRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *RefundOrderResponse `json:"response,omitempty"`
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

// 查询订单 Rsp
type CombineQueryOrderRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *CombineQueryOrder `json:"response,omitempty"`
	Error    string             `json:"-"`
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
	Refund        int    `json:"refund,omitempty"`         // 退款金额，币种的最小单位，只能为整数，不能超过原订单支付金额。
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
	MerchantGoodsID string `json:"merchant_goods_id,omitempty"` // 商户侧商品编码
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
	Payer             *Payer       `json:"combine_payer_info"`   // 支付者信息
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
	RefundID            string             `json:"refund_id" valid:"Required"`             // 微信支付退款号
	OutRefundNo         string             `json:"out_refund_no" valid:"Required"`         // 商户退款单号
	TransactionID       string             `json:"transaction_id" valid:"Required"`        // 微信支付系统生成的订单号
	OutTradeNo          string             `json:"out_trade_no" valid:"Required"`          // 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
	Channel             string             `json:"channel" valid:"Required"`               // 退款渠道
	UserReceivedAccount string             `json:"user_received_account" valid:"Required"` // 退款入账账户
	SuccessTime         string             `json:"success_time"`                           // 退款成功时间
	CreateTime          string             `json:"create_time" valid:"Required"`           // 退款创建时间
	Status              string             `json:"status" valid:"Required"`                // 退款状态
	FundsAccount        string             `json:"funds_account"`                          // 资金账户
	Amount              *Amount            `json:"amount" valid:"Required"`                // 金额信息
	PromotionDetail     []*PromotionDetail `json:"promotion_detail"`                       // 优惠退款信息
}
