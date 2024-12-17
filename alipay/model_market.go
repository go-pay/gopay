package alipay

type MarketingCampaignCashCreateRsp struct {
	Response     *MarketingCampaignCashCreate `json:"alipay_marketing_campaign_cash_create_response"`
	AlipayCertSn string                       `json:"alipay_cert_sn,omitempty"`
	SignData     string                       `json:"-"`
	Sign         string                       `json:"sign"`
}

type MarketingCampaignCashTriggerRsp struct {
	Response     *MarketingCampaignCashTrigger `json:"alipay_marketing_campaign_cash_trigger_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type MarketingCampaignCashStatusModifyRsp struct {
	Response     *MarketingCampaignCashStatusModify `json:"alipay_marketing_campaign_cash_status_modify_response"`
	AlipayCertSn string                             `json:"alipay_cert_sn,omitempty"`
	SignData     string                             `json:"-"`
	Sign         string                             `json:"sign"`
}

type MarketingCampaignCashListQueryRsp struct {
	Response     *MarketingCampaignCashListQuery `json:"alipay_marketing_campaign_cash_list_query_response"`
	AlipayCertSn string                          `json:"alipay_cert_sn,omitempty"`
	SignData     string                          `json:"-"`
	Sign         string                          `json:"sign"`
}

type MarketingCampaignCashDetailQueryRsp struct {
	Response     *MarketingCampaignCashDetailQuery `json:"alipay_marketing_campaign_cash_detail_query_response"`
	AlipayCertSn string                            `json:"alipay_cert_sn,omitempty"`
	SignData     string                            `json:"-"`
	Sign         string                            `json:"sign"`
}

type MarketingActivityDeliveryStopRsp struct {
	Response     *MarketingActivityDeliveryStop `json:"alipay_marketing_activity_delivery_stop_response"`
	AlipayCertSn string                         `json:"alipay_cert_sn,omitempty"`
	SignData     string                         `json:"-"`
	Sign         string                         `json:"sign"`
}

type MarketingActivityDeliveryQueryRsp struct {
	Response     *MarketingActivityDeliveryQuery `json:"alipay_marketing_activity_delivery_query_response"`
	AlipayCertSn string                          `json:"alipay_cert_sn,omitempty"`
	SignData     string                          `json:"-"`
	Sign         string                          `json:"sign"`
}

type MarketingActivityDeliveryCreateRsp struct {
	Response     *MarketingActivityDeliveryCreate `json:"alipay_marketing_activity_delivery_create_response"`
	AlipayCertSn string                           `json:"alipay_cert_sn,omitempty"`
	SignData     string                           `json:"-"`
	Sign         string                           `json:"sign"`
}

type MarketingMaterialImageUploadRsp struct {
	Response     *MarketingMaterialImageUpload `json:"alipay_marketing_material_image_upload_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type MarketingActivityVoucherCreateRsp struct {
	Response     *MarketingActivityVoucherCreate `json:"alipay_marketing_activity_voucher_create_response"`
	AlipayCertSn string                          `json:"alipay_cert_sn,omitempty"`
	SignData     string                          `json:"-"`
	Sign         string                          `json:"sign"`
}

type MarketingActivityVoucherPublishRsp struct {
	Response     *MarketingActivityVoucherPublish `json:"alipay_marketing_activity_voucher_publish_response"`
	AlipayCertSn string                           `json:"alipay_cert_sn,omitempty"`
	SignData     string                           `json:"-"`
	Sign         string                           `json:"sign"`
}

type MarketingActivityVoucherQueryRsp struct {
	Response     *MarketingActivityVoucherQuery `json:"alipay_marketing_activity_voucher_query_response"`
	AlipayCertSn string                         `json:"alipay_cert_sn,omitempty"`
	SignData     string                         `json:"-"`
	Sign         string                         `json:"sign"`
}

type MarketingActivityVoucherModifyRsp struct {
	Response     *MarketingActivityVoucherModify `json:"alipay_marketing_activity_voucher_modify_response"`
	AlipayCertSn string                          `json:"alipay_cert_sn,omitempty"`
	SignData     string                          `json:"-"`
	Sign         string                          `json:"sign"`
}

type MarketingActivityVoucherAppendRsp struct {
	Response     *MarketingActivityVoucherAppend `json:"alipay_marketing_activity_voucher_append_response"`
	AlipayCertSn string                          `json:"alipay_cert_sn,omitempty"`
	SignData     string                          `json:"-"`
	Sign         string                          `json:"sign"`
}

