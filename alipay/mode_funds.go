package alipay

type FundTransUniTransferResponse struct {
	Response     *TransUniTransfer `json:"alipay_fund_trans_uni_transfer_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type FundAccountQueryResponse struct {
	Response     *FundAccountQuery `json:"alipay_fund_account_query_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type FundTransCommonQueryResponse struct {
	Response     *FundTransCommonQuery `json:"alipay_fund_trans_common_query_response"`
	AlipayCertSn string                `json:"alipay_cert_sn,omitempty"`
	SignData     string                `json:"-"`
	Sign         string                `json:"sign"`
}

type FundTransOrderQueryResponse struct {
	Response     *FundTransOrderQuery `json:"alipay_fund_trans_order_query_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type FundTransRefundResponse struct {
	Response     *FundTransRefund `json:"alipay_fund_trans_refund_response"`
	AlipayCertSn string           `json:"alipay_cert_sn,omitempty"`
	SignData     string           `json:"-"`
	Sign         string           `json:"sign"`
}

type FundAuthOrderFreezeResponse struct {
	Response     *FundAuthOrderFreeze `json:"alipay_fund_auth_order_freeze_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type FundAuthOrderVoucherCreateResponse struct {
	Response     *FundAuthOrderVoucherCreate `json:"alipay_fund_auth_order_voucher_create_response"`
	AlipayCertSn string                      `json:"alipay_cert_sn,omitempty"`
	SignData     string                      `json:"-"`
	Sign         string                      `json:"sign"`
}

type FundAuthOrderUnfreezeResponse struct {
	Response     *FundAuthOrderUnfreeze `json:"alipay_fund_auth_order_unfreeze_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type FundAuthOperationDetailQueryResponse struct {
	Response     *FundAuthOperationDetailQuery `json:"alipay_fund_auth_operation_detail_query_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type FundAuthOperationCancelResponse struct {
	Response     *FundAuthOperationCancel `json:"alipay_fund_auth_operation_cancel_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type FundBatchCreateResponse struct {
	Response     *FundBatchCreate `json:"alipay_fund_batch_create_response"`
	AlipayCertSn string           `json:"alipay_cert_sn,omitempty"`
	SignData     string           `json:"-"`
	Sign         string           `json:"sign"`
}

type FundBatchCloseResponse struct {
	Response     *FundBatchClose `json:"alipay_fund_batch_close_response"`
	AlipayCertSn string          `json:"alipay_cert_sn,omitempty"`
	SignData     string          `json:"-"`
	Sign         string          `json:"sign"`
}

type FundBatchDetailQueryResponse struct {
	Response     *FundBatchDetailQuery `json:"alipay_fund_batch_detail_query_response"`
	AlipayCertSn string                `json:"alipay_cert_sn,omitempty"`
	SignData     string                `json:"-"`
	Sign         string                `json:"sign"`
}

type FundTransAppPayResponse struct {
	Response     *FundTransAppPay `json:"alipay_fund_trans_app_pay_response"`
	AlipayCertSn string           `json:"alipay_cert_sn,omitempty"`
	SignData     string           `json:"-"`
	Sign         string           `json:"sign"`
}

type FundTransPayeeBindQueryRsp struct {
	Response     *FundTransPayeeBindQuery `json:"alipay_fund_trans_payee_bind_query_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type FundTransPagePayRsp struct {
	Response     *FundTransPagePay `json:"alipay_fund_trans_page_pay_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

// =========================================================分割=========================================================

type TransUniTransfer struct {
	ErrorResponse
	OutBizNo       string `json:"out_biz_no,omitempty"`
	OrderId        string `json:"order_id,omitempty"`
	PayFundOrderId string `json:"pay_fund_order_id,omitempty"`
	Status         string `json:"status,omitempty"`
	TransDate      string `json:"trans_date,omitempty"`
}

type FundAccountQuery struct {
	ErrorResponse
	AvailableAmount string       `json:"available_amount,omitempty"`
	FreezeAmount    string       `json:"freeze_amount,omitempty"`
	ExtCardInfo     *ExtCardInfo `json:"ext_card_info,omitempty"`
}

type ExtCardInfo struct {
	CardNo       string `json:"card_no,omitempty"`
	BankAccName  string `json:"bank_acc_name,omitempty"`
	CardBranch   string `json:"card_branch,omitempty"`
	CardBank     string `json:"card_bank,omitempty"`
	CardLocation string `json:"card_location,omitempty"`
	CardDeposit  string `json:"card_deposit,omitempty"`
	Status       string `json:"status,omitempty"`
}

type FundTransCommonQuery struct {
	ErrorResponse
	OrderId        string `json:"order_id,omitempty"`
	PayFundOrderId string `json:"pay_fund_order_id,omitempty"`
	OutBizNo       string `json:"out_biz_no,omitempty"`
	TransAmount    string `json:"trans_amount,omitempty"`
	Status         string `json:"status,omitempty"`
	PayDate        string `json:"pay_date,omitempty"`
	ArrivalTimeEnd string `json:"arrival_time_end,omitempty"`
	OrderFee       string `json:"order_fee,omitempty"`
	ErrorCode      string `json:"error_code,omitempty"`
	FailReason     string `json:"fail_reason,omitempty"`
}

type FundTransOrderQuery struct {
	ErrorResponse
	OrderId        string `json:"order_id,omitempty"`
	Status         string `json:"status,omitempty"`
	PayDate        string `json:"pay_date,omitempty"`
	ArrivalTimeEnd string `json:"arrival_time_end,omitempty"`
	OrderFee       string `json:"order_fee,omitempty"`
	FailReason     string `json:"fail_reason,omitempty"`
	OutBizNo       string `json:"out_biz_no,omitempty"`
	ErrorCode      string `json:"error_code,omitempty"`
}

type FundTransRefund struct {
	ErrorResponse
	RefundOrderId string `json:"refund_order_id"`
	OrderId       string `json:"order_id"`
	OutRequestNo  string `json:"out_request_no"`
	Status        string `json:"status"`
	RefundAmount  string `json:"refund_amount"`
	RefundDate    string `json:"refund_date"`
}

type FundAuthOrderFreeze struct {
	ErrorResponse
	AuthNo       string `json:"auth_no,omitempty"`
	OutOrderNo   string `json:"out_order_no,omitempty"`
	OperationId  string `json:"operation_id,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
	Amount       string `json:"amount,omitempty"`
	Status       string `json:"status,omitempty"`
	PayerUserId  string `json:"payer_user_id,omitempty"`
	PayerOpenId  string `json:"payer_open_id,omitempty"`
	PayerLogonId string `json:"payer_logon_id,omitempty"`
	GmtTrans     string `json:"gmt_trans,omitempty"`
}

