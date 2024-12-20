package wechat

type PayGiftActivityCreateRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *PayGiftActivityCreate `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

type PayGiftActivityListRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *PayGiftActivityList `json:"response,omitempty"`
	Error    string               `json:"-"`
}

type PayGiftActivityDetailRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *PayGiftActivityDetail `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

type PayGiftActivityGoodsRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *PayGiftActivityGoods `json:"response,omitempty"`
	Error    string                `json:"-"`
}

type PayGiftActivityTerminateRsp struct {
	Code     int                       `json:"-"`
	SignInfo *SignInfo                 `json:"-"`
	Response *PayGiftActivityTerminate `json:"response,omitempty"`
	Error    string                    `json:"-"`
}

type PayGiftActivityMerchantRsp struct {
	Code     int                      `json:"-"`
	SignInfo *SignInfo                `json:"-"`
	Response *PayGiftActivityMerchant `json:"response,omitempty"`
	Error    string                   `json:"-"`
}

type PayGiftActivityMerchantAddRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *PayGiftActivityMerchantAdd `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

type PayGiftActivityMerchantDeleteRsp struct {
	Code     int                            `json:"-"`
	SignInfo *SignInfo                      `json:"-"`
	Response *PayGiftActivityMerchantDelete `json:"response,omitempty"`
	Error    string                         `json:"-"`
}

// =========================================================分割=========================================================

type PayGiftActivityCreate struct {
	ActivityId string `json:"activity_id"`
	CreateTime string `json:"create_time"`
}

type PayGiftActivityList struct {
	Data []struct {
		ActivityId       string `json:"activity_id"`
		ActivityType     string `json:"activity_type"`
		ActivityBaseInfo struct {
			ActivityName        string `json:"activity_name"`
			ActivitySecondTitle string `json:"activity_second_title"`
			MerchantLogoUrl     string `json:"merchant_logo_url"`
			BackgroundColor     string `json:"background_color"`
			BeginTime           string `json:"begin_time"`
			EndTime             string `json:"end_time"`
			AvailablePeriods    struct {
				AvailableTime []struct {
					BeginTime string `json:"begin_time"`
					EndTime   string `json:"end_time"`
				} `json:"available_time"`
				AvailableDayTime []struct {
					BeginDayTime string `json:"begin_day_time"`
					EndDayTime   string `json:"end_day_time"`
				} `json:"available_day_time"`
			} `json:"available_periods"`
			OutRequestNo      string `json:"out_request_no"`
			DeliveryPurpose   string `json:"delivery_purpose"`
			MiniProgramsAppid string `json:"mini_programs_appid"`
			MiniProgramsPath  string `json:"mini_programs_path"`
		} `json:"activity_base_info"`
		AwardSendRule struct {
			FullSendRule struct {
				TransactionAmountMinimum int    `json:"transaction_amount_minimum"`
				SendContent              string `json:"send_content"`
				AwardType                string `json:"award_type"`
				AwardList                []struct {
					StockId          string `json:"stock_id"`
					OriginalImageUrl string `json:"original_image_url"`
					ThumbnailUrl     string `json:"thumbnail_url"`
				} `json:"award_list"`
				MerchantOption string   `json:"merchant_option"`
				MerchantIdList []string `json:"merchant_id_list"`
			} `json:"full_send_rule"`
		} `json:"award_send_rule"`
		AdvancedSetting struct {
			DeliveryUserCategory string `json:"delivery_user_category"`
			MerchantMemberAppid  string `json:"merchant_member_appid"`
			PaymentMode          struct {
				PaymentSceneList []string `json:"payment_scene_list"`
			} `json:"payment_mode"`
			PaymentMethodInformation struct {
				PaymentMethod    string `json:"payment_method"`
				BankAbbreviation string `json:"bank_abbreviation"`
			} `json:"payment_method_information"`
			GoodsTags []string `json:"goods_tags"`
		} `json:"advanced_setting"`
		ActivityStatus    string `json:"activity_status"`
		CreatorMerchantId string `json:"creator_merchant_id"`
		BelongMerchantId  string `json:"belong_merchant_id"`
		CreateTime        string `json:"create_time"`
		UpdateTime        string `json:"update_time"`
	} `json:"data"`
	TotalCount int `json:"total_count"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
}

