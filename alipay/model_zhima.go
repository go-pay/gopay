package alipay

type ZhimaCreditScoreGetResponse struct {
	Response     *ScoreGet `json:"zhima_credit_score_get_response"`
	AlipayCertSn string    `json:"alipay_cert_sn,omitempty"`
	SignData     string    `json:"-"`
	Sign         string    `json:"sign"`
}

type ZhimaCreditEpSceneRatingInitializeRsp struct {
	Response     *ZhimaCreditEpSceneRatingInitialize `json:"zhima_credit_ep_scene_rating_initialize_response"`
	AlipayCertSn string                              `json:"alipay_cert_sn,omitempty"`
	SignData     string                              `json:"-"`
	Sign         string                              `json:"sign"`
}

type ZhimaCreditEpSceneFulfillmentSyncRsp struct {
	Response     *ZhimaCreditEpSceneFulfillmentSync `json:"zhima_credit_ep_scene_fulfillment_sync_response"`
	AlipayCertSn string                             `json:"alipay_cert_sn,omitempty"`
	SignData     string                             `json:"-"`
	Sign         string                             `json:"sign"`
}

type ZhimaCreditEpSceneAgreementUseRsp struct {
	Response     *ZhimaCreditEpSceneAgreementUse `json:"zhima_credit_ep_scene_agreement_use_response"`
	AlipayCertSn string                          `json:"alipay_cert_sn,omitempty"`
	SignData     string                          `json:"-"`
	Sign         string                          `json:"sign"`
}

type ZhimaCreditEpSceneAgreementCancelRsp struct {
	Response     *ZhimaCreditEpSceneAgreementCancel `json:"zhima_credit_ep_scene_agreement_cancel_response"`
	AlipayCertSn string                             `json:"alipay_cert_sn,omitempty"`
	SignData     string                             `json:"-"`
	Sign         string                             `json:"sign"`
}

type ZhimaCreditEpSceneFulfillmentlistSyncRsp struct {
	Response     *ZhimaCreditEpSceneFulfillmentlistSync `json:"zhima_credit_ep_scene_fulfillmentlist_sync_response"`
	AlipayCertSn string                                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                                 `json:"-"`
	Sign         string                                 `json:"sign"`
}

type ZhimaCreditPeZmgoCumulationSyncRsp struct {
	Response     *ZhimaCreditPeZmgoCumulationSync `json:"zhima_credit_pe_zmgo_cumulation_sync_response"`
	AlipayCertSn string                           `json:"alipay_cert_sn,omitempty"`
	SignData     string                           `json:"-"`
	Sign         string                           `json:"sign"`
}

type ZhimaMerchantZmgoCumulateSyncRsp struct {
	Response     *ZhimaMerchantZmgoCumulateSync `json:"zhima_merchant_zmgo_cumulate_sync_response"`
	AlipayCertSn string                         `json:"alipay_cert_sn,omitempty"`
	SignData     string                         `json:"-"`
	Sign         string                         `json:"sign"`
}

type ZhimaMerchantZmgoCumulateQueryRsp struct {
	Response     *ZhimaMerchantZmgoCumulateQuery `json:"zhima_merchant_zmgo_cumulate_query_response"`
	AlipayCertSn string                          `json:"alipay_cert_sn,omitempty"`
	SignData     string                          `json:"-"`
	Sign         string                          `json:"sign"`
}

