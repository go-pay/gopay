## 微信v3

> #### 推荐使用v3接口，官方在v3接口实现未覆盖或gopay未开发的接口，还继续用v2接口，欢迎参与完善v3接口。

- 已实现API列表附录：[API 列表附录](https://github.com/go-pay/gopay/blob/main/doc/wechat_v3.md#%E9%99%84%E5%BD%95)

- 微信官方文档：[官方文档](https://pay.weixin.qq.com/wiki/doc/apiv3/index.shtml)

- 接口规则：[平台证书](https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_0.shtml)

- 接入规范：[最佳实践](https://pay.weixin.qq.com/wiki/doc/apiv3/Practices/chapter1_1_1.shtml)

- GoPay微信v2文档：[GoPay微信v2文档](https://github.com/go-pay/gopay/blob/main/doc/wechat_v2.md) （部分接口仅v2版本支持）

---

### 1、初始化微信v3客户端并做配置

> 注意：v3 版本接口持续增加中，不支持沙箱支付，测试请用1分钱测试法

> 具体API使用介绍，请参考 `gopay/wechat/v3/client_test.go`

```go
import (
    "github.com/go-pay/gopay/pkg/xlog"
    "github.com/go-pay/gopay/wechat/v3"
)

// NewClientV3 初始化微信客户端 v3
//	mchid：商户ID 或者服务商模式的 sp_mchid
// 	serialNo：商户证书的证书序列号
//	apiV3Key：apiV3Key，商户平台获取
//	privateKey：私钥 apiclient_key.pem 读取后的内容
client, err = wechat.NewClientV3(MchId, SerialNo, APIv3Key, PrivateKey)
if err != nil {
    xlog.Error(err)
    return
}

// 设置微信平台API证书和序列号（如开启自动验签，请忽略此步骤）
//client.SetPlatformCert([]byte(""), "")

// 启用自动同步返回验签，并定时更新微信平台API证书（开启自动验签时，无需单独设置微信平台API证书和序列号）
err = client.AutoVerifySign()
if err != nil {
    xlog.Error(err)
    return
}

// 打开Debug开关，输出日志，默认是关闭的
client.DebugSwitch = gopay.DebugOn
```

### 2、API 方法调用及入参

> 具体参数请根据不同接口查看：[微信支付V3的API字典概览](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/index.shtml)

- JSAPI下单 示例
```go
import (
    "github.com/go-pay/gopay"
)

expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
// 初始化 BodyMap
bm := make(gopay.BodyMap)
bm.Set("sp_appid", "sp_appid").
    Set("sp_mchid", "sp_mchid").
    Set("sub_mchid", "sub_mchid").
    Set("description", "测试Jsapi支付商品").
    Set("out_trade_no", tradeNo).
    Set("time_expire", expire).
    Set("notify_url", "https://www.fmm.ink").
    SetBodyMap("amount", func(bm gopay.BodyMap) {
        bm.Set("total", 1).
            Set("currency", "CNY")
    }).
    SetBodyMap("payer", func(bm gopay.BodyMap) {
        bm.Set("sp_openid", "asdas")
    })

wxRsp, err := client.V3TransactionJsapi(bm)
if err != nil {
    xlog.Error(err)
    return
}
if wxRsp.Code == Success {
    xlog.Debugf("wxRsp: %#v", wxRsp.Response)
    return
}
xlog.Errorf("wxRsp:%s", wxRsp.Error)
```

### 3、下单后，获取微信小程序支付、APP支付、JSAPI支付所需要的 pay sign

> 小程序调起支付API：[小程序调起支付API](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_5_4.shtml)

> APP调起支付API：[APP调起支付API](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_2_4.shtml)

> JSAPI调起支付API：[JSAPI调起支付API](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_4.shtml)

```go
// 小程序
applet, err := client.PaySignOfApplet("appid", "prepayid")
// app
app, err := client.PaySignOfApp("appid", "prepayid")
// jsapi
jsapi, err := client.PaySignOfJSAPI("appid", "prepayid")
```

### 4、同步返回参数验签Sign、异步通知参数解析和验签Sign、异步通知返回

> 异步通知请求参数需要先解析，解析出来的结构体或BodyMap再验签（此处需要注意，`http.Request.Body` 只能解析一次，如果需要解析前调试，请处理好Body复用问题）

[Gin Web框架（推荐）](https://github.com/gin-gonic/gin)

[Echo Web框架](https://github.com/labstack/echo)

- 同步返回验签，手动验签（如已开启自动验签，则无需手动验签操作）

```go
import (
    "github.com/go-pay/gopay/wechat/v3"
    "github.com/go-pay/gopay/pkg/xlog"
)

wxRsp, err := client.V3TransactionJsapi(bm)
if err != nil {
    xlog.Error(err)
    return
}
// wxPublicKey 通过 client.WxPublicKey() 获取
err = wechat.V3VerifySignByPK(wxRsp.SignInfo.HeaderTimestamp, wxRsp.SignInfo.HeaderNonce, wxRsp.SignInfo.SignBody, wxRsp.SignInfo.HeaderSignature, wxPublicKey)
if err != nil {
    xlog.Error(err)
    return
}
```

- 异步通知验签 及 敏感参数解密

```go
import (
    "github.com/go-pay/gopay/wechat/v3"
    "github.com/go-pay/gopay/pkg/xlog"
)

notifyReq, err := wechat.V3ParseNotify()
if err != nil {
    xlog.Error(err)
    return
}

// wxPublicKey 通过 client.WxPublicKey() 获取
err = notifyReq.VerifySignByPK(wxPublicKey)
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

### 5、微信v3 公共API（仅部分说明）

```go
import (
    "github.com/go-pay/gopay/wechat/v3"
)

// 获取微信平台证书和序列号信息，推荐使用后者
wechat.GetPlatformCerts()
 或
client.GetAndSelectNewestCert()

// 请求参数 敏感信息加密，推荐使用后者
wechat.V3EncryptText() 或 client.V3EncryptText()

// 返回参数 敏感信息解密，推荐使用后者
wechat.V3DecryptText() 或 client.V3DecryptText()

// 回调通知敏感信息解密
wechat.V3DecryptNotifyCipherText()
wechat.V3DecryptRefundNotifyCipherText()
wechat.V3DecryptCombineNotifyCipherText()
wechat.V3DecryptScoreNotifyCipherText()
...
```

---

## 附录：

### 微信支付v3 API

* <font color='#07C160' size='4'>基础支付</font>
    * APP下单：`client.V3TransactionApp()`
    * JSAPI/小程序下单：`client.V3TransactionJsapi()`
    * Native下单：`client.V3TransactionNative()`
    * H5下单：`client.V3TransactionH5()`
    * 查询订单：`client.V3TransactionQueryOrder()`
    * 关闭订单：`client.V3TransactionCloseOrder()`
* <font color='#07C160' size='4'>基础支付（服务商）</font>
    * APP下单：`client.V3PartnerTransactionApp()`
    * JSAPI/小程序下单：`client.V3PartnerTransactionJsapi()`
    * Native下单：`client.V3PartnerTransactionNative()`
    * H5下单：`client.V3PartnerTransactionH5()`
    * 查询订单：`client.V3PartnerQueryOrder()`
    * 关闭订单：`client.V3PartnerCloseOrder()`
* <font color='#07C160' size='4'>合单支付</font>
    * 合单APP下单：`client.V3CombineTransactionApp()`
    * 合单JSAPI/小程序下单：`client.V3CombineTransactionJsapi()`
    * 合单Native下单：`client.V3CombineTransactionNative()`
    * 合单H5下单：`client.V3CombineTransactionH5()`
    * 合单查询订单：`client.V3CombineQueryOrder()`
    * 合单关闭订单：`client.V3CombineCloseOrder()`
* <font color='#07C160' size='4'>退款</font>
    * 申请退款：`client.V3Refund()`
    * 查询单笔退款：`client.V3RefundQuery()`
* <font color='#07C160' size='4'>账单</font>
    * 申请交易账单：`client.V3BillTradeBill()`
    * 申请资金账单：`client.V3BillFundFlowBill()`
    * 申请特约商户资金账单：`client.V3BillEcommerceFundFlowBill()`
    * 下载账单：`client.V3BillDownLoadBill()`
* <font color='#07C160' size='4'>提现（服务商、电商）</font>
    * 特约商户余额提现/二级商户预约提现：`client.V3Withdraw()`
    * 查询特约商户提现状态/二级商户查询预约提现状态：`client.V3WithdrawStatus()`
    * 电商平台预约提现：`client.V3EcommerceWithdraw()`
    * 电商平台查询预约提现状态：`client.V3EcommerceWithdrawStatus()`
    * 按日下载提现异常文件：`client.V3WithdrawDownloadErrBill()`
* <font color='#07C160' size='4'>微信支付分（公共API）</font>
    * 创建支付分订单：`client.V3ScoreOrderCreate()`
    * 查询支付分订单：`client.V3ScoreOrderQuery()`
    * 取消支付分订单：`client.V3ScoreOrderCancel()`
    * 修改订单金额：`client.V3ScoreOrderModify()`
    * 完结支付分订单：`client.V3ScoreOrderComplete()`
    * 商户发起催收扣款：`client.V3ScoreOrderPay()`
    * 同步服务订单信息：`client.V3ScoreOrderSync()`
* <font color='#07C160' size='4'>微信支付分（免确认模式）</font>
    * 创单结单合并：`client.V3ScoreDirectComplete()`
* <font color='#07C160' size='4'>微信支付分（免确认预授权模式）</font>
    * 商户预授权：`client.V3ScorePermission()`
    * 查询用户授权记录（授权协议号）：`client.V3ScorePermissionQuery()`
    * 解除用户授权关系（授权协议号）：`client.V3ScorePermissionTerminate()`
    * 查询用户授权记录（openid）：`client.V3ScorePermissionOpenidQuery()`
    * 解除用户授权关系（openid）：`client.V3ScorePermissionOpenidTerminate()`
* <font color='#07C160' size='4'>微信先享卡</font>
    * 预受理领卡请求：`client.V3DiscountCardApply()`
    * 增加用户记录：`client.V3DiscountCardAddUser()`
    * 查询先享卡订单：`client.V3DiscountCardQuery()`
* <font color='#07C160' size='4'>支付即服务</font>
    * 服务人员注册：`client.V3SmartGuideReg()`
    * 服务人员分配：`client.V3SmartGuideAssign()`
    * 服务人员查询：`client.V3SmartGuideQuery()`
    * 服务人员信息更新：`client.V3SmartGuideUpdate()`
* <font color='#07C160' size='4'>智慧商圈</font>
    * 商圈积分同步：`client.V3BusinessPointsSync()`
    * 商圈积分授权查询：`client.V3BusinessAuthPointsQuery()`
* <font color='#07C160' size='4'>微信支付分停车服务</font>
    * 待实现-[文档](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_8_1.shtml)
* <font color='#07C160' size='4'>代金券</font>
    * 创建代金券批次：`client.V3FavorBatchCreate()`
    * 激活代金券批次：`client.V3FavorBatchStart()`
    * 发放代金券批次：`client.V3FavorBatchGrant()`
    * 暂停代金券批次：`client.V3FavorBatchPause()`
    * 重启代金券批次：`client.V3FavorBatchRestart()`
    * 条件查询批次列表：`client.V3FavorBatchList()`
    * 查询批次详情：`client.V3FavorBatchDetail()`
    * 查询代金券详情：`client.V3FavorDetail()`
    * 查询代金券可用商户：`client.V3FavorMerchant()`
    * 查询代金券可用单品：`client.V3FavorItems()`
    * 根据商户号查用户的券：`client.V3FavorUserCoupons()`
    * 下载批次核销明细：`client.V3FavorUseFlowDownload()`
    * 下载批次退款明细：`client.V3FavorRefundFlowDownload()`
    * 设置消息通知地址：`client.V3FavorCallbackUrlSet()`
* <font color='#07C160' size='4'>商家券</font>
    * 创建商家券：`client.V3BusiFavorBatchCreate()`
    * 查询商家券详情：`client.V3BusiFavorBatchDetail()`
    * 核销用户券：`client.V3BusiFavorUse()`
    * 根据过滤条件查询用户券：`client.V3BusiFavorUserCoupons()`
    * 查询用户单张券详情：`client.V3BusiFavorUserCouponDetail()`
    * 上传预存code：`client.V3BusiFavorCodeUpload()`
    * 设置商家券事件通知地址：`client.V3BusiFavorCallbackUrlSet()`
    * 查询商家券事件通知地址：`client.V3BusiFavorCallbackUrl()`
    * 关联订单信息：`client.V3BusiFavorAssociate()`
    * 取消关联订单信息：`client.V3BusiFavorDisassociate()`
    * 修改批次预算：`client.V3BusiFavorBatchUpdate()`
    * 修改商家券基本信息：`client.V3BusiFavorInfoUpdate()`
    * 发放消费卡：`client.V3BusiFavorSend()`
    * 申请退券：`client.V3BusiFavorReturn()`
    * 使券失效：`client.V3BusiFavorDeactivate()`
    * 营销补差付款：`client.V3BusiFavorSubsidyPay()`
    * 查询营销补差付款单详情：`client.V3BusiFavorSubsidyPayDetail()`
* <font color='#07C160' size='4'>委托营销</font>
    * 建立合作关系：`client.V3PartnershipsBuild()`
    * 终止合作关系：`client.V3PartnershipsTerminate()`
    * 查询合作关系列表：`client.V3PartnershipsList()`
* <font color='#07C160' size='4'>支付有礼</font>
    * 待实现-[文档](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_7_2.shtml)
* <font color='#07C160' size='4'>分账</font>
    * 请求分账：`client.V3ProfitShareOrder()`
    * 查询分账结果：`client.V3ProfitShareOrderQuery()`
    * 请求分账回退：`client.V3ProfitShareReturn()`
    * 查询分账回退结果：`client.V3ProfitShareReturnResult()`
    * 解冻剩余资金：`client.V3ProfitShareOrderUnfreeze()`
    * 查询剩余待分金额：`client.V3ProfitShareUnsplitAmount()`
    * 添加分账接收方：`client.V3ProfitShareAddReceiver()`
    * 删除分账接收方：`client.V3ProfitShareDeleteReceiver()`
* <font color='#07C160' size='4'>消费者投诉2.0</font>
    * 查询投诉单列表：`client.V3ComplaintList()`
    * 查询投诉单详情：`client.V3ComplaintDetail()`
    * 查询投诉协商历史：`client.V3ComplaintNegotiationHistory()`
    * 创建投诉通知回调地址：`client.V3ComplaintNotifyUrlCreate()`
    * 查询投诉通知回调地址：`client.V3ComplaintNotifyUrlQuery()`
    * 更新投诉通知回调地址：`client.V3ComplaintNotifyUrlUpdate()`
    * 删除投诉通知回调地址：`client.V3ComplaintNotifyUrlDelete()`
    * 提交回复：`client.V3ComplaintResponse()`
    * 反馈处理完成：`client.V3ComplaintComplete()`
    * 商户上传反馈图片：`client.V3ComplaintUploadImage()`
* <font color='#07C160' size='4'>其他能力</font>
    * 图片上传：`client.V3MediaUploadImage()`
    * 视频上传：`client.V3MediaUploadVideo()`
    * 图片上传（营销专用）：`client.V3FavorMediaUploadImage()`
* <font color='#07C160' size='4'>商家转账到零钱（直连商户）</font>
    * 发起商家转账：`client.V3Transfer()`
    * 微信批次单号查询批次单：`client.V3TransferQuery()`
    * 微信明细单号查询明细单：`client.V3TransferDetail()`
    * 商家批次单号查询批次单：`client.V3TransferMerchantQuery()`
    * 商家明细单号查询明细单：`client.V3TransferMerchantDetail()`
    * 转账电子回单申请受理：`client.V3TransferReceipt()`
    * 查询转账电子回单：`client.V3TransferReceiptQuery()`
    * 转账明细电子回单受理：`client.V3TransferDetailReceipt()`
    * 查询转账明细电子回单受理结果：`client.V3TransferDetailReceiptQuery()`
* <font color='#07C160' size='4'>转账（服务商）</font>
    * 发起批量转账：`client.V3PartnerTransfer()`
    * 微信批次单号查询批次单：`client.V3PartnerTransferQuery()`
    * 微信明细单号查询明细单：`client.V3PartnerTransferDetail()`
    * 商家批次单号查询批次单：`client.V3PartnerTransferMerchantQuery()`
    * 商家明细单号查询明细单：`client.V3PartnerTransferMerchantDetail()`
* <font color='#07C160' size='4'>余额查询</font>
    * 查询特约商户账户实时余额：`client.V3EcommerceBalance()`
    * 查询二级商户账户日终余额：`client.V3EcommerceDayBalance()`
    * 查询账户实时余额：`client.V3MerchantBalance()`
    * 查询账户日终余额：`client.V3MerchantDayBalance()`
* <font color='#07C160' size='4'>来账识别</font>
    * 商户银行来账查询：`client.V3MerchantIncomeRecord()`
    * 特约商户银行来账查询：`client.V3EcommerceIncomeRecord()`
* <font color='#07C160' size='4'>特约商户进件（服务商）</font>
    * 提交申请单：`client.V3Apply4SubSubmit()`
    * 查询申请单状态（BusinessCode）：`client.V3Apply4SubQueryByBusinessCode()`
    * 查询申请单状态（ApplyId）：`client.V3Apply4SubQueryByApplyId()`
    * 修改结算账号：`client.V3Apply4SubModifySettlement()`
    * 查询结算账户：`client.V3Apply4SubQuerySettlement()`
* <font color='#07C160' size='4'>点金计划（服务商）</font>
    * 点金计划管理：`client.V3GoldPlanManage()`
    * 商家小票管理：`client.V3GoldPlanBillManage()`
    * 同业过滤标签管理：`client.V3GoldPlanFilterManage()`
    * 开通广告展示：`client.V3GoldPlanOpenAdShow()`
    * 关闭广告展示：`client.V3GoldPlanCloseAdShow()`
* <font color='#07C160' size='4'>电商收付通（商户进件）</font>
    * 二级商户进件：`client.V3EcommerceApply()`
    * 查询申请状态：`client.V3EcommerceApplyStatus()`
* <font color='#07C160' size='4'>电商收付通（分账）</font>
    * 请求分账：`client.V3EcommerceProfitShare()`
    * 查询分账结果：`client.V3EcommerceProfitShareResult()`
    * 请求分账回退：`client.V3EcommerceProfitShareReturn()`
    * 查询分账回退结果：`client.V3EcommerceProfitShareReturnResult()`
    * 完结分账：`client.V3EcommerceProfitShareFinish()`
    * 查询订单剩余待分金额：`client.V3EcommerceProfitShareUnsplitAmount()`
    * 添加分账接收方：`client.V3EcommerceProfitShareAddReceiver()`
    * 删除分账接收方：`client.V3EcommerceProfitShareDeleteReceiver()`
* <font color='#07C160' size='4'>电商收付通（补差）</font>
    * 请求补差：`client.V3EcommerceProfitShare()`
    * 请求补差回退：`client.V3EcommerceSubsidiesReturn()`
    * 取消补差：`client.V3EcommerceSubsidiesCancel()`
* <font color='#07C160' size='4'>电商收付通（退款）</font>
    * 申请退款：`client.V3EcommerceRefund()`
    * 查询退款ById：`client.V3EcommerceRefundQueryById()`
    * 查询退款ByNo：`client.V3EcommerceRefundQueryByNo()`
    * 垫付退款回补：`client.V3EcommerceRefundAdvance()`
    * 查询垫付回补结果：`client.V3EcommerceRefundAdvanceResult()`
* <font color='#07C160' size='4'>银行组件（服务商）</font>
    * 获取对私银行卡号开户银行：`client.V3BankSearchBank()`
    * 查询支持个人业务的银行列表：`client.V3BankSearchPersonalList()`
    * 查询支持对公业务的银行列表：`client.V3BankSearchCorporateList()`
    * 查询省份列表：`client.V3BankSearchProvinceList()`
    * 查询城市列表：`client.V3BankSearchCityList()`
    * 查询支行列表：`client.V3BankSearchBranchList()`


### 微信v3公共 API

* `wechat.GetPlatformCerts()` => 获取微信平台证书公钥
* `client.GetAndSelectNewestCert()` => 获取并选择最新的有效证书
* `wechat.V3VerifySignByPK()` => 微信V3 同步验签（推荐直接打开自动验签功能）
* `wechat.V3ParseNotify()` => 解析微信回调请求的参数到 V3NotifyReq 结构体
* `notify.VerifySignByPK()` => 微信V3 异步通知验签
* `client.V3EncryptText()` => 敏感参数信息加密
* `client.V3DecryptText()` =>  敏感参数信息解密
* `wechat.V3EncryptText()` => 敏感参数信息加密
* `wechat.V3DecryptText()` =>  敏感参数信息解密
* `wechat.V3DecryptNotifyCipherText()` => 解密 普通支付 回调中的加密信息
* `wechat.V3DecryptRefundNotifyCipherText()` => 解密 普通退款 回调中的加密信息
* `wechat.V3DecryptCombineNotifyCipherText()` => 解密 合单支付 回调中的加密信息
* `wechat.V3DecryptScoreNotifyCipherText()` => 解密 支付分 回调中的加密信息
* `client.PaySignOfJSAPI()` => 获取 JSAPI 支付 paySign
* `client.PaySignOfApp()` => 获取 APP 支付 paySign
* `client.PaySignOfApplet()` => 获取 小程序 支付 paySign

