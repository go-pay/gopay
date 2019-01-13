package gopay

type aliPayClient struct {
	AppId     string
	MchId     string
	secretKey string
	Params    *AliPayParams
	isProd    bool
}

//初始化支付宝客户端
//    appId：应用ID
//    mchID：商户ID
//    isProd：是否是正式环境
//    secretKey：key，（当isProd为true时，此参数必传；false时，此参数为空）
func NewAlipayClient(appId, mchId string, isProd bool, secretKey ...string) *aliPayClient {
	client := new(aliPayClient)
	client.AppId = appId
	client.MchId = mchId
	client.isProd = isProd
	if isProd && len(secretKey) > 0 {
		client.secretKey = secretKey[0]
	}
	return client
}

//统一下单
func (this aliPayClient) UnifiedOrder() {

}

//查询订单
func (this aliPayClient) QueryOrder() {

}

//关闭订单
func (this aliPayClient) CloseOrder() {

}

//申请退款
func (this aliPayClient) Refund() {

}

//查询退款
func (this aliPayClient) QueryRefund() {

}

//下载对账单
func (this aliPayClient) DownloadBill() {

}

//下载资金账单
func (this aliPayClient) DownloadFundFlow() {

}

//拉取订单评价数据
func (this aliPayClient) BatchQueryComment() {

}
