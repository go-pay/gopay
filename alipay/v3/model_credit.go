package alipay

type ZmGoPreorderCreateRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	OrderStr string `json:"orderStr"`
}

type ZmGoCumulateSyncRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	AgreementId string `json:"agreement_id"`
	OutBizNo    string `json:"out_biz_no"`
	FailReason  string `json:"fail_reason"`
}

type ZmGoCumulateQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	AgreementId        string `json:"agreement_id"`
	AggrAmount         int64  `json:"aggr_amount"`
	AggrTimes          int64  `json:"aggr_times"`
	AggrDiscountAmount int64  `json:"aggr_discount_amount"`
	PageNo             int64  `json:"page_no"`
	PageSize           int64  `json:"page_size"`
}

type ZmGoSettleApplyRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	OutRequestNo   string `json:"out_request_no"`
	AgreementId    string `json:"agreement_id"`
	WithholdPlanNo string `json:"withhold_plan_no"`
	SettleStatus   string `json:"settle_status"`
	FailReason     string `json:"fail_reason"`
}

type ZmGoSettleRefundRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	OutRequestNo   string `json:"out_request_no"`
	WithholdPlanNo string `json:"withhold_plan_no"`
	RefundAmount   string `json:"refund_amount"`
	RefundOptNo    string `json:"refund_opt_no"`
	FailReason     string `json:"fail_reason"`
	Retry          bool   `json:"retry"`
}

type ZmGoAgreementQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	AgreementId     string `json:"agreement_id"`
	AgreementName   string `json:"agreement_name"`
	AgreementStatus string `json:"agreement_status"`
	BizType         string `json:"biz_type"`
	SignTime        string `json:"sign_time"`
	StartTime       string `json:"start_time"`
}

type ZmGoAgreementQueryUnsignRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	AgreementId    string `json:"agreement_id"`
	WithholdPlanNo string `json:"withhold_plan_no"`
}

type ZmGoTemplateCreateRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	TemplateNo string `json:"template_no"`
}

type ZmGoTemplateQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	BasicConfig struct {
		TemplateName       string `json:"template_name"`
		PartnerId          string `json:"partner_id"`
		IsvPid             string `json:"isv_pid"`
		BizType            string `json:"biz_type"`
		OutBizNo           string `json:"out_biz_no"`
		MerchantCustomLogo string `json:"merchant_custom_logo"`
		Contact            string `json:"contact"`
		TemplateNo         string `json:"template_no"`
	} `json:"basic_config"`
	ObligationConfig struct {
		ObligationTemplate         string  `json:"obligation_template"`
		ObligationTimes            string  `json:"obligation_times"`
		ObligationAmount           float64 `json:"obligation_amount"`
		PromiseTypeDesc            string  `json:"promise_type_desc"`
		BenefitUrl                 string  `json:"benefit_url"`
		TaskProgressRedirectSchema string  `json:"task_progress_redirect_schema"`
	} `json:"obligation_config"`
	RightConfig struct {
		CustomBenefitDesc                    string `json:"custom_benefit_desc"`
		CustomSubBenefitDesc                 string `json:"custom_sub_benefit_desc"`
		CumulativePreferentialRedirectSchema string `json:"cumulative_preferential_redirect_schema"`
	} `json:"right_config"`
	OpenConfig struct {
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
	} `json:"open_config"`
	SettlementConfig struct {
		SettlementType          string `json:"settlement_type"`
		CustomFeeName           string `json:"custom_fee_name"`
		CycleFlexWithholdConfig struct {
			CycleFlexWithholdTotalPeriodCount int     `json:"cycle_flex_withhold_total_period_count"`
			CycleFlexWithholdMaxPrice         float64 `json:"cycle_flex_withhold_max_price"`
			CycleFlexWithholdFeeName          string  `json:"cycle_flex_withhold_fee_name"`
		} `json:"cycle_flex_withhold_config"`
		FulfilltimesCustomSettlementPlan string `json:"fulfilltimes_custom_settlement_plan"`
		ExpStopTimeMode                  string `json:"exp_stop_time_mode"`
		ExpStopTime                      string `json:"exp_stop_time"`
		ExpStopDelayDays                 int    `json:"exp_stop_delay_days"`
		CycleWithholdConfig              struct {
			WithholdMode                 string   `json:"withhold_mode"`
			PeriodType                   string   `json:"period_type"`
			Period                       string   `json:"period"`
			SupportCycleWithholdHighMode bool     `json:"support_cycle_withhold_high_mode"`
			DeductPlan                   []string `json:"deduct_plan"`
			SupportExemptionPeriod       bool     `json:"support_exemption_period"`
			ExemptionPeriod              string   `json:"exemption_period"`
		} `json:"cycle_withhold_config"`
	} `json:"settlement_config"`
	QuitConfig struct {
		QuitType    string `json:"quit_type"`
		QuitJumpUrl string `json:"quit_jump_url"`
		QuitDesc    string `json:"quit_desc"`
	} `json:"quit_config"`
	ExtConfig struct {
		TextContentFillRuleId   string `json:"text_content_fill_rule_id"`
		TextContentFillVariable string `json:"text_content_fill_variable"`
	} `json:"ext_config"`
}
