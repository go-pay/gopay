package cmbpay

// 本文件为其余接口补充类型化的请求/响应结构体（models.go 之外的接口）。
// 金额单位均为【分】，字符串传输；非必填字段使用 omitempty。

// ServPayReq 服务窗支付请求参数（详见接口文档 4.11.2）。
type ServPayReq struct {
	MerID               string `json:"merId"`
	SubMerID            string `json:"subMerId,omitempty"`
	SubStoreID          string `json:"subStoreId,omitempty"`
	DeviceInfo          string `json:"deviceInfo,omitempty"`
	OrderID             string `json:"orderId"`
	UserID              string `json:"userId"`
	Body                string `json:"body,omitempty"`
	ItemDiscount        string `json:"itemDiscount,omitempty"`
	MchReserved         string `json:"mchReserved,omitempty"`
	TimeoutExpress      string `json:"timeoutExpress,omitempty"`
	NotifyURL           string `json:"notifyUrl"`
	TxnAmt              string `json:"txnAmt"`
	CurrencyCode        string `json:"currencyCode,omitempty"`
	OrderOrigAmt        string `json:"orderOrigAmt,omitempty"`
	OrderCouponAmt      string `json:"orderCouponAmt,omitempty"`
	DisablePayChannels  string `json:"disablePayChannels,omitempty"`
	BuyerLogonID        string `json:"buyerLogonId,omitempty"`
	BuyerID             string `json:"buyerId,omitempty"`
	EncryptIdentity     string `json:"encryptIdentity,omitempty"`
	PolicyNo            string `json:"policyNo,omitempty"`
	Region              string `json:"region,omitempty"`
	PaymentNo           string `json:"paymentNo,omitempty"`
	AlipayExtendParams  string `json:"alipayExtendParams,omitempty"`
	EncryptTerminalInfo string `json:"encryptTerminalInfo,omitempty"`
	BusinessParams      string `json:"businessParams,omitempty"`
	EncryptTradeAddress string `json:"encryptTradeAddressInfo,omitempty"`
	ReturnURL           string `json:"returnUrl,omitempty"`
}

// ServPayBizContent 服务窗支付业务返回数据。
type ServPayBizContent struct {
	MerID          string `json:"merId"`
	OrderID        string `json:"orderId"`
	ErrDescription string `json:"errDescription,omitempty"`
	PayInfo        string `json:"payInfo"` // 唤起支付宝支付使用
	CmbOrderID     string `json:"cmbOrderId"`
	TxnTime        string `json:"txnTime"`
}

// ServPayResp 服务窗支付返回参数（详见接口文档 4.11.3）。
type ServPayResp struct {
	*PubResp
	BizContent ServPayBizContent `json:"biz_content"`
}

// ZfbNativeReq 支付宝 native 码支付请求参数（详见接口文档 4.12.2）。
type ZfbNativeReq struct {
	MerID                string `json:"merId"`
	SubMerID             string `json:"subMerId,omitempty"`
	SubStoreID           string `json:"subStoreId,omitempty"`
	DeviceInfo           string `json:"deviceInfo,omitempty"`
	OrderID              string `json:"orderId"`
	UserID               string `json:"userId,omitempty"`
	Body                 string `json:"body,omitempty"`
	GoodsDetail          string `json:"goodsDetail,omitempty"`
	MchReserved          string `json:"mchReserved,omitempty"`
	TimeoutExpress       string `json:"timeoutExpress,omitempty"`
	QrCodeTimeoutExpress string `json:"qrCodeTimeoutExpress,omitempty"`
	NotifyURL            string `json:"notifyUrl"`
	TxnAmt               string `json:"txnAmt"`
	OrderOrigAmt         string `json:"orderOrigAmt,omitempty"`
	OrderCouponAmt       string `json:"orderCouponAmt,omitempty"`
	CurrencyCode         string `json:"currencyCode,omitempty"`
	DisablePayChannels   string `json:"disablePayChannels,omitempty"`
	BuyerLogonID         string `json:"buyerLogonId,omitempty"`
	EncryptIdentity      string `json:"encryptIdentity,omitempty"`
	PolicyNo             string `json:"policyNo,omitempty"`
	Region               string `json:"region,omitempty"`
	PaymentNo            string `json:"paymentNo,omitempty"`
	AlipayExtendParams   string `json:"alipayExtendParams,omitempty"`
	EncryptTerminalInfo  string `json:"encryptTerminalInfo,omitempty"`
	BusinessParams       string `json:"businessParams,omitempty"`
	EncryptTradeAddress  string `json:"encryptTradeAddressInfo,omitempty"`
}