type MarketingActivityVoucherStopRsp struct {
	Response     *MarketingActivityVoucherStop `json:"alipay_marketing_activity_voucher_stop_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type MarketingActivityBatchQueryRsp struct {
	Response     *MarketingActivityBatchQuery `json:"alipay_marketing_activity_batchquery_response"`
	AlipayCertSn string                       `json:"alipay_cert_sn,omitempty"`
	SignData     string                       `json:"-"`
	Sign         string                       `json:"sign"`
}

type MarketingActivityConsultRsp struct {
	Response     *MarketingActivityConsult `json:"alipay_marketing_activity_consult_response"`
	AlipayCertSn string                    `json:"alipay_cert_sn,omitempty"`
	SignData     string                    `json:"-"`
	Sign         string                    `json:"sign"`
}

type MarketingActivityQueryRsp struct {
	Response     *MarketingActivityQuery `json:"alipay_marketing_activity_query_response"`
	AlipayCertSn string                  `json:"alipay_cert_sn,omitempty"`
	SignData     string                  `json:"-"`
	Sign         string                  `json:"sign"`
}

type MarketingActivityQueryMerchantBatchQueryRsp struct {
	Response     *MarketingActivityQueryMerchantBatchQuery `json:"alipay_marketing_activity_merchant_batchquery_response"`
	AlipayCertSn string                                    `json:"alipay_cert_sn,omitempty"`
	SignData     string                                    `json:"-"`
	Sign         string                                    `json:"sign"`
}

type MarketingActivityQueryAppBatchQueryRsp struct {
	Response     *MarketingActivityQueryAppBatchQuery `json:"alipay_marketing_activity_app_batchquery_response"`
	AlipayCertSn string                               `json:"alipay_cert_sn,omitempty"`
	SignData     string                               `json:"-"`
	Sign         string                               `json:"sign"`
}

type MarketingActivityQueryShopBatchQueryRsp struct {
	Response     *MarketingActivityQueryShopBatchQuery `json:"alipay_marketing_activity_shop_batchquery_response"`
	AlipayCertSn string                                `json:"alipay_cert_sn,omitempty"`
	SignData     string                                `json:"-"`
	Sign         string                                `json:"sign"`
}

type MarketingActivityQueryGoodsBatchQueryRsp struct {
	Response     *MarketingActivityQueryGoodsBatchQuery `json:"alipay_marketing_activity_goods_batchquery_response"`
	AlipayCertSn string                                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                                 `json:"-"`
	Sign         string                                 `json:"sign"`
}

type MarketingActivityQueryUserBatchQueryVoucherRsp struct {
	Response     *MarketingActivityQueryUserBatchQueryVoucher `json:"alipay_marketing_activity_user_batchqueryvoucher_response"`
	AlipayCertSn string                                       `json:"alipay_cert_sn,omitempty"`
	SignData     string                                       `json:"-"`
	Sign         string                                       `json:"sign"`
}

type MarketingActivityQueryUserQueryVoucherRsp struct {
	Response     *MarketingActivityQueryUserQueryVoucher `json:"alipay_marketing_activity_user_queryvoucher_response"`
	AlipayCertSn string                                  `json:"alipay_cert_sn,omitempty"`
	SignData     string                                  `json:"-"`
	Sign         string                                  `json:"sign"`
}

type MarketingCampaignOrderVoucherConsultRsp struct {
	Response     *MarketingCampaignOrderVoucherConsult `json:"alipay_marketing_campaign_order_voucher_consult_response"`
	AlipayCertSn string                                `json:"alipay_cert_sn,omitempty"`
	SignData     string                                `json:"-"`
	Sign         string                                `json:"sign"`
}

type MarketingActivityOrderVoucherCreateRsp struct {
	Response     *MarketingActivityOrderVoucherCreate `json:"alipay_marketing_activity_ordervoucher_create_response"`
	AlipayCertSn string                               `json:"alipay_cert_sn,omitempty"`
	SignData     string                               `json:"-"`
	Sign         string                               `json:"sign"`
}

type MarketingActivityOrderVoucherCodeDepositRsp struct {
	Response     *MarketingActivityOrderVoucherCodeDeposit `json:"alipay_marketing_activity_ordervoucher_codedeposit_response"`
	AlipayCertSn string                                    `json:"alipay_cert_sn,omitempty"`
	SignData     string                                    `json:"-"`
	Sign         string                                    `json:"sign"`
}

type MarketingActivityOrderVoucherModifyRsp struct {
	Response     *MarketingActivityOrderVoucherModify `json:"alipay_marketing_activity_ordervoucher_modify_response"`
	AlipayCertSn string                               `json:"alipay_cert_sn,omitempty"`
	SignData     string                               `json:"-"`
	Sign         string                               `json:"sign"`
}

type MarketingActivityOrderVoucherStopRsp struct {
	Response     *MarketingActivityOrderVoucherStop `json:"alipay_marketing_activity_ordervoucher_stop_response"`
	AlipayCertSn string                             `json:"alipay_cert_sn,omitempty"`
	SignData     string                             `json:"-"`
	Sign         string                             `json:"sign"`
}

type MarketingActivityOrderVoucherAppendRsp struct {
	Response     *MarketingActivityOrderVoucherAppend `json:"alipay_marketing_activity_ordervoucher_append_response"`
	AlipayCertSn string                               `json:"alipay_cert_sn,omitempty"`
	SignData     string                               `json:"-"`
	Sign         string                               `json:"sign"`
}

type MarketingActivityOrderVoucherUseRsp struct {
	Response     *MarketingActivityOrderVoucherUse `json:"alipay_marketing_activity_ordervoucher_use_response"`
	AlipayCertSn string                            `json:"alipay_cert_sn,omitempty"`
	SignData     string                            `json:"-"`
	Sign         string                            `json:"sign"`
}

type MarketingActivityOrderVoucherRefundRsp struct {
	Response     *MarketingActivityOrderVoucherRefund `json:"alipay_marketing_activity_ordervoucher_refund_response"`
	AlipayCertSn string                               `json:"alipay_cert_sn,omitempty"`
	SignData     string                               `json:"-"`
	Sign         string                               `json:"sign"`
}

type MarketingActivityOrderVoucherQueryRsp struct {
	Response     *MarketingActivityOrderVoucherQuery `json:"alipay_marketing_activity_ordervoucher_query_response"`
	AlipayCertSn string                              `json:"alipay_cert_sn,omitempty"`
	SignData     string                              `json:"-"`
	Sign         string                              `json:"sign"`
}

type MarketingActivityOrderVoucherCodeCountRsp struct {
	Response     *MarketingActivityOrderVoucherCodeCount `json:"alipay_marketing_activity_ordervoucher_codecount_response"`
	AlipayCertSn string                                  `json:"alipay_cert_sn,omitempty"`
	SignData     string                                  `json:"-"`
	Sign         string                                  `json:"sign"`
}

type MarketingCardTemplateCreateRsp struct {
	Response     *MarketingCardTemplateCreate `json:"alipay_marketing_card_template_create_response"`
	AlipayCertSn string                       `json:"alipay_cert_sn,omitempty"`
	SignData     string                       `json:"-"`
	Sign         string                       `json:"sign"`
}

type MarketingCardTemplateModifyRsp struct {
	Response     *MarketingCardTemplateModify `json:"alipay_marketing_card_template_modify_response"`
	AlipayCertSn string                       `json:"alipay_cert_sn,omitempty"`
	SignData     string                       `json:"-"`
	Sign         string                       `json:"sign"`
}

type MarketingCardTemplateQueryRsp struct {
	Response     *MarketingCardTemplateQuery `json:"alipay_marketing_card_template_query_response"`
	AlipayCertSn string                      `json:"alipay_cert_sn,omitempty"`
	SignData     string                      `json:"-"`
	Sign         string                      `json:"sign"`
}

type MarketingCardUpdateRsp struct {
	Response     *MarketingCardUpdate `json:"alipay_marketing_card_update_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type MarketingCardQueryRsp struct {
	Response     *MarketingCardQuery `json:"alipay_marketing_card_query_response"`
	AlipayCertSn string              `json:"alipay_cert_sn,omitempty"`
	SignData     string              `json:"-"`
	Sign         string              `json:"sign"`
}

