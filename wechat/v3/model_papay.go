package wechat

// ===================== 扣费服务（预约扣费） · 状态枚举 =====================

// 委托代扣协议状态。
// 文档：https://pay.weixin.qq.com/doc/v3/merchant/4012489245
const (
	PapayContractStateSigned      = "SIGNED"        // 签约生效中
	PapayContractStateTerminated  = "TERMINATED"    // 已解约
	PapayContractStateSignFailed  = "SIGN_FAILED"   // 用户同意签约但因业务规则限制导致签约失败
	PapayContractStateToBeRenewed = "TO_BE_RENEWED" // 可续期的计划，等待商户传入续期计划
)

// 协议解约方式。
const (
	PapayTerminationModeUser            = "USER_TERMINATE"             // 用户发起的解约
	PapayTerminationModeMchAPI          = "MCH_API_TERMINATE"          // 商户通过 API 发起的解约
	PapayTerminationModeWepayWeb        = "WEPAY_WEB_TERMINATE"        // 商户在商户平台发起的解约
	PapayTerminationModeCustomerService = "CUSTOMER_SERVICE_TERMINATE" // 用户联系微信支付客服发起的解约
	PapayTerminationModeSystem          = "SYSTEM_TERMINATE"           // 微信支付系统主动发起的解约
)

// 扣费预约状态。
const (
	PapayScheduleStateNoScheduled = "NO_SCHEDULED" // 未进行预约
	PapayScheduleStateScheduled   = "SCHEDULED"    // 已预约成功，未发起扣费或已发起扣费但扣费失败
	PapayScheduleStatePaid        = "PAID"         // 已发起扣费且扣费成功
	PapayScheduleStateExpired     = "EXPIRED"      // 超过预计扣费时间且没有扣费成功
)

// ===================== Rsp 包装 =====================

// 小程序预签约 Rsp
type PapayScheduledPreSignMiniProgramRsp struct {
	Code        int                               `json:"-"`
	SignInfo    *SignInfo                         `json:"-"`
	Response    *PapayScheduledPreSignMiniProgram `json:"response,omitempty"`
	ErrResponse ErrResponse                       `json:"err_response,omitempty"`
	Error       string                            `json:"-"`
}

// APP 预签约 Rsp
type PapayScheduledPreSignAppRsp struct {
	Code        int                       `json:"-"`
	SignInfo    *SignInfo                 `json:"-"`
	Response    *PapayScheduledPreSignApp `json:"response,omitempty"`
	ErrResponse ErrResponse               `json:"err_response,omitempty"`
	Error       string                    `json:"-"`
}

// H5 预签约 Rsp
type PapayScheduledPreSignH5Rsp struct {
	Code        int                      `json:"-"`
	SignInfo    *SignInfo                `json:"-"`
	Response    *PapayScheduledPreSignH5 `json:"response,omitempty"`
	ErrResponse ErrResponse              `json:"err_response,omitempty"`
	Error       string                   `json:"-"`
}

// JSAPI 预签约 Rsp
type PapayScheduledPreSignJsapiRsp struct {
	Code        int                         `json:"-"`
	SignInfo    *SignInfo                   `json:"-"`
	Response    *PapayScheduledPreSignJsapi `json:"response,omitempty"`
	ErrResponse ErrResponse                 `json:"err_response,omitempty"`
	Error       string                      `json:"-"`
}

// 协议查询 / 协议解约 Rsp（响应结构一致，复用同一 wrapper）
type PapayScheduledContractRsp struct {
	Code        int                     `json:"-"`
	SignInfo    *SignInfo               `json:"-"`
	Response    *PapayScheduledContract `json:"response,omitempty"`
	ErrResponse ErrResponse             `json:"err_response,omitempty"`
	Error       string                  `json:"-"`
}

// 创建预约扣费 / 查询预约扣费 Rsp（响应结构一致）
type PapayScheduledScheduleRsp struct {
	Code        int                     `json:"-"`
	SignInfo    *SignInfo               `json:"-"`
	Response    *PapayScheduledSchedule `json:"response,omitempty"`
	ErrResponse ErrResponse             `json:"err_response,omitempty"`
	Error       string                  `json:"-"`
}

// 受理扣款 Rsp
type PapayScheduledApplyRsp struct {
	Code        int                  `json:"-"`
	SignInfo    *SignInfo            `json:"-"`
	Response    *PapayScheduledApply `json:"response,omitempty"`
	ErrResponse ErrResponse          `json:"err_response,omitempty"`
	Error       string               `json:"-"`
}

// ===================== 公共子类型 =====================

// 金额信息。文档统一使用 total。
// 部分回调文档表格里写作 "amount"，但 JSON 实例仍是 "total"，按 JSON 实例为准。
type PapayScheduledAmount struct {
	Total    int    `json:"total"`              // 总金额，单位：分
	Currency string `json:"currency,omitempty"` // 币种，默认 CNY
}

