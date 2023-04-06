package wechat

// 微信支付分订单类型：1-微信订单号，2-商户订单号，3-微信侧回跳到商户前端时用于查单的单据查询id（查询支付分订单中会使用）
type OrderNoType uint8

// 微信证书类型：RSA、SM2
type CertType string

type PlatformCertRsp struct {
	Code  int                 `json:"-"`
	Certs []*PlatformCertItem `json:"certs"`
	Error string              `json:"-"`
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

// 请求分账 Rsp
type ProfitShareOrderRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *ProfitShareOrder `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 查询分账结果 Rsp
type ProfitShareOrderQueryRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *ProfitShareOrderQuery `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

// 请求分账回退 Rsp
type ProfitShareReturnRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *ProfitShareReturn `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 请求分账回退 Rsp
type ProfitShareReturnResultRsp struct {
	Code     int                      `json:"-"`
	SignInfo *SignInfo                `json:"-"`
	Response *ProfitShareReturnResult `json:"response,omitempty"`
	Error    string                   `json:"-"`
}

// 解冻剩余资金 Rsp
type ProfitShareOrderUnfreezeRsp struct {
	Code     int                       `json:"-"`
	SignInfo *SignInfo                 `json:"-"`
	Response *ProfitShareOrderUnfreeze `json:"response,omitempty"`
	Error    string                    `json:"-"`
}

// 查询剩余待分金额 Rsp
type ProfitShareUnsplitAmountRsp struct {
	Code     int                       `json:"-"`
	SignInfo *SignInfo                 `json:"-"`
	Response *ProfitShareUnsplitAmount `json:"response,omitempty"`
	Error    string                    `json:"-"`
}

// 添加分账接收方 Rsp
type ProfitShareAddReceiverRsp struct {
	Code     int                     `json:"-"`
	SignInfo *SignInfo               `json:"-"`
	Response *ProfitShareAddReceiver `json:"response,omitempty"`
	Error    string                  `json:"-"`
}

// 删除分账接收方 Rsp
type ProfitShareDeleteReceiverRsp struct {
	Code     int                        `json:"-"`
	SignInfo *SignInfo                  `json:"-"`
	Response *ProfitShareDeleteReceiver `json:"response,omitempty"`
	Error    string                     `json:"-"`
}

type ProfitShareMerchantConfigsRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *ProfitShareMerchantConfigs `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

// 申请分账账单 Rsp
type ProfitShareBillsRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *ProfitShareBills `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 服务人员注册 Rsp
type SmartGuideRegRsp struct {
	Code     int            `json:"-"`
	SignInfo *SignInfo      `json:"-"`
	Response *SmartGuideReg `json:"response,omitempty"`
	Error    string         `json:"-"`
}

// 服务人员查询 Rsp
type SmartGuideQueryRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *SmartGuideQuery `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 发起批量转账 Rsp
type TransferRsp struct {
	Code     int       `json:"-"`
	SignInfo *SignInfo `json:"-"`
	Response *Transfer `json:"response,omitempty"`
	Error    string    `json:"-"`
}

// 微信批次单号查询批次单 Rsp
type TransferQueryRsp struct {
	Code     int            `json:"-"`
	SignInfo *SignInfo      `json:"-"`
	Response *TransferQuery `json:"response,omitempty"`
	Error    string         `json:"-"`
}

// 微信批次单号查询批次单（服务商） Rsp
type PartnerTransferQueryRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *PartnerTransferQuery `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 微信明细单号查询明细单 Rsp
type TransferDetailRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *TransferDetailQuery `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 微信明细单号查询明细单（服务商） Rsp
type PartnerTransferDetailRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *PartnerTransferDetail `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

// 商家批次单号查询批次单 Rsp
type TransferMerchantQueryRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *TransferMerchantQuery `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

// 商家批次单号查询批次单（服务商） Rsp
type PartnerTransferMerchantQueryRsp struct {
	Code     int                           `json:"-"`
	SignInfo *SignInfo                     `json:"-"`
	Response *PartnerTransferMerchantQuery `json:"response,omitempty"`
	Error    string                        `json:"-"`
}

// 商家明细单号查询明细单 Rsp
type TransferMerchantDetailRsp struct {
	Code     int                     `json:"-"`
	SignInfo *SignInfo               `json:"-"`
	Response *TransferMerchantDetail `json:"response,omitempty"`
	Error    string                  `json:"-"`
}

// 商家明细单号查询明细单（服务商） Rsp
type PartnerTransferMerchantDetailRsp struct {
	Code     int                            `json:"-"`
	SignInfo *SignInfo                      `json:"-"`
	Response *PartnerTransferMerchantDetail `json:"response,omitempty"`
	Error    string                         `json:"-"`
}

// 转账电子回单申请受理 Rsp
type TransferReceiptRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *TransferReceipt `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 查询转账电子回单 Rsp
type TransferReceiptQueryRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *TransferReceiptQuery `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 转账明细电子回单受理 Rsp
type TransferDetailReceiptRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *TransferDetailReceipt `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

// 查询转账明细电子回单受理结果 Rsp
type TransferDetailReceiptQueryRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *TransferDetailReceiptQuery `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

// 查询特约商户账户实时余额 Rsp
type EcommerceBalanceRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *EcommerceBalance `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 查询账户实时余额 Rsp
type MerchantBalanceRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *MerchantBalance `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 商户银行来账查询 Rsp
type MerchantIncomeRecordRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *MerchantIncomeRecord `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 特约商户银行来账查询 Rsp
type PartnerIncomeRecordRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *PartnerIncomeRecord `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 点金计划管理 Rsp
type GoldPlanManageRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *GoldPlanManage `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 特约商户余额提现 Rsp
type WithdrawRsp struct {
	Code     int       `json:"-"`
	SignInfo *SignInfo `json:"-"`
	Response *Withdraw `json:"response,omitempty"`
	Error    string    `json:"-"`
}

// 查询特约商户提现状态 Rsp
type WithdrawStatusRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *WithdrawStatus `json:"response,omitempty"`
	Error    string          `json:"-"`
}

type EntrustPayNotifyRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *BankSearchBranch `json:"response,omitempty"`
	Error    string            `json:"-"`
}

type VehicleParkingQueryRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *VehicleParkingQuery `json:"response,omitempty"`
	Error    string               `json:"-"`
}

type VehicleParkingInRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *VehicleParkingIn `json:"response,omitempty"`
	Error    string            `json:"-"`
}

type VehicleParkingFeeRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *VehicleParkingFee `json:"response,omitempty"`
	Error    string             `json:"-"`
}

type VehicleParkingOrderRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *VehicleParkingOrder `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// =========================================================分割=========================================================

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
	Sign      string `json:"sign"`
}

type AppletParams struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

