package alipay

type MarketingActivityOrderVoucherCreateRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ActivityId                      string `json:"activity_id"`
	VoucherAvailableScopeResultInfo struct {
		VoucherAvailableGeographyScopeResultInfo struct {
			AvailableGeographyShopResultInfo struct {
				SuccessAvailableShopIds []string `json:"success_available_shop_ids"`
				FailAvailableShopInfos  []struct {
					ShopId      string   `json:"shop_id"`
					FailReasons []string `json:"fail_reasons"`
					FailMessage string   `json:"fail_message"`
				} `json:"fail_available_shop_infos"`
				AvailableGeographyAllShopResultInfo struct {
					SuccessExcludeShopIds []string `json:"success_exclude_shop_ids"`
					FailExcludeShopInfos  []struct {
						ShopId      string   `json:"shop_id"`
						RealShopId  string   `json:"real_shop_id"`
						FailReasons []string `json:"fail_reasons"`
						FailMessage string   `json:"fail_message"`
					} `json:"fail_exclude_shop_infos"`
				} `json:"available_geography_all_shop_result_info"`
			} `json:"available_geography_shop_result_info"`
		} `json:"voucher_available_geography_scope_result_info"`
	} `json:"voucher_available_scope_result_info"`
}

type MarketingActivityOrderVoucherCodeDepositRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	SuccessCount              int      `json:"success_count"`
	FailCount                 int      `json:"fail_count"`
	SuccessVoucherCodeList    []string `json:"success_voucher_code_list"`
	FailVoucherCodeDetailList []struct {
		VoucherCode string `json:"voucher_code"`
		ErrorCode   string `json:"error_code"`
		ErrorMsg    string `json:"error_msg"`
	} `json:"fail_voucher_code_detail_list"`
}

type MarketingActivityOrderVoucherModifyRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	VoucherAvailableScopeResultInfo struct {
		VoucherAvailableGeographyScopeResultInfo struct {
			AvailableGeographyShopResultInfo struct {
				SuccessAvailableShopIds []string `json:"success_available_shop_ids"`
				FailAvailableShopInfos  []struct {
					ShopId      string   `json:"shop_id"`
					FailReasons []string `json:"fail_reasons"`
					FailMessage string   `json:"fail_message"`
				} `json:"fail_available_shop_infos"`
				AvailableGeographyAllShopResultInfo struct {
					SuccessExcludeShopIds []string `json:"success_exclude_shop_ids"`
					FailExcludeShopInfos  []struct {
						ShopId      string   `json:"shop_id"`
						FailReasons []string `json:"fail_reasons"`
						FailMessage string   `json:"fail_message"`
					} `json:"fail_exclude_shop_infos"`
				} `json:"available_geography_all_shop_result_info"`
			} `json:"available_geography_shop_result_info"`
		} `json:"voucher_available_geography_scope_result_info"`
	} `json:"voucher_available_scope_result_info"`
}

type MarketingActivityOrderVoucherStopRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

type MarketingActivityOrderVoucherAppendRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`
}

type MarketingActivityOrderVoucherUseRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ActivityId                 string `json:"activity_id"`
	VoucherUseDetailResultInfo struct {
		VoucherMaxUnUseTimes int `json:"voucher_max_un_use_times"`
	} `json:"voucher_use_detail_result_info"`
}

type MarketingActivityOrderVoucherRefundRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ActivityId                 string `json:"activity_id"`
	VoucherUseDetailResultInfo struct {
		VoucherMaxUnUseTimes int `json:"voucher_max_un_use_times"`
	} `json:"voucher_use_detail_result_info"`
}

type MarketingActivityConsultRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	UserId                string `json:"user_id"`
	OpenId                string `json:"open_id"`
	ConsultResultInfoList []struct {
		ActivityId        string `json:"activity_id"`
		ConsultResultCode string `json:"consult_result_code"`
	} `json:"consult_result_info_list"`
}

type MarketingActivityOrderVoucherQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ActivityBaseInfo struct {
		ActivityId         string `json:"activity_id"`
		ActivityName       string `json:"activity_name"`
		BelongMerchantInfo struct {
			MerchantId string `json:"merchant_id"`
		} `json:"belong_merchant_info"`
		CodeMode                string `json:"code_mode"`
		ActivityOperationStatus string `json:"activity_operation_status"`
		ActivityStatus          string `json:"activity_status"`
	} `json:"activity_base_info"`
	VoucherSendModeInfo struct {
		VoucherSendMode     string `json:"voucher_send_mode"`
		VoucherSendRuleInfo struct {
			Quantity             int    `json:"quantity"`
			QuantityLimitPerUser int    `json:"quantity_limit_per_user"`
			NaturalPersonLimit   bool   `json:"natural_person_limit"`
			PhoneNumberLimit     bool   `json:"phone_number_limit"`
			PublishStartTime     string `json:"publish_start_time"`
			PublishEndTime       string `json:"publish_end_time"`
		} `json:"voucher_send_rule_info"`
	} `json:"voucher_send_mode_info"`
	VoucherDeductInfo struct {
		VoucherType    string `json:"voucher_type"`
		FixVoucherInfo struct {
			Amount      string `json:"amount"`
			FloorAmount string `json:"floor_amount"`
		} `json:"fix_voucher_info"`
		DiscountVoucherInfo struct {
			Discount      string `json:"discount"`
			CeilingAmount string `json:"ceiling_amount"`
			FloorAmount   string `json:"floor_amount"`
		} `json:"discount_voucher_info"`
		SpecialVoucherInfo struct {
			SpecialAmount string `json:"special_amount"`
			FloorAmount   string `json:"floor_amount"`
		} `json:"special_voucher_info"`
	} `json:"voucher_deduct_info"`
	VoucherAvailableScopeInfo struct {
		VoucherAvailableGeographyScopeInfo struct {
			AvailableGeographyScopeType string `json:"available_geography_scope_type"`
			AvailableGeographyShopInfo  struct {
				AvailableShopIds          []string `json:"available_shop_ids"`
				AvailableGeographyAllShop struct {
					MerchantIds      []string `json:"merchant_ids"`
					ExcludeShopIds   []string `json:"exclude_shop_ids"`
					AvailableBrandId string   `json:"available_brand_id"`
				} `json:"available_geography_all_shop"`
			} `json:"available_geography_shop_info"`
			AvailableGeographyCityInfo struct {
				AllCity            bool     `json:"all_city"`
				AvailableCityCodes []string `json:"available_city_codes"`
			} `json:"available_geography_city_info"`
		} `json:"voucher_available_geography_scope_info"`
		VoucherAvailableGoodsInfo struct {
			GoodsName    string `json:"goods_name"`
			OriginAmount string `json:"origin_amount"`
		} `json:"voucher_available_goods_info"`
	} `json:"voucher_available_scope_info"`
	VoucherUseRuleInfo struct {
		VoucherUseTimeInfo struct {
			PeriodType         string `json:"period_type"`
			AbsolutePeriodInfo struct {
				ValidBeginTime   string `json:"valid_begin_time"`
				ValidEndTime     string `json:"valid_end_time"`
				TimeRestrictInfo struct {
					UsablePeriodInfo []struct {
						RuleType     string `json:"rule_type"`
						WeekRuleInfo struct {
							WeekDay       string `json:"week_day"`
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"week_rule_info"`
					} `json:"usable_period_info"`
					DisablePeriodInfo []struct {
						RuleType     string `json:"rule_type"`
						DateRuleInfo struct {
							DateRangeInfo struct {
								BeginDate string `json:"begin_date"`
								EndDate   string `json:"end_date"`
							} `json:"date_range_info"`
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"date_rule_info"`
						HolidayRuleInfo struct {
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"holiday_rule_info"`
					} `json:"disable_period_info"`
				} `json:"time_restrict_info"`
			} `json:"absolute_period_info"`
			RelativePeriodInfo struct {
				WaitDaysAfterReceive  int `json:"wait_days_after_receive"`
				ValidDaysAfterReceive int `json:"valid_days_after_receive"`
				TimeRestrictInfo      struct {
					UsablePeriodInfo []struct {
						RuleType     string `json:"rule_type"`
						WeekRuleInfo struct {
							WeekDay       string `json:"week_day"`
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"week_rule_info"`
					} `json:"usable_period_info"`
					DisablePeriodInfo []struct {
						RuleType     string `json:"rule_type"`
						DateRuleInfo struct {
							DateRangeInfo struct {
								BeginDate string `json:"begin_date"`
								EndDate   string `json:"end_date"`
							} `json:"date_range_info"`
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"date_rule_info"`
						HolidayRuleInfo struct {
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"holiday_rule_info"`
					} `json:"disable_period_info"`
				} `json:"time_restrict_info"`
			} `json:"relative_period_info"`
		} `json:"voucher_use_time_info"`
	} `json:"voucher_use_rule_info"`
	VoucherDisplayPatternInfo struct {
		BrandName              string   `json:"brand_name"`
		BrandLogo              string   `json:"brand_logo"`
		BrandLogoUrl           string   `json:"brand_logo_url"`
		VoucherDescription     string   `json:"voucher_description"`
		VoucherImage           string   `json:"voucher_image"`
		VoucherImageUrl        string   `json:"voucher_image_url"`
		VoucherDetailImages    []string `json:"voucher_detail_images"`
		VoucherDetailImageUrls []string `json:"voucher_detail_image_urls"`
		CustomerServiceMobile  string   `json:"customer_service_mobile"`
		CustomerServiceUrl     string   `json:"customer_service_url"`
	} `json:"voucher_display_pattern_info"`
	VoucherCustomerGuideInfo struct {
		VoucherUseGuideInfo struct {
			UseGuideMode        []string `json:"use_guide_mode"`
			MiniAppUseGuideInfo struct {
				MiniAppUrl          string   `json:"mini_app_url"`
				MiniAppServiceCodes []string `json:"mini_app_service_codes"`
			} `json:"mini_app_use_guide_info"`
		} `json:"voucher_use_guide_info"`
	} `json:"voucher_customer_guide_info"`
	VoucherInventoryInfo struct {
		SendCount int `json:"send_count"`
		UseCount  int `json:"use_count"`
	} `json:"voucher_inventory_info"`
}

type MarketingActivityQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ActivityBaseInfo struct {
		ActivityId         string `json:"activity_id"`
		BelongMerchantInfo struct {
			MerchantId string `json:"merchant_id"`
		} `json:"belong_merchant_info"`
		ActivityStatus string `json:"activity_status"`
	} `json:"activity_base_info"`
	VoucherSendModeInfo struct {
		VoucherSendMode     string `json:"voucher_send_mode"`
		VoucherSaleModeInfo struct {
			SaleAmount        string `json:"sale_amount"`
			Refundable        bool   `json:"refundable"`
			OverdueRefundable bool   `json:"overdue_refundable"`
		} `json:"voucher_sale_mode_info"`
		VoucherPackageModeInfo struct {
			VoucherPackageId string `json:"voucher_package_id"`
		} `json:"voucher_package_mode_info"`
		VoucherSendRuleInfo struct {
			Quantity                       int    `json:"quantity"`
			MaxQuantityByDay               int    `json:"max_quantity_by_day"`
			QuantityLimitPerUser           int    `json:"quantity_limit_per_user"`
			QuantityLimitPerUserPeriodType string `json:"quantity_limit_per_user_period_type"`
			NaturalPersonLimit             bool   `json:"natural_person_limit"`
			PhoneNumberLimit               bool   `json:"phone_number_limit"`
			RealNameLimit                  bool   `json:"real_name_limit"`
			PublishStartTime               string `json:"publish_start_time"`
			PublishEndTime                 string `json:"publish_end_time"`
		} `json:"voucher_send_rule_info"`
	} `json:"voucher_send_mode_info"`
	VoucherDeductInfo struct {
		VoucherType    string `json:"voucher_type"`
		FixVoucherInfo struct {
			Amount      string `json:"amount"`
			FloorAmount string `json:"floor_amount"`
		} `json:"fix_voucher_info"`
		DiscountVoucherInfo struct {
			Discount      string `json:"discount"`
			CeilingAmount string `json:"ceiling_amount"`
			FloorAmount   string `json:"floor_amount"`
		} `json:"discount_voucher_info"`
		SpecialVoucherInfo struct {
			SpecialAmount string `json:"special_amount"`
			FloorAmount   string `json:"floor_amount"`
		} `json:"special_voucher_info"`
		ExchangeVoucherInfo struct {
			Amount      string `json:"amount"`
			FloorAmount string `json:"floor_amount"`
			BizType     string `json:"biz_type"`
		} `json:"exchange_voucher_info"`
	} `json:"voucher_deduct_info"`
	VoucherUseRuleInfo struct {
		VoucherMaxUseTimes int `json:"voucher_max_use_times"`
		VoucherUseTimeInfo struct {
			PeriodType         string `json:"period_type"`
			AbsolutePeriodInfo struct {
				ValidBeginTime   string `json:"valid_begin_time"`
				ValidEndTime     string `json:"valid_end_time"`
				TimeRestrictInfo struct {
					UsablePeriodInfo []struct {
						RuleType     string `json:"rule_type"`
						WeekRuleInfo struct {
							WeekDay       string `json:"week_day"`
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"week_rule_info"`
					} `json:"usable_period_info"`
					DisablePeriodInfo []struct {
						RuleType     string `json:"rule_type"`
						DateRuleInfo struct {
							DateRangeInfo struct {
								BeginDate string `json:"begin_date"`
								EndDate   string `json:"end_date"`
							} `json:"date_range_info"`
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"date_rule_info"`
						HolidayRuleInfo struct {
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"holiday_rule_info"`
					} `json:"disable_period_info"`
				} `json:"time_restrict_info"`
			} `json:"absolute_period_info"`
			RelativePeriodInfo struct {
				WaitDaysAfterReceive  int `json:"wait_days_after_receive"`
				ValidDaysAfterReceive int `json:"valid_days_after_receive"`
				TimeRestrictInfo      struct {
					UsablePeriodInfo []struct {
						RuleType     string `json:"rule_type"`
						WeekRuleInfo struct {
							WeekDay       string `json:"week_day"`
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"week_rule_info"`
					} `json:"usable_period_info"`
					DisablePeriodInfo []struct {
						RuleType     string `json:"rule_type"`
						DateRuleInfo struct {
							DateRangeInfo struct {
								BeginDate string `json:"begin_date"`
								EndDate   string `json:"end_date"`
							} `json:"date_range_info"`
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"date_rule_info"`
						HolidayRuleInfo struct {
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"holiday_rule_info"`
					} `json:"disable_period_info"`
				} `json:"time_restrict_info"`
			} `json:"relative_period_info"`
		} `json:"voucher_use_time_info"`
	} `json:"voucher_use_rule_info"`
	VoucherDisplayPatternInfo struct {
		BrandName              string   `json:"brand_name"`
		BrandLogoUrl           string   `json:"brand_logo_url"`
		VoucherName            string   `json:"voucher_name"`
		VoucherDescription     string   `json:"voucher_description"`
		VoucherImageUrl        string   `json:"voucher_image_url"`
		VoucherDetailImageUrls []string `json:"voucher_detail_image_urls"`
		CustomerServiceMobile  string   `json:"customer_service_mobile"`
		CustomerServiceUrl     string   `json:"customer_service_url"`
	} `json:"voucher_display_pattern_info"`
	VoucherAvailableScopeInfo struct {
		VoucherAvailableGoodsInfo struct {
			GoodsName        string `json:"goods_name"`
			GoodsDescription string `json:"goods_description"`
			OriginAmount     string `json:"origin_amount"`
		} `json:"voucher_available_goods_info"`
	} `json:"voucher_available_scope_info"`
	VoucherCustomerGuideInfo struct {
		VoucherUseGuideInfo struct {
			UseGuideMode        []string `json:"use_guide_mode"`
			MiniAppUseGuideInfo struct {
				MiniAppUrl          string   `json:"mini_app_url"`
				MiniAppServiceCodes []string `json:"mini_app_service_codes"`
			} `json:"mini_app_use_guide_info"`
		} `json:"voucher_use_guide_info"`
	} `json:"voucher_customer_guide_info"`
}

type MarketingActivityOrderVoucherCodeCountRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	SuccessCount int `json:"success_count"`
}

type MarketingActivityBatchQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ActivityLiteInfos []struct {
		ActivityBaseInfo struct {
			ActivityId         string `json:"activity_id"`
			BelongMerchantInfo struct {
				MerchantId string `json:"merchant_id"`
			} `json:"belong_merchant_info"`
			ActivityStatus string `json:"activity_status"`
		} `json:"activity_base_info"`
		VoucherDeductInfo struct {
			VoucherType    string `json:"voucher_type"`
			FixVoucherInfo struct {
				Amount      string `json:"amount"`
				FloorAmount string `json:"floor_amount"`
			} `json:"fix_voucher_info"`
			DiscountVoucherInfo struct {
				Discount      string `json:"discount"`
				CeilingAmount string `json:"ceiling_amount"`
				FloorAmount   string `json:"floor_amount"`
			} `json:"discount_voucher_info"`
			SpecialVoucherInfo struct {
				SpecialAmount string `json:"special_amount"`
				FloorAmount   string `json:"floor_amount"`
			} `json:"special_voucher_info"`
			ExchangeVoucherInfo struct {
				Amount      string `json:"amount"`
				FloorAmount string `json:"floor_amount"`
				BizType     string `json:"biz_type"`
			} `json:"exchange_voucher_info"`
		} `json:"voucher_deduct_info"`
		VoucherDisplayPatternInfo struct {
			BrandName    string `json:"brand_name"`
			BrandLogoUrl string `json:"brand_logo_url"`
		} `json:"voucher_display_pattern_info"`
		VoucherAvailableScopeInfo struct {
			VoucherAvailableGoodsInfo struct {
				GoodsName    string `json:"goods_name"`
				OriginAmount string `json:"origin_amount"`
			} `json:"voucher_available_goods_info"`
		} `json:"voucher_available_scope_info"`
	} `json:"activity_lite_infos"`
	PageNum   int    `json:"page_num"`
	PageSize  int    `json:"page_size"`
	TotalSize string `json:"total_size"`
}

type MarketingActivityQueryUserBatchQueryVoucherRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	UserVoucherInfos []struct {
		UserVoucherBaseInfo struct {
			VoucherId        string `json:"voucher_id"`
			VoucherCode      string `json:"voucher_code"`
			VoucherName      string `json:"voucher_name"`
			VoucherStatus    string `json:"voucher_status"`
			CreateTime       string `json:"create_time"`
			ValidBeginTime   string `json:"valid_begin_time"`
			ValidEndTime     string `json:"valid_end_time"`
			BelongMerchantId string `json:"belong_merchant_id"`
		} `json:"user_voucher_base_info"`
		ActivityBaseInfo struct {
			ActivityId         string `json:"activity_id"`
			BelongMerchantInfo struct {
				MerchantId string `json:"merchant_id"`
			} `json:"belong_merchant_info"`
			ActivityProductType string `json:"activity_product_type"`
		} `json:"activity_base_info"`
		VoucherSendModeInfo struct {
			VoucherSendMode string `json:"voucher_send_mode"`
		} `json:"voucher_send_mode_info"`
		VoucherDeductInfo struct {
			VoucherType    string `json:"voucher_type"`
			FixVoucherInfo struct {
				Amount      string `json:"amount"`
				FloorAmount string `json:"floor_amount"`
			} `json:"fix_voucher_info"`
			DiscountVoucherInfo struct {
				Discount      string `json:"discount"`
				CeilingAmount string `json:"ceiling_amount"`
				FloorAmount   string `json:"floor_amount"`
			} `json:"discount_voucher_info"`
			SpecialVoucherInfo struct {
				SpecialAmount string `json:"special_amount"`
				FloorAmount   string `json:"floor_amount"`
			} `json:"special_voucher_info"`
			ExchangeVoucherInfo struct {
				Amount      string `json:"amount"`
				FloorAmount string `json:"floor_amount"`
				BizType     string `json:"biz_type"`
			} `json:"exchange_voucher_info"`
		} `json:"voucher_deduct_info"`
		VoucherDisplayPatternInfo struct {
			BrandName    string `json:"brand_name"`
			BrandLogoUrl string `json:"brand_logo_url"`
		} `json:"voucher_display_pattern_info"`
		VoucherAvailableScopeInfo struct {
			VoucherAvailableGoodsInfo struct {
				GoodsName    string `json:"goods_name"`
				OriginAmount string `json:"origin_amount"`
			} `json:"voucher_available_goods_info"`
		} `json:"voucher_available_scope_info"`
		VoucherCustomerGuideInfo struct {
			VoucherUseGuideInfo struct {
				UseGuideMode        []string `json:"use_guide_mode"`
				MiniAppUseGuideInfo struct {
					MiniAppUrl          string   `json:"mini_app_url"`
					MiniAppServiceCodes []string `json:"mini_app_service_codes"`
				} `json:"mini_app_use_guide_info"`
			} `json:"voucher_use_guide_info"`
		} `json:"voucher_customer_guide_info"`
	} `json:"user_voucher_infos"`
	PageNum   int `json:"page_num"`
	PageSize  int `json:"page_size"`
	TotalSize int `json:"total_size"`
}

type MarketingActivityQueryUserQueryVoucherRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	UserVoucherBaseInfo struct {
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
	} `json:"user_voucher_base_info"`
	ActivityBaseInfo struct {
		ActivityId         string `json:"activity_id"`
		BelongMerchantInfo struct {
			MerchantId string `json:"merchant_id"`
		} `json:"belong_merchant_info"`
		ActivityStatus string `json:"activity_status"`
	} `json:"activity_base_info"`
	VoucherSendModeInfo struct {
		VoucherSendMode     string `json:"voucher_send_mode"`
		VoucherSaleModeInfo struct {
			SaleAmount        string `json:"sale_amount"`
			Refundable        bool   `json:"refundable"`
			OverdueRefundable bool   `json:"overdue_refundable"`
		} `json:"voucher_sale_mode_info"`
		VoucherPackageModeInfo struct {
			VoucherPackageId string `json:"voucher_package_id"`
		} `json:"voucher_package_mode_info"`
		VoucherSendRuleInfo struct {
			Quantity                       int    `json:"quantity"`
			MaxQuantityByDay               int    `json:"max_quantity_by_day"`
			QuantityLimitPerUser           int    `json:"quantity_limit_per_user"`
			QuantityLimitPerUserPeriodType string `json:"quantity_limit_per_user_period_type"`
			NaturalPersonLimit             bool   `json:"natural_person_limit"`
			PhoneNumberLimit               bool   `json:"phone_number_limit"`
			RealNameLimit                  bool   `json:"real_name_limit"`
			PublishStartTime               string `json:"publish_start_time"`
			PublishEndTime                 string `json:"publish_end_time"`
		} `json:"voucher_send_rule_info"`
	} `json:"voucher_send_mode_info"`
	VoucherUseRuleInfo struct {
		QuantityLimitPerUser           int    `json:"quantity_limit_per_user"`
		QuantityLimitPerUserPeriodType string `json:"quantity_limit_per_user_period_type"`
		VoucherMaxUseTimes             int    `json:"voucher_max_use_times"`
		VoucherUseTimeInfo             struct {
			PeriodType         string `json:"period_type"`
			AbsolutePeriodInfo struct {
				ValidBeginTime   string `json:"valid_begin_time"`
				ValidEndTime     string `json:"valid_end_time"`
				TimeRestrictInfo struct {
					UsablePeriodInfo []struct {
						RuleType     string `json:"rule_type"`
						WeekRuleInfo struct {
							WeekDay       string `json:"week_day"`
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"week_rule_info"`
					} `json:"usable_period_info"`
					DisablePeriodInfo []struct {
						RuleType     string `json:"rule_type"`
						DateRuleInfo struct {
							DateRangeInfo struct {
								BeginDate string `json:"begin_date"`
								EndDate   string `json:"end_date"`
							} `json:"date_range_info"`
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"date_rule_info"`
						HolidayRuleInfo struct {
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"holiday_rule_info"`
					} `json:"disable_period_info"`
				} `json:"time_restrict_info"`
			} `json:"absolute_period_info"`
			RelativePeriodInfo struct {
				WaitDaysAfterReceive  int `json:"wait_days_after_receive"`
				ValidDaysAfterReceive int `json:"valid_days_after_receive"`
				TimeRestrictInfo      struct {
					UsablePeriodInfo []struct {
						RuleType     string `json:"rule_type"`
						WeekRuleInfo struct {
							WeekDay       string `json:"week_day"`
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"week_rule_info"`
					} `json:"usable_period_info"`
					DisablePeriodInfo []struct {
						RuleType     string `json:"rule_type"`
						DateRuleInfo struct {
							DateRangeInfo struct {
								BeginDate string `json:"begin_date"`
								EndDate   string `json:"end_date"`
							} `json:"date_range_info"`
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"date_rule_info"`
						HolidayRuleInfo struct {
							TimeRangeInfo struct {
								BeginTime   string `json:"begin_time"`
								EndTimeInfo struct {
									EndTimeType string `json:"end_time_type"`
									EndTime     string `json:"end_time"`
								} `json:"end_time_info"`
							} `json:"time_range_info"`
						} `json:"holiday_rule_info"`
					} `json:"disable_period_info"`
				} `json:"time_restrict_info"`
			} `json:"relative_period_info"`
		} `json:"voucher_use_time_info"`
	} `json:"voucher_use_rule_info"`
	VoucherDeductInfo struct {
		VoucherType    string `json:"voucher_type"`
		FixVoucherInfo struct {
			Amount      string `json:"amount"`
			FloorAmount string `json:"floor_amount"`
		} `json:"fix_voucher_info"`
		DiscountVoucherInfo struct {
			Discount      string `json:"discount"`
			CeilingAmount string `json:"ceiling_amount"`
			FloorAmount   string `json:"floor_amount"`
		} `json:"discount_voucher_info"`
		SpecialVoucherInfo struct {
			SpecialAmount string `json:"special_amount"`
			FloorAmount   string `json:"floor_amount"`
		} `json:"special_voucher_info"`
		ExchangeVoucherInfo struct {
			Amount      string `json:"amount"`
			FloorAmount string `json:"floor_amount"`
			BizType     string `json:"biz_type"`
		} `json:"exchange_voucher_info"`
	} `json:"voucher_deduct_info"`
	VoucherDisplayPatternInfo struct {
		BrandName              string   `json:"brand_name"`
		BrandLogoUrl           string   `json:"brand_logo_url"`
		VoucherDescription     string   `json:"voucher_description"`
		VoucherImageUrl        string   `json:"voucher_image_url"`
		VoucherDetailImageUrls []string `json:"voucher_detail_image_urls"`
		CustomerServiceMobile  string   `json:"customer_service_mobile"`
		CustomerServiceUrl     string   `json:"customer_service_url"`
	} `json:"voucher_display_pattern_info"`
	VoucherCustomerGuideInfo struct {
		VoucherUseGuideInfo struct {
			UseGuideMode        []string `json:"use_guide_mode"`
			MiniAppUseGuideInfo struct {
				MiniAppUrl          string   `json:"mini_app_url"`
				MiniAppServiceCodes []string `json:"mini_app_service_codes"`
			} `json:"mini_app_use_guide_info"`
		} `json:"voucher_use_guide_info"`
	} `json:"voucher_customer_guide_info"`
}

type MarketingActivityQueryAppBatchQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ActivityId string `json:"activity_id"`
	AppInfos   []struct {
		MiniAppId string `json:"mini_app_id"`
	} `json:"app_infos"`
	PageNum   int `json:"page_num"`
	PageSize  int `json:"page_size"`
	TotalSize int `json:"total_size"`
}

type MarketingActivityQueryShopBatchQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ActivityId string `json:"activity_id"`
	ShopInfos  []struct {
		ShopId    string `json:"shop_id"`
		ShopType  string `json:"shop_type"`
		ShopName  string `json:"shop_name"`
		Longitude string `json:"longitude"`
		Latitude  string `json:"latitude"`
	} `json:"shop_infos"`
	PageNum   int `json:"page_num"`
	PageSize  int `json:"page_size"`
	TotalSize int `json:"total_size"`
}

type MarketingActivityQueryGoodsBatchQueryRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	ActivityId   string `json:"activity_id"`
	AppItemInfos []struct {
		ItemId      string `json:"item_id"`
		ItemUseType string `json:"item_use_type"`
		OutItemId   string `json:"out_item_id"`
		MiniAppId   string `json:"mini_app_id"`
	} `json:"app_item_infos"`
	GoodsInfos []struct {
		GoodsId      string `json:"goods_id"`
		GoodsUseType string `json:"goods_use_type"`
	} `json:"goods_infos"`
	PageNum   int `json:"page_num"`
	PageSize  int `json:"page_size"`
	TotalSize int `json:"total_size"`
}
