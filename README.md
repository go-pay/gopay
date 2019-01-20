
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
* 查询订单
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

未完成

## 微信

<font color='#0088ff'>注意：具体参数根据请求的不同而不同，请参考微信官方文档的参数说明！</font>

参考文档：[微信支付文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)


### 统一下单
```go
//初始化微信客户端
//    appId：应用ID
//    mchID：商户ID
//    secretKey：Key值
//    isProd：是否是正式环境
client := gopay.NewWeChatClient("wxd678efh567hg6787", "1230000109", "192006250b4c09247ec02edce69f6a2d", true)

//初始化参数Map
body := make(gopay.BodyMap)
body.Set("nonce_str", gopay.GetRandomString(32))
body.Set("body", "测试支付")
number := gopay.GetRandomString(32)
log.Println("Number:", number)
body.Set("out_trade_no", number)
body.Set("total_fee", 1)
body.Set("spbill_create_ip", "127.0.0.1")   //终端IP
body.Set("notify_url", "http://www.igoogle.ink")
body.Set("trade_type", gopay.TradeType_JsApi)
body.Set("device_info", "WEB")
body.Set("sign_type", gopay.SignType_MD5)
//body.Set("scene_info", `{"h5_info": {"type":"Wap","wap_url": "http://www.igoogle.ink","wap_name": "测试支付"}}`)
body.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu6")

//发起下单请求
wxRsp, err := client.UnifiedOrder(body)
if err != nil {
	fmt.Println("Error:", err)
	return
}
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
fmt.Println("CodeUrl:", wxRsp.CodeUrl)
fmt.Println("MwebUrl:", wxRsp.MwebUrl)
```

### 查询订单
```go
//初始化微信客户端
//    appId：应用ID
//    mchID：商户ID
//    secretKey：Key值
//    isProd：是否是正式环境
client := gopay.NewWeChatClient("wxd678efh567hg6787", "1230000109", "192006250b4c09247ec02edce69f6a2d", true)

//初始化参数结构体
body := make(gopay.BodyMap)
body.Set("out_trade_no", "CC68aTofMIwVKkVR5UruoBLFFXTAqBfv")
body.Set("nonce_str", gopay.GetRandomString(32))
body.Set("sign_type", gopay.SignType_MD5)

//请求订单查询
wxRsp, err := client.QueryOrder(body)
if err != nil {
	fmt.Println("Error:", err)
	return
}
fmt.Println("Response：", wxRsp)
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