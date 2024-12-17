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

type ZhimaMerchantZmgoTemplateCreateRsp struct {
	Response     *ZhimaMerchantZmgoTemplateCreate `json:"zhima_merchant_zmgo_template_create_response"`
	AlipayCertSn string                           `json:"alipay_cert_sn,omitempty"`
	SignData     string                           `json:"-"`
	Sign         string                           `json:"sign"`
}

type ZhimaMerchantZmgoTemplateQueryRsp struct {
	Response     *ZhimaMerchantZmgoTemplateQuery `json:"zhima_merchant_zmgo_template_query_response"`
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

type ZhimaCustomerJobworthAuthQueryRsp struct {
	Response     *ZhimaCustomerJobworthAuthQuery `json:"zhima_customer_jobworth_authentication_query_response"`
	AlipayCertSn string                          `json:"alipay_cert_sn,omitempty"`
	SignData     string                          `json:"-"`
	Sign         string                          `json:"sign"`
}

type ZhimaCustomerJobworthAuthPreConsultRsp struct {
	Response     *ZhimaCustomerJobworthAuthPreConsult `json:"zhima_customer_jobworth_authentication_preconsult_response"`
	AlipayCertSn string                               `json:"alipay_cert_sn,omitempty"`
	SignData     string                               `json:"-"`
	Sign         string                               `json:"sign"`
}

type ZhimaCreditPeZmgoSettleApplyRsp struct {
	Response     *ZhimaCreditPeZmgoSettleApply `json:"zhima_credit_pe_zmgo_settle_apply_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type ZhimaCreditPayAfterUseAgreementQueryRsp struct {
	Response     *ZhimaCreditPayAfterUseAgreementQuery `json:"zhima_credit_payafteruse_creditagreement_query_response"`
	AlipayCertSn string                                `json:"alipay_cert_sn,omitempty"`
	SignData     string                                `json:"-"`
	Sign         string                                `json:"sign"`
}

type ZhimaCreditPayAfterUseCreditBizOrderRsp struct {
	Response     *ZhimaCreditPayAfterUseCreditBizOrder `json:"zhima_credit_payafteruse_creditbizorder_order_response"`
	AlipayCertSn string                                `json:"alipay_cert_sn,omitempty"`
	SignData     string                                `json:"-"`
	Sign         string                                `json:"sign"`
}

type ZhimaCreditPayAfterUseCreditBizOrderQueryRsp struct {
	Response     *ZhimaCreditPayAfterUseCreditBizOrderQuery `json:"zhima_credit_payafteruse_creditbizorder_query_response"`
	AlipayCertSn string                                     `json:"alipay_cert_sn,omitempty"`
	SignData     string                                     `json:"-"`
	Sign         string                                     `json:"sign"`
}

type ZhimaCreditPayAfterUseCreditBizOrderFinishRsp struct {
	Response     *ZhimaCreditPayAfterUseCreditBizOrderFinish `json:"zhima_credit_payafteruse_creditbizorder_finish_response"`
	AlipayCertSn string                                      `json:"alipay_cert_sn,omitempty"`
	SignData     string                                      `json:"-"`
	Sign         string                                      `json:"sign"`
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

type ZhimaMerchantZmgoTemplateCreate struct {
	ErrorResponse
	TemplateNo string `json:"template_no"`
}

type ZhimaMerchantZmgoTemplateQuery struct {
	ErrorResponse
	BasicConfig      *BasicConfig      `json:"basic_config"`
	ObligationConfig *ObligationConfig `json:"obligation_config"`
	RightConfig      *RightConfig      `json:"right_config"`
	OpenConfig       *OpenConfig       `json:"open_config"`
	SettlementConfig *SettlementConfig `json:"settlement_config"`
	QuitConfig       *QuitConfig       `json:"quit_config"`
	ExtConfig        *ExtConfig        `json:"ext_config"`
}

type BasicConfig struct {
	TemplateName       string `json:"template_name"`
	PartnerId          string `json:"partner_id"`
	IsvPid             string `json:"isv_pid"`
	BizType            string `json:"biz_type"`
	OutBizNo           string `json:"out_biz_no"`
	MerchantCustomLogo string `json:"merchant_custom_logo"`
	Contact            string `json:"contact"`
	TemplateNo         string `json:"template_no"`
}

type ObligationConfig struct {
	ObligationTemplate         string `json:"obligation_template"`
	ObligationTimes            string `json:"obligation_times"`
	ObligationAmount           int    `json:"obligation_amount"`
	PromiseTypeDesc            string `json:"promise_type_desc"`
	BenefitUrl                 string `json:"benefit_url"`
	TaskProgressRedirectSchema string `json:"task_progress_redirect_schema"`
}

type RightConfig struct {
	CustomBenefitDesc                    string `json:"custom_benefit_desc"`
	CustomSubBenefitDesc                 string `json:"custom_sub_benefit_desc"`
	CumulativePreferentialRedirectSchema string `json:"cumulative_preferential_redirect_schema"`
}

type OpenConfig struct {
	FreezeAmount              string `json:"freeze_amount"`
	PeriodMode                string `json:"period_mode"`
	PeriodTime                string `json:"period_time"`
	AppointDate               string `json:"appoint_date"`
	MinSignInterval           string `json:"min_sign_interval"`
	SupportExpireDeferral     bool   `json:"support_expire_deferral"`
	CustomOpenTipList         string `json:"custom_open_tip_list"`
	CardColorScheme           string `json:"card_color_scheme"`
	CustomOpenTips            string `json:"custom_open_tips"`
	ShowSignSuccessPage       bool   `json:"show_sign_success_page"`
	SignSuccessTaskButtonDesc string `json:"sign_success_task_button_desc"`
	SignAgainSchema           string `json:"sign_again_schema"`
	ApplyButtonDesc           string `json:"apply_button_desc"`
}

type SettlementConfig struct {
	SettlementType                   string                   `json:"settlement_type"`
	CustomFeeName                    string                   `json:"custom_fee_name"`
	CycleFlexWithholdConfig          *CycleFlexWithholdConfig `json:"cycle_flex_withhold_config"`
	FulfilltimesCustomSettlementPlan string                   `json:"fulfilltimes_custom_settlement_plan"`
	ExpStopTimeMode                  string                   `json:"exp_stop_time_mode"`
	ExpStopTime                      string                   `json:"exp_stop_time"`
	ExpStopDelayDays                 int                      `json:"exp_stop_delay_days"`
	CycleWithholdConfig              *CycleWithholdConfig     `json:"cycle_withhold_config"`
}

type CycleFlexWithholdConfig struct {
	CycleFlexWithholdTotalPeriodCount int    `json:"cycle_flex_withhold_total_period_count"`
	CycleFlexWithholdMaxPrice         int    `json:"cycle_flex_withhold_max_price"`
	CycleFlexWithholdFeeName          string `json:"cycle_flex_withhold_fee_name"`
}

type CycleWithholdConfig struct {
	WithholdMode                 string   `json:"withhold_mode"`
	PeriodType                   string   `json:"period_type"`
	Period                       string   `json:"period"`
	SupportCycleWithholdHighMode bool     `json:"support_cycle_withhold_high_mode"`
	DeductPlan                   []string `json:"deduct_plan"`
	SupportExemptionPeriod       bool     `json:"support_exemption_period"`
	ExemptionPeriod              string   `json:"exemption_period"`
}

type QuitConfig struct {
	QuitType    string `json:"quit_type"`
	QuitJumpUrl string `json:"quit_jump_url"`
	QuitDesc    string `json:"quit_desc"`
}

type ExtConfig struct {
	TextContentFillRuleId   string `json:"text_content_fill_rule_id"`
	TextContentFillVariable string `json:"text_content_fill_variable"`
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
	Url          string `json:"url,omitempty"`
}

type ZhimaCustomerJobworthSceneUse struct {
	ErrorResponse
}

type ZhimaCustomerJobworthAuthQuery struct {
	ErrorResponse
	IdentityResult        string `json:"identity_result"`
	IdentityResultSkipUrl string `json:"identity_result_skip_url"`
	TokenStatus           string `json:"token_status"`
	UserId                string `json:"user_id"`
	OpenId                string `json:"open_id"`
}

type ZhimaCustomerJobworthAuthPreConsult struct {
	ErrorResponse
	PreCheckSuccess bool   `json:"pre_check_success"`
	Reason          string `json:"reason"`
}

type ZhimaCreditPeZmgoSettleApply struct {
	ErrorResponse
	OutRequestNo   string `json:"out_request_no"`
	AgreementId    string `json:"agreement_id"`
	WithholdPlanNo string `json:"withhold_plan_no"`
	SettleStatus   string `json:"settle_status"`
	FailReason     string `json:"fail_reason"`
}

type ZhimaCreditPayAfterUseAgreementQuery struct {
	ErrorResponse
	OutAgreementNo    string `json:"out_agreement_no"`
	AgreementStatus   string `json:"agreement_status"`
	CreditAgreementId string `json:"credit_agreement_id"`
	BizTime           string `json:"biz_time"`
	ExtInfo           string `json:"ext_info"`
	AlipayUserId      string `json:"alipay_user_id"`
	OpenId            string `json:"open_id"`
}

type ZhimaCreditPayAfterUseCreditBizOrder struct {
	ErrorResponse
	OutOrderNo       string `json:"out_order_no"`
	CreditBizOrderId string `json:"credit_biz_order_id"`
}

type ZhimaCreditPayAfterUseCreditBizOrderQuery struct {
	ErrorResponse
	CreditBizOrderId  string `json:"credit_biz_order_id"`
	CreditAgreementId string `json:"credit_agreement_id"`
	TotalAmount       string `json:"total_amount"`
	CreateTime        string `json:"create_time"`
	ZmServiceId       string `json:"zm_service_id"`
	ProductCode       string `json:"product_code"`
	OrderStatus       string `json:"order_status"`
	TradeNo           string `json:"trade_no"`
}

type ZhimaCreditPayAfterUseCreditBizOrderFinish struct {
	ErrorResponse
	OutRequestNo     string `json:"out_request_no"`
	CreditBizOrderId string `json:"credit_biz_order_id"`
	OrderStatus      string `json:"order_status"`
}
