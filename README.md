
<div align=center><img width="250" height="250" src="https://raw.githubusercontent.com/iGoogle-ink/gopay/master/logo.png"/></div>

# GoPay

<a href="https://golang.org" target="_blank"><img src="https://img.shields.io/badge/Golang-1.11+-brightgreen.svg"/></a>
<a href="https://doc.gopay.ink" target="_blank"><img src="https://img.shields.io/badge/Doc-doc.gopay.ink-blue.svg"/></a>
<img src="https://img.shields.io/badge/build-passing-brightgreen.svg"/>
<a href="https://www.apache.org/licenses/LICENSE-2.0.html" target="_blank"><img src="https://img.shields.io/badge/License-Apache-blue.svg"/></a>

## 安装

```bash
$ go get github.com/iGoogle-ink/gopay
```

## 文档

[GoPay使用手册](https://doc.gopay.ink)

## 微信支付 example

* 初始化客户端
    * AppId：应用ID
    * mchID：商户ID
    * isProd：是否是正式环境
    * secretKey：key，（当isProd为true时，此参数必传；false时，此参数为空）
```go
//正式环境 
client := gopay.NewWeChatClient("wxd678efh567hg6787", "1230000109", true, "192006250b4c09247ec02edce69f6a2d")

//沙箱环境
client := gopay.NewWeChatClient("wxd678efh567hg6787", "1230000109", false)
```

* 初始化统一下单参数
> 以下参数设置皆为必选参数，如需其他参数，请参考API文档。
>
> 参考文档：[微信支付文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)，[JSAPI支付：统一下单文档](https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1)
```go
params := new(gopay.WeChatPayParams)
params.NonceStr = "dyUNIkNS29hvDUC1CmoF0alSdfCQGg9I"
params.Body = "支付测试"
params.OutTradeNo = "GYsadfjk4dhg3fkh3ffgnlsdkf"
params.TotalFee = 10 //单位为分
params.SpbillCreateIp = "127.0.0.1"
params.NotifyUrl = "http://www.igoogle.ink"
params.TradeType = gopay.WX_PayType_JsApi //目前只支持JSAPI有效
params.DeviceInfo = "WEB"
params.SignType = gopay.WX_SignType_HMAC_SHA256 //如不设置此参数，默认为 MD5
params.Openid = "o0Df70H2Q0fY8JXh1aFPIRyOBgu8"
```

* 发起统一下单请求
    * param：统一下单请求参数
> 请求成功后，获取下单结果
```go
wxRsp, err := client.UnifiedOrder(params)
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
	fmt.Println("ErrCode：", wxRsp.ErrCode)
	fmt.Println("ErrCodeDes：", wxRsp.ErrCodeDes)
	fmt.Println("PrepayId：", wxRsp.PrepayId)
	fmt.Println("TradeType：", wxRsp.TradeType)
}
```

## 支付宝支付 example

## License

[Apache License](https://www.apache.org/licenses/LICENSE-2.0.html)