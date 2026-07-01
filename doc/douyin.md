## 抖音支付

> 抖音支付 SDK 面向直连商户，覆盖基础支付、账单、分账、转账及相关回调通知。

- 已实现API列表附录：[API 列表附录](https://github.com/go-pay/gopay/blob/main/doc/douyin.md#%E9%99%84%E5%BD%95)

- 抖音官方文档入口：[抖音支付开放平台](https://pay.douyinpay.com/wiki/)

- 通用签名认证 / 敏感字段加密 / 回调解密规则：请以官方文档为准

---

### 1、初始化抖音支付客户端并做配置

> 抖音支付无沙箱环境，测试请使用正式环境 1 分钱测试法。

> 具体 API 使用介绍，请参考 `gopay/examples/douyin/douyin.go`。

```go
import (
    "github.com/go-pay/xlog"
    "github.com/go-pay/gopay"
    "github.com/go-pay/gopay/douyin"
)

// NewClient 初始化抖音支付客户端
// mchid：商户号
// serialNo：商户 API 证书序列号
// apiKey：接口加密密钥（32 字节，用于回调 AES-256-GCM 解密）
// privateKey：商户 API 私钥 apiclient_key.pem 内容
client, err := douyin.NewClient(Mchid, SerialNo, ApiKey, PrivateKey)
if err != nil {
    xlog.Error(err)
    return
}

// 注册抖音支付平台证书（可多次调用注册多张，签名头 Douyinpay-Serial 会自动带最新一张）
// 抖音无平台证书自动拉取接口，请从抖音支付商户后台或运营处获取后手动注册
if err = client.SetPlatformCert([]byte(PlatformCertPEM), PlatformCertSerial); err != nil {
    xlog.Error(err)
    return
}

// 打开 Debug 开关，输出日志（默认关闭）
client.DebugSwitch = gopay.DebugOn

// 可选：自定义 xhttp.Client / Logger / 代理
// client.SetHttpClient(hc)
// client.SetLogger(logger)
// client.SetProxyHost("https://your-proxy.internal")
```

---

### 2、基础支付

四类下单接口结构一致（`bm gopay.BodyMap` 直接组装官方 body 字段，`mchid` 未设置时自动填 `client.Mchid`）：

```go
// App 下单
rsp, err := client.AppOrder(ctx, bm)         // rsp.Response.PrepayId
// JSAPI 下单
rsp, err := client.JsapiOrder(ctx, bm)       // rsp.Response.PrepayId
// H5 下单（响应直接返回 h5_url）
rsp, err := client.H5Order(ctx, bm)          // rsp.Response.H5Url
// Native 下单（响应直接返回二维码 code_url）
rsp, err := client.NativeOrder(ctx, bm)      // rsp.Response.CodeUrl
```

前端调起签名：

```go
appParams, err   := client.PaySignOfApp(appid, prepayId)   // App 端
jsapiParams, err := client.PaySignOfJSAPI(appid, prepayId) // JSAPI/小程序
```

订单查询 & 关单：

```go
q1, err := client.OrderQueryByTransactionId(ctx, "TP20221013...") // 抖音支付订单号
q2, err := client.OrderQueryByOutTradeNo(ctx, "OUT_XXX")          // 商户订单号
_,  err := client.CloseOrder(ctx, "OUT_XXX", nil)                 // 关单
```

退款：

```go
r1, err := client.Refund(ctx, bm)                            // 申请退款
r2, err := client.RefundQuery(ctx, "REF_XXX", mchid, appid)  // 查询退款
```

---

### 3、账单

三种账单 + 通用下载器 + 完整性校验：

```go
tradeBM  := gopay.BodyMap{"bill_date": "2026-06-30"} // bill_type 默认 TRADE，tar_type 默认 GZIP
fundBM   := gopay.BodyMap{"bill_date": "2026-06-30", "account_type": "BaseAccount"}
profitBM := gopay.BodyMap{"bill_date": "2026-06-30"}

tRsp, _ := client.ApplyTradeBill(ctx, tradeBM)
fRsp, _ := client.ApplyFundBill(ctx, fundBM)
pRsp, _ := client.ApplyProfitBill(ctx, profitBM)

// 下载 + 解压 + SHA1 校验
gz,   _ := client.DownloadBillFile(ctx, tRsp.Response.DownloadUrl)
raw,  _ := douyin.UngzipBill(gz)
err     := douyin.VerifyBillHash(raw, tRsp.Response.HashType, tRsp.Response.HashValue)
```

---

### 4、分账

```go
// 敏感字段先加密（type=MERCHANT_ID 时 name 必传）
encName, _ := client.EncryptText("接收方商户名称")

// 请求分账（异步受理）
rsp, _ := client.ProfitRequest(ctx, gopay.BodyMap{
    "appid":            "your_appid",
    "transaction_id":   "TP20221013...",
    "out_order_no":     "SPLIT_XXX",
    "unfreeze_unsplit": false,
    "notify_url":       "https://your.callback",
    "receivers": []douyin.ProfitReceiverReq{
        {Type: "MERCHANT_ID", Account: "6020...", Name: encName, Amount: 100, Description: "分账"},
    },
})

// 查询分账
_, _ = client.ProfitQuery(ctx, "SPLIT_XXX", gopay.BodyMap{"transaction_id": "TP..."})
// 完结分账
_, _ = client.ProfitComplete(ctx, gopay.BodyMap{"transaction_id": "TP...", "out_order_no": "FIN_XXX", "description": "完结"})
// 分账回退 / 回退查询
_, _ = client.ProfitRollback(ctx, gopay.BodyMap{...})
_, _ = client.ProfitRollbackQuery(ctx, "OUT_RETURN_XXX", gopay.BodyMap{"out_order_no": "SPLIT_XXX"})
// 剩余待分账余额
_, _ = client.ProfitBalanceQuery(ctx, "TP...", "") // mchid 传 "" 默认使用 client.Mchid
// 接收方管理
_, _ = client.ProfitReceiverAdd(ctx, gopay.BodyMap{"appid": "...", "type": "MERCHANT_ID", "account": "...", "name": encName, "relation_type": "STORE"})
_, _ = client.ProfitReceiverDelete(ctx, gopay.BodyMap{"appid": "...", "type": "MERCHANT_ID", "account": "..."})
```

---

### 5、转账（商户转账到零钱）

```go
// ≥2000.00 元时 user_name 必填，需先加密
encUserName, _ := client.EncryptText("张三")

rsp, _ := client.Transfer(ctx, gopay.BodyMap{
    "appid":             "your_appid",
    "out_bill_no":       "OUT_XXX",
    "transfer_scene_id": "SCENE_001",
    "openid":            "oUpF8uMEB4jR",
    "user_name":         encUserName,
    "transfer_amount":   100,
    "transfer_remark":   "商户转账",
    "notify_url":        "https://your.callback",
})

// 查询（两种维度）
_, _ = client.TransferQueryByTransferBillNo(ctx, rsp.Response.TransferBillNo)
_, _ = client.TransferQueryByOutBillNo(ctx, rsp.Response.OutBillNo)
```

---

### 6、回调通知处理

```go
// 1) 解析请求为 NotifyReq（内含 SignInfo）
notifyReq, err := douyin.ParseNotify(req)
if err != nil { /* ... */ }

// 2) 验签（按 header Douyinpay-Serial 匹配已注册的平台证书）
if err := notifyReq.VerifySignByPKMap(client.PlatformCertMap()); err != nil {
    // 验签失败：返回失败应答
    return
}

// 3) 根据 event_type 解密对应载荷
switch notifyReq.EventType {
case douyin.EventTransactionSuccess:      // TRANSACTION.SUCCESS
    pay, err := notifyReq.DecryptPayCipherText(ApiKey)
    _ = pay; _ = err
case douyin.EventRefundSuccess:           // REFUND.SUCCESS
    r, err := notifyReq.DecryptRefundCipherText(ApiKey)
    _ = r; _ = err
case douyin.EventAsyncSplitFinish:        // ASYNC_SPLIT.FINISH（分账结果）
    s, err := notifyReq.DecryptProfitResultCipherText(ApiKey)
    _ = s; _ = err
case douyin.EventSplitSuccess:            // SPLIT.SUCCESS（分账动账）
    d, err := notifyReq.DecryptProfitDynamicCipherText(ApiKey)
    _ = d; _ = err
case douyin.EventTransferSuccess:         // TRANSFER.SUCCESS
    t, err := notifyReq.DecryptTransferCipherText(ApiKey)
    _ = t; _ = err
}
```

---

### 7、敏感字段加解密（RSA-PKCS1v15）

抖音支付敏感字段（接收方姓名、转账收款人姓名等）走 **RSA-PKCS1v15**（注意与微信 V3 的 OAEP 不同）：

```go
// 加密：使用抖音支付平台证书公钥
cipher, err := client.EncryptText("张三")

// 解密：使用商户私钥
plain,  err := client.DecryptText(cipher)
```

---

## 附录：

### 抖音支付 API

* <font color='#07C160' size='4'>基础支付</font>
    * APP下单：`client.AppOrder()`
    * JSAPI/小程序下单：`client.JsapiOrder()`
    * H5下单：`client.H5Order()`
    * Native下单：`client.NativeOrder()`
    * App端调起支付签名：`client.PaySignOfApp()`
    * JSAPI前端调起支付签名：`client.PaySignOfJSAPI()`
    * 抖音支付订单号查询订单：`client.OrderQueryByTransactionId()`
    * 商户订单号查询订单：`client.OrderQueryByOutTradeNo()`
    * 关闭订单：`client.CloseOrder()`
* <font color='#07C160' size='4'>退款</font>
    * 申请退款：`client.Refund()`
    * 查询单笔退款（通过商户退款单号）：`client.RefundQuery()`
* <font color='#07C160' size='4'>账单</font>
    * 申请交易账单：`client.ApplyTradeBill()`
    * 申请资金账单：`client.ApplyFundBill()`
    * 申请分账账单：`client.ApplyProfitBill()`
    * 下载账单文件（跨域名，自动签名）：`client.DownloadBillFile()`
    * GZIP账单解压：`douyin.UngzipBill()`
    * SHA1完整性校验：`douyin.VerifyBillHash()`
* <font color='#07C160' size='4'>分账</font>
    * 请求分账（异步受理）：`client.ProfitRequest()`
    * 查询分账结果：`client.ProfitQuery()`
    * 请求分账回退：`client.ProfitRollback()`
    * 查询分账回退结果：`client.ProfitRollbackQuery()`
    * 完结分账：`client.ProfitComplete()`
    * 查询订单剩余待分账金额：`client.ProfitBalanceQuery()`
    * 添加分账接收方：`client.ProfitReceiverAdd()`
    * 删除分账接收方：`client.ProfitReceiverDelete()`
* <font color='#07C160' size='4'>转账（商户转账到零钱）</font>
    * 发起转账：`client.Transfer()`
    * 商户订单号查询转账单：`client.TransferQueryByOutBillNo()`
    * 抖音转账单号查询转账单：`client.TransferQueryByTransferBillNo()`
* <font color='#07C160' size='4'>客户端基础设施</font>
    * 初始化客户端：`douyin.NewClient()`
    * 注册抖音支付平台证书（多证书）：`client.SetPlatformCert()`
    * 获取已注册平台证书 map：`client.PlatformCertMap()`
    * 获取最新平台证书序列号：`client.NewestPlatformSerialNo()`
    * 设置自定义 xhttp.Client：`client.SetHttpClient()`
    * 设置自定义 Logger：`client.SetLogger()`
    * 设置代理域名：`client.SetProxyHost()`
    * 设置响应体大小上限：`client.SetBodySize()`
* <font color='#07C160' size='4'>敏感字段加解密（RSA-PKCS1v15）</font>
    * 加密（平台证书公钥）：`client.EncryptText()`
    * 解密（商户私钥）：`client.DecryptText()`
    * 包级加密：`douyin.EncryptText()`
    * 包级解密：`douyin.DecryptText()`
* <font color='#07C160' size='4'>回调通知（解析 / 验签 / 解密）</font>
    * 解析回调为结构体：`douyin.ParseNotify()`
    * 解析回调为 BodyMap：`douyin.ParseNotifyToBodyMap()`
    * 指定公钥验签：`(v *NotifyReq).VerifySignByPK()`
    * 按 Serial 匹配公钥 map 验签：`(v *NotifyReq).VerifySignByPKMap()`
    * 通用密文解密到任意结构体：`(v *NotifyReq).DecryptCipherTextToStruct()`
    * 支付成功通知解密（`TRANSACTION.SUCCESS`）：`(v *NotifyReq).DecryptPayCipherText()`
    * 退款结果通知解密（`REFUND.SUCCESS`）：`(v *NotifyReq).DecryptRefundCipherText()`
    * 分账结果通知解密（`ASYNC_SPLIT.FINISH`）：`(v *NotifyReq).DecryptProfitResultCipherText()`
    * 分账动账通知解密（`SPLIT.SUCCESS`）：`(v *NotifyReq).DecryptProfitDynamicCipherText()`
    * 转账结果通知解密（`TRANSFER.SUCCESS`）：`(v *NotifyReq).DecryptTransferCipherText()`
    * 通用回调密文解字节：`douyin.DecryptNotifyCipherTextToBytes()`
    * 通用回调密文解到结构体：`douyin.DecryptNotifyCipherTextToStruct()`
