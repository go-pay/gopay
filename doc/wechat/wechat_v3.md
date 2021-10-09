### 微信支付V3 API

> #### 推荐使用V3接口，官方在V3接口实现未覆盖或gopay未开发的接口，还继续用V2接口，欢迎参与完善V3接口。

* <font color='#07C160' size='4'>基础支付</font>
    * APP下单：`client.V3TransactionApp()`
    * JSAPI/小程序下单：`client.V3TransactionJsapi()`
    * Native下单：`client.V3TransactionNative()`
    * H5下单：`client.V3TransactionH5()`
    * 查询订单：`client.V3TransactionQueryOrder()`
    * 关闭订单：`client.V3TransactionCloseOrder()`
* <font color='#07C160' size='4'>基础支付（服务商）</font>
    * APP下单：`client.V3PartnerTransactionApp()`
    * JSAPI/小程序下单：`client.V3PartnerTransactionJsapi()`
    * Native下单：`client.V3PartnerTransactionNative()`
    * H5下单：`client.V3PartnerTransactionH5()`
    * 查询订单：`client.V3PartnerQueryOrder()`
    * 关闭订单：`client.V3PartnerCloseOrder()`
* <font color='#07C160' size='4'>合单支付</font>
    * 合单APP下单：`client.V3CombineTransactionApp()`
    * 合单JSAPI/小程序下单：`client.V3CombineTransactionJsapi()`
    * 合单Native下单：`client.V3CombineTransactionNative()`
    * 合单H5下单：`client.V3CombineTransactionH5()`
    * 合单查询订单：`client.V3CombineQueryOrder()`
    * 合单关闭订单：`client.V3CombineCloseOrder()`
* <font color='#07C160' size='4'>退款</font>
    * 申请退款：`client.V3Refund()`
    * 查询单笔退款：`client.V3RefundQuery()`
* <font color='#07C160' size='4'>账单</font>
    * 申请交易账单：`client.V3BillTradeBill()`
    * 申请资金账单：`client.V3BillFundFlowBill()`
    * 申请特约商户资金账单：`client.V3BillEcommerceFundFlowBill()`
    * 下载账单：`client.V3BillDownLoadBill()`
* <font color='#07C160' size='4'>提现（服务商）</font>
    * 待实现-[文档](https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer_partner/chapter6_1.shtml)
* <font color='#07C160' size='4'>微信支付分（公共API）</font>
    * 创建支付分订单：`client.V3ScoreOrderCreate()`
    * 查询支付分订单：`client.V3ScoreOrderQuery()`
    * 取消支付分订单：`client.V3ScoreOrderCancel()`
    * 修改订单金额：`client.V3ScoreOrderModify()`
    * 完结支付分订单：`client.V3ScoreOrderComplete()`
    * 商户发起催收扣款：`client.V3ScoreOrderPay()`
    * 同步服务订单信息：`client.V3ScoreOrderSync()`
* <font color='#07C160' size='4'>微信支付分（免确认模式）</font>
    * 创单结单合并：`client.V3ScoreDirectComplete()`
* <font color='#07C160' size='4'>微信支付分（免确认预授权模式）</font>
    * 商户预授权：`client.V3ScorePermission()`
    * 查询用户授权记录（授权协议号）：`client.V3ScorePermissionQuery()`
    * 解除用户授权关系（授权协议号）：`client.V3ScorePermissionTerminate()`
    * 查询用户授权记录（openid）：`client.V3ScorePermissionOpenidQuery()`
    * 解除用户授权关系（openid）：`client.V3ScorePermissionOpenidTerminate()`
* <font color='#07C160' size='4'>微信先享卡</font>
    * 预受理领卡请求：`client.V3DiscountCardApply()`
    * 增加用户记录：`client.V3DiscountCardAddUser()`
    * 查询先享卡订单：`client.V3DiscountCardQuery()`
* <font color='#07C160' size='4'>支付即服务</font>
    * 服务人员注册：`client.V3SmartGuideReg()`
    * 服务人员分配：`client.V3SmartGuideAssign()`
    * 服务人员查询：`client.V3SmartGuideQuery()`
    * 服务人员信息更新：`client.V3SmartGuideUpdate()`
