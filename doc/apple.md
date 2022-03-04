## Apple
### Apple Pay 支付校验收据

* [苹果校验收据文档](https://developer.apple.com/documentation/appstorereceipts/verifyreceipt)

> url 请选择 apple.UrlSandbox 或 apple.UrlProd

* `apple.VerifyReceipt()` => 苹果支付校验收据API

---

### 校验示例

```go
import (
    "github.com/go-pay/gopay/apple"
    "github.com/go-pay/gopay/pkg/xlog"
)

pwd := ""
receipt := ""
rsp, err := apple.VerifyReceipt(UrlSandbox, pwd, receipt)
if err != nil {
    xlog.Error(err)
    return
}
/**
    response body:
    {
"receipt":{"original_purchase_date_pst":"2021-08-14 05:28:17 America/Los_Angeles", "purchase_date_ms":"1628944097586", "unique_identifier":"13f339a765b706f8775f729723e9b889b0cbb64e", "original_transaction_id":"1000000859439868", "bvrs":"10", "transaction_id":"1000000859439868", "quantity":"1", "in_app_ownership_type":"PURCHASED", "unique_vendor_identifier":"6DFDEA8B-38CE-4710-A1E1-BAEB8B66FEBD", "item_id":"1581250870", "version_external_identifier":"0", "bid":"com.huochai.main", "is_in_intro_offer_period":"false", "product_id":"10002", "purchase_date":"2021-08-14 12:28:17 Etc/GMT", "is_trial_period":"false", "purchase_date_pst":"2021-08-14 05:28:17 America/Los_Angeles", "original_purchase_date":"2021-08-14 12:28:17 Etc/GMT", "original_purchase_date_ms":"1628944097586"}, "status":0}
*/
if rsp.Receipt != nil {
    xlog.Infof("receipt:%+v", rsp.Receipt)
}
```


* [苹果服务端通知V2版本](https://developer.apple.com/documentation/appstoreservernotifications)

> 苹果支付服务服务端通知数据解析
> 对应App下 `[App 信息] --> [App Store 服务器通知]  --> [版本 2]` 配置对应的服务器地址,支付状态发生变化时Apple 将通过POST请求推送消息至配置的地址

### 示例

```go
import (
    "github.com/go-pay/gopay/apple"
    "github.com/go-pay/gopay/pkg/xlog"
    "encoding/json"
)

// Apple 通知请求体
body := "{\"signedPayload\":\"eyJhbGciOiJFUzI1NiIsIng1YyI6WyJNSUlOW...mnpo2QrItvA\"}"

var payload *NotificationV2SignedPayload
err := json.Unmarshal([]byte(body), &payload)
if err != nil {
    xlog.Error(err)
    return
}
rsp, err := payload.Decode()
if err != nil {
    xlog.Error(err)
    return
}
/*
* rsp结构如下
* {"notificationType":"DID_RENEW","subtype":"","notificationUUID":"7a33cb9d-2503-4104-a1e9-24bb9ea8bb75","notificationVersion":"","data":{"appAppleId":0,"bundleId":"bundleId","bundleVersion":"7","environment":"Sandbox","signedRenewalInfo":"","signedTransactionInfo":""},"renewalInfo":{"autoRenewProductId":"bundleId.productId","autoRenewStatus":1,"expirationIntent":0,"gracePeriodExpiresDate":0,"isInBillingRetryPeriod":false,"offerIdentifier":"","offerType":0,"originalTransactionId":"2000000000842607","priceIncreaseStatus":0,"productId":"bundleId.productId","signedDate":1646327704113},"transactionInfo":{"appAccountToken":"","bundleId":"bundleId","expiresDate":1646329907000,"inAppOwnershipType":"PURCHASED","isUpgraded":false,"offerIdentifier":"","offerType":0,"originalPurchaseDate":1646046037000,"originalTransactionId":"2000000000842607","productId":"bundleId.productId","purchaseDate":1646327747000,"quantity":1,"revocationDate":0,"revocationReason":"","signedDate":1646327704135,"subscriptionGroupIdentifier":"20929536","transactionId":"2000000003595767","type":"Auto-Renewable Subscription","webOrderLineItemId":"2000000000264942"}}
*/
xlog.Debugf("notify data: %s", rsp)
```
