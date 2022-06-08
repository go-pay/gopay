## PayPal

> 具体API使用介绍，请参考`gopay/paypal/client_test.go`,`gopay/paypal/order_test.go`,`gopay/paypal/payment_test.go` 等xxx_test.go

- 已实现API列表附录：[API List](https://github.com/go-pay/gopay/blob/main/doc/paypal.md#%E9%99%84%E5%BD%95)

- PayPal官方文档：[Official Document](https://developer.paypal.com/docs/api/overview)

---

### 1、初始化PayPal客户端并做配置（Init PayPal Client）

```go
import (
    "github.com/go-pay/gopay/paypal"
    "github.com/go-pay/gopay/pkg/xlog"
)

// 初始化PayPal支付客户端
client, err := paypal.NewClient(Clientid, Secret, false)
if err != nil {
    xlog.Error(err)
    return
}
// 打开Debug开关，输出日志，默认关闭
client.DebugSwitch = gopay.DebugOn
```

### 2、API 方法调用及入参（Call API）

> Orders：[Orders API](https://developer.paypal.com/api/orders/v2/)

> Payments：[Payments API](https://developer.paypal.com/api/payments/v2/)

> Subscriptions：[Subscriptions API](https://developer.paypal.com/docs/api/subscriptions/v1/)

- Create Orders example

```go
import (
"github.com/go-pay/gopay"
"github.com/go-pay/gopay/paypal"
"github.com/go-pay/gopay/pkg/util"
"github.com/go-pay/gopay/pkg/xlog"
)

// Create Orders example
var pus []*paypal.PurchaseUnit
var item = &paypal.PurchaseUnit{
ReferenceId: util.GetRandomString(16),
Amount: &paypal.Amount{
CurrencyCode: "USD",
Value:        "8",
},
}
pus = append(pus, item)

bm := make(gopay.BodyMap)
bm.Set("intent", "CAPTURE").
Set("purchase_units", pus).
SetBodyMap("application_context", func(b gopay.BodyMap) {
b.Set("brand_name", "gopay").
Set("locale", "en-PT").
Set("return_url", "https://example.com/returnUrl").
Set("cancel_url", "https://example.com/cancelUrl")
})
ppRsp, err := client.CreateOrder(ctx, bm)
if err != nil {
xlog.Error(err)
return
}
if ppRsp.Code != paypal.Success {
// do something
return
}
```

- Capture payment for order

```go
import (
"github.com/go-pay/gopay"
"github.com/go-pay/gopay/pkg/xlog"
)

// Capture payment for order
//bm := make(gopay.BodyMap)
//bm.SetBodyMap("payment_source", func(b gopay.BodyMap) {
//	b.SetBodyMap("token", func(b gopay.BodyMap) {
//		b.Set("id", "The PayPal-generated ID for the token").
//			Set("type", "BILLING_AGREEMENT")
//	})
//})
ppRsp, err := client.OrderCapture(ctx, "4X223967G91314611", nil)
if err != nil {
xlog.Error(err)
return
}
if ppRsp.Code != paypal.Success {
// do something
return
}
```

---

## 附录：

### PayPal API

* <font color='#003087' size='4'>AccessToken</font>
    * 获取AccessToken（Get AccessToken）：`client.GetAccessToken()`
* <font color='#003087' size='4'>订单</font>
    * 创建订单（Create order）：`client.CreateOrder()`
    * 订单详情（Show order details）：`client.OrderDetail()`
    * 更新订单（Update order）：`client.UpdateOrder()`
    * 订单支付授权（Authorize payment for order）：`client.OrderAuthorize()`
    * 订单支付捕获（Capture payment for order）：`client.OrderCapture()`
    * 订单支付确认（Confirm the Order）：`client.OrderConfirm()`
* <font color='#003087' size='4'>支付</font>
    * 支付授权详情（Show details for authorized payment）：`client.PaymentAuthorizeDetail()`
    * 重新授权支付授权（Reauthorize authorized payment）：`client.PaymentReauthorize()`
    * 作废支付授权（Void authorized payment）：`client.PaymentAuthorizeVoid()`
    * 支付授权捕获（Capture authorized payment）：`client.PaymentAuthorizeCapture()`
    * 支付捕获详情（Show captured payment details）：`client.PaymentCaptureDetail()`
    * 支付捕获退款（Refund captured payment）：`client.PaymentCaptureRefund()`
    * 支付退款详情（Show refund details）：`client.PaymentRefundDetail()`

* <font color='#003087' size='4'>支出</font>
    * 创建批量支出（Create batch payout）：`client.CreateBatchPayout()`
    * 批量支出详情（Show payout batch details）：`client.ShowPayoutBatchDetails()`
    * 批量支出项目详情（Show Payout Item Details）：`client.ShowPayoutItemDetails()`
    * 取消批量支付中收款人无PayPal账号的项目（Cancel Unclaimed Payout Item）：`client.CancelUnclaimedPayoutItem()`

* <font color='#003087' size='4'>订阅</font>
	* 创建订阅计划（Create plan）：`client.CreateBillingPlan()`