type MarketingCardDeleteRsp struct {
	Response     *MarketingCardDelete `json:"alipay_marketing_card_delete_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type MarketingCardMessageNotifyRsp struct {
	Response     *MarketingCardMessageNotify `json:"alipay_marketing_card_message_notify_response"`
	AlipayCertSn string                      `json:"alipay_cert_sn,omitempty"`
	SignData     string                      `json:"-"`
	Sign         string                      `json:"sign"`
}

type MarketingCardFormTemplateSetRsp struct {
	Response     *MarketingCardFormTemplateSet `json:"alipay_marketing_card_formtemplate_set_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type OfflineMaterialImageUploadRsp struct {
	Response     *OfflineMaterialImageUpload `json:"alipay_offline_material_image_upload_response"`
	AlipayCertSn string                      `json:"alipay_cert_sn,omitempty"`
	SignData     string                      `json:"-"`
	Sign         string                      `json:"sign"`
}

// =========================================================分割=========================================================

type MarketingCampaignCashCreate struct {
	ErrorResponse
	CrowdNo       string `json:"crowd_no"`
	PayURL        string `json:"pay_url"`
	OriginCrowdNo string `json:"origin_crowd_no"`
}

type MarketingCampaignCashTrigger struct {
	ErrorResponse
	TriggerResult     string `json:"trigger_result"`
	PrizeAmount       string `json:"prize_amount"`
	RepeatTriggerFlag string `json:"repeat_trigger_flag"`
	PartnerId         string `json:"partner_id"`
	ErrorMsg          string `json:"error_msg"`
	CouponName        string `json:"coupon_name"`
	PrizeMsg          string `json:"prize_msg"`
	MerchantLogo      string `json:"merchant_logo"`
	BizNo             string `json:"biz_no"`
	OutBizNo          string `json:"out_biz_no"`
}

type MarketingCampaignCashStatusModify struct {
	ErrorResponse
}

type MarketingCampaignCashListQuery struct {
	ErrorResponse
	PageSize  string             `json:"page_size"`
	CampList  []CashCampaignInfo `json:"camp_list"`
	PageIndex string             `json:"page_index"`
	TotalSize string             `json:"total_size"`
}

type CashCampaignInfo struct {
	CrowdNo       string `json:"crowd_no"`
	OriginCrowdNo string `json:"origin_crowd_no"`
	CampStatus    string `json:"camp_status"`
	CouponName    string `json:"coupon_name"`
}

type MarketingCampaignCashDetailQuery struct {
	ErrorResponse
	CrowdNo       string  `json:"crowd_no"`
	CouponName    string  `json:"coupon_name"`
	PrizeMsg      string  `json:"prize_msg"`
	PrizeType     string  `json:"prize_type"`
	StartTime     string  `json:"start_time"`
	EndTime       string  `json:"end_time"`
	TotalAmount   float64 `json:"total_amount"`
	SendAmount    float64 `json:"send_amount"`
	TotalNum      int     `json:"total_num"`
	TotalCount    int     `json:"total_count"`
	OriginCrowdNo string  `json:"origin_crowd_no"`
	CampStatus    string  `json:"camp_status"`
}

type MarketingActivityDeliveryStop struct {
	ErrorResponse
	DeliveryId string `json:"delivery_id"`
}

type MarketingActivityDeliveryQuery struct {
	ErrorResponse
	DeliveryId         string              `json:"delivery_id"`
	DeliveryStatus     string              `json:"delivery_status"`
	DeliveryBoothCode  string              `json:"delivery_booth_code"`
	DeliveryErrorMsg   string              `json:"delivery_error_msg"`
	DeliveryBaseInfo   *DeliveryBaseInfo   `json:"delivery_base_info"`
	DeliveryPlayConfig *DeliveryPlayConfig `json:"delivery_play_config"`
	DeliveryTargetRule *DeliveryTargetRule `json:"delivery_target_rule"`
}

type DeliveryBaseInfo struct {
	DeliveryName      string `json:"delivery_name"`
	DeliveryBeginTime string `json:"delivery_begin_time"`
	DeliveryEndTime   string `json:"delivery_end_time"`
}

type DeliveryPlayConfig struct {
	DeliverySingleSendConfig *DeliverySingleSendConfig `json:"delivery_single_send_config"`
	DeliveryFullSendConfig   *DeliveryFullSendConfig   `json:"delivery_full_send_config"`
}

type DeliverySingleSendConfig struct {
	DeliveryContentInfo *DeliveryContentInfo `json:"delivery_content_info"`
}

type DeliveryFullSendConfig struct {
	DeliveryFloorAmount string               `json:"delivery_floor_amount"`
	DeliveryContentInfo *DeliveryContentInfo `json:"delivery_content_info"`
}