// 协议解约信息。仅在 contract_state=TERMINATED 时返回。
type PapayScheduledContractTerminateInfo struct {
	ContractTerminationMode   string `json:"contract_termination_mode"`             // 协议解约方式，见 PapayTerminationMode* 常量
	ContractTerminatedTime    string `json:"contract_terminated_time"`              // 协议解约时间，rfc3339
	ContractTerminationRemark string `json:"contract_termination_remark,omitempty"` // 解约备注
}

// 预约扣费场景的预约信息。仅在预约扣费类型的协议中返回。
type PapayScheduledDeductScheduleDetail struct {
	EstimatedDeductDate   string                `json:"estimated_deduct_date,omitempty"`   // 预计扣费日期，yyyy-MM-DD
	EstimatedDeductAmount *PapayScheduledAmount `json:"estimated_deduct_amount,omitempty"` // 预计扣费金额
	ScheduleState         string                `json:"schedule_state,omitempty"`          // 扣费预约状态，见 PapayScheduleState* 常量
	ScheduledAmount       *PapayScheduledAmount `json:"scheduled_amount,omitempty"`        // 已预约的扣费金额，状态为已预约/已扣费时返回
	DeductAmount          *PapayScheduledAmount `json:"deduct_amount,omitempty"`           // 实际扣费金额，状态为已扣费时返回
	DeductDate            string                `json:"deduct_date,omitempty"`             // 实际扣费日期，状态为已扣费时返回，yyyy-MM-DD
}

// ===================== 预签约响应 =====================

// 小程序预签约响应。文档：4012525209
// 调用方据此向客户端下发用于 navigateToMiniProgram 的参数。
type PapayScheduledPreSignMiniProgram struct {
	PreEntrustwebId string `json:"pre_entrustweb_id"` // 预签约 ID，10 分钟内有效
	RedirectAppid   string `json:"redirect_appid"`    // 跳转签约小程序的 AppID
	RedirectPath    string `json:"redirect_path"`     // 跳转签约小程序的路径
}

// APP 预签约响应。文档：4012524934
// 后两字段仅在满足模板条件时返回，否则只返回 PreEntrustwebId。
type PapayScheduledPreSignApp struct {
	PreEntrustwebId     string `json:"pre_entrustweb_id"`              // 预签约 ID，10 分钟内有效
	MiniprogramUsername string `json:"miniprogram_username,omitempty"` // 跳转签约小程序的 username（仅 2025-09-23 后或申请 WXLaunchMiniProgram 权限的模板返回）
	MiniprogramPath     string `json:"miniprogram_path,omitempty"`     // 跳转签约小程序的 path（同上）
}

// H5 预签约响应。文档：4012489208
type PapayScheduledPreSignH5 struct {
	RedirectUrl string `json:"redirect_url"` // 拉起微信支付客户端的签约页面 URL
}

// JSAPI 预签约响应。文档：4012525133
type PapayScheduledPreSignJsapi struct {
	RedirectUrl string `json:"redirect_url"` // 跳转签约流程的 URL
}

// ===================== 协议详情（查询 / 解约 复用） =====================

// 委托代扣协议详情。文档：4012489245（查询）、4012489295（解约）
// 查询与解约接口返回字段完全一致。
type PapayScheduledContract struct {
	Mchid                  string                               `json:"mchid"`
	ContractId             string                               `json:"contract_id"`
	Appid                  string                               `json:"appid"`
	PlanId                 int                                  `json:"plan_id"` // 委托代扣模板 ID，文档定义为 integer
	OutContractCode        string                               `json:"out_contract_code"`
	ContractDisplayAccount string                               `json:"contract_display_account"`
	ContractState          string                               `json:"contract_state"`                  // 见 PapayContractState* 常量
	ContractSignedTime     string                               `json:"contract_signed_time,omitempty"`  // rfc3339；签约失败时不返回
	ContractExpiredTime    string                               `json:"contract_expired_time,omitempty"` // rfc3339；签约失败时不返回
	Openid                 string                               `json:"openid"`
	ContractTerminateInfo  *PapayScheduledContractTerminateInfo `json:"contract_terminate_info,omitempty"` // 仅 TERMINATED 状态返回
	OutUserCode            string                               `json:"out_user_code,omitempty"`
	DeductSchedule         *PapayScheduledDeductScheduleDetail  `json:"deduct_schedule,omitempty"` // 仅预约扣费类型模板返回
}

// ===================== 预约扣费（创建 / 查询 复用） =====================

// 预约扣费详情。文档：4012467036（创建）、4012466997（查询）
type PapayScheduledSchedule struct {
	ScheduleState   string                `json:"schedule_state"`              // 见 PapayScheduleState* 常量
	DeductStartDate string                `json:"deduct_start_date,omitempty"` // 可扣费开始日期，yyyy-MM-DD
	DeductEndDate   string                `json:"deduct_end_date,omitempty"`   // 可扣费结束日期，yyyy-MM-DD
	ScheduledAmount *PapayScheduledAmount `json:"scheduled_amount,omitempty"`  // 已预约的扣费金额
	DeductAmount    *PapayScheduledAmount `json:"deduct_amount,omitempty"`     // 实际扣费金额，预约状态为已扣费时返回
	DeductDate      string                `json:"deduct_date,omitempty"`       // 实际扣费日期，预约状态为已扣费时返回
}

