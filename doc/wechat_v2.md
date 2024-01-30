## 微信v2

> #### 推荐使用v3接口，官方在v3接口实现未覆盖或gopay未开发的接口，还继续用v2接口。

- 已实现API列表附录：[API 列表附录](https://github.com/go-pay/gopay/blob/main/doc/wechat_v2.md#%E9%99%84%E5%BD%95)

---

### 1、初始化微信v2客户端并做配置

微信v2官方文档：[微信v2官方文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)

> 注意：微信支付下单等操作可用沙箱环境测试是否成功，但真正支付时，请使用正式环境 `isProd = true`，不然会报错。

> 微信证书二选一：只传 `apiclient_cert.pem` 和 `apiclient_key.pem` 或者只传 `apiclient_cert.p12`

```go
import (
    "github.com/go-pay/gopay/wechat"
)

// 初始化微信客户端
//    appId：应用ID
//    mchId：商户ID
//    apiKey：API秘钥值
//    isProd：是否是正式环境
client := wechat.NewClient("wxdaa2ab9ef87b5497", mchId, apiKey, false)

// 打开Debug开关，输出请求日志，默认关闭
client.DebugSwitch = gopay.DebugOn

// 自定义配置http请求接收返回结果body大小，默认 10MB
client.SetBodySize() // 没有特殊需求，可忽略此配置

// 设置国家：不设置默认 中国国内
//    wechat.China：中国国内
//    wechat.China2：中国国内备用
//    wechat.SoutheastAsia：东南亚
//    wechat.Other：其他国家
client.SetCountry(wechat.China)

// 添加微信pem证书
client.AddCertPemFilePath()
client.AddCertPemFileContent()
 或
// 添加微信pkcs12证书
client.AddCertPkcs12FilePath()
client.AddCertPkcs12FileContent()
```

### 2、API 方法调用及入参

- #### 微信请求参数

> 微信V2接口通用参数（mch_id、appid、sign）无需传入，client 请求时会默认处理

具体参数请根据不同接口查看：[微信支付接口文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)

```go
import (
    "github.com/go-pay/util"
    "github.com/go-pay/gopay/wechat"
)

// 初始化 BodyMap
bm := make(gopay.BodyMap)
bm.Set("nonce_str", util.RandomString(32)).
    Set("body", "H5支付").
    Set("out_trade_no", number).
    Set("total_fee", 1).
    Set("spbill_create_ip", "127.0.0.1").
    Set("notify_url", "https://www.fmm.ink").
    Set("trade_type", TradeType_H5).
    Set("device_info", "WEB").
    Set("sign_type", SignType_MD5).
    SetBodyMap("scene_info", func(bm gopay.BodyMap) {
        bm.SetBodyMap("h5_info", func(bm gopay.BodyMap) {
            bm.Set("type", "Wap")
            bm.Set("wap_url", "https://www.fmm.ink")
            bm.Set("wap_name", "H5测试支付")
        })
    }) /*.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")*/
```

- #### client 方法调用

```go
wxRsp, err := client.UnifiedOrder(bm)
wxRsp, err := client.Micropay(bm)
wxRsp, err := client.QueryOrder(bm)
wxRsp, err := client.CloseOrder(bm)
wxRsp, err := client.Reverse(bm)
wxRsp, err := client.Refund(bm)
wxRsp, err := client.QueryRefund(bm)
wxRsp, err := client.DownloadBill(bm)
wxRsp, err := client.DownloadFundFlow(bm)
wxRsp, err := client.BatchQueryComment(bm)
wxRsp, err := client.Transfer(bm)
...
```

### 3、微信统一下单后，获取微信小程序支付、APP支付、微信内H5支付所需要的 paySign

> 微信小程序支付官方文档：[微信小程序支付API](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/payment/wx.requestPayment.html)

> APP支付官方文档：[APP端调起支付的参数列表文档](https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_12)

> 微信内H5支付官方文档：[微信内H5支付文档](https://pay.weixin.qq.com/wiki/doc/api/wxpay/ch/pay/OfficialPayMent/chapter5_5.shtml)

```go
import (
    "github.com/go-pay/gopay/wechat"
)

// ====微信小程序 paySign====
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
packages := "prepay_id=" + wxRsp.PrepayId   // 此处的 wxRsp.PrepayId ,统一下单成功后得到
// 获取微信小程序支付的 paySign
//    appId：AppID
//    nonceStr：随机字符串
//    packages：统一下单成功后拼接得到的值
//    signType：签名方式，务必与统一下单时用的签名方式一致
//    timeStamp：时间
//    apiKey：API秘钥值
paySign := wechat.GetMiniPaySign(AppID, wxRsp.NonceStr, packages, wechat.SignType_MD5, timeStamp, apiKey)

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

### 4、同步返回参数验签Sign、异步通知参数解析和验签Sign、异步通知返回

> 异步通知请求参数需要先解析，解析出来的结构体或BodyMap再验签（此处需要注意，`http.Request.Body` 只能解析一次，如果需要解析前调试，请处理好Body复用问题）

[Gin Web框架（推荐）](https://github.com/gin-gonic/gin)

[Echo Web框架](https://github.com/labstack/echo)

```go
import (
    "github.com/go-pay/gopay"
    "github.com/go-pay/gopay/wechat"
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

// 此写法是 gin 框架返回微信的写法
c.String(http.StatusOK, "%s", rsp.ToXmlString())
// 此写法是 echo 框架返回微信的写法
return c.String(http.StatusOK, rsp.ToXmlString())
```

### 5、公共API（仅部分说明）

---

## 附录：

### 微信支付v2 API

* 统一下单：`client.UnifiedOrder()`
    * JSAPI - JSAPI支付（或小程序支付）
    * NATIVE - Native支付
    * APP - app支付
    * MWEB - H5支付
* 提交付款码支付：`client.Micropay()`
* 查询订单：`client.QueryOrder()`
* 关闭订单：`client.CloseOrder()`
* 撤销订单：`client.Reverse()`
* 申请退款：`client.Refund()`
* 查询退款：`client.QueryRefund()`
* 下载对账单：`client.DownloadBill()`
* 下载资金账单（正式）：`client.DownloadFundFlow()`
* 交易保障：`client.Report()`
* 拉取订单评价数据（正式）：`client.BatchQueryComment()`
* 企业付款（正式）：`client.Transfer()`
* 查询企业付款（正式）：`client.GetTransferInfo()`
* 授权码查询OpenId（正式）：`client.AuthCodeToOpenId()`
* 公众号纯签约（正式）：`client.EntrustPublic()`
* APP纯签约-预签约接口-获取预签约ID（正式）：`client.EntrustAppPre()`
* H5纯签约（正式）：`client.EntrustH5()`
* 支付中签约（正式）：`client.EntrustPaying()`
* 请求单次分账（正式）：`client.ProfitSharing()`
* 请求多次分账（正式）：`client.MultiProfitSharing()`
* 查询分账结果（正式）：`client.ProfitSharingQuery()`
* 查询订单待分账金额 （正式）：`client.ProfitSharingOrderAmountQuery()`
* 查询最大分账比例 （正式）：`client.ProfitSharingMerchantRatioQuery()`
* 添加分账接收方（正式）：`client.ProfitSharingAddReceiver()`
* 删除分账接收方（正式）：`client.ProfitSharingRemoveReceiver()`
* 完结分账（正式）：`client.ProfitSharingFinish()`
* 分账回退（正式）：`client.ProfitSharingReturn()`
* 分账回退结果查询（正式）：`client.ProfitSharingReturnQuery()`
* 企业付款到银行卡API（正式）：`client.PayBank()`
* 查询企业付款到银行卡API（正式）：`client.QueryBank()`
* 获取RSA加密公钥API（正式）：`client.GetRSAPublicKey()`
* 发放现金红包：`client.SendCashRed()`
* 发放现金裂变红包：`client.SendGroupCashRed()`
* 发放小程序红包：`client.SendAppletRed()`
* 查询红包记录：`client.QueryRedRecord()`
* 订单附加信息提交（海关）：`client.CustomsDeclareOrder()`
* 订单附加信息查询（海关）：`client.CustomsDeclareQuery()`
* 订单附加信息重推（海关）：`client.CustomsReDeclareOrder()`
* 自定义方法请求微信API接口：`client.PostWeChatAPISelf()`

### 微信公共v2 API

* `wechat.GetParamSign()` => 获取微信支付所需参数里的Sign值（通过支付参数计算Sign值）
* `wechat.GetSanBoxParamSign()` => 获取微信支付沙箱环境所需参数里的Sign值（通过支付参数计算Sign值）
* `wechat.GetMiniPaySign()` => 获取微信小程序支付所需要的paySign
* `wechat.GetH5PaySign()` => 获取微信内H5支付所需要的paySign
* `wechat.GetAppPaySign()` => 获取APP支付所需要的paySign
* `wechat.ParseNotifyToBodyMap()` => 解析微信支付异步通知的参数到BodyMap
* `wechat.ParseNotify()` => 解析微信支付异步通知的参数
* `wechat.ParseRefundNotify()` => 解析微信退款异步通知的参数
* `wechat.VerifySign()` => 微信同步返回参数验签或异步通知参数验签
* `wechat.GetOpenIdByAuthCode()` => 授权码查询openid
* `wechat.GetOauth2AccessToken()` => 微信第三方登录，code 换取 access_token
* `wechat.RefreshOauth2AccessToken()` => 刷新微信第三方登录后，获取到的 access_token
* `wechat.CheckOauth2AccessToken()` => 检验授权凭证（access_token）是否有效
* `wechat.DecryptRefundNotifyReqInfo()` => 解密微信退款异步通知的加密数据
