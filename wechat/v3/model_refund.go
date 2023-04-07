package wechat

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

// 申请退款 Rsp
type EcommerceRefundRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *EcommerceRefund `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 申请退款 Rsp
type EcommerceRefundQueryRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *EcommerceRefundQuery `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 垫付退款回补 Rsp
type EcommerceRefundAdvanceRsp struct {
	Code     int                     `json:"-"`
	SignInfo *SignInfo               `json:"-"`
	Response *EcommerceRefundAdvance `json:"response,omitempty"`
	Error    string                  `json:"-"`
}

// =========================================================分割=========================================================

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

type EcommerceRefund struct {
	RefundId        string                 `json:"refund_id"`        // 微信支付退款号
	OutRefundNo     string                 `json:"out_refund_no"`    // 商户退款单号
	CreateTime      string                 `json:"create_time"`      // 退款创建时间
	RefundAccount   string                 `json:"refund_account"`   // 退款资金来源
	Amount          *EcommerceRefundAmount `json:"amount"`           // 金额信息
	PromotionDetail []*PromotionDetailItem `json:"promotion_detail"` // 优惠退款信息
}

type EcommerceRefundAmount struct {
	Refund         int    `json:"refund"`          // 退款金额
	PayerRefund    int    `json:"payer_refund"`    // 用户退款金额
	DiscountRefund int    `json:"discount_refund"` // 优惠退款金额
	Currency       string `json:"currency"`        // 退款币种
}

type PromotionDetailItem struct {
	PromotionId  string `json:"promotion_id"`  // 券Id，券或立减金额
	Scope        string `json:"scope"`         // 优惠范围，GLOBAL：全场代金券，SINGLE：单品优惠
	Type         string `json:"type"`          // 优惠类型，COUPON：代金券，DISCOUNT：优惠券
	Amount       int    `json:"amount"`        // 优惠券面额，用户享受优惠的金额（优惠券面额=微信出资金额+商家出资金额+其他出资方金额），单位为分
	RefundAmount int    `json:"refund_amount"` // 优惠退款金额，单位为分
}

type EcommerceRefundQuery struct {
	RefundId            string                 `json:"refund_id"`                // 微信支付退款号
	OutRefundNo         string                 `json:"out_refund_no"`            // 商户退款单号
	TransactionId       string                 `json:"transaction_id"`           // 微信支付系统生成的订单号
	OutTradeNo          string                 `json:"out_trade_no"`             // 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
	Channel             string                 `json:"channel"`                  // 退款渠道
	UserReceivedAccount string                 `json:"user_received_account"`    // 退款入账账户
	SuccessTime         string                 `json:"success_time"`             // 退款成功时间
	CreateTime          string                 `json:"create_time"`              // 退款创建时间
	Status              string                 `json:"status"`                   // 退款状态
	RefundAccount       string                 `json:"refund_account,omitempty"` // 退款出资商户
	FundsAccount        string                 `json:"funds_account"`            // 资金账户
	Amount              *EcommerceRefundAmount `json:"amount"`                   // 金额信息
	PromotionDetail     []*PromotionDetailItem `json:"promotion_detail"`         // 优惠退款信息
}

type EcommerceRefundAdvance struct {
	RefundId        string `json:"refund_id"`         // 微信支付退款号
	AdvanceReturnId string `json:"advance_return_id"` // 微信回补单号
	ReturnAmount    int    `json:"return_amount"`     // 垫付回补金额
	PayerMchid      string `json:"payer_mchid"`       // 出款方商户号
	PayerAccount    string `json:"payer_account"`     // 出款方账户
	PayeeMchid      string `json:"payee_mchid"`       // 入账方商户号
	PayeeAccount    string `json:"payee_account"`     // 入账方账户
	Result          string `json:"result"`            // 垫付回补结果
	SuccessTime     string `json:"success_time"`      // 垫付回补完成时间
}
