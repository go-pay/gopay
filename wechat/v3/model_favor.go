package wechat


// 创建代金券批次 Rsp
type FavorBatchCreateRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *FavorBatchCreate `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 发放代金券批次 Rsp
type FavorBatchGrantRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *FavorBatchGrant `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 激活代金券批次 Rsp
type FavorBatchStartRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *FavorBatchStart `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 条件查询批次列表 Rsp
type FavorBatchListRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *FavorBatchList `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 查询批次详情 Rsp
type FavorBatchDetailRsp struct {
	Code     int         `json:"-"`
	SignInfo *SignInfo   `json:"-"`
	Response *FavorBatch `json:"response,omitempty"`
	Error    string      `json:"-"`
}

// 查询批次详情 Rsp
type FavorDetailRsp struct {
	Code     int          `json:"-"`
	SignInfo *SignInfo    `json:"-"`
	Response *FavorDetail `json:"response,omitempty"`
	Error    string       `json:"-"`
}

// 查询代金券可用商户 Rsp
type FavorMerchantRsp struct {
	Code     int            `json:"-"`
	SignInfo *SignInfo      `json:"-"`
	Response *FavorMerchant `json:"response,omitempty"`
	Error    string         `json:"-"`
}

// 查询代金券可用单品 Rsp
type FavorItemsRsp struct {
	Code     int         `json:"-"`
	SignInfo *SignInfo   `json:"-"`
	Response *FavorItems `json:"response,omitempty"`
	Error    string      `json:"-"`
}

// 根据商户号查用户的券 Rsp
type FavorUserCouponsRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *FavorUserCoupons `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 下载批次核销明细 Rsp
type FavorUseFlowDownloadRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *FavorFlowDownload `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 下载批次退款明细 Rsp
type FavorRefundFlowDownloadRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *FavorFlowDownload `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 设置消息通知地址 Rsp
type FavorCallbackUrlSetRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *FavorCallbackUrl `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 暂停代金券批次 Rsp
type FavorBatchPauseRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *FavorBatchPause `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 重启代金券批次 Rsp
type FavorBatchRestartRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *FavorBatchRestart `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 创建商家券 Rsp
type BusiFavorCreateRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *FavorBatchCreate `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 查询商家券详情 Rsp
type BusiFavorBatchDetailRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *BusiFavorBatchDetail `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 核销用户券 Rsp
type BusiFavorUseRsp struct {
	Code     int           `json:"-"`
	SignInfo *SignInfo     `json:"-"`
	Response *BusiFavorUse `json:"response,omitempty"`
	Error    string        `json:"-"`
}

// 核销用户券 Rsp
type BusiFavorUserCouponsRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *BusiFavorUserCoupons `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 查询用户单张券详情 Rsp
type BusiFavorUserCouponDetailRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *BusiUserCoupon `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 上传预存code Rsp
type BusiFavorCodeUploadRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *BusiFavorCodeUpload `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 设置商家券事件通知地址 Rsp
type BusiFavorCallbackUrlSetRsp struct {
	Code     int                      `json:"-"`
	SignInfo *SignInfo                `json:"-"`
	Response *BusiFavorCallbackUrlSet `json:"response,omitempty"`
	Error    string                   `json:"-"`
}

// 查询商家券事件通知地址 Rsp
type BusiFavorCallbackUrlRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *BusiFavorCallbackUrl `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 关联订单信息 Rsp
type BusiFavorAssociateRsp struct {
	Code     int                 `json:"-"`
	SignInfo *SignInfo           `json:"-"`
	Response *BusiFavorAssociate `json:"response,omitempty"`
	Error    string              `json:"-"`
}

// 取消关联订单信息 Rsp
type BusiFavorDisassociateRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *BusiFavorDisassociate `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

// 修改批次预算 Rsp
type BusiFavorBatchUpdateRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *BusiFavorBatchUpdate `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 发放消费卡 Rsp
type BusiFavorSendRsp struct {
	Code     int            `json:"-"`
	SignInfo *SignInfo      `json:"-"`
	Response *BusiFavorSend `json:"response,omitempty"`
	Error    string         `json:"-"`
}