type FundAuthOrderVoucherCreate struct {
	ErrorResponse
	OutOrderNo   string `json:"out_order_no,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
	CodeType     string `json:"code_type,omitempty"`
	CodeValue    string `json:"code_value,omitempty"`
	CodeUrl      string `json:"code_url,omitempty"`
}

type FundAuthOrderUnfreeze struct {
	ErrorResponse
	AuthNo       string `json:"auth_no,omitempty"`
	OutOrderNo   string `json:"out_order_no,omitempty"`
	OperationId  string `json:"operation_id,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
	Amount       string `json:"amount,omitempty"`
	Status       string `json:"status,omitempty"`
	GmtTrans     string `json:"gmt_trans,omitempty"`
	CreditAmount string `json:"credit_amount,omitempty"`
	FundAmount   string `json:"fund_amount,omitempty"`
}

type FundAuthOperationDetailQuery struct {
	ErrorResponse
	AuthNo                  string `json:"auth_no,omitempty"`
	OutOrderNo              string `json:"out_order_no,omitempty"`
	OrderStatus             string `json:"order_status,omitempty"`
	TotalFreezeAmount       string `json:"total_freeze_amount,omitempty"`
	RestAmount              string `json:"rest_amount,omitempty"`
	TotalPayAmount          string `json:"total_pay_amount,omitempty"`
	OrderTitle              string `json:"order_title,omitempty"`
	PayerLogonId            string `json:"payer_logon_id,omitempty"`
	PayerUserId             string `json:"payer_user_id,omitempty"`
	PayerOpenId             string `json:"payer_open_id,omitempty"`
	ExtraParam              string `json:"extra_param,omitempty"`
	OperationId             string `json:"operation_id,omitempty"`
	OutRequestNo            string `json:"out_request_no,omitempty"`
	Amount                  string `json:"amount,omitempty"`
	OperationType           string `json:"operation_type,omitempty"`
	Status                  string `json:"status,omitempty"`
	Remark                  string `json:"remark,omitempty"`
	GmtCreate               string `json:"gmt_create,omitempty"`
	GmtTrans                string `json:"gmt_trans,omitempty"`
	PreAuthType             string `json:"pre_auth_type,omitempty"`
	TransCurrency           string `json:"trans_currency,omitempty"`
	TotalFreezeCreditAmount string `json:"total_freeze_credit_amount,omitempty"`
	TotalFreezeFundAmount   string `json:"total_freeze_fund_amount,omitempty"`
	TotalPayCreditAmount    string `json:"total_pay_credit_amount,omitempty"`
	TotalPayFundAmount      string `json:"total_pay_fund_amount,omitempty"`
	RestCreditAmount        string `json:"rest_credit_amount,omitempty"`
	RestFundAmount          string `json:"rest_fund_amount,omitempty"`
	CreditAmount            string `json:"credit_amount,omitempty"`
	FundAmount              string `json:"fund_amount,omitempty"`
}

