package alipay

type TradePayResponse struct {
	Response     *TradePay `json:"alipay_trade_pay_response"`
	AlipayCertSn string    `json:"alipay_cert_sn,omitempty"`
	SignData     string    `json:"-"`
	Sign         string    `json:"sign"`
}

type TradePrecreateResponse struct {
	Response     *TradePrecreate `json:"alipay_trade_precreate_response"`
	NullResponse *ErrorResponse  `json:"null_response,omitempty"`
	AlipayCertSn string          `json:"alipay_cert_sn,omitempty"`
	SignData     string          `json:"-"`
	Sign         string          `json:"sign"`
}

type TradeCreateResponse struct {
	Response     *TradeCreate `json:"alipay_trade_create_response"`
	AlipayCertSn string       `json:"alipay_cert_sn,omitempty"`
	SignData     string       `json:"-"`
	Sign         string       `json:"sign"`
}

type TradeQueryResponse struct {
	Response     *TradeQuery `json:"alipay_trade_query_response"`
	AlipayCertSn string      `json:"alipay_cert_sn,omitempty"`
	SignData     string      `json:"-"`
	Sign         string      `json:"sign"`
}

type TradeCancelResponse struct {
	Response     *TradeCancel `json:"alipay_trade_cancel_response"`
	AlipayCertSn string       `json:"alipay_cert_sn,omitempty"`
	SignData     string       `json:"-"`
	Sign         string       `json:"sign"`
}

type TradeCloseResponse struct {
	Response     *TradeClose `json:"alipay_trade_close_response"`
	AlipayCertSn string      `json:"alipay_cert_sn,omitempty"`
	SignData     string      `json:"-"`
	Sign         string      `json:"sign"`
}

type TradeRefundResponse struct {
	Response     *TradeRefund `json:"alipay_trade_refund_response"`
	AlipayCertSn string       `json:"alipay_cert_sn,omitempty"`
	SignData     string       `json:"-"`
	Sign         string       `json:"sign"`
}

type TradePageRefundResponse struct {
	Response     *TradePageRefund `json:"alipay_trade_page_refund_response"`
	AlipayCertSn string           `json:"alipay_cert_sn,omitempty"`
	SignData     string           `json:"-"`
	Sign         string           `json:"sign"`
}

type TradeFastpayRefundQueryResponse struct {
	Response     *TradeRefundQuery `json:"alipay_trade_fastpay_refund_query_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type TradeOrderInfoSyncRsp struct {
	Response     *TradeOrderInfoSync `json:"alipay_trade_orderinfo_sync_response"`
	AlipayCertSn string              `json:"alipay_cert_sn,omitempty"`
	SignData     string              `json:"-"`
	Sign         string              `json:"sign"`
}

type TradeAdvanceConsultRsp struct {
	Response     *TradeAdvanceConsult `json:"alipay_trade_advance_consult_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type PcreditHuabeiAuthSettleApplyRsp struct {
	Response     *PcreditHuabeiAuthSettleApply `json:"alipay_pcredit_huabei_auth_settle_apply_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type PaymentTradeOrderCreateRsp struct {
	Response     *PaymentTradeOrderCreate `json:"mybank_payment_trade_order_create_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type TradeRepaybillQueryRsp struct {
	Response     *TradeRepaybillQuery `json:"alipay_trade_repaybill_query_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

// =========================================================分割=========================================================

type TradePay struct {
	ErrorResponse
	TradeNo             string           `json:"trade_no,omitempty"`
	OutTradeNo          string           `json:"out_trade_no,omitempty"`
	BuyerLogonId        string           `json:"buyer_logon_id,omitempty"`
	TotalAmount         string           `json:"total_amount,omitempty"`
	ReceiptAmount       string           `json:"receipt_amount,omitempty"`
	BuyerPayAmount      string           `json:"buyer_pay_amount,omitempty"`
	PointAmount         string           `json:"point_amount,omitempty"`
	InvoiceAmount       string           `json:"invoice_amount,omitempty"`
	FundBillList        []*TradeFundBill `json:"fund_bill_list"`
	StoreName           string           `json:"store_name,omitempty"`
	BuyerUserId         string           `json:"buyer_user_id,omitempty"`
	BuyerOpenId         string           `json:"buyer_open_id,omitempty"`
	DiscountGoodsDetail string           `json:"discount_goods_detail,omitempty"`
	AsyncPaymentMode    string           `json:"async_payment_mode,omitempty"`
	VoucherDetailList   []*VoucherDetail `json:"voucher_detail_list"`
	AdvanceAmount       string           `json:"advance_amount,omitempty"`
	AuthTradePayMode    string           `json:"auth_trade_pay_mode,omitempty"`
	MdiscountAmount     string           `json:"mdiscount_amount,omitempty"`
	DiscountAmount      string           `json:"discount_amount,omitempty"`
	CreditPayMode       string           `json:"credit_pay_mode"`
	CreditBizOrderId    string           `json:"credit_biz_order_id"`
}

type TradeFundBill struct {
	FundChannel string `json:"fund_channel,omitempty"` // 同步通知里是 fund_channel
	Amount      string `json:"amount,omitempty"`
	RealAmount  string `json:"real_amount,omitempty"`
	FundType    string `json:"fund_type,omitempty"`
}

type VoucherDetail struct {
	Id                         string `json:"id,omitempty"`
	Name                       string `json:"name,omitempty"`
	Type                       string `json:"type,omitempty"`
	Amount                     string `json:"amount,omitempty"`
	MerchantContribute         string `json:"merchant_contribute,omitempty"`
	OtherContribute            string `json:"other_contribute,omitempty"`
	Memo                       string `json:"memo,omitempty"`
	TemplateId                 string `json:"template_id,omitempty"`
	PurchaseBuyerContribute    string `json:"purchase_buyer_contribute,omitempty"`
	PurchaseMerchantContribute string `json:"purchase_merchant_contribute,omitempty"`
	PurchaseAntContribute      string `json:"purchase_ant_contribute,omitempty"`
}

type TradePrecreate struct {
	ErrorResponse
	OutTradeNo string `json:"out_trade_no,omitempty"`
	QrCode     string `json:"qr_code,omitempty"`
}

type TradeCreate struct {
	ErrorResponse
	TradeNo    string `json:"trade_no,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
}

