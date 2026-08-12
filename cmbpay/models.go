package cmbpay

// 本文件定义常用接口的请求/响应结构体。字段命名与接口文档保持一致（JSON tag），
// 非必填字段统一使用 `omitempty`，以满足"非必填字段不上送即不出现"的要求
// （详见接口文档 2.3 第 7 条）。金额类字段单位均为【分】，以字符串传输。

// QrCodeApplyReq 收款码申请请求参数（详见接口文档 4.1.2）。
type QrCodeApplyReq struct {
	MerID               string `json:"merId"`                         // 商户号
	SubMerID            string `json:"subMerId,omitempty"`            // 附加经营商户号
	SubStoreID          string `json:"subStoreId,omitempty"`          // 经营商户门店号
	OrderID             string `json:"orderId"`                       // 商户订单号，商户下唯一
	UserID              string `json:"userId"`                        // 收银员
	TermID              string `json:"termId,omitempty"`              // 终端号
	PayValidTime        string `json:"payValidTime,omitempty"`        // 支付有效时间（秒），默认 900
	NotifyURL           string `json:"notifyUrl"`                     // 交易通知地址
	TxnAmt              string `json:"txnAmt"`                        // 交易金额（分）
	CurrencyCode        string `json:"currencyCode,omitempty"`        // 交易币种，默认 156
	Body                string `json:"body,omitempty"`                // 商品描述
	MchReserved         string `json:"mchReserved,omitempty"`         // 商户保留域
	EncryptIdentity     string `json:"encryptIdentity,omitempty"`     // 实名支付信息（需加密）
	PolicyNo            string `json:"policyNo,omitempty"`            // 保单单号
	Region              string `json:"region,omitempty"`              // 地区码
	EncryptTerminalInfo string `json:"encryptTerminalInfo,omitempty"` // 终端信息（需加密）
	BusinessParams      string `json:"businessParams,omitempty"`      // 商户传入业务信息
	EncryptTradeAddress string `json:"encryptTradeAddressInfo,omitempty"`
}

// QrCodeApplyBizContent 收款码申请业务返回数据。
type QrCodeApplyBizContent struct {
	MerID      string `json:"merId"`      // 商户号
	OrderID    string `json:"orderId"`    // 商户订单号
	QrCode     string `json:"qrCode"`     // 二维码
	CmbOrderID string `json:"cmbOrderId"` // 招行平台订单号
	TxnTime    string `json:"txnTime"`    // 订单发送时间 yyyyMMddHHmmss
}

// QrCodeApplyResp 收款码申请返回参数（详见接口文档 4.1.3）。
type QrCodeApplyResp struct {
	*PubResp
	BizContent QrCodeApplyBizContent `json:"biz_content"`
}

// OrderQueryReq 支付结果查询请求参数（详见接口文档 4.2.2）。
// orderId、cmbOrderId 至少上送一个，两者都上送时以 cmbOrderId 为准。
type OrderQueryReq struct {
	MerID        string `json:"merId"`                  // 商户号
	OrderID      string `json:"orderId,omitempty"`      // 商户订单号
	OutOrderID   string `json:"outOrderId,omitempty"`   // 外部商户订单号（先享后付）
	CmbOrderID   string `json:"cmbOrderId,omitempty"`   // 平台订单号
	UserID       string `json:"userId"`                 // 收银员
	QueryOptions string `json:"queryOptions,omitempty"` // 查询选项（支付宝）
}

