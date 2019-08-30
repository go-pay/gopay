
<div align=center><img width="220" height="220" alt="Logo was Loading Faild!" src="https://raw.githubusercontent.com/iGoogle-ink/gopay/master/logo.png"/></div>

# GoPay

[English Document](https://github.com/iGoogle-ink/gopay/blob/master/README_EN.md)

微信和支付宝的Golang版本SDK

[![](https://img.shields.io/badge/Author-Jerry-blue.svg)](https://www.gopay.ink)
[![](https://img.shields.io/badge/Golang-1.11+-brightgreen.svg)](https://golang.org)
[![](https://img.shields.io/badge/Version-1.3.2-blue.svg)](https://github.com/iGoogle-ink/gopay)
[![](https://api.travis-ci.org/iGoogle-ink/gopay.svg?branch=master)]()
[![](https://img.shields.io/badge/License-Apache_2-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0)

### 微信Client

* client := gopay.NewWeChatClient() => 初始化微信支付客户端

### 微信支付API

* 统一下单：client.UnifiedOrder()
    * JSAPI - JSAPI支付（或小程序支付）
    * NATIVE - Native支付
    * APP - app支付
    * MWEB - H5支付
* 提交付款码支付：client.Micropay()
* 查询订单：client.QueryOrder()
* 关闭订单：client.CloseOrder()
* 撤销订单：client.Reverse()
* 申请退款：client.Refund()
* 查询退款：client.QueryRefund()
* 下载对账单：client.DownloadBill()
* 下载资金账单：client.DownloadFundFlow()
* 拉取订单评价数据：client.BatchQueryComment()
* 企业向微信用户个人付款：client.Transfer()

### 微信公共API

* gopay.GetWeChatParamSign() => 获取微信支付所需参数里的Sign值（通过支付参数计算Sign值）
* gopay.GetWeChatSanBoxParamSign() => 获取微信支付沙箱环境所需参数里的Sign值（通过支付参数计算Sign值）
* gopay.GetMiniPaySign() => 获取微信小程序支付所需要的paySign
* gopay.GetH5PaySign() => 获取微信内H5支付所需要的paySign
* gopay.GetAppPaySign() => 获取APP支付所需要的paySign
* gopay.ParseWeChatNotifyResultToBodyMap() => 解析微信支付异步通知的参数到BodyMap
* gopay.ParseWeChatNotifyResult() => 解析微信支付异步通知的参数
* gopay.VerifyWeChatSign() => 微信同步返回参数验签或异步通知参数验签
* gopay.Code2Session() => 登录凭证校验：获取微信用户OpenId、UnionId、SessionKey
* gopay.GetAccessToken() => 获取小程序全局唯一后台接口调用凭据
* gopay.GetPaidUnionId() => 微信小程序用户支付完成后，获取该用户的 UnionId，无需用户授权
* gopay.GetWeChatUserInfo() => 微信公众号：获取用户基本信息(UnionID机制)
* gopay.DecryptOpenDataToStruct() => 加密数据，解密到指定结构体
* gopay.GetOpenIdByAuthCode() => 授权码查询openid

---

### 支付宝Client

* client := gopay.NewAliPayClient() => 初始化支付宝支付客户端

### 支付宝支付API

* 手机网站支付接口2.0（手机网站支付）：client.AliPayTradeWapPay()
* 统一收单下单并支付页面接口（电脑网站支付）：client.AliPayTradePagePay()
* APP支付接口2.0（APP支付）：client.AliPayTradeAppPay()
* 统一收单交易支付接口（商家扫用户付款码）：client.AliPayTradePay()
* 统一收单交易创建接口（小程序支付）：client.AliPayTradeCreate()
* 统一收单线下交易查询：client.AliPayTradeQuery()
* 统一收单交易关闭接口：client.AliPayTradeClose()
* 统一收单交易撤销接口：client.AliPayTradeCancel()
* 统一收单交易退款接口：client.AliPayTradeRefund()
* 统一收单退款页面接口：client.AliPayTradePageRefund()
* 统一收单交易退款查询：client.AliPayTradeFastPayRefundQuery()
* 统一收单交易结算接口：client.AliPayTradeOrderSettle()
* 统一收单线下交易预创建（用户扫商品收款码）：client.AliPayTradePrecreate()
* 单笔转账到支付宝账户接口（商户给支付宝用户转账）：client.AlipayFundTransToaccountTransfer()
* 换取授权访问令牌（获取access_token，user_id等信息）：client.AliPaySystemOauthToken()
* 换取应用授权令牌（获取app_auth_token，auth_app_id，user_id等信息）：client.AlipayOpenAuthTokenApp()
* 获取芝麻信用分：client.ZhimaCreditScoreGet()

### 支付宝公共API

* gopay.AliPaySystemOauthToken() => 换取授权访问令牌（得到access_token，user_id等信息）
* gopay.FormatPrivateKey() => 格式化应用私钥
* gopay.FormatAliPayPublicKey() => 格式化支付宝公钥
* gopay.ParseAliPayNotifyResult() => 解析并返回支付宝支付异步通知的参数
* gopay.VerifyAliPaySign() => 支付宝同步返回参数验签或异步通知参数验签
* gopay.DecryptAliPayOpenDataToStruct() => 支付宝小程序敏感加密数据解析到结构体

# 一、安装

```bash
$ go get -u github.com/iGoogle-ink/gopay
```

* #### 查看 GoPay 版本

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

# 二、文档说明

* [GoDoc](https://godoc.org/github.com/iGoogle-ink/gopay)

* 有问题请加微信群。在此，非常感谢那些加群后，提出意见和反馈问题的同志们！另外，仅 Fork 的同志顺手点个星呗，您的支持给了我源源不断的动力

讨论群（如过期请加微信）：
<img width="226" height="300" alt="Photo was Loading Faild!" src="https://raw.githubusercontent.com/iGoogle-ink/gopay/master/wechat_gopay.png"/>
微信：
<img width="226" height="300" alt="Photo was Loading Faild!" src="https://raw.githubusercontent.com/iGoogle-ink/gopay/master/wechat_jerry.png"/>

---

## 1、初始化GoPay客户端并做配置

* #### 微信客户端

微信官方文档：[官方文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)
```go
//初始化微信客户端
//    appId：应用ID
//    mchId：商户ID
//    apiKey：API秘钥值
//    isProd：是否是正式环境
client := gopay.NewWeChatClient("wxdaa2ab9ef87b5497", mchId, apiKey, false)

//设置国家：不设置默认 中国国内
//    gopay.China：中国国内
//    gopay.China2：中国国内备用
//    gopay.SoutheastAsia：东南亚
//    gopay.Other：其他国家
client.SetCountry(gopay.China)
```

* #### 支付宝

支付宝官方文档：[官方文档](https://docs.open.alipay.com/catalog)

支付宝RSA秘钥生成文档：[生成 RSA 密钥](https://docs.open.alipay.com/291/105971/)

沙箱环境使用说明：[文档地址](https://docs.open.alipay.com/200/105311)

```go
//初始化支付宝客户端
//    appId：应用ID
//    privateKey：应用秘钥
//    isProd：是否是正式环境
client := gopay.NewAliPayClient("2016091200494382", privateKey, false)

//设置支付宝请求 公共参数
//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所以设置参数
client.SetCharset("utf-8").                 //设置字符编码，不设置默认 utf-8
	SetSignType("RSA2").                    //设置签名类型，不设置默认 RSA2
	SetReturnUrl("https://www.gopay.ink").  //设置返回URL
	SetNotifyUrl("https://www.gopay.ink").  //设置异步通知URL
	SetAppAuthToken("").                    //设置第三方应用授权
	SetAuthToken("")                        //设置个人信息授权
```

## 2、初始化并赋值BodyMap（client的方法所需的入参）

* #### 微信请求参数

具体参数请根据不同接口查看：[微信支付接口文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)
```go
//初始化 BodyMap
bm := make(gopay.BodyMap)
bm.Set("nonce_str", gopay.GetRandomString(32))
bm.Set("body", "小程序测试支付")
bm.Set("out_trade_no", number)
bm.Set("total_fee", 1)
bm.Set("spbill_create_ip", "127.0.0.1")
bm.Set("notify_url", "http://www.gopay.ink")
bm.Set("trade_type", gopay.TradeType_Mini)
bm.Set("device_info", "WEB")
bm.Set("sign_type", gopay.SignType_MD5)
bm.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")

//嵌套json格式数据（例如：H5支付的 scene_info 参数）
h5Info := make(map[string]string)
h5Info["type"] = "Wap"
h5Info["wap_url"] = "http://www.gopay.ink"
h5Info["wap_name"] = "H5测试支付"

sceneInfo := make(map[string]map[string]string)
sceneInfo["h5_info"] = h5Info

bm.Set("scene_info", sceneInfo)

//参数 sign ，可单独生成赋值到BodyMap中；也可不传sign参数，client内部会自动获取
//如需单独赋值 sign 参数，需通过下面方法，最后获取sign值并在最后赋值此参数
sign := gopay.GetWeChatParamSign("wxdaa2ab9ef87b5497", mchId, apiKey, body)
//sign, _ := gopay.GetWeChatSanBoxParamSign("wxdaa2ab9ef87b5497", mchId, apiKey, body)
bm.Set("sign", sign)
```

* #### 支付宝请求参数

具体参数请根据不同接口查看：[支付宝支付API接口文档](https://docs.open.alipay.com/api_1/alipay.trade.wap.pay)
```go
//此时李
//初始化 BodyMap
bm := make(gopay.BodyMap)
bm.Set("subject", "手机网站测试支付")
bm.Set("out_trade_no", "GZ201901301040355703")
bm.Set("quit_url", "https://www.gopay.ink")
bm.Set("total_amount", "100.00")
bm.Set("product_code", "QUICK_WAP_WAY")
```

## 3、client 方法调用

* #### 微信 client 
```go
wxRsp, err := client.UnifiedOrder(bm)
wxRsp, err := client.Micropay(bm)
wxRsp, err := client.QueryOrder(bm)
wxRsp, err := client.CloseOrder(bm)
wxRsp, err := client.Reverse(bm, "apiclient_cert.pem", "apiclient_key.pem", "apiclient_cert.p12")
wxRsp, err := client.Refund(bm, "apiclient_cert.pem", "apiclient_key.pem", "apiclient_cert.p12")
wxRsp, err := client.QueryRefund(bm)
wxRsp, err := client.DownloadBill(bm)
wxRsp, err := client.DownloadFundFlow(bm, "apiclient_cert.pem", "apiclient_key.pem", "apiclient_cert.p12")
wxRsp, err := client.BatchQueryComment(bm, "apiclient_cert.pem", "apiclient_key.pem", "apiclient_cert.p12")
wxRsp, err := client.Transfer(bm, "apiclient_cert.pem", "apiclient_key.pem", "apiclient_cert.p12")
```

* #### 支付宝 client
```go
//手机网站支付是通过服务端获取支付URL后，然后返回给客户端，请求URL地址即可打开支付页面
payUrl, err := client.AliPayTradeWapPay(bm)

//电脑网站支付是通过服务端获取支付URL后，然后返回给客户端，请求URL地址即可打开支付页面
payUrl, err := client.AliPayTradePagePay(bm)

//APP支付是通过服务端获取支付参数后，然后通过Android/iOS客户端的SDK调用支付功能
payParam, err := client.AliPayTradeAppPay(bm)

//商家使用扫码枪等条码识别设备扫描用户支付宝钱包上的条码/二维码，完成收款
aliRsp, err := client.AliPayTradePay(bm)

//支付宝小程序支付时 buyer_id 为必传参数，需要提前获取，获取方法如下两种
//    1、gopay.AliPaySystemOauthToken()     返回取值：rsp.AliPaySystemOauthTokenResponse.UserId
//    2、client.AliPaySystemOauthToken()    返回取值：aliRsp.AliPaySystemOauthTokenResponse.UserId
aliRsp, err := client.AliPayTradeCreate(bm)

aliRsp, err := client.AliPayTradeQuery(bm)
aliRsp, err := client.AliPayTradeClose(bm)
aliRsp, err := client.AliPayTradeCancel(bm)
aliRsp, err := client.AliPayTradeRefund(bm)
aliRsp, err := client.AliPayTradePageRefund(bm)
aliRsp, err := client.AliPayTradeFastPayRefundQuery(bm)
aliRsp, err := client.AliPayTradeOrderSettle(bm)
aliRsp, err := client.AliPayTradePrecreate(bm)
aliRsp, err := client.AlipayFundTransToaccountTransfer(bm)
aliRsp, err := client.AliPaySystemOauthToken(bm)
aliRsp, err := client.AlipayOpenAuthTokenApp(bm)
aliRsp, err := client.ZhimaCreditScoreGet(bm)
```

## 4、微信统一下单后，获取微信小程序支付、APP支付、微信内H5支付所需要的 paySign

* #### 微信（只有微信需要此操作）
微信小程序支付官方文档：[微信小程序支付API](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/payment/wx.requestPayment.html)

APP支付官方文档：[APP端调起支付的参数列表文档](https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_12)

微信内H5支付官方文档：[微信内H5支付文档](https://pay.weixin.qq.com/wiki/doc/api/external/jsapi.php?chapter=7_7&index=6)
```go
//====微信小程序 paySign====
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
prepayId := "prepay_id=" + wxRsp.PrepayId   //此处的 wxRsp.PrepayId ,统一下单成功后得到
//获取微信小程序支付的 paySign
//    appId：APPID
//    nonceStr：随即字符串
//    prepayId：统一下单成功后得到的值
//    signType：签名方式，务必与统一下单时用的签名方式一致
//    timeStamp：时间
//    apiKey：API秘钥值
paySign := gopay.GetMiniPaySign(AppID, wxRsp.NonceStr, prepayId, gopay.SignType_MD5, timeStamp, apiKey)

//====APP支付 paySign====
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
//获取APP支付的 paySign
//注意：package 参数因为是固定值，无需开发者再传入
//    appId：APPID
//    partnerid：partnerid
//    nonceStr：随即字符串
//    prepayId：统一下单成功后得到的值
//    signType：签名方式，务必与统一下单时用的签名方式一致
//    timeStamp：时间
//    apiKey：API秘钥值
paySign := gopay.GetAppPaySign(appid, partnerid, wxRsp.NonceStr, wxRsp.PrepayId, gopay.SignType_MD5, timeStamp, apiKey)

//====微信内H5支付 paySign====
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
packages := "prepay_id=" + wxRsp.PrepayId   //此处的 wxRsp.PrepayId ,统一下单成功后得到
//获取微信内H5支付 paySign
//    appId：APPID
//    nonceStr：随即字符串
//    packages：统一下单成功后拼接得到的值
//    signType：签名方式，务必与统一下单时用的签名方式一致
//    timeStamp：时间
//    apiKey：API秘钥值
paySign := gopay.GetMiniPaySign(AppID, wxRsp.NonceStr, packages, gopay.SignType_MD5, timeStamp, apiKey)
```

## 5、同步返回参数验签Sign、异步通知参数解析和验签Sign、异步通知返回

异步参数需要先解析，解析出来的结构体或BodyMap再验签

[Echo Web框架](https://github.com/labstack/echo)，有兴趣的可以尝试一下

异步通知处理完后，需回复平台固定数据

* #### 微信
```go
//====同步返回参数验签Sign====
wxRsp, err := client.UnifiedOrder(bm)
//微信同步返回参数验签或异步通知参数验签
//    apiKey：API秘钥值
//    signType：签名类型（调用API方法时填写的类型）
//    bean：微信同步返回的结构体 wxRsp 或 异步通知解析的结构体 notifyReq
//    返回参数 ok：是否验签通过
//    返回参数 err：错误信息
ok, err := gopay.VerifyWeChatSign(apiKey, gopay.SignType_MD5, wxRsp)

//====异步通知参数解析和验签Sign====
//解析异步通知的参数
//    req：*http.Request
//    返回参数 notifyReq：通知的参数
//    返回参数 err：错误信息
notifyReq, err := gopay.ParseWeChatNotifyResult(c.Request())    //c.Request()是 echo 框架的获取 *http.Request 的写法
//验签操作
ok, err := gopay.VerifyWeChatSign(apiKey, gopay.SignType_MD5, notifyReq)

//==异步通知，返回给微信平台的信息==
rsp := new(gopay.WeChatNotifyResponse) //回复微信的数据
rsp.ReturnCode = gopay.SUCCESS
rsp.ReturnMsg = gopay.OK
return c.String(http.StatusOK, rsp.ToXmlString())   //此写法是 echo 框架返回客户端数据的写法
```

* #### 支付宝

支付宝的**同步返回**验签，参数请注意看注释

APP支付，手机网站支付，电脑网站支付 暂不支持同步返回验签

支付宝支付后的同步/异步通知验签文档：[支付结果通知](https://docs.open.alipay.com/200/106120)
```go
//====同步返回参数验签Sign====
aliRsp, err := client.AliPayTradePay(bm)
//支付宝同步返回验签或异步通知验签
//    注意：APP支付，手机网站支付，电脑网站支付 暂不支持同步返回验签
//    aliPayPublicKey：支付宝公钥
//    bean： 同步返回验签时，此参数为 aliRsp.SignData ；异步通知验签时，此参数为异步通知解析的结构体 notifyReq
//    syncSign：同步返回验签时，此参数必传，即：aliRsp.Sign ；异步通知验签时，不要传此参数，否则会出错。
//    返回参数ok：是否验签通过
//    返回参数err：错误信息
ok, err := gopay.VerifyAliPaySign(alipayPublicKey, aliRsp.SignData, aliRsp.Sign)

//====异步通知参数解析和验签Sign====
//解析异步通知的参数
//    req：*http.Request
//    返回参数 notifyReq：通知的参数
//    返回参数 err：错误信息
notifyReq, err = gopay.ParseAliPayNotifyResult(c.Request())     //c.Request()是 echo 框架的获取
//验签操作
ok, err = gopay.VerifyAliPaySign(alipayPublicKey, notifyReq)

//==异步通知，返回支付宝平台的信息==
//    文档：https://docs.open.alipay.com/203/105286
//    程序执行完后必须打印输出“success”（不包含引号）。如果商户反馈给支付宝的字符不是success这7个字符，支付宝服务器会不断重发通知，直到超过24小时22分钟。一般情况下，25小时以内完成8次通知（通知的间隔频率一般是：4m,10m,10m,1h,2h,6h,15h）
return c.String(http.StatusOK, "success")   //此写法是 echo 框架返回客户端数据的写法
```

## 6、微信、支付宝 公共API（仅部分说明）

* #### 微信 公共API

官方文档：[code2Session](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html)

button按钮获取手机号码：[button组件文档](https://developers.weixin.qq.com/miniprogram/dev/component/button.html)

微信解密算法文档：[解密算法文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)
```go
//获取微信小程序用户的OpenId、SessionKey、UnionId
//    appId：微信小程序的APPID
//    appSecret：微信小程序的AppSecret
//    wxCode：小程序调用wx.login 获取的code
sessionRsp, err := gopay.Code2Session(appId, appSecret, wxCode)

//解密微信加密数据到指定结构体
//    以小程序获取手机号为例
phone := new(gopay.WeChatUserPhone)
//解密开放数据
//    encryptedData：包括敏感数据在内的完整用户信息的加密数据，小程序获取到
//    iv：加密算法的初始向量，小程序获取到
//    sessionKey：会话密钥，通过 gopay.Code2Session() 方法获取到
//    beanPtr：需要解析到的结构体指针，操作完后，声明的结构体会被赋值
err := gopay.DecryptOpenDataToStruct(encryptedData, iv, sessionKey, phone)
fmt.Println(*phone)
```

* #### 支付宝 公共API

支付宝换取授权访问令牌文档：[换取授权访问令牌](https://docs.open.alipay.com/api_9/alipay.system.oauth.token)

获取用户手机号文档：[获取用户手机号](https://docs.alipay.com/mini/api/getphonenumber)

支付宝加解密文档：[AES配置文档](https://docs.alipay.com/mini/introduce/aes)，[AES加解密文档](https://docs.open.alipay.com/common/104567)
```go
//换取授权访问令牌（默认使用utf-8，RSA2）
//    appId：应用ID
//    privateKey：应用私钥
//    grantType：值为 authorization_code 时，代表用code换取；值为 refresh_token 时，代表用refresh_token换取，传空默认code换取
//    codeOrToken：支付宝授权码或refresh_token
rsp, err := gopay.AlipaySystemOauthToken(appId, privateKey, grantType, codeOrToken)

//解密支付宝开放数据带到指定结构体
//    以小程序获取手机号为例
phone := new(gopay.AliPayUserPhone)
//解密支付宝开放数据
//    encryptedData:包括敏感数据在内的完整用户信息的加密数据
//    secretKey:AES密钥，支付宝管理平台配置
//    beanPtr:需要解析到的结构体指针
err := gopay.DecryptAliPayOpenDataToStruct(encryptedData, secretKey, phone)
fmt.Println(*phone)
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