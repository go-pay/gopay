package wechat

// 创建支付分订单 Rsp
type ScoreOrderCreateRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *ScoreOrderCreate `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 查询支付分订单 Rsp
type ScoreOrderQueryRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *ScoreOrderQuery `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 取消支付分订单 Rsp
type ScoreOrderCancelRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *ScoreOrderCancel `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 修改订单金额 Rsp
type ScoreOrderModifyRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *ScoreOrderModify `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 完结支付分订单 Rsp
type ScoreOrderCompleteRsp struct {
	Code     int                 `json:"-"`
	SignInfo *SignInfo           `json:"-"`
	Response *ScoreOrderComplete `json:"response,omitempty"`
	Error    string              `json:"-"`
}

// 商户发起催收扣款 Rsp
type ScoreOrderPayRsp struct {
	Code     int            `json:"-"`
	SignInfo *SignInfo      `json:"-"`
	Response *ScoreOrderPay `json:"response,omitempty"`
	Error    string         `json:"-"`
}

// 同步服务订单信息 Rsp
type ScoreOrderSyncRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *ScoreOrderSync `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 创单结单合并 Rsp
type ScoreDirectCompleteRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *ScoreDirectComplete `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 创单结单合并 Rsp
type ScorePermissionRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *ScorePermission `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 查询用户授权记录（授权协议号） Rsp
type ScorePermissionQueryRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *ScorePermissionQuery `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 查询用户授权记录（openid） Rsp
type ScorePermissionOpenidQueryRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *ScorePermissionOpenidQuery `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

// =========================================================分割=========================================================

type ScoreOrderCreate struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	PostPayments        []*PostPayments  `json:"post_payments,omitempty"`     // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	RiskFund            *RiskFund        `json:"risk_fund"`                   // 订单风险金信息
	TimeRange           *TimeRange       `json:"time_range"`                  // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	Attach              string           `json:"attach,omitempty"`            // 商户数据包,可存放本订单所需信息，需要先urlencode后传入，总长度不大于256字符,超出报错处理。
	NotifyUrl           string           `json:"notify_url,omitempty"`        // 商户接收用户确认订单或扣款成功回调通知的地址。
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	Package             string           `json:"package"`                     // 用户跳转到微信侧小程序订单数据，需确认模式特有API中调起支付分-确认订单传入。该数据一小时内有效。
}

type PostPayments struct {
	Name        string `json:"name"`        // 付费项目名称
	Amount      int    `json:"amount"`      // 此付费项目总金额，大于等于0，单位为分，等于0时代表不需要扣费，只能为整数
	Description string `json:"description"` // 描述计费规则，不超过30个字符，超出报错处理。
	Count       int    `json:"count"`       // 付费项目的数量。
}

type PostDiscounts struct {
	Name        string `json:"name"`        // 优惠名称说明。
	Description string `json:"description"` // 优惠使用条件说明。
	Amount      int    `json:"amount"`      // 优惠金额，只能为整数
	Count       int    `json:"count"`       // 优惠的数量。
}

type RiskFund struct {
	Name        string `json:"name"`                  // 风险金名称。DEPOSIT：押金，ADVANCE：预付款，CASH_DEPOSIT：保证金，ESTIMATE_ORDER_COST：预估订单费用
	Description string `json:"description,omitempty"` // 风险说明
	Amount      int    `json:"amount"`                // 风险金额
}

type TimeRange struct {
	StartTime       string `json:"start_time"`                  // 服务开始时间，20091225091010
	StartTimeRemark string `json:"start_time_remark,omitempty"` // 服务开始时间备注
	EndTime         string `json:"end_time,omitempty"`          // 预计服务结束时间，20091225121010
	EndTimeRemark   string `json:"end_time_remark,omitempty"`   // 预计服务结束时间备注
}

type Location struct {
	StartLocation string `json:"start_location,omitempty"` // 服务开始地点
	EndLocation   string `json:"end_location,omitempty"`   // 服务结束位置
}

type ScoreOrderQuery struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	TotalAmount         int              `json:"total_amount,omitempty"`      // 总金额，大于等于0的数字，单位为分，只能为整数
	PostPayments        []*PostPayments  `json:"post_payments,omitempty"`     // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	RiskFund            *RiskFund        `json:"risk_fund"`                   // 订单风险金信息
	TimeRange           *TimeRange       `json:"time_range"`                  // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	Attach              string           `json:"attach,omitempty"`            // 商户数据包,可存放本订单所需信息，需要先urlencode后传入，总长度不大于256字符,超出报错处理。
	NotifyUrl           string           `json:"notify_url"`                  // 商户接收用户确认订单或扣款成功回调通知的地址。
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	NeedCollection      bool             `json:"need_collection,omitempty"`   // 是否需要收款
	Collection          *Collection      `json:"collection,omitempty"`        // 收款信息
	Openid              string           `json:"openid,omitempty"`            // 微信用户在商户对应appid下的唯一标识
}

type Collection struct {
	State        string     `json:"state"`                   // 收款状态，USER_PAYING：待支付，USER_PAID：已支付
	TotalAmount  int        `json:"total_amount,omitempty"`  // 总金额，大于等于0的数字，单位为分，只能为整数
	PayingAmount int        `json:"paying_amount,omitempty"` // 等待用户付款金额，只能为整数
	PaidAmount   int        `json:"paid_amount,omitempty"`   // 用户已付款的金额，只能为整数
	Details      []*Details `json:"details,omitempty"`       // 收款明细列表
}

type Details struct {
	Seq             int                `json:"seq,omitempty"`              // 收款序号
	Amount          int                `json:"amount,omitempty"`           // 单笔收款动作的金额，只能为整数
	PaidType        string             `json:"paid_type,omitempty"`        // 收款成功渠道，NEWTON：微信支付分，MCH：商户渠道
	PaidTime        string             `json:"paid_time,omitempty"`        // 支付成功时间，支持两种格式:yyyyMMddHHmmss和yyyyMMdd
	TransactionId   string             `json:"transaction_id,omitempty"`   // 结单交易单号，等于普通支付接口中的transaction_id
	PromotionDetail []*PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能，享受优惠时返回该字段
}

type ScoreOrderCancel struct {
	Appid      string `json:"appid"`        // 调用接口提交的公众账号Id。
	Mchid      string `json:"mchid"`        // 调用接口提交的商户号。
	ServiceId  string `json:"service_id"`   // 调用该接口提交的服务Id。
	OutOrderNo string `json:"out_order_no"` // 调用接口提交的商户服务订单号。
	OrderId    string `json:"order_id"`     // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
}

type ScoreOrderModify struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	TotalAmount         int              `json:"total_amount,omitempty"`      // 总金额，大于等于0的数字，单位为分，只能为整数
	PostPayments        []*PostPayments  `json:"post_payments,omitempty"`     // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	RiskFund            *RiskFund        `json:"risk_fund"`                   // 订单风险金信息
	TimeRange           *TimeRange       `json:"time_range"`                  // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	Attach              string           `json:"attach,omitempty"`            // 商户数据包,可存放本订单所需信息，需要先urlencode后传入，总长度不大于256字符,超出报错处理。
	NotifyUrl           string           `json:"notify_url,omitempty"`        // 商户接收用户确认订单或扣款成功回调通知的地址。
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	NeedCollection      bool             `json:"need_collection,omitempty"`   // 是否需要收款
	Collection          *Collection      `json:"collection,omitempty"`        // 收款信息
}

type ScoreOrderComplete struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	TotalAmount         int              `json:"total_amount"`                // 总金额，大于等于0的数字，单位为分，只能为整数
	PostPayments        []*PostPayments  `json:"post_payments"`               // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	RiskFund            *RiskFund        `json:"risk_fund"`                   // 订单风险金信息
	TimeRange           *TimeRange       `json:"time_range,omitempty"`        // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	NeedCollection      bool             `json:"need_collection,omitempty"`   // 是否需要收款
}

type ScoreOrderPay struct {
	Appid      string `json:"appid"`        // 调用接口提交的公众账号Id。
	Mchid      string `json:"mchid"`        // 调用接口提交的商户号。
	ServiceId  string `json:"service_id"`   // 调用该接口提交的服务Id。
	OutOrderNo string `json:"out_order_no"` // 调用接口提交的商户服务订单号。
	OrderId    string `json:"order_id"`     // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
}

type ScoreOrderSync struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	TotalAmount         int              `json:"total_amount"`                // 总金额，大于等于0的数字，单位为分，只能为整数
	PostPayments        []*PostPayments  `json:"post_payments,omitempty"`     // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	RiskFund            *RiskFund        `json:"risk_fund,omitempty"`         // 订单风险金信息
	TimeRange           *TimeRange       `json:"time_range,omitempty"`        // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	Attach              string           `json:"attach,omitempty"`            // 商户数据包,可存放本订单所需信息，需要先urlencode后传入，总长度不大于256字符,超出报错处理。
	NotifyUrl           string           `json:"notify_url,omitempty"`        // 商户接收用户确认订单或扣款成功回调通知的地址。
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	NeedCollection      bool             `json:"need_collection,omitempty"`   // 是否需要收款
	Collection          *Collection      `json:"collection,omitempty"`        // 收款信息
	Openid              string           `json:"openid"`                      // 微信用户在商户对应appid下的唯一标识
}

type ScoreDirectComplete struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	PostPayments        []*PostPayments  `json:"post_payments"`               // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	TimeRange           *TimeRange       `json:"time_range"`                  // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	Attach              string           `json:"attach,omitempty"`            // 商户数据包,可存放本订单所需信息，需要先urlencode后传入，总长度不大于256字符,超出报错处理。
	NotifyUrl           string           `json:"notify_url,omitempty"`        // 商户接收用户确认订单或扣款成功回调通知的地址。
	TotalAmount         int              `json:"total_amount"`                // 金额：数字，必须≥0（单位：分）
}

type ScorePermission struct {
	ApplyPermissionsToken string `json:"apply_permissions_token"` // 用于跳转到微信侧小程序授权数据,跳转到微信侧小程序传入，时效性为1小时
}

type ScorePermissionQuery struct {
	Appid                    string `json:"appid"`                               // 调用接口提交的公众账号Id。
	Mchid                    string `json:"mchid"`                               // 调用接口提交的商户号。
	ServiceId                string `json:"service_id"`                          // 调用该接口提交的服务Id。
	Openid                   string `json:"openid,omitempty"`                    // 微信用户在商户对应appid下的唯一标识
	AuthorizationCode        string `json:"authorization_code"`                  // 预授权成功时的授权协议号。
	AuthorizationState       string `json:"authorization_state"`                 // 标识用户授权服务情况：UNAVAILABLE：用户未授权服务，AVAILABLE：用户已授权服务，UNBINDUSER：未绑定用户（已经预授权但未完成正式授权）
	NotifyUrl                string `json:"notify_url,omitempty"`                // 商户接收用户确认订单或扣款成功回调通知的地址。
	CancelAuthorizationTime  string `json:"cancel_authorization_time,omitempty"` // 最近一次解除授权时间, 示例值：2015-05-20T13:29:35.120+08:00
	AuthorizationSuccessTime string `json:"authorization_success_time"`          // 最近一次授权成功时间, 示例值：2015-05-20T13:29:35.120+08:00
}

type ScorePermissionOpenidQuery struct {
	Appid                    string `json:"appid"`                               // 调用接口提交的公众账号Id。
	Mchid                    string `json:"mchid"`                               // 调用接口提交的商户号。
	ServiceId                string `json:"service_id"`                          // 调用该接口提交的服务Id。
	Openid                   string `json:"openid,omitempty"`                    // 微信用户在商户对应appid下的唯一标识
	AuthorizationCode        string `json:"authorization_code"`                  // 预授权成功时的授权协议号。
	AuthorizationState       string `json:"authorization_state"`                 // 标识用户授权服务情况：UNAVAILABLE：用户未授权服务，AVAILABLE：用户已授权服务，UNBINDUSER：未绑定用户（已经预授权但未完成正式授权）
	CancelAuthorizationTime  string `json:"cancel_authorization_time,omitempty"` // 最近一次解除授权时间, 示例值：2015-05-20T13:29:35.120+08:00
	AuthorizationSuccessTime string `json:"authorization_success_time"`          // 最近一次授权成功时间, 示例值：2015-05-20T13:29:35.120+08:00
}
