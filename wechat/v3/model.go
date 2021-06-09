package wechat

type OrderNoType uint8

type PlatformCertRsp struct {
	Code     int                 `json:"-"`
	SignInfo *SignInfo           `json:"-"`
	Certs    []*PlatformCertItem `json:"certs"`
	Error    string              `json:"-"`
}

type EmptyRsp struct {
	Code     int       `json:"-"`
	SignInfo *SignInfo `json:"-"`
	Error    string    `json:"-"`
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

// 取消支付分订单 Rsp
type ScoreOrderCancelRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *ScoreOrderCancel `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 修改订单金额 Rsp
type ScoreOrderModifyRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *ScoreOrderModify `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 完结支付分订单 Rsp
type ScoreOrderCompleteRsp struct {
	Code     int                 `json:"-"`
	SignInfo *SignInfo           `json:"-"`
	Response *ScoreOrderComplete `json:"response,omitempty"`
	Error    string              `json:"-"`
}

// 商户发起催收扣款 Rsp
type ScoreOrderPayRsp struct {
	Code     int            `json:"-"`
	SignInfo *SignInfo      `json:"-"`
	Response *ScoreOrderPay `json:"response,omitempty"`
	Error    string         `json:"-"`
}

// 同步服务订单信息 Rsp
type ScoreOrderSyncRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *ScoreOrderSync `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 创单结单合并 Rsp
type ScoreDirectCompleteRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *ScoreDirectComplete `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 创单结单合并 Rsp
type ScorePermissionRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *ScorePermission `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 查询用户授权记录（授权协议号） Rsp
type ScorePermissionQueryRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *ScorePermissionQuery `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 查询用户授权记录（openid） Rsp
type ScorePermissionOpenidQueryRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *ScorePermissionOpenidQuery `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

// 查询投诉单列表 Rsp
type ComplaintListRsp struct {
	Code     int            `json:"-"`
	SignInfo *SignInfo      `json:"-"`
	Response *ComplaintList `json:"response,omitempty"`
	Error    string         `json:"-"`
}

// 查询投诉单详情 Rsp
type ComplaintDetailRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *ComplaintDetail `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 查询投诉单详情 Rsp
type ComplaintNegotiationHistoryRsp struct {
	Code     int                          `json:"-"`
	SignInfo *SignInfo                    `json:"-"`
	Response *ComplaintNegotiationHistory `json:"response,omitempty"`
	Error    string                       `json:"-"`
}

// 创建、查询、更新投诉通知回调地址 Rsp
type ComplaintNotifyUrlRsp struct {
	Code     int                 `json:"-"`
	SignInfo *SignInfo           `json:"-"`
	Response *ComplaintNotifyUrl `json:"response,omitempty"`
	Error    string              `json:"-"`
}

// 商户上传反馈图片 Rsp
type ComplaintUploadImageRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *ComplaintUploadImage `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 请求分账 Rsp
type ProfitSharingOrderRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *ProfitSharingOrderResponse `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

// 查询分账结果 Rsp
type ProfitSharingOrderResultRsp struct {
	Code     int                               `json:"-"`
	SignInfo *SignInfo                         `json:"-"`
	Response *ProfitSharingOrderResultResponse `json:"response,omitempty"`
	Error    string                            `json:"-"`
}

// 解冻剩余资金 Rsp
type ProfitSharingOrderUnfreezeRsp struct {
	Code     int                                 `json:"-"`
	SignInfo *SignInfo                           `json:"-"`
	Response *ProfitSharingOrderUnfreezeResponse `json:"response,omitempty"`
	Error    string                              `json:"-"`
}

// 添加分账接收方 Rsp
type ProfitSharingAddReceiverRsp struct {
	Code     int                               `json:"-"`
	SignInfo *SignInfo                         `json:"-"`
	Response *ProfitSharingAddReceiverResponse `json:"response,omitempty"`
	Error    string                            `json:"-"`
}

// 删除分账接收方 Rsp
type ProfitSharingDeleteReceiverRsp struct {
	Code     int                                  `json:"-"`
	SignInfo *SignInfo                            `json:"-"`
	Response *ProfitSharingDeleteReceiverResponse `json:"response,omitempty"`
	Error    string                               `json:"-"`
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
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号ID。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务ID。
	OutTradeNo          string           `json:"out_trade_no"`                // 调用接口提交的商户服务订单号。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	TotalAmount         int              `json:"total_amount,omitempty"`      // 总金额，大于等于0的数字，单位为分，只能为整数
	PostPayments        []*PostPayments  `json:"post_payments,omitempty"`     // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	RiskFund            *RiskFund        `json:"risk_fund"`                   // 订单风险金信息
	TimeRange           *TimeRange       `json:"time_range"`                  // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	Attach              string           `json:"attach,omitempty"`            // 商户数据包,可存放本订单所需信息，需要先urlencode后传入，总长度不大于256字符,超出报错处理。
	NotifyUrl           string           `json:"notify_url"`                  // 商户接收用户确认订单或扣款成功回调通知的地址。
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	NeedCollection      bool             `json:"need_collection,omitempty"`   // 是否需要收款
	Collection          *Collection      `json:"collection,omitempty"`        // 收款信息
	Openid              string           `json:"openid,omitempty"`            // 微信用户在商户对应appid下的唯一标识
}

type Collection struct {
	State        string     `json:"state"`                   // 收款状态，USER_PAYING：待支付，USER_PAID：已支付
	TotalAmount  int        `json:"total_amount,omitempty"`  // 总金额，大于等于0的数字，单位为分，只能为整数
	PayingAmount int        `json:"paying_amount,omitempty"` // 等待用户付款金额，只能为整数
	PaidAmount   int        `json:"paid_amount,omitempty"`   // 用户已付款的金额，只能为整数
	Details      []*Details `json:"details,omitempty"`       // 收款明细列表
}

type Details struct {
	Seq             int                `json:"seq,omitempty"`              // 收款序号
	Amount          int                `json:"amount,omitempty"`           // 单笔收款动作的金额，只能为整数
	PaidType        string             `json:"paid_type,omitempty"`        // 收款成功渠道，NEWTON：微信支付分，MCH：商户渠道
	PaidTime        string             `json:"paid_time,omitempty"`        // 支付成功时间，支持两种格式:yyyyMMddHHmmss和yyyyMMdd
	TransactionId   string             `json:"transaction_id,omitempty"`   // 结单交易单号，等于普通支付接口中的transaction_id
	PromotionDetail []*PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能，享受优惠时返回该字段
}

type ScoreOrderCancel struct {
	Appid      string `json:"appid"`        // 调用接口提交的公众账号ID。
	Mchid      string `json:"mchid"`        // 调用接口提交的商户号。
	ServiceId  string `json:"service_id"`   // 调用该接口提交的服务ID。
	OutTradeNo string `json:"out_trade_no"` // 调用接口提交的商户服务订单号。
	OrderId    string `json:"order_id"`     // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
}

type ScoreOrderModify struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号ID。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务ID。
	OutTradeNo          string           `json:"out_trade_no"`                // 调用接口提交的商户服务订单号。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	TotalAmount         int              `json:"total_amount,omitempty"`      // 总金额，大于等于0的数字，单位为分，只能为整数
	PostPayments        []*PostPayments  `json:"post_payments,omitempty"`     // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	RiskFund            *RiskFund        `json:"risk_fund"`                   // 订单风险金信息
	TimeRange           *TimeRange       `json:"time_range"`                  // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	Attach              string           `json:"attach,omitempty"`            // 商户数据包,可存放本订单所需信息，需要先urlencode后传入，总长度不大于256字符,超出报错处理。
	NotifyUrl           string           `json:"notify_url,omitempty"`        // 商户接收用户确认订单或扣款成功回调通知的地址。
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	NeedCollection      bool             `json:"need_collection,omitempty"`   // 是否需要收款
	Collection          *Collection      `json:"collection,omitempty"`        // 收款信息
}

