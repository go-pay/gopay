## 支付宝

> #### 支付宝v3版本已支持，接口持续完善中。

> #### 因支付宝接口太多，如没实现的接口，还请开发者自行调用 `client.DoAliPayAPISelfV3()`方法实现！请参考 `client_test.go` 内的 `TestDoAliPayAPISelfV3()` 方法

> #### 希望有时间的伙伴儿Fork完后，补充并提交Pull Request，一起完善支付宝各个类别下的接口到相应的go文件中

- 已实现API列表附录：[API 列表附录](https://github.com/go-pay/gopay/blob/main/doc/alipay_v3.md#%E9%99%84%E5%BD%95)

- 开发文档（V3版）：[开发文档（V3版）](https://opendocs.alipay.com/open-v3) 

- 接口规则：[接口规则](https://opendocs.alipay.com/open-v3/053sd1)

- 支付宝RSA秘钥生成文档：[生成RSA密钥](https://opendocs.alipay.com/common/02kipl) （推荐使用 RSA2）

- 沙箱环境(新) 使用说明：[新版沙箱文档](https://opendocs.alipay.com/common/02kkv7)

---

### 1、初始化支付宝客户端并做配置

> 具体API使用介绍，请参考 `gopay/alipay/v3/client_test.go`

>  

```go
import (
    "github.com/go-pay/gopay/alipay/v3"
    "github.com/go-pay/xlog"
)

// 初始化支付宝客V3户端
// appid：应用ID
// privateKey：应用私钥，支持PKCS1和PKCS8
// isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。
client, err := alipay.NewClientV3("2016091200494382", privateKey, false)
if err != nil {
    xlog.Error(err)
    return
}

// 设置自定义配置（如需要）
//client.
//	SetAppAuthToken("app_auth_token").    // 设置授权token
//	SetBodySize().                        // 自定义配置http请求接收返回结果body大小，默认 10MB，没有特殊需求，可忽略此配置
//	SetRequestIdFunc().                   // 设置自定义RequestId生成方法
//	SetAESKey("KvKUTqSVZX2fUgmxnFyMaQ==") // 设置biz_content加密KEY，设置此参数默认开启加密（目前不可用）

// 打开Debug开关，输出日志，默认关闭
client.DebugSwitch = gopay.DebugOn

// 传入证书内容
err := client.SetCert("appPublicCert.crt bytes", "alipayRootCert bytes", "alipayPublicCert.crt bytes")
```

### 2、API 方法调用及入参

> 具体参数请根据不同接口查看：[支付宝V3版API接口文档](https://opendocs.alipay.com/open-v3)

> ★入参 BodyMap中，支持如下公共参数在当次请求中自定义设置：`alipay-app-auth-token`

- 统一收单线下交易预创建 - 示例

```go
import (
    "github.com/go-pay/gopay"
    "github.com/go-pay/gopay/pkg/js"
    "github.com/go-pay/util"
    "github.com/go-pay/xlog"
    "github.com/go-pay/gopay/alipay/v3"
)

// 请求参数
bm := make(gopay.BodyMap)
bm.Set("subject", "预创建创建订单").
    Set("out_trade_no", util.RandomString(32)).
    Set("total_amount", "0.01")
    Set(alipay.HeaderAppAuthToken, "i_am_app_auth_token") // 如果需要，可以设置自定义应用授权

// 创建订单
aliRsp, err := client.TradePrecreate(ctx, bm)
if err != nil {
    xlog.Errorf("client.TradePrecreate(), err:%v", err)
    return
}
xlog.Warnf("aliRsp:%s", js.Marshal(aliRsp))

if aliRsp.StatusCode != Success {
    xlog.Errorf("aliRsp.StatusCode:%d", aliRsp.StatusCode)
    return
}
xlog.Warnf("aliRsp.QrCode:", aliRsp.QrCode)
xlog.Warnf("aliRsp.OutTradeNo:", aliRsp.OutTradeNo)
```

- 自定义接口调用 - 示例

```go
import (
    "github.com/go-pay/gopay"
    "github.com/go-pay/gopay/pkg/js"
    "github.com/go-pay/util"
    "github.com/go-pay/xlog"
)

// 请求参数
bm := make(gopay.BodyMap)
bm.Set("subject", "预创建创建订单").
    Set("out_trade_no", util.RandomString(32)).
    Set("total_amount", "0.01")

rsp := new(struct {
    OutTradeNo string `json:"out_trade_no"`
    QrCode     string `json:"qr_code"`
})
// 创建订单
res, err := client.DoAliPayAPISelfV3(ctx, alipay.MethodPost, "/v3/alipay/trade/precreate", bm, rsp)
if err != nil {
    xlog.Errorf("client.TradePrecreate(), err:%v", err)
    return
}
xlog.Warnf("aliRsp:%s", js.Marshal(rsp))
if res.StatusCode != Success {
    xlog.Errorf("aliRsp.StatusCode:%d", res.StatusCode)
    return
}
```

### 3、异步通知参数解析和验签Sign、异步通知返回（复用非V3版方式）

> 异步通知请求参数需要先解析，解析出来的结构体或BodyMap再验签（此处需要注意，`http.Request.Body` 只能解析一次，如果需要解析前调试，请处理好Body复用问题）

[Gin Web框架（推荐）](https://github.com/gin-gonic/gin)

[Echo Web框架](https://github.com/labstack/echo)

> 支付宝支付后的异步通知验签文档：[支付结果通知](https://opendocs.alipay.com/common/02mse7)

- 异步通知验签（复用非V3版方式）

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
ok, err = alipay.VerifySignWithCert([]byte("alipayPublicCert.crt content"), notifyReq)

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

---

## 附录：

### 支付宝支付 API

* 支付宝接口自行实现方法：`client.DoAliPayAPISelfV3()`
* <font color='#027AFF' size='4'>支付</font>
  * 统一收单交易支付接口：`client.TradePay()`
  * 统一收单交易查询：`client.TradeQuery()`
  * 统一收单交易退款接口：`client.TradeRefund()`
  * 统一收单交易退款查询：`client.TradeFastPayRefundQuery()`
  * 统一收单交易撤销接口：`client.TradeCancel()`
  * 统一收单交易关闭接口：`client.TradeClose()`
  * 查询对账单下载地址：`client.DataBillDownloadUrlQuery()`
  * 统一收单线下交易预创建：`client.TradePrecreate()`
  * 支付宝APP支付（v2版本）：`client.TradeAppPay()`
  * 统一收单下单并支付页面接口（v2版本）：`client.TradePagePay()`
  * 统一收单交易创建接口：`client.TradeCreate()`
  * 支付宝订单信息同步接口：`client.TradeOrderInfoSync()`
  * 资金授权操作查询接口：`client.FundAuthOperationDetailQuery()`
  * 资金授权冻结接口：`client.FundAuthOrderFreeze()`
  * 资金授权解冻接口：`client.FundAuthOrderUnfreeze()`
  * 资金授权发码接口：`client.FundAuthOrderVoucherCreate()`
  * 刷脸支付初始化接口：`client.ZolozAuthenticationSmilepayInitialize()`
  * 查询刷脸结果信息接口：`client.ZolozAuthenticationCustomerFtokenQuery()`
* <font color='#027AFF' size='4'>商家扣款</font>
  * 支付宝个人代扣协议查询接口：`client.UserAgreementQuery()`
  * 支付宝个人代扣协议解约接口：`client.UserAgreementPageUnSign()`
* <font color='#027AFF' size='4'>商家分账</font>
  * 分账关系绑定接口：`client.TradeRelationBind()`
  * 分账关系解绑：`client.TradeRelationUnbind()`
  * 分账关系查询：`client.TradeRelationBatchQuery()`
  * 分账比例查询：`client.TradeRoyaltyRateQuery()`
  * 统一收单交易结算接口：`client.TradeOrderSettle()`
  * 分账剩余金额查询：`client.TradeOrderOnSettleQuery()`
* <font color='#027AFF' size='4'>商家转账</font>
  * 支付宝资金账户资产查询接口：`client.FundAccountQuery()`
  * 转账额度查询接口：`client.FundQuotaQuery()`
  * 单笔转账接口：`client.FundTransUniTransfer()`
  * 申请电子回单(incubating)：`client.DataBillEreceiptApply()`
  * 查询电子回单状态(incubating)：`client.DataBillEreceiptQuery()`
  * 转账业务单据查询接口：`client.FundTransCommonQuery()`
  * 多步转账创建并支付：`client.FundTransMultistepTransfer()`
  * 多步转账查询接口：`client.FundTransMultistepQuery()`
* <font color='#027AFF' size='4'>会员</font>
  * 换取授权访问令牌：`client.SystemOauthToken()`
  * 身份认证记录查询：`client.UserCertifyOpenQuery()`
  * 身份认证初始化服务：`client.UserCertifyOpenInitialize()`
  * 支付宝会员授权信息查询接口：`client.UserInfoShare()`
  * 用户授权关系查询：`client.UserAuthRelationshipQuery()`
  * 查询解除授权明细：`client.UserDelOauthDetailQuery()`
* <font color='#027AFF' size='4'>人脸认证</font>
  * 人脸核身初始化：`client.FaceVerificationInitialize()`
  * 人脸核身结果查询：`client.FaceVerificationQuery()`
  * 跳转支付宝人脸核身初始化：`client.FaceCertifyInitialize()`
  * 跳转支付宝人脸核身开始认证：`client.FaceCertifyVerify()`
  * 跳转支付宝人脸核身查询记录：`client.FaceCertifyQuery()`
  * 纯服务端人脸核身：`client.FaceSourceCertify()`
  * 活体检测初始化：`client.FaceCheckInitialize()`
  * 活体检测结果查询：`client.FaceCheckQuery()`
  * 身份证二要素核验：`client.IDCardTwoMetaCheck()`
  * 银行卡核验：`client.BankCardCheck()`
  * 手机号三要素核验简版：`client.MobileThreeMetaSimpleCheck()`
  * 手机号三要素核验详版：`client.MobileThreeMetaDetailCheck()`
  * 服务端OCR：`client.OcrServerDetect()`
  * App端OCR初始化：`client.OcrMobileInitialize()`
* <font color='#027AFF' size='4'>推广计划</font>
  * 创建推广计划：`client.MarketingActivityDeliveryCreate()`
  * 查询推广计划：`client.MarketingActivityDeliveryQuery()`
  * 停止推广计划：`client.MarketingActivityDeliveryStop()`
* <font color='#027AFF' size='4'>营销</font>
  * 创建现金活动：`client.MarketingCampaignCashCreate()`
  * 触发现金红包：`client.MarketingCampaignCashTrigger()`
  * 更改现金活动状态：`client.MarketingCampaignCashStatusModify()`
  * 现金活动列表查询：`client.MarketingCampaignCashListQuery()`
  * 现金活动详情查询：`client.MarketingCampaignCashDetailQuery()`
  * 创建商家券活动：`client.MarketingActivityOrderVoucherCreate()`
  * 同步商家券券码：`client.MarketingActivityOrderVoucherCodeDeposit()`
  * 修改商家活动券基本信息：`client.MarketingActivityOrderVoucherModify()`
  * 停止商家活动券：`client.MarketingActivityOrderVoucherStop()`
  * 修改商家活动券发券数量上限：`client.MarketingActivityOrderVoucherAppend()`
  * 同步券核销状态：`client.MarketingActivityOrderVoucherUse()`
  * 取消券核销状态：`client.MarketingActivityOrderVoucherRefund()`
  * 活动领取咨询接口：`client.MarketingActivityConsult()`
  * 查询商家券活动：`client.MarketingActivityOrderVoucherQuery()`
  * 查询活动详情：`client.MarketingActivityQuery()`
  * 统计商家券券码数量：`client.MarketingActivityOrderVoucherCodeCount()`
  * 条件查询活动列表：`client.MarketingActivityBatchQuery()`
  * 条件查询用户券：`client.MarketingActivityQueryUserBatchQueryVoucher()`
  * 查询用户券详情：`client.MarketingActivityQueryUserQueryVoucher()`
  * 查询活动可用小程序：`client.MarketingActivityQueryAppBatchQuery()`
  * 查询活动可用门店：`client.MarketingActivityQueryShopBatchQuery()`
  * 查询活动适用商品：`client.MarketingActivityQueryGoodsBatchQuery()`
  * 蚂蚁店铺创建：`client.AntMerchantShopCreate()`
  * 店铺查询接口：`client.AntMerchantShopQuery()`
  * 修改蚂蚁店铺：`client.AntMerchantShopModify()`
  * 蚂蚁店铺关闭：`client.AntMerchantShopClose()`
  * 商户申请单查询：`client.AntMerchantOrderQuery()`
  * 店铺分页查询接口：`client.AntMerchantShopPageQuery()`
  * 图片上传：`client.AntMerchantExpandIndirectImageUpload()`
  * 商户mcc信息查询：`client.AntMerchantExpandMccQuery()`
  * 店铺增加收单账号：`client.AntMerchantExpandShopReceiptAccountSave()`
  * 会员卡模板创建：`client.MarketingCardTemplateCreate()`
  * 会员卡模板查询接口：`client.MarketingCardTemplateQuery()`
  * 会员卡模板修改：`client.MarketingCardTemplateModify()`
  * 会员卡开卡表单模板配置：`client.MarketingCardFormTemplateSet()`
  * 会员卡查询：`client.MarketingCardQuery()`
  * 会员卡更新：`client.MarketingCardUpdate()`
  * 会员卡删卡：`client.MarketingCardDelete()`
  * 会员卡消息通知：`client.MarketingCardMessageNotify()`
  * 上传门店照片和视频接口：`client.OfflineMaterialImageUpload()`
  * 资金退回接口：`client.FundTransRefund()`
* <font color='#027AFF' size='4'>图片素材</font>
  * 营销图片资源上传接口：`client.MarketingMaterialImageUpload()`
* <font color='#027AFF' size='4'>信用</font>
  * 芝麻GO签约预创单：`client.ZmGoPreorderCreate()`
  * 商家芝麻GO累计数据回传接口：`client.ZmGoCumulateSync()`
  * 商家芝麻GO累计数据查询接口：`client.ZmGoCumulateQuery()`
  * 芝麻GO结算申请：`client.ZmGoSettleApply()`
  * 芝麻GO结算退款：`client.ZmGoSettleRefund()`
  * 芝麻Go协议查询接口：`client.ZmGoAgreementQuery()`
  * 芝麻GO协议解约：`client.ZmGoAgreementQueryUnsign()`
  * 商户创建芝麻GO模板接口：`client.ZmGoTemplateCreate()`
  * 商家芝麻GO模板查询：`client.ZmGoTemplateQuery()`
* <font color='#027AFF' size='4'>广告</font>
  * 转化数据回传：`client.AdConversionUpload()`
  * 广告投放数据通用查询：`client.AdReportdataQuery()`
  * 自建推广页列表批量查询：`client.AdPromotepageBatchquery()`
  * 自建推广页留资数据查询：`client.AdPromotepageDownload()`
  * 任务广告完成状态查询接口：`client.XlightTaskQuery()`
  * 消费明细查询接口：`client.AdConsumehistoryQuery()`
  * 商品落地页信息创建或更新：`client.ProductLandinginfoCreateOrModify()`
  * 商品落地页信息查询：`client.ProductLandinginfoQuery()`
  * 广告代理商投放数据查询：`client.AdAgentreportdataQuery()`
* <font color='#027AFF' size='4'>小程序开发</font>
  * 小程序退回开发：`client.OpenMiniVersionAuditedCancel()`
  * 小程序灰度上架：`client.OpenMiniVersionGrayOnline()`
  * 小程序结束灰度：`client.OpenMiniVersionGrayCancel()`
  * 小程序上架：`client.OpenMiniVersionOnline()`
  * 小程序下架：`client.OpenMiniVersionOffline()`
  * 小程序回滚：`client.OpenMiniVersionRollback()`
  * 小程序删除版本：`client.OpenMiniVersionDelete()`
  * 小程序提交审核：`client.OpenMiniVersionAuditApply()`
  * 小程序基于模板上传版本：`client.OpenMiniVersionUpload()`
  * 查询使用模板的小程序列表：`client.OpenMiniTemplateUsageQuery()`
  * 小程序查询版本构建状态：`client.OpenMiniVersionBuildQuery()`
  * 小程序版本详情查询：`client.OpenMiniVersionDetailQuery()`
  * 小程序版本列表查询：`client.OpenMiniVersionListQuery()`
  * 小程序生成体验版：`client.OpenMiniExperienceCreate()`
  * 小程序体验版状态查询接口：`client.OpenMiniExperienceQuery()`
  * 小程序取消体验版：`client.OpenMiniExperienceCancel()`

### 支付宝公共 API
