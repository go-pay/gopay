//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/13 14:03
//==================================
package gopay

type weChatPayResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	Appid      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	DeviceInfo string `xml:"device_info"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	TradeType  string `xml:"trade_type"`
	PrepayId   string `xml:"prepay_id"`
	CodeUrl    string `xml:"code_url"`
	MwebUrl    string `xml:"mweb_url"`
}

type getSignKeyResponse struct {
	ReturnCode     string `xml:"return_code"`
	ReturnMsg      string `xml:"return_msg"`
	Retmsg         string `xml:"retmsg"`
	Retcode        string `xml:"retcode"`
	MchId          string `xml:"mch_id"`
	SandboxSignkey string `xml:"sandbox_signkey"`
}
