
<div align=center><img width="220" height="220" alt="Logo was Loading Faild!" src="https://raw.githubusercontent.com/iGoogle-ink/gopay/master/logo.png"/></div>

# GoPay

[中文文档](https://github.com/iGoogle-ink/gopay/blob/master/README.md)

The Golang SDK for WeChat and AliPay

[![](https://img.shields.io/badge/Author-Jerry-blue.svg)](https://www.gopay.ink)
[![](https://img.shields.io/badge/Golang-1.11+-brightgreen.svg)](https://golang.org)
[![](https://img.shields.io/badge/Version-1.3.1-blue.svg)](https://www.gopay.ink)
[![](https://api.travis-ci.org/iGoogle-ink/gopay.svg?branch=master)]()
[![](https://img.shields.io/badge/License-Apache_2-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0)

## WeChat Client

* client := gopay.NewWeChatClient() => Init WeChat Payment Client
* client's other configuration
    * client.SetCountry() => set country area，default China（gopay.China、gopay.China2、gopay.SoutheastAsia、gopay.Other）

## WeChat Payment API

* UnifiedOrder：client.UnifiedOrder()
    * JSAPI - Official Account Payment（or WeChat Applet Payment）
    * NATIVE - Native Payment
    * APP - In-App Payment
    * MWEB - H5 Payment
* Quick Pay：client.Micropay()
* QueryOrder：client.QueryOrder()
* CloseOrder：client.CloseOrder()
* Reverse：client.Reverse()
* Refund：client.Refund()
* QueryRefund：client.QueryRefund()
* DownloadBill：client.DownloadBill()
* DownloadFundFlow：client.DownloadFundFlow()
* BatchQueryComment：client.BatchQueryComment()

## WeChat Public API

* gopay.GetWeChatParamSign() => 获取微信支付所需参数里的Sign值（通过支付参数计算Sign值）
* gopay.GetWeChatSanBoxParamSign() => 获取微信支付沙箱环境所需参数里的Sign值（通过支付参数计算Sign值）
* gopay.GetMiniPaySign() => Obtain the paySign required for WeChat Applet Payment
* gopay.GetH5PaySign() => Obtain the paySign required for H5 Payment in WeChat
* gopay.GetAppPaySign() => Obtain the paySign required for App Payment
* gopay.ParseWeChatNotifyResultToBodyMap() => Parse the parameters of WeChat Payment asynchronous notification to BodyMap
* gopay.ParseWeChatNotifyResult() => Parse the parameters of WeChat Payment asynchronous notification to Struct
* gopay.VerifyWeChatSign() =>Verify WeChat Response Sign
* gopay.Code2Session() => Login certificate verification：Obtain WeChat user's OpenId, UnionId, SessionKey
* gopay.GetAccessToken() => Obtain WeChat Applet's global unique access token
* gopay.GetPaidUnionId() => After the WeChat Applet user's payment is completed, obtain the UnionId of the user without authorization
* gopay.GetWeChatUserInfo() => WeChat Official Account：Obtain Basic User Information (UnionID System)
* gopay.DecryptOpenDataToStruct() => Decrypt encrypted data to the specified struct
* gopay.GetOpenIdByAuthCode() => Authorization code query openid

---

## Alipay Client

* client := gopay.NewAliPayClient() => Init Alipay Payment Client
* client request some settings of Common Parameters
    * client.SetReturnUrl() => set the return URL
    * client.SetNotifyUrl() => set the asynchronous notification URL
    * client.SetCharset() => set the character,default utf-8
    * client.SetSignType() => set the sign type,default RSA2
    * client.SetAppAuthToken() => set the app auth token
    * client.SetAuthToken() => set the personal information auth token

## Alipay Payment API

* alipay.trade.wap.pay（Wap Payment）：client.AliPayTradeWapPay()
* alipay.trade.page.pay（PC Web Payment）：client.AliPayTradePagePay()
* alipay.trade.app.pay（In-App Payment）：client.AliPayTradeAppPay()
* alipay.trade.pay（Merchant Scan User）：client.AliPayTradePay()
* alipay.trade.create（Alipay Applet Payment）：client.AliPayTradeCreate()
* alipay.trade.query：client.AliPayTradeQuery()
* alipay.trade.close：client.AliPayTradeClose()
* alipay.trade.cancel：client.AliPayTradeCancel()
* alipay.trade.refund：client.AliPayTradeRefund()
* alipay.trade.page.refund：client.AliPayTradePageRefund()
* alipay.trade.fastpay.refund.query：client.AliPayTradeFastPayRefundQuery()
* alipay.trade.order.settle：client.AliPayTradeOrderSettle()
* alipay.trade.precreate（User Scan Merchant）：client.AliPayTradePrecreate()
* alipay.fund.trans.toaccount.transfer：client.AlipayFundTransToaccountTransfer()
* alipay.system.oauth.token（obtain access_token, user_id and so on）：client.AliPaySystemOauthToken()
* zhima.credit.score.get：client.ZhimaCreditScoreGet()

## Alipay Public API

* gopay.AliPaySystemOauthToken() => Obtain authorized access token（obtain access_token, user_id and so on）
* gopay.FormatPrivateKey() => Format private key
* gopay.FormatAliPayPublicKey() => Format alipay public key
* gopay.ParseAliPayNotifyResult() => Parse the parameters of Alipay Payment asynchronous notification to Struct
* gopay.VerifyAliPaySign() => Verify Alipay Response Sign
* gopay.DecryptAliPayOpenDataToStruct() => Decrypt alipay applet encrypted data to the specified struct

# Install

```bash
$ go get -u github.com/iGoogle-ink/gopay
```

## View the GoPay version

```go
package main

import (
    "fmt"
    "github.com/iGoogle-ink/gopay"
)

func main() {
    fmt.Println("GoPay Version: ", gopay.Version)
}
```

# Document

* [GoDoc](https://godoc.org/github.com/iGoogle-ink/gopay)

* If you have any questions, Please add Wechat Group. Please click a star.

WeChat QrCode：
<img width="226" height="300" alt="Photo was Loading Faild!" src="https://raw.githubusercontent.com/iGoogle-ink/gopay/master/wechat_jerry.png"/>
WeChat Group：
<img width="226" height="300" alt="Photo was Loading Faild!" src="https://raw.githubusercontent.com/iGoogle-ink/gopay/master/wechat_gopay.png"/>

# WeChat Payment

<font color='#0088ff'>Attention：Specific parameters vary depending on the request. Please refer to the parameter description in the official document of WeChat.</font>

Reference Documents：[WeChat Payment Documents](https://pay.weixin.qq.com/wiki/doc/api/index.html)

---
### Obtain WeChat user's OpenId, UnionId, SessionKey

```go
userIdRsp, err := gopay.Code2Session(appID, secretKey, "")
if err != nil {
	fmt.Println("Error:", err)
	return
}
fmt.Println("OpenID:", userIdRsp.Openid)
fmt.Println("UnionID:", userIdRsp.Unionid)
fmt.Println("SessionKey:", userIdRsp.SessionKey)
```

### WeChat Applet Payment, After the success of the unified order, Obtain the paySign required for WeChat Applet Payment

* GetMiniPaySign
    * timeStamp
    * nonceStr
    * package 
    * signType
    * paySign

> WeChat Applet Payment Document：[Document](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/payment/wx.requestPayment.html)
```go
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
packages := "prepay_id=" + wxRsp.PrepayId   // wxRsp.PrepayId
//    appId：APPID
//    nonceStr：
//    prepayId：
//    signType：
//    timeStamp：
//    apiKey：API KEY
paySign := gopay.GetMiniPaySign(AppID, wxRsp.NonceStr, packages, gopay.SignType_MD5, timeStamp, ApiKey)

fmt.Println("paySign：", paySign)
```

### WeChat H5 Payment, After the success of the unified order, Obtain the paySign required for WeChat H5 Payment

* GetH5PaySign
    * appId
    * timeStamp
    * nonceStr
    * package 
    * signType
    * paySign
> WeChat H5 Payment Document：[Document](https://pay.weixin.qq.com/wiki/doc/api/external/jsapi.php?chapter=7_7&index=6)
```go
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
packages := "prepay_id=" + wxRsp.PrepayId   // wxRsp.PrepayId
//    appId：APPID
//    nonceStr：
//    prepayId：
//    signType：
//    timeStamp：
//    apiKey：API KEY
paySign := gopay.GetH5PaySign(AppID, wxRsp.NonceStr, packages, gopay.SignType_MD5, timeStamp, ApiKey)

fmt.Println("paySign：", paySign)
```

### WeChat In-App Payment, After the success of the unified order, Obtain the paySign required for WeChat In-App Payment

* GetAppPaySign
    * appid
    * partnerid
    * noncestr
    * prepayid
    * package 
    * timestamp
    * sign
> WeChat In-App Payment Document：[Document](https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_12)
```go
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
//Attention：signType：sign type must be the same as when unified order
//Attention：package：The parameter is a fixed value and does not need to be passed in again.
//    appId：APPID
//    partnerid：
//    nonceStr：
//    prepayId：
//    signType：sign type must be the same as when unified order
//    timeStamp：
//    apiKey：API KEY
paySign := gopay.GetAppPaySign(appid, partnerid, wxRsp.NonceStr, prepayid, gopay.SignType_MD5, timeStamp, apiKey)

fmt.Println("paySign：", paySign)
```

### 1、Parse the parameters of WeChat Payment asynchronous notification，2、 Verify the Sign of WeChat Payment asynchronous notification

> WeChat Payment asynchronous notification document：[Notification of Payment Result](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_7&index=8)

```go
//解析微信支付异步通知的参数
//    req：*http.Request
//    返回参数notifyRsp：Notify请求的参数
//    返回参数err：错误信息
notifyRsp, err := gopay.ParseWeChatNotifyResult(c.Request())
if err != nil {
    fmt.Println("err:", err)
    return
}
fmt.Println("notifyRsp:", notifyRsp)

//验证微信API返回结果或异步通知结果的Sign值
//    apiKey：API秘钥值
//    signType：签名类型（调用API方法时填写的类型）
//    bean：微信API返回的结构体 wxRsp 或 异步通知解析的结构体 notifyReq
//    返回参数ok：是否验签通过
//    返回参数err：错误信息
ok, err := gopay.VerifyWeChatSign("192006250b4c09247ec02edce69f6a2d", "MD5", notifyRsp)
if err != nil {
    fmt.Println("err:", err)
}
fmt.Println("ok:", ok)
```
或者
```go
//解析微信支付异步通知的结果到BodyMap
//    req：*http.Request
//    返回参数bm：Notify请求的参数
//    返回参数err：错误信息
bm, err := gopay.ParseWeChatNotifyResultToBodyMap(c.Request())
if err != nil {
    fmt.Println("err:", err)
    return
}

//验证微信API返回结果或异步通知结果的Sign值
//    apiKey：API秘钥值
//    signType：签名类型（调用API方法时填写的类型）
//    bean：微信API返回的结构体 wxRsp 或 异步通知解析的结构体 notifyReq
//    返回参数ok：是否验签通过
//    返回参数err：错误信息
ok, err := gopay.VerifyWeChatSign("192006250b4c09247ec02edce69f6a2d", gopay.SignType_MD5, bm)
if err != nil {
    fmt.Println("err:", err)
}
fmt.Println("ok:", ok)
```

### Decrypt encrypted data to the specified struct

> Take WeChat Applet to get phone number as an example

button get phone number：[button](https://developers.weixin.qq.com/miniprogram/dev/component/button.html)

WeChat decryption algorithm document：[Decryption Algorithm](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)
```go
data := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
iv := "Cds8j3VYoGvnTp1BrjXdJg=="
sessionKey := "lyY4HPQbaOYzZdG+JcYK9w=="

phone := new(gopay.WeChatUserPhone)
err := gopay.DecryptOpenDataToStruct(data, iv, sessionKey, phone)
if err != nil {
	fmt.Println("err:", err)
	return
}
fmt.Println("PhoneNumber:", phone.PhoneNumber)
fmt.Println("PurePhoneNumber:", phone.PurePhoneNumber)
fmt.Println("CountryCode:", phone.CountryCode)
fmt.Println("Watermark:", phone.Watermark)
```

### WeChat Payment asynchronous notification, Respond to WeChat Platform Success

> Method of return in code, Because of I use [Echo](https://github.com/labstack/echo), I Recommend it

```go
rsp := new(gopay.WeChatNotifyResponse) // the struct return WeChat Platform

rsp.ReturnCode = gopay.SUCCESS
rsp.ReturnMsg = gopay.OK

return c.String(http.StatusOK, rsp.ToXmlString())
```

### UnifiedOrder
```go
//Init WeChat Payment Client
//    appId：
//    mchID：
//    apiKey：API KEY
//    isProd：
client := gopay.NewWeChatClient("wxd678efh567hg6787", "1230000109", "192006250b4c09247ec02edce69f6a2d", false)

number := gopay.GetRandomString(32)
fmt.Println("out_trade_no:", number)
//Init BodyMap
body := make(gopay.BodyMap)
body.Set("nonce_str", gopay.GetRandomString(32))
body.Set("body", "测试支付")
body.Set("out_trade_no", number)
body.Set("total_fee", 1)
body.Set("spbill_create_ip", "127.0.0.1")
body.Set("notify_url", "http://www.gopay.ink")
body.Set("trade_type", gopay.TradeType_H5)
body.Set("device_info", "WEB")
body.Set("sign_type", gopay.SignType_MD5)

//sceneInfo := make(map[string]map[string]string)
//h5Info := make(map[string]string)
//h5Info["type"] = "Wap"
//h5Info["wap_url"] = "http://www.gopay.ink"
//h5Info["wap_name"] = "H5测试支付"
//sceneInfo["h5_info"] = h5Info
//body.Set("scene_info", sceneInfo)

body.Set("openid", OpenID)

//Request
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

### Quick Pay
```go
//Init WeChat Payment Client
//    appId：
//    mchID：
//    apiKey：API KEY
//    isProd：
client := gopay.NewWeChatClient("wxd678efh567hg6787", "1230000109", "192006250b4c09247ec02edce69f6a2d", false)

//Init BodyMap
body := make(gopay.BodyMap)
body.Set("nonce_str", gopay.GetRandomString(32))
body.Set("body", "扫用户付款码支付")
number := gopay.GetRandomString(32)
log.Println("Number:", number)
body.Set("out_trade_no", number)
body.Set("total_fee", 1)
body.Set("spbill_create_ip", "127.0.0.1")
body.Set("notify_url", "http://www.gopay.ink")
body.Set("auth_code", "120061098828009406")
body.Set("sign_type", gopay.SignType_MD5)

//Request
wxRsp, err := client.Micropay(body)
if err != nil {
	fmt.Println("Error:", err)
}
fmt.Println("Response:", wxRsp)
```

### Refund
```go
//Init WeChat Payment Client
//    appId：
//    mchID：
//    apiKey：API KEY
//    isProd：
client := gopay.NewWeChatClient("wxd678efh567hg6787", "1230000109", "192006250b4c09247ec02edce69f6a2d", false)

//Init BodyMap
body := make(gopay.BodyMap)
body.Set("out_trade_no", "MfZC2segKxh0bnJSELbvKNeH3d9oWvvQ")
body.Set("nonce_str", gopay.GetRandomString(32))
body.Set("sign_type", gopay.SignType_MD5)
s := gopay.GetRandomString(64)
fmt.Println("s:", s)
body.Set("out_refund_no", s)
body.Set("total_fee", 101)
body.Set("refund_fee", 101)

//Request Refund（Sandbox environment，Certificate path parameters can be null）
//    body：BodyMap
//    certFilePath：Certificate path
//    keyFilePath：Key path
//    pkcs12FilePath：p12 path
wxRsp, err := client.Refund(body, "", "", "")
if err != nil {
	fmt.Println("Error:", err)
}
fmt.Println("Response：", wxRsp)
```
---

# Alipay Payment

<font color='#0088ff'>Attention：Specific parameters vary depending on the request. Please refer to the parameter description in the official document of Alipay.</font>

Alipay official documents：[官方文档](https://docs.open.alipay.com/catalog)

Alipay RSA secret key generation document：[Generating RSA Key](https://docs.open.alipay.com/291/105971/)

Alipay online debugging：[Online Debugging](https://openhome.alipay.com/platform/demoManage.htm)

Instructions for Environmental Use of Sandbox：[Document](https://docs.open.alipay.com/200/105311)

---

### Obtain authorized access token（obtain access_token, user_id and so on）

> Obtain alipay authorized access token document：[Obtain Access Token](https://docs.open.alipay.com/api_9/alipay.system.oauth.token)

```go
privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
//Obtain authorized access token（default utf-8, RSA2）
//    appId：
//    privateKey：
//    grantType：When the value is authorization_code, it is obtained by Code. When the value is refresh_token, it is obtained by refresh_token, and null is obtained by code.
//    codeOrToken：authorization_code or refresh_token
rsp, err := gopay.AlipaySystemOauthToken("2016091200494382", privateKey, "authorization_code", "06e8961891d647c0ac99bb1cebe7SE69")
if err != nil {
	fmt.Println("gopay.AlipaySystemOauthToken:",err)
	return
}
fmt.Println("rsp:", *rsp)
```

### Decrypt alipay applet encrypted data to the specified struct

> Take Alipay Applet to get phone number as an example

Obtain user phone number document：[Obtain Phone Number](https://docs.alipay.com/mini/api/getphonenumber)

Alipay decryption algorithm document：[Decryption Algorithm](https://docs.alipay.com/mini/introduce/aes)
```go
data := "MkvuiIZsGOC8S038cu/JIpoRKnF+ZFjoIRGf5d/K4+ctYjCtb/eEkwgrdB5TeH/93bxff1Ylb+SE+UGStlpvcg=="
key := "TDftre9FpItr46e9BVNJcw=="
rsp := new(gopay.PhoneNumberResponse)
err := gopay.DecryptAliPayOpenDataToStruct(data, key, rsp)
if err != nil {
	fmt.Println("err:", err)
	return
}
fmt.Println("rsp.Code:", rsp.Code)
fmt.Println("rsp.Msg:", rsp.Msg)
fmt.Println("rsp.SubCode:", rsp.SubCode)
fmt.Println("rsp.SubMsg:", rsp.SubMsg)
fmt.Println("rsp.Mobile:", rsp.Mobile)
```

### 1、Parse the parameters of Alipay Payment asynchronous notification，2、 Verify the Sign of Alipay Payment asynchronous notification

> Alipay Payment asynchronous notification document：[Notification of Payment Result](https://docs.open.alipay.com/200/106120)

```go
//Parse the parameters of WeChat Payment asynchronous notification
// c.Request() is *http.Request
notifyRsp, err := gopay.ParseAliPayNotifyResult(c.Request())
if err != nil {
    fmt.Println("gopay.ParseAliPayNotifyResult:", err)
    return
}
fmt.Println("notifyRsp:", notifyRsp)

aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"
//Verify the Sign of Alipay Payment asynchronous notification
//    aliPayPublicKey：
//    notifyRsp：Struct obtained by gopay.ParseAliPayNotifyResult()
//    return：Is passed
//    return：error
ok, err := gopay.VerifyAliPayResultSign(aliPayPublicKey, notifyRsp)
if err != nil {
	log.Println("gopay.VerifyAliPayResultSign:", err)
	return
}
fmt.Println("ok:", ok)
```

### Alipay Payment asynchronous notification, Respond to Alipay Platform Success

* The output "success" (without quotes) must be printed after the program has finished executing. If the merchant's feedback to Alipay is not the 7 characters of success, the Alipay server will continue to resend the notification until it exceeds 24 hours and 22 minutes. Under normal circumstances, 8 notifications are completed within 25 hours (the interval between notifications is generally: 4m, 10m, 10m, 1h, 2h, 6h, 15h)

> Method of return in code, Because of I use [Echo](https://github.com/labstack/echo), I Recommend it

```go
return c.String(http.StatusOK, "success")
```

### AliPayTradeWapPay

* The mobile website payment is obtained by the server after obtaining the payment URL, and then returned to the client, requesting the URL address to open the payment page.

> Document：[Wap Payment](https://docs.open.alipay.com/203/107090/) 

> Document：[alipay.trade.wap.pay2.0](https://docs.open.alipay.com/api_1/alipay.trade.wap.pay/) 

```go
privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
//Init Alipay Payment Client
//    appId：
//    privateKey：
//    isProd：
client := gopay.NewAliPayClient("2016091200494382", privateKey, false)
//set public parameters
client.SetCharset("utf-8").
	SetSignType("RSA2").
	//SetAppAuthToken("").
	//SetReturnUrl("https://www.gopay.ink").
	SetNotifyUrl("https://www.gopay.ink")
//BodyMap
body := make(gopay.BodyMap)
body.Set("subject", "测试支付")
body.Set("out_trade_no", "GYWX201901301040355706100409")
body.Set("quit_url", "https://www.gopay.ink")
body.Set("total_amount", "10.00")
body.Set("product_code", "QUICK_WAP_WAY")
//Request
payUrl, err := client.AliPayTradeWapPay(body)
if err != nil {
	log.Println("err:", err)
	return
}
fmt.Println("payUrl:", payUrl)
```

### AliPayTradeAppPay

* APP payment is obtained after the server obtains the payment parameters, and then calls the payment function through the SDK of the Android/iOS client.

> Document：[Parameter Description](https://docs.open.alipay.com/204/105465/) 

> Document：[alipay.trade.app.pay](https://docs.open.alipay.com/api_1/alipay.trade.app.pay/) 

```go
privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
//Init Alipay Payment Client
//    appId：
//    privateKey：
//    isProd：
client := gopay.NewAliPayClient("2016091200494382", privateKey, false)
//set public parameters
client.SetCharset("utf-8").
	SetSignType("RSA2").
	//SetAppAuthToken("").
	//SetReturnUrl("https://www.gopay.ink").
	SetNotifyUrl("https://www.gopay.ink")
//BodyMap
body := make(gopay.BodyMap)
body.Set("subject", "测试APP支付")
body.Set("out_trade_no", "GYWX201901301040355706100411")
body.Set("total_amount", "1.00")
//Request
payParam, err := client.AliPayTradeAppPay(body)
if err != nil {
	fmt.Println("err:", err)
	return
}
fmt.Println("payParam:", payParam)
```

### AliPayTradePagePay

* The computer website payment is obtained after the server obtains the payment URL, and then returns to the client, and requests the URL address to open the payment page.

> Document：[PC Web Payment](https://docs.open.alipay.com/270) 

> Document：[alipay.trade.page.pay](https://docs.open.alipay.com/api_1/alipay.trade.page.pay) 

```go
privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
//Init Alipay Payment Client
//    appId：
//    privateKey：
//    isProd：
client := gopay.NewAliPayClient("2016091200494382", privateKey, false)
//set public parameters
client.SetCharset("utf-8").
	SetSignType("RSA2").
	//SetAppAuthToken("").
	//SetReturnUrl("https://www.gopay.ink").
	SetNotifyUrl("https://www.gopay.ink")
//BodyMap
body := make(gopay.BodyMap)
body.Set("subject", "网站测试支付")
body.Set("out_trade_no", "GYWX201901301040355706100418")
body.Set("total_amount", "88.88")
body.Set("product_code", "FAST_INSTANT_TRADE_PAY")

//Request
payUrl, err := client.AliPayTradePagePay(body)
if err != nil {
	fmt.Println("err:", err)
	return
}
fmt.Println("payUrl:", payUrl)
```

### AliPayTradePay

* The merchant uses a bar code recognition device such as a scan code gun to scan the barcode/two-dimensional code on the user's Alipay wallet to complete the payment.

> Document：[Scan Payment](https://docs.open.alipay.com/194) 

> Document：[alipay.trade.pay](https://docs.open.alipay.com/api_1/alipay.trade.pay) 

```go
privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
//Init Alipay Payment Client
//    appId：
//    privateKey：
//    isProd：
client := gopay.NewAliPayClient("2016091200494382", privateKey, false)
//set public parameters
client.SetCharset("utf-8").
	SetSignType("RSA2").
	//SetAppAuthToken("").
	//SetReturnUrl("https://www.gopay.ink").
	SetNotifyUrl("https://www.gopay.ink")
//BodyMap
body := make(gopay.BodyMap)
body.Set("subject", "条码支付")
body.Set("scene", "bar_code")
body.Set("auth_code", "285860185283886370")
body.Set("out_trade_no", "GYWX201901301040355706100456")
body.Set("total_amount", "10.00")
body.Set("timeout_express", "2m")

//Request
aliRsp, err := client.AliPayTradePay(body)
if err != nil {
	fmt.Println("err:", err)
	return
}
fmt.Println("aliRsp:", *aliRsp)
```

### AliPayTradeRefund

* Trade order refund interface, please refer to the official document for specific conditions.

> Document：[alipay.trade.refund](https://docs.open.alipay.com/api_1/alipay.trade.refund) 

```go
privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
//Init Alipay Payment Client
//    appId：
//    privateKey：
//    isProd：
client := gopay.NewAliPayClient("2016091200494382", privateKey, false)
//set public parameters
client.SetCharset("utf-8").
	SetSignType("RSA2")
//BodyMap
body := make(gopay.BodyMap)
body.Set("out_trade_no", "GZ201907261437329516")
body.Set("refund_amount", "100.00")
body.Set("refund_reason", "测试支付退款")
//Request
aliRsp, err := client.AliPayTradeRefund(body)
if err != nil {
	fmt.Println("err:", err)
	return
}
fmt.Println("aliRsp:", *aliRsp)
```

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