// ZfbNativeBizContent 支付宝 native 码支付业务返回数据。
type ZfbNativeBizContent struct {
	MerID          string `json:"merId"`
	OrderID        string `json:"orderId"`
	RespMsg        string `json:"respMsg,omitempty"`
	ErrDescription string `json:"errDescription,omitempty"`
	QrCode         string `json:"qrCode"` // 二维码码串
	CmbOrderID     string `json:"cmbOrderId"`
	TxnTime        string `json:"txnTime"`
}

// ZfbNativeResp 支付宝 native 码支付返回参数（详见接口文档 4.12.3）。
type ZfbNativeResp struct {
	*PubResp
	BizContent ZfbNativeBizContent `json:"biz_content"`
}

// StatementURLReq 对账单下载地址获取请求参数（详见接口文档 4.13.2）。
type StatementURLReq struct {
	MerID    string `json:"merId"`
	BillDate string `json:"billDate"`           // 格式 20180803
	BillType string `json:"billType,omitempty"` // zip 表示打包对账单
	URLType  string `json:"urlType,omitempty"`  // inner 表示行内业务网下载
}

// StatementURLBizContent 对账单下载地址获取业务返回数据。
type StatementURLBizContent struct {
	MerID           string `json:"merId"`
	FileDownloadURL string `json:"fileDownloadUrl"` // 对账单下载 URL
}

// StatementURLResp 对账单下载地址获取返回参数（详见接口文档 4.13.3）。
type StatementURLResp struct {
	*PubResp
	BizContent StatementURLBizContent `json:"biz_content"`
}

// OrderQrCodeApplyReq 订单二维码申请请求参数（详见接口文档 4.14.2）。
type OrderQrCodeApplyReq struct {
	MerID               string `json:"merId"`
	SubMerID            string `json:"subMerId,omitempty"`
	SubStoreID          string `json:"subStoreId,omitempty"`
	OrderID             string `json:"orderId"`
	UserID              string `json:"userId"`
	TermID              string `json:"termId,omitempty"`
	PayValidTime        string `json:"payValidTime,omitempty"` // 默认 1200 秒
	NotifyURL           string `json:"notifyUrl"`
	TxnAmt              string `json:"txnAmt"`
	CurrencyCode        string `json:"currencyCode,omitempty"`
	Body                string `json:"body,omitempty"`
	MchReserved         string `json:"mchReserved,omitempty"`
	EncryptTerminalInfo string `json:"encryptTerminalInfo,omitempty"`
	EncryptTradeAddress string `json:"encryptTradeAddressInfo,omitempty"`
}

// OrderQrCodeApplyBizContent 订单二维码申请业务返回数据。
type OrderQrCodeApplyBizContent struct {
	MerID   string `json:"merId"`
	OrderID string `json:"orderId"`
	QrCode  string `json:"qrCode"`
}

// OrderQrCodeApplyResp 订单二维码申请返回参数（详见接口文档 4.14.3）。
type OrderQrCodeApplyResp struct {
	*PubResp
	BizContent OrderQrCodeApplyBizContent `json:"biz_content"`
}