type DeliveryContentInfo struct {
	DeliveryContentType     string                   `json:"delivery_content_type"`
	DeliveryActivityContent *DeliveryActivityContent `json:"delivery_activity_content"`
	DeliveryAppContent      *DeliveryAppContent      `json:"delivery_app_content"`
	DeliveryDisplayInfo     *DeliveryDisplayInfo     `json:"delivery_display_info"`
}

type DeliveryActivityContent struct {
	ActivityId string `json:"activity_id"`
}

type DeliveryAppContent struct {
	MiniAppId           string   `json:"mini_app_id"`
	ServiceCodeList     []string `json:"service_code_list"`
	MiniAppDeliveryType string   `json:"mini_app_delivery_type"`
}

type DeliveryDisplayInfo struct {
	MainTitle string `json:"main_title"`
	SubTitle  string `json:"sub_title"`
}

type DeliveryTargetRule struct {
	DeliveryMerchantRule *DeliveryMerchantRule `json:"delivery_merchant_rule"`
	DeliveryCityCodeRule *DeliveryCityCodeRule `json:"delivery_city_code_rule"`
	DeliveryPromoTags    string                `json:"delivery_promo_tags"`
}

type DeliveryMerchantRule struct {
	DeliveryMerchantMode  string                  `json:"delivery_merchant_mode"`
	DeliveryMerchantInfos []*DeliveryMerchantInfo `json:"delivery_merchant_infos"`
}

type DeliveryMerchantInfo struct {
	MerchantIdType string `json:"merchant_id_type"`
	MerchantId     string `json:"merchant_id"`
}

type DeliveryCityCodeRule struct {
	CityCodes []string `json:"city_codes"`
	AllCity   bool     `json:"all_city"`
}

type MarketingActivityDeliveryCreate struct {
	ErrorResponse
	DeliveryId string `json:"delivery_id"`
}

type MarketingMaterialImageUpload struct {
	ErrorResponse
	ResourceId      string `json:"resource_id"`
	ResourceEnhance bool   `json:"resource_enhance"`
	ResourceUrl     string `json:"resource_url"`
}

type MarketingActivityVoucherCreate struct {
	ErrorResponse
	ActivityId string `json:"activity_id"`
}

type MarketingActivityVoucherPublish struct {
	ErrorResponse
	RechargeUrl string `json:"recharge_url"`
}

type MarketingActivityVoucherQuery struct {
	ErrorResponse
	ActivityBaseInfo          *ActivityBaseInfo          `json:"activity_base_info"`
	VoucherSendModeInfo       *VoucherSendModeInfo       `json:"voucher_send_mode_info"`
	VoucherDeductInfo         *VoucherDeductInfo         `json:"voucher_deduct_info"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo `json:"voucher_available_scope_info"`
	VoucherUseRuleInfo        *VoucherUseRuleInfo        `json:"voucher_use_rule_info"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo `json:"voucher_display_pattern_info"`
	VoucherCustomerGuideInfo  *VoucherCustomerGuideInfo  `json:"voucher_customer_guide_info"`
	VoucherBudgetSupplyInfo   *VoucherBudgetSupplyInfo   `json:"voucher_budget_supply_info"`
	VoucherInventoryInfo      *VoucherInventoryInfo      `json:"voucher_inventory_info"`
}

type ActivityBaseInfo struct {
	ActivityId              string              `json:"activity_id"`
	ActivityName            string              `json:"activity_name,omitempty"`
	BelongMerchantInfo      *BelongMerchantInfo `json:"belong_merchant_info"`
	CodeMode                string              `json:"code_mode,omitempty"`
	ActivityOperationStatus string              `json:"activity_operation_status,omitempty"`
	ActivityStatus          string              `json:"activity_status"`
	ActivityProductType     string              `json:"activity_product_type,omitempty"`
}

type BelongMerchantInfo struct {
	MerchantIdType string `json:"merchant_id_type"`
	MerchantId     string `json:"merchant_id"`
}

type VoucherSendModeInfo struct {
	VoucherSendMode        string                  `json:"voucher_send_mode"`
	VoucherSendRuleInfo    *VoucherSendRuleInfo    `json:"voucher_send_rule_info"`
	VoucherSaleModeInfo    *VoucherSaleModeInfo    `json:"voucher_sale_mode_info"`
	VoucherPackageModeInfo *VoucherPackageModeInfo `json:"voucher_package_mode_info"`
}

type VoucherSendRuleInfo struct {
	Quantity                       int    `json:"quantity"`
	MaxQuantityByDay               int    `json:"max_quantity_by_day"`
	QuantityLimitPerUser           int    `json:"quantity_limit_per_user"`
	QuantityLimitPerUserPeriodType string `json:"quantity_limit_per_user_period_type"`
	NaturalPersonLimit             bool   `json:"natural_person_limit"`
	PhoneNumberLimit               bool   `json:"phone_number_limit"`
	RealNameLimit                  bool   `json:"real_name_limit"`
	PublishStartTime               string `json:"publish_start_time"`
	PublishEndTime                 string `json:"publish_end_time"`
}

type VoucherSaleModeInfo struct {
	SaleAmount        string `json:"sale_amount"`
	Refundable        bool   `json:"refundable"`
	OverdueRefundable bool   `json:"overdue_refundable"`
}

type VoucherPackageModeInfo struct {
	VoucherPackageId string `json:"voucher_package_id"`
}

type VoucherDeductInfo struct {
	VoucherType         string               `json:"voucher_type"`
	FixVoucherInfo      *FixVoucherInfo      `json:"fix_voucher_info"`
	DiscountVoucherInfo *DiscountVoucherInfo `json:"discount_voucher_info,omitempty"`
	SpecialVoucherInfo  *SpecialVoucherInfo  `json:"special_voucher_info,omitempty"`
	ExchangeVoucherInfo *ExchangeVoucherInfo `json:"exchange_voucher_info,omitempty"`
}

type FixVoucherInfo struct {
	Amount      string `json:"amount"`
	FloorAmount string `json:"floor_amount"`
}

