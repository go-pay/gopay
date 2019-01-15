
<div align=center><a href="https://doc.gopay.ink" target="_blank"><img width="250" height="250" alt="Logo was Loading Faild!" src="https://raw.githubusercontent.com/iGoogle-ink/gopay/master/logo.png"/></a></div>

# GoPay

<a href="https://www.igoogle.ink" target="_blank"><img src="https://img.shields.io/badge/Author-Jerry-blue.svg"/></a>
<a href="https://golang.org" target="_blank"><img src="https://img.shields.io/badge/Golang-1.11+-brightgreen.svg"/></a>
<a href="https://doc.gopay.ink" target="_blank"><img src="https://img.shields.io/badge/Doc-doc.gopay.ink-blue.svg"/></a>
<img src="https://img.shields.io/badge/Build-passing-brightgreen.svg"/>
<a href="http://www.apache.org/licenses/LICENSE-2.0" target="_blank"><img src="https://img.shields.io/badge/License-Apache 2-blue.svg"/></a>

## 微信
* 统一下单
    * JSAPI - JSAPI支付（或小程序支付）
    * NATIVE - Native支付
    * APP - app支付
    * MWEB - H5支付
* 查询订单(开发中)
* 关闭订单(开发中)
* 申请退款(开发中)
* 查询退款(开发中)
* 下载对账单(开发中)
* 下载资金账单(开发中)
* 拉取订单评价数据(开发中)


## 安装

```bash
$ go get github.com/iGoogle-ink/gopay
```

## 文档

[GoPay使用手册](https://doc.gopay.ink)

## 微信统一下单 example

* 初始化客户端
    * 参数：AppId：应用ID
    * 参数：mchID：商户ID
    * 参数：secretKey：Key值
    * 参数：isProd：是否正式环境
```go
//正式环境 
client := gopay.NewWeChatClient("wxd678efh567hg6787", "1230000109", "192006250b4c09247ec02edce69f6a2d", true)

//沙箱环境
client := gopay.NewWeChatClient("wxd678efh567hg6787", "1230000109", "192006250b4c09247ec02edce69f6a2d", false)
```

* 初始化统一下单参数
> 以下参数设置皆为必选参数，如需其他参数，请参考API文档。
>
> 参考文档：[微信支付文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)
```go
params := new(gopay.WeChatPayParams)
params.NonceStr = "dyUNIkNS29hvDUC1CmoF0alSdfCQGg9I"
params.Body = "支付测试"
params.OutTradeNo = "GYsadfjk4dhg3fkh3ffgnlsdkf"
params.TotalFee = 10 //单位为分，如沙箱环境，则默认为101
params.SpbillCreateIp = "127.0.0.1"
params.NotifyUrl = "http://www.igoogle.ink"
params.TradeType = gopay.WX_PayType_JsApi
params.DeviceInfo = "WEB"
params.SignType = gopay.WX_SignType_HMAC_SHA256 //如不设置此参数，默认为MD5，如沙箱环境，则默认为MD5
params.Openid = "o0Df70H2Q0fY8JXh1aFPIRyOBgu8" //JSAPI 方式时，此参数必填
```

* 发起统一下单请求
    * 参数：param：统一下单请求参数
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

* Coming soon.

## License
```
Copyright 2019 Jerry

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```