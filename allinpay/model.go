package allinpay

const (
	// URL
	baseUrl        = "https://vsp.allinpay.com/apiweb"
	sandboxBaseUrl = "https://syb-test.allinpay.com/apiweb"

	RSA = "RSA"
	SM2 = "SM2"

	// OrderTypeReqSN 商家自由订单号
	OrderTypeReqSN = "reqsn"
	// OrderTypeTrxId 通联生产支付号
	OrderTypeTrxId = "trxid"

	// PayTypeWXScan 微信扫码支付
	PayTypeWXScan = "W01"
	// PayTypeWXJS 微信JS支付
	PayTypeWXJS = "W02"
	// PayTypeWXMini 微信小程序支付
	// Deprecated: 使用 PayTypeWXMiniV2 代替
	PayTypeWXMini = "W03"
	// PayTypeAliScan 支付宝扫码支付
	PayTypeAliScan = "A01"
	// PayTypeAliJS 支付宝JS支付
	PayTypeAliJS = "A02"
	// PayTypeAliApp 支付宝APP支付
	PayTypeAliApp = "A03"
	// PayTypeQQScan 手机QQ扫码支付
	PayTypeQQScan = "Q01"
	// PayTypeQQJS 手机QQ JS支付
	PayTypeQQJS = "Q02"
	// PayTypeUniPay 银联扫码支付(CSB)
	PayTypeUniPay = "U01"
	// PayTypeUniJS 银联JS支付
	PayTypeUniJS = "U02"
	// PayTypeNumH5 数字货币H5
	PayTypeNumH5 = "S01"
)

// 交易方式 https://prodoc.allinpay.com/doc/273/
const (
	PayTypeWXApp    = "W03" // 微信APP支付
	PayTypeWXMiniV2 = "W06" // 微信小程序支付
	PayTypeWXNative = "W11" // 微信原生扫码支付
	PayTypeUniApp   = "U03" // 银联APP支付
	PayTypeNumHA    = "S03" // 数字货币H5/APP
	PayTypeNUCC03   = "N03" // 网联支付
)

type RspBase struct {
	RetCode string `json:"retcode"`
	RetMsg  string `json:"retmsg"`
	Sign    string `json:"sign"`
	Cusid   string `json:"cusid"`
	Appid   string `json:"appid"`
}

// PayRsp 通用支付响应
type PayRsp struct {
	RspBase
	Trxid     string `json:"trxid"`
	ChnlTrxId string `json:"chnltrxid"`
	Reqsn     string `json:"reqsn"`
	RandomStr string `json:"randomstr"`
	TrxStatus string `json:"trxstatus"`
	FinTime   string `json:"fintime"`
	ErrMsg    string `json:"errmsg"`
	PayInfo   string `json:"payinfo"`
}

// ScanPayRsp 扫码支付、订单查询响应
type ScanPayRsp struct {
	RspBase
	Trxid     string `json:"trxid"`
	ChnlTrxId string `json:"chnltrxid"`
	Reqsn     string `json:"reqsn"`
	TrxStatus string `json:"trxstatus"`
	Acct      string `json:"acct"`
	TrxCode   string `json:"trxcode"`
	FinTime   string `json:"fintime"`
	ErrMsg    string `json:"errmsg"`
	RandomStr string `json:"randomstr"`
	InitAmt   string `json:"initamt"`
	TrxAmt    string `json:"trxamt"`
	Fee       string `json:"fee"`
	Cmid      string `json:"cmid"`
	Chnlid    string `json:"chnlid"`
	ChnlData  string `json:"chnldata"`
	AcctType  string `json:"accttype"`
}

// NativePayRsp 支付主扫响应
type NativePayRsp struct {
	RspBase
	ReqSn     string `json:"reqsn"`     // 商户交易单号
	TrxStatus string `json:"trxstatus"` // 交易状态
	ErrMsg    string `json:"errmsg"`    // 错误原因
	PayInfo   string `json:"payinfo"`   // 支付串 (二维码串)
	RandomStr string `json:"randomstr"` // 随机字符串
}

// NativeCloseRsp 主扫交易关闭响应 https://prodoc.allinpay.com/doc/2439/
type NativeCloseRsp struct {
	RspBase
	TrxStatus string `json:"trxstatus"` // 交易状态
	ErrMsg    string `json:"errmsg"`    // 错误原因
	RandomStr string `json:"randomstr"` // 随机字符串
}

// RefundRsp 退款响应
type RefundRsp struct {
	RspBase
	Trxid     string `json:"trxid"`
	Reqsn     string `json:"reqsn"`
	TrxStatus string `json:"trxstatus"`
	FinTime   string `json:"fintime"`
	ErrMsg    string `json:"errmsg"`
	Fee       string `json:"fee"`
	TrxCode   string `json:"trxCode"`
	RandomStr string `json:"randomstr"`
	ChnlTrxId string `json:"chnltrxid"`
}

// CloseRsp 关闭响应
type CloseRsp struct {
	RspBase
	TrxStatus string `json:"trxstatus"`
}

// QueryConfirmRsp 交易确认查询响应 https://prodoc.allinpay.com/doc/2590/
type QueryConfirmRsp struct {
	RspBase
	Trxid     string `json:"trxid"`     // 平台的交易流水号
	ChnlTrxId string `json:"chnltrxid"` // 支付渠道交易单号
	Reqsn     string `json:"reqsn"`     // 商户的交易订单号
	TrxStatus string `json:"trxstatus"` // 交易状态
	Acct      string `json:"acct"`      // 支付平台用户标识
	TrxCode   string `json:"trxcode"`   // 交易类型
	FinTime   string `json:"fintime"`   // 交易完成时间-[yyyyMMddHHmmss]
	ErrMsg    string `json:"errmsg"`    // 错误原因
	RandomStr string `json:"randomstr"` // 随机字符串
	TrxAmt    string `json:"trxamt"`    // 交易金额
	Cmid      string `json:"cmid"`      // 渠道子商户号
	Chnlid    string `json:"chnlid"`    // 渠道商号

	InitAmt  string `json:"initamt"`  // ??, 文档未提供
	Fee      string `json:"fee"`      // ??, 文档未提供
	ChnlData string `json:"chnldata"` // ??, 文档未提供
	AcctType string `json:"accttype"` // ??, 文档未提供
}