// MiniAppOrderReq 微信小程序下单请求参数（详见接口文档 4.15.2）。
type MiniAppOrderReq struct {
	MerID               string `json:"merId"`
	SubMerID            string `json:"subMerId,omitempty"`
	SubStoreID          string `json:"subStoreId,omitempty"`
	DeviceInfo          string `json:"deviceInfo,omitempty"`
	OrderID             string `json:"orderId"`
	TradeType           string `json:"tradeType"` // 小程序支付：JSAPI
	UserID              string `json:"userId"`
	Body                string `json:"body"`
	GoodsDetail         string `json:"goodsDetail,omitempty"`
	GoodsTag            string `json:"goodsTag,omitempty"`
	Attach              string `json:"attach,omitempty"`
	MchReserved         string `json:"mchReserved,omitempty"`
	PayValidTime        string `json:"payValidTime,omitempty"`
	NotifyURL           string `json:"notifyUrl"`
	TxnAmt              string `json:"txnAmt"`
	OrderOrigAmt        string `json:"orderOrigAmt,omitempty"`
	OrderCouponAmt      string `json:"orderCouponAmt,omitempty"`
	CurrencyCode        string `json:"currencyCode,omitempty"`
	SpbillCreateIP      string `json:"spbillCreateIp"` // 用户端 IP
	LimitPay            string `json:"limitPay,omitempty"`
	EncryptIdentity     string `json:"encryptIdentity,omitempty"`
	PolicyNo            string `json:"policyNo,omitempty"`
	Region              string `json:"region,omitempty"`
	EncryptTerminalInfo string `json:"encryptTerminalInfo,omitempty"`
	LimitPayer          string `json:"limitPayer,omitempty"`
	EncryptTradeAddress string `json:"encryptTradeAddressInfo,omitempty"`
}

// MiniAppOrderBizContent 微信小程序下单业务返回数据。
type MiniAppOrderBizContent struct {
	MerID               string `json:"merId"`
	OrderID             string `json:"orderId"`
	TradeType           string `json:"tradeType"`
	ErrDescription      string `json:"errDescription,omitempty"`
	CmbOrderID          string `json:"cmbOrderId"`
	TxnTime             string `json:"txnTime"`
	EncryptedCmbOrderID string `json:"encryptedCmbOrderId"` // 加密后的招行订单号
	EncryptedTradeInfo  string `json:"encryptedTradeInfo"`  // 加密交易数据
	CmbMiniAppID        string `json:"cmbMiniAppId"`        // 招行小程序原始 id
}

// MiniAppOrderResp 微信小程序下单返回参数（详见接口文档 4.15.3）。
type MiniAppOrderResp struct {
	*PubResp
	BizContent MiniAppOrderBizContent `json:"biz_content"`
}

// CloudPayReq 银联云闪付请求参数（详见接口文档 4.16.2）。
type CloudPayReq struct {
	MerID               string `json:"merId"`
	SubMerID            string `json:"subMerId,omitempty"`
	SubStoreID          string `json:"subStoreId,omitempty"`
	OrderID             string `json:"orderId"`
	TradeScene          string `json:"tradeScene"` // ONLINE：线上场景
	UserID              string `json:"userId"`
	Body                string `json:"body"`
	MchReserved         string `json:"mchReserved,omitempty"`
	NotifyURL           string `json:"notifyUrl"`
	TxnAmt              string `json:"txnAmt"`
	CurrencyCode        string `json:"currencyCode,omitempty"`
	FrontURL            string `json:"frontUrl,omitempty"`
	PayTimeout          string `json:"payTimeout,omitempty"`
	PayCardType         string `json:"payCardType,omitempty"`
	EncryptTradeAddress string `json:"encryptTradeAddressInfo,omitempty"`
}

// CloudPayBizContent 银联云闪付业务返回数据。
type CloudPayBizContent struct {
	MerID          string `json:"merId"`
	OrderID        string `json:"orderId"`
	TradeType      string `json:"tradeType,omitempty"`
	ErrDescription string `json:"errDescription,omitempty"`
	CmbOrderID     string `json:"cmbOrderId"`
	TxnTime        string `json:"txnTime"`
	TN             string `json:"tn"` // 银联受理订单号，调用支付控件时使用
}

// CloudPayResp 银联云闪付返回参数（详见接口文档 4.16.3）。
type CloudPayResp struct {
	*PubResp
	BizContent CloudPayBizContent `json:"biz_content"`
}

