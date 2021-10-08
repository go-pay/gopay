

### QQ支付 API

* 提交付款码支付：`client.MicroPay()`
* 撤销订单：`client.Reverse()`
* 统一下单：`client.UnifiedOrder()`
* 订单查询：`client.OrderQuery()`
* 关闭订单：`client.CloseOrder()`
* 申请退款：`client.Refund()`
* 退款查询：`client.RefundQuery()`
* 交易账单：`client.StatementDown()`
* 资金账单：`client.AccRoll()`
* 创建现金红包（未测试可用性）：`client.SendCashRed()`
* 对账单下载（未测试可用性）：`client.DownloadRedListFile()`
* 查询红包详情（未测试可用性）：`client.QueryRedInfo()`
* 自定义方法请求微信API接口：`client.PostQQAPISelf()`

### QQ公共 API

* `qq.ParseNotifyToBodyMap()` => 解析QQ支付异步通知的结果到BodyMap
* `qq.ParseNotify()` => 解析QQ支付异步通知的参数
* `qq.VerifySign()` => QQ同步返回参数验签或异步通知参数验签

---