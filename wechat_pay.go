package go_pay

type WechatPayClient struct {
	Params *WechatParams
	Sign   string
}

func NewWechatPayClient() *WechatPayClient {
	client := new(WechatPayClient)
	return client
}

func (this *WechatPayClient) SetParams(param *WechatParams) {
	this.Params = param
}

func (this *WechatPayClient) SetSign(secretKey string) error {
	this.Sign = getSign(secretKey, this.Params)
	return nil
}