type ScoreOrderComplete struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号ID。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务ID。
	OutTradeNo          string           `json:"out_trade_no"`                // 调用接口提交的商户服务订单号。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	TotalAmount         int              `json:"total_amount"`                // 总金额，大于等于0的数字，单位为分，只能为整数
	PostPayments        []*PostPayments  `json:"post_payments"`               // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	RiskFund            *RiskFund        `json:"risk_fund"`                   // 订单风险金信息
	TimeRange           *TimeRange       `json:"time_range,omitempty"`        // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	NeedCollection      bool             `json:"need_collection,omitempty"`   // 是否需要收款
}

type ScoreOrderPay struct {
	Appid      string `json:"appid"`        // 调用接口提交的公众账号ID。
	Mchid      string `json:"mchid"`        // 调用接口提交的商户号。
	ServiceId  string `json:"service_id"`   // 调用该接口提交的服务ID。
	OutTradeNo string `json:"out_trade_no"` // 调用接口提交的商户服务订单号。
	OrderId    string `json:"order_id"`     // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
}

type ScoreOrderSync struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号ID。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务ID。
	OutTradeNo          string           `json:"out_trade_no"`                // 调用接口提交的商户服务订单号。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	TotalAmount         int              `json:"total_amount"`                // 总金额，大于等于0的数字，单位为分，只能为整数
	PostPayments        []*PostPayments  `json:"post_payments,omitempty"`     // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	RiskFund            *RiskFund        `json:"risk_fund,omitempty"`         // 订单风险金信息
	TimeRange           *TimeRange       `json:"time_range,omitempty"`        // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	Attach              string           `json:"attach,omitempty"`            // 商户数据包,可存放本订单所需信息，需要先urlencode后传入，总长度不大于256字符,超出报错处理。
	NotifyUrl           string           `json:"notify_url,omitempty"`        // 商户接收用户确认订单或扣款成功回调通知的地址。
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	NeedCollection      bool             `json:"need_collection,omitempty"`   // 是否需要收款
	Collection          *Collection      `json:"collection,omitempty"`        // 收款信息
	Openid              string           `json:"openid"`                      // 微信用户在商户对应appid下的唯一标识
}

