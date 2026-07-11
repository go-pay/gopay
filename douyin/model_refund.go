package douyin

// RefundRsp 退款申请/查询响应壳
type RefundRsp struct {
	Code        int         `json:"-"`
	SignInfo    *SignInfo   `json:"-"`
	Response    *Refund     `json:"-"`
	ErrResponse ErrResponse `json:"-"`
	Error       string      `json:"-"`
}

// Refund 退款响应体（Refund / RefundQuery 共用）
type Refund struct {
	RefundId            string                    `json:"refund_id"`
	OutRefundNo         string                    `json:"out_refund_no"`
	TransactionId       string                    `json:"transaction_id"`
	OutTradeNo          string                    `json:"out_trade_no"`
	Channel             string                    `json:"channel,omitempty"`
	UserReceivedAccount string                    `json:"user_received_account,omitempty"`
	SuccessTime         string                    `json:"success_time,omitempty"`
	CreateTime          string                    `json:"create_time"`
	Status              string                    `json:"status"`
	FundsAccount        string                    `json:"funds_account,omitempty"`
	Amount              *RefundResponseAmount     `json:"amount,omitempty"`
	PromotionDetail     []*RefundPromotionDetail  `json:"promotion_detail,omitempty"`
}

type RefundResponseAmount struct {
	Refund           int              `json:"refund"`
	Total            int              `json:"total"`
	PayerTotal       int              `json:"payer_total"`
	PayerRefund      int              `json:"payer_refund"`
	SettlementRefund int              `json:"settlement_refund,omitempty"`
	SettlementTotal  int              `json:"settlement_total,omitempty"`
	DiscountRefund   int              `json:"discount_refund"`
	Currency         string           `json:"currency"`
	RefundFee        int              `json:"refund_fee,omitempty"`
	From             []*RefundAcctSrc `json:"from,omitempty"`
}

type RefundAcctSrc struct {
	Account string `json:"account"`
	Amount  int    `json:"amount"`
}

type RefundPromotionDetail struct {
	PromotionId  string                 `json:"promotion_id"`
	Type         string                 `json:"type"`
	Amount       int                    `json:"amount"`
	RefundAmount int                    `json:"refund_amount"`
	Scope        string                 `json:"scope"`
	GoodsDetail  []*RefundGoodsDetail   `json:"goods_detail,omitempty"`
}

type RefundGoodsDetail struct {
	MerchantGoodsId   string `json:"merchant_goods_id"`
	DouyinpayGoodsId  string `json:"douyinpay_goods_id,omitempty"`
	GoodsName         string `json:"goods_name,omitempty"`
	UnitPrice         int    `json:"unit_price"`
	RefundAmount      int    `json:"refund_amount"`
	RefundQuantity    int    `json:"refund_quantity"`
}

// ===================== 退款成功回调解密结构 =====================

// DecryptRefundResult 退款结果通知（event_type=REFUND.SUCCESS）密文解密后的明文
type DecryptRefundResult struct {
	Mchid               string                `json:"mchid"`
	OutTradeNo          string                `json:"out_trade_no"`
	TransactionId       string                `json:"transaction_id"`
	OutRefundNo         string                `json:"out_refund_no"`
	RefundId            string                `json:"refund_id"`
	RefundStatus        string                `json:"refund_status"`
	SuccessTime         string                `json:"success_time,omitempty"`
	UserReceivedAccount string                `json:"user_received_account,omitempty"`
	Amount              *RefundNotifyAmount   `json:"amount"`
	PromotionDetail     []*RefundPromotionDetail `json:"promotion_detail,omitempty"`
}

type RefundNotifyAmount struct {
	Total       int `json:"total"`
	Refund      int `json:"refund"`
	PayerTotal  int `json:"payer_total"`
	PayerRefund int `json:"payer_refund"`
}
