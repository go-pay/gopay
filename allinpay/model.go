package allinpay

const (
	// URL
	baseUrl        = "https://vsp.allinpay.com/apiweb"
	sandboxBaseUrl = "https://syb-test.allinpay.com/apiweb"

	RSA = "RSA"
	SM2 = "SM2"

	OrderTypeReqSN = "reqsn"
	OrderTypeTrxId = "trxid"

	//PayTypeWXScan 微信扫码支付
	PayTypeWXScan = "W01"
	//PayTypeWXJs 微信JS支付
	PayTypeWXJS = "W02"
	//PayTypeWXMini 微信小程序支付
	PayTypeWXMini = "W03"
	//PayTypeAliScan 支付宝扫码支付
	PayTypeAliScan = "A01"
	//PayTypeAliJS 支付宝JS支付
	PayTypeAliJS = "A02"
	//PayTypeAliApp 支付宝APP支付
	PayTypeAliApp = "A03"
	//PayTypeQQScan 手机QQ扫码支付
	PayTypeQQScan = "Q01"
	//PayTypeQQJS 手机QQ JS支付
	PayTypeQQJS = "Q02"
	//PayTypeUniPay 银联扫码支付(CSB)
	PayTypeUniPay = "U01"
	//PayTypeUniJS 银联JS支付
	PayTypeUniJS = "U02"
	//PayTypeNumH5 数字货币H5
	PayTypeNumH5 = "S01"
)

type RespBase struct {
	RetCode string `json:"retcode"`
	RetMsg  string `json:"retmsg"`
	Sign    string `json:"sign"`
	Cusid   string `json:"cusid"`
	Appid   string `json:"appid"`
}