// EcnyUnifiedOrderReq 数字人民币统一下单请求参数（详见接口文档 4.17.2）。
// TransactionType：TT01 扫码支付、TT03 APP 拉起、TT04 H5 拉起、TT13 小程序拉起。
type EcnyUnifiedOrderReq struct {
	MerID               string `json:"merId"`
	OrderID             string `json:"orderId"`
	UserID              string `json:"userId"`
	NotifyURL           string `json:"notifyUrl"`
	CurrencyCode        string `json:"currencyCode"`
	TransactionType     string `json:"transactionType"`
	TxnAmt              string `json:"txnAmt"`
	TerminalNo          string `json:"terminalNo"`
	TerminalIP          string `json:"terminalIp"`
	GoodsName           string `json:"goodsName"`
	OrderDetails        string `json:"orderDetails,omitempty"`
	PlatformName        string `json:"platformName,omitempty"`
	TradePlace          string `json:"tradePlace"`
	OrderTimeExpire     string `json:"orderTimeExpire"` // yyyy-mm-ddTHH:MM:SS
	MchReserved         string `json:"mchReserved,omitempty"`
	EncryptTradeAddress string `json:"encryptTradeAddressInfo,omitempty"`
	PolicyNo            string `json:"policyNo,omitempty"`
	Region              string `json:"region,omitempty"`
	ContractReq         string `json:"contractReq,omitempty"` // 4.21 带合约时上送
}

// EcnyUnifiedOrderBizContent 数字人民币统一下单业务返回数据。
type EcnyUnifiedOrderBizContent struct {
	MerID          string            `json:"merId"`
	OrderID        string            `json:"orderId"`
	ErrDescription string            `json:"errDescription,omitempty"`
	CmbOrderID     string            `json:"cmbOrderId"`
	TxnTime        string            `json:"txnTime"`
	QrCode         string            `json:"qrCode,omitempty"`    // TT01 返回
	InvokeURL      string            `json:"invokeUrl,omitempty"` // TT03/TT04/TT13 返回
	Context        *EcnyOrderContext `json:"context,omitempty"`
	ContractResp   string            `json:"contractResp,omitempty"` // 带合约时返回
}

// EcnyUnifiedOrderResp 数字人民币统一下单返回参数（详见接口文档 4.17.3 / 4.21.3）。
type EcnyUnifiedOrderResp struct {
	*PubResp
	BizContent EcnyUnifiedOrderBizContent `json:"biz_content"`
}

// EcnyOrderContext 拉起数币 APP 的业务信息（详见接口文档 4.17.3 context 内容）。
type EcnyOrderContext struct {
	CdtrPtyID   string `json:"cdtrPtyId"`            // 收款运营机构金融编码
	MrchntNo    string `json:"mrchntNo"`             // 商户号
	EncryptKey  string `json:"encryptKey,omitempty"` // 秘钥密文
	EncryptInfo string `json:"encryptInfo,omitempty"`
	NcrptnSN    string `json:"ncrptnSN,omitempty"` // 加密证书序列号
}

// EcnyUnifiedPaymentReq 数字人民币统一支付请求参数（详见接口文档 4.18.2）。
type EcnyUnifiedPaymentReq struct {
	MerID               string `json:"merId"`
	OrderID             string `json:"orderId"`
	UserID              string `json:"userId"`
	NotifyURL           string `json:"notifyUrl"`
	CurrencyCode        string `json:"currencyCode"`
	TransactionType     string `json:"transactionType"` // 目前仅 TT01
	TxnAmt              string `json:"txnAmt"`
	TerminalNo          string `json:"terminalNo"`
	TerminalIP          string `json:"terminalIp"`
	GoodsName           string `json:"goodsName"`
	OrderDetails        string `json:"orderDetails,omitempty"`
	PlatformName        string `json:"platformName,omitempty"`
	TradePlace          string `json:"tradePlace"`
	OrderTimeExpire     string `json:"orderTimeExpire"`
	AuthCode            string `json:"authCode"` // 付款码信息
	MchReserved         string `json:"mchReserved,omitempty"`
	EncryptTradeAddress string `json:"encryptTradeAddressInfo,omitempty"`
}

