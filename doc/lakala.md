## 拉卡拉网关支付

- 拉卡拉网关支付：[官方文档中心](https://payjp.lakala.com/docs/cn)

### 回调通知解析

```
notifyReq, err := lakala.ParseNotify()
if err != nil {
    xlog.Error(err)
    return
}
```

### 拉卡拉 API

* <font color='#07C160' size='4'>QRCode</font>
    * 创建QRCode支付单：`client.CreateQRCodeOrder()`
    * 创建Native QRCode支付单：`client.CreateNativeQRCodeOrder()`
    * QRCode支付跳转页：`client.QRCodePay()`
* <font color='#07C160' size='4'>JSAPI</font>
    * 创建JSAPI订单：`client.CreateJSAPIOrder()`
    * 创建Native JSAPI订单(offline)：`client.CreateNativeJSApiOrder()`
    * 微信JSAPI支付跳转页：`client.JSAPIWechatPay()`
    * 支付宝JSAPI支付跳转页：`client.JSAPIAlipayPay()`
    * Alipay+ JSAPI支付跳转页：`client.JSAPIAlipayPlusPay()`
* <font color='#07C160' size='4'>MobileH5</font>
    * 创建H5支付单：`client.CreateH5PayOrder()`
    * H5支付跳转页：`client.H5Pay()`
    * H5支付跳转页(Alipay+)：`client.H5AlipayPlusPay()`
* <font color='#07C160' size='4'>Miniprogram Payment</font>
    * 创建小程序订单：`client.CreateMiniProgramOrder()`
* <font color='#07C160' size='4'>Channel Web Gateway</font>
    * 创建渠道Web网关订单：`client.CreateWebGatewayOrder()`
* <font color='#07C160' size='4'>SDKPayment</font>
    * 创建SDK订单(Online)：`client.CreateSDKPaymentOrder()`
* <font color='#07C160' size='4'>CommonApi</font>
    * 获取当前汇率：`client.GetExchangeRate()`
    * 获取加密密钥：`client.GetEncrypt()`
    * 关闭订单：`client.CloseOrder()`
    * 查询订单状态：`client.OrderStatus()`
    * 申请退款：`client.ApplyRefund()`
    * 查询退款状态：`client.RefundQuery()`
    * 查看订单：`client.OrderList()`
    * 查看账单流水：`client.TransactionList()`
    * 查看清算详情：`client.Settlements()`
    * 查询可用钱包：`client.ConsultPayment()`
    * 获取优惠券信息：`client.GetCoupon()`
* <font color='#07C160' size='4'>Custom</font>
    * 创建报关单（非拆单）：`client.CreateReportSingle()`
    * 创建报关单（拆单）：`client.CreateReportSeparate()`
    * 报关状态查询：`client.ReportStatus()`
    * 报关子单状态查询：`client.ReportSubStatus()`
    * 修改报关信息（非拆单）：`client.ModifyReportSingle()`
    * 修改报关信息（拆单）：`client.ModifyReportSeparate()`
    * 重推报关（非拆单）：`client.ResendReportSingle()`
    * 报关单子单重推：`client.ResendReportSeparate()`