type ScoreDirectComplete struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号ID。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	OutTradeNo          string           `json:"out_trade_no"`                // 调用接口提交的商户服务订单号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务ID。
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	PostPayments        []*PostPayments  `json:"post_payments"`               // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	TimeRange           *TimeRange       `json:"time_range"`                  // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	Attach              string           `json:"attach,omitempty"`            // 商户数据包,可存放本订单所需信息，需要先urlencode后传入，总长度不大于256字符,超出报错处理。
	NotifyUrl           string           `json:"notify_url,omitempty"`        // 商户接收用户确认订单或扣款成功回调通知的地址。
	TotalAmount         int              `json:"total_amount"`                // 金额：数字，必须≥0（单位：分）
}

type ScorePermission struct {
	ApplyPermissionsToken string `json:"apply_permissions_token"` // 用于跳转到微信侧小程序授权数据,跳转到微信侧小程序传入，时效性为1小时
}

type ScorePermissionQuery struct {
	Appid                    string `json:"appid"`                               // 调用接口提交的公众账号ID。
	Mchid                    string `json:"mchid"`                               // 调用接口提交的商户号。
	ServiceId                string `json:"service_id"`                          // 调用该接口提交的服务ID。
	Openid                   string `json:"openid,omitempty"`                    // 微信用户在商户对应appid下的唯一标识
	AuthorizationCode        string `json:"authorization_code"`                  // 预授权成功时的授权协议号。
	AuthorizationState       string `json:"authorization_state"`                 // 标识用户授权服务情况：UNAVAILABLE：用户未授权服务，AVAILABLE：用户已授权服务，UNBINDUSER：未绑定用户（已经预授权但未完成正式授权）
	NotifyUrl                string `json:"notify_url,omitempty"`                // 商户接收用户确认订单或扣款成功回调通知的地址。
	CancelAuthorizationTime  string `json:"cancel_authorization_time,omitempty"` // 最近一次解除授权时间, 示例值：2015-05-20T13:29:35.120+08:00
	AuthorizationSuccessTime string `json:"authorization_success_time"`          // 最近一次授权成功时间, 示例值：2015-05-20T13:29:35.120+08:00
}

type ScorePermissionOpenidQuery struct {
	Appid                    string `json:"appid"`                               // 调用接口提交的公众账号ID。
	Mchid                    string `json:"mchid"`                               // 调用接口提交的商户号。
	ServiceId                string `json:"service_id"`                          // 调用该接口提交的服务ID。
	Openid                   string `json:"openid,omitempty"`                    // 微信用户在商户对应appid下的唯一标识
	AuthorizationCode        string `json:"authorization_code"`                  // 预授权成功时的授权协议号。
	AuthorizationState       string `json:"authorization_state"`                 // 标识用户授权服务情况：UNAVAILABLE：用户未授权服务，AVAILABLE：用户已授权服务，UNBINDUSER：未绑定用户（已经预授权但未完成正式授权）
	CancelAuthorizationTime  string `json:"cancel_authorization_time,omitempty"` // 最近一次解除授权时间, 示例值：2015-05-20T13:29:35.120+08:00
	AuthorizationSuccessTime string `json:"authorization_success_time"`          // 最近一次授权成功时间, 示例值：2015-05-20T13:29:35.120+08:00
}

type ComplaintList struct {
	Data       []*ComplaintListItem `json:"data,omitempty"`        // 用户投诉信息详情
	Limit      int                  `json:"limit"`                 // 设置该次请求返回的最大投诉条数，范围【1,50】
	Offset     int                  `json:"offset"`                // 该次请求的分页开始位置，从0开始计数，例如offset=10，表示从第11条记录开始返回。
	TotalCount int                  `json:"total_count,omitempty"` // 投诉总条数，当offset=0时返回
}