* <font color='#07C160' size='4'>点金计划（服务商）</font>
    * 待实现-[文档](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_5_1.shtml)
* <font color='#07C160' size='4'>智慧商圈</font>
    * 商圈积分同步：`client.V3BusinessPointsSync()`
    * 商圈积分授权查询：`client.V3BusinessAuthPointsQuery()`
* <font color='#07C160' size='4'>微信支付分停车服务</font>
    * 待实现-[文档](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_8_1.shtml)
* <font color='#07C160' size='4'>代金券</font>
    * 创建代金券批次：`client.V3FavorBatchCreate()`
    * 激活代金券批次：`client.V3FavorBatchStart()`
    * 发放代金券批次：`client.V3FavorBatchGrant()`
    * 暂停代金券批次：`client.V3FavorBatchPause()`
    * 重启代金券批次：`client.V3FavorBatchRestart()`
    * 条件查询批次列表：`client.V3FavorBatchList()`
    * 查询批次详情：`client.V3FavorBatchDetail()`
    * 查询代金券详情：`client.V3FavorDetail()`
    * 查询代金券可用商户：`client.V3FavorMerchant()`
    * 查询代金券可用单品：`client.V3FavorItems()`
    * 根据商户号查用户的券：`client.V3FavorUserCoupons()`
    * 下载批次核销明细：`client.V3FavorUseFlowDownload()`
    * 下载批次退款明细：`client.V3FavorRefundFlowDownload()`
    * 设置消息通知地址：`client.V3FavorCallbackUrlSet()`
* <font color='#07C160' size='4'>商家券</font>
    * 创建商家券：`client.V3BusiFavorBatchCreate()`
    * 查询商家券详情：`client.V3BusiFavorBatchDetail()`
    * 核销用户券：`client.V3BusiFavorUse()`
    * 根据过滤条件查询用户券：`client.V3BusiFavorUserCoupons()`
    * 查询用户单张券详情：`client.V3BusiFavorUserCouponDetail()`
    * 上传预存code：`client.V3BusiFavorCodeUpload()`
    * 设置商家券事件通知地址：`client.V3BusiFavorCallbackUrlSet()`
    * 查询商家券事件通知地址：`client.V3BusiFavorCallbackUrl()`
    * 关联订单信息：`client.V3BusiFavorAssociate()`
    * 取消关联订单信息：`client.V3BusiFavorDisassociate()`
    * 修改批次预算：`client.V3BusiFavorBatchUpdate()`
    * 修改商家券基本信息：`client.V3BusiFavorInfoUpdate()`
    * 发放消费卡：`client.V3BusiFavorSend()`
    * 申请退券：`client.V3BusiFavorReturn()`
    * 使券失效：`client.V3BusiFavorDeactivate()`
    * 营销补差付款：`client.V3BusiFavorSubsidyPay()`
    * 查询营销补差付款单详情：`client.V3BusiFavorSubsidyPayDetail()`
* <font color='#07C160' size='4'>委托营销</font>
    * 建立合作关系：`client.V3PartnershipsBuild()`
    * 终止合作关系：`client.V3PartnershipsTerminate()`
    * 查询合作关系列表：`client.V3PartnershipsList()`
* <font color='#07C160' size='4'>支付有礼</font>
    * 待实现-[文档](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_7_2.shtml)
* <font color='#07C160' size='4'>分账</font>
    * 请求分账：`client.V3ProfitShareOrder()`
    * 查询分账结果：`client.V3ProfitShareOrderQuery()`
    * 请求分账回退：`client.V3ProfitShareReturn()`
    * 查询分账回退结果：`client.V3ProfitShareReturnResult()`
    * 解冻剩余资金：`client.V3ProfitShareOrderUnfreeze()`
    * 查询剩余待分金额：`client.V3ProfitShareUnsplitAmount()`
    * 添加分账接收方：`client.V3ProfitShareAddReceiver()`
    * 删除分账接收方：`client.V3ProfitShareDeleteReceiver()`
