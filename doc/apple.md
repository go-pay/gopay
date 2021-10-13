## Apple
### Apple Pay 支付校验收据

* [苹果校验收据文档](https://developer.apple.com/documentation/appstorereceipts/verifyreceipt)

> url 请选择 apple.UrlSandbox 或 apple.UrlProd

* `apple.VerifyReceipt()` => 苹果支付校验收据API

---

### 校验示例

```go
import (
    "github.com/cedarwu/gopay/apple"
    "github.com/cedarwu/gopay/pkg/xlog"
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