type DiscountVoucherInfo struct {
	Discount      string `json:"discount"`
	CeilingAmount string `json:"ceiling_amount"`
	FloorAmount   string `json:"floor_amount"`
}

type SpecialVoucherInfo struct {
	SpecialAmount string `json:"special_amount"`
	FloorAmount   string `json:"floor_amount"`
}

type ExchangeVoucherInfo struct {
	Amount      string `json:"amount"`
	FloorAmount string `json:"floor_amount"`
	BizType     string `json:"biz_type"`
}

type VoucherAvailableScopeInfo struct {
	VoucherAvailableAccountInfo        *VoucherAvailableAccountInfo        `json:"voucher_available_account_info,omitempty"`
	VoucherAvailableAppInfo            *VoucherAvailableAppInfo            `json:"voucher_available_app_info,omitempty"`
	VoucherAvailableGeographyScopeInfo *VoucherAvailableGeographyScopeInfo `json:"voucher_available_geography_scope_info,omitempty"`
	VoucherAvailableGoodsInfo          *VoucherAvailableGoodsInfo          `json:"voucher_available_goods_info,omitempty"`
}

type VoucherAvailableAccountInfo struct {
	AvailablePids  []string `json:"available_pids"`
	AvailableSmids []string `json:"available_smids"`
}

type VoucherAvailableAppInfo struct {
	AvailableAppIds []string `json:"available_app_ids"`
}

type VoucherAvailableGeographyScopeInfo struct {
	AvailableGeographyScopeType string                      `json:"available_geography_scope_type"`
	AvailableGeographyShopInfo  *AvailableGeographyShopInfo `json:"available_geography_shop_info"`
}

type AvailableGeographyShopInfo struct {
	AvailableShopIds []string `json:"available_shop_ids"`
}

type VoucherAvailableGoodsInfo struct {
	GoodsName            string   `json:"goods_name"`
	GoodsDescription     string   `json:"goods_description,omitempty"`
	OriginAmount         string   `json:"origin_amount"`
	AvailableGoodsSkuIds []string `json:"available_goods_sku_ids,omitempty"`
	ExcludeGoodsSkuIds   []string `json:"exclude_goods_sku_ids,omitempty"`
}

type VoucherUseRuleInfo struct {
	VoucherMaxUseTimes             int                 `json:"voucher_max_use_times"`
	QuantityLimitPerUser           int                 `json:"quantity_limit_per_user"`
	QuantityLimitPerUserPeriodType string              `json:"quantity_limit_per_user_period_type"`
	VoucherUseTimeInfo             *VoucherUseTimeInfo `json:"voucher_use_time_info"`
}

type VoucherUseTimeInfo struct {
	PeriodType         string              `json:"period_type"`
	AbsolutePeriodInfo *AbsolutePeriodInfo `json:"absolute_period_info"`
	RelativePeriodInfo *RelativePeriodInfo `json:"relative_period_info"`
}

type AbsolutePeriodInfo struct {
	ValidBeginTime   string            `json:"valid_begin_time"`
	ValidEndTime     string            `json:"valid_end_time"`
	TimeRestrictInfo *TimeRestrictInfo `json:"time_restrict_info"`
}

type TimeRestrictInfo struct {
	UsablePeriodInfo  []*UsablePeriodInfoItem  `json:"usable_period_info"`
	DisablePeriodInfo []*DisablePeriodInfoItem `json:"disable_period_info"`
}

type UsablePeriodInfoItem struct {
	RuleType     string        `json:"rule_type"`
	WeekRuleInfo *WeekRuleInfo `json:"week_rule_info"`
}

type WeekRuleInfo struct {
	WeekDay       string         `json:"week_day"`
	TimeRangeInfo *TimeRangeInfo `json:"time_range_info"`
}

type TimeRangeInfo struct {
	BeginTime   string       `json:"begin_time"`
	EndTimeInfo *EndTimeInfo `json:"end_time_info"`
}

type EndTimeInfo struct {
	EndTimeType string `json:"end_time_type"`
	EndTime     string `json:"end_time"`
}

type DisablePeriodInfoItem struct {
	RuleType        string           `json:"rule_type"`
	DateRuleInfo    *DateRuleInfo    `json:"date_rule_info"`
	HolidayRuleInfo *HolidayRuleInfo `json:"holiday_rule_info"`
}

type DateRuleInfo struct {
	DateRangeInfo *DateRangeInfo `json:"date_range_info"`
	TimeRangeInfo *TimeRangeInfo `json:"time_range_info"`
}