type ZhimaCreditPeZmgoBizoptCloseRsp struct {
	Response     *ZhimaCreditPeZmgoBizoptClose `json:"zhima_credit_pe_zmgo_bizopt_close_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type ZhimaCreditPeZmgoSettleRefundRsp struct {
	Response     *ZhimaCreditPeZmgoSettleRefund `json:"zhima_credit_pe_zmgo_settle_refund_response"`
	AlipayCertSn string                         `json:"alipay_cert_sn,omitempty"`
	SignData     string                         `json:"-"`
	Sign         string                         `json:"sign"`
}

type ZhimaCreditPeZmgoPreorderCreateRsp struct {
	Response     *ZhimaCreditPeZmgoPreorderCreate `json:"zhima_credit_pe_zmgo_preorder_create_response"`
	AlipayCertSn string                           `json:"alipay_cert_sn,omitempty"`
	SignData     string                           `json:"-"`
	Sign         string                           `json:"sign"`
}

type ZhimaCreditPeZmgoAgreementUnsignRsp struct {
	Response     *ZhimaCreditPeZmgoAgreementUnsign `json:"zhima_credit_pe_zmgo_agreement_unsign_response"`
	AlipayCertSn string                            `json:"alipay_cert_sn,omitempty"`
	SignData     string                            `json:"-"`
	Sign         string                            `json:"sign"`
}

type ZhimaCreditPeZmgoAgreementQueryRsp struct {
	Response     *ZhimaCreditPeZmgoAgreementQuery `json:"zhima_credit_pe_zmgo_agreement_query_response"`
	AlipayCertSn string                           `json:"alipay_cert_sn,omitempty"`
	SignData     string                           `json:"-"`
	Sign         string                           `json:"sign"`
}

type ZhimaCreditPeZmgoSettleUnfreezeRsp struct {
	Response     *ZhimaCreditPeZmgoSettleUnfreeze `json:"zhima_credit_pe_zmgo_settle_unfreeze_response"`
	AlipayCertSn string                           `json:"alipay_cert_sn,omitempty"`
	SignData     string                           `json:"-"`
	Sign         string                           `json:"sign"`
}

type ZhimaCreditPeZmgoPaysignApplyRsp struct {
	Response     *ZhimaCreditPeZmgoPaysignApply `json:"zhima_credit_pe_zmgo_paysign_apply_response"`
	AlipayCertSn string                         `json:"alipay_cert_sn,omitempty"`
	SignData     string                         `json:"-"`
	Sign         string                         `json:"sign"`
}

type ZhimaCreditPeZmgoPaysignConfirmRsp struct {
	Response     *ZhimaCreditPeZmgoPaysignConfirm `json:"zhima_credit_pe_zmgo_paysign_confirm_response"`
	AlipayCertSn string                           `json:"alipay_cert_sn,omitempty"`
	SignData     string                           `json:"-"`
	Sign         string                           `json:"sign"`
}

type ZhimaCustomerJobworthAdapterQueryRsp struct {
	Response     *ZhimaCustomerJobworthAdapterQuery `json:"zhima_customer_jobworth_adapter_query_response"`
	AlipayCertSn string                             `json:"alipay_cert_sn,omitempty"`
	SignData     string                             `json:"-"`
	Sign         string                             `json:"sign"`
}

type ZhimaCustomerJobworthSceneUseRsp struct {
	Response     *ZhimaCustomerJobworthSceneUse `json:"zhima_customer_jobworth_scene_use_response"`
	AlipayCertSn string                         `json:"alipay_cert_sn,omitempty"`
	SignData     string                         `json:"-"`
	Sign         string                         `json:"sign"`
}

// =========================================================分割=========================================================

type ScoreGet struct {
	ErrorResponse
	BizNo   string `json:"biz_no,omitempty"`
	ZmScore string `json:"zm_score,omitempty"`
}

type ZhimaCreditEpSceneRatingInitialize struct {
	ErrorResponse
	OrderNo string `json:"order_no"`
}

type ZhimaCreditEpSceneFulfillmentSync struct {
	ErrorResponse
	FulfillmentOrderNo string `json:"fulfillment_order_no"`
}

type ZhimaCreditEpSceneAgreementUse struct {
	ErrorResponse
	CreditOrderNo string `json:"credit_order_no"`
}

type ZhimaCreditEpSceneAgreementCancel struct {
	ErrorResponse
	CreditOrderNo string `json:"credit_order_no"`
}

type ZhimaCreditEpSceneFulfillmentlistSync struct {
	ErrorResponse
	FulfillmentResultList []*FulfillmentResult `json:"fulfillment_result_list"`
}

type FulfillmentResult struct {
	FulfillmentOrderNo string `json:"fulfillment_order_no"`
	OutOrderNo         string `json:"out_order_no"`
}

type ZhimaCreditPeZmgoCumulationSync struct {
	ErrorResponse
	OutBizNo     string `json:"out_biz_no,omitempty"`
	AagreementNo string `json:"aagreement_no,omitempty"`
	UserId       string `json:"user_id,omitempty"`
	OpenId       string `json:"open_id,omitempty"`
	FailReason   string `json:"fail_reason,omitempty"`
}

type ZhimaMerchantZmgoCumulateSync struct {
	ErrorResponse
	AgreementId string `json:"agreement_id"`
	OutBizNo    string `json:"out_biz_no"`
	FailReason  string `json:"fail_reason,omitempty"`
}

type ZhimaMerchantZmgoCumulateQuery struct {
	ErrorResponse
	AgreementId        string                `json:"agreement_id"`
	AggrAmount         string                `json:"aggr_amount"`
	AggrTimes          int64                 `json:"aggr_times"`
	AggrDiscountAmount string                `json:"aggr_discount_amount"`
	PageNo             int64                 `json:"page_no"`
	PageSize           int64                 `json:"page_size"`
	DetailList         []*CumulateDataDetail `json:"detail_list,omitempty"`
	FailReason         string                `json:"fail_reason,omitempty"`
}

type CumulateDataDetail struct {
	OutBizNo       string `json:"out_biz_no,omitempty"`
	ReferOutBizNo  string `json:"refer_out_biz_no,omitempty"`
	BizTime        string `json:"biz_time,omitempty"`
	ActionType     string `json:"action_type,omitempty"`
	DataType       string `json:"data_type,omitempty"`
	SubDataType    string `json:"sub_data_type,omitempty"`
	TaskDesc       string `json:"task_desc,omitempty"`
	TaskAmount     string `json:"task_amount,omitempty"`
	TaskTimes      int64  `json:"task_times,omitempty"`
	DiscountDesc   string `json:"discount_desc,omitempty"`
	DiscountAmount string `json:"discount_amount,omitempty"`
}

type ZhimaCreditPeZmgoBizoptClose struct {
	ErrorResponse
	UserId       string `json:"user_id"`
	OpenId       string `json:"open_id,omitempty"`
	BizOptNo     string `json:"biz_opt_no,omitempty"`
	PartnerId    string `json:"partner_id"`
	OutRequestNo string `json:"out_request_no"`
}

type ZhimaCreditPeZmgoSettleRefund struct {
	ErrorResponse
	OutRequestNo   string `json:"out_request_no"`
	WithholdPlanNo string `json:"withhold_plan_no"`
	RefundAmount   string `json:"refund_amount"`
	FailReason     string `json:"fail_reason,omitempty"`
	Retry          bool   `json:"retry,omitempty"`
}

type ZhimaCreditPeZmgoPreorderCreate struct {
	ErrorResponse
	PreorderNo string `json:"preorder_no"`
	PartnerId  string `json:"partner_id"`
	BizType    string `json:"biz_type"`
}

type ZhimaCreditPeZmgoAgreementUnsign struct {
	ErrorResponse
	AgreementId    string `json:"agreement_id"`
	WithholdPlanNo string `json:"withhold_plan_no"`
}

type ZhimaCreditPeZmgoAgreementQuery struct {
	ErrorResponse
	AgreementId     string `json:"agreement_id"`
	AgreementName   string `json:"agreement_name"`
	AlipayUserId    string `json:"alipay_user_id"`
	OpenId          string `json:"open_id,omitempty"`
	AgreementStatus string `json:"agreement_status"`
}

type ZhimaCreditPeZmgoSettleUnfreeze struct {
	ErrorResponse
	UnfreezeStatus string `json:"unfreeze_status"`
	FailReaseon    string `json:"fail_reaseon,omitempty"`
	Retry          string `json:"retry,omitempty"`
	UnfreezeAmount string `json:"unfreeze_amount,omitempty"`
}

type ZhimaCreditPeZmgoPaysignApply struct {
	ErrorResponse
	BizType     string `json:"biz_type"`
	ZmgoOptNo   string `json:"zmgo_opt_no,omitempty"`
	Idempotent  bool   `json:"idempotent,omitempty"`
	AgreementId string `json:"agreement_id,omitempty"`
}

type ZhimaCreditPeZmgoPaysignConfirm struct {
	ErrorResponse
	AgreementId string `json:"agreement_id,omitempty"`
}

type ZhimaCustomerJobworthAdapterQuery struct {
	ErrorResponse
	AdapterScore string `json:"adapter_score,omitempty"`
	SubCode      string `json:"sub_code,omitempty"`
	SubMsg       string `json:"sub_msg,omitempty"`
	Url          string `json:"url,omitempty"`
}

type ZhimaCustomerJobworthSceneUse struct {
	ErrorResponse
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}
