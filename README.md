
<div align=center><img width="220" height="220" alt="Logo was Loading Faild!" src="https://raw.githubusercontent.com/iGoogle-ink/gopay/master/logo.png"/></div>

# GoPay

#### QQ、微信、支付宝的Golang版本SDK

[![Golang](https://img.shields.io/badge/golang-1.14+-brightgreen.svg)](https://golang.google.cn)
[![GoDoc](https://img.shields.io/badge/doc-go.dev-informational.svg)](https://pkg.go.dev/github.com/iGoogle-ink/gopay)
[![Drone CI](https://cloud.drone.io/api/badges/iGoogle-ink/gopay/status.svg)](https://cloud.drone.io/iGoogle-ink/gopay)
[![GitHub Release](https://img.shields.io/github/v/release/iGoogle-ink/gopay)](https://github.com/iGoogle-ink/gopay/releases)
[![License](https://img.shields.io/github/license/iGoogle-ink/gopay)](https://www.apache.org/licenses/LICENSE-2.0)

# 一、安装

```bash
$ go get github.com/iGoogle-ink/gopay
```

* #### 查看 GoPay 版本
    * [版本更新记录](https://github.com/iGoogle-ink/gopay/blob/master/release_note.txt)
```go
import (
    "fmt"

    "github.com/iGoogle-ink/gopay"
)

func main() {
    fmt.Println("GoPay Version: ", gopay.Version)
}
```

---

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
* 下载资金账单（正式）：client.DownloadFundFlow()
* 交易保障：client.Report()
* 拉取订单评价数据（正式）：client.BatchQueryComment()
* 企业付款（正式）：client.Transfer()
* 查询企业付款（正式）：client.GetTransferInfo()
* 授权码查询OpenId（正式）：client.AuthCodeToOpenId()
* 公众号纯签约（正式）：client.EntrustPublic()
* APP纯签约-预签约接口-获取预签约ID（正式）：client.EntrustAppPre()
* H5纯签约（正式）：client.EntrustH5()
* 支付中签约（正式）：client.EntrustPaying()
* 请求单次分账（正式）：client.ProfitSharing() 
* 请求多次分账（正式）：client.MultiProfitSharing()
* 查询分账结果（正式）：client.ProfitSharingQuery()
* 添加分账接收方（正式）：client.ProfitSharingAddReceiver()
* 删除分账接收方（正式）：client.ProfitSharingRemoveReceiver()
* 完结分账（正式）：client.ProfitSharingFinish()
* 分账回退（正式）：client.ProfitSharingReturn()
* 分账回退结果查询（正式）：client.ProfitSharingReturnQuery()
* 企业付款到银行卡API（正式）：client.PayBank()
* 查询企业付款到银行卡API（正式）：client.QueryBank()
* 获取RSA加密公钥API（正式）：client.GetRSAPublicKey()
* 自定义方法请求微信API接口：client.PostWeChatAPISelf()

### 微信公共API

* wechat.GetParamSign() => 获取微信支付所需参数里的Sign值（通过支付参数计算Sign值）
* wechat.GetSanBoxParamSign() => 获取微信支付沙箱环境所需参数里的Sign值（通过支付参数计算Sign值）
* wechat.GetMiniPaySign() => 获取微信小程序支付所需要的paySign
* wechat.GetH5PaySign() => 获取微信内H5支付所需要的paySign
* wechat.GetAppPaySign() => 获取APP支付所需要的paySign
* wechat.ParseNotifyToBodyMap() => 解析微信支付异步通知的参数到BodyMap
* wechat.ParseNotify() => 解析微信支付异步通知的参数
* wechat.ParseRefundNotify() => 解析微信退款异步通知的参数
* wechat.VerifySign() => 微信同步返回参数验签或异步通知参数验签
* wechat.Code2Session() => 登录凭证校验：获取微信用户OpenId、UnionId、SessionKey
* wechat.GetAppletAccessToken() => 获取微信小程序全局唯一后台接口调用凭据
* wechat.GetAppletPaidUnionId() => 微信小程序用户支付完成后，获取该用户的 UnionId，无需用户授权
* wechat.GetUserInfo() => 微信公众号：获取用户基本信息(UnionID机制)
* wechat.GetUserInfoOpen() => 微信开放平台：获取用户个人信息(UnionID机制)
* wechat.DecryptOpenDataToStruct() => 加密数据，解密到指定结构体
* wechat.DecryptOpenDataToBodyMap() => 加密数据，解密到 BodyMap
* wechat.GetOpenIdByAuthCode() => 授权码查询openid
* wechat.GetAppLoginAccessToken() => App应用微信第三方登录，code换取access_token
* wechat.RefreshAppLoginAccessToken() => 刷新App应用微信第三方登录后，获取的 access_token
* wechat.DecryptRefundNotifyReqInfo() => 解密微信退款异步通知的加密数据

---

### QQ支付API

* 提交付款码支付：client.MicroPay()
* 撤销订单：client.Reverse()
* 统一下单：client.UnifiedOrder()
* 订单查询：client.OrderQuery()
* 关闭订单：client.CloseOrder()
* 申请退款：client.Refund()
* 退款查询：client.RefundQuery()
* 交易账单：client.StatementDown()
* 资金账单：client.AccRoll()
* 自定义方法请求微信API接口：client.PostQQAPISelf()

### QQ公共API

* qq.ParseNotifyToBodyMap() => 解析QQ支付异步通知的结果到BodyMap
* qq.ParseNotify() => 解析QQ支付异步通知的参数
* qq.VerifySign() => QQ同步返回参数验签或异步通知参数验签

---

### 支付宝支付API
> #### 因支付宝接口太多，如没实现的接口，还请开发者自行调用client.PostAliPayAPISelf()方法实现！
* 支付宝接口自行实现方法：client.PostAliPayAPISelf()
* 手机网站支付接口2.0（手机网站支付）：client.TradeWapPay()
* 统一收单下单并支付页面接口（电脑网站支付）：client.TradePagePay()
* APP支付接口2.0（APP支付）：client.TradeAppPay()
* 统一收单交易支付接口（商家扫用户付款码）：client.TradePay()
* 统一收单交易创建接口（小程序支付）：client.TradeCreate()
* 统一收单线下交易查询：client.TradeQuery()
* 统一收单交易关闭接口：client.TradeClose()
* 统一收单交易撤销接口：client.TradeCancel()
* 统一收单交易退款接口：client.TradeRefund()
* 统一收单退款页面接口：client.TradePageRefund()
* 统一收单交易退款查询：client.TradeFastPayRefundQuery()
* 统一收单交易结算接口：client.TradeOrderSettle()
* 统一收单线下交易预创建（用户扫商品收款码）：client.TradePrecreate()
* 单笔转账接口：client.FundTransUniTransfer()
* 转账业务单据查询接口：client.FundTransCommonQuery()
* 支付宝资金账户资产查询接口：client.FundAccountQuery()
* 换取授权访问令牌（获取access_token，user_id等信息）：client.SystemOauthToken()
* 支付宝会员授权信息查询接口（App支付宝登录）：client.UserInfoShare()
* 换取应用授权令牌（获取app_auth_token，auth_app_id，user_id等信息）：client.OpenAuthTokenApp()
* 获取芝麻信用分：client.ZhimaCreditScoreGet()
* 身份认证初始化服务：client.UserCertifyOpenInit()
* 身份认证开始认证（获取认证链接）：client.UserCertifyOpenCertify()
* 身份认证记录查询：client.UserCertifyOpenQuery()
* 用户登陆授权：client.UserInfoAuth()
* 支付宝商家账户当前余额查询：client.DataBillBalanceQuery()
* 查询对账单下载地址：client.DataBillDownloadUrlQuery()

### 支付宝公共API

* alipay.GetCertSN() => 获取证书SN号（app_cert_sn、alipay_cert_sn）
* alipay.GetRootCertSN() => 获取证书SN号（alipay_root_cert_sn）
* alipay.GetRsaSign() => 获取支付宝参数签名（参数sign值）
* alipay.SystemOauthToken() => 换取授权访问令牌（得到access_token，user_id等信息）
* alipay.FormatPrivateKey() => 格式化应用私钥
* alipay.FormatPublicKey() => 格式化支付宝公钥
* alipay.FormatURLParam() => 格式化支付宝请求URL参数
* alipay.ParseNotifyToBodyMap() => 解析支付宝支付异步通知的参数到BodyMap
* alipay.ParseNotifyByURLValues() => 通过 url.Values 解析支付宝支付异步通知的参数到BodyMap
* alipay.VerifySign() => 支付宝异步通知参数验签
* alipay.VerifySignWithCert() => 支付宝异步通知参数验签（证书方式）
* alipay.VerifySyncSign() => 支付宝同步返回参数验签
* alipay.DecryptOpenDataToStruct() => 解密支付宝开放数据到 结构体
* alipay.DecryptOpenDataToBodyMap() => 解密支付宝开放数据到 BodyMap
* alipay.MonitorHeartbeatSyn() => 验签接口

---

# 二、文档说明

* [GoDoc](https://godoc.org/github.com/iGoogle-ink/gopay)
* QQ支付 使用方法请参考微信的
* 所有方法，如有问题，请仔细查看 wechat_client_test.go、alipay_client_test.go 或 examples
* 有问题请加QQ群（加群验证答案：gopay），微信加好友拉群（微信群比较活跃）。在此，非常感谢那些加群后，提出意见和反馈问题的同志们！

QQ群：
<img width="226" height="300" src="https://raw.githubusercontent.com/iGoogle-ink/gopay/master/qq_gopay.png"/>
加微信拉群：
<img width="226" height="300" src="https://raw.githubusercontent.com/iGoogle-ink/gopay/master/wechat_jerry.png"/>

---

## 1、初始化GoPay客户端并做配置（HTTP请求均默认设置tls.Config{InsecureSkipVerify: true}）

* #### 微信

微信官方文档：[官方文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)

> 注意：微信支付下单等操作可用沙箱环境测试是否成功，但真正支付时，请使用正式环境 isProd = true，不然会报错。

```go
import (
	"github.com/iGoogle-ink/gopay/wechat"
)

// 初始化微信客户端
//    appId：应用ID
//    mchId：商户ID
//    apiKey：API秘钥值
//    isProd：是否是正式环境
client := wechat.NewClient("wxdaa2ab9ef87b5497", mchId, apiKey, false)

// 设置国家：不设置默认 中国国内
//    wechat.China：中国国内
//    wechat.China2：中国国内备用
//    wechat.SoutheastAsia：东南亚
//    wechat.Other：其他国家
client.SetCountry(wechat.China)

// 添加微信证书 Path 路径
//    certFilePath：apiclient_cert.pem 路径
//    keyFilePath：apiclient_key.pem 路径
//    pkcs12FilePath：apiclient_cert.p12 路径
//    返回err
client.AddCertFilePath()

```

* #### 支付宝

支付宝官方文档：[官方文档](https://openhome.alipay.com/docCenter/docCenter.htm)

支付宝RSA秘钥生成文档：[生成RSA密钥](https://opendocs.alipay.com/open/291/105971) （推荐使用 RSA2）

沙箱环境使用说明：[文档地址](https://opendocs.alipay.com/open/200/105311)

```go
import (
	"github.com/iGoogle-ink/gopay/alipay"
)

// 初始化支付宝客户端
//    appId：应用ID
//    privateKey：应用私钥，支持PKCS1和PKCS8
//    isProd：是否是正式环境
client := alipay.NewClient("2016091200494382", privateKey, false)

// 设置支付宝请求 公共参数
//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
client.SetLocation().                       // 设置时区，不设置或出错均为默认服务器时间
    SetPrivateKeyType().                    // 设置 支付宝 私钥类型，alipay.PKCS1 或 alipay.PKCS8，默认 PKCS1
    SetAliPayRootCertSN().                  // 设置支付宝根证书SN，通过 alipay.GetRootCertSN() 获取
    SetAppCertSN().                         // 设置应用公钥证书SN，通过 alipay.GetCertSN() 获取
    SetAliPayPublicCertSN().                // 设置支付宝公钥证书SN，通过 alipay.GetCertSN() 获取
    SetCharset("utf-8").                    // 设置字符编码，不设置默认 utf-8
    SetSignType(alipay.RSA2).               // 设置签名类型，不设置默认 RSA2
    SetReturnUrl("https://www.gopay.ink").  // 设置返回URL
    SetNotifyUrl("https://www.gopay.ink").  // 设置异步通知URL
    SetAppAuthToken().                      // 设置第三方应用授权
    SetAuthToken()                          // 设置个人信息授权

err := client.SetCertSnByPath("appCertPublicKey.crt", "alipayRootCert.crt", "alipayCertPublicKey_RSA2.crt")
```

## 2、初始化并赋值BodyMap（client的方法所需的入参）

* #### 微信请求参数

具体参数请根据不同接口查看：[微信支付接口文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)
```go
import (
	"github.com/iGoogle-ink/gopay/wechat"
)

// 初始化 BodyMap
bm := make(gopay.BodyMap)
bm.Set("nonce_str", gotil.GetRandomString(32))
bm.Set("body", "小程序测试支付")
bm.Set("out_trade_no", number)
bm.Set("total_fee", 1)
bm.Set("spbill_create_ip", "127.0.0.1")
bm.Set("notify_url", "http://www.gopay.ink")
bm.Set("trade_type", wechat.TradeType_Mini)
bm.Set("device_info", "WEB")
bm.Set("sign_type", wechat.SignType_MD5)
bm.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")

// 嵌套json格式数据（例如：H5支付的 scene_info 参数）
h5Info := make(map[string]string)
h5Info["type"] = "Wap"
h5Info["wap_url"] = "http://www.gopay.ink"
h5Info["wap_name"] = "H5测试支付"

sceneInfo := make(map[string]map[string]string)
sceneInfo["h5_info"] = h5Info

bm.Set("scene_info", sceneInfo)

// 参数 sign ，可单独生成赋值到BodyMap中；也可不传sign参数，client内部会自动获取
// 如需单独赋值 sign 参数，需通过下面方法，最后获取sign值并在最后赋值此参数
sign := wechat.GetParamSign("wxdaa2ab9ef87b5497", mchId, apiKey, body)
// sign, _ := wechat.GetSanBoxParamSign("wxdaa2ab9ef87b5497", mchId, apiKey, body)
bm.Set("sign", sign)
```

* #### 支付宝请求参数

具体参数请根据不同接口查看：[支付宝支付API接口文档](https://opendocs.alipay.com/apis/api_1/alipay.trade.wap.pay)
```go
// 初始化 BodyMap
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
// 手机网站支付是通过服务端获取支付URL后，然后返回给客户端，请求URL地址即可打开支付页面
payUrl, err := client.TradeWapPay(bm)

// 电脑网站支付是通过服务端获取支付URL后，然后返回给客户端，请求URL地址即可打开支付页面
payUrl, err := client.TradePagePay(bm)

// APP支付是通过服务端获取支付参数后，然后通过Android/iOS客户端的SDK调用支付功能
payParam, err := client.TradeAppPay(bm)

// 商家使用扫码枪等条码识别设备扫描用户支付宝钱包上的条码/二维码，完成收款
aliRsp, err := client.TradePay(bm)

// 支付宝小程序支付时 buyer_id 为必传参数，需要提前获取，获取方法如下两种
//    1、alipay.SystemOauthToken()     返回取值：rsp.SystemOauthTokenResponse.UserId
//    2、client.SystemOauthToken()    返回取值：aliRsp.SystemOauthTokenResponse.UserId
aliRsp, err := client.TradeCreate(bm)
aliRsp, err := client.TradeQuery(bm)
aliRsp, err := client.TradeClose(bm)
aliRsp, err := client.TradeCancel(bm)
aliRsp, err := client.TradeRefund(bm)
aliRsp, err := client.TradePageRefund(bm)
aliRsp, err := client.TradeFastPayRefundQuery(bm)
aliRsp, err := client.TradeOrderSettle(bm)
aliRsp, err := client.TradePrecreate(bm)
aliRsp, err := client.FundTransUniTransfer(bm)
aliRsp, err := client.FundTransCommonQuery(bm)
aliRsp, err := client.FundAccountQuery(bm)
aliRsp, err := client.SystemOauthToken(bm)
aliRsp, err := client.OpenAuthTokenApp(bm)
aliRsp, err := client.ZhimaCreditScoreGet(bm)
aliRsp, err := client.UserCertifyOpenInit(bm)
aliRsp, err := client.UserCertifyOpenCertify(bm)
aliRsp, err := client.UserCertifyOpenQuery(bm)
```

## 4、微信统一下单后，获取微信小程序支付、APP支付、微信内H5支付所需要的 paySign

* #### 微信（只有微信需要此操作）
微信小程序支付官方文档：[微信小程序支付API](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/payment/wx.requestPayment.html)

APP支付官方文档：[APP端调起支付的参数列表文档](https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_12)

微信内H5支付官方文档：[微信内H5支付文档](https://pay.weixin.qq.com/wiki/doc/api/external/jsapi.php?chapter=7_7&index=6)
```go
import (
	"github.com/iGoogle-ink/gopay/wechat"
)

// ====微信小程序 paySign====
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
prepayId := "prepay_id=" + wxRsp.PrepayId   // 此处的 wxRsp.PrepayId ,统一下单成功后得到
// 获取微信小程序支付的 paySign
//    appId：AppID
//    nonceStr：随机字符串
//    prepayId：统一下单成功后得到的值
//    signType：签名方式，务必与统一下单时用的签名方式一致
//    timeStamp：时间
//    apiKey：API秘钥值
paySign := wechat.GetMiniPaySign(AppID, wxRsp.NonceStr, prepayId, wechat.SignType_MD5, timeStamp, apiKey)

// ====APP支付 paySign====
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
// 获取APP支付的 paySign
// 注意：package 参数因为是固定值，无需开发者再传入
//    appId：AppID
//    partnerid：partnerid
//    nonceStr：随机字符串
//    prepayId：统一下单成功后得到的值
//    signType：签名方式，务必与统一下单时用的签名方式一致
//    timeStamp：时间
//    apiKey：API秘钥值
paySign := wechat.GetAppPaySign(appid, partnerid, wxRsp.NonceStr, wxRsp.PrepayId, wechat.SignType_MD5, timeStamp, apiKey)

// ====微信内H5支付 paySign====
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
packages := "prepay_id=" + wxRsp.PrepayId   // 此处的 wxRsp.PrepayId ,统一下单成功后得到
// 获取微信内H5支付 paySign
//    appId：AppID
//    nonceStr：随机字符串
//    packages：统一下单成功后拼接得到的值
//    signType：签名方式，务必与统一下单时用的签名方式一致
//    timeStamp：时间
//    apiKey：API秘钥值
paySign := wechat.GetH5PaySign(AppID, wxRsp.NonceStr, packages, wechat.SignType_MD5, timeStamp, apiKey)
```

## 5、同步返回参数验签Sign、异步通知参数解析和验签Sign、异步通知返回

异步参数需要先解析，解析出来的结构体或BodyMap再验签

[Gin Web框架](https://github.com/gin-gonic/gin)

[Echo Web框架](https://github.com/labstack/echo)

异步通知处理完后，需回复平台固定数据

* #### 微信
```go
import (
	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/wechat"
)

// ====同步返回参数验签Sign====
wxRsp, err := client.UnifiedOrder(bm)
// 微信同步返回参数验签或异步通知参数验签
//    apiKey：API秘钥值
//    signType：签名类型（调用API方法时填写的类型）
//    bean：微信同步返回的结构体 wxRsp 或 异步通知解析的结构体 notifyReq
//    返回参数 ok：是否验签通过
//    返回参数 err：错误信息
ok, err := wechat.VerifySign(apiKey, wechat.SignType_MD5, wxRsp)

// ====支付异步通知参数解析和验签Sign====
// 解析支付异步通知的参数
//    req：*http.Request
//    ctx.Request   是 gin 框架的获取 *http.Request
//    ctx.Request() 是 echo 框架的获取 *http.Request
//    返回参数 notifyReq：通知的参数
//    返回参数 err：错误信息
notifyReq, err := wechat.ParseNotifyToBodyMap(ctx.Request)

// 验签操作
ok, err := wechat.VerifySign(apiKey, wechat.SignType_MD5, notifyReq)

// ====退款异步通知参数解析，退款通知无sign，不用验签====
// 
// 解析退款异步通知的参数，解析出来的 req_info 是加密数据，需解密
//    req：*http.Request
//    ctx.Request   是 gin 框架的获取 *http.Request
//    ctx.Request() 是 echo 框架的获取 *http.Request
//    返回参数 notifyReq：通知的参数
//    返回参数 err：错误信息
notifyReq, err := wechat.ParseNotifyToBodyMap(c.Request)
 或
notifyReq, err := wechat.ParseRefundNotify(c.Request)

// ==解密退款异步通知的加密参数 req_info ==
refundNotify, err := wechat.DecryptRefundNotifyReqInfo(notifyReq.ReqInfo, apiKey)

// ==异步通知，返回给微信平台的信息==
rsp := new(wechat.NotifyResponse) // 回复微信的数据
rsp.ReturnCode = gopay.SUCCESS
rsp.ReturnMsg = gopay.OK

return c.String(http.StatusOK, rsp.ToXmlString())   // 此写法是 echo 框架返回微信的写法
c.String(http.StatusOK, "%s", rsp.ToXmlString())    // 此写法是 gin 框架返回微信的写法
```

* #### 支付宝

注意：APP支付、手机网站支付、电脑网站支付 暂不支持同步返回验签

支付宝支付后的同步/异步通知验签文档：[支付结果通知](https://opendocs.alipay.com/open/200/106120)
```go
import (
	"github.com/iGoogle-ink/gopay/alipay"
)

// ====同步返回参数验签Sign====
aliRsp, err := client.TradePay(bm)
// 支付宝同步返回验签
//    注意：APP支付，手机网站支付，电脑网站支付 暂不支持同步返回验签
//    aliPayPublicKey：支付宝公钥
//    signData：待验签参数，aliRsp.SignData
//    sign：待验签sign，aliRsp.Sign
//    返回参数ok：是否验签通过
//    返回参数err：错误信息
ok, err := alipay.VerifySyncSign(aliPayPublicKey, aliRsp.SignData, aliRsp.Sign)

// ====异步通知参数解析和验签Sign====
// 解析异步通知的参数
//    req：*http.Request
//    返回参数 notifyReq：通知的参数
//    返回参数 err：错误信息
notifyReq, err = alipay.ParseNotifyToBodyMap(c.Request())     // c.Request()是 echo 框架的获取
 或
notifyReq, err = alipay.ParseNotifyByURLValues()

// 验签操作
ok, err = alipay.VerifySign(aliPayPublicKey, notifyReq)
// 证书验签操作
ok, err = alipay.VerifySignWithCert("alipayCertPublicKey_RSA2.crt", aliPayPublicKey, notifyReq)

// ==异步通知，返回支付宝平台的信息==
//    文档：https://opendocs.alipay.com/open/203/105286
//    程序执行完后必须打印输出“success”（不包含引号）。如果商户反馈给支付宝的字符不是success这7个字符，支付宝服务器会不断重发通知，直到超过24小时22分钟。一般情况下，25小时以内完成8次通知（通知的间隔频率一般是：4m,10m,10m,1h,2h,6h,15h）
return c.String(http.StatusOK, "success")   // 此写法是 echo 框架返回支付宝的写法
c.String(http.StatusOK, "%s", "success")    // 此写法是 gin 框架返回支付宝的写法
```

## 6、微信、支付宝 公共API（仅部分说明）

* #### 微信 公共API

官方文档：[code2Session](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html)

button按钮获取手机号码：[button组件文档](https://developers.weixin.qq.com/miniprogram/dev/component/button.html)

微信解密算法文档：[解密算法文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)
```go
import (
	"github.com/iGoogle-ink/gopay/wechat"
)

// 获取微信小程序用户的OpenId、SessionKey、UnionId
//    appId：微信小程序的APPID
//    appSecret：微信小程序的AppSecret
//    wxCode：小程序调用wx.login 获取的code
sessionRsp, err := wechat.Code2Session(appId, appSecret, wxCode)

// ====解密微信加密数据到指定结构体====

//小程序获取手机号
data := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
iv := "Cds8j3VYoGvnTp1BrjXdJg=="
session := "lyY4HPQbaOYzZdG+JcYK9w=="
phone := new(wechat.UserPhone)
// 解密开放数据
//    encryptedData：包括敏感数据在内的完整用户信息的加密数据，小程序获取到
//    iv：加密算法的初始向量，小程序获取到
//    sessionKey：会话密钥，通过 wechat.Code2Session() 方法获取到
//    beanPtr：需要解析到的结构体指针，操作完后，声明的结构体会被赋值
err := wechat.DecryptOpenDataToStruct(data, iv, session, phone)
fmt.Println(*phone)
// 获取微信小程序用户信息
sessionKey := "tiihtNczf5v6AKRyjwEUhQ=="
encryptedData := "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="
iv2 := "r7BXXKkLb8qrSNn05n0qiA=="

// 微信小程序 用户信息
userInfo := new(wechat.AppletUserInfo)
err = wechat.DecryptOpenDataToStruct(encryptedData, iv2, sessionKey, userInfo)
fmt.Println(*userInfo)

data := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
iv := "Cds8j3VYoGvnTp1BrjXdJg=="
session := "lyY4HPQbaOYzZdG+JcYK9w=="
    
// 解密开放数据到 BodyMap
//    encryptedData:包括敏感数据在内的完整用户信息的加密数据
//    iv:加密算法的初始向量
//    sessionKey:会话密钥
bm, err := wechat.DecryptOpenDataToBodyMap(data, iv, session)
if err != nil {
     fmt.Println("err:", err)
     return
}
fmt.Println("WeChatUserPhone:", bm)
```

* #### 支付宝 公共API

支付宝换取授权访问令牌文档：[换取授权访问令牌](https://opendocs.alipay.com/apis/api_9/alipay.system.oauth.token)

获取用户手机号文档：[获取用户手机号](https://opendocs.alipay.com/mini/api/getphonenumber)

支付宝加解密文档：[AES配置文档](https://opendocs.alipay.com/mini/introduce/aes) ，[AES加解密文档](https://opendocs.alipay.com/open/common/104567)
```go
import (
	"github.com/iGoogle-ink/gopay/alipay"
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
fmt.Println(*phone)
```

## 开源不易，讲究的朋友可以给个赞赏
<font color='#0088ff'>微信：</font>
<img width="200" height="200" src="https://raw.githubusercontent.com/iGoogle-ink/gopay/master/zanshang_wx.png"/>
<font color='#0088ff'>支付宝：</font>
<img width="200" height="200" src="https://raw.githubusercontent.com/iGoogle-ink/gopay/master/zanshang_zfb.png"/>

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