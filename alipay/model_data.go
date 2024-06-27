package alipay

type DataBillBalanceQueryResponse struct {
	Response     *DataBillBalanceQuery `json:"alipay_data_bill_balance_query_response"`
	AlipayCertSn string                `json:"alipay_cert_sn,omitempty"`
	SignData     string                `json:"-"`
	Sign         string                `json:"sign"`
}

type DataBillAccountLogQueryResponse struct {
	Response     *DataBillAccountLogQuery `json:"alipay_data_bill_accountlog_query_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type DataBillEreceiptApplyRsp struct {
	Response     *DataBillEreceiptApply `json:"alipay_data_bill_ereceipt_apply_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type DataBillEreceiptQueryRsp struct {
	Response     *DataBillEreceiptQuery `json:"alipay_data_bill_ereceipt_query_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type DataBillDownloadUrlQueryResponse struct {
	Response     *DataBillDownloadUrlQuery `json:"alipay_data_dataservice_bill_downloadurl_query_response"`
	AlipayCertSn string                    `json:"alipay_cert_sn,omitempty"`
	SignData     string                    `json:"-"`
	Sign         string                    `json:"sign"`
}

// =========================================================分割=========================================================

type DataBillBalanceQuery struct {
	ErrorResponse
	TotalAmount     string `json:"total_amount,omitempty"`
	AvailableAmount string `json:"available_amount,omitempty"`
	FreezeAmount    string `json:"freeze_amount,omitempty"`
	SettleAmount    string `json:"settle_amount,omitempty"`
}

type DataBillAccountLogQuery struct {
	ErrorResponse
	PageNo     string                  `json:"page_no,omitempty"`
	PageSize   string                  `json:"page_size,omitempty"`
	TotalSize  string                  `json:"total_size,omitempty"`
	DetailList []*AccountLogItemResult `json:"detail_list,omitempty"`
}

type AccountLogItemResult struct {
	TransDt             string `json:"trans_dt,omitempty"`
	AccountLogId        string `json:"account_log_id,omitempty"`
	AlipayOrderNo       string `json:"alipay_order_no,omitempty"`
	MerchantOrderNo     string `json:"merchant_order_no,omitempty"`
	TransAmount         string `json:"trans_amount,omitempty"`
	Balance             string `json:"balance,omitempty"`
	Type                string `json:"type,omitempty"`
	OtherAccount        string `json:"other_account,omitempty"`
	TransMemo           string `json:"trans_memo,omitempty"`
	Direction           string `json:"direction,omitempty"`
	BillSource          string `json:"bill_source,omitempty"`
	BizNos              string `json:"biz_nos,omitempty"`
	BizOrigNo           string `json:"biz_orig_no,omitempty"`
	BizDesc             string `json:"biz_desc,omitempty"`
	MerchantOutRefundNo string `json:"merchant_out_refund_no,omitempty"`
	ComplementInfo      string `json:"complement_info,omitempty"`
	StoreName           string `json:"store_name,omitempty"`
}

type DataBillDownloadUrlQuery struct {
	ErrorResponse
	BillDownloadUrl string `json:"bill_download_url,omitempty"`
}

type DataBillEreceiptApply struct {
	ErrorResponse
	FileId string `json:"file_id"`
}

type DataBillEreceiptQuery struct {
	ErrorResponse
	Status       string `json:"status"`
	DownloadUrl  string `json:"download_url"`
	ErrorMessage string `json:"error_message"`
}