type TradeQuery struct {
	ErrorResponse
	TradeNo               string           `json:"trade_no,omitempty"`
	OutTradeNo            string           `json:"out_trade_no,omitempty"`
	BuyerLogonId          string           `json:"buyer_logon_id,omitempty"`
	TradeStatus           string           `json:"trade_status,omitempty"`
	TotalAmount           string           `json:"total_amount,omitempty"`
	TransCurrency         string           `json:"trans_currency,omitempty"`
	SettleCurrency        string           `json:"settle_currency,omitempty"`
	SettleAmount          string           `json:"settle_amount,omitempty"`
	PayCurrency           string           `json:"pay_currency,omitempty"`
	PayAmount             string           `json:"pay_amount,omitempty"`
	SettleTransRate       string           `json:"settle_trans_rate,omitempty"`
	TransPayRate          string           `json:"trans_pay_rate,omitempty"`
	BuyerPayAmount        string           `json:"buyer_pay_amount,omitempty"`
	PointAmount           string           `json:"point_amount,omitempty"`
	InvoiceAmount         string           `json:"invoice_amount,omitempty"`
	SendPayDate           string           `json:"send_pay_date,omitempty"`
	ReceiptAmount         string           `json:"receipt_amount,omitempty"`
	StoreId               string           `json:"store_id,omitempty"`
	TerminalId            string           `json:"terminal_id,omitempty"`
	FundBillList          []*TradeFundBill `json:"fund_bill_list"`
	StoreName             string           `json:"store_name,omitempty"`
	BuyerUserId           string           `json:"buyer_user_id,omitempty"`
	BuyerOpenId           string           `json:"buyer_open_id,omitempty"`
	DiscountGoodsDetail   string           `json:"discount_goods_detail,omitempty"`
	IndustrySepcDetail    string           `json:"industry_sepc_detail,omitempty"`
	IndustrySepcDetailGov string           `json:"industry_sepc_detail_gov,omitempty"`
	IndustrySepcDetailAcc string           `json:"industry_sepc_detail_acc,omitempty"`
	ChargeAmount          string           `json:"charge_amount,omitempty"`
	ChargeFlags           string           `json:"charge_flags,omitempty"`
	SettlementId          string           `json:"settlement_id,omitempty"`
	TradeSettleInfo       *TradeSettleInfo `json:"trade_settle_info,omitempty"`
	AuthTradePayMode      string           `json:"auth_trade_pay_mode,omitempty"`
	BuyerUserType         string           `json:"buyer_user_type,omitempty"`
	MdiscountAmount       string           `json:"mdiscount_amount,omitempty"`
	DiscountAmount        string           `json:"discount_amount,omitempty"`
	Subject               string           `json:"subject,omitempty"`
	Body                  string           `json:"body,omitempty"`
	AlipaySubMerchantId   string           `json:"alipay_sub_merchant_id,omitempty"`
	ExtInfos              string           `json:"ext_infos,omitempty"`
	PassbackParams        string           `json:"passback_params,omitempty"`
	HbFqPayInfo           *HbFqPayInfo     `json:"hb_fq_pay_info,omitempty"`
	CreditPayMode         string           `json:"credit_pay_mode,omitempty"`
	CreditBizOrderId      string           `json:"credit_biz_order_id,omitempty"`
	HybAmount             string           `json:"hyb_amount,omitempty"`
	BkagentRespInfo       *BkAgentRespInfo `json:"bkagent_resp_info,omitempty"`
	ChargeInfoList        []*ChargeInfo    `json:"charge_info_list,omitempty"`
	BizSettleMode         string           `json:"biz_settle_mode,omitempty"`
}