type PayGiftActivityDetail struct {
	ActivityId       string `json:"activity_id"`
	ActivityType     string `json:"activity_type"`
	ActivityBaseInfo struct {
		ActivityName        string `json:"activity_name"`
		ActivitySecondTitle string `json:"activity_second_title"`
		MerchantLogoUrl     string `json:"merchant_logo_url"`
		BackgroundColor     string `json:"background_color"`
		BeginTime           string `json:"begin_time"`
		EndTime             string `json:"end_time"`
		AvailablePeriods    struct {
			AvailableTime []struct {
				BeginTime string `json:"begin_time"`
				EndTime   string `json:"end_time"`
			} `json:"available_time"`
			AvailableDayTime []struct {
				BeginDayTime string `json:"begin_day_time"`
				EndDayTime   string `json:"end_day_time"`
			} `json:"available_day_time"`
		} `json:"available_periods"`
		OutRequestNo      string `json:"out_request_no"`
		DeliveryPurpose   string `json:"delivery_purpose"`
		MiniProgramsAppid string `json:"mini_programs_appid"`
		MiniProgramsPath  string `json:"mini_programs_path"`
	} `json:"activity_base_info"`
	AwardSendRule struct {
		FullSendRule struct {
			TransactionAmountMinimum int    `json:"transaction_amount_minimum"`
			SendContent              string `json:"send_content"`
			AwardType                string `json:"award_type"`
			AwardList                []struct {
				StockId          string `json:"stock_id"`
				OriginalImageUrl string `json:"original_image_url"`
				ThumbnailUrl     string `json:"thumbnail_url"`
			} `json:"award_list"`
			MerchantOption string   `json:"merchant_option"`
			MerchantIdList []string `json:"merchant_id_list"`
		} `json:"full_send_rule"`
	} `json:"award_send_rule"`
	AdvancedSetting struct {
		DeliveryUserCategory string `json:"delivery_user_category"`
		MerchantMemberAppid  string `json:"merchant_member_appid"`
		PaymentMode          struct {
			PaymentSceneList []string `json:"payment_scene_list"`
		} `json:"payment_mode"`
		PaymentMethodInformation struct {
			PaymentMethod    string `json:"payment_method"`
			BankAbbreviation string `json:"bank_abbreviation"`
		} `json:"payment_method_information"`
		GoodsTags []string `json:"goods_tags"`
	} `json:"advanced_setting"`
	ActivityStatus    string `json:"activity_status"`
	CreatorMerchantId string `json:"creator_merchant_id"`
	BelongMerchantId  string `json:"belong_merchant_id"`
	PauseTime         string `json:"pause_time"`
	RecoveryTime      string `json:"recovery_time"`
	CreateTime        string `json:"create_time"`
	UpdateTime        string `json:"update_time"`
}

type PayGiftActivityGoods struct {
	Data []struct {
		GoodsId    string `json:"goods_id"`
		CreateTime string `json:"create_time"`
		UpdateTime string `json:"update_time"`
	} `json:"data"`
	TotalCount int    `json:"total_count"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	ActivityId string `json:"activity_id"`
}

type PayGiftActivityTerminate struct {
	ActivityId    string `json:"activity_id"`
	TerminateTime string `json:"terminate_time"`
}

type PayGiftActivityMerchant struct {
	Data []struct {
		Mchid        string `json:"mchid"`
		MerchantName string `json:"merchant_name"`
		CreateTime   string `json:"create_time"`
		UpdateTime   string `json:"update_time"`
	} `json:"data"`
	TotalCount int    `json:"total_count"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	ActivityId string `json:"activity_id"`
}

type PayGiftActivityMerchantAdd struct {
	ActivityId            string `json:"activity_id"`
	InvalidMerchantIdList []struct {
		Mchid         string `json:"mchid"`
		InvalidReason string `json:"invalid_reason"`
	} `json:"invalid_merchant_id_list"`
	AddTime string `json:"add_time"`
}

type PayGiftActivityMerchantDelete struct {
	ActivityId string `json:"activity_id"`
	DeleteTime string `json:"delete_time"`
}
