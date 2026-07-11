package alipay

// =========================================================分割=========================================================

// 商品创建 Response
type TradeProductCreateResponse struct {
	Response     *TradeProductCreate `json:"alipay_trade_product_create_response"`
	AlipayCertSn string              `json:"alipay_cert_sn,omitempty"`
	SignData     string              `json:"-"`
	Sign         string              `json:"sign"`
}

// 商品修改 Response
type TradeProductModifyResponse struct {
	Response     *TradeProductModify `json:"alipay_trade_product_modify_response"`
	AlipayCertSn string              `json:"alipay_cert_sn,omitempty"`
	SignData     string              `json:"-"`
	Sign         string              `json:"sign"`
}

// 商品查询 Response
type TradeProductQueryResponse struct {
	Response     *TradeProductQuery `json:"alipay_trade_product_query_response"`
	AlipayCertSn string             `json:"alipay_cert_sn,omitempty"`
	SignData     string             `json:"-"`
	Sign         string             `json:"sign"`
}

// 价格创建 Response
type TradePriceCreateResponse struct {
	Response     *TradePriceCreate `json:"alipay_trade_price_create_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

// 价格查询 Response
type TradePriceQueryResponse struct {
	Response     *TradePriceQuery `json:"alipay_trade_price_query_response"`
	AlipayCertSn string           `json:"alipay_cert_sn,omitempty"`
	SignData     string           `json:"-"`
	Sign         string           `json:"sign"`
}

// 客户创建 Response
type TradeCustomerCreateResponse struct {
	Response     *TradeCustomerCreate `json:"alipay_trade_customer_create_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

// 订阅创建 Response
type TradeSubscriptionCreateResponse struct {
	Response     *TradeSubscriptionCreate `json:"alipay_trade_subscription_create_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

// 订阅修改 Response
type TradeSubscriptionModifyResponse struct {
	Response     *TradeSubscriptionModify `json:"alipay_trade_subscription_modify_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

// 订阅查询 Response
type TradeSubscriptionQueryResponse struct {
	Response     *TradeSubscriptionQuery `json:"alipay_trade_subscription_query_response"`
	AlipayCertSn string                  `json:"alipay_cert_sn,omitempty"`
	SignData     string                  `json:"-"`
	Sign         string                  `json:"sign"`
}

// =========================================================分割=========================================================

type TradeProductCreate struct {
	ErrorResponse
	ProductId string `json:"product_id,omitempty"` // 商品id
}

type TradeProductModify struct {
	ErrorResponse
	ProductId string `json:"product_id,omitempty"` // 商品id
}

type TradeProductQuery struct {
	ErrorResponse
	ProductList []*NexusPayProduct `json:"product_list,omitempty"` // 商品信息列表
}

type TradePriceCreate struct {
	ErrorResponse
	PriceId   string `json:"price_id,omitempty"`   // 价格实例id
	ProductId string `json:"product_id,omitempty"` // 商品id
}

type TradePriceQuery struct {
	ErrorResponse
	Id         string           `json:"id,omitempty"`          // 价格id
	Active     bool             `json:"active,omitempty"`      // 是否可用
	ProductId  string           `json:"product_id,omitempty"`  // 关联商品id
	UnitAmount int              `json:"unit_amount,omitempty"` // 单位金额，单位：分
	Type       string           `json:"type,omitempty"`        // 价格类型：recurring/one_time
	Metadata   string           `json:"metadata,omitempty"`    // 价格信息元数据
	GmtCreate  string           `json:"gmt_create,omitempty"`  // 创建时间
	Product    *NexusPayProduct `json:"product,omitempty"`     // 关联的商品信息
	Recurring  *RecurringConfig `json:"recurring,omitempty"`   // 循环计价配置
}

type TradeCustomerCreate struct {
	ErrorResponse
	CustomerId string `json:"customer_id,omitempty"` // 客户id
}

type TradeSubscriptionCreate struct {
	ErrorResponse
	PayAmount        int    `json:"pay_amount,omitempty"`         // 支付金额，单位分
	AlipayJumpSchema string `json:"alipay_jump_schema,omitempty"` // 支付宝长链跳端schema
	AlipaySchema     string `json:"alipay_schema,omitempty"`      // 支付宝短链schema
	SubscriptionId   string `json:"subscription_id,omitempty"`    // 订阅id
}

type TradeSubscriptionModify struct {
	ErrorResponse
	AlipayJumpSchema string `json:"alipay_jump_schema,omitempty"` // 支付宝长链跳端schema
	PayAmount        int    `json:"pay_amount,omitempty"`         // 支付金额，单位分
	SubscriptionId   string `json:"subscription_id,omitempty"`    // 订阅id
	AlipaySchema     string `json:"alipay_schema,omitempty"`      // 支付宝短链schema
}

type TradeSubscriptionQuery struct {
	ErrorResponse
	Subscriptions []*Subscription `json:"subscriptions,omitempty"` // 订阅详情信息
}

// =========================================================分割=========================================================

// NexusPayProduct 商品信息
type NexusPayProduct struct {
	Id             string         `json:"id,omitempty"`               // 商品id
	Name           string         `json:"name,omitempty"`             // 商品名称
	Description    string         `json:"description,omitempty"`      // 商品描述
	Active         bool           `json:"active,omitempty"`           // 是否可用
	DefaultPriceId string         `json:"default_price_id,omitempty"` // 默认价格id
	DefaultPrice   *NexusPayPrice `json:"default_price,omitempty"`    // 默认价格信息
	Metadata       string         `json:"metadata,omitempty"`         // 商品元数据
	GmtCreate      string         `json:"gmt_create,omitempty"`       // 创建时间
}

// NexusPayPrice 价格信息
type NexusPayPrice struct {
	Id         string           `json:"id,omitempty"`          // 价格id
	Active     bool             `json:"active,omitempty"`      // 是否可用
	Type       string           `json:"type,omitempty"`        // 价格类型
	UnitAmount int              `json:"unit_amount,omitempty"` // 单位金额，单位：分
	ProductId  string           `json:"product_id,omitempty"`  // 关联商品id
	Product    *NexusPayProduct `json:"product,omitempty"`     // 关联的商品信息
	Recurring  *RecurringConfig `json:"recurring,omitempty"`   // 循环计价配置
	Metadata   string           `json:"metadata,omitempty"`    // 价格信息元数据
	GmtCreate  string           `json:"gmt_create,omitempty"`  // 创建时间
}

// RecurringConfig 循环计价配置
type RecurringConfig struct {
	Interval        string `json:"interval,omitempty"`          // 计价周期：MONTH等
	IntervalCount   int    `json:"interval_count,omitempty"`    // 计价周期数
	UsageType       string `json:"usage_type,omitempty"`        // 使用类型
	TrialPeriodDays int    `json:"trial_period_days,omitempty"` // 试用天数
}

// Subscription 订阅信息
type Subscription struct {
	SubscriptionId     string              `json:"subscription_id,omitempty"`      // 订阅id
	CustomerId         string              `json:"customer_id,omitempty"`          // 客户id
	SubscribeTitle     string              `json:"subscribe_title,omitempty"`      // 订阅标题
	SubscriptionStatus string              `json:"subscription_status,omitempty"`  // 订阅状态
	CurrentPeriodStart string              `json:"current_period_start,omitempty"` // 当前周期开始时间
	CurrentPeriodEnd   string              `json:"current_period_end,omitempty"`   // 当前周期结束时间
	CancelAtPeriodEnd  bool                `json:"cancel_at_period_end,omitempty"` // 周期结束是否失效
	StartDate          string              `json:"start_date,omitempty"`           // 开始时间
	CanceledDate       string              `json:"canceled_date,omitempty"`        // 取消时间
	Created            string              `json:"created,omitempty"`              // 创建时间
	Items              []*SubscriptionItem `json:"items,omitempty"`                // 订阅项目信息
}

// SubscriptionItem 订阅项目
type SubscriptionItem struct {
	ItemId   string         `json:"item_id,omitempty"`  // 项目id
	PriceId  string         `json:"price_id,omitempty"` // 价格id
	Quantity int            `json:"quantity,omitempty"` // 数量
	Created  string         `json:"created,omitempty"`  // 创建时间
	Price    *NexusPayPrice `json:"price,omitempty"`    // 价格信息
}
