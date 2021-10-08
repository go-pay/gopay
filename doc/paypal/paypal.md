### PayPal支付 API

* [PayPal文档概览](https://developer.paypal.com/docs/api/overview)

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