type BkAgentRespInfo struct {
	BindtrxId        string `json:"bindtrx_id,omitempty"`
	BindclrissrId    string `json:"bindclrissr_id,omitempty"`
	BindpyeracctbkId string `json:"bindpyeracctbk_id,omitempty"`
	BkpyeruserCode   string `json:"bkpyeruser_code,omitempty"`
	EstterLocation   string `json:"estter_location,omitempty"`
}

type ChargeInfo struct {
	ChargeFee               string    `json:"charge_fee,omitempty"`
	OriginalChargeFee       string    `json:"original_charge_fee,omitempty"`
	SwitchFeeRate           string    `json:"switch_fee_rate,omitempty"`
	IsRatingOnTradeReceiver string    `json:"is_rating_on_trade_receiver,omitempty"`
	IsRatingOnSwitch        string    `json:"is_rating_on_switch,omitempty"`
	ChargeType              string    `json:"charge_type,omitempty"`
	SubFeeDetailList        []*SubFee `json:"sub_fee_detail_list,omitempty"`
}

type SubFee struct {
	ChargeFee         string `json:"charge_fee,omitempty"`
	OriginalChargeFee string `json:"original_charge_fee,omitempty"`
	SwitchFeeRate     string `json:"switch_fee_rate,omitempty"`
}

type TradeSettleInfo struct {
	TradeSettleDetailList []*TradeSettleDetail `json:"trade_settle_detail_list,omitempty"`
}

type TradeSettleDetail struct {
	OperationType     string `json:"operation_type,omitempty"`
	OperationSerialNo string `json:"operation_serial_no,omitempty"`
	OperationDt       string `json:"operation_dt,omitempty"`
	TransOut          string `json:"trans_out,omitempty"`
	TransIn           string `json:"trans_in,omitempty"`
	Amount            string `json:"amount,omitempty"`
	OriTransOut       string `json:"ori_trans_out,omitempty"`
	OriTransIn        string `json:"ori_trans_in,omitempty"`
}

type HbFqPayInfo struct {
	UserInstallNum string `json:"user_install_num,omitempty"`
}

type TradeCancel struct {
	ErrorResponse
	TradeNo            string `json:"trade_no,omitempty"`
	OutTradeNo         string `json:"out_trade_no,omitempty"`
	RetryFlag          string `json:"retry_flag,omitempty"`
	Action             string `json:"action,omitempty"`
	GmtRefundPay       string `json:"gmt_refund_pay,omitempty"`
	RefundSettlementId string `json:"refund_settlement_id,omitempty"`
}

type TradeClose struct {
	ErrorResponse
	TradeNo    string `json:"trade_no,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
}

type TradeRefund struct {
	ErrorResponse
	TradeNo                      string                 `json:"trade_no,omitempty"`
	OutTradeNo                   string                 `json:"out_trade_no,omitempty"`
	BuyerLogonId                 string                 `json:"buyer_logon_id,omitempty"`
	FundChange                   string                 `json:"fund_change,omitempty"`
	RefundFee                    string                 `json:"refund_fee,omitempty"`
	RefundDetailItemList         []*TradeFundBill       `json:"refund_detail_item_list,omitempty"`
	StoreName                    string                 `json:"store_name,omitempty"`
	BuyerUserId                  string                 `json:"buyer_user_id,omitempty"`
	BuyerOpenId                  string                 `json:"buyer_open_id,omitempty"`
	SendBackFee                  string                 `json:"send_back_fee,omitempty"`
	OpenId                       string                 `json:"open_id,omitempty"`
	RefundCurrency               string                 `json:"refund_currency,omitempty"`
	GmtRefundPay                 string                 `json:"gmt_refund_pay,omitempty"`
	RefundPresetPaytoolList      []*RefundPresetPaytool `json:"refund_preset_paytool_list,omitempty"`
	RefundChargeAmount           string                 `json:"refund_charge_amount,omitempty"`
	RefundSettlementId           string                 `json:"refund_settlement_id,omitempty"`
	PresentRefundBuyerAmount     string                 `json:"present_refund_buyer_amount,omitempty"`
	PresentRefundDiscountAmount  string                 `json:"present_refund_discount_amount,omitempty"`
	PresentRefundMdiscountAmount string                 `json:"present_refund_mdiscount_amount,omitempty"`
	HasDepositBack               string                 `json:"has_deposit_back,omitempty"`
	RefundHybAmount              string                 `json:"refund_hyb_amount,omitempty"`
}

type RefundPresetPaytool struct {
	Amount         []string `json:"amount,omitempty"`
	AssertTypeCode string   `json:"assert_type_code,omitempty"`
}

type TradePageRefund struct {
	ErrorResponse
	TradeNo      string `json:"trade_no,omitempty"`
	OutTradeNo   string `json:"out_trade_no,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
	RefundAmount string `json:"refund_amount,omitempty"`
}

