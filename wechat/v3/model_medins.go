package wechat

// 医保支付下单 Rsp
type MedInsOrderRsp struct {
	Code        int          `json:"-"`
	SignInfo    *SignInfo    `json:"-"`
	Response    *MedInsOrder `json:"response,omitempty"`
	ErrResponse ErrResponse  `json:"err_response,omitempty"`
	Error       string       `json:"-"`
}

// 医保支付查询订单 Rsp
type MedInsQueryOrderRsp struct {
	Code        int               `json:"-"`
	SignInfo    *SignInfo         `json:"-"`
	Response    *MedInsQueryOrder `json:"response,omitempty"`
	ErrResponse ErrResponse       `json:"err_response,omitempty"`
	Error       string            `json:"-"`
}

// =========================================================分割=========================================================

// MedInsOrder 医保支付下单响应
type MedInsOrder struct {
	MixTradeNo                 string `json:"mix_trade_no"`                           // 医保自费混合订单号
	MixPayStatus               string `json:"mix_pay_status"`                         // 混合订单支付状态
	SelfPayStatus              string `json:"self_pay_status"`                        // 自费部分状态
	MedInsPayStatus            string `json:"med_ins_pay_status"`                     // 医保部分状态
	PaidTime                   string `json:"paid_time,omitempty"`                    // 订单支付时间
	PassthroughResponseContent string `json:"passthrough_response_content,omitempty"` // 医保局返回内容
}

// MedInsQueryOrder 医保支付查询订单响应
type MedInsQueryOrder struct {
	MixTradeNo                 string        `json:"mix_trade_no"`                           // 医保自费混合订单号
	OutTradeNo                 string        `json:"out_trade_no"`                           // 商户订单号
	SerialNo                   string        `json:"serial_no"`                              // 医疗机构订单号
	MixPayStatus               string        `json:"mix_pay_status"`                         // 混合订单支付状态
	SelfPayStatus              string        `json:"self_pay_status"`                        // 自费部分状态
	MedInsPayStatus            string        `json:"med_ins_pay_status"`                     // 医保部分状态
	TotalFee                   int           `json:"total_fee"`                              // 总金额
	MedInsGovFee               int           `json:"med_ins_gov_fee,omitempty"`              // 医保统筹支付金额
	MedInsSelfFee              int           `json:"med_ins_self_fee,omitempty"`             // 医保个账支付金额
	MedInsOtherFee             int           `json:"med_ins_other_fee,omitempty"`            // 医保其他支付金额
	MedInsCashFee              int           `json:"med_ins_cash_fee,omitempty"`             // 医保结算后需自费金额
	WechatPayCashFee           int           `json:"wechat_pay_cash_fee,omitempty"`          // 实际微信支付金额
	CashAddDetail              []*CashDetail `json:"cash_add_detail,omitempty"`              // 现金补充列表
	CashReduceDetail           []*CashDetail `json:"cash_reduce_detail,omitempty"`           // 现金减免列表
	Appid                      string        `json:"appid"`                                  // 医疗机构公众号ID
	Openid                     string        `json:"openid"`                                 // 用户标识
	MedInstName                string        `json:"med_inst_name"`                          // 医疗机构名称
	MedInstNo                  string        `json:"med_inst_no"`                            // 医疗机构编码
	PayForRelatives            bool          `json:"pay_for_relatives,omitempty"`            // 是否代亲属支付
	PayOrderId                 string        `json:"pay_order_id,omitempty"`                 // 医保局返回的支付单ID
	PayAuthNo                  string        `json:"pay_auth_no,omitempty"`                  // 医保局返回的支付授权码
	GeoLocation                string        `json:"geo_location,omitempty"`                 // 用户定位信息
	CityId                     string        `json:"city_id"`                                // 城市ID
	MedInsOrderCreateTime      string        `json:"med_ins_order_create_time,omitempty"`    // 医保下单时间
	PaidTime                   string        `json:"paid_time,omitempty"`                    // 订单支付时间
	MixPayType                 string        `json:"mix_pay_type"`                           // 混合支付类型
	OrderType                  string        `json:"order_type"`                             // 订单类型
	PrepayId                   string        `json:"prepay_id,omitempty"`                    // 自费预下单ID
	CallbackUrl                string        `json:"callback_url"`                           // 回调通知URL
	PassthroughResponseContent string        `json:"passthrough_response_content,omitempty"` // 医保局返回内容
	PassthroughRequestContent  string        `json:"passthrough_request_content,omitempty"`  // 医疗机构透传数据
	Extends                    string        `json:"extends,omitempty"`                      // 扩展字段
	Attach                     string        `json:"attach,omitempty"`                       // 附加数据
	ChannelNo                  string        `json:"channel_no,omitempty"`                   // 渠道号
	MedInsTestEnv              bool          `json:"med_ins_test_env,omitempty"`             // 是否为医保测试环境
}

// CashDetail 现金补充/减免明细
type CashDetail struct {
	Amount int    `json:"amount"` // 金额
	Type   string `json:"type"`   // 类型
}