type ComplaintListItem struct {
	ComplaintId           string                `json:"complaint_id"`                   // 投诉单对应的投诉单号
	ComplaintTime         string                `json:"complaint_time"`                 // 投诉时间, 例如：2015-05-20T13:29:35.120+08:00表示北京时间2015年05月20日13点29分35秒
	ComplaintDetail       string                `json:"complaint_detail"`               // 投诉的具体描述
	ComplaintState        string                `json:"complaint_state"`                // 投诉单状态, PENDING：待处理, PROCESSING：处理中, PROCESSED：已处理完成
	PayerPhone            string                `json:"payer_phone,omitempty"`          // 投诉人联系方式。该字段已做加密处理
	ComplaintOrderInfo    []*ComplaintOrderInfo `json:"complaint_order_info,omitempty"` // 投诉单关联订单信息
	ComplaintFullRefunded bool                  `json:"complaint_full_refunded"`        // 投诉单下所有订单是否已全部全额退款
	IncomingUserResponse  bool                  `json:"incoming_user_response"`         // 投诉单是否有待回复的用户留言
	UserComplaintTimes    int                   `json:"user_complaint_times"`           // 用户投诉次数
}

type ComplaintOrderInfo struct {
	TransactionId string `json:"transaction_id"` // 投诉单关联的微信订单号
	OutTradeNo    string `json:"out_trade_no"`   // 投诉单关联的商户订单号
	Amount        int    `json:"amount"`         // 订单金额，单位（分）
}

type ComplaintDetail struct {
	ComplaintId           string                `json:"complaint_id"`                   // 投诉单对应的投诉单号
	ComplaintTime         string                `json:"complaint_time"`                 // 投诉时间, 例如：2015-05-20T13:29:35.120+08:00表示北京时间2015年05月20日13点29分35秒
	ComplaintDetail       string                `json:"complaint_detail"`               // 投诉的具体描述
	ComplaintedMchid      string                `json:"complainted_mchid,omitempty"`    // 投诉单对应的被诉商户号。
	ComplaintState        string                `json:"complaint_state"`                // 投诉单状态, PENDING：待处理, PROCESSING：处理中, PROCESSED：已处理完成
	PayerPhone            string                `json:"payer_phone,omitempty"`          // 投诉人联系方式。该字段已做加密处理
	PayerOpenid           string                `json:"payer_openid"`                   // 投诉人在商户appid下的唯一标识
	ComplaintOrderInfo    []*ComplaintOrderInfo `json:"complaint_order_info,omitempty"` // 投诉单关联订单信息
	ComplaintFullRefunded bool                  `json:"complaint_full_refunded"`        // 投诉单下所有订单是否已全部全额退款
	IncomingUserResponse  bool                  `json:"incoming_user_response"`         // 投诉单是否有待回复的用户留言
	UserComplaintTimes    int                   `json:"user_complaint_times"`           // 用户投诉次数
}

type ComplaintNegotiationHistory struct {
	Data       []*ComplaintNegotiationHistoryItem `json:"data,omitempty"`        // 投诉协商历史
	Limit      int                                `json:"limit"`                 // 设置该次请求返回的最大投诉条数，范围【1,50】
	Offset     int                                `json:"offset"`                // 该次请求的分页开始位置，从0开始计数，例如offset=10，表示从第11条记录开始返回。
	TotalCount int                                `json:"total_count,omitempty"` // 投诉总条数，当offset=0时返回
}

type ComplaintNegotiationHistoryItem struct {
	LogId          string   `json:"log_id"`          // 操作流水号
	Operator       string   `json:"operator"`        // 当前投诉协商记录的操作人
	OperateTime    string   `json:"operate_time"`    // 当前投诉协商记录的操作时间
	OperateType    string   `json:"operate_type"`    // 当前投诉协商记录的操作类型
	OperateDetails string   `json:"operate_details"` // 当前投诉协商记录的具体内容
	ImageList      []string `json:"image_list"`      // 当前投诉协商记录提交的图片凭证（url格式），最多返回4张图片，url有效时间为1小时。如未查询到协商历史图片凭证，则返回空数组。
}

type ComplaintNotifyUrl struct {
	Mchid string `json:"mchid"` // 返回创建回调地址的商户号，由微信支付生成并下发。
	Url   string `json:"url"`   // 通知地址，仅支持https。
}

type ComplaintUploadImage struct {
	MediaId string `json:"media_id"` // 微信返回的媒体文件标识ID。
}