// 申请退券 Rsp
type BusiFavorReturnRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *BusiFavorReturn `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 使券失效 Rsp
type BusiFavorDeactivateRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *BusiFavorDeactivate `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 营销补差付款 Rsp
type BusiFavorSubsidyPayRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *BusiFavorSubsidyPay `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 查询营销补差付款单详情 Rsp
type BusiFavorSubsidyPayDetailRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *BusiFavorSubsidyPay `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// =========================================================分割=========================================================

type FavorBatchCreate struct {
	StockId    string `json:"stock_id"`    // 批次号
	CreateTime string `json:"create_time"` // 创建时间，遵循rfc3339标准格式
}

type FavorBatchGrant struct {
	CouponId string `json:"coupon_id"` // 微信为代金券唯一分配的id
}

type FavorBatchStart struct {
	StockId   string `json:"stock_id"`   // 微信为每个代金券批次分配的唯一Id
	StartTime string `json:"start_time"` // 生效时间，遵循rfc3339标准格式
}

type FavorBatchList struct {
	Data       []*FavorBatch `json:"data,omitempty"` // 批次详情列表
	TotalCount int           `json:"total_count"`    // 批次总数
	Offset     int           `json:"offset"`         // 分页页码
	Limit      int           `json:"limit"`          // 分页大小
}

type FavorBatch struct {
	StockId            string        `json:"stock_id"`             // 微信为每个代金券批次分配的唯一Id
	StockCreatorMchid  string        `json:"stock_creator_mchid"`  // 创建批次的商户号
	StockName          string        `json:"stock_name"`           // 批次名称
	Status             string        `json:"status"`               // 批次状态
	CreateTime         string        `json:"create_time"`          // 创建时间，遵循rfc3339标准格式
	Description        string        `json:"description"`          // 使用说明
	StockUseRule       *StockUseRule `json:"stock_use_rule"`       // 普通发券批次特定信息
	AvailableBeginTime string        `json:"available_begin_time"` // 可用开始时间，遵循rfc3339标准格式
	AvailableEndTime   string        `json:"available_end_time"`   // 可用结束时间，遵循rfc3339标准格式
	DistributedCoupons int           `json:"distributed_coupons"`  // 已发券数量
	NoCash             bool          `json:"no_cash"`              // 是否无资金流
	StartTime          string        `json:"start_time"`           // 激活批次的时间
	StopTime           string        `json:"stop_time"`            // 终止批次的时间
	CutToMessage       *CutToMessage `json:"cut_to_message"`       // 单品优惠特定信息
	Singleitem         bool          `json:"singleitem"`           // 是否单品优惠
	StockType          string        `json:"stock_type"`           // 批次类型
}

type StockUseRule struct {
	MaxCoupons        int                `json:"max_coupons"`          // 发放总上限
	MaxAmount         int                `json:"max_amount"`           // 总预算
	MaxAmountByDay    int                `json:"max_amount_by_day"`    // 当天发放上限金额
	FixedNormalCoupon *FixedNormalCoupon `json:"fixed_normal_coupon"`  // 固定面额发券批次特定信息
	MaxCouponsPerUser int                `json:"max_coupons_per_user"` // 单个用户可领个数
	CouponType        string             `json:"coupon_type"`          // 券类型
	GoodsTag          []string           `json:"goods_tag,omitempty"`  // 订单优惠标记
	TradeType         []string           `json:"trade_type"`           // 支付方式
	CombineUse        bool               `json:"combine_use"`          // 是否可叠加其他优惠
}

type FixedNormalCoupon struct {
	CouponAmount       int `json:"coupon_amount"`       // 面额，单位：分
	TransactionMinimum int `json:"transaction_minimum"` // 使用券金额门槛，单位：分
}

type CutToMessage struct {
	SinglePriceMax int `json:"single_price_max"` // 可用优惠的商品最高单价，单位：分
	CutToPrice     int `json:"cut_to_price"`     // 减至后的优惠单价，单位：分
}