type FundAuthOperationCancel struct {
	ErrorResponse
	AuthNo       string `json:"auth_no,omitempty"`
	OutOrderNo   string `json:"out_order_no,omitempty"`
	OperationId  string `json:"operation_id,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
	Action       string `json:"action,omitempty"`
}

type FundBatchCreate struct {
	ErrorResponse
	OutBatchNo   string `json:"out_batch_no,omitempty"`
	BatchTransId string `json:"batch_trans_id,omitempty"`
	Status       string `json:"status,omitempty"`
}

type FundBatchClose struct {
	ErrorResponse
	BatchTransId string `json:"batch_trans_id,omitempty"`
	Status       string `json:"status,omitempty"`
}

type FundBatchDetailQuery struct {
	ErrorResponse
	BatchTransId    string `json:"batch_trans_id,omitempty"`
	BatchNo         string `json:"batch_no,omitempty"`
	BizCode         string `json:"biz_code,omitempty"`
	BizScene        string `json:"biz_scene,omitempty"`
	BatchStatus     string `json:"batch_status,omitempty"`
	ApprovalStatus  string `json:"approval_status,omitempty"`
	ErrorCode       string `json:"error_code,omitempty"`
	FailReason      string `json:"fail_reason,omitempty"`
	SignPrincipal   string `json:"sign_principal,omitempty"`
	PaymentAmount   string `json:"payment_amount,omitempty"`
	PaymentCurrency string `json:"payment_currency,omitempty"`
	PageSize        int    `json:"page_size,omitempty"`
	PageNum         int    `json:"page_num,omitempty"`
	ProductCode     string `json:"product_code,omitempty"`
	TotalPageCount  string `json:"total_page_count,omitempty"`
	OutBatchNo      string `json:"out_batch_no,omitempty"`
	GmtFinish       string `json:"gmt_finish,omitempty"`
	TotalAmount     string `json:"total_amount,omitempty"`
	GmtPayFinish    string `json:"gmt_pay_finish,omitempty"`
	PayerId         string `json:"payer_id,omitempty"`
	SuccessAmount   string `json:"success_amount,omitempty"`
	FailAmount      string `json:"fail_amount,omitempty"`
	FailCount       string `json:"fail_count,omitempty"`
	SuccessCount    string `json:"success_count,omitempty"`
	TotalItemCount  string `json:"total_item_count,omitempty"`
	AccDetailList   []*struct {
		DetailNo           string `json:"detail_no,omitempty"`
		PaymentAmount      string `json:"payment_amount,omitempty"`
		PaymentCurrency    string `json:"payment_currency,omitempty"`
		TransAmount        string `json:"trans_amount,omitempty"`
		TransCurrency      string `json:"trans_currency,omitempty"`
		SettlementAmount   string `json:"settlement_amount,omitempty"`
		SettlementCurrency string `json:"settlement_currency,omitempty"`
		PayeeInfo          *struct {
			PayeeAccount string `json:"payee_account,omitempty"`
			PayeeType    string `json:"payee_type,omitempty"`
			PayeeName    string `json:"payee_name,omitempty"`
		} `json:"payee_info,omitempty"`
		CertInfo *struct {
			CertNo   string `json:"cert_no,omitempty"`
			CertType string `json:"cert_type,omitempty"`
		} `json:"cert_info,omitempty"`
		Remark       string `json:"remark,omitempty"`
		Status       string `json:"status,omitempty"`
		ExchangeRate *struct {
			Rate             string `json:"rate,omitempty"`
			BaseCurrency     string `json:"base_currency,omitempty"`
			ExchangeCurrency string `json:"exchange_currency,omitempty"`
		} `json:"exchange_rate,omitempty"`
		NeedRetry     string `json:"need_retry,omitempty"`
		AlipayOrderNo string `json:"alipay_order_no,omitempty"`
		OutBizNo      string `json:"out_biz_no,omitempty"`
		DetailId      string `json:"detail_id,omitempty"`
		ErrorCode     string `json:"error_code,omitempty"`
		ErrorMsg      string `json:"error_msg,omitempty"`
		GmtCreate     string `json:"gmt_create,omitempty"`
		GmtFinish     string `json:"gmt_finish,omitempty"`
		SubStatus     string `json:"sub_status,omitempty"`
	} `json:"acc_detail_list,omitempty"`
}

type FundTransAppPay struct {
	ErrorResponse
	OutBizNo string `json:"out_biz_no,omitempty"`
	OrderId  string `json:"order_id,omitempty"`
	Status   string `json:"status,omitempty"`
}

type FundTransPayeeBindQuery struct {
	ErrorResponse
	Bind string `json:"bind"` // 是否绑定收款账号。true：已绑定；false：未绑定
}

type FundTransPagePay struct {
	ErrorResponse
	OutBizNo string `json:"out_biz_no"`
	OrderID  string `json:"order_id,omitempty"`
	Status   string `json:"status"`
}
