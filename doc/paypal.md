## PayPal

PayPal官方文档：[官方文档](https://developer.paypal.com/docs/api/overview)

### PayPal支付 API

> 具体API使用介绍，请参考`gopay/paypal/client_test.go`,`gopay/paypal/order_test.go`,`gopay/paypal/payment_test.go`

* <font color='#003087' size='4'>AccessToken</font>
    * 获取AccessToken：`client.GetAccessToken()`
* <font color='#003087' size='4'>订单</font>
    * 创建订单（Create order）：`client.CreateOrder()`
    * 订单详情（Show order details）：`client.OrderDetail()`
    * 更新订单（Update order）：`client.UpdateOrder()`
    * 订单支付授权（Authorize payment for order）：`client.OrderAuthorize()`
    * 订单支付捕获（Capture payment for order）：`client.OrderCapture()`
* <font color='#003087' size='4'>支付</font>
    * 支付授权详情（Show details for authorized payment）：`client.PaymentAuthorizeDetail()`
    * 重新授权支付授权（Reauthorize authorized payment）：`client.PaymentReauthorize()`
    * 作废支付授权（Void authorized payment）：`client.PaymentAuthorizeVoid()`
    * 支付授权捕获（Capture authorized payment）：`client.PaymentAuthorizeCapture()`
    * 支付捕获详情（Show captured payment details）：`client.PaymentCaptureDetail()`
    * 支付捕获退款（Refund captured payment）：`client.PaymentCaptureRefund()`
    * 支付退款详情（Show refund details）：`client.PaymentRefundDetail()`

---

### 1、初始化PayPal客户端并做配置

> 具体API使用介绍，请参考 `gopay/paypal/client_test.go`

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

### 2、API 方法调用及入参

> Orders：[Orders API](https://developer.paypal.com/docs/api/orders/v2)

> Payments：[Payments API](https://developer.paypal.com/docs/api/payments/v2)

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
    Set("purchase_units", pus)

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