// EcnyUnifiedPaymentBizContent 数字人民币统一支付业务返回数据。
type EcnyUnifiedPaymentBizContent struct {
	MerID          string `json:"merId"`
	OrderID        string `json:"orderId"`
	ErrDescription string `json:"errDescription,omitempty"`
	TradeState     string `json:"tradeState"` // P 支付中 / F 交易失败
	CmbOrderID     string `json:"cmbOrderId"`
	TxnTime        string `json:"txnTime"`
}

// EcnyUnifiedPaymentResp 数字人民币统一支付返回参数（详见接口文档 4.18.3）。
type EcnyUnifiedPaymentResp struct {
	*PubResp
	BizContent EcnyUnifiedPaymentBizContent `json:"biz_content"`
}

// EcnySubwalletPayReq 数字人民币子钱包支付请求参数（详见接口文档 4.19.2）。
// 认证方式 AuthenticCode：AC00 协议、AC01 在线认证、AC02 动态密码、AC03 短信认证。
type EcnySubwalletPayReq struct {
	MerID               string `json:"merId"`
	OrderID             string `json:"orderId"`
	UserID              string `json:"userId"`
	CurrencyCode        string `json:"currencyCode"`
	TxnAmt              string `json:"txnAmt"`
	DebtorAgentID       string `json:"debtorAgentId"` // 付款运营机构
	AuthenticCode       string `json:"authenticCode"`
	AuthenticInfo       string `json:"authenticInfo"` // AC00 时为签约协议号
	OrderTimeExpire     string `json:"orderTimeExpire,omitempty"`
	GoodsName           string `json:"goodsName"`
	SceneID             string `json:"sceneId"`
	MchReserved         string `json:"mchReserved,omitempty"`
	EncryptTradeAddress string `json:"encryptTradeAddressInfo,omitempty"`
	ContractReq         string `json:"contractReq,omitempty"` // 4.20 带合约时上送
}

// EcnySubwalletPayBizContent 数字人民币子钱包支付业务返回数据。
type EcnySubwalletPayBizContent struct {
	MerID               string `json:"merId"`
	OrderID             string `json:"orderId"`
	ErrDescription      string `json:"errDescription,omitempty"`
	CmbOrderID          string `json:"cmbOrderId"`
	TxnTime             string `json:"txnTime"`
	TxnAmt              string `json:"txnAmt"`
	DscAmt              string `json:"dscAmt"`
	TradeState          string `json:"tradeState"` // F 失败 / S 成功 / P 已受理
	ThirdOrderID        string `json:"thirdOrderId,omitempty"`
	Remark              string `json:"remark,omitempty"`
	DebtorAgentID       string `json:"debtorAgentId,omitempty"`
	DebtorAgentName     string `json:"debtorAgentName,omitempty"`
	ContractResp        string `json:"contractResp,omitempty"` // 带合约时返回
	EcnyPromotionDetail string `json:"ecnyPromotionDetail,omitempty"`
}

// EcnySubwalletPayResp 数字人民币子钱包支付返回参数（详见接口文档 4.19.3 / 4.20.3）。
type EcnySubwalletPayResp struct {
	*PubResp
	BizContent EcnySubwalletPayBizContent `json:"biz_content"`
}