* <font color='#07C160' size='4'>消费者投诉2.0</font>
    * 查询投诉单列表：`client.V3ComplaintList()`
    * 查询投诉单详情：`client.V3ComplaintDetail()`
    * 查询投诉协商历史：`client.V3ComplaintNegotiationHistory()`
    * 创建投诉通知回调地址：`client.V3ComplaintNotifyUrlCreate()`
    * 查询投诉通知回调地址：`client.V3ComplaintNotifyUrlQuery()`
    * 更新投诉通知回调地址：`client.V3ComplaintNotifyUrlUpdate()`
    * 删除投诉通知回调地址：`client.V3ComplaintNotifyUrlDelete()`
    * 提交回复：`client.V3ComplaintResponse()`
    * 反馈处理完成：`client.V3ComplaintComplete()`
    * 商户上传反馈图片：`client.V3ComplaintUploadImage()`
* <font color='#07C160' size='4'>其他能力</font>
    * 图片上传：`client.V3MediaUploadImage()`
    * 视频上传：`client.V3MediaUploadVideo()`
    * 图片上传（营销专用）：`client.V3FavorMediaUploadImage()`
* <font color='#07C160' size='4'>批量转账</font>
    * 发起批量转账：`client.V3Transfer()`
    * 微信批次单号查询批次单：`client.V3TransferQuery()`
    * 微信明细单号查询明细单：`client.V3TransferDetailQuery()`
    * 商家批次单号查询批次单：`client.V3TransferMerchantQuery()`
    * 商家明细单号查询明细单：`client.V3TransferMerchantDetailQuery()`
    * 转账电子回单申请受理：`client.V3TransferReceipt()`
    * 查询转账电子回单：`client.V3TransferReceiptQuery()`
    * 转账明细电子回单受理：`client.V3TransferDetailReceipt()`
    * 查询转账明细电子回单受理结果：`client.V3TransferDetailReceiptQuery()`
* <font color='#07C160' size='4'>余额查询</font>
    * 查询特约商户账户实时余额（服务商）：`client.V3EcommerceBalance()`
    * 查询账户实时余额：`client.V3MerchantBalance()`
    * 查询账户日终余额：`client.V3MerchantDayBalance()`
* <font color='#07C160' size='4'>来账识别</font>
    * 商户银行来账查询：`client.V3MerchantIncomeRecord()`
    * 特约商户银行来账查询：`client.V3EcommerceIncomeRecord()`
* <font color='#07C160' size='4'>特约商户进件（服务商）</font>
    * 提交申请单：`client.V3Apply4SubSubmit()`
    * 查询申请单状态（BusinessCode）：`client.V3Apply4SubQueryByBusinessCode()`
    * 查询申请单状态（ApplyId）：`client.V3Apply4SubQueryByApplyId()`
    * 修改结算账号：`client.V3Apply4SubModifySettlement()`
    * 查询结算账户：`client.V3Apply4SubQuerySettlement()`

### 微信V3公共 API

* `wechat.GetPlatformCerts()` => 获取微信平台证书公钥
* `wechat.V3VerifySign()` => 微信V3 版本验签（同步/异步）
* `wechat.V3ParseNotify()` => 解析微信回调请求的参数到 V3NotifyReq 结构体
* `client.V3EncryptText()` => 敏感参数信息加密
* `client.V3DecryptText()` =>  敏感参数信息解密
* `wechat.V3EncryptText()` => 敏感参数信息加密
* `wechat.V3DecryptText()` =>  敏感参数信息解密
* `wechat.V3DecryptNotifyCipherText()` => 解密 普通支付 回调中的加密信息
* `wechat.V3DecryptRefundNotifyCipherText()` => 解密 普通退款 回调中的加密信息
* `wechat.V3DecryptCombineNotifyCipherText()` => 解密 合单支付 回调中的加密信息
* `wechat.V3DecryptScoreNotifyCipherText()` => 解密 支付分 回调中的加密信息
* `client.PaySignOfJSAPI()` => 获取 JSAPI 支付 paySign
* `client.PaySignOfApp()` => 获取 APP 支付 paySign
* `client.PaySignOfApplet()` => 获取 小程序 支付 paySign