type FavorDetail struct {
	StockId                 string                   `json:"stock_id"`                  // 微信为每个代金券批次分配的唯一Id
	StockCreatorMchid       string                   `json:"stock_creator_mchid"`       // 创建批次的商户号
	CouponId                string                   `json:"coupon_id"`                 // 微信为代金券唯一分配的id
	CutToMessage            *CutToMessage            `json:"cut_to_message"`            // 单品优惠特定信息
	CouponName              string                   `json:"coupon_name"`               // 代金券名称
	Status                  string                   `json:"status"`                    // 代金券状态
	Description             string                   `json:"description"`               // 使用说明
	CreateTime              string                   `json:"create_time"`               // 领券时间，遵循rfc3339标准格式
	CouponType              string                   `json:"coupon_type"`               // 券类型
	NoCash                  bool                     `json:"no_cash"`                   // 是否无资金流
	AvailableBeginTime      string                   `json:"available_begin_time"`      // 可用开始时间，遵循rfc3339标准格式
	AvailableEndTime        string                   `json:"available_end_time"`        // 可用结束时间，遵循rfc3339标准格式
	Singleitem              bool                     `json:"singleitem"`                // 是否单品优惠
	NormalCouponInformation *NormalCouponInformation `json:"normal_coupon_information"` // 普通满减券面额、门槛信息
}

type NormalCouponInformation struct {
	CouponAmount       int `json:"coupon_amount"`       // 面额，单位：分
	TransactionMinimum int `json:"transaction_minimum"` // 使用券金额门槛，单位：分
}

type FavorMerchant struct {
	StockId    string   `json:"stock_id"`       // 批次号
	Data       []string `json:"data,omitempty"` // 可用商户列表
	TotalCount int      `json:"total_count"`    // 批次总数
	Offset     int      `json:"offset"`         // 分页页码
	Limit      int      `json:"limit"`          // 分页大小
}

type FavorItems struct {
	StockId    string   `json:"stock_id"`       // 批次号
	Data       []string `json:"data,omitempty"` // 可用商户列表
	TotalCount int      `json:"total_count"`    // 批次总数
	Offset     int      `json:"offset"`         // 分页页码
	Limit      int      `json:"limit"`          // 分页大小
}

type FavorUserCoupons struct {
	Data       []*UserCoupon `json:"data,omitempty"` // 批次详情列表
	TotalCount int           `json:"total_count"`    // 批次总数
	Offset     int           `json:"offset"`         // 分页页码
	Limit      int           `json:"limit"`          // 分页大小
}

type UserCoupon struct {
	StockCreatorMchid       string                   `json:"stock_creator_mchid"` // 创建批次的商户号
	StockId                 string                   `json:"stock_id"`            // 批次号
	CouponId                string                   `json:"coupon_id"`
	CouponName              string                   `json:"coupon_name"`
	CouponType              string                   `json:"coupon_type"`
	CutToMessage            *CutToMessage            `json:"cut_to_message"`            // 单品优惠特定信息
	Status                  string                   `json:"status"`                    // 代金券状态
	Description             string                   `json:"description"`               // 使用说明
	CreateTime              string                   `json:"create_time"`               // 领券时间，遵循rfc3339标准格式
	NoCash                  bool                     `json:"no_cash"`                   // 是否无资金流
	AvailableBeginTime      string                   `json:"available_begin_time"`      // 可用开始时间，遵循rfc3339标准格式
	AvailableEndTime        string                   `json:"available_end_time"`        // 可用结束时间，遵循rfc3339标准格式
	Singleitem              bool                     `json:"singleitem"`                // 是否单品优惠
	NormalCouponInformation *NormalCouponInformation `json:"normal_coupon_information"` // 普通满减券面额、门槛信息
	ConsumeInformation      *ConsumeInformation      `json:"consume_information"`       // 已实扣代金券信息
}

type ConsumeInformation struct {
	ConsumeTime   string               `json:"consume_time"`           // 核销时间，遵循rfc3339标准格式
	ConsumeMchid  string               `json:"consume_mchid"`          // 核销商户号
	TransactionId string               `json:"transaction_id"`         // 核销订单号
	GoodsDetail   []*CouponGoodsDetail `json:"goods_detail,omitempty"` // 商户下单单品信息
}

type CouponGoodsDetail struct {
	GoodsId        string `json:"goods_id"`        // 商品编码
	Quantity       int    `json:"quantity"`        // 商品数量
	Price          int    `json:"price"`           // 单品价格，单位为分
	DiscountAmount int    `json:"discount_amount"` // 商品优惠金额
}

