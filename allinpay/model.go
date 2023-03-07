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
