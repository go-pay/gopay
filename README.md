
<div align=center><img width="240" height="240" alt="Logo was Loading Faild!" src="https://raw.githubusercontent.com/iGoogle-ink/gopay/master/logo.png"/></div>

# GoPay

<a href="https://www.igoogle.ink" target="_blank"><img src="https://img.shields.io/badge/Author-Jerry-blue.svg"/></a>
<a href="https://golang.org" target="_blank"><img src="https://img.shields.io/badge/Golang-1.11+-brightgreen.svg"/></a>
<img src="https://img.shields.io/badge/Build-passing-brightgreen.svg"/>
<a href="http://www.apache.org/licenses/LICENSE-2.0" target="_blank"><img src="https://img.shields.io/badge/License-Apache 2-blue.svg"/></a>

## 微信支付
* 统一下单：gopay.UnifiedOrder()
    * JSAPI - JSAPI支付（或小程序支付）
    * NATIVE - Native支付
    * APP - app支付
    * MWEB - H5支付
* 提交付款码支付：gopay.Micropay()
* 查询订单：gopay.QueryOrder()
* 关闭订单：gopay.CloseOrder()
* 撤销订单：gopay.Reverse()
* 申请退款：gopay.Refund()
* 查询退款：gopay.QueryRefund()
* 下载对账单：gopay.DownloadBill()
* 下载资金账单：gopay.DownloadFundFlow()
* 拉取订单评价数据：gopay.BatchQueryComment()

## 公共API
* gopay.ParseNotifyResult() => 解析并返回微信支付通知的参数
* gopay.VerifyPayResultSign() => 支付通知的签名验证和参数签名后的Sign
* gopay.Code2Session() => 登录凭证校验：获取微信用户OpenId、UnionId、SessionKey
* gopay.GetAccessToken() => 获取小程序全局唯一后台接口调用凭据
* gopay.GetPaidUnionId() => 用户支付完成后，获取该用户的 UnionId，无需用户授权
* gopay.GetWeChatUserInfo() => 微信公众号：获取用户基本信息(UnionID机制)
* gopay.DecryptOpenDataToStruct() => 加密数据，解密到指定结构体

## 安装

```bash
$ go get -u github.com/iGoogle-ink/gopay
```

## 文档

* 未完成，有问题+微信

微信：
<img width="300" height="300" src="https://raw.githubusercontent.com/iGoogle-ink/gopay/master/wechat_jerry.png"/>

## 微信支付

<font color='#0088ff'>注意：具体参数根据请求的不同而不同，请参考微信官方文档的参数说明！</font>

参考文档：[微信支付文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)

### 获取微信用户OpenId、UnionId、SessionKey

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

### 微信小程序支付，需要进一步获取微信小程序支付所需要的参数

* 小程序支付所需要的参数，paySign由后端计算
    * timeStamp
    * nonceStr
    * package 
    * signType
    * paySign

> 官方文档说明[微信小程序支付API](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/payment/wx.requestPayment.html)
```go
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
packages := "prepay_id=" + wxRsp.PrepayId   //此处的 wxRsp.PrepayId ,统一下单成功后得到
paySign := gopay.GetMiniPaySign("wxd678efh567hg6787", wxRsp.NonceStr, packages, gopay.SignType_MD5, timeStamp, "192006250b4c09247ec02edce69f6a2d")

//微信小程序支付需要的参数信息
payRsp := new(vm.WeChatPayRsp)
fmt.Println("timeStamp：", timeStamp)
fmt.Println("nonceStr：", wxRsp.NonceStr)
fmt.Println("package：", packages)
fmt.Println("signType：", gopay.SignType_MD5)
fmt.Println("paySign：", paySign)
```

### 微信内H5支付，同样需要进一步获取支付所需要的参数（与微信小程序支付类似）

* 微信内H5支付所需要的参数，paySign由后端计算
    * appId
    * timeStamp
    * nonceStr
    * package 
    * signType
    * paySign
> 官方文档说明[微信内H5支付文档](https://pay.weixin.qq.com/wiki/doc/api/external/jsapi.php?chapter=7_7&index=6)
```go
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
packages := "prepay_id=" + wxRsp.PrepayId   //此处的 wxRsp.PrepayId ,统一下单成功后得到
paySign := gopay.GetH5PaySign("wxd678efh567hg6787", wxRsp.NonceStr, packages, gopay.SignType_MD5, timeStamp, "192006250b4c09247ec02edce69f6a2d")

//微信内H5支付需要的参数信息
payRsp := new(vm.WeChatPayRsp)
fmt.Println("appId:","wxd678efh567hg6787")
fmt.Println("timeStamp：", timeStamp)
fmt.Println("nonceStr：", wxRsp.NonceStr)
fmt.Println("package：", packages)
fmt.Println("signType：", gopay.SignType_MD5)
fmt.Println("paySign：", paySign)
```

### 支付结果通知回调：参数解析和Sign值的验证

> 微信支付后的回调通知文档[支付结果通知](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_7&index=8)

```go
//解析支付完成后的回调信息
notifyRsp, err := gopay.ParseNotifyResult(c.Request())
if err != nil {
    fmt.Println("err:", err)
}
fmt.Println("notifyRsp:", notifyRsp)

//支付通知的签名验证和参数签名后的Sign
//    apiKey：API秘钥值
//    signType：签名类型 MD5 或 HMAC-SHA256（默认请填写 MD5）
//    notifyRsp：利用 gopay.ParseNotifyResult() 得到的结构体
//    返回参数ok：是否验证通过
//    返回参数sign：根据参数计算的sign值，非微信返回参数中的Sign
ok, sign := gopay.VerifyPayResultSign("192006250b4c09247ec02edce69f6a2d", "MD5", notifyRsp)
log.Println("ok:", ok)
log.Println("sign:", sign)
```


### 加密数据，解密到指定结构体

> 拿小程序获取手机号为例

button按钮获取手机号码:[button组件文档](https://developers.weixin.qq.com/miniprogram/dev/component/button.html)

微信解密算法文档:[解密算法文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)
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

### 付款结果回调,需回复微信平台是否成功

> 代码中return写法，由于本人用的[Echo Web框架](https://github.com/labstack/echo)，有兴趣的可以尝试一下
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
//    apiKey：API秘钥值
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
//    apiKey：API秘钥值
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
//    apiKey：API秘钥值
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
//    apiKey：API秘钥值
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