type FavorFlowDownload struct {
	Url       string `json:"url"`        // 流水文件下载链接，30s内有效
	HashValue string `json:"hash_value"` // 文件内容的哈希值，防止篡改
	HashType  string `json:"hash_type"`  // 哈希算法类型，目前只支持sha1
}

type FavorCallbackUrl struct {
	UpdateTime string `json:"update_time"` // 修改时间，遵循rfc3339标准格式
	NotifyUrl  string `json:"notify_url"`  // 通知地址
}

type FavorBatchPause struct {
	PauseTime string `json:"pause_time"` // 暂停时间，遵循rfc3339标准格式
	StockId   string `json:"stock_id"`   // 批次号
}

type FavorBatchRestart struct {
	RestartTime string `json:"restart_time"` // 生效时间，遵循rfc3339标准格式
	StockId     string `json:"stock_id"`     // 批次号
}

type BusiFavorBatchDetail struct {
	StockId              string                `json:"stock_id"`                         // 批次号
	StockName            string                `json:"stock_name"`                       // 商家券批次名称
	BelongMerchant       string                `json:"belong_merchant"`                  // 批次归属商户号
	Comment              string                `json:"comment"`                          // 批次备注
	GoodsName            string                `json:"goods_name"`                       // 适用商品范围
	StockType            string                `json:"stock_type"`                       // 批次类型
	StockState           string                `json:"stock_state"`                      // 批次状态
	CouponUseRule        *CouponUseRule        `json:"coupon_use_rule"`                  // 券核销相关规则
	StockSendRule        *StockSendRule        `json:"stock_send_rule"`                  // 券发放相关规则
	CustomEntrance       *CustomEntrance       `json:"custom_entrance,omitempty"`        // 自定义入口
	DisplayPatternInfo   *DisplayPatternInfo   `json:"display_pattern_info,omitempty"`   // 创建批次时的样式信息
	CouponCodeMode       string                `json:"coupon_code_mode"`                 // 券code模式
	CouponCodeCount      *CouponCodeCount      `json:"coupon_code_count,omitempty"`      // 券code数量
	NotifyConfig         *NotifyConfig         `json:"notify_config,omitempty"`          // 事件通知配置
	SendCountInformation *SendCountInformation `json:"send_count_information,omitempty"` // 批次发放情况
}

type CouponUseRule struct {
	CouponAvailableTime *CouponAvailableTime `json:"coupon_available_time"`
	FixedNormalCoupon   *struct {
		DiscountAmount     int `json:"discount_amount"`     // 优惠金额，单位：分
		TransactionMinimum int `json:"transaction_minimum"` // 消费门槛，单位：分
	} `json:"fixed_normal_coupon,omitempty"` // 固定面额满减券使用规则
	DiscountCoupon *struct {
		DiscountPercent    int `json:"discount_percent"`    // 折扣百分比，例如：86为八六折
		TransactionMinimum int `json:"transaction_minimum"` // 消费门槛，单位：分
	} `json:"discount_coupon,omitempty"` // 折扣券使用规则
	ExchangeCoupon *struct {
		ExchangePrice      int `json:"exchange_price"`      // 单品换购价，单位：分
		TransactionMinimum int `json:"transaction_minimum"` // 消费门槛，单位：分
	} `json:"exchange_coupon,omitempty"` // 换购券使用规则
	UseMethod         string `json:"use_method"`          // 核销方式
	MiniProgramsAppid string `json:"mini_programs_appid"` // 小程序appid
	MiniProgramsPath  string `json:"mini_programs_path"`  // 小程序path
}

type CouponAvailableTime struct {
	AvailableBeginTime       string           `json:"available_begin_time"`        // 批次开始时间，遵循rfc3339标准格式
	AvailableEndTime         string           `json:"available_end_time"`          // 批次结束时间，遵循rfc3339标准格式
	AvailableDayAfterReceive int              `json:"available_day_after_receive"` // 生效后N天内有效
	AvailableWeek            *AvailableWeek   `json:"available_week"`              // 固定周期有效时间段
	IrregularyAvaliableTime  []*AvailableTime `json:"irregulary_avaliable_time"`   // 无规律的有效时间段
	WaitDaysAfterReceive     int              `json:"wait_days_after_receive"`     // 领取后N天开始生效
}

type AvailableWeek struct {
	WeekDay          []int            `json:"week_day"` // 可用星期数，0代表周日，1代表周一，以此类推
	AvailableDayTime []*AvailableTime `json:"available_day_time"`
}

