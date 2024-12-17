## Apple

- App Store Server API：[官方文档](https://developer.apple.com/documentation/appstoreserverapi)

### 初始化Apple客户端

> 具体介绍，请参考 `gopay/apple/client_test.go`

```go
import (
    "github.com/go-pay/xlog"
    "github.com/go-pay/gopay/apple"
)

// 初始化通联客户端
// iss：issuer ID
// bid：bundle ID
// kid：private key ID
// privateKey：私钥文件读取后的字符串内容
// isProd：是否是正式环境
client, err = NewClient(iss, bid, kid, "privateKey", false)
if err != nil {
    xlog.Error(err)
    return
}
```


### Apple Pay 支付校验收据

* [苹果校验收据文档](https://developer.apple.com/documentation/appstorereceipts/verifyreceipt)

> url 请选择 apple.UrlSandbox 或 apple.UrlProd

* `apple.VerifyReceipt()` => 苹果支付校验收据API

---

#### 校验示例

```go
import (
    "github.com/go-pay/gopay/apple"
    "github.com/go-pay/xlog"
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
  {"receipt":{"original_purchase_date_pst":"2021-08-14 05:28:17 America/Los_Angeles", "purchase_date_ms":"1628944097586", "unique_identifier":"13f339a765b706f8775f729723e9b889b0cbb64e", "original_transaction_id":"1000000859439868", "bvrs":"10", "transaction_id":"1000000859439868", "quantity":"1", "in_app_ownership_type":"PURCHASED", "unique_vendor_identifier":"6DFDEA8B-38CE-4710-A1E1-BAEB8B66FEBD", "item_id":"1581250870", "version_external_identifier":"0", "bid":"com.huochai.main", "is_in_intro_offer_period":"false", "product_id":"10002", "purchase_date":"2021-08-14 12:28:17 Etc/GMT", "is_trial_period":"false", "purchase_date_pst":"2021-08-14 05:28:17 America/Los_Angeles", "original_purchase_date":"2021-08-14 12:28:17 Etc/GMT", "original_purchase_date_ms":"1628944097586"}, "status":0}
*/
if rsp.Receipt != nil {
    xlog.Infof("receipt:%+v", rsp.Receipt)
}
```

* [苹果服务端通知V2版本](https://developer.apple.com/documentation/appstoreservernotifications)

> 苹果支付服务服务端通知数据解析
> 对应App下 `[App 信息] --> [App Store 服务器通知]  --> [版本 2]` 配置对应的服务器地址,支付状态发生变化时Apple 将通过POST请求推送消息至配置的地址

#### 示例

- 请参考 `notification_v2_test.go`

```go
import (
    "github.com/go-pay/gopay/apple"
    "github.com/go-pay/xlog"
)

// decode signedPayload
payload, err := apple.DecodeSignedPayload(signedPayload)
if err != nil {
    xlog.Error(err)
    return
}
xlog.Debugf("payload.NotificationType: %s", payload.NotificationType)
xlog.Debugf("payload.Subtype: %s", payload.Subtype)
xlog.Debugf("payload.NotificationUUID: %s", payload.NotificationUUID)
xlog.Debugf("payload.NotificationVersion: %s", payload.NotificationVersion)
xlog.Debugf("payload.Data: %+v", payload.Data)
bs1, _ := json.Marshal(payload)
xlog.Info(string(bs1))
/*
   {
       "notificationType":"DID_RENEW",
       "subtype":"",
       "notificationUUID":"469bf30e-7715-4f9f-aae3-a7bfc12aea77",
       "notificationVersion":"",
       "data":{
           "appAppleId":0,
           "bundleId":"com.audaos.audarecorder",
           "bundleVersion":"7",
           "environment":"Sandbox",
           "signedRenewalInfo":"xxxxxxxxxx",
           "signedTransactionInfo":"xxxxxxxxxxx"
       }
   }
*/

// decode renewalInfo
renewalInfo, err := payload.DecodeRenewalInfo()
if err != nil {
    xlog.Error(err)
    return
}
xlog.Debugf("data.renewalInfo: %+v", renewalInfo)
bs, _ := json.Marshal(renewalInfo)
xlog.Info(string(bs))
/*
   {
       "autoRenewProductId":"com.audaos.audarecorder.vip.m2",
       "autoRenewStatus":1,
       "expirationIntent":0,
       "gracePeriodExpiresDate":0,
       "isInBillingRetryPeriod":false,
       "offerIdentifier":"",
       "offerType":0,
       "originalTransactionId":"2000000000842607",
       "priceIncreaseStatus":0,
       "productId":"com.audaos.audarecorder.vip.m2",
       "signedDate":1646387008228
   }
*/

// decode transactionInfo
transactionInfo, err := payload.DecodeTransactionInfo()
if err != nil {
    xlog.Error(err)
    return
}
xlog.Debugf("data.transactionInfo: %+v", transactionInfo)
bs2, _ := json.Marshal(transactionInfo)
xlog.Info(string(bs2))
/*
{
    "appAccountToken":"",
    "bundleId":"com.audaos.audarecorder",
    "expiresDate":1646387196000,
    "inAppOwnershipType":"PURCHASED",
    "isUpgraded":false,
    "offerIdentifier":"",
    "offerType":0,
    "originalPurchaseDate":1646046037000,
    "originalTransactionId":"2000000000842607",
    "productId":"com.audaos.audarecorder.vip.m2",
    "purchaseDate":1646387016000,
    "quantity":1,
    "revocationDate":0,
    "revocationReason":"",
    "signedDate":1646387008254,
    "subscriptionGroupIdentifier":"20929536",
    "transactionId":"2000000004047119",
    "type":"Auto-Renewable Subscription",
    "webOrderLineItemId":"2000000000302832"
}
*/
```

### App Store Server API Client Function

* `client.GetTransactionInfo()` => Get Transaction Info
* `client.GetTransactionHistory()` => Get Transaction History
* `client.GetAllSubscriptionStatuses()` => GetAllSubscriptionStatuses
* `client.SendConsumptionInformation()` => Send Consumption Information
* `client.GetNotificationHistory()` => Get Notification History
* `client.LookUpOrderId()` => Look Up Order ID
* `client.GetRefundHistory()` => Get Refund History

### Apple Function

* `apple.VerifyReceipt()` => 验证支付凭证
* `apple.ExtractClaims()` => 解析signedPayload
* `apple.DecodeSignedPayload()` => 解析notification signedPayload
