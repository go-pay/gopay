## 通联支付


- 通联支付：[官方文档中心](https://aipboss.allinpay.com/know/devhelp/index.php)



> 具体API使用介绍，请参考`gopay/allinpay/client_test.go`,`gopay/allinpay/pay_test.go` 等xxx_test.go

### 通联支付 API

* 统一支付接口(暂无账号为测试可用性)：`client.Pay()`
* 统一扫码接口: `client.ScanPay()`
* 撤销订单：`client.Cancel()`
* 交易退款：`client.Refund()`
* 交易结果查询：`client.Query()`
* 关闭订单：`client.Close()`
* 申请退款：`client.Refund()`