// ZfbAppReq 支付宝 APP 支付请求参数（详见接口文档 4.24.2）。
type ZfbAppReq struct {
	MerID               string `json:"merId"`
	SubMerID            string `json:"subMerId,omitempty"`
	SubStoreID          string `json:"subStoreId,omitempty"`
	OrderID             string `json:"orderId"`
	UserID              string `json:"userId"`
	TxnAmt              string `json:"txnAmt"`
	CurrencyCode        string `json:"currencyCode,omitempty"`
	NotifyURL           string `json:"notifyUrl"`
	ProductCode         string `json:"productCode"` // QUICK_MSECURITY_PAY 等
	TimeExpire          string `json:"timeExpire,omitempty"`
	TimeoutExpress      string `json:"timeoutExpress,omitempty"`
	GoodsType           string `json:"goodsType,omitempty"`
	PromoParams         string `json:"promoParams,omitempty"`
	SellerID            string `json:"sellerId,omitempty"`
	Subject             string `json:"subject"`
	Body                string `json:"body,omitempty"`
	ItemDiscount        string `json:"itemDiscount,omitempty"`
	DisablePayChannels  string `json:"disablePayChannels,omitempty"`
	AlipayExtendParams  string `json:"alipayExtendParams,omitempty"`
	BusinessParams      string `json:"businessParams,omitempty"`
	EncryptIdentity     string `json:"encryptIdentity,omitempty"`
	MchReserved         string `json:"mchReserved,omitempty"`
	PolicyNo            string `json:"policyNo,omitempty"`
	Region              string `json:"region,omitempty"`
	OrderOrigAmt        string `json:"orderOrigAmt,omitempty"`
	OrderCouponAmt      string `json:"orderCouponAmt,omitempty"`
	PaymentNo           string `json:"paymentNo,omitempty"`
	EncryptTerminalInfo string `json:"encryptTerminalInfo,omitempty"`
	EncryptTradeAddress string `json:"encryptTradeAddressInfo,omitempty"`
	ReturnURL           string `json:"returnUrl,omitempty"`
}

// ZfbAppBizContent 支付宝 APP 支付业务返回数据。
type ZfbAppBizContent struct {
	MerID          string `json:"merId"`
	OrderID        string `json:"orderId"`
	ErrDescription string `json:"errDescription,omitempty"`
	CmbOrderID     string `json:"cmbOrderId"`
	TxnTime        string `json:"txnTime"`
	OrderStr       string `json:"orderStr"` // APP 支付请求参数字符串
}

// ZfbAppResp 支付宝 APP 支付返回参数（详见接口文档 4.24.3）。
type ZfbAppResp struct {
	*PubResp
	BizContent ZfbAppBizContent `json:"biz_content"`
}

// ZfbWapReq 支付宝手机网站支付请求参数（详见接口文档 4.25.2）。
type ZfbWapReq struct {
	MerID               string `json:"merId"`
	SubMerID            string `json:"subMerId,omitempty"`
	SubStoreID          string `json:"subStoreId,omitempty"`
	OrderID             string `json:"orderId"`
	UserID              string `json:"userId"`
	TxnAmt              string `json:"txnAmt"`
	CurrencyCode        string `json:"currencyCode,omitempty"`
	NotifyURL           string `json:"notifyUrl"`
	ProductCode         string `json:"productCode"` // 目前仅 QUICK_WAP_WAY
	TimeExpire          string `json:"timeExpire,omitempty"`
	TimeoutExpress      string `json:"timeoutExpress,omitempty"`
	GoodsType           string `json:"goodsType,omitempty"`
	PromoParams         string `json:"promoParams,omitempty"`
	AuthToken           string `json:"authToken,omitempty"`
	QuitURL             string `json:"quitUrl"`
	SellerID            string `json:"sellerId,omitempty"`
	Subject             string `json:"subject"`
	Body                string `json:"body,omitempty"`
	ItemDiscount        string `json:"itemDiscount,omitempty"`
	DisablePayChannels  string `json:"disablePayChannels,omitempty"`
	AlipayExtendParams  string `json:"alipayExtendParams,omitempty"`
	BusinessParams      string `json:"businessParams,omitempty"`
	EncryptIdentity     string `json:"encryptIdentity,omitempty"`
	MchReserved         string `json:"mchReserved,omitempty"`
	PolicyNo            string `json:"policyNo,omitempty"`
	Region              string `json:"region,omitempty"`
	OrderOrigAmt        string `json:"orderOrigAmt,omitempty"`
	OrderCouponAmt      string `json:"orderCouponAmt,omitempty"`
	PaymentNo           string `json:"paymentNo,omitempty"`
	EncryptTerminalInfo string `json:"encryptTerminalInfo,omitempty"`
	EncryptTradeAddress string `json:"encryptTradeAddressInfo,omitempty"`
	ReturnURL           string `json:"returnUrl,omitempty"`
}

