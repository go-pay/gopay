package alipay

type TradePayRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	TradeNo             string           `json:"trade_no"`
	OutTradeNo          string           `json:"out_trade_no"`
	BuyerLogonId        string           `json:"buyer_logon_id"`
	TotalAmount         string           `json:"total_amount"`
	ReceiptAmount       string           `json:"receipt_amount"`
	BuyerPayAmount      string           `json:"buyer_pay_amount"`
	PointAmount         string           `json:"point_amount"`
	InvoiceAmount       string           `json:"invoice_amount"`
	GmtPayment          string           `json:"gmt_payment"`
	FundBillList        []*FundBill      `json:"fund_bill_list"`
	StoreName           string           `json:"store_name"`
	DiscountGoodsDetail string           `json:"discount_goods_detail"`
	BuyerUserId         string           `json:"buyer_user_id"`
	BuyerOpenId         string           `json:"buyer_open_id"`
	VoucherDetailList   []*VoucherDetail `json:"voucher_detail_list"`
	MdiscountAmount     string           `json:"mdiscount_amount"`
	DiscountAmount      string           `json:"discount_amount"`
}

type FundBill struct {
	FundChannel string `json:"fund_channel"`
	Amount      string `json:"amount"`
	RealAmount  string `json:"real_amount"`
}

type VoucherDetail struct {
	Id                         string `json:"id"`
	Name                       string `json:"name"`
	Type                       string `json:"type"`
	Amount                     string `json:"amount"`
	MerchantContribute         string `json:"merchant_contribute"`
	OtherContribute            string `json:"other_contribute"`
	Memo                       string `json:"memo"`
	TemplateId                 string `json:"template_id"`
	PurchaseBuyerContribute    string `json:"purchase_buyer_contribute"`
	PurchaseMerchantContribute string `json:"purchase_merchant_contribute"`
	PurchaseAntContribute      string `json:"purchase_ant_contribute"`
}

type TradeQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	TradeNo         string      `json:"trade_no"`
	OutTradeNo      string      `json:"out_trade_no"`
	BuyerLogonId    string      `json:"buyer_logon_id"`
	TradeStatus     string      `json:"trade_status"`
	TotalAmount     string      `json:"total_amount"`
	BuyerPayAmount  string      `json:"buyer_pay_amount"`
	PointAmount     string      `json:"point_amount"`
	InvoiceAmount   string      `json:"invoice_amount"`
	SendPayDate     string      `json:"send_pay_date"`
	ReceiptAmount   string      `json:"receipt_amount"`
	StoreId         string      `json:"store_id"`
	TerminalId      string      `json:"terminal_id"`
	FundBillList    []*FundBill `json:"fund_bill_list"`
	StoreName       string      `json:"store_name"`
	BuyerUserId     string      `json:"buyer_user_id"`
	BuyerOpenId     string      `json:"buyer_open_id"`
	BuyerUserType   string      `json:"buyer_user_type"`
	MdiscountAmount string      `json:"mdiscount_amount"`
	DiscountAmount  string      `json:"discount_amount"`
	ExtInfos        string      `json:"ext_infos"`
}

type TradeRefundRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	TradeNo                 string                 `json:"trade_no"`
	OutTradeNo              string                 `json:"out_trade_no"`
	BuyerLogonId            string                 `json:"buyer_logon_id"`
	FundChange              string                 `json:"fund_change"`
	RefundFee               string                 `json:"refund_fee"`
	RefundDetailItemList    []*RefundDetailItem    `json:"refund_detail_item_list"`
	StoreName               string                 `json:"store_name"`
	BuyerUserId             string                 `json:"buyer_user_id"`
	BuyerOpenId             string                 `json:"buyer_open_id"`
	SendBackFee             string                 `json:"send_back_fee"`
	RefundHybAmount         string                 `json:"refund_hyb_amount"`
	RefundChargeInfoList    []*RefundChargeInfo    `json:"refund_charge_info_list"`
	RefundVoucherDetailList []*RefundVoucherDetail `json:"refund_voucher_detail_list"`
}

type RefundDetailItem struct {
	FundChannel string `json:"fund_channel"`
	Amount      string `json:"amount"`
	RealAmount  string `json:"real_amount"`
	FundType    string `json:"fund_type"`
}

type RefundChargeInfo struct {
	RefundChargeFee        string                `json:"refund_charge_fee"`
	SwitchFeeRate          string                `json:"switch_fee_rate"`
	ChargeType             string                `json:"charge_type"`
	RefundSubFeeDetailList []*RefundSubFeeDetail `json:"refund_sub_fee_detail_list"`
}

type RefundSubFeeDetail struct {
	RefundChargeFee string `json:"refund_charge_fee"`
	SwitchFeeRate   string `json:"switch_fee_rate"`
}