// OrderQueryBizContent 支付结果查询业务返回数据。
type OrderQueryBizContent struct {
	MerID               string `json:"merId"`
	OrderID             string `json:"orderId,omitempty"`
	OutOrderID          string `json:"outOrderId,omitempty"`
	CmbOrderID          string `json:"cmbOrderId"`
	ErrDescription      string `json:"errDescription,omitempty"`      // 第三方错误描述
	TxnAmt              string `json:"txnAmt"`                        // 交易金额（分）
	DscAmt              string `json:"dscAmt"`                        // 优惠金额（分）
	CouponInfo          string `json:"couponInfo,omitempty"`          // 银联优惠信息
	PromotionDetail     string `json:"promotionDetail,omitempty"`     // 微信/支付宝优惠券信息
	IssAddnData         string `json:"issAddnData,omitempty"`         // 银联付款方附加数据
	OrderOrigAmt        string `json:"orderOrigAmt,omitempty"`        // 订单原始金额（分）
	OrderCouponAmt      string `json:"orderCouponAmt,omitempty"`      // 订单优惠金额（分）
	DiscountGoodsDetail string `json:"discountGoodsDetail,omitempty"` // 支付宝商品优惠信息
	CurrencyCode        string `json:"currencyCode"`                  // 交易币种
	PayType             string `json:"payType"`                       // 支付方式 ZF/WX/YL/EC
	OpenID              string `json:"openId,omitempty"`              // 用户标识
	OriOpenID           string `json:"oriOpenId,omitempty"`           // 用户标识
	PayBank             string `json:"payBank,omitempty"`             // 付款银行
	ThirdOrderID        string `json:"thirdOrderId,omitempty"`        // 第三方订单号
	BuyerLogonID        string `json:"buyerLogonId,omitempty"`        // 支付宝账户
	TradeState          string `json:"tradeState"`                    // 交易状态 C/D/P/F/S/R
	TxnTime             string `json:"txnTime"`                       // 订单发送时间
	EndDate             string `json:"endDate,omitempty"`             // 订单完成日期 yyyyMMdd
	EndTime             string `json:"endTime,omitempty"`             // 订单完成时间 HHmmss
	MchReserved         string `json:"mchReserved,omitempty"`         // 商户保留域
	ContractResp        string `json:"contractResp,omitempty"`        // 合约响应内容
	DebtorAgentID       string `json:"debtorAgentId,omitempty"`       // 付款运营机构编码
	DebtorAgentName     string `json:"debtorAgentName,omitempty"`     // 付款运营机构名称
	EcnyPromotionDetail string `json:"ecnyPromotionDetail,omitempty"` // 数字人民币优惠详情
}

// OrderQueryResp 支付结果查询返回参数（详见接口文档 4.2.3）。
type OrderQueryResp struct {
	*PubResp
	BizContent OrderQueryBizContent `json:"biz_content"`
}

// 交易状态取值（OrderQueryResp.TradeState / NotifyData.TradeState，详见接口文档 4.2.3）。
const (
	TradeStateClosed    = "C" // 订单已关闭
	TradeStateCanceled  = "D" // 交易已撤销
	TradeStateProcess   = "P" // 交易在进行
	TradeStateFail      = "F" // 交易失败
	TradeStateSuccess   = "S" // 交易成功
	TradeStateRefunding = "R" // 转入退款
)

// 支付方式取值（PayType）。
const (
	PayTypeAlipay = "ZF" // 支付宝
	PayTypeWeixin = "WX" // 微信
	PayTypeUnion  = "YL" // 银联
	PayTypeECNY   = "EC" // 数字人民币
)

// RefundReq 退款申请请求参数（详见接口文档 4.4.2）。
type RefundReq struct {
	MerID           string `json:"merId"`                     // 商户号
	SubMerID        string `json:"subMerId,omitempty"`        // 附加经营商户号
	SubStoreID      string `json:"subStoreId,omitempty"`      // 经营商户门店号
	OrderID         string `json:"orderId"`                   // 商户退款订单号，商户下唯一
	CmbOrderID      string `json:"cmbOrderId,omitempty"`      // 原支付平台订单号
	OrigOrderID     string `json:"origOrderId,omitempty"`     // 原商户订单号
	OrigOutOrderId  string `json:"origOutOrderId,omitempty"`  // 外部交易单号
	OrigCmbOrderId  string `json:"origCmbOrderId,omitempty"`  // 招行平台单号
	UserID          string `json:"userId"`                    // 收银员
	RefundAmt       string `json:"refundAmt"`                 // 退款金额（分）
	TxnAmt          string `json:"txnAmt,omitempty"`          // 原交易金额（分）
	NotifyURL       string `json:"notifyUrl,omitempty"`       // 退款结果通知地址
	RefundReason    string `json:"refundReason,omitempty"`    // 退款原因
	MchReserved     string `json:"mchReserved,omitempty"`     // 商户保留域
	RefundOrigAmt   string `json:"refundOrigAmt,omitempty"`   // 退单原始金额，与RefundCouponAmt同时出现
	RefundCouponAmt string `json:"refundCouponAmt,omitempty"` // 退单优惠金额，与RefundOrigAmt同时出现
}

