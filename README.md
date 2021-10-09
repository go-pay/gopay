<div align=center><img width="240" height="240" alt="Logo was Loading Faild!" src="https://raw.githubusercontent.com/go-pay/gopay/main/logo.png"/></div>

# GoPay

#### 微信、支付宝、PayPal、QQ 的Golang版本SDK

[![Github](https://img.shields.io/github/followers/iGoogle-ink?label=Follow&style=social)](https://github.com/iGoogle-ink)
[![Github](https://img.shields.io/github/forks/go-pay/gopay?label=Fork&style=social)](https://github.com/go-pay/gopay/fork)

[![Golang](https://img.shields.io/badge/golang-1.16-brightgreen.svg)](https://golang.google.cn)
[![GoDoc](https://img.shields.io/badge/doc-go.dev-informational.svg)](https://pkg.go.dev/github.com/go-pay/gopay)
[![Drone CI](https://cloud.drone.io/api/badges/go-pay/gopay/status.svg)](https://cloud.drone.io/go-pay/gopay)
[![GitHub Release](https://img.shields.io/github/v/release/go-pay/gopay)](https://github.com/go-pay/gopay/releases)
[![License](https://img.shields.io/github/license/go-pay/gopay)](https://www.apache.org/licenses/LICENSE-2.0)

---

# 一、安装

- v1.5.42 开始，仓库从 `github.com/iGoogle-ink/gopay` 迁移到 `github.com/go-pay/gopay`
```bash
go get github.com/go-pay/gopay
```

#### 查看 GoPay 版本

  [版本更新记录](https://github.com/go-pay/gopay/blob/main/release_note.txt)

```go
import (
    "github.com/go-pay/gopay"
    "github.com/go-pay/gopay/pkg/xlog"
)

func main() {
    xlog.Debug("GoPay Version: ", gopay.Version)
}
```

---

# 二、API 列表

* [GoPay 已实现 API 文档链接地址](https://github.com/go-pay/gopay/tree/main/doc)
    - `alipay` [点击这里跳转](https://github.com/go-pay/gopay/tree/main/doc/alipay)
    - `wechat` [点击这里跳转](https://github.com/go-pay/gopay/tree/main/doc/wechat)
    - `qq` [点击这里跳转](https://github.com/go-pay/gopay/tree/main/doc/qq)
    - `paypal` [点击这里跳转](https://github.com/go-pay/gopay/tree/main/doc/paypal)
    - `apple` [点击这里跳转](https://github.com/go-pay/gopay/tree/main/doc/apple)

---

# 三、文档说明
* 使用示例请参考[example](https://github.com/go-pay/gopay/tree/main/examples]) 或者各类单元测试`_test.go`文件
* [GoPay 文档地址](https://pkg.go.dev/github.com/go-pay/gopay)
* 所有方法，如有问题，请仔细查看 `wechat/client_test.go`、`alipay/client_test.go` 或 examples
* 有问题请加QQ群（加群验证答案：gopay），微信加好友拉群。在此，非常感谢那些加群后，提出意见和反馈问题的同志们！

QQ群：
<img width="280" height="280" src="https://raw.githubusercontent.com/go-pay/gopay/main/qq_gopay.png"/>
加微信拉群：
<img width="280" height="280" src="https://raw.githubusercontent.com/go-pay/gopay/main/wechat_jerry.png"/>

---

## 1、初始化GoPay客户端并做配置（HTTP请求均默认设置tls.Config{InsecureSkipVerify: true}）

* #### 微信V3

> 注意：V3 版本接口持续增加中，不支持沙箱支付，测试请用1分钱测试法

> 注意：`微信平台证书` 和 `微信平台证书序列号`，请自行通过 `wechat.GetPlatformCerts()` 方法维护

> 具体使用介绍，请参考 `gopay/wechat/v3/client_test.go`

```go
import (
    "github.com/go-pay/gopay/pkg/xlog"
    "github.com/go-pay/gopay/wechat/v3"
)

// NewClientV3 初始化微信客户端 V3
//	mchid：商户ID 或者服务商模式的 sp_mchid
// 	serialNo：商户证书的证书序列号
//	apiV3Key：apiV3Key，商户平台获取
//	privateKey：私钥 apiclient_key.pem 读取后的内容
client, err = wechat.NewClientV3(MchId, SerialNo, APIv3Key, PrivateKey)
if err != nil {
    xlog.Error(err)
    return
}

// 设置微信平台证书和序列号，并启用自动同步返回验签
//	注意：请预先通过 wechat.GetPlatformCerts() 获取并维护微信平台证书和证书序列号
client.SetPlatformCert([]byte(WxPkContent), WxPkSerialNo).AutoVerifySign()

// 打开Debug开关，输出日志，默认是关闭的
client.DebugSwitch = gopay.DebugOn
```

* #### 支付宝

支付宝官方文档：[官方文档](https://openhome.alipay.com/docCenter/docCenter.htm)

支付宝RSA秘钥生成文档：[生成RSA密钥](https://opendocs.alipay.com/open/291/105971) （推荐使用 RSA2）

技术支持 & 案例 FAQ：[秘钥问题](https://opendocs.alipay.com/support/01rauw)

沙箱环境使用说明：[文档地址](https://opendocs.alipay.com/open/200/105311)

```go
import (
    "github.com/go-pay/gopay/alipay"
    "github.com/go-pay/gopay/pkg/xlog"
)

// 初始化支付宝客户端
//    appId：应用ID
//    privateKey：应用私钥，支持PKCS1和PKCS8
//    isProd：是否是正式环境
client, err := alipay.NewClient("2016091200494382", privateKey, false)
if err != nil {
    xlog.Error(err)
    return
}
// 打开Debug开关，输出日志，默认关闭
client.DebugSwitch = gopay.DebugOn

// 设置支付宝请求 公共参数
//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
client.SetLocation(alipay.LocationShanghai).    // 设置时区，不设置或出错均为默认服务器时间
    SetCharset(alipay.UTF8).                    // 设置字符编码，不设置默认 utf-8
    SetSignType(alipay.RSA2).                   // 设置签名类型，不设置默认 RSA2
    SetReturnUrl("https://www.fmm.ink").        // 设置返回URL
    SetNotifyUrl("https://www.fmm.ink").        // 设置异步通知URL
    SetAppAuthToken()                           // 设置第三方应用授权

// 自动同步验签（只支持证书模式）
// 传入 alipayCertPublicKey_RSA2.crt 内容
client.AutoVerifySign([]byte("alipayCertPublicKey_RSA2 bytes"))

// 公钥证书模式，需要传入证书，以下两种方式二选一
// 证书路径
err := client.SetCertSnByPath("appCertPublicKey.crt", "alipayRootCert.crt", "alipayCertPublicKey_RSA2.crt")
// 证书内容
err := client.SetCertSnByContent("appCertPublicKey bytes", "alipayRootCert bytes", "alipayCertPublicKey_RSA2 bytes")
```

* #### PayPal 支付

PayPal官方文档：[官方文档](https://developer.paypal.com/docs/api/overview)

> 具体API使用介绍，请参考 `gopay/paypal/client_test.go`

```go
import (
    "github.com/go-pay/gopay/paypal"
    "github.com/go-pay/gopay/pkg/xlog"
)

// 初始化PayPal支付客户端
client, err := paypal.NewClient(Clientid, Secret, false)
if err != nil {
    xlog.Error(err)
    return
}
// 打开Debug开关，输出日志，默认关闭
client.DebugSwitch = gopay.DebugOn
```

## 2、初始化并赋值BodyMap（client的方法所需的入参）

* #### 微信请求参数

具体参数请根据不同接口查看：[微信支付V3的API字典概览](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/index.shtml)

```go
import (
    "github.com/go-pay/gopay"
)

// JSAPI下单 示例
expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
// 初始化 BodyMap
bm := make(gopay.BodyMap)
bm.Set("description", "测试Jsapi支付商品").
    Set("out_trade_no", tradeNo).
    Set("time_expire", expire).
    Set("notify_url", "https://www.fmm.ink").
    SetBodyMap("amount", func(bm gopay.BodyMap) {
        bm.Set("total", 1).
            Set("currency", "CNY")
    }).
    SetBodyMap("payer", func(bm gopay.BodyMap) {
        bm.Set("openid", "asdas")
    })
```

* #### 支付宝请求参数

具体参数请根据不同接口查看：[支付宝支付API接口文档](https://opendocs.alipay.com/apis)

```go
import (
    "github.com/go-pay/gopay"
)

// 统一收单交易支付接口 示例
// 初始化 BodyMap
bm := make(gopay.BodyMap)
bm.Set("subject", "条码支付").
    Set("scene", "bar_code").
    Set("auth_code", "286248566432274952").
    Set("out_trade_no", "GZ201909081743431443").
    Set("total_amount", "0.01").
    Set("timeout_express", "2m")
```

## 3、client 方法调用

* #### 微信V3

```go
// 公共方法
client.SetPlatformCert()
client.AutoVerifySign()
client.V3EncryptText()
client.V3DecryptText()

// 直连商户
wxRsp, err := client.V3TransactionApp()
wxRsp, err := client.V3TransactionJsapi()
wxRsp, err := client.V3TransactionNative()

// 服务商
wxRsp, err := client.V3PartnerTransactionApp()
wxRsp, err := client.V3PartnerTransactionJsapi()

// 合单
wxRsp, err := client.V3CombineTransactionApp()
wxRsp, err := client.V3CombineTransactionJsapi()
...
```

* #### 支付宝

```go
// 手机网站支付是通过服务端获取支付URL后，然后返回给客户端，请求URL地址即可打开支付页面
payUrl, err := client.TradeWapPay()

// 电脑网站支付是通过服务端获取支付URL后，然后返回给客户端，请求URL地址即可打开支付页面
payUrl, err := client.TradePagePay()

// APP支付是通过服务端获取支付参数后，然后通过Android/iOS客户端的SDK调用支付功能
payParam, err := client.TradeAppPay()

// 商家使用扫码枪等条码识别设备扫描用户支付宝钱包上的条码/二维码，完成收款
aliRsp, err := client.TradePay()

// 支付宝小程序支付时 buyer_id 为必传参数，需要提前获取，获取方法如下两种
//  1、alipay.SystemOauthToken()     返回取值：aliRsp.SystemOauthTokenResponse.UserId
//  2、client.SystemOauthToken()     返回取值：aliRsp.SystemOauthTokenResponse.UserId
aliRsp, err := client.TradeCreate()
aliRsp, err := client.TradeQuery()
...
```

## 4、微信下单后，获取微信小程序支付、APP支付、JSAPI支付所需要的 pay sign

* #### 微信V3

小程序调起支付API：[小程序调起支付API](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_5_4.shtml)

APP调起支付API：[APP调起支付API](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_2_4.shtml)

JSAPI调起支付API：[JSAPI调起支付API](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_4.shtml)

H5调起支付API：[H5调起支付API](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_3_4.shtml)

```go
// jsapi
jsapi, err := client.PaySignOfJSAPI("prepayid")
// app
app, err := client.PaySignOfApp("prepayid")
// 小程序
applet, err := client.PaySignOfApplet("prepayid")
```

## 5、同步返回参数验签Sign、异步通知参数解析和验签Sign、异步通知返回

> 异步通知请求参数需要先解析，解析出来的结构体或BodyMap再验签（此处需要注意，`http.Request.Body` 只能解析一次，如果需要解析前调试，请处理好Body复用问题）

[Gin Web框架（推荐）](https://github.com/gin-gonic/gin)

[Echo Web框架](https://github.com/labstack/echo)

* #### 微信V3

```go
import (
    "github.com/go-pay/gopay/wechat/v3"
    "github.com/go-pay/gopay/pkg/xlog"
)

// ========同步返回 手动验签（如已开启自动验签，则无需手动验签操作）========
wxRsp, err := client.V3TransactionJsapi(bm)
if err != nil {
    xlog.Error(err)
    return
}
err = wechat.V3VerifySign(wxRsp.SignInfo.HeaderTimestamp, wxRsp.SignInfo.HeaderNonce, wxRsp.SignInfo.SignBody, wxRsp.SignInfo.HeaderSignature, WxPkContent)
if err != nil {
    xlog.Error(err)
    return
}

// ========异步通知验签========
notifyReq, err := wechat.V3ParseNotify()
if err != nil {
    xlog.Error(err)
    return
}
// WxPkContent 是通过 wechat.GetPlatformCerts() 接口向微信获取的微信平台公钥证书内容
err = notifyReq.VerifySign(WxPkContent)
if err != nil {
    xlog.Error(err)
    return
}

// ========异步通知敏感信息解密========
// 普通支付通知解密
result, err := notifyReq.DecryptCipherText(apiV3Key)
// 合单支付通知解密
result, err := notifyReq.DecryptCombineCipherText(apiV3Key)
// 退款通知解密
result, err := notifyReq.DecryptRefundCipherText(apiV3Key)

// ========异步通知应答========
// 退款通知http应答码为200且返回状态码为SUCCESS才会当做商户接收成功，否则会重试。
// 注意：重试过多会导致微信支付端积压过多通知而堵塞，影响其他正常通知。

// 此写法是 gin 框架返回微信的写法
c.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})
// 此写法是 echo 框架返回微信的写法
return c.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})
```

* #### 支付宝

注意：APP支付、手机网站支付、电脑网站支付 不支持同步返回验签

支付宝支付后的同步/异步通知验签文档：[支付结果通知](https://opendocs.alipay.com/open/200/106120)

```go
import (
    "github.com/go-pay/gopay/alipay"
)

// ====同步返回参数 手动验签（如已开启自动验签，则无需手动验签操作）====
aliRsp, err := client.TradePay(bm)
// 公钥模式验签
//    注意：APP支付，手机网站支付，电脑网站支付 不支持同步返回验签
//    aliPayPublicKey：支付宝平台获取的支付宝公钥
//    signData：待验签参数，aliRsp.SignData
//    sign：待验签sign，aliRsp.Sign
ok, err := alipay.VerifySyncSign(aliPayPublicKey, aliRsp.SignData, aliRsp.Sign)
// 公钥证书模式验签
//    aliPayPublicKeyCert：支付宝公钥证书存放路径 alipayCertPublicKey_RSA2.crt 或文件内容[]byte
//    signData：待验签参数，aliRsp.SignData
//    sign：待验签sign，aliRsp.Sign
ok, err := alipay.VerifySyncSignWithCert(aliPayPublicKeyCert, aliRsp.SignData, aliRsp.Sign)

// ====异步通知参数解析和验签Sign====
// 解析异步通知的参数
//    req：*http.Request
notifyReq, err = alipay.ParseNotifyToBodyMap(c.Request)     // c.Request 是 gin 框架的写法
 或
//    value：url.Values
notifyReq, err = alipay.ParseNotifyByURLValues()

// 支付宝异步通知验签（公钥模式）
ok, err = alipay.VerifySign(aliPayPublicKey, notifyReq)
// 支付宝异步通知验签（公钥证书模式）
ok, err = alipay.VerifySignWithCert("alipayCertPublicKey_RSA2.crt content", notifyReq)

// ====异步通知，返回支付宝平台的信息====
//    文档：https://opendocs.alipay.com/open/203/105286
//    程序执行完后必须打印输出“success”（不包含引号）。如果商户反馈给支付宝的字符不是success这7个字符，支付宝服务器会不断重发通知，直到超过24小时22分钟。一般情况下，25小时以内完成8次通知（通知的间隔频率一般是：4m,10m,10m,1h,2h,6h,15h）

// 此写法是 gin 框架返回支付宝的写法
c.String(http.StatusOK, "%s", "success")
// 此写法是 echo 框架返回支付宝的写法
return c.String(http.StatusOK, "success")
```

## 6、微信、支付宝 公共API（仅部分说明）

* #### 微信V3 公共API

> 微信敏感信息加解密、回调接口敏感信息解密

```go
import (
    "github.com/go-pay/gopay/wechat/v3"
)

// 获取微信平台证书和序列号信息
wechat.GetPlatformCerts()
// 请求参数 敏感信息加密
wechat.V3EncryptText() 或 client.V3EncryptText()
// 返回参数 敏感信息解密
wechat.V3DecryptText() 或 client.V3DecryptText()
// 回调通知敏感信息解密
wechat.V3DecryptNotifyCipherText()
wechat.V3DecryptRefundNotifyCipherText()
wechat.V3DecryptCombineNotifyCipherText()
...
```

* #### 支付宝 公共API

支付宝换取授权访问令牌文档：[换取授权访问令牌](https://opendocs.alipay.com/apis/api_9/alipay.system.oauth.token)

小程序获取用户手机号文档：[获取用户手机号](https://opendocs.alipay.com/mini/api/getphonenumber)

支付宝加解密文档：[AES配置文档](https://opendocs.alipay.com/mini/introduce/aes) ，[AES加解密文档](https://opendocs.alipay.com/open/common/104567)

```go
import (
    "github.com/go-pay/gopay/alipay"
    "github.com/go-pay/gopay/pkg/xlog"
)

// 换取授权访问令牌（默认使用utf-8，RSA2）
//    appId：应用ID
//    privateKey：应用私钥，支持PKCS1和PKCS8
//    grantType：值为 authorization_code 时，代表用code换取；值为 refresh_token 时，代表用refresh_token换取，传空默认code换取
//    codeOrToken：支付宝授权码或refresh_token
rsp, err := alipay.SystemOauthToken(appId, privateKey, grantType, codeOrToken)

// 解密支付宝开放数据带到指定结构体
//    以小程序获取手机号为例
phone := new(alipay.UserPhone)
// 解密支付宝开放数据
//    encryptedData:包括敏感数据在内的完整用户信息的加密数据
//    secretKey:AES密钥，支付宝管理平台配置
//    beanPtr:需要解析到的结构体指针
err := alipay.DecryptOpenDataToStruct(encryptedData, secretKey, phone)
xlog.Debugf("%+v", phone)
```

## 赞赏多少是您的心意，感谢！

<font color='#07C160' size='4'>微信：</font>
<img width="200" height="200" src="https://raw.githubusercontent.com/go-pay/gopay/main/zanshang_wx.png"/>
<font color='#027AFF' size='4'>支付宝：</font>
<img width="200" height="200" src="https://raw.githubusercontent.com/go-pay/gopay/main/zanshang_zfb.png"/>

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