type RefundVoucherDetail struct {
	Id                         string                   `json:"id"`
	Name                       string                   `json:"name"`
	Type                       string                   `json:"type"`
	Amount                     string                   `json:"amount"`
	MerchantContribute         string                   `json:"merchant_contribute"`
	OtherContribute            string                   `json:"other_contribute"`
	Memo                       string                   `json:"memo"`
	TemplateId                 string                   `json:"template_id"`
	OtherContributeDetail      []*OtherContributeDetail `json:"other_contribute_detail"`
	PurchaseBuyerContribute    string                   `json:"purchase_buyer_contribute"`
	PurchaseMerchantContribute string                   `json:"purchase_merchant_contribute"`
	PurchaseAntContribute      string                   `json:"purchase_ant_contribute"`
}

type OtherContributeDetail struct {
	ContributeType   string `json:"contribute_type"`
	ContributeAmount string `json:"contribute_amount"`
}

type TradeFastPayRefundQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	TradeNo                 string                 `json:"trade_no"`
	OutTradeNo              string                 `json:"out_trade_no"`
	OutRequestNo            string                 `json:"out_request_no"`
	TotalAmount             string                 `json:"total_amount"`
	RefundAmount            string                 `json:"refund_amount"`
	RefundStatus            string                 `json:"refund_status"`
	RefundRoyaltys          []*RefundRoyalty       `json:"refund_royaltys"`
	GmtRefundPay            string                 `json:"gmt_refund_pay"`
	RefundDetailItemList    []*RefundDetailItem    `json:"refund_detail_item_list"`
	SendBackFee             string                 `json:"send_back_fee"`
	DepositBackInfo         *DepositBackInfo       `json:"deposit_back_info"`
	RefundHybAmount         string                 `json:"refund_hyb_amount"`
	RefundChargeInfoList    []*RefundChargeInfo    `json:"refund_charge_info_list"`
	DepositBackInfoList     []*DepositBackInfo     `json:"deposit_back_info_list"`
	RefundVoucherDetailList []*RefundVoucherDetail `json:"refund_voucher_detail_list"`
}

type RefundRoyalty struct {
	RefundAmount  string `json:"refund_amount"`
	RoyaltyType   string `json:"royalty_type"`
	ResultCode    string `json:"result_code"`
	TransOut      string `json:"trans_out"`
	TransOutEmail string `json:"trans_out_email"`
	TransIn       string `json:"trans_in"`
	TransInEmail  string `json:"trans_in_email"`
	OriTransOut   string `json:"ori_trans_out"`
	OriTransIn    string `json:"ori_trans_in"`
}

type DepositBackInfo struct {
	HasDepositBack     string `json:"has_deposit_back"`
	DbackStatus        string `json:"dback_status"`
	DbackAmount        string `json:"dback_amount"`
	BankAckTime        string `json:"bank_ack_time"`
	EstBankReceiptTime string `json:"est_bank_receipt_time"`
}

type TradeCancelRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	TradeNo    string `json:"trade_no"`
	OutTradeNo string `json:"out_trade_no"`
	RetryFlag  string `json:"retry_flag"`
	Action     string `json:"action"`
}

type TradeCloseRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	TradeNo    string `json:"trade_no"`
	OutTradeNo string `json:"out_trade_no"`
}

type DataBillDownloadUrlQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	BillDownloadUrl string `json:"bill_download_url"`
	BillFileCode    string `json:"bill_file_code"`
}

type TradePrecreateRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	OutTradeNo string `json:"out_trade_no"`
	QrCode     string `json:"qr_code"`
}

type TradeCreateRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	TradeNo    string `json:"trade_no"`
	OutTradeNo string `json:"out_trade_no"`
}

type TradeOrderInfoSyncRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	TradeNo     string `json:"trade_no"`
	OutTradeNo  string `json:"out_trade_no"`
	BuyerUserId string `json:"buyer_user_id"`
	BuyerOpenId string `json:"buyer_open_id"`
}

type ZolozAuthenticationSmilepayInitializeRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	RetCodeSub        string `json:"ret_code_sub"`
	RetMessageSub     string `json:"ret_message_sub"`
	ZimId             string `json:"zim_id"`
	ZimInitClientData string `json:"zim_init_client_data"`
}

type ZolozAuthenticationCustomerFtokenQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	Uid            string        `json:"uid"`
	OpenId         string        `json:"open_id"`
	UidTelPairList []*UidTelPair `json:"uid_tel_pair_list"`
	AgeCheckResult string        `json:"age_check_result"`
	CertNo         string        `json:"cert_no"`
	CertName       string        `json:"cert_name"`
	FaceId         string        `json:"face_id"`
}

type UidTelPair struct {
	UserId string `json:"user_id"`
	OpenId string `json:"open_id"`
}