type DateRangeInfo struct {
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

type HolidayRuleInfo struct {
	TimeRangeInfo *TimeRangeInfo `json:"time_range_info"`
}

type RelativePeriodInfo struct {
	WaitDaysAfterReceive  int               `json:"wait_days_after_receive"`
	ValidDaysAfterReceive int               `json:"valid_days_after_receive"`
	TimeRestrictInfo      *TimeRestrictInfo `json:"time_restrict_info"`
}

type VoucherDisplayPatternInfo struct {
	BrandName              string   `json:"brand_name"`
	BrandLogoUrl           string   `json:"brand_logo_url"`
	VoucherName            string   `json:"voucher_name,omitempty"`
	VoucherDescription     string   `json:"voucher_description,omitempty"`
	VoucherImage           string   `json:"voucher_image,omitempty"`
	VoucherImageUrl        string   `json:"voucher_image_url,omitempty"`
	VoucherDetailImages    []string `json:"voucher_detail_images,omitempty"`
	VoucherDetailImageUrls []string `json:"voucher_detail_image_urls,omitempty"`
	CustomerServiceMobile  string   `json:"customer_service_mobile,omitempty"`
	CustomerServiceUrl     string   `json:"customer_service_url,omitempty"`
}

type VoucherCustomerGuideInfo struct {
	VoucherUseGuideInfo *VoucherUseGuideInfo `json:"voucher_use_guide_info"`
}

type VoucherUseGuideInfo struct {
	UseGuideMode        []string             `json:"use_guide_mode"`
	MiniAppUseGuideInfo *MiniAppUseGuideInfo `json:"mini_app_use_guide_info"`
}

type MiniAppUseGuideInfo struct {
	MiniAppUrl          string   `json:"mini_app_url"`
	MiniAppServiceCodes []string `json:"mini_app_service_codes"`
}

type VoucherBudgetSupplyInfo struct {
	BudgetType          string               `json:"budget_type"`
	VoucherRechargeInfo *VoucherRechargeInfo `json:"voucher_recharge_info"`
}

type VoucherRechargeInfo struct {
	RechargeType               string                      `json:"recharge_type"`
	VoucherBalanceRechargeInfo *VoucherBalanceRechargeInfo `json:"voucher_balance_recharge_info"`
}

type VoucherBalanceRechargeInfo struct {
	LogonId   string `json:"logon_id"`
	PartnerId string `json:"partner_id"`
	Amount    string `json:"amount"`
}

type VoucherInventoryInfo struct {
	SendCount int `json:"send_count"`
	UseCount  int `json:"use_count"`
}

type MarketingActivityVoucherModify struct {
	ErrorResponse
}

type MarketingActivityVoucherAppend struct {
	ErrorResponse
	RechargeUrl string `json:"recharge_url"`
}

type MarketingActivityVoucherStop struct {
	ErrorResponse
}

type MarketingActivityBatchQuery struct {
	ErrorResponse
	ActivityLiteInfos []*ActivityLiteInfo `json:"activity_lite_infos"`
	PageNum           int                 `json:"page_num"`
	PageSize          int                 `json:"page_size"`
	TotalSize         string              `json:"total_size"`
}

type ActivityLiteInfo struct {
	ActivityBaseInfo          *ActivityBaseInfo          `json:"activity_base_info"`
	VoucherDeductInfo         *VoucherDeductInfo         `json:"voucher_deduct_info"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo `json:"voucher_display_pattern_info"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo `json:"voucher_available_scope_info"`
}

type MarketingActivityConsult struct {
	ErrorResponse
	UserId                string               `json:"user_id"`
	OpenId                string               `json:"open_id"`
	ConsultResultInfoList []*ConsultResultInfo `json:"consult_result_info_list"`
}

type ConsultResultInfo struct {
	ActivityId        string `json:"activity_id"`
	ConsultResultCode string `json:"consult_result_code"`
}

type MarketingActivityQuery struct {
	ErrorResponse
	ActivityBaseInfo          *ActivityBaseInfo          `json:"activity_base_info"`
	VoucherSendModeInfo       *VoucherSendModeInfo       `json:"voucher_send_mode_info"`
	VoucherDeductInfo         *VoucherDeductInfo         `json:"voucher_deduct_info"`
	VoucherUseRuleInfo        *VoucherUseRuleInfo        `json:"voucher_use_rule_info"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo `json:"voucher_display_pattern_info"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo `json:"voucher_available_scope_info"`
	VoucherCustomerGuideInfo  *VoucherCustomerGuideInfo  `json:"voucher_customer_guide_info"`
}

type MarketingActivityQueryMerchantBatchQuery struct {
	ErrorResponse
	ActivityId    string          `json:"activity_id"`
	MerchantInfos []*MerchantInfo `json:"merchant_infos"`
	PageNum       int             `json:"page_num"`
	PageSize      int             `json:"page_size"`
	TotalSize     int             `json:"total_size"`
}

type MerchantInfo struct {
	MerchantId     string `json:"merchant_id"`
	MerchantIdType string `json:"merchant_id_type"`
}

type MarketingActivityQueryAppBatchQuery struct {
	ErrorResponse
	ActivityId string     `json:"activity_id"`
	AppInfos   []*AppInfo `json:"app_infos"`
	PageNum    int        `json:"page_num"`
	PageSize   int        `json:"page_size"`
	TotalSize  int        `json:"total_size"`
}

type AppInfo struct {
	MiniAppId string `json:"mini_app_id"`
}

type MarketingActivityQueryShopBatchQuery struct {
	ErrorResponse
	ActivityId string      `json:"activity_id"`
	ShopInfos  []*ShopInfo `json:"shop_infos"`
	PageNum    int         `json:"page_num"`
	PageSize   int         `json:"page_size"`
	TotalSize  int         `json:"total_size"`
}

type ShopInfo struct {
	ShopId    string `json:"shop_id"`
	ShopType  string `json:"shop_type"`
	ShopName  string `json:"shop_name"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type MarketingActivityQueryGoodsBatchQuery struct {
	ErrorResponse
	ActivityId   string         `json:"activity_id"`
	AppItemInfos []*AppItemInfo `json:"app_item_infos"`
	GoodsInfos   []*GoodsInfo   `json:"goods_infos"`
	PageNum      int            `json:"page_num"`
	PageSize     int            `json:"page_size"`
	TotalSize    int            `json:"total_size"`
}

type AppItemInfo struct {
	ItemId      string `json:"item_id"`
	ItemUseType string `json:"item_use_type"`
	OutItemId   string `json:"out_item_id"`
	MiniAppId   string `json:"mini_app_id"`
}

type GoodsInfo struct {
	GoodsId      string `json:"goods_id"`
	GoodsUseType string `json:"goods_use_type"`
}

type MarketingActivityQueryUserBatchQueryVoucher struct {
	ErrorResponse
	UserVoucherInfos []*UserVoucherInfo `json:"user_voucher_infos"`
	PageNum          int                `json:"page_num"`
	PageSize         int                `json:"page_size"`
	TotalSize        int                `json:"total_size"`
}

type UserVoucherInfo struct {
	UserVoucherBaseInfo       *UserVoucherBaseInfo       `json:"user_voucher_base_info"`
	ActivityBaseInfo          *ActivityBaseInfo          `json:"activity_base_info"`
	VoucherSendModeInfo       *VoucherSendModeInfo       `json:"voucher_send_mode_info"`
	VoucherDeductInfo         *VoucherDeductInfo         `json:"voucher_deduct_info"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo `json:"voucher_display_pattern_info"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo `json:"voucher_available_scope_info"`
	VoucherCustomerGuideInfo  *VoucherCustomerGuideInfo  `json:"voucher_customer_guide_info"`
}

type UserVoucherBaseInfo struct {
	VoucherId          string `json:"voucher_id"`
	VoucherCode        string `json:"voucher_code"`
	VoucherName        string `json:"voucher_name"`
	VoucherStatus      string `json:"voucher_status"`
	CreateTime         string `json:"create_time"`
	ValidBeginTime     string `json:"valid_begin_time"`
	ValidEndTime       string `json:"valid_end_time"`
	AssociateTradeNo   string `json:"associate_trade_no"`
	VoucherMaxUseTimes int    `json:"voucher_max_use_times"`
	VoucherUsedTimes   int    `json:"voucher_used_times"`
	BelongMerchantId   string `json:"belong_merchant_id"`
}

type MarketingActivityQueryUserQueryVoucher struct {
	ErrorResponse
	UserVoucherBaseInfo       *UserVoucherBaseInfo       `json:"user_voucher_base_info"`
	ActivityBaseInfo          *ActivityBaseInfo          `json:"activity_base_info"`
	VoucherSendModeInfo       *VoucherSendModeInfo       `json:"voucher_send_mode_info"`
	VoucherUseRuleInfo        *VoucherUseRuleInfo        `json:"voucher_use_rule_info"`
	VoucherDeductInfo         *VoucherDeductInfo         `json:"voucher_deduct_info"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo `json:"voucher_display_pattern_info"`
	VoucherCustomerGuideInfo  *VoucherCustomerGuideInfo  `json:"voucher_customer_guide_info"`
}

type MarketingCampaignOrderVoucherConsult struct {
	ErrorResponse
	OptimalTotalPromoAmount string            `json:"optimal_total_promo_amount"`
	VoucherConsultList      []*VoucherConsult `json:"voucher_consult_list"`
}

type VoucherConsult struct {
	VoucherId       string         `json:"voucher_id"`
	VoucherName     string         `json:"voucher_name"`
	VoucherType     string         `json:"voucher_type"`
	PromoType       string         `json:"promo_type"`
	ReductionAmount string         `json:"reduction_amount"`
	SpecifiedAmount string         `json:"specified_amount"`
	ReductionRatio  string         `json:"reduction_ratio"`
	CeilingAmount   string         `json:"ceiling_amount"`
	ThresholdAmount string         `json:"threshold_amount"`
	PromoAmount     string         `json:"promo_amount"`
	PromoText       string         `json:"promo_text"`
	ItemPromoInfo   *ItemPromoInfo `json:"item_promo_info"`
}

type ItemPromoInfo struct {
	ItemName          string         `json:"item_name"`
	ItemDesc          string         `json:"item_desc"`
	ItemCoverPic      string         `json:"item_cover_pic"`
	ItemDetailPicList []string       `json:"item_detail_pic_list"`
	ItemConsultList   []*ItemConsult `json:"item_consult_list"`
}

type ItemConsult struct {
	ItemId      string `json:"item_id"`
	PromoAmount string `json:"promo_amount"`
	PromoCount  string `json:"promo_count"`
}

type MarketingActivityOrderVoucherCreate struct {
	ErrorResponse
	ActivityId                      string                           `json:"activity_id"`
	VoucherAvailableScopeResultInfo *VoucherAvailableScopeResultInfo `json:"voucher_available_scope_result_info"`
}

type VoucherAvailableScopeResultInfo struct {
	VoucherAvailableGeographyScopeResultInfo *VoucherAvailableGeographyScopeResultInfo `json:"voucher_available_geography_scope_result_info"`
}

type VoucherAvailableGeographyScopeResultInfo struct {
	AvailableGeographyShopResultInfo *AvailableGeographyShopResultInfo `json:"available_geography_shop_result_info"`
}

type AvailableGeographyShopResultInfo struct {
	SuccessAvailableShopIds             []string                             `json:"success_available_shop_ids"`
	FailAvailableShopInfos              []*FailAvailableShopInfo             `json:"fail_available_shop_infos"`
	AvailableGeographyAllShopResultInfo *AvailableGeographyAllShopResultInfo `json:"available_geography_all_shop_result_info"`
}

type FailAvailableShopInfo struct {
	ShopId      string   `json:"shop_id"`
	FailReasons []string `json:"fail_reasons"`
	FailMessage string   `json:"fail_message"`
}

type AvailableGeographyAllShopResultInfo struct {
	SuccessExcludeShopIds []string               `json:"success_exclude_shop_ids"`
	FailExcludeShopInfos  []*FailExcludeShopInfo `json:"fail_exclude_shop_infos"`
}

type FailExcludeShopInfo struct {
	ShopId      string   `json:"shop_id"`
	RealShopId  string   `json:"real_shop_id"`
	FailReasons []string `json:"fail_reasons"`
	FailMessage string   `json:"fail_message"`
}

type MarketingActivityOrderVoucherCodeDeposit struct {
	ErrorResponse
	SuccessCount              int                      `json:"success_count"`
	FailCount                 int                      `json:"fail_count"`
	SuccessVoucherCodeList    []string                 `json:"success_voucher_code_list"`
	FailVoucherCodeDetailList []*FailVoucherCodeDetail `json:"fail_voucher_code_detail_list"`
}

type FailVoucherCodeDetail struct {
	VoucherCode string `json:"voucher_code"`
	ErrorCode   string `json:"error_code"`
	ErrorMsg    string `json:"error_msg"`
}

type MarketingActivityOrderVoucherModify struct {
	ErrorResponse
	VoucherAvailableScopeResultInfo *VoucherAvailableScopeResultInfo `json:"voucher_available_scope_result_info"`
}

type MarketingActivityOrderVoucherStop struct {
	ErrorResponse
}

type MarketingActivityOrderVoucherAppend struct {
	ErrorResponse
}

type MarketingActivityOrderVoucherUse struct {
	ErrorResponse
	ActivityId                 string                      `json:"activity_id"`
	VoucherUseDetailResultInfo *VoucherUseDetailResultInfo `json:"voucher_use_detail_result_info"`
}

type VoucherUseDetailResultInfo struct {
	VoucherMaxUnUseTimes int `json:"voucher_max_un_use_times"`
}

type MarketingActivityOrderVoucherRefund struct {
	ErrorResponse
	ActivityId                 string                      `json:"activity_id"`
	VoucherUseDetailResultInfo *VoucherUseDetailResultInfo `json:"voucher_use_detail_result_info"`
}

type MarketingActivityOrderVoucherQuery struct {
	ErrorResponse
	ActivityBaseInfo          *ActivityBaseInfo          `json:"activity_base_info"`
	VoucherSendModeInfo       *VoucherSendModeInfo       `json:"voucher_send_mode_info"`
	VoucherDeductInfo         *VoucherDeductInfo         `json:"voucher_deduct_info"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo `json:"voucher_available_scope_info"`
	VoucherUseRuleInfo        *VoucherUseRuleInfo        `json:"voucher_use_rule_info"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo `json:"voucher_display_pattern_info"`
	VoucherCustomerGuideInfo  *VoucherCustomerGuideInfo  `json:"voucher_customer_guide_info"`
	VoucherInventoryInfo      *VoucherInventoryInfo      `json:"voucher_inventory_info"`
}

type MarketingActivityOrderVoucherCodeCount struct {
	ErrorResponse
	SuccessCount int `json:"success_count"`
}

type MarketingCardTemplateCreate struct {
	ErrorResponse
	TemplateId string `json:"template_id"`
}

type MarketingCardTemplateModify struct {
	ErrorResponse
	TemplateId string `json:"template_id"`
}

type MarketingCardTemplateQuery struct {
	ErrorResponse
	TemplateStyleInfo  *TemplateStyleInfo  `json:"template_style_info"`
	AccessVersion      string              `json:"access_version"`
	CardLevelConfs     []*CardLevelConf    `json:"card_level_confs"`
	TemplateFormConfig *TemplateFormConfig `json:"template_form_config"`
	SpiAppId           string              `json:"spi_app_id"`
}

type TemplateStyleInfo struct {
	CardShowName string `json:"card_show_name"`
	LogoId       string `json:"logo_id"`
	BackgroundId string `json:"background_id"`
	BrandName    string `json:"brand_name"`
}

type CardLevelConf struct {
	Level         string `json:"level"`
	LevelShowName string `json:"level_show_name"`
	LevelIcon     string `json:"level_icon"`
	LevelDesc     string `json:"level_desc"`
}

type TemplateFormConfig struct {
	Fields            *Fields `json:"fields"`
	OpenCardMiniAppId string  `json:"open_card_mini_app_id"`
}

type Fields struct {
	Required []string `json:"required"`
	Optional []string `json:"optional"`
}

type MarketingCardUpdate struct {
	ErrorResponse
	ResultCode string `json:"result_code"`
}

type MarketingCardQuery struct {
	ErrorResponse
	CardInfo          *CardInfo          `json:"card_info"`
	SchemaUrl         string             `json:"schema_url"`
	PassId            string             `json:"pass_id"`
	PaidOuterCardInfo *PaidOuterCardInfo `json:"paid_outer_card_info"`
}

type CardInfo struct {
	BizCardNo      string       `json:"biz_card_no"`
	ExternalCardNo string       `json:"external_card_no"`
	OpenDate       string       `json:"open_date"`
	ValidDate      string       `json:"valid_date"`
	Level          string       `json:"level"`
	Point          string       `json:"point"`
	Balance        string       `json:"balance"`
	TemplateId     string       `json:"template_id"`
	MdcodeInfo     *MdcodeInfo  `json:"mdcode_info"`
	FrontTextList  []*FrontText `json:"front_text_list"`
	FrontImageId   string       `json:"front_image_id"`
}

type MdcodeInfo struct {
	CodeStatus string `json:"code_status"`
	CodeValue  string `json:"code_value"`
	ExpireTime string `json:"expire_time"`
	TimeStamp  int    `json:"time_stamp"`
}

type FrontText struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type PaidOuterCardInfo struct {
	Action       string        `json:"action"`
	PurchaseInfo *PurchaseInfo `json:"purchase_info"`
	CycleInfo    *CycleInfo    `json:"cycle_info"`
}

type PurchaseInfo struct {
	Source        string `json:"source"`
	Price         string `json:"price"`
	ActionDate    string `json:"action_date"`
	OutTradeNo    string `json:"out_trade_no"`
	AlipayTradeNo string `json:"alipay_trade_no"`
}

type CycleInfo struct {
	OpenStatus              string `json:"open_status"`
	CloseReason             string `json:"close_reason"`
	CycleType               string `json:"cycle_type"`
	AlipayDeductScene       string `json:"alipay_deduct_scene"`
	AlipayDeductProductCode string `json:"alipay_deduct_product_code"`
	AlipayDeductAgreement   string `json:"alipay_deduct_agreement"`
}

type MarketingCardDelete struct {
	ErrorResponse
	BizSerialNo string `json:"biz_serial_no"`
}

type MarketingCardMessageNotify struct {
	ErrorResponse
	ResultCode string `json:"result_code"`
}

type MarketingCardFormTemplateSet struct {
	ErrorResponse
}

type OfflineMaterialImageUpload struct {
	ErrorResponse
	ImageId  string `json:"image_id"`
	ImageUrl string `json:"image_url"`
}