type AvailableTime struct {
	BeginTime int `json:"begin_time"` // 开始时间
	EndTime   int `json:"end_time"`   // 结束时间
}

type StockSendRule struct {
	MaxAmount          int  `json:"max_amount"`           // 批次总预算
	MaxCoupons         int  `json:"max_coupons"`          // 批次最大发放个数
	MaxCouponsPerUser  int  `json:"max_coupons_per_user"` // 用户最大可领个数
	MaxAmountByDay     int  `json:"max_amount_by_day"`    // 单天发放上限金额
	MaxCouponsByDay    int  `json:"max_coupons_by_day"`   // 单天发放上限个数
	NaturalPersonLimit bool `json:"natural_person_limit"` // 是否开启自然人限制
	PreventApiAbuse    bool `json:"prevent_api_abuse"`    // 可疑账号拦截
	Transferable       bool `json:"transferable"`         // 是否允许转赠
	Shareable          bool `json:"shareable"`            // 是否允许分享链接
}

type CustomEntrance struct {
	MiniProgramsInfo *MiniProgramsInfo `json:"mini_programs_info"` // 小程序入口
	Appid            string            `json:"appid"`              // 商户公众号appid
	HallId           string            `json:"hall_id"`            // 营销馆id
	StoreId          string            `json:"store_id"`           // 可用门店id
	CodeDisplayMode  string            `json:"code_display_mode"`  // code展示模式
}

type MiniProgramsInfo struct {
	MiniProgramsAppid string `json:"mini_programs_appid"` // 商家小程序appid
	MiniProgramsPath  string `json:"mini_programs_path"`  // 商家小程序path
	EntranceWords     string `json:"entrance_words"`      // 入口文案
	GuidingWords      string `json:"guiding_words"`       // 引导文案
}

type DisplayPatternInfo struct {
	Description     string `json:"description,omitempty"` // 使用说明
	MerchantLogoUrl string `json:"merchant_logo_url"`     // 商户logo
	MerchantName    string `json:"merchant_name"`         // 商户名称
	BackgroundColor string `json:"background_color"`      // 背景颜色
	CouponImageUrl  string `json:"coupon_image_url"`      // 券详情图片
}

type CouponCodeCount struct {
	TotalCount     int `json:"total_count"`     // 该批次总共已上传的code总数
	AvailableCount int `json:"available_count"` // 该批次当前可用的code数
}

type NotifyConfig struct {
	NotifyAppid string `json:"notify_appid"` // 事件通知appid
}

type SendCountInformation struct {
	TotalSendNum    int `json:"total_send_num"`    // 已发放券张数
	TotalSendAmount int `json:"total_send_amount"` // 已发放券金额
	TodaySendNum    int `json:"today_send_num"`    // 单天已发放券张数
	TodaySendAmount int `json:"today_send_amount"` // 单天已发放券金额
}

type BusiFavorUse struct {
	StockId          string `json:"stock_id"`           // 批次号
	Openid           string `json:"openid"`             // 用户在公众号内的唯一身份标识
	WechatpayUseTime string `json:"wechatpay_use_time"` // 系统成功核销券的时间，遵循rfc3339标准
}

type BusiFavorUserCoupons struct {
	Data       []*BusiUserCoupon `json:"data,omitempty"` // 结果集
	TotalCount int               `json:"total_count"`    // 批次总数
	Offset     int               `json:"offset"`         // 分页页码
	Limit      int               `json:"limit"`          // 分页大小
}

