
<div align=center><img width="240" height="240" alt="Logo was Loading Faild!" src="https://raw.githubusercontent.com/iGoogle-ink/gopay/master/logo.png"/></div>

# GoPay

<a href="https://www.igoogle.ink" target="_blank"><img src="https://img.shields.io/badge/Author-Jerry-blue.svg"/></a>
<a href="https://golang.org" target="_blank"><img src="https://img.shields.io/badge/Golang-1.11+-brightgreen.svg"/></a>
<img src="https://img.shields.io/badge/Build-passing-brightgreen.svg"/>
<a href="http://www.apache.org/licenses/LICENSE-2.0" target="_blank"><img src="https://img.shields.io/badge/License-Apache 2-blue.svg"/></a>

## 微信
* 统一下单
    * JSAPI - JSAPI支付（或小程序支付）
    * NATIVE - Native支付
    * APP - app支付
    * MWEB - H5支付
* 提交付款码支付
* 查询订单
* 关闭订单
* 撤销订单
* 申请退款
* 查询退款
* 下载对账单
* 下载资金账单
* 拉取订单评价数据


## 安装

```bash
$ go get -u github.com/iGoogle-ink/gopay
```

## 文档

未完成，有问题请QQ或微信讨论（同号）：85411418

## 微信

<font color='#0088ff'>注意：具体参数根据请求的不同而不同，请参考微信官方文档的参数说明！</font>

参考文档：[微信支付文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)

### 付款结果回调,需回复微信平台是否成功
```go
rsp := new(gopay.WeChatNotifyResponse) //回复微信的数据

rsp.ReturnCode = "SUCCESS"
rsp.ReturnMsg = "OK"
return c.String(http.StatusOK, rsp.ToXmlString())
```

### 统一下单
```go
//初始化微信客户端
//    appId：应用ID
//    mchID：商户ID
//    secretKey：Key值
//    isProd：是否是正式环境
client := gopay.NewWeChatClient("wxd678efh567hg6787", "1230000109", "192006250b4c09247ec02edce69f6a2d", false)

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

### 提交付款码支付
```go
//初始化微信客户端
//    appId：应用ID
//    mchID：商户ID
//    secretKey：Key值
//    isProd：是否是正式环境
client := gopay.NewWeChatClient("wxd678efh567hg6787", "1230000109", "192006250b4c09247ec02edce69f6a2d", false)

//初始化参数Map
body := make(gopay.BodyMap)
body.Set("nonce_str", gopay.GetRandomString(32))
body.Set("body", "扫用户付款码支付")
number := gopay.GetRandomString(32)
log.Println("Number:", number)
body.Set("out_trade_no", number)
body.Set("total_fee", 1)
body.Set("spbill_create_ip", "127.0.0.1")
body.Set("notify_url", "http://www.igoogle.ink")
body.Set("auth_code", "120061098828009406")
body.Set("sign_type", gopay.SignType_MD5)

//请求支付，成功后得到结果
wxRsp, err := client.Micropay(body)
if err != nil {
	fmt.Println("Error:", err)
}
fmt.Println("Response:", wxRsp)
```

### 申请退款
```go
//初始化微信客户端
//    appId：应用ID
//    mchID：商户ID
//    secretKey：Key值
//    isProd：是否是正式环境
client := gopay.NewWeChatClient("wxd678efh567hg6787", "1230000109", "192006250b4c09247ec02edce69f6a2d", false)

//初始化参数结构体
body := make(gopay.BodyMap)
body.Set("out_trade_no", "MfZC2segKxh0bnJSELbvKNeH3d9oWvvQ")
body.Set("nonce_str", gopay.GetRandomString(32))
body.Set("sign_type", gopay.SignType_MD5)
s := gopay.GetRandomString(64)
fmt.Println("s:", s)
body.Set("out_refund_no", s)
body.Set("total_fee", 101)
body.Set("refund_fee", 101)

//请求申请退款（沙箱环境下，证书路径参数可传空）
//    body：参数Body
//    certFilePath：cert证书路径
//    keyFilePath：Key证书路径
//    pkcs12FilePath：p12证书路径
wxRsp, err := client.Refund(body, "", "", "")
if err != nil {
	fmt.Println("Error:", err)
}
fmt.Println("Response：", wxRsp)
```

### 查询订单
```go
client := gopay.NewWeChatClient("wxd678efh567hg6787", "1230000109", "192006250b4c09247ec02edce69f6a2d", false)

//初始化参数结构体
body := make(gopay.BodyMap)
body.Set("out_trade_no", "CC68aTofMIwVKkVR5UruoBLFFXTAqBfv")
body.Set("nonce_str", gopay.GetRandomString(32))
body.Set("sign_type", gopay.SignType_MD5)

//请求查询订单
wxRsp, err := client.QueryOrder(body)
if err != nil {
	fmt.Println("Error:", err)
	return
}
fmt.Println("Response：", wxRsp)
```

### 下载账单
```go
//初始化微信客户端
//    appId：应用ID
//    mchID：商户ID
//    secretKey：Key值
//    isProd：是否是正式环境
client := gopay.NewWeChatClient("wxd678efh567hg6787", "1230000109", "192006250b4c09247ec02edce69f6a2d", false)

//初始化参数结构体
body := make(gopay.BodyMap)
body.Set("nonce_str", gopay.GetRandomString(32))
body.Set("sign_type", gopay.SignType_MD5)
body.Set("bill_date", "20190122")
body.Set("bill_type", "ALL")

//请求下载账单，成功后得到结果（string类型）
wxRsp, err := client.DownloadBill(body)
if err != nil {
	fmt.Println("Error:", err)
}
fmt.Println("Response：", wxRsp)
```

## 支付宝支付（由于没有支付条件，暂停更新，后续补上）


* Coming soon.
* 手机网站支付流程
<div align=center><a href="https://docs.open.alipay.com/203/105285" target="_blank"><img alt="Logo was Loading Faild!" src="https://raw.githubusercontent.com/iGoogle-ink/gopay/master/alipay.jpg"/></a></div>

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