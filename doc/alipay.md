## 支付宝

> #### 因支付宝接口太多，如没实现的接口，还请开发者自行调用 `client.PostAliPayAPISelfV2()`方法实现！请参考 `client_test.go` 内的 `TestClient_PostAliPayAPISelfV2()` 方法

> #### 希望有时间的伙伴儿Fork完后，补充并提交Pull Request，一起完善支付宝各个类别下的接口到相应的go文件中

- 已实现API列表附录：[API 列表附录](https://github.com/go-pay/gopay/blob/main/doc/alipay.md#%E9%99%84%E5%BD%95)

- 支付宝官方文档：[官方文档](https://openhome.alipay.com/docCenter/docCenter.htm)

- 支付宝RSA秘钥生成文档：[生成RSA密钥](https://opendocs.alipay.com/common/02kipl) （推荐使用 RSA2）

- 沙箱环境(新) 使用说明：[新版沙箱文档](https://opendocs.alipay.com/common/05yvy1)

---

### 1、初始化支付宝客户端并做配置

> 具体API使用介绍，请参考 `gopay/alipay/client_test.go`

```go
import (
    "github.com/go-pay/gopay/alipay"
    "github.com/go-pay/xlog"
)

// 初始化支付宝客户端
// appid：应用ID
// privateKey：应用私钥，支持PKCS1和PKCS8
// isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。
client, err := alipay.NewClient("2016091200494382", privateKey, false)
if err != nil {
    xlog.Error(err)
    return
}

// 自定义配置http请求接收返回结果body大小，默认 10MB
client.SetBodySize() // 没有特殊需求，可忽略此配置

// 打开Debug开关，输出日志，默认关闭
client.DebugSwitch = gopay.DebugOn

// 设置支付宝请求 公共参数
//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
SetCharset(alipay.UTF8).  // 设置字符编码，不设置默认 utf-8
SetSignType(alipay.RSA2). // 设置签名类型，不设置默认 RSA2
SetReturnUrl("https://www.fmm.ink").            // 设置返回URL
    SetNotifyUrl("https://www.fmm.ink").        // 设置异步通知URL
    SetAppAuthToken()                           // 设置第三方应用授权

// 设置biz_content加密KEY，设置此参数默认开启加密（目前不可用，设置后会报错）
//client.SetAESKey("1234567890123456")

// 自动同步验签（只支持证书模式）
// 传入 alipayPublicCert.crt 内容
client.AutoVerifySign([]byte("alipayPublicCert.crt bytes"))

// 公钥证书模式，需要传入证书，以下两种方式二选一
// 证书路径
err := client.SetCertSnByPath("appPublicCert.crt", "alipayRootCert.crt", "alipayPublicCert.crt")
// 证书内容
err := client.SetCertSnByContent("appPublicCert.crt bytes", "alipayRootCert bytes", "alipayPublicCert.crt bytes")
```

### 2、API 方法调用及入参

> 具体参数请根据不同接口查看：[支付宝支付API接口文档](https://opendocs.alipay.com/apis)

> 业务错误处理：当 `err != nil` 时，可通过 `alipay.IsBizError()` 捕获业务错误状态码和说明。
> 不在乎 `BizError` 的可忽略统一判错处理

> ★入参 BodyMap中，支持如下公共参数在当次请求中自定义设置：`version`、`return_url`、`notify_url`、`app_auth_token`

- 统一收单交易支付接口 - 示例

```go
import (
    "github.com/go-pay/gopay"
)

// 初始化 BodyMap
bm := make(gopay.BodyMap)
bm.Set("subject", "条码支付").
    Set("scene", "bar_code").
    Set("auth_code", "286248566432274952").
    Set("out_trade_no", "GZ201909081743431443").
    Set("total_amount", "0.01").
    Set("timeout_express", "2m")

aliRsp, err := client.TradePay(bm)
if err != nil {
    if bizErr, ok := alipay.IsBizError(err); ok {
        xlog.Errorf("%+v", bizErr)
        // do something
        return
    }
    xlog.Errorf("client.TradePay(%+v),err:%+v", bm, err)
    return
}
```

### 3、同步返回参数验签Sign、异步通知参数解析和验签Sign、异步通知返回

> 异步通知请求参数需要先解析，解析出来的结构体或BodyMap再验签（此处需要注意，`http.Request.Body` 只能解析一次，如果需要解析前调试，请处理好Body复用问题）

[Gin Web框架（推荐）](https://github.com/gin-gonic/gin)

[Echo Web框架](https://github.com/labstack/echo)

> 注意：APP支付、手机网站支付、电脑网站支付 不支持同步返回验签

> 支付宝支付后的同步/异步通知验签文档：[支付结果通知](https://opendocs.alipay.com/open/200/106120)

- 同步返回验签，手动验签（如已开启自动验签，则无需手动验签操作）

```go
import (
    "github.com/go-pay/gopay/alipay"
)

aliRsp, err := client.TradePay(bm)
if err != nil {
    xlog.Error("err:", err)
    return
}

// 公钥模式验签
//    注意：APP支付，手机网站支付，电脑网站支付 不支持同步返回验签
//    aliPayPublicKey：支付宝平台获取的支付宝公钥
//    signData：待验签参数，aliRsp.SignData
//    sign：待验签sign，aliRsp.Sign
ok, err := alipay.VerifySyncSign(aliPayPublicKey, aliRsp.SignData, aliRsp.Sign)

// 公钥证书模式验签
//    aliPayPublicKeyCert：支付宝公钥证书存放路径 alipayPublicCert.crt 或文件内容[]byte
//    signData：待验签参数，aliRsp.SignData
//    sign：待验签sign，aliRsp.Sign
ok, err := alipay.VerifySyncSignWithCert(aliPayPublicKeyCert, aliRsp.SignData, aliRsp.Sign)
```

- 异步通知验签

```go
import (
    "github.com/go-pay/gopay/alipay"
)

// 解析异步通知的参数
// req：*http.Request
notifyReq, err = alipay.ParseNotifyToBodyMap(c.Request)     // c.Request 是 gin 框架的写法
if err != nil {
    xlog.Error(err)
    return
}
 或
// value：url.Values
notifyReq, err = alipay.ParseNotifyByURLValues()
if err != nil {
    xlog.Error(err)
    return
}

// 支付宝异步通知验签（公钥模式）
ok, err = alipay.VerifySign(aliPayPublicKey, notifyReq)

// 支付宝异步通知验签（公钥证书模式）
ok, err = alipay.VerifySignWithCert("alipayPublicCert.crt content", notifyReq)

// 如果需要，可将 BodyMap 内数据，Unmarshal 到指定结构体指针 ptr
err = notifyReq.Unmarshal(ptr)

// ====异步通知，返回支付宝平台的信息====
// 文档：https://opendocs.alipay.com/open/203/105286
// 程序执行完后必须打印输出“success”（不包含引号）。如果商户反馈给支付宝的字符不是success这7个字符，支付宝服务器会不断重发通知，直到超过24小时22分钟。一般情况下，25小时以内完成8次通知（通知的间隔频率一般是：4m,10m,10m,1h,2h,6h,15h）

// 此写法是 gin 框架返回支付宝的写法
c.String(http.StatusOK, "%s", "success")

// 此写法是 echo 框架返回支付宝的写法
return c.String(http.StatusOK, "success")
```

### 4、支付宝 公共API（仅部分说明）

> 支付宝换取授权访问令牌文档：[换取授权访问令牌](https://opendocs.alipay.com/apis/api_9/alipay.system.oauth.token)

> 支付宝小程序 获取用户手机号文档：[获取用户手机号](https://opendocs.alipay.com/mini/api/getphonenumber)

> 支付宝加解密文档：[AES加解密文档](https://opendocs.alipay.com/open/common/104567)

```go
import (
    "github.com/go-pay/gopay/alipay"
    "github.com/go-pay/xlog"
)

// 换取授权访问令牌（默认使用utf-8，RSA2）
// appId：应用ID
// privateKey：应用私钥，支持PKCS1和PKCS8
// grantType：值为 authorization_code 时，代表用code换取；值为 refresh_token 时，代表用refresh_token换取，传空默认code换取
// codeOrToken：支付宝授权码或refresh_token
rsp, err := alipay.SystemOauthToken(appId, privateKey, grantType, codeOrToken)
if err != nil {
    xlog.Error(err)
    return
}

// 解密支付宝开放数据带到指定结构体
// 以小程序获取手机号为例
phone := new(alipay.UserPhone)
// 解密支付宝开放数据
// encryptedData:包括敏感数据在内的完整用户信息的加密数据
// secretKey:AES密钥，支付宝管理平台配置
// beanPtr:需要解析到的结构体指针
err := alipay.DecryptOpenDataToStruct(encryptedData, secretKey, phone)
xlog.Infof("%+v", phone)
```

---

## 附录：

### 支付宝支付 API

* 支付宝接口自行实现方法：`client.PostAliPayAPISelfV2()`
* <font color='#027AFF' size='4'>支付产品</font>
  * 当面付
    * 付款码支付接口(商家扫用户付款码)：`client.TradePay()`
    * 统一收单线下交易预创建接口(用户扫商品收款码)：`client.TradePrecreate()`
  * App支付
    * APP支付接口2.0(APP支付)：`client.TradeAppPay()`
  * 手机网站支付
    * 手机网站支付接口2.0(手机网站支付)：`client.TradeWapPay()`
  * 电脑网站支付
    * 统一收单下单并支付页面接口(电脑网站支付)：`client.TradePagePay()`
  * 刷脸付
    * 刷脸支付初始化接口：`client.ZolozAuthenticationSmilepayInitialize()`
    * 查询刷脸结果信息接口：`client.ZolozAuthenticationCustomerFtokenQuery()`
  * 预授权支付
    * 线上资金授权冻结接口：`client:FundAuthOrderAppFreeze()`
    * 资金授权操作查询接口：`client.FundAuthOperationDetailQuery()`
    * 资金授权撤销接口：`client.FundAuthOperationCancel()`
    * 资金授权解冻接口：`client.FundAuthOrderUnfreeze()`
    * 资金授权发码接口：`client.FundAuthOrderVoucherCreate()`
    * 资金授权冻结接口：`client.FundAuthOrderFreeze()`
  * 交易
    * 统一收单交易创建接口：`client.TradeCreate()`
    * 统一收单交易订单支付接口：`client.TradeOrderPay()`
    * 统一收单线下交易查询: `client.TradeQuery()`
    * 统一收单交易退款接口: `client.TradeRefund()`
    * 统一收单交易退款查询: `client.TradeFastPayRefundQuery()`
    * 统一收单交易关闭接口: `client.TradeClose()`
    * 统一收单交易撤销接口: `client.TradeCancel()`
    * 支付宝订单信息同步接口: `client.TradeOrderInfoSync()`
  * 账单
    * 查询对账单下载地址：`client.DataBillDownloadUrlQuery()`
  * 商家分账
    * 分账关系维护
      * 分账关系绑定接口：`client.TradeRelationBind()`
      * 分账关系解绑接口：`client.TradeRelationUnbind()`
      * 分账关系查询接口：`client.TradeRelationBatchQuery()`
    * 分账请求
      * 统一收单交易结算接口：`client.TradeOrderSettle()`
      * 统一收单确认结算接口：`client.TradeSettleConfirm()`
    * 分账查询 
      * 分账比例查询：`client.TradeRoyaltyRateQuery()`
      * 分账剩余金额查询：`client.TradeOrderOnSettleQuery()`
      * 交易分账查询接口：`client.TradeOrderSettleQuery()`
  * 商家扣款
    * 支付宝个人协议页面签约接口: `client.UserAgreementPageSign()`
    * 支付宝个人协议页面签约接口(App 专用,生成唤醒签约页面链接): `client.UserAgreementPageSignInApp()`
    * 支付宝个人代扣协议查询接口: `client.UserAgreementQuery()`
    * 支付宝个人代扣协议解约接口: `client.UserAgreementPageUnSign()`
    * 周期性扣款协议执行计划修改接口: `client.UserAgreementExecutionplanModify()`
* <font color='#027AFF' size='4'>私域产品</font>
  * 支付有礼
    * 推广计划
* <font color='#027AFF' size='4'>公域产品</font>
  * 经营推广
    * 推广计划
      * 推广计划状态变更消息接口：`client.MarketingActivityDeliveryChanged()`
      * 创建推广计划接口：`client.MarketingActivityDeliveryCreate()`
      * 查询推广计划接口：`client.MarketingActivityDeliveryQuery()`
      * 停止推广计划接口：`client.MarketingActivityDeliveryStop()` 
      * 营销图片资源上传接口：`client.MarketingMaterialImageUpload()`
* <font color='#027AFF' size='4'>营销产品</font>
  * 支付券
    * 创建支付券接口：TODO：https://opendocs.alipay.com/open/049d65
    * 激活支付券接口：TODO：https://opendocs.alipay.com/open/049d66
    * 查询支付券详情接口：TODO：https://opendocs.alipay.com/open/049d6g
    * 修改支付券基本信息接口：TODO：https://opendocs.alipay.com/open/049d67
    * 追加支付券预算接口：TODO：https://opendocs.alipay.com/open/049d68
    * 停止支付券接口：TODO：https://opendocs.alipay.com/open/049d69
    * 券活动创建通知接口：TODO：https://opendocs.alipay.com/open/049d6a
    * 券活动修改通知接口：TODO：https://opendocs.alipay.com/open/049d6b
    * 券活动停止通知接口：TODO：https://opendocs.alipay.com/open/049d6c
    * 券活动预算追加通知接口：TODO：https://opendocs.alipay.com/open/049d6d
  * 私域营销
    * 条件查询活动列表接口：TODO：https://opendocs.alipay.com/open/04fgw9
    * 活动领取咨询接口：TODO：https://opendocs.alipay.com/open/04fgwa
    * 查询活动详情接口：TODO：https://opendocs.alipay.com/open/04fgwb
    * 查询活动可用商户接口：TODO：https://opendocs.alipay.com/open/04fgwc
    * 查询活动可用小程序接口：TODO：https://opendocs.alipay.com/open/04fgwd
    * 查询活动可用门店接口：TODO：https://opendocs.alipay.com/open/04fgwe
    * 查询活动适用商品接口：TODO：https://opendocs.alipay.com/open/04fgwf
    * 条件查询用户券接口：TODO：https://opendocs.alipay.com/open/04fgwg
    * 查询用户券详情接口：TODO：https://opendocs.alipay.com/open/04fgwh
    * 订单优惠前置咨询接口：TODO：https://opendocs.alipay.com/open/04fgwi
    * 券领取通知接口：TODO：https://opendocs.alipay.com/open/049d6f
    * 券核销消息接口：TODO：https://opendocs.alipay.com/open/049d6e
    * 券过期消息接口：TODO：https://opendocs.alipay.com/open/04fgwj
  * 蚂蚁门店管理
    * 蚂蚁店铺创建接口: `client.AntMerchantShopCreate()`
    * 店铺保存审核通过消息接口：TODO：https://opendocs.alipay.com/open/04fgwm
    * 店铺保存拒绝消息：TODO：https://opendocs.alipay.com/open/04fgwn
    * 修改蚂蚁店铺接口: `client.AntMerchantShopModify()`
    * 商户申请单查询接口: `client.AntMerchantOrderQuery()`
    * 店铺分页查询接口：TODO：https://opendocs.alipay.com/open/04fgwq
    * 店铺查询接口：`client.AntMerchantShopQuery()`
    * 蚂蚁店铺关闭接口：`client.AntMerchantShopClose()`
    * 图片上传接口：TODO：https://opendocs.alipay.com/open/04fgwt
    * 商户mcc信息查询接口：TODO：https://opendocs.alipay.com/open/04fgwu
    * 店铺增加收单账号接口：TODO：https://opendocs.alipay.com/open/54b69b89_ant.merchant.expand.shop.receiptaccount.save
  * 商家券 2.0
    * 活动创建
      * 创建商家券活动接口：TODO：https://opendocs.alipay.com/open/7ad3a7bf_alipay.marketing.activity.ordervoucher.create
      * 同步商家券券码接口：TODO：https://opendocs.alipay.com/open/7ed0450d_alipay.marketing.activity.ordervoucher.codedeposit
      * 券活动创建通知接口：TODO：https://opendocs.alipay.com/open/f28e36e2_alipay.marketing.activity.message.created
    * 活动修改
      * 修改商家券活动基本信息接口：TODO：https://opendocs.alipay.com/open/528f83f6_alipay.marketing.activity.ordervoucher.modify
      * 券活动修改通知接口：TODO：https://opendocs.alipay.com/open/f80939f9_alipay.marketing.activity.message.modified
    * 活动停止 
      * 停止商家券活动接口：TODO：https://opendocs.alipay.com/open/16803efe_alipay.marketing.activity.ordervoucher.stop
      * 券活动停止通知接口：TODO：https://opendocs.alipay.com/open/a055ad21_alipay.marketing.activity.message.stopped
    * 预算追加 
      * 修改商家券活动发券数量上限接口：TODO：https://opendocs.alipay.com/open/4e2acff5_alipay.marketing.activity.ordervoucher.append
      * 券活动预算追加通知接口：TODO：https://opendocs.alipay.com/open/d4697d94_alipay.marketing.activity.message.appended
    * 活动发放
      * 活动领取咨询接口：TODO：https://opendocs.alipay.com/open/cdf2e908_alipay.marketing.activity.consult
    * 优惠券核销
      * 同步券核销状态接口：TODO：https://opendocs.alipay.com/open/3ffce87f_alipay.marketing.activity.ordervoucher.use
      * 取消券核销状态接口：TODO：https://opendocs.alipay.com/open/4682759b_alipay.marketing.activity.ordervoucher.refund
    * 活动查询
      * 查询商家券活动接口：TODO：https://opendocs.alipay.com/open/51f5946e_alipay.marketing.activity.ordervoucher.query
      * 查询活动详情接口：TODO：https://opendocs.alipay.com/open/f76b603f_alipay.marketing.activity.query
      * 统计商家券券码数量接口：TODO：https://opendocs.alipay.com/open/f6e49e82_alipay.marketing.activity.ordervoucher.codecount
      * 条件查询活动列表接口：TODO：https://opendocs.alipay.com/open/55aa8ee8_alipay.marketing.activity.batchquery
      * 条件查询用户券接口：TODO：https://opendocs.alipay.com/open/300dff38_alipay.marketing.activity.user.batchqueryvoucher
      * 查询用户券详情接口：TODO：https://opendocs.alipay.com/open/abffd3f6_alipay.marketing.activity.user.queryvoucher
      * 券领取通知接口：TODO：https://opendocs.alipay.com/open/0ed08444_alipay.marketing.activity.message.received
      * 券核销消息接口：TODO：https://opendocs.alipay.com/open/6a0f4680_alipay.marketing.activity.message.used
      * 券过期消息接口：TODO：https://opendocs.alipay.com/open/ac03373c_alipay.marketing.activity.message.expired
      * 查询活动可用小程序接口：TODO：https://opendocs.alipay.com/open/634212b2_alipay.marketing.activity.app.batchquery
      * 查询活动可用门店接口：TODO：https://opendocs.alipay.com/open/cf4e0873_alipay.marketing.activity.shop.batchquery
      * 查询活动适用商品接口：TODO：https://opendocs.alipay.com/open/1bee5767_alipay.marketing.activity.goods.batchquery
  * 商家会员卡
    * 基础功能
      * 会员卡模板创建接口：TODO：https://opendocs.alipay.com/open/b2854ad3_alipay.marketing.card.template.create
      * 会员卡开通，获取会员卡信息接口：TODO：https://opendocs.alipay.com/open/03sx83
      * 会员卡更新接口：TODO：https://opendocs.alipay.com/open/03sx89
    * 扩展功能 
      * 会员卡开卡表单模板配置接口：TODO：https://opendocs.alipay.com/open/03sx82
      * 会员卡开卡结果通知接口：TODO：https://opendocs.alipay.com/open/03sx84
      * 会员卡模板修改接口：TODO：https://opendocs.alipay.com/open/e3227c82_alipay.marketing.card.template.modify
      * 会员卡模板查询接口：TODO：https://opendocs.alipay.com/open/690f3d16_alipay.marketing.card.template.query
      * 会员卡查询接口：TODO：https://opendocs.alipay.com/open/03sx88
      * 会员卡删卡接口：TODO：https://opendocs.alipay.com/open/03sx8a
      * 会员卡消息通知接口：TODO：https://opendocs.alipay.com/open/06ruek
      * 上传门店照片和视频接口：TODO：https://opendocs.alipay.com/open/03sx81
  * 营销活动送红包 
    * 创建现金活动接口：TODO：https://opendocs.alipay.com/open/029yy9
    * 触发现金红包活动接口：TODO：https://opendocs.alipay.com/open/029yya
    * 更改现金活动状态接口：TODO：https://opendocs.alipay.com/open/029yyb
    * 现金活动列表查询接口：TODO：https://opendocs.alipay.com/open/02a1f9
    * 现金活动详情查询接口：TODO：https://opendocs.alipay.com/open/02a1fa
  * 红包
    * 资金转账页面支付接口: `client.FundTransPagePay()`
    * 现金红包无线支付接口: `client.FundTransAppPay()`
    * 单笔转账接口：`client.FundTransUniTransfer()`
    * 转账业务单据查询接口：`client.FundTransCommonQuery()`
    * 资金单据状态变更通知接口：TODO：https://opendocs.alipay.com/open/02byvf
    * 资金退回成功通知接口：TODO：https://opendocs.alipay.com/open/02byvg
    * 资金退回接口: `client.FundTransRefund()`
  * 棋盘密云
    * 接口上传人群
      * 上传创建人群接口：TODO：https://opendocs.alipay.com/open/e93d9a54_alipay.merchant.qipan.crowd.create
      * 人群中追加用户接口：TODO：https://opendocs.alipay.com/open/04330914_alipay.merchant.qipan.crowduser.add
      * 人群中删除用户接口：TODO：https://opendocs.alipay.com/open/80646633_alipay.merchant.qipan.crowduser.delete
    * 标签圈选创建人群
      * 查询圈选标签列表接口：TODO：https://opendocs.alipay.com/open/8e411cec_alipay.marketing.qipan.crowdtag.query
      * 标签圈选创建人群接口：TODO：https://opendocs.alipay.com/open/cbf0efa4_alipay.marketing.qipan.crowdwithtag.create
      * 标签圈选预估人群规模接口：TODO：https://opendocs.alipay.com/open/adb5dd04_alipay.marketing.qipan.crowdwithtag.query
    * 管理人群
      * 查询人群列表接口：TODO：https://opendocs.alipay.com/open/b7f3caec_alipay.merchant.qipan.crowd.batchquery
      * 查询人群详情接口：TODO：https://opendocs.alipay.com/open/be384367_alipay.merchant.qipan.crowd.query
      * 修改人群接口：TODO：https://opendocs.alipay.com/open/6c8ebb31_alipay.merchant.qipan.crowd.modify
    * 数据洞察 
      * 看板分析接口：TODO：https://opendocs.alipay.com/open/aa5066aa_alipay.merchant.qipan.board.query
      * 画像分析接口：TODO：https://opendocs.alipay.com/open/46a75f4e_alipay.merchant.qipan.insight.query
      * 行为分析接口：TODO：https://opendocs.alipay.com/open/56735ac5_alipay.merchant.qipan.behavior.query
      * 趋势分析接口：TODO：https://opendocs.alipay.com/open/8ee7795a_alipay.merchant.qipan.trend.query
      * 常住省市查询接口：TODO：https://opendocs.alipay.com/open/f7d99821_alipay.merchant.qipan.insightcity.query
* <font color='#027AFF' size='4'>资金产品</font>
  * 转账到支付宝账户
    * 支付宝资金账户资产查询接口：`client.FundAccountQuery()`
    * 申请电子回单(incubating)接口：TODO：https://opendocs.alipay.com/open/1aad1956_alipay.data.bill.ereceipt.apply
    * 查询电子回单状态(incubating)接口：TODO：https://opendocs.alipay.com/open/30b94a2f_alipay.data.bill.ereceipt.query
    * 查询转账订单接口: `client.FundTransOrderQuery()`
    * 批次下单接口: `client.FundBatchCreate()`
    * 批量转账关单接口: `client.FundBatchClose()`
    * 批量转账明细查询接口: `client.FundBatchDetailQuery()`
    * 资金收款账号绑定关系查询: `client.FundTransPayeeBindQuery()`
* <font color='#027AFF' size='4'>会员产品</font>
  * 人脸验证
    * 人脸核身
      * APP人脸核身初始化接口：TODO：https://opendocs.alipay.com/open/04jg6r
      * APP人脸核身结果查询接口：TODO：https://opendocs.alipay.com/open/04jg6s
      * H5人脸核身初始化接口：TODO：https://opendocs.alipay.com/open/02zloa
      * H5人脸核身开始认证接口：TODO：https://opendocs.alipay.com/open/02zlob
      * H5人脸核身查询记录接口：TODO：https://opendocs.alipay.com/open/02zloc
      * 纯服务端人脸核身接口：TODO：https://opendocs.alipay.com/open/04pxq6
    * 活体检测
      * 人脸检测初始化接口：TODO：https://opendocs.alipay.com/open/03nisu
      * 人脸检测结果数据查询接口：TODO：https://opendocs.alipay.com/open/03nisv
    * OCR文字识别
      * 服务端OCR接口：TODO：https://opendocs.alipay.com/open/05ut8h
      * App端OCR初始化接口：TODO：https://opendocs.alipay.com/open/043ksf
      * 文字识别OCR接口：TODO：https://opendocs.alipay.com/open/0776c2cb_datadigital.fincloud.generalsaas.ocr.common.detect
  * 获取会员信息
    * 支付宝会员授权信息查询接口：`client.UserInfoShare()`
    * 换取授权访问令牌接口：`client.SystemOauthToken()`
    * 用户授权取消消息接口：TODO：https://opendocs.alipay.com/open/b4f0d4cf_alipay.open.auth.userauth.cancelled
    * 用户授权关系查询接口：TODO：https://opendocs.alipay.com/open/6b97edd1_alipay.open.auth.userauth.relationship.query
    * 查询解除授权明细接口：TODO：https://opendocs.alipay.com/open/77e7fec5_alipay.user.deloauth.detail.query
  * 支付宝身份验证
    * 身份认证记录查询: `client.UserCertifyOpenQuery()`
    * 身份认证初始化服务接口: `client.UserCertifyOpenInit()`
    * 身份认证开始认证: `client.UserCertifyOpenCertify()`
  * APP支付宝登录
    * 用户登录授权接口：`client.UserInfoAuth()`
* <font color='#027AFF' size='4'>信用产品</font>
  * 芝麻GO
    * 芝麻GO签约预创单接口：`client.ZhimaCreditPeZmgoPreorderCreate()`
    * 芝麻GO页面签约接口：TODO：https://opendocs.alipay.com/open/03u934
    * 芝麻GO业务签约变更商户通知接口：TODO：https://opendocs.alipay.com/open/03ums1
    * 商家芝麻GO累计数据回传接口：`client.ZhimaMerchantZmgoCumulateSync()`
    * 商家芝麻GO累计数据查询接口：`client.ZhimaMerchantZmgoCumulateQuery()`
    * 芝麻GO业务协议变更商户通知：TODO：https://opendocs.alipay.com/open/03up3c
    * 芝麻GO结算申请接口：TODO：https://opendocs.alipay.com/open/03usxk
    * 芝麻GO业务结算变更商户通知接口：TODO：https://opendocs.alipay.com/open/03ub1d
    * 芝麻GO结算退款接口：`client.ZhimaCreditPeZmgoSettleRefund()`
    * 芝麻Go协议查询接口：`client.ZhimaCreditPeZmgoAgreementQuery()`
    * 芝麻GO协议解约接口：`client.ZhimaCreditPeZmgoAgreementUnsign()`
    * 商户创建芝麻GO模板接口：TODO：https://opendocs.alipay.com/open/03uq08
    * 芝麻GO模板查询接口：TODO：https://opendocs.alipay.com/open/04m8ci
    * 芝麻GO用户数据回传: `client.ZhimaCreditPeZmgoCumulationSync()`
    * 芝麻GO签约关单: `client.ZhimaCreditPeZmgoBizoptClose()`
    * 芝麻GO解冻接口: `client.ZhimaCreditPeZmgoSettleUnfreeze()`
    * 芝麻GO支付下单链路签约申请: `client.ZhimaCreditPeZmgoPaysignApply()`
    * 芝麻GO支付下单链路签约确认: `client.ZhimaCreditPeZmgoPaysignConfirm()`
  * 芝麻先享
    * 服务开通
      * 信用服务开通/授权接口：TODO：https://opendocs.alipay.com/open/03uloz
      * 查询服务开通/授权信息接口：TODO：https://opendocs.alipay.com/open/03ulp0
      * 服务开通/授权状态变更通知：TODO：https://opendocs.alipay.com/open/50885f82_zhima.credit.payafteruse.creditagreement.changed
    * 信用下单 
      * 芝麻先享信用服务下单（免用户确认场景）接口：TODO：https://opendocs.alipay.com/open/03ulpo
      * 芝麻先享信用服务下单（用户确认场景）接口：TODO：https://opendocs.alipay.com/open/03ulpp
      * 信用服务订单查询接口：TODO：https://opendocs.alipay.com/open/03vtet
      * 结束信用服务订单接口：TODO：https://opendocs.alipay.com/open/03vteu
      * 芝麻先享信用服务订单状态变更通知接口：TODO：https://opendocs.alipay.com/open/6c3a11df_zhima.credit.payafteruse.creditbizorder.changed
  * 扣款
    * 统一收单交易订单支付接口：TODO：https://opendocs.alipay.com/open/2e81a85f_alipay.trade.order.pay
  * 芝麻免押
    * 创建免押订单接口：`client:FundAuthOrderAppFreeze()` 
    * 完结免押订单接口：`client.FundAuthOrderUnfreeze()`
    * 查询免押订单接口：`client.FundAuthOperationDetailQuery()`
    * 取消免押订单接口：`client.FundAuthOperationCancel()`
  * 芝麻身份信息验证
    * 信用服务开通/授权接口：TODO：https://opendocs.alipay.com/open/63612668_zhima.credit.payafteruse.creditagreement.sign
    * 职得身份认证查询接口：TODO：https://opendocs.alipay.com/open/351177b5_zhima.customer.jobworth.authentication.query
    * 职得工作证信息匹配度查询: `client.ZhimaCustomerJobworthAdapterQuery()`
    * 职得工作证外部渠道应用数据回流: `client.ZhimaCustomerJobworthSceneUse()`
    * 身份验真预咨询服务接口：TODO：https://opendocs.alipay.com/open/16f72e43_zhima.customer.jobworth.authentication.preconsult
* <font color='#027AFF' size='4'>安全产品</font>
  * 交易安全防护
    * 商户数据同步：TODO：https://opendocs.alipay.com/open/02qth4
* <font color='#027AFF' size='4'>其他产品</font>
  * 花芝轻会员结算申请: `client.PcreditHuabeiAuthSettleApply()`
  * NFC用户卡信息同步: `client.CommerceTransportNfccardSend()`
  * 广告投放数据查询: `client.DataDataserviceAdDataQuery()`
  * 航司电话订票待申请接口: `client.CommerceAirCallcenterTradeApply()`
  * 网商银行全渠道收单业务订单创建: `client.PaymentTradeOrderCreate()`
  * 聚合支付订单咨询服务: `client.KoubeiTradeOrderAggregateConsult()`
  * 口碑订单预下单: `client.KoubeiTradeOrderPrecreate()`
  * 口碑商品交易购买接口: `client.KoubeiTradeItemorderBuy()`
  * 口碑订单预咨询: `client.KoubeiTradeOrderConsult()`
  * 口碑商品交易退货接口: `client.KoubeiTradeItemorderRefund()`
  * 口碑商品交易查询接口: `client.KoubeiTradeItemorderQuery()`
  * 码商发码成功回调接口: `client.KoubeiTradeTicketTicketcodeSend()`
  * 口碑凭证延期接口: `client.KoubeiTradeTicketTicketcodeDelay()`
  * 口碑凭证码查询: `client.KoubeiTradeTicketTicketcodeQuery()`
  * 口碑凭证码撤销核销: `client.KoubeiTradeTicketTicketcodeCancel()`
  * 蚂蚁店铺创建咨询: `client.AntMerchantShopConsult()`
  * 申请权益发放: `client.CommerceBenefitApply()`
  * 权益核销: `client.CommerceBenefitVerify()`
  * 还款账单查询: `client.TradeRepaybillQuery()`
* <font color='#027AFF' size='4'>会员</font>
  * 协议由普通通用代扣协议产品转移到周期扣协议产品: `client.UserAgreementTransfer()`
  * 通用当面付二阶段接口: `client.UserTwostageCommonUse()`
  * 芝麻企业征信基于身份的协议授权: `client.UserAuthZhimaorgIdentityApply()`
  * 查询是否在支付宝公益捐赠的接口: `client.UserCharityRecordexistQuery()`
  * 集分宝发放接口: `client.UserAlipaypointSend()`
  * isv 会员CRM数据回流: `client.MemberDataIsvCreate()`
  * 询家人信息档案(选人授权)组件已选的家人档案信息: `client.UserFamilyArchiveQuery()`
  * 初始化家人信息档案(选人授权)组件: `client.UserFamilyArchiveInitialize()`
  * 实名证件信息比对验证预咨询: `client.UserCertdocCertverifyPreconsult()`
  * 实名证件信息比对验证咨询: `client.UserCertdocCertverifyConsult()`
  * 初始化家庭芝麻GO共享组件: `client.UserFamilyShareZmgoInitialize()`
  * 数字分行银行码明细数据查询: `client.UserDtbankQrcodedataQuery()`
  * 查询集分宝预算库详情: `client.UserAlipaypointBudgetlibQuery()`
* <font color='#027AFF' size='4'>营销</font>
  * 小程序生成推广二维码接口：`client.OpenAppQrcodeCreate()`
* <font color='#027AFF' size='4'>工具类</font>
  * 换取应用授权令牌：`client.OpenAuthTokenApp()`
  * 应用支付宝公钥证书下载：`client.PublicCertDownload()`
* <font color='#027AFF' size='4'>芝麻信用</font>
  * 芝麻企业信用信用评估初始化: `client.ZhimaCreditEpSceneRatingInitialize()`
  * 信用服务履约同步: `client.ZhimaCreditEpSceneFulfillmentSync()`
  * 加入信用服务: `clinet.ZhimaCreditEpSceneAgreementUse()`
  * 取消信用服务: `client.ZhimaCreditEpSceneAgreementCancel()`
  * 信用服务履约同步(批量): `client.ZhimaCreditEpSceneFulfillmentlistSync()`

### 支付宝公共 API

* `alipay.IsBizError()` => 判断并捕获业务错误 BizError
* `alipay.GetCertSN()` => 获取证书SN号（app_cert_sn、alipay_cert_sn）
* `alipay.GetRootCertSN()` => 获取证书SN号（alipay_root_cert_sn）
* `alipay.GetRsaSign()` => 获取支付宝参数签名（参数sign值）
* `alipay.SystemOauthToken()` => 换取授权访问令牌（得到access_token，user_id等信息）
* `alipay.FormatPrivateKey()` => 格式化应用私钥
* `alipay.FormatPublicKey()` => 格式化支付宝公钥
* `alipay.FormatURLParam()` => 格式化支付宝请求URL参数
* `alipay.ParseNotifyToBodyMap()` => 解析支付宝支付异步通知的参数到BodyMap
* `alipay.ParseNotifyByURLValues()` => 通过 url.Values 解析支付宝支付异步通知的参数到BodyMap
* `alipay.VerifySign()` => 支付宝异步通知参数验签
* `alipay.VerifySignWithCert()` => 支付宝异步通知参数验签（证书方式）
* `alipay.VerifySyncSign()` => 支付宝同步返回参数验签
* `alipay.DecryptOpenDataToStruct()` => 解密支付宝开放数据到 结构体
* `alipay.DecryptOpenDataToBodyMap()` => 解密支付宝开放数据到 BodyMap
* `alipay.MonitorHeartbeatSyn()` => 验签接口
