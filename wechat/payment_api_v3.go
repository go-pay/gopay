package wechat

// GetAppPaySignV3 APP调起支付，所需要的的paySign获取
//	官方文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transactions/chapter3_7.shtml
func GetAppPaySignV3(appid, prepayid, noncestr, timestamp, serialNo string) (paySign string) {
	//var (
	//	buffer strings.Builder
	//	h      hash.Hash
	//)
	//h.Write([]byte(buffer.String()))
	//paySign = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return
}
