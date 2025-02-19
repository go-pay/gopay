## PayPal

> 具体API使用介绍，请参考`gopay/paypal/client_test.go`,`gopay/paypal/order_test.go`,`gopay/paypal/payment_test.go` 等xxx_test.go

- 已实现API列表附录：[API List](https://github.com/go-pay/gopay/blob/main/doc/paypal.md#%E9%99%84%E5%BD%95)

- PayPal官方文档：[Official Document](https://developer.paypal.com/api/rest)

---

### 1、初始化PayPal客户端并做配置（Init PayPal Client）

```go
import (
    "github.com/go-pay/gopay/paypal"
    "github.com/go-pay/xlog"
)

// 初始化PayPal支付客户端
client, err := paypal.NewClient(Clientid, Secret, false)
if err != nil {
    xlog.Error(err)
    return
}

// 自定义配置http请求接收返回结果body大小，默认 10MB
client.SetBodySize() // 没有特殊需求，可忽略此配置

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
    "github.com/go-pay/util"
    "github.com/go-pay/xlog"
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
	SetBodyMap("payment_source", func(b gopay.BodyMap) {
		b.SetBodyMap("paypal", func(bb gopay.BodyMap) {
			bb.SetBodyMap("experience_context", func(bbb gopay.BodyMap) {
				bbb.Set("brand_name", "gopay").
					Set("locale", "en-US").
					Set("shipping_preference", "NO_SHIPPING").
					Set("user_action", "PAY_NOW").
					Set("return_url", "http://xxx/return").
					Set("cancel_url", "http://xxx/cancel")
			})
		})
	})
ppRsp, err := client.CreateOrder(ctx, bm)
if err != nil {
	xlog.Error(err)
	return
}
if ppRsp.Code != 200 {
	// do something
	return
}
```

- Capture payment for order

```go
import (
    "github.com/go-pay/gopay"
    "github.com/go-pay/xlog"
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
* <font color='#003087' size='4'>Invoices</font>
	* 生成发票号码（Generate invoice number）：`client.InvoiceNumberGenerate()`
	* 发票列表（List invoices）：`client.InvoiceList()`
	* 创建虚拟发票（Create draft invoice）：`client.InvoiceCreate()`
	* 删除发票（Delete invoice）：`client.InvoiceDelete()`
	* 更新发票（Fully update invoice）：`client.InvoiceUpdate()`
	* 获取发票详情（Show invoice details）：`client.InvoiceDetail()`
	* 生成发票二维码（Generate QR code）：`client.InvoiceGenerateQRCode()`
	* 发票付款记录（Record payment for invoice）：`client.InvoicePaymentRecord()`
	* 发票付款删除（Delete external payment）：`client.InvoicePaymentDelete()`
	* 发票退款记录（Record refund for invoice）：`client.InvoiceRefundRecord()`
	* 发票退款删除（Delete external refund）：`client.InvoiceRefundDelete()`
	* 发送发票提醒（Send invoice reminder）：`client.InvoiceSendRemind()`
	* 发送发票（Send invoice）：`client.InvoiceSend()`
	* 发票搜索（Search for invoices）：`client.InvoiceSearch()`
	* 发票模板列表（List templates）：`client.InvoiceTemplateList()`
	* 创建发票模板（Create template）：`client.InvoiceTemplateCreate()`
	* 删除发票模板（Delete template）：`client.InvoiceTemplateDelete()`
	* 更新发票模板（Fully update template）：`client.InvoiceTemplateUpdate()`
* <font color='#003087' size='4'>Orders</font>
    * 创建订单（Create order）：`client.CreateOrder()`
    * 订单详情（Show order details）：`client.OrderDetail()`
    * 更新订单（Update order）：`client.UpdateOrder()`
    * 订单支付授权（Authorize payment for order）：`client.OrderAuthorize()`
    * 订单支付捕获（Capture payment for order）：`client.OrderCapture()`
    * 订单支付确认（Confirm the Order）：`client.OrderConfirm()`
* <font color='#003087' size='4'>Payments</font>
    * 支付授权详情（Show details for authorized payment）：`client.PaymentAuthorizeDetail()`
    * 重新授权支付授权（Reauthorize authorized payment）：`client.PaymentReauthorize()`
    * 作废支付授权（Void authorized payment）：`client.PaymentAuthorizeVoid()`
    * 支付授权捕获（Capture authorized payment）：`client.PaymentAuthorizeCapture()`
    * 支付捕获详情（Show captured payment details）：`client.PaymentCaptureDetail()`
    * 支付捕获退款（Refund captured payment）：`client.PaymentCaptureRefund()`
    * 支付退款详情（Show refund details）：`client.PaymentRefundDetail()`
* <font color='#003087' size='4'>Payment Method Tokens</font>
    * 为给定的支付来源创建支付令牌（Create payment token for a given payment source）：`client.CreatePaymentToken()`
    * 列出所有支付令牌（List all payment tokens）：`client.ListAllPaymentTokens()`
    * 检索付款令牌（Retrieve a payment token）：`client.RetrievePaymentToken()`
    * 删除付款令牌（Delete payment token）：`client.DeletePaymentToken()`
    * 创建设置令牌（Create a setup token）：`client.CreateSetupToken()`
    * 检索设置令牌（Retrieve a setup token）：`client.RetrieveSetupToken()`
* <font color='#003087' size='4'>Payouts</font>
    * 创建批量支出（Create batch payout）：`client.CreateBatchPayout()`
    * 批量支出详情（Show payout batch details）：`client.ShowPayoutBatchDetails()`
    * 批量支出项目详情（Show Payout Item Details）：`client.ShowPayoutItemDetails()`
    * 取消批量支付中收款人无PayPal账号的项目（Cancel Unclaimed Payout Item）：`client.CancelUnclaimedPayoutItem()`
* <font color='#003087' size='4'>Subscriptions</font>
    * 创建订阅计划（Create plan）：`client.CreateBillingPlan()`