type BusiUserCoupon struct {
	BelongMerchant     string              `json:"belong_merchant"`                // 批次归属商户号
	StockName          string              `json:"stock_name"`                     // 批次名称
	Comment            string              `json:"comment"`                        // 批次备注
	GoodsName          string              `json:"goods_name"`                     // 适用商品范围
	StockType          string              `json:"stock_type"`                     // 批次类型
	Transferable       bool                `json:"transferable"`                   // 是否允许转赠
	Shareable          bool                `json:"shareable"`                      // 是否允许分享链接
	CouponState        string              `json:"coupon_state"`                   // 商家券状态
	DisplayPatternInfo *DisplayPatternInfo `json:"display_pattern_info,omitempty"` // 创建批次时的样式信息
	CouponUseRule      *CouponUseRule      `json:"coupon_use_rule"`                // 券核销相关规则
	CustomEntrance     *CustomEntrance     `json:"custom_entrance,omitempty"`      // 自定义入口
	CouponCode         string              `json:"coupon_code"`                    // 券的唯一标识
	StockId            string              `json:"stock_id"`                       // 批次号
	AvailableStartTime string              `json:"available_start_time"`           // 券可使用开始时间
	ExpireTime         string              `json:"expire_time"`                    // 券过期时间
	ReceiveTime        string              `json:"receive_time"`                   // 券领券时间
	SendRequestNo      string              `json:"send_request_no"`                // 发券请求单号
	UseRequestNo       string              `json:"use_request_no"`                 // 核销请求单号
	UseTime            string              `json:"use_time"`                       // 券核销时间
}

type BusiFavorCodeUpload struct {
	StockId        string      `json:"stock_id"`        // 批次号
	TotalCount     int         `json:"total_count"`     // 去重后上传code总数
	SuccessCount   int         `json:"success_count"`   // 上传成功code个数
	SuccessCodes   []string    `json:"success_codes"`   // 上传成功的code列表
	SuccessTime    string      `json:"success_time"`    // 上传成功时间
	FailCount      int         `json:"fail_count"`      // 上传失败code个数
	FailCodes      []*FailCode `json:"fail_codes"`      // 上传失败的code及原因
	ExistCodes     []string    `json:"exist_codes"`     // 已存在的code列表
	DuplicateCodes []string    `json:"duplicate_codes"` // 本次请求中重复的code列表
}

type FailCode struct {
	CouponCode string `json:"coupon_code"` // 上传失败的券code
	Code       string `json:"code"`        // 上传失败错误码
	Message    string `json:"message"`     // 上传失败错误信息
}

type BusiFavorCallbackUrlSet struct {
	UpdateTime string `json:"update_time"` // 修改时间
	NotifyUrl  string `json:"notify_url"`  // 通知URL地址
	Mchid      string `json:"mchid"`       // 商户号
}

type BusiFavorCallbackUrl struct {
	NotifyUrl string `json:"notify_url"` // 通知URL地址
	Mchid     string `json:"mchid"`      // 商户号
}

type BusiFavorAssociate struct {
	WechatpayAssociateTime string `json:"wechatpay_associate_time"` // 关联成功时间
}

type BusiFavorDisassociate struct {
	WechatpayDisassociateTime string `json:"wechatpay_disassociate_time"` // 取消关联成功时间
}

type BusiFavorBatchUpdate struct {
	MaxCoupons      int `json:"max_coupons"`        // 批次当前最大发放个数
	MaxCouponsByDay int `json:"max_coupons_by_day"` // 当前单天发放上限个数
}

type BusiFavorSend struct {
	CardCode string `json:"card_code"` // 消费卡code
}

type BusiFavorReturn struct {
	WechatpayReturnTime string `json:"wechatpay_return_time"` // 微信退券成功的时间
}

type BusiFavorDeactivate struct {
	WechatpayDeactivateTime string `json:"wechatpay_deactivate_time"` // 券成功失效的时间
}

type BusiFavorSubsidyPay struct {
	SubsidyReceiptId string `json:"subsidy_receipt_id"` // 补差付款单号
	StockId          string `json:"stock_id"`           // 商家券批次号
	CouponCode       string `json:"coupon_code"`        // 券的唯一标识
	TransactionId    string `json:"transaction_id"`     // 微信支付系统生成的订单号
	PayerMerchant    string `json:"payer_merchant"`     // 营销补差扣款商户号
	PayeeMerchant    string `json:"payee_merchant"`     // 营销补差入账商户号
	Amount           int    `json:"amount"`             // 补差付款金额
	Description      string `json:"description"`        // 补差付款描述
	Status           string `json:"status"`             // 补差付款单据状态
	FailReason       string `json:"fail_reason"`        // 补差付款失败原因
	SuccessTime      string `json:"success_time"`       // 补差付款完成时间
	OutSubsidyNo     string `json:"out_subsidy_no"`     // 业务请求唯一单号
	CreateTime       string `json:"create_time"`        // 补差付款发起时间
}