// RefundBizContent 退款申请业务返回数据。
type RefundBizContent struct {
	MerID         string `json:"merId"`
	OrderID       string `json:"orderId"`                 // 商户退款订单号
	CmbOrderID    string `json:"cmbOrderId,omitempty"`    // 退款平台订单号
	OriCmbOrderID string `json:"oriCmbOrderId,omitempty"` // 原支付平台订单号
	RefundAmt     string `json:"refundAmt,omitempty"`     // 退款金额（分）
	TxnTime       string `json:"txnTime,omitempty"`       // 退款发送时间
	RefundState   string `json:"refundState,omitempty"`   // 退款状态
}

// RefundResp 退款申请返回参数（详见接口文档 4.4.3）。
type RefundResp struct {
	*PubResp
	BizContent RefundBizContent `json:"biz_content"`
}

// RefundQueryReq 退款结果查询请求参数（详见接口文档 4.5.2）。
type RefundQueryReq struct {
	MerID      string `json:"merId"`
	OrderID    string `json:"orderId,omitempty"`    // 商户退款订单号
	CmbOrderID string `json:"cmbOrderId,omitempty"` // 退款平台订单号
	UserID     string `json:"userId"`
}

// RefundQueryBizContent 退款结果查询业务返回数据。
type RefundQueryBizContent struct {
	MerID        string `json:"merId"`
	OrderID      string `json:"orderId,omitempty"`
	CmbOrderID   string `json:"cmbOrderId,omitempty"`
	ThirdOrderID string `json:"thirdOrderId,omitempty"` // 第三方订单号
	RefundAmt    string `json:"refundAmt,omitempty"`    // 退款金额（分）
	RefundDscAmt string `json:"refundDscAmt,omitempty"` // 退款优惠金额（分）
	CurrencyCode string `json:"currencyCode,omitempty"` // 交易币种
	TradeState   string `json:"tradeState,omitempty"`   // 退款状态 S/F/P 等
	TxnTime      string `json:"txnTime,omitempty"`      // 退款发送时间
	EndDate      string `json:"endDate,omitempty"`      // 退款完成日期 yyyyMMdd
	EndTime      string `json:"endTime,omitempty"`      // 退款完成时间 HHmmss
}

// RefundQueryResp 退款结果查询返回参数（详见接口文档 4.5.3）。
type RefundQueryResp struct {
	*PubResp
	BizContent RefundQueryBizContent `json:"biz_content"`
}

// CloseReq 关闭订单请求参数（详见接口文档 4.7.2）。
type CloseReq struct {
	MerID      string `json:"merId"`
	OrderID    string `json:"orderId,omitempty"`
	CmbOrderID string `json:"cmbOrderId,omitempty"`
	UserID     string `json:"userId"`
}

// CloseBizContent 关闭订单业务返回数据。
type CloseBizContent struct {
	MerID      string `json:"merId"`
	OrderID    string `json:"orderId,omitempty"`
	CmbOrderID string `json:"cmbOrderId,omitempty"`
}

// CloseResp 关闭订单返回参数（详见接口文档 4.7.3）。
type CloseResp struct {
	*PubResp
	BizContent CloseBizContent `json:"biz_content"`
}

// PayReq 付款码收款请求参数（详见接口文档 4.8.2）。
type PayReq struct {
	MerID               string `json:"merId"`
	SubMerID            string `json:"subMerId,omitempty"`
	SubStoreID          string `json:"subStoreId,omitempty"`
	OrderID             string `json:"orderId"`
	UserID              string `json:"userId"`
	TermID              string `json:"termId,omitempty"`
	AuthCode            string `json:"authCode"` // 付款码（用户出示的条码/二维码）
	NotifyURL           string `json:"notifyUrl,omitempty"`
	TxnAmt              string `json:"txnAmt"` // 交易金额（分）
	CurrencyCode        string `json:"currencyCode,omitempty"`
	Body                string `json:"body,omitempty"`
	MchReserved         string `json:"mchReserved,omitempty"`
	EncryptTerminalInfo string `json:"encryptTerminalInfo,omitempty"` // 终端信息（需加密）
}

