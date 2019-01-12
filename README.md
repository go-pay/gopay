
![](https://img02.sogoucdn.com/app/a/100520146/2D860B130504E7780A8F94CFACC023DD)
# GoPay

### 文档：[https://doc.gopay.ink](https://doc.gopay.ink)

## 微信支付 example

* 初始化客户端
    * AppId：应用ID
    * mchID：商户ID
    * isProd：是否是正式环境
    * secretKey：key，（当isProd为true时，此参数必传；false时，此参数为空）

```go
	client := NewWeChatClient(AppID, MchID, false)
```

* 初始化统一下单参数
> 参数说明请参考文档：[微信支付文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)，[JSAPI支付：统一下单文档](https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1)
```go
	params := new(WeChatPayParams)
	params.NonceStr = "dyUNIkNS29hvDUC1CmoF0alSdfCQGg9I"
	params.Body = "测试充值"
	params.OutTradeNo = "GYsadfjk4dhg3fkhffgnlsdkf"
	params.TotalFee = 10 //单位为分
	params.SpbillCreateIp = "127.0.0.1"
	params.NotifyUrl = "http://www.igoogle.ink"
	params.TradeType = WX_PayType_JsApi //目前只支持JSAPI有效
	params.DeviceInfo = "WEB"
	params.SignType = WX_SignType_MD5 //如不设置此参数，默认为 MD5
	params.Openid = OpenID

	//请求支付，成功后得到结果

```

* 发起统一下单请求
    * param：统一下单请求参数
> 请求成后，获取下单结果
```go
	wxRsp, err := client.GoUnifiedOrder(params)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("ReturnCode：", wxRsp.ReturnCode)
		fmt.Println("ReturnMsg：", wxRsp.ReturnMsg)
		fmt.Println("Appid：", wxRsp.Appid)
		fmt.Println("MchId：", wxRsp.MchId)
		fmt.Println("DeviceInfo：", wxRsp.DeviceInfo)
		fmt.Println("NonceStr：", wxRsp.NonceStr)
		fmt.Println("Sign：", wxRsp.Sign)
		fmt.Println("ResultCode：", wxRsp.ResultCode)
		fmt.Println("PrepayId：", wxRsp.PrepayId)
		fmt.Println("TradeType：", wxRsp.TradeType)
	}
```

## 支付宝支付 example

### License

[MIT](https://github.com/labstack/echo/blob/master/LICENSE)