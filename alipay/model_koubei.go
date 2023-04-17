package alipay

type KoubeiTradeOrderAggregateConsultRsp struct {
	Response     *KoubeiTradeOrderAggregateConsult `json:"koubei_trade_order_aggregate_consult_response"`
	AlipayCertSn string                            `json:"alipay_cert_sn,omitempty"`
	SignData     string                            `json:"-"`
	Sign         string                            `json:"sign"`
}

type KoubeiTradeOrderPrecreateRsp struct {
	Response     *KoubeiTradeOrderPrecreate `json:"koubei_trade_order_precreate_response"`
	AlipayCertSn string                     `json:"alipay_cert_sn,omitempty"`
	SignData     string                     `json:"-"`
	Sign         string                     `json:"sign"`
}

type KoubeiTradeItemorderBuyRsp struct {
	Response     *TradeItemorderBuy `json:"koubei_trade_itemorder_buy_response"`
	AlipayCertSn string             `json:"alipay_cert_sn,omitempty"`
	SignData     string             `json:"-"`
	Sign         string             `json:"sign"`
}

type KoubeiTradeOrderConsultRsp struct {
	Response     *KoubeiTradeOrderConsult `json:"koubei_trade_order_consult_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type KoubeiTradeItemorderRefundRsp struct {
	Response     *KoubeiTradeItemorderRefund `json:"koubei_trade_itemorder_refund_response"`
	AlipayCertSn string                      `json:"alipay_cert_sn,omitempty"`
	SignData     string                      `json:"-"`
	Sign         string                      `json:"sign"`
}

type KoubeiTradeItemorderQueryRsp struct {
	Response     *KoubeiTradeItemorderQuery `json:"koubei_trade_itemorder_query_response"`
	AlipayCertSn string                     `json:"alipay_cert_sn,omitempty"`
	SignData     string                     `json:"-"`
	Sign         string                     `json:"sign"`
}

type KoubeiTradeTicketTicketcodeSendRsp struct {
	Response     *KoubeiTradeTicketTicketcodeSend `json:"koubei_trade_ticket_ticketcode_send_response"`
	AlipayCertSn string                           `json:"alipay_cert_sn,omitempty"`
	SignData     string                           `json:"-"`
	Sign         string                           `json:"sign"`
}

type KoubeiTradeTicketTicketcodeDelayRsp struct {
	Response     *KoubeiTradeTicketTicketcodeDelay `json:"koubei_trade_ticket_ticketcode_delay_response"`
	AlipayCertSn string                            `json:"alipay_cert_sn,omitempty"`
	SignData     string                            `json:"-"`
	Sign         string                            `json:"sign"`
}

type KoubeiTradeTicketTicketcodeQueryRsp struct {
	Response     *KoubeiTradeTicketTicketcodeQuery `json:"koubei_trade_ticket_ticketcode_query_response"`
	AlipayCertSn string                            `json:"alipay_cert_sn,omitempty"`
	SignData     string                            `json:"-"`
	Sign         string                            `json:"sign"`
}

type KoubeiTradeTicketTicketcodeCancelRsp struct {
	Response     *KoubeiTradeTicketTicketcodeCancel `json:"koubei_trade_ticket_ticketcode_cancel_response"`
	AlipayCertSn string                             `json:"alipay_cert_sn,omitempty"`
	SignData     string                             `json:"-"`
	Sign         string                             `json:"sign"`
}

// =========================================================分割=========================================================

type KoubeiTradeOrderAggregateConsult struct {
	ErrorResponse
	OutOrderNo             string                `json:"out_order_no,omitempty"`
	OrderNo                string                `json:"order_no,omitempty"`
	TradeNo                string                `json:"trade_no,omitempty"`
	BuyerId                string                `json:"buyer_id,omitempty"`
	BuyerIdType            string                `json:"buyer_id_type,omitempty"`
	TotalAmount            string                `json:"total_amount,omitempty"`
	ReceiptAmount          string                `json:"receipt_amount,omitempty"`
	BuyerPayAmount         string                `json:"buyer_pay_amount,omitempty"`
	MerchantDiscountAmount string                `json:"merchant_discount_amount,omitempty"`
	PlatformDiscountAmount string                `json:"platform_discount_amount,omitempty"`
	DiscountDetailList     []*DiscountDetailInfo `json:"discount_detail_list,omitempty"`
	OrderStatus            string                `json:"order_status,omitempty"`
	PayChannel             string                `json:"pay_channel,omitempty"`
	CreateTime             string                `json:"create_time"`
	GmtPaymentTime         string                `json:"gmt_payment_time,omitempty"`
}

type DiscountDetailInfo struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Type   string `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type KoubeiTradeOrderPrecreate struct {
	ErrorResponse
	OrderNo string `json:"order_no"`
	QrCode  string `json:"qr_code,omitempty"`
}

type TradeItemorderBuy struct {
	ErrorResponse
	OrderNo        string `json:"order_no"`
	TradeNo        string `json:"trade_no"`
	CashierOrderId string `json:"cashier_order_id,omitempty"`
}

type KoubeiTradeOrderConsult struct {
	ErrorResponse
	BuyerPayAmount string `json:"buyer_pay_amount"`
	RequestId      string `json:"request_id,omitempty"`
	MCardDetail    *struct {
		Name            string `json:"name"`
		AvailableAmount string `json:"available_amount"`
		PayAmount       string `json:"pay_amount"`
	} `json:"m_card_detail,omitempty"`
	DiscountDetail *struct {
		Id             string   `json:"id"`
		DiscountDesc   []string `json:"discount_desc,omitempty"`
		DiscountType   string   `json:"discount_type"`
		IsHit          string   `json:"is_hit"`
		IsPurchased    string   `json:"is_purchased"`
		Name           string   `json:"name"`
		DiscountAmount string   `json:"discount_amount,omitempty"`
	} `json:"discount_detail,omitempty"`
}

type KoubeiTradeItemorderRefund struct {
	ErrorResponse
	OrderNo          string `json:"order_no"`
	OutRequestNo     string `json:"out_request_no"`
	RealRefundAmount string `json:"real_refund_amount"`
}

type KoubeiTradeItemorderQuery struct {
	ErrorResponse
	OrderNo                 string `json:"order_no"`
	OutOrderNo              string `json:"out_order_no"`
	PartnerID               string `json:"partner_id"`
	TradeNo                 string `json:"trade_no"`
	Status                  string `json:"status"`
	BuyerID                 string `json:"buyer_id"`
	BizProduct              string `json:"biz_product"`
	GmtCreate               string `json:"gmt_create"`
	SellerID                string `json:"seller_id,omitempty"`
	GmtPayment              string `json:"gmt_payment,omitempty"`
	GmtModified             string `json:"gmt_modified"`
	TotalAmount             string `json:"total_amount"`
	RealPayAmount           string `json:"real_pay_amount"`
	DiscountAmount          string `json:"discount_amount,omitempty"`
	DeliverSellerRealAmount string `json:"deliver_seller_real_amount"`
	ItemOrderVo             []*struct {
		ItemOrderNo  string `json:"item_order_no"`
		SkuID        string `json:"sku_id"`
		Quantity     int    `json:"quantity"`
		Price        string `json:"price"`
		Status       string `json:"status"`
		MerchantFund string `json:"merchant_fund,omitempty"`
		PlatformFund string `json:"platform_fund,omitempty"`
		ExtInfo      string `json:"ext_info,omitempty"`
	} `json:"item_order_vo"`
}

type KoubeiTradeTicketTicketcodeSend struct {
	ErrorResponse
	RequestId string `json:"request_id"`
	BizCode   string `json:"biz_code,omitempty"`
}

type KoubeiTradeTicketTicketcodeDelay struct {
	ErrorResponse
	RequestId string `json:"request_id"`
	BizCode   string `json:"biz_code,omitempty"`
}

type KoubeiTradeTicketTicketcodeQuery struct {
	ErrorResponse
	TicketCode          string             `json:"ticket_code"`
	TicketStatus        string             `json:"ticket_status"`
	TicketStatusDesc    string             `json:"ticket_status_desc"`
	ItemName            string             `json:"item_name"`
	ItemID              string             `json:"item_id"`
	OriginalPrice       string             `json:"original_price"`
	CurrentPrice        string             `json:"current_price"`
	EffectDate          string             `json:"effect_date"`
	ExpireDate          string             `json:"expire_date"`
	VoucherID           string             `json:"voucher_id"`
	OrderNo             string             `json:"order_no"`
	AvailableQuantity   string             `json:"available_quantity,omitempty"`
	TotalQuantity       string             `json:"total_quantity,omitempty"`
	ItemAlias           string             `json:"item_alias,omitempty"`
	TimeCards           string             `json:"time_cards"`
	TicketTransInfoList []*TicketTransInfo `json:"ticket_trans_info_list,omitempty"`
}

type TicketTransInfo struct {
	TicketTransID   string `json:"ticket_trans_id"`
	TicketTransType string `json:"ticket_trans_type"`
	CreateTime      string `json:"create_time"`
	LastModifyTime  string `json:"last_modify_time"`
	Quantity        string `json:"quantity"`
}

type KoubeiTradeTicketTicketcodeCancel struct {
	ErrorResponse
	RequestId string `json:"request_id"`
	BizCode   string `json:"biz_code,omitempty"`
}
