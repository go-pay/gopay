## 扫呗支付 API


- 扫呗支付：[官方文档中心](https://help.lcsw.cn/xrmpic/q6imdiojes7iq5y1/qg52lx)



> 具体API使用介绍，请参考`gopay/saobei/client_test.go`


### 支付2.0接口 
> 请参考`gopay/saobei/pay_test.go`,
* 小程序支付接口(暂无账号为测试可用性)：`client.MiniPay()`
* 付款码支付 `client.BarcodePay()`
* 支付查询  `client.Query()`
* 退款申请 `client.Refund()`
* 退款订单查询 `client.QueryRefund()`

### 资金接口
> 请参考`gopay/saobei/merchant_test.go`,

### CBK企业钱包分账
> 请参考`gopay/saobei/account_test.go`,