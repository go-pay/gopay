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

type JSAPIScoreQuery struct {
	MchId     string `json:"mch_id"`
	Package   string `json:"package"`
	TimeStamp string `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	SignType  string `json:"sign_type"`
	Sign      string `json:"sign"`
}

type APPScoreQuery struct {
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
	Data []*CertData `json:"data"`
}

type CertData struct {
	EffectiveTime      string       `json:"effective_time"`
	EncryptCertificate *EncryptCert `json:"encrypt_certificate"`
	ExpireTime         string       `json:"expire_time"`
	SerialNo           string       `json:"serial_no"`
}

type EncryptCert struct {
	Algorithm      string `json:"algorithm"`
	AssociatedData string `json:"associated_data"`
	Ciphertext     string `json:"ciphertext"`
	Nonce          string `json:"nonce"`
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