type ProfitSharingOrderResponse struct {
	TransactionId string                   `json:"transaction_id,omitempty"` // 微信订单号
	OutOrderNo    string                   `json:"out_order_no,omitempty"`   // 商户分账单号
	OrderId       string                   `json:"order_id,omitempty"`       // 微信分账单号
	State         string                   `json:"state"`                    // 分账单状态（每个接收方的分账结果请查看receivers中的result字段）:PROCESSING：处理中,FINISHED：分账完成
	Receivers     []*ProfitSharingReceiver `json:"receivers"`                // 分账接收方列表
}

type ProfitSharingOrderResultResponse struct {
	TransactionId string                   `json:"transaction_id"` // 微信订单号
	OutOrderNo    string                   `json:"out_order_no"`   // 商户分账单号
	OrderId       string                   `json:"order_id"`       // 微信分账单号
	State         string                   `json:"state"`          // 分账单状态（每个接收方的分账结果请查看receivers中的result字段）:PROCESSING：处理中,FINISHED：分账完成
	Receivers     []*ProfitSharingReceiver `json:"receivers"`      // 分账接收方列表
}

type ProfitSharingOrderUnfreezeResponse struct {
	TransactionId string                   `json:"transaction_id,omitempty"` // 微信支付订单号
	OutOrderNo    string                   `json:"out_order_no,omitempty"`   // 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OrderId       string                   `json:"order_id,omitempty"`       // 微信分账单号，微信系统返回的唯一标识
	State         string                   `json:"state"`                    // 分账单状态（每个接收方的分账结果请查看receivers中的result字段）:PROCESSING：处理中,FINISHED：分账完成
	Receivers     []*ProfitSharingReceiver `json:"receivers"`                // 分账接收方列表
}

// 查询剩余待分金额 Rsp
type ProfitSharingUnsplitAmountRsp struct {
	Code     int                                 `json:"-"`
	SignInfo *SignInfo                           `json:"-"`
	Response *ProfitSharingUnsplitAmountResponse `json:"response,omitempty"`
	Error    string                              `json:"-"`
}

type ProfitSharingUnsplitAmountResponse struct {
	TransactionId string `json:"transaction_id"` // 微信支付订单号
	UnsplitAmount int    `json:"unsplit_amount"` // 订单剩余待分金额，整数，单位为分
}

// ProfitSharingReceiver 分账接收方
type ProfitSharingReceiver struct {
	Amount      int    `json:"amount"`      // 分账金额
	Description string `json:"description"` // 分账描述
	Type        string `json:"type"`        // 分账接收方类型
	Account     string `json:"account"`     // 分账接收方帐号
	Result      string `json:"result"`      // 分账结果,PENDING：待分账,SUCCESS：分账成功,CLOSED：已关闭
	FailReason  string `json:"fail_reason"` // 分账失败原因ACCOUNT_ABNORMAL : 分账接收账户异常、NO_RELATION : 分账关系已解除、RECEIVER_HIGH_RISK : 高风险接收方、RECEIVER_REAL_NAME_NOT_VERIFIED : 接收方未实名、NO_AUTH : 分账权限已解除
	CreateTime  string `json:"create_time"` // 分账创建时间,遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE
	FinishTime  string `json:"finish_time"` // 分账完成时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE
}

type ProfitSharingAddReceiverResponse struct {
	Type           string `json:"type"`            // 分账接收方类型MERCHANT_ID：商户ID,PERSONAL_OPENID：个人openid（由父商户APPID转换得到）
	Account        string `json:"account"`         // 分账接收方帐号
	Name           string `json:"name"`            // 分账接收方类型是MERCHANT_ID时，是商户全称（必传），当商户是小微商户或个体户时，是开户人姓名 分账接收方类型是PERSONAL_OPENID时，是个人姓名（选传，传则校验）
	RelationType   string `json:"relation_type"`   // 商户与接收方的关系。STORE：门店STAFF：员工	STORE_OWNER：店主	PARTNER：合作伙伴	HEADQUARTER：总部	BRAND：品牌方	DISTRIBUTOR：分销商	USER：用户	SUPPLIER： 供应商	CUSTOM：自定义
	CustomRelation string `json:"custom_relation"` // 子商户与接收方具体的关系，本字段最多10个字。当字段relation_type的值为CUSTOM时，本字段必填;当字段relation_type的值不为CUSTOM时，本字段无需填写。
}

type ProfitSharingDeleteReceiverResponse struct {
	Type    string `json:"type"`    // 分账接收方类型MERCHANT_ID：商户ID,PERSONAL_OPENID：个人openid（由父商户APPID转换得到）
	Account string `json:"account"` // 分账接收方帐号
}