type AppletScoreExtraData struct {
	MchId     string `json:"mch_id"`
	Package   string `json:"package"`
	TimeStamp string `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	SignType  string `json:"sign_type"`
	Sign      string `json:"sign"`
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
	Openid    string `json:"openid"`               // 用户在直连商户appid下的唯一标识
	SubOpenid string `json:"sub_openid,omitempty"` // 用户在sub_appid下的标识
}

type Amount struct {
	Total         int    `json:"total,omitempty"`          // 订单总金额，单位为分
	PayerTotal    int    `json:"payer_total,omitempty"`    // 用户支付金额，单位为分
	DiscountTotal int    `json:"discount_total,omitempty"` // 订单折扣
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
	CouponId            string         `json:"coupon_id"`                      // 券Id
	Name                string         `json:"name,omitempty"`                 // 优惠名称
	Scope               string         `json:"scope,omitempty"`                // 优惠范围：GLOBAL：全场代金券, SINGLE：单品优惠
	Type                string         `json:"type,omitempty"`                 // 优惠类型：CASH：充值, NOCASH：预充值
	StockId             string         `json:"stock_id,omitempty"`             // 活动Id
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
	MerchantGoodsId string `json:"merchant_goods_id,omitempty"` // 商户侧商品编码，服务商模式下无此字段
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

type RefundOrderResponse struct {
	RefundId            string                        `json:"refund_id"`             // 微信支付退款号
	OutRefundNo         string                        `json:"out_refund_no"`         // 商户退款单号
	TransactionId       string                        `json:"transaction_id"`        // 微信支付系统生成的订单号
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
	RefundId            string                        `json:"refund_id"`             // 微信支付退款号
	OutRefundNo         string                        `json:"out_refund_no"`         // 商户退款单号
	TransactionId       string                        `json:"transaction_id"`        // 微信支付系统生成的订单号
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
	PromotionId  string                    `json:"promotion_id"`           // 券Id，券或立减金额
	Scope        string                    `json:"scope"`                  // 优惠范围，GLOBAL：全场代金券，SINGLE：单品优惠
	Type         string                    `json:"type"`                   // 优惠类型，COUPON：代金券，DISCOUNT：优惠券
	Amount       int                       `json:"amount"`                 // 优惠券面额，用户享受优惠的金额（优惠券面额=微信出资金额+商家出资金额+其他出资方金额），单位为分
	RefundAmount int                       `json:"refund_amount"`          // 优惠退款金额，单位为分
	GoodsDetail  []*RefundQueryGoodsDetail `json:"goods_detail,omitempty"` // 商品列表，优惠商品发送退款时返回商品信息
}

type RefundQueryGoodsDetail struct {
	MerchantGoodsId  string `json:"merchant_goods_id"`            // 商户侧商品编码
	WechatpayGoodsId string `json:"wechatpay_goods_id,omitempty"` // 微信侧商品编码
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
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
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
	Name        string `json:"name"`        // 付费项目名称
	Amount      int    `json:"amount"`      // 此付费项目总金额，大于等于0，单位为分，等于0时代表不需要扣费，只能为整数
	Description string `json:"description"` // 描述计费规则，不超过30个字符，超出报错处理。
	Count       int    `json:"count"`       // 付费项目的数量。
}

type PostDiscounts struct {
	Name        string `json:"name"`        // 优惠名称说明。
	Description string `json:"description"` // 优惠使用条件说明。
	Amount      int    `json:"amount"`      // 优惠金额，只能为整数
	Count       int    `json:"count"`       // 优惠的数量。
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
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
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
	Appid      string `json:"appid"`        // 调用接口提交的公众账号Id。
	Mchid      string `json:"mchid"`        // 调用接口提交的商户号。
	ServiceId  string `json:"service_id"`   // 调用该接口提交的服务Id。
	OutOrderNo string `json:"out_order_no"` // 调用接口提交的商户服务订单号。
	OrderId    string `json:"order_id"`     // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
}

type ScoreOrderModify struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
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
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
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
	Appid      string `json:"appid"`        // 调用接口提交的公众账号Id。
	Mchid      string `json:"mchid"`        // 调用接口提交的商户号。
	ServiceId  string `json:"service_id"`   // 调用该接口提交的服务Id。
	OutOrderNo string `json:"out_order_no"` // 调用接口提交的商户服务订单号。
	OrderId    string `json:"order_id"`     // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
}

type ScoreOrderSync struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
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
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
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
	Appid                    string `json:"appid"`                               // 调用接口提交的公众账号Id。
	Mchid                    string `json:"mchid"`                               // 调用接口提交的商户号。
	ServiceId                string `json:"service_id"`                          // 调用该接口提交的服务Id。
	Openid                   string `json:"openid,omitempty"`                    // 微信用户在商户对应appid下的唯一标识
	AuthorizationCode        string `json:"authorization_code"`                  // 预授权成功时的授权协议号。
	AuthorizationState       string `json:"authorization_state"`                 // 标识用户授权服务情况：UNAVAILABLE：用户未授权服务，AVAILABLE：用户已授权服务，UNBINDUSER：未绑定用户（已经预授权但未完成正式授权）
	NotifyUrl                string `json:"notify_url,omitempty"`                // 商户接收用户确认订单或扣款成功回调通知的地址。
	CancelAuthorizationTime  string `json:"cancel_authorization_time,omitempty"` // 最近一次解除授权时间, 示例值：2015-05-20T13:29:35.120+08:00
	AuthorizationSuccessTime string `json:"authorization_success_time"`          // 最近一次授权成功时间, 示例值：2015-05-20T13:29:35.120+08:00
}

type ScorePermissionOpenidQuery struct {
	Appid                    string `json:"appid"`                               // 调用接口提交的公众账号Id。
	Mchid                    string `json:"mchid"`                               // 调用接口提交的商户号。
	ServiceId                string `json:"service_id"`                          // 调用该接口提交的服务Id。
	Openid                   string `json:"openid,omitempty"`                    // 微信用户在商户对应appid下的唯一标识
	AuthorizationCode        string `json:"authorization_code"`                  // 预授权成功时的授权协议号。
	AuthorizationState       string `json:"authorization_state"`                 // 标识用户授权服务情况：UNAVAILABLE：用户未授权服务，AVAILABLE：用户已授权服务，UNBINDUSER：未绑定用户（已经预授权但未完成正式授权）
	CancelAuthorizationTime  string `json:"cancel_authorization_time,omitempty"` // 最近一次解除授权时间, 示例值：2015-05-20T13:29:35.120+08:00
	AuthorizationSuccessTime string `json:"authorization_success_time"`          // 最近一次授权成功时间, 示例值：2015-05-20T13:29:35.120+08:00
}

type ProfitShareOrder struct {
	SubMchid      string                   `json:"sub_mchid"`           // 二级商户号
	TransactionId string                   `json:"transaction_id"`      // 微信订单号
	OutOrderNo    string                   `json:"out_order_no"`        // 商户分账单号
	OrderId       string                   `json:"order_id"`            // 微信分账单号
	State         string                   `json:"state"`               // 分账单状态（每个接收方的分账结果请查看receivers中的result字段）:PROCESSING：处理中,FINISHED：分账完成
	Receivers     []*ProfitSharingReceiver `json:"receivers,omitempty"` // 分账接收方列表
}

type ProfitShareOrderQuery struct {
	SubMchid      string                   `json:"sub_mchid,omitempty"` // 子商户号，即分账的出资商户号【服务商模式】
	TransactionId string                   `json:"transaction_id"`      // 微信订单号
	OutOrderNo    string                   `json:"out_order_no"`        // 商户分账单号
	OrderId       string                   `json:"order_id"`            // 微信分账单号
	State         string                   `json:"state"`               // 分账单状态（每个接收方的分账结果请查看receivers中的result字段）:PROCESSING：处理中,FINISHED：分账完成
	Receivers     []*ProfitSharingReceiver `json:"receivers,omitempty"` // 分账接收方列表
}

type ProfitShareReturn struct {
	SubMchid    string `json:"sub_mchid"`             // 子商户号，分账回退的接收商户，对应原分账出资的分账方商户【服务商模式】
	OrderId     string `json:"order_id"`              // 微信分账单号，微信系统返回的唯一标识
	OutOrderNo  string `json:"out_order_no"`          // 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OutReturnNo string `json:"out_return_no"`         // 此回退单号是商户在自己后台生成的一个新的回退单号，在商户后台唯一
	ReturnId    string `json:"return_id"`             // 微信分账回退单号，微信系统返回的唯一标识
	ReturnMchid string `json:"return_mchid"`          // 只能对原分账请求中成功分给商户接收方进行回退
	Amount      int    `json:"amount"`                // 需要从分账接收方回退的金额，单位为分，只能为整数
	Description string `json:"description"`           // 分账回退的原因描述
	Result      string `json:"result"`                // 回退结果: PROCESSING：处理中，SUCCESS：已成功，FAILED：已失败
	FailReason  string `json:"fail_reason,omitempty"` // 失败原因: ACCOUNT_ABNORMAL : 分账接收方账户异常，TIME_OUT_CLOSED : 超时关单
	CreateTime  string `json:"create_time"`           // 分账回退创建时间
	FinishTime  string `json:"finish_time"`           // 分账回退完成时间
}

type ProfitShareReturnResult struct {
	SubMchid    string `json:"sub_mchid,omitempty"`   // 子商户号，分账回退的接收商户，对应原分账出资的分账方商户【服务商模式】
	OrderId     string `json:"order_id"`              // 微信分账单号，微信系统返回的唯一标识
	OutOrderNo  string `json:"out_order_no"`          // 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OutReturnNo string `json:"out_return_no"`         // 此回退单号是商户在自己后台生成的一个新的回退单号，在商户后台唯一
	ReturnId    string `json:"return_id"`             // 微信分账回退单号，微信系统返回的唯一标识
	ReturnMchid string `json:"return_mchid"`          // 只能对原分账请求中成功分给商户接收方进行回退
	Amount      int    `json:"amount"`                // 需要从分账接收方回退的金额，单位为分，只能为整数
	Description string `json:"description"`           // 分账回退的原因描述
	Result      string `json:"result"`                // 回退结果: PROCESSING：处理中，SUCCESS：已成功，FAILED：已失败
	FailReason  string `json:"fail_reason,omitempty"` // 失败原因: ACCOUNT_ABNORMAL : 分账接收方账户异常，TIME_OUT_CLOSED : 超时关单
	CreateTime  string `json:"create_time"`           // 分账回退创建时间
	FinishTime  string `json:"finish_time"`           // 分账回退完成时间
}

type ProfitShareOrderUnfreeze struct {
	SubMchid      string                   `json:"sub_mchid,omitempty"` // 子商户号，分账回退的接收商户，对应原分账出资的分账方商户【服务商模式】
	TransactionId string                   `json:"transaction_id"`      // 微信支付订单号
	OutOrderNo    string                   `json:"out_order_no"`        // 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OrderId       string                   `json:"order_id"`            // 微信分账单号，微信系统返回的唯一标识
	State         string                   `json:"state"`               // 分账单状态（每个接收方的分账结果请查看receivers中的result字段）:PROCESSING：处理中,FINISHED：分账完成
	Receivers     []*ProfitSharingReceiver `json:"receivers,omitempty"` // 分账接收方列表
}

type ProfitShareUnsplitAmount struct {
	TransactionId string `json:"transaction_id"` // 微信支付订单号
	UnsplitAmount int    `json:"unsplit_amount"` // 订单剩余待分金额，整数，单位为分
}

// 分账接收方
type ProfitSharingReceiver struct {
	Amount      int    `json:"amount"`      // 分账金额
	Description string `json:"description"` // 分账描述
	Type        string `json:"type"`        // 分账接收方类型
	Account     string `json:"account"`     // 分账接收方帐号
	Result      string `json:"result"`      // 分账结果,PENDING：待分账,SUCCESS：分账成功,CLOSED：已关闭
	FailReason  string `json:"fail_reason"` // 分账失败原因ACCOUNT_ABNORMAL : 分账接收账户异常、NO_RELATION : 分账关系已解除、RECEIVER_HIGH_RISK : 高风险接收方、RECEIVER_REAL_NAME_NOT_VERIFIED : 接收方未实名、NO_AUTH : 分账权限已解除
	CreateTime  string `json:"create_time"` // 分账创建时间,遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE
	FinishTime  string `json:"finish_time"` // 分账完成时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE
	DetailId    string `json:"detail_id"`   // 分账明细单号, 微信分账明细单号，每笔分账业务执行的明细单号，可与资金账单对账使用
}

type ProfitShareAddReceiver struct {
	SubMchid       string `json:"sub_mchid,omitempty"`       // 子商户号，分账回退的接收商户，对应原分账出资的分账方商户【服务商模式】
	Type           string `json:"type"`                      // 分账接收方类型MERCHANT_ID：商户ID,PERSONAL_OPENID：个人openid（由父商户APPID转换得到）
	Account        string `json:"account"`                   // 分账接收方帐号
	Name           string `json:"name,omitempty"`            // 分账接收方类型是MERCHANT_ID时，是商户全称（必传），当商户是小微商户或个体户时，是开户人姓名 分账接收方类型是PERSONAL_OPENID时，是个人姓名（选传，传则校验）
	RelationType   string `json:"relation_type"`             // 商户与接收方的关系。STORE：门店STAFF：员工	STORE_OWNER：店主	PARTNER：合作伙伴	HEADQUARTER：总部	BRAND：品牌方	DISTRIBUTOR：分销商	USER：用户	SUPPLIER： 供应商	CUSTOM：自定义
	CustomRelation string `json:"custom_relation,omitempty"` // 子商户与接收方具体的关系，本字段最多10个字。当字段relation_type的值为CUSTOM时，本字段必填;当字段relation_type的值不为CUSTOM时，本字段无需填写。
}

type ProfitShareDeleteReceiver struct {
	SubMchid string `json:"sub_mchid,omitempty"` // 子商户号，分账回退的接收商户，对应原分账出资的分账方商户【服务商模式】
	Type     string `json:"type"`                // 分账接收方类型MERCHANT_ID：商户ID,PERSONAL_OPENID：个人openid（由父商户APPID转换得到）
	Account  string `json:"account"`             // 分账接收方帐号
}

type ProfitShareMerchantConfigs struct {
	SubMchId string `json:"sub_mchid"` // 子商户号
	MaxRatio int    `json:"max_ratio"` // 最大分账比例 (单位万分比，比如2000表示20%)
}

type ProfitShareBills struct {
	DownloadUrl string `json:"download_url"` // 下载地址	原始账单（gzip需要解压缩）的摘要值，用于校验文件的完整性。
	HashType    string `json:"hash_type"`    // 哈希类型 原始账单（gzip需要解压缩）的摘要算法，用于校验文件的完整性
	HashValue   string `json:"hash_value"`   // 哈希值	供下一步请求账单文件的下载地址，该地址30s内有效。
}

type PayInformation struct {
	PayAmount     int    `json:"pay_amount"`               // 用户需要退回优惠而付款的金额，单位为：分
	PayState      string `json:"pay_state"`                // 用户付款状态：PAYING：付款中，PAID：已付款
	TransactionId string `json:"transaction_id,omitempty"` // 微信支付订单号，仅在订单成功收款时才返回
	PayTime       string `json:"pay_time,omitempty"`       // 用户成功支付的时间，仅在订单成功收款时才返回
}

type Objective struct {
	ObjectiveId                string                       `json:"objective_id"`                           // 由先享卡平台生成，唯一标识一个先享卡目标。商户需要记录该目标Id，进行同步用户记录
	Name                       string                       `json:"name"`                                   // 目标的名称
	Count                      int                          `json:"count"`                                  // 履约目标需要完成的数量，必须大于0
	Unit                       string                       `json:"unit"`                                   // 目标的单位
	Description                string                       `json:"description"`                            // 对先享卡目标的补充信息
	ObjectiveCompletionRecords []*ObjectiveCompletionRecord `json:"objective_completion_records,omitempty"` // 用户完成的目标明细列表
}

type ObjectiveCompletionRecord struct {
	ObjectiveCompletionSerialNo string `json:"objective_completion_serial_no"` // 目标完成流水号
	ObjectiveId                 string `json:"objective_id"`                   // 微信先享卡为每个先享卡目标分配的唯一Id
	CompletionTime              string `json:"completion_time"`                // 用户履约行为发生的时间
	CompletionType              string `json:"completion_type"`                // 目标完成类型： INCREASE：增加数量，DECREASE：减少数量
	Description                 string `json:"description"`                    // 用户本次履约的描述
	CompletionCount             int    `json:"completion_count"`               // 用户本次履约的数量，必须大于0
	Remark                      string `json:"remark,omitempty"`               // 对于用户履约情况的一些补充信息
}

type Reward struct {
	RewardId           string               `json:"reward_id"`                      // 由先享卡平台生成，唯一标识一个先享卡目标。商户需要记录该优惠Id，进行同步用户记录
	Name               string               `json:"name"`                           // 优惠名称
	CountType          string               `json:"count_type"`                     // 优惠数量的类型标识：COUNT_UNLIMITED：不限数量，COUNT_LIMIT：有限数量
	Count              int                  `json:"count"`                          // 本项优惠可使用的数量，必须大于0
	Unit               string               `json:"unit"`                           // 优惠的单位
	Amount             int                  `json:"amount"`                         // 优惠金额，此项优惠对应的优惠总金额，单位：分，必须大于0
	Description        string               `json:"description,omitempty"`          // 对先享卡优惠的补充信息
	RewardUsageRecords []*RewardUsageRecord `json:"reward_usage_records,omitempty"` // 优惠使用记录列表
}

type RewardUsageRecord struct {
	RewardUsageSerialNo string `json:"reward_usage_serial_no"` // 优惠使用记录流水号
	RewardId            string `json:"reward_id"`              // 微信先享卡为每个先享卡优惠分配的唯一Id
	UsageTime           string `json:"usage_time"`             // 用户使用优惠的时间
	UsageType           string `json:"usage_type"`             // 目标完成类型：INCREASE：增加数量，DECREASE：减少数量
	Description         string `json:"description"`            // 用户获得奖励的描述
	UsageCount          int    `json:"usage_count"`            // 用户本次获得的奖励数量，必须大于0
	Amount              int    `json:"amount"`                 // 优惠金额，用户此项本次享受的优惠对应的优惠总金额，单位：分，必须大于0
	Remark              string `json:"remark,omitempty"`       // 对于用户奖励情况的一些补充信息
}

type SmartGuideReg struct {
	GuideId string `json:"guide_id"` // 服务人员在服务人员系统中的唯一标识
}

type SmartGuideQuery struct {
	Data       []*SmartGuide `json:"data"`        // 服务人员列表
	TotalCount int           `json:"total_count"` // 服务人员数量
	Limit      int           `json:"limit"`       // 该次请求可返回的最大资源条数，不大于10
	Offset     int           `json:"offset"`      // 该次请求资源的起始位置，默认值为0
}

type SmartGuide struct {
	GuideId string `json:"guide_id"`          // 服务人员在服务人员系统中的唯一标识
	StoreId int    `json:"store_id"`          // 门店在微信支付商户平台的唯一标识
	Name    string `json:"name"`              // 服务人员姓名
	Mobile  string `json:"mobile"`            // 员工在商户个人/企业微信通讯录上设置的手机号码（加密信息，需解密）
	Userid  string `json:"userid,omitempty"`  // 员工在商户企业微信通讯录使用的唯一标识，使用企业微信商家时返回
	WorkId  string `json:"work_id,omitempty"` // 服务人员通过小程序注册时填写的工号，使用个人微信商家时返回
}

type Transfer struct {
	OutBatchNo string `json:"out_batch_no"` // 商户系统内部的商家批次单号
	BatchId    string `json:"batch_id"`     // 微信批次单号，微信商家转账系统返回的唯一标识
	CreateTime string `json:"create_time"`  // 批次受理成功时返回
}

type TransferQuery struct {
	TransferBatch      *TransferBatch    `json:"transfer_batch"`                 // 转账批次单基本信息
	TransferDetailList []*TransferDetail `json:"transfer_detail_list,omitempty"` // 当批次状态为“FINISHED”（已完成），且成功查询到转账明细单时返回
}

type TransferBatch struct {
	Mchid         string `json:"mchid"`                    // 微信支付分配的商户号
	OutBatchNo    string `json:"out_batch_no"`             // 商户系统内部的商家批次单号
	BatchId       string `json:"batch_id"`                 // 微信批次单号，微信商家转账系统返回的唯一标识
	Appid         string `json:"appid"`                    // 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	BatchStatus   string `json:"batch_status"`             // 批次状态
	BatchType     string `json:"batch_type"`               // 批次类型
	BatchName     string `json:"batch_name"`               // 该笔批量转账的名称
	BatchRemark   string `json:"batch_remark"`             // 转账说明，UTF8编码，最多允许32个字符
	CloseReason   string `json:"close_reason,omitempty"`   // 如果批次单状态为“CLOSED”（已关闭），则有关闭原因
	TotalAmount   int    `json:"total_amount"`             // 转账金额单位为分
	TotalNum      int    `json:"total_num"`                // 一个转账批次单最多发起三千笔转账
	CreateTime    string `json:"create_time,omitempty"`    // 批次受理成功时返回
	UpdateTime    string `json:"update_time,omitempty"`    // 批次最近一次状态变更的时间
	SuccessAmount int    `json:"success_amount,omitempty"` // 转账成功的金额，单位为分
	SuccessNum    int    `json:"success_num,omitempty"`    // 转账成功的笔数
	FailAmount    int    `json:"fail_amount,omitempty"`    // 转账失败的金额，单位为分
	FailNum       int    `json:"fail_num,omitempty"`       // 转账失败的笔数
}

type TransferDetail struct {
	DetailId     string `json:"detail_id"`     // 微信明细单号
	OutDetailNo  string `json:"out_detail_no"` // 商家明细单号
	DetailStatus string `json:"detail_status"` // 明细状态：PROCESSING：转账中，SUCCESS：转账成功，FAIL：转账失败
}

type PartnerTransferQuery struct {
	SpMchid            string            `json:"sp_mchid"`                       // 微信支付分配的服务商商户号
	SubMchid           string            `json:"sub_mchid"`                      // 微信支付分配的特约商户号
	OutBatchNo         string            `json:"out_batch_no"`                   // 商户系统内部的商家批次单号
	BatchId            string            `json:"batch_id"`                       // 微信批次单号，微信商家转账系统返回的唯一标识
	SpAppid            string            `json:"sp_appid,omitempty"`             // 微信分配的服务商商户公众账号Id，特约商户授权类型为FUND_AUTHORIZATION_TYPE时才有该字段
	SubAppid           string            `json:"sub_appid"`                      // 微信分配的特约商户公众账号Id。特约商户appid
	BatchStatus        string            `json:"batch_status"`                   // 批次状态
	BatchType          string            `json:"batch_type"`                     // 批次类型
	AuthorizationType  string            `json:"authorization_type"`             // 特约商户授权类型
	BatchName          string            `json:"batch_name"`                     // 该笔批量转账的名称
	BatchRemark        string            `json:"batch_remark"`                   // 转账说明，UTF8编码，最多允许32个字符
	CloseReason        string            `json:"close_reason,omitempty"`         // 如果批次单状态为“CLOSED”（已关闭），则有关闭原因
	TotalAmount        int               `json:"total_amount"`                   // 转账金额单位为分
	TotalNum           int               `json:"total_num"`                      // 一个转账批次单最多发起三千笔转账
	CreateTime         string            `json:"create_time,omitempty"`          // 批次受理成功时返回
	UpdateTime         string            `json:"update_time,omitempty"`          // 批次最近一次状态变更的时间
	SuccessAmount      int               `json:"success_amount,omitempty"`       // 转账成功的金额，单位为分
	SuccessNum         int               `json:"success_num,omitempty"`          // 转账成功的笔数
	FailAmount         int               `json:"fail_amount,omitempty"`          // 转账失败的金额，单位为分
	FailNum            int               `json:"fail_num,omitempty"`             // 转账失败的笔数
	TransferPurpose    string            `json:"transfer_purpose"`               // 批量转账用途
	TransferDetailList []*TransferDetail `json:"transfer_detail_list,omitempty"` // 当批次状态为“FINISHED”（已完成），且成功查询到转账明细单时返回
}

type TransferDetailQuery struct {
	Mchid          string `json:"mchid"`                 // 微信支付分配的商户号
	OutBatchNo     string `json:"out_batch_no"`          // 商户系统内部的商家批次单号
	BatchId        string `json:"batch_id"`              // 微信批次单号，微信商家转账系统返回的唯一标识
	Appid          string `json:"appid"`                 // 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	OutDetailNo    string `json:"out_detail_no"`         // 商家明细单号
	DetailId       string `json:"detail_id"`             // 微信明细单号
	DetailStatus   string `json:"detail_status"`         // 明细状态：PROCESSING：转账中，SUCCESS：转账成功，FAIL：转账失败
	TransferAmount int    `json:"transfer_amount"`       // 转账金额单位为分
	TransferRemark string `json:"transfer_remark"`       // 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符
	FailReason     string `json:"fail_reason,omitempty"` // 如果转账失败则有失败原因
	Openid         string `json:"openid"`                // 用户在直连商户appid下的唯一标识
	UserName       string `json:"user_name"`             // 收款方姓名（加密）
	InitiateTime   string `json:"initiate_time"`         // 转账发起的时间
	UpdateTime     string `json:"update_time"`           // 明细最后一次状态变更的时间
}

type PartnerTransferDetail struct {
	SpMchid        string `json:"sp_mchid"`              // 微信支付分配的服务商商户号
	OutBatchNo     string `json:"out_batch_no"`          // 商户系统内部的商家批次单号
	BatchId        string `json:"batch_id"`              // 微信批次单号，微信商家转账系统返回的唯一标识
	Appid          string `json:"appid"`                 // 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	OutDetailNo    string `json:"out_detail_no"`         // 商家明细单号
	DetailId       string `json:"detail_id"`             // 微信明细单号
	DetailStatus   string `json:"detail_status"`         // 明细状态：PROCESSING：转账中，SUCCESS：转账成功，FAIL：转账失败
	TransferAmount int    `json:"transfer_amount"`       // 转账金额单位为分
	TransferRemark string `json:"transfer_remark"`       // 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符
	FailReason     string `json:"fail_reason,omitempty"` // 如果转账失败则有失败原因
	Openid         string `json:"openid"`                // 用户在直连商户appid下的唯一标识
	Username       string `json:"username"`              // 收款方姓名（加密）
	InitiateTime   string `json:"initiate_time"`         // 转账发起的时间
	UpdateTime     string `json:"update_time"`           // 明细最后一次状态变更的时间
}

type TransferMerchantQuery struct {
	TransferBatch      *TransferBatch    `json:"transfer_batch"`                 // 转账批次单基本信息
	TransferDetailList []*TransferDetail `json:"transfer_detail_list,omitempty"` // 当批次状态为“FINISHED”（已完成），且成功查询到转账明细单时返回
	Offset             int               `json:"offset,omitempty"`               // 该次请求资源（转账明细单）的起始位置
	Limit              int               `json:"limit,omitempty"`                // 该次请求可返回的最大资源（转账明细单）条数
}

type PartnerTransferMerchantQuery struct {
	SpMchid            string            `json:"sp_mchid"`                       // 微信支付分配的服务商商户号
	SubMchid           string            `json:"sub_mchid"`                      // 微信支付分配的特约商户号
	OutBatchNo         string            `json:"out_batch_no"`                   // 商户系统内部的商家批次单号
	BatchId            string            `json:"batch_id"`                       // 微信批次单号，微信商家转账系统返回的唯一标识
	SpAppid            string            `json:"sp_appid,omitempty"`             // 微信分配的服务商商户公众账号Id，特约商户授权类型为FUND_AUTHORIZATION_TYPE时才有该字段
	SubAppid           string            `json:"sub_appid"`                      // 微信分配的特约商户公众账号Id。特约商户appid
	BatchStatus        string            `json:"batch_status"`                   // 批次状态
	BatchType          string            `json:"batch_type"`                     // 批次类型
	AuthorizationType  string            `json:"authorization_type"`             // 特约商户授权类型
	BatchName          string            `json:"batch_name"`                     // 该笔批量转账的名称
	BatchRemark        string            `json:"batch_remark"`                   // 转账说明，UTF8编码，最多允许32个字符
	CloseReason        string            `json:"close_reason,omitempty"`         // 如果批次单状态为“CLOSED”（已关闭），则有关闭原因
	TotalAmount        int               `json:"total_amount"`                   // 转账金额单位为分
	TotalNum           int               `json:"total_num"`                      // 一个转账批次单最多发起三千笔转账
	CreateTime         string            `json:"create_time,omitempty"`          // 批次受理成功时返回
	UpdateTime         string            `json:"update_time,omitempty"`          // 批次最近一次状态变更的时间
	SuccessAmount      int               `json:"success_amount,omitempty"`       // 转账成功的金额，单位为分
	SuccessNum         int               `json:"success_num,omitempty"`          // 转账成功的笔数
	FailAmount         int               `json:"fail_amount,omitempty"`          // 转账失败的金额，单位为分
	FailNum            int               `json:"fail_num,omitempty"`             // 转账失败的笔数
	TransferPurpose    string            `json:"transfer_purpose"`               // 批量转账用途
	TransferDetailList []*TransferDetail `json:"transfer_detail_list,omitempty"` // 当批次状态为“FINISHED”（已完成），且成功查询到转账明细单时返回
}

type TransferMerchantDetail struct {
	OutBatchNo     string `json:"out_batch_no"`          // 商户系统内部的商家批次单号
	BatchId        string `json:"batch_id"`              // 微信批次单号，微信商家转账系统返回的唯一标识
	Appid          string `json:"appid"`                 // 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	OutDetailNo    string `json:"out_detail_no"`         // 商家明细单号
	DetailId       string `json:"detail_id"`             // 微信明细单号
	DetailStatus   string `json:"detail_status"`         // 明细状态：PROCESSING：转账中，SUCCESS：转账成功，FAIL：转账失败
	TransferAmount int    `json:"transfer_amount"`       // 转账金额单位为分
	TransferRemark string `json:"transfer_remark"`       // 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符
	FailReason     string `json:"fail_reason,omitempty"` // 如果转账失败则有失败原因
	Openid         string `json:"openid"`                // 用户在直连商户appid下的唯一标识
	UserName       string `json:"user_name"`             // 收款方姓名（加密）
	InitiateTime   string `json:"initiate_time"`         // 转账发起的时间
	UpdateTime     string `json:"update_time"`           // 明细最后一次状态变更的时间
}

type PartnerTransferMerchantDetail struct {
	SpMchid        string `json:"sp_mchid"`              // 微信支付分配的服务商商户号
	OutBatchNo     string `json:"out_batch_no"`          // 商户系统内部的商家批次单号
	BatchId        string `json:"batch_id"`              // 微信批次单号，微信商家转账系统返回的唯一标识
	Appid          string `json:"appid"`                 // 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	OutDetailNo    string `json:"out_detail_no"`         // 商家明细单号
	DetailId       string `json:"detail_id"`             // 微信明细单号
	DetailStatus   string `json:"detail_status"`         // 明细状态：PROCESSING：转账中，SUCCESS：转账成功，FAIL：转账失败
	TransferAmount int    `json:"transfer_amount"`       // 转账金额单位为分
	TransferRemark string `json:"transfer_remark"`       // 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符
	FailReason     string `json:"fail_reason,omitempty"` // 如果转账失败则有失败原因
	Openid         string `json:"openid"`                // 用户在直连商户appid下的唯一标识
	Username       string `json:"username"`              // 收款方姓名（加密）
	InitiateTime   string `json:"initiate_time"`         // 转账发起的时间
	UpdateTime     string `json:"update_time"`           // 明细最后一次状态变更的时间
}

type TransferReceipt struct {
	OutBatchNo      string `json:"out_batch_no"`               // 商户系统内部的商家批次单号
	SignatureNo     string `json:"signature_no"`               // 电子回单申请单号，申请单据的唯一标识
	SignatureStatus string `json:"signature_status,omitempty"` // 电子回单状态：ACCEPTED:已受理，电子签章已受理成功，FINISHED:已完成。电子签章已处理完成
	HashType        string `json:"hash_type,omitempty"`        // 电子回单文件的hash方法，回单状态为：FINISHED时返回。
	HashValue       string `json:"hash_value,omitempty"`       // 电子回单文件的hash值，用于下载之后验证文件的完整、正确性，回单状态为：FINISHED时返回。
	DownloadUrl     string `json:"download_url,omitempty"`     // 电子回单文件的下载地址，回单状态为：FINISHED时返回
	CreateTime      string `json:"create_time,omitempty"`      // 电子签章单创建时间
	UpdateTime      string `json:"update_time,omitempty"`      // 电子签章单最近一次状态变更的时间
}

type TransferReceiptQuery struct {
	OutBatchNo      string `json:"out_batch_no"`               // 商户系统内部的商家批次单号
	SignatureNo     string `json:"signature_no"`               // 电子回单申请单号，申请单据的唯一标识
	SignatureStatus string `json:"signature_status,omitempty"` // 电子回单状态：ACCEPTED:已受理，电子签章已受理成功，FINISHED:已完成。电子签章已处理完成
	HashType        string `json:"hash_type,omitempty"`        // 电子回单文件的hash方法，回单状态为：FINISHED时返回。
	HashValue       string `json:"hash_value,omitempty"`       // 电子回单文件的hash值，用于下载之后验证文件的完整、正确性，回单状态为：FINISHED时返回。
	DownloadUrl     string `json:"download_url,omitempty"`     // 电子回单文件的下载地址，回单状态为：FINISHED时返回
	CreateTime      string `json:"create_time,omitempty"`      // 电子签章单创建时间
	UpdateTime      string `json:"update_time,omitempty"`      // 电子签章单最近一次状态变更的时间
}

type TransferDetailReceipt struct {
	AcceptType      string `json:"accept_type"`                // 电子回单受理类型
	OutBatchNo      string `json:"out_batch_no,omitempty"`     // 商户系统内部的商家批次单号
	OutDetailNo     string `json:"out_detail_no"`              // 商家明细单号
	SignatureNo     string `json:"signature_no"`               // 电子回单申请单号，申请单据的唯一标识
	SignatureStatus string `json:"signature_status,omitempty"` // 电子回单状态：ACCEPTED:已受理，电子签章已受理成功，FINISHED:已完成。电子签章已处理完成
	HashType        string `json:"hash_type,omitempty"`        // 电子回单文件的hash方法，回单状态为：FINISHED时返回。
	HashValue       string `json:"hash_value,omitempty"`       // 电子回单文件的hash值，用于下载之后验证文件的完整、正确性，回单状态为：FINISHED时返回。
	DownloadUrl     string `json:"download_url,omitempty"`     // 电子回单文件的下载地址，回单状态为：FINISHED时返回
}

type TransferDetailReceiptQuery struct {
	AcceptType      string `json:"accept_type"`                // 电子回单受理类型
	OutBatchNo      string `json:"out_batch_no,omitempty"`     // 商户系统内部的商家批次单号
	OutDetailNo     string `json:"out_detail_no"`              // 商家明细单号
	SignatureNo     string `json:"signature_no"`               // 电子回单申请单号，申请单据的唯一标识
	SignatureStatus string `json:"signature_status,omitempty"` // 电子回单状态：ACCEPTED:已受理，电子签章已受理成功，FINISHED:已完成。电子签章已处理完成
	HashType        string `json:"hash_type,omitempty"`        // 电子回单文件的hash方法，回单状态为：FINISHED时返回。
	HashValue       string `json:"hash_value,omitempty"`       // 电子回单文件的hash值，用于下载之后验证文件的完整、正确性，回单状态为：FINISHED时返回。
	DownloadUrl     string `json:"download_url,omitempty"`     // 电子回单文件的下载地址，回单状态为：FINISHED时返回
}

type EcommerceBalance struct {
	SubMchid        string `json:"sub_mchid"`                // 特约商户号
	AccountType     string `json:"account_type,omitempty"`   // 账户类型
	AvailableAmount int    `json:"available_amount"`         // 可用余额（单位：分），此余额可做提现操作
	PendingAmount   int    `json:"pending_amount,omitempty"` // 不可用余额（单位：分）
}

type MerchantBalance struct {
	AvailableAmount int `json:"available_amount"`         // 可用余额（单位：分），此余额可做提现操作
	PendingAmount   int `json:"pending_amount,omitempty"` // 不可用余额（单位：分）
}

type MerchantIncomeRecord struct {
	Data       []*IncomeData `json:"data,omitempty"` // 单次查询返回的银行来账记录列表结果数组，如果查询结果为空时，则为空数组
	Links      *Link         `json:"links"`          // 返回前后页和当前页面的访问链接
	Offset     int           `json:"offset"`         // 该次请求资源的起始位置，请求中包含偏移量时应答消息返回相同偏移量，否则返回默认值0
	Limit      int           `json:"limit"`          // 经过条件筛选，本次查询到的银行来账记录条数
	TotalCount int           `json:"total_count"`    // 经过条件筛选，查询到的银行来账记录总数
}

type IncomeData struct {
	Mchid             string `json:"mchid"`               // 微信支付分配的商户号
	AccountType       string `json:"account_type"`        // 需查询银行来账记录商户的账户类型：BASIC：基本账户，OPERATION：运营账户，FEES：手续费账户
	IncomeRecordType  string `json:"income_record_type"`  // 银行来账类型
	IncomeRecordId    string `json:"income_record_id"`    // 银行来账的微信单号
	Amount            int    `json:"amount"`              // 银行来账金额，单位为分，只能为整数
	SuccessTime       string `json:"success_time"`        // 银行来账完成时间
	BankName          string `json:"bank_name"`           // 银行来账的付款方银行名称，由于部分银行的数据获取限制，该字段有可能为空
	BankAccountName   string `json:"bank_account_name"`   // 银行来账的付款方银行账户信息，户名为全称、明文，由于部分银行的数据获取限制，该字段有可能为空
	BankAccountNumber string `json:"bank_account_number"` // 四位掩码+付款方银行卡尾号后四位
	RechargeRemark    string `json:"recharge_remark"`     // 银行备注
}

type Link struct {
	Next string `json:"next"` // 下一页链接
	Prev string `json:"prev"` // 上一页链接
	Self string `json:"self"` // 当前链接
}

type PartnerIncomeRecord struct {
	Data       []*PartnerIncomeData `json:"data,omitempty"` // 单次查询返回的银行来账记录列表结果数组，如果查询结果为空时，则为空数组
	Links      *Link                `json:"links"`          // 返回前后页和当前页面的访问链接
	Offset     int                  `json:"offset"`         // 该次请求资源的起始位置，请求中包含偏移量时应答消息返回相同偏移量，否则返回默认值0
	Limit      int                  `json:"limit"`          // 经过条件筛选，本次查询到的银行来账记录条数
	TotalCount int                  `json:"total_count"`    // 经过条件筛选，查询到的银行来账记录总数
}

type PartnerIncomeData struct {
	SubMchid          string `json:"sub_mchid"`           // 特约商户号
	AccountType       string `json:"account_type"`        // 需查询银行来账记录商户的账户类型：BASIC：基本账户，OPERATION：运营账户，FEES：手续费账户
	IncomeRecordType  string `json:"income_record_type"`  // 银行来账类型
	IncomeRecordId    string `json:"income_record_id"`    // 银行来账的微信单号
	Amount            int    `json:"amount"`              // 银行来账金额，单位为分，只能为整数
	SuccessTime       string `json:"success_time"`        // 银行来账完成时间
	BankName          string `json:"bank_name"`           // 银行来账的付款方银行名称，由于部分银行的数据获取限制，该字段有可能为空
	BankAccountName   string `json:"bank_account_name"`   // 银行来账的付款方银行账户信息，户名为全称、明文，由于部分银行的数据获取限制，该字段有可能为空
	BankAccountNumber string `json:"bank_account_number"` // 四位掩码+付款方银行卡尾号后四位
	RechargeRemark    string `json:"recharge_remark"`     // 银行备注
}

type AuditDetail struct {
	ParamName    string `json:"param_name"`
	RejectReason string `json:"reject_reason"`
}

type GoldPlanManage struct {
	SubMchid string `json:"sub_mchid"`
}

type Withdraw struct {
	SubMchid    string `json:"sub_mchid"`              // 服务商特约商户号，由微信支付生成并下发。
	WithdrawId  string `json:"withdraw_id"`            // 微信支付提现单号
	AccountType string `json:"account_type,omitempty"` // 出款账户类型
}

type WithdrawStatus struct {
	SubMchid      string `json:"sub_mchid"`      // 服务商特约商户号，由微信支付生成并下发。
	SpMchid       string `json:"sp_mchid"`       // 服务商户号
	Status        string `json:"status"`         // 提现单状态：CREATE_SUCCESS：受理成功，SUCCESS：提现成功，FAIL：提现失败，REFUND：提现退票，CLOSE：关单，INIT：业务单已创建
	WithdrawId    string `json:"withdraw_id"`    // 微信支付提现单号
	OutRequestNo  string `json:"out_request_no"` // 商户提现单号
	Amount        int    `json:"amount"`         // 提现金额
	CreateTime    string `json:"create_time"`    // 创建时间
	UpdateTime    string `json:"update_time"`    // 更新时间
	Reason        string `json:"reason"`         // 提现失败原因，仅在提现失败、退票、关单时有值
	Remark        string `json:"remark"`         // 商户对提现单的备注，若发起提现时未传入相应值或输入不合法，则该值为空
	BankMemo      string `json:"bank_memo"`      // 展示在收款银行系统中的附言，由数字、字母、汉字组成（能否成功展示依赖银行系统支持）。若发起提现时未传入相应值或输入不合法，则该值为空
	AccountType   string `json:"account_type"`   // 出款账户类型
	AccountNumber string `json:"account_number"` // 服务商提现入账的银行账号，仅显示后四位。
	AccountBank   string `json:"account_bank"`   // 服务商提现入账的开户银行
	BankName      string `json:"bank_name"`      // 服务商提现入账的开户银行全称（含支行）
}

type VehicleParkingQuery struct {
	PlateNumber     string `json:"plate_number"` // 车牌号，仅包括省份+车牌，不包括特殊字符。
	PlateColor      string `json:"plate_color"`  // 车牌颜色，BLUE：蓝色，GREEN：绿色，YELLOW：黄色，BLACK：黑色，WHITE：白色，LIMEGREEN：黄绿色
	ServiceOpenTime string `json:"service_open_time,omitempty"`
	Openid          string `json:"openid"`
	ServiceState    string `json:"service_state"`
}

type VehicleParkingIn struct {
	Id           string `json:"id"`
	OutParkingNo string `json:"out_parking_no"` // 商户入场id
	PlateNumber  string `json:"plate_number"`   // 车牌号，仅包括省份+车牌，不包括特殊字符。
	PlateColor   string `json:"plate_color"`    // 车牌颜色，BLUE：蓝色，GREEN：绿色，YELLOW：黄色，BLACK：黑色，WHITE：白色，LIMEGREEN：黄绿色
	StartTime    string `json:"start_time"`     // 入场时间
	ParkingName  string `json:"parking_name"`   // 所在停车位车场的名称
	FreeDuration int    `json:"free_duration"`  // 停车场的免费停车时长，单位为秒
	State        string `json:"state"`
	BlockReason  string `json:"block_reason"`
}

type VehicleParkingFee struct {
	Appid                 string             `json:"appid"`
	SubAppid              string             `json:"sub_appid,omitempty"`
	SpMchid               string             `json:"sp_mchid"`
	SubMchid              string             `json:"sub_mchid,omitempty"`
	Description           string             `json:"description"`
	CreateTime            string             `json:"create_time"`
	OutTradeNo            string             `json:"out_trade_no"`
	TransactionId         string             `json:"transaction_id"`
	TradeState            string             `json:"trade_state"`
	TradeStateDescription string             `json:"trade_state_description"`
	SuccessTime           string             `json:"success_time"`
	BankType              string             `json:"bank_type"`
	UserRepaid            string             `json:"user_repaid"`
	Attach                string             `json:"attach"`
	TradeScene            string             `json:"trade_scene"`
	ParkingInfo           *ParkingInfo       `json:"parking_info"`
	Payer                 *Payer             `json:"payer"`
	Amount                *Amount            `json:"amount"`
	PromotionDetail       []*PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能，享受优惠时返回该字段
}

type ParkingInfo struct {
	ParkingId        string `json:"parking_id"`        // 停车入场id
	PlateNumber      string `json:"plate_number"`      // 车牌号，仅包括省份+车牌，不包括特殊字符。
	PlateColor       string `json:"plate_color"`       // 车牌颜色，BLUE：蓝色，GREEN：绿色，YELLOW：黄色，BLACK：黑色，WHITE：白色，LIMEGREEN：黄绿色
	StartTime        string `json:"start_time"`        // 入场时间
	EndTime          string `json:"end_time"`          // 出场时间
	ParkingName      string `json:"parking_name"`      // 所在停车位车场的名称
	ChargingDuration int    `json:"charging_duration"` // 计费的时间长，单位为秒
	DeviceId         string `json:"device_id"`         // 停车场设备id
}

type VehicleParkingOrder struct {
	Appid                 string             `json:"appid"`
	SubAppid              string             `json:"sub_appid,omitempty"`
	SpMchid               string             `json:"sp_mchid"`
	SubMchid              string             `json:"sub_mchid,omitempty"`
	Description           string             `json:"description"`
	CreateTime            string             `json:"create_time"`
	OutTradeNo            string             `json:"out_trade_no"`
	TransactionId         string             `json:"transaction_id"`
	TradeState            string             `json:"trade_state"`
	TradeStateDescription string             `json:"trade_state_description"`
	SuccessTime           string             `json:"success_time"`
	BankType              string             `json:"bank_type"`
	UserRepaid            string             `json:"user_repaid"`
	Attach                string             `json:"attach"`
	TradeScene            string             `json:"trade_scene"`
	ParkingInfo           *ParkingInfo       `json:"parking_info"`
	Payer                 *Payer             `json:"payer"`
	Amount                *Amount            `json:"amount"`
	PromotionDetail       []*PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能，享受优惠时返回该字段
}