// ===================== 受理扣款 =====================

// 受理扣款响应。文档：4012467087
type PapayScheduledApply struct {
	OutTradeNo string                `json:"out_trade_no"`
	Amount     *PapayScheduledAmount `json:"amount"`
}

// ===================== 回调通知载荷 =====================

// 代扣签解约结果通知载荷。文档：4012286323
// 商户用 APIv3 密钥 AES-GCM 解密 resource.ciphertext 后 Unmarshal 至此。
// 与协议查询响应（PapayScheduledContract）字段一致，但 contract_state 仅返回 SIGNED / TERMINATED。
type PapayScheduledSignNotifyResource struct {
	Mchid                  string                               `json:"mchid"`
	ContractId             string                               `json:"contract_id"`
	Appid                  string                               `json:"appid"`
	PlanId                 int                                  `json:"plan_id"`
	OutContractCode        string                               `json:"out_contract_code"`
	ContractDisplayAccount string                               `json:"contract_display_account"`
	ContractState          string                               `json:"contract_state"` // SIGNED / TERMINATED
	ContractSignedTime     string                               `json:"contract_signed_time,omitempty"`
	ContractExpiredTime    string                               `json:"contract_expired_time,omitempty"`
	Openid                 string                               `json:"openid"`
	ContractTerminateInfo  *PapayScheduledContractTerminateInfo `json:"contract_terminate_info,omitempty"`
	OutUserCode            string                               `json:"out_user_code,omitempty"`
	DeductSchedule         *PapayScheduledDeductScheduleDetail  `json:"deduct_schedule,omitempty"`
}

// 代扣支付者信息
type PapayPayer struct {
	Openid string `json:"openid"` // 用户在直连商户 AppID 下的唯一标识
}

// 代扣订单金额信息（支付通知使用）
type PapayPayNotifyAmount struct {
	Total         int    `json:"total"`          // 订单总金额，单位：分
	PayerTotal    int    `json:"payer_total"`    // 用户支付金额，单位：分
	Currency      string `json:"currency"`       // CNY
	PayerCurrency string `json:"payer_currency"` // 用户支付币种
}

// 支付场景信息
type PapayPayNotifySceneInfo struct {
	DeviceId string `json:"device_id,omitempty"`
}

// 商品列表
type PapayPayNotifyGoodsDetail struct {
	GoodsId        string `json:"goods_id"`
	Quantity       int    `json:"quantity"`
	UnitPrice      int    `json:"unit_price"`
	DiscountAmount int    `json:"discount_amount"`
	GoodsRemark    string `json:"goods_remark,omitempty"`
}

// 优惠功能
type PapayPayNotifyPromotionDetail struct {
	CouponId            string                       `json:"coupon_id"`
	Name                string                       `json:"name,omitempty"`
	Scope               string                       `json:"scope,omitempty"` // GLOBAL / SINGLE
	Type                string                       `json:"type,omitempty"`  // CASH / NOCASH
	Amount              int                          `json:"amount"`
	StockId             string                       `json:"stock_id,omitempty"`
	WechatpayContribute int                          `json:"wechatpay_contribute,omitempty"`
	MerchantContribute  int                          `json:"merchant_contribute,omitempty"`
	OtherContribute     int                          `json:"other_contribute,omitempty"`
	Currency            string                       `json:"currency,omitempty"`
	GoodsDetail         []*PapayPayNotifyGoodsDetail `json:"goods_detail,omitempty"`
}

// 代扣支付结果通知载荷。文档：4012286313
// 商户用 APIv3 密钥 AES-GCM 解密 resource.ciphertext 后 Unmarshal 至此。
type PapayScheduledPayNotifyResource struct {
	Appid           string                           `json:"appid,omitempty"`
	Mchid           string                           `json:"mchid"`
	OutTradeNo      string                           `json:"out_trade_no"`
	TransactionId   string                           `json:"transaction_id"`
	TradeType       string                           `json:"trade_type"`  // PAP: 委托代扣
	TradeState      string                           `json:"trade_state"` // SUCCESS
	TradeStateDesc  string                           `json:"trade_state_desc"`
	BankType        string                           `json:"bank_type,omitempty"`
	Attach          string                           `json:"attach,omitempty"`
	SuccessTime     string                           `json:"success_time,omitempty"` // rfc3339
	Payer           *PapayPayer                      `json:"payer,omitempty"`
	Amount          *PapayPayNotifyAmount            `json:"amount,omitempty"`
	SceneInfo       *PapayPayNotifySceneInfo         `json:"scene_info,omitempty"`
	PromotionDetail []*PapayPayNotifyPromotionDetail `json:"promotion_detail,omitempty"`
}