// ZfbWapBizContent 支付宝手机网站支付业务返回数据。
type ZfbWapBizContent struct {
	MerID          string `json:"merId"`
	OrderID        string `json:"orderId"`
	ErrDescription string `json:"errDescription,omitempty"`
	CmbOrderID     string `json:"cmbOrderId"`
	TxnTime        string `json:"txnTime"`
	FormData       string `json:"formData"` // 手机网站支付 form 表单 html
}

// ZfbWapResp 支付宝手机网站支付返回参数（详见接口文档 4.25.3）。
type ZfbWapResp struct {
	*PubResp
	BizContent ZfbWapBizContent `json:"biz_content"`
}

// PapQueryReq 微信委托代扣查询请求参数（详见接口文档 4.23.2）。
type PapQueryReq struct {
	MerID      string `json:"merId"`
	OrderID    string `json:"orderId,omitempty"`
	CmbOrderID string `json:"cmbOrderId,omitempty"`
	UserID     string `json:"userId"`
}

// PapQueryBizContent 微信委托代扣查询业务返回数据。
type PapQueryBizContent struct {
	MerID                  string `json:"merId"`
	OrderID                string `json:"orderId,omitempty"`
	CmbOrderID             string `json:"cmbOrderId"`
	ErrDescription         string `json:"errDescription,omitempty"`
	TxnAmt                 string `json:"txnAmt"`
	DscAmt                 string `json:"dscAmt"`
	CurrencyCode           string `json:"currencyCode"`
	OpenID                 string `json:"openId,omitempty"`
	PayBank                string `json:"payBank,omitempty"`
	ThirdOrderID           string `json:"thirdOrderId,omitempty"`
	TradeState             string `json:"tradeState"` // P 未知 / F 失败 / S 成功
	PromotionDetail        string `json:"promotionDetail,omitempty"`
	TxnTime                string `json:"txnTime"`
	EndDate                string `json:"endDate,omitempty"`
	EndTime                string `json:"endTime,omitempty"`
	AcceptTime             string `json:"acceptTime,omitempty"`
	BusiData               string `json:"busiData,omitempty"`
	ErrorType              string `json:"errorType,omitempty"`
	DebtState              string `json:"debtState,omitempty"`
	RepaymentTransActionID string `json:"repaymentTransActionId,omitempty"`
	MchReserved            string `json:"mchReserved,omitempty"`
}

// PapQueryResp 微信委托代扣查询返回参数（详见接口文档 4.23.3）。
type PapQueryResp struct {
	*PubResp
	BizContent PapQueryBizContent `json:"biz_content"`
}

// OpenIDQryReq 微信授权码查询 openid 请求参数（详见接口文档 4.37.2）。
type OpenIDQryReq struct {
	MerID    string `json:"merId"`
	SubAppID string `json:"subAppId,omitempty"`
	AuthCode string `json:"authCode"`
}

// OpenIDQryBizContent 微信授权码查询 openid 业务返回数据。
type OpenIDQryBizContent struct {
	MerID          string `json:"merId"`
	SubAppID       string `json:"subAppId,omitempty"`
	ErrDescription string `json:"errDescription,omitempty"`
	OpenID         string `json:"openId"`
	SubOpenID      string `json:"subOpenId,omitempty"`
}

// OpenIDQryResp 微信授权码查询 openid 返回参数（详见接口文档 4.37.3）。
type OpenIDQryResp struct {
	*PubResp
	BizContent OpenIDQryBizContent `json:"biz_content"`
}

// 数字人民币交易类型（EcnyUnifiedOrderReq.TransactionType）。
const (
	EcnyTxnScanPay = "TT01" // 扫码支付
	EcnyTxnAppPay  = "TT03" // APP 拉起支付
	EcnyTxnH5Pay   = "TT04" // H5 拉起支付
	EcnyTxnMiniPay = "TT13" // 小程序拉起支付
)

// 数字人民币子钱包认证方式（EcnySubwalletPayReq.AuthenticCode）。
const (
	EcnyAuthProtocol = "AC00" // 协议方式
	EcnyAuthOnline   = "AC01" // 在线认证方式
	EcnyAuthDynamic  = "AC02" // 动态密码方式
	EcnyAuthSMS      = "AC03" // 短信认证方式
)