type TradeRefundQuery struct {
	ErrorResponse
	TradeNo              string           `json:"trade_no,omitempty"`
	OutTradeNo           string           `json:"out_trade_no,omitempty"`
	OutRequestNo         string           `json:"out_request_no,omitempty"`
	RefundReason         string           `json:"refund_reason,omitempty"`
	TotalAmount          string           `json:"total_amount,omitempty"`
	RefundAmount         string           `json:"refund_amount,omitempty"`
	RefundStatus         string           `json:"refund_status,omitempty"`
	RefundRoyaltys       []*RefundRoyalty `json:"refund_royaltys,omitempty"`
	GmtRefundPay         string           `json:"gmt_refund_pay,omitempty"`
	RefundDetailItemList []*TradeFundBill `json:"refund_detail_item_list,omitempty"`
	SendBackFee          string           `json:"send_back_fee,omitempty"`
	DepositBackInfo      *DepositBackInfo `json:"deposit_back_info,omitempty"`
}

type RefundRoyalty struct {
	RefundAmount  string `json:"refund_amount,omitempty"`
	RoyaltyType   string `json:"royalty_type,omitempty"`
	ResultCode    string `json:"result_code,omitempty"`
	TransOut      string `json:"trans_out,omitempty"`
	TransOutEmail string `json:"trans_out_email,omitempty"`
	TransIn       string `json:"trans_in,omitempty"`
	TransInEmail  string `json:"trans_in_email,omitempty"`
}

type DepositBackInfo struct {
	HasDepositBack     string `json:"has_deposit_back,omitempty"`
	DbackStatus        string `json:"dback_status,omitempty"`
	DbackAmount        string `json:"dback_amount,omitempty"`
	BankAckTime        string `json:"bank_ack_time,omitempty"`
	EstBankReceiptTime string `json:"est_bank_receipt_time,omitempty"`
}

type TradeOrderInfoSync struct {
	ErrorResponse
	TradeNo     string `json:"trade_no"`
	OutTradeNo  string `json:"out_trade_no"`
	BuyerUserId string `json:"buyer_user_id"`
	BuyerOpenId string `json:"buyer_open_id,omitempty"`
}

type TradeAdvanceConsult struct {
	ErrorResponse
	ReferResult             bool                      `json:"refer_result"`
	WaitRepaymentOrderInfos []*WaitRepaymentOrderInfo `json:"wait_repayment_order_infos,omitempty"`
	WaitRepaymentAmount     string                    `json:"wait_repayment_amount,omitempty"`
	WaitRepaymentOrderCount string                    `json:"wait_repayment_order_count,omitempty"`
	RiskLevel               string                    `json:"risk_level,omitempty"`
	ResultMessage           string                    `json:"result_message"`
	ResultCode              string                    `json:"result_code"`
}

type WaitRepaymentOrderInfo struct {
	AdvanceOrderId      string `json:"advance_order_id"`
	AlipayUserId        string `json:"alipay_user_id"`
	OrigBizOrderId      string `json:"orig_biz_order_id"`
	BizProduct          string `json:"biz_product"`
	WaitRepaymentAmount string `json:"wait_repayment_amount"`
}

type PcreditHuabeiAuthSettleApply struct {
	ErrorResponse
	OutRequestNo string `json:"out_request_no"`
	FailReason   string `json:"fail_reason,omitempty"`
}

type PaymentTradeOrderCreate struct {
	ErrorResponse
}

type TradeRepaybillQuery struct {
	ErrorResponse
	BillNo                string `json:"bill_no"`
	BillAmount            string `json:"bill_amount"`
	BillOverdueAmount     string `json:"bill_overdue_amount"`
	BillPaidAmount        string `json:"bill_paid_amount"`
	BillPaidRevokedAmount string `json:"bill_paid_revoked_amount"`
	BillRevokedAmount     string `json:"bill_revoked_amount"`
	BillStatus            string `json:"bill_status"`
}