// PayBizContent 付款码收款业务返回数据。
type PayBizContent struct {
	MerID          string `json:"merId"`
	OrderID        string `json:"orderId"`
	CmbOrderID     string `json:"cmbOrderId,omitempty"`
	ErrDescription string `json:"errDescription,omitempty"`
	TxnAmt         string `json:"txnAmt,omitempty"`
	PayType        string `json:"payType,omitempty"`
	TradeState     string `json:"tradeState,omitempty"`
	OpenID         string `json:"openId,omitempty"`
	ThirdOrderID   string `json:"thirdOrderId,omitempty"`
	TxnTime        string `json:"txnTime,omitempty"`
}

// PayResp 付款码收款返回参数（详见接口文档 4.8.3）。
type PayResp struct {
	*PubResp
	BizContent PayBizContent `json:"biz_content"`
}

// CancelReq 付款码支付撤销请求参数（详见接口文档 4.9.2）。
type CancelReq struct {
	MerID      string `json:"merId"`
	OrderID    string `json:"orderId,omitempty"`
	CmbOrderID string `json:"cmbOrderId,omitempty"`
	UserID     string `json:"userId"`
}

// CancelBizContent 付款码支付撤销业务返回数据。
type CancelBizContent struct {
	MerID      string `json:"merId"`
	OrderID    string `json:"orderId,omitempty"`
	CmbOrderID string `json:"cmbOrderId,omitempty"`
}

// CancelResp 付款码支付撤销返回参数（详见接口文档 4.9.3）。
type CancelResp struct {
	*PubResp
	BizContent CancelBizContent `json:"biz_content"`
}

// OnlinePayReq 微信统一下单请求参数（详见接口文档 4.10.2）。
type OnlinePayReq struct {
	MerID          string `json:"merId"`
	SubMerID       string `json:"subMerId,omitempty"`
	OrderID        string `json:"orderId"`
	UserID         string `json:"userId"`
	NotifyURL      string `json:"notifyUrl"`
	TxnAmt         string `json:"txnAmt"`
	CurrencyCode   string `json:"currencyCode,omitempty"`
	TradeType      string `json:"tradeType"`           // JSAPI 等
	SubAppID       string `json:"subAppId,omitempty"`  // 子商户公众账号 ID
	SubOpenID      string `json:"subOpenId,omitempty"` // 用户在 subAppid 下的 openid
	Body           string `json:"body,omitempty"`
	PayValidTime   string `json:"payValidTime,omitempty"`
	MchReserved    string `json:"mchReserved,omitempty"`
	SpbillCreateIP string `json:"spbillCreateIp,omitempty"`
}

// OnlinePayBizContent 微信统一下单业务返回数据。
type OnlinePayBizContent struct {
	MerID      string                      `json:"merId"`
	OrderID    string                      `json:"orderId"`
	CmbOrderID string                      `json:"cmbOrderId,omitempty"`
	PrepayID   string                      `json:"prepayId,omitempty"` // 预支付交易会话标识
	PayData    *OnlinePayBizContentPayData `json:"payData,omitempty"`  // 拉起支付所需数据
	TxnTime    string                      `json:"txnTime,omitempty"`
}

// OnlinePayBizContentPayData 统一下单 payData 字段，前端拉起支付所需参数。
type OnlinePayBizContentPayData struct {
	AppID     string `json:"appId,omitempty"`
	TimeStamp string `json:"timeStamp,omitempty"`
	NonceStr  string `json:"nonceStr,omitempty"`
	Package   string `json:"package,omitempty"`
	SignType  string `json:"signType,omitempty"`
	PaySign   string `json:"paySign,omitempty"`
}

// OnlinePayResp 微信统一下单返回参数（详见接口文档 4.10.3）。
type OnlinePayResp struct {
	*PubResp
	BizContent OnlinePayBizContent `json:"biz_content"`
}

type PubResp struct {
	ReturnCode     string `json:"returnCode,omitempty"`
	RespCode       string `json:"respCode,omitempty"`
	ErrCode        string `json:"errCode,omitempty"`
	RespMsg        string `json:"respMsg,omitempty"`
	ErrDescription string `json:"errDescription,omitempty"`
	Version        string `json:"version,omitempty"`
	Encoding       string `json:"encoding,omitempty"`
	Sign           string `json:"sign,omitempty"`
	SignMethod     string `json:"signMethod,omitempty"`
}
