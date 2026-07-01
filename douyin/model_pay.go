package douyin

// ===================== 下单响应 =====================

// PrepayRsp App / JSAPI 下单响应壳
type PrepayRsp struct {
	Code        int         `json:"-"`
	SignInfo    *SignInfo   `json:"-"`
	Response    *Prepay     `json:"-"`
	ErrResponse ErrResponse `json:"-"`
	Error       string      `json:"-"`
}

type Prepay struct {
	PrepayId string `json:"prepay_id"`
}

// H5OrderRsp H5 下单响应壳
type H5OrderRsp struct {
	Code        int         `json:"-"`
	SignInfo    *SignInfo   `json:"-"`
	Response    *H5Order    `json:"-"`
	ErrResponse ErrResponse `json:"-"`
	Error       string      `json:"-"`
}

type H5Order struct {
	H5Url string `json:"h5_url"`
}

// NativeOrderRsp Native 下单响应壳
type NativeOrderRsp struct {
	Code        int          `json:"-"`
	SignInfo    *SignInfo    `json:"-"`
	Response    *NativeOrder `json:"-"`
	ErrResponse ErrResponse  `json:"-"`
	Error       string       `json:"-"`
}

type NativeOrder struct {
	CodeUrl string `json:"code_url"`
}

// ===================== 查询/关闭订单 =====================

// OrderQueryRsp 订单查询响应壳
type OrderQueryRsp struct {
	Code        int         `json:"-"`
	SignInfo    *SignInfo   `json:"-"`
	Response    *OrderQuery `json:"-"`
	ErrResponse ErrResponse `json:"-"`
	Error       string      `json:"-"`
}

type OrderQuery struct {
	Mchid            string             `json:"mchid"`
	Appid            string             `json:"appid,omitempty"`
	OutTradeNo       string             `json:"out_trade_no"`
	TransactionId    string             `json:"transaction_id,omitempty"`
	TradeType        string             `json:"trade_type,omitempty"`
	TradeState       string             `json:"trade_state"`
	TradeStateDesc   string             `json:"trade_state_desc"`
	BankType         string             `json:"bank_type,omitempty"`
	Attach           string             `json:"attach,omitempty"`
	SuccessTime      string             `json:"success_time,omitempty"`
	Payer            *Payer             `json:"payer,omitempty"`
	Amount           *AmountInfo        `json:"amount,omitempty"`
	SceneInfo        *SceneInfo         `json:"scene_info,omitempty"`
	PromotionDetail  []*PromotionDetail `json:"promotion_detail,omitempty"`
}

type Payer struct {
	Openid string `json:"openid"`
}

type AmountInfo struct {
	Total         int    `json:"total"`
	PayerTotal    int    `json:"payer_total,omitempty"`
	Currency      string `json:"currency,omitempty"`
	PayerCurrency string `json:"payer_currency,omitempty"`
}

type PromotionDetail struct {
	CouponId            string         `json:"coupon_id"`
	Name                string         `json:"name,omitempty"`
	Scope               string         `json:"scope,omitempty"`
	Type                string         `json:"type,omitempty"`
	Amount              int            `json:"amount"`
	StockId             string         `json:"stock_id,omitempty"`
	DouyinpayContribute int            `json:"douyinpay_contribute,omitempty"`
	MerchantContribute  int            `json:"merchant_contribute,omitempty"`
	OtherContribute     int            `json:"other_contribute,omitempty"`
	Currency            string         `json:"currency,omitempty"`
	GoodsDetail         []*GoodsDetail `json:"goods_detail,omitempty"`
}

type GoodsDetail struct {
	GoodsId        string `json:"goods_id,omitempty"`
	Quantity       int    `json:"quantity,omitempty"`
	UnitPrice      int    `json:"unit_price,omitempty"`
	DiscountAmount int    `json:"discount_amount,omitempty"`
	GoodsRemark    string `json:"goods_remark,omitempty"`
}

// ===================== 支付成功回调解密结构 =====================

// DecryptPayResult 支付成功通知（event_type=TRANSACTION.SUCCESS）密文解密后的明文
type DecryptPayResult struct {
	Mchid           string             `json:"mchid"`
	Appid           string             `json:"appid"`
	OutTradeNo      string             `json:"out_trade_no"`
	TransactionId   string             `json:"transaction_id"`
	TradeType       string             `json:"trade_type"`
	TradeState      string             `json:"trade_state"`
	TradeStateDesc  string             `json:"trade_state_desc"`
	BankType        string             `json:"bank_type,omitempty"`
	Attach          string             `json:"attach,omitempty"`
	SuccessTime     string             `json:"success_time"`
	Payer           *Payer             `json:"payer"`
	Amount          *AmountInfo        `json:"amount"`
	SceneInfo       *SceneInfo         `json:"scene_info,omitempty"`
	PromotionDetail []*PromotionDetail `json:"promotion_detail,omitempty"`
}
