package alipay

type FundTransUniTransferRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	OutBizNo       string `json:"out_biz_no"`        // 商户订单号
	OrderId        string `json:"order_id"`          // 支付宝转账订单号
	PayFundOrderId string `json:"pay_fund_order_id"` // 支付宝支付资金流水号
	TransDate      string `json:"trans_date"`        // 订单支付时间
	Status         string `json:"status"`
}

type FundAccountQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	AvailableAmount string `json:"available_amount"`
	FreezeAmount    string `json:"freeze_amount"`
	TotalAmount     string `json:"total_amount"`
	AmountDetail    struct {
		Acs  string `json:"acs"`
		Bank string `json:"bank"`
	} `json:"amount_detail"`
}

type FundQuotaQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ToCorporateDailyAvailableAmount   string `json:"to_corporate_daily_available_amount"`
	ToPrivateDailyAvailableAmount     string `json:"to_private_daily_available_amount"`
	ToCorporateMonthlyAvailableAmount string `json:"to_corporate_monthly_available_amount"`
	ToPrivateMonthlyAvailableAmount   string `json:"to_private_monthly_available_amount"`
}

type DataBillEreceiptApplyRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	FileId string `json:"file_id"`
}

type DataBillEreceiptQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	Status       string `json:"status"`
	DownloadUrl  string `json:"download_url"`
	ErrorMessage string `json:"error_message"`
}

type FundTransCommonQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	OrderId              string `json:"order_id"`
	InflowSettleSerialNo string `json:"inflow_settle_serial_no"`
	PayFundOrderId       string `json:"pay_fund_order_id"`
	OutBizNo             string `json:"out_biz_no"`
	TransAmount          string `json:"trans_amount"`
	Status               string `json:"status"`
	PayDate              string `json:"pay_date"`
	SubStatus            string `json:"sub_status"`
	ArrivalTimeEnd       string `json:"arrival_time_end"`
	OrderFee             string `json:"order_fee"`
	ErrorCode            string `json:"error_code"`
	FailReason           string `json:"fail_reason"`
	SubOrderErrorCode    string `json:"sub_order_error_code"`
	SubOrderFailReason   string `json:"sub_order_fail_reason"`
	SubOrderStatus       string `json:"sub_order_status"`
	SettleSerialNo       string `json:"settle_serial_no"`
	ReceiverUserId       string `json:"receiver_user_id"`
	ReceiverOpenId       string `json:"receiver_open_id"`
}

type FundTransMultistepTransferRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	OrderId     string `json:"order_id"`
	OrderStatus string `json:"order_status"`
}

type FundTransMultistepQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	OrderId       string         `json:"order_id"`
	OutBizNo      string         `json:"out_biz_no"`
	RequestUserId string         `json:"request_user_id"`
	TotalAmount   string         `json:"total_amount"`
	TotalCount    string         `json:"total_count"`
	Remark        string         `json:"remark"`
	OrderDetails  []*OrderDetail `json:"order_details"`
}

type OrderDetail struct {
	OutBizNo      string     `json:"out_biz_no"`
	RequestUserId string     `json:"request_user_id"`
	OpenId        string     `json:"open_id"`
	OrderId       string     `json:"order_id"`
	Amount        string     `json:"amount"`
	PayerInfo     *PayerInfo `json:"payer_info"`
	PayeeInfo     *PayeeInfo `json:"payee_info"`
	Status        string     `json:"status"`
	GmtPay        string     `json:"gmt_pay"`
	OrderTitle    string     `json:"order_title"`
	OrderSeq      int        `json:"order_seq"`
	Remark        string     `json:"remark"`
}

type PayerInfo struct {
	Identity        string `json:"identity"`
	IdentityType    string `json:"identity_type"`
	Name            string `json:"name"`
	AgreementNo     string `json:"agreement_no"`
	RentAgreementNo string `json:"rent_agreement_no"`
}

type PayeeInfo struct {
	Identity        string `json:"identity"`
	IdentityType    string `json:"identity_type"`
	Name            string `json:"name"`
	AgreementNo     string `json:"agreement_no"`
	RentAgreementNo string `json:"rent_agreement_no"`
}
