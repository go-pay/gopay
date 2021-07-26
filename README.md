<div align=center><img width="240" height="240" alt="Logo was Loading Faild!" src="https://raw.githubusercontent.com/go-pay/gopay/main/logo.png"/></div>

# GoPay

#### QQ、微信、支付宝的Golang版本SDK

[![Github](https://img.shields.io/github/followers/iGoogle-ink?label=Follow&style=social)](https://github.com/iGoogle-ink)
[![Github](https://img.shields.io/github/forks/go-pay/gopay?label=Fork&style=social)](https://github.com/go-pay/gopay/fork)

[![Golang](https://img.shields.io/badge/golang-1.16-brightgreen.svg)](https://golang.google.cn)
[![GoDoc](https://img.shields.io/badge/doc-go.dev-informational.svg)](https://pkg.go.dev/github.com/go-pay/gopay)
[![Drone CI](https://cloud.drone.io/api/badges/go-pay/gopay/status.svg)](https://cloud.drone.io/go-pay/gopay)
[![GitHub Release](https://img.shields.io/github/v/release/go-pay/gopay)](https://github.com/go-pay/gopay/releases)
[![License](https://img.shields.io/github/license/go-pay/gopay)](https://www.apache.org/licenses/LICENSE-2.0)

---

# 一、安装

- v1.5.42 开始，仓库从 `github.com/iGoogle-ink/gopay` 迁移到 `github.com/go-pay/gopay`
```bash
go get github.com/go-pay/gopay
```

* #### 查看 GoPay 版本
    * [版本更新记录](https://github.com/go-pay/gopay/blob/main/release_note.txt)

```go
import (
    "fmt"

    "github.com/go-pay/gopay"
    "github.com/go-pay/gopay/pkg/xlog"
)

func main() {
    xlog.Debug("GoPay Version: ", gopay.Version)
}
```

---

### 微信支付V3 API

> #### 推荐使用V3接口，官方在V3接口实现未覆盖或gopay未开发的接口，还继续用V2接口，欢迎参与完善V3接口。

* <font color='#07C160' size='4'>基础支付（商户平台）</font>
    * APP下单：`client.V3TransactionApp()`
    * JSAPI/小程序下单：`client.V3TransactionJsapi()`
    * Native下单：`client.V3TransactionNative()`
    * H5下单：`client.V3TransactionH5()`
    * 查询订单：`client.V3TransactionQueryOrder()`
    * 关闭订单：`client.V3TransactionCloseOrder()`
* <font color='#07C160' size='4'>基础支付（服务商平台）</font>
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
    * 申请二级商户资金账单：`client.V3BillLevel2FundFlowBill()`
    * 下载账单：`client.V3BillDownLoadBill()`
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
* <font color='#07C160' size='4'>智慧商圈</font>
    * 商圈积分同步：`client.V3BusinessPointsSync()`
    * 商圈积分授权查询：`client.V3BusinessAuthPointsQuery()`
* <font color='#07C160' size='4'>代金券</font>
    * 待实现
* <font color='#07C160' size='4'>商家券</font>
    * 待实现
* <font color='#07C160' size='4'>委托营销</font>
    * 待实现
* <font color='#07C160' size='4'>消费卡</font>
    * 待实现
* <font color='#07C160' size='4'>支付有礼</font>
    * 待实现
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
* <font color='#07C160' size='4'>公共API</font>
    * 发起批量转账：`client.V3Transfer()`
    * 微信批次单号查询批次单：`client.V3TransferQuery()`
    * 微信明细单号查询明细单：`client.V3TransferDetailQuery()`
    * 商家批次单号查询批次单：`client.V3TransferMerchantQuery()`
    * 商家明细单号查询明细单：`client.V3TransferMerchantDetailQuery()`
    * 转账电子回单申请受理：`client.V3TransferReceipt()`
    * 查询转账电子回单：`client.V3TransferReceiptQuery()`
    * 转账明细电子回单受理：`client.V3TransferDetailReceipt()`
    * 查询转账明细电子回单受理结果：`client.V3TransferDetailReceiptQuery()`
    * 查询账户实时余额：`client.V3MerchantBalance()`
    * 查询账户日终余额：`client.V3MerchantDayBalance()`
* <font color='#07C160' size='4'>来账识别API</font>
    * 商户银行来账查询：`client.V3MerchantIncomeRecord()`
* <font color='#07C160' size='4'>特约商户进件（服务商平台）</font>
    * 提交申请单：`client.V3Apply4SubSubmit()`
    * 查询申请单状态（BusinessCode）：`client.V3Apply4SubQueryByBusinessCode()`
    * 查询申请单状态（ApplyId）：`client.V3Apply4SubQueryByApplyId()`
    * 修改结算账号：`client.V3Apply4SubModifySettlement()`
    * 查询结算账户：`client.V3Apply4SubQuerySettlement()`
  
### 微信支付V2 API

> #### 推荐使用V3接口，官方在V3接口实现未覆盖或gopay未开发的接口，还继续用V2接口。

* 统一下单：`client.UnifiedOrder()`
    * JSAPI - JSAPI支付（或小程序支付）
    * NATIVE - Native支付
    * APP - app支付
    * MWEB - H5支付
* 提交付款码支付：`client.Micropay()`
* 查询订单：`client.QueryOrder()`
* 关闭订单：`client.CloseOrder()`
* 撤销订单：`client.Reverse()`
* 申请退款：`client.Refund()`
* 查询退款：`client.QueryRefund()`
* 下载对账单：`client.DownloadBill()`
* 下载资金账单（正式）：`client.DownloadFundFlow()`
* 交易保障：`client.Report()`
* 拉取订单评价数据（正式）：`client.BatchQueryComment()`
* 企业付款（正式）：`client.Transfer()`
* 查询企业付款（正式）：`client.GetTransferInfo()`
* 授权码查询OpenId（正式）：`client.AuthCodeToOpenId()`
* 公众号纯签约（正式）：`client.EntrustPublic()`
* APP纯签约-预签约接口-获取预签约ID（正式）：`client.EntrustAppPre()`
* H5纯签约（正式）：`client.EntrustH5()`
* 支付中签约（正式）：`client.EntrustPaying()`
* 请求单次分账（正式）：`client.ProfitSharing()`
* 请求多次分账（正式）：`client.MultiProfitSharing()`
* 查询分账结果（正式）：`client.ProfitSharingQuery()`
* 添加分账接收方（正式）：`client.ProfitSharingAddReceiver()`
* 删除分账接收方（正式）：`client.ProfitSharingRemoveReceiver()`
* 完结分账（正式）：`client.ProfitSharingFinish()`
* 分账回退（正式）：`client.ProfitSharingReturn()`
* 分账回退结果查询（正式）：`client.ProfitSharingReturnQuery()`
* 企业付款到银行卡API（正式）：`client.PayBank()`
* 查询企业付款到银行卡API（正式）：`client.QueryBank()`
* 获取RSA加密公钥API（正式）：`client.GetRSAPublicKey()`
* 发放现金红包：`client.SendCashRed()`
* 发放现金裂变红包：`client.SendGroupCashRed()`
* 发放小程序红包：`client.SendAppletRed()`
* 查询红包记录：`client.QueryRedRecord()`
* 订单附加信息提交（海关）：`client.CustomsDeclareOrder()`
* 订单附加信息查询（海关）：`client.CustomsDeclareQuery()`
* 订单附加信息重推（海关）：`client.CustomsReDeclareOrder()`
* 自定义方法请求微信API接口：`client.PostWeChatAPISelf()`

### 微信公共V2 API

* `wechat.GetParamSign()` => 获取微信支付所需参数里的Sign值（通过支付参数计算Sign值）
* `wechat.GetSanBoxParamSign()` => 获取微信支付沙箱环境所需参数里的Sign值（通过支付参数计算Sign值）
* `wechat.GetMiniPaySign()` => 获取微信小程序支付所需要的paySign
* `wechat.GetH5PaySign()` => 获取微信内H5支付所需要的paySign
* `wechat.GetAppPaySign()` => 获取APP支付所需要的paySign
* `wechat.ParseNotifyToBodyMap()` => 解析微信支付异步通知的参数到BodyMap
* `wechat.ParseNotify()` => 解析微信支付异步通知的参数
* `wechat.ParseRefundNotify()` => 解析微信退款异步通知的参数
* `wechat.VerifySign()` => 微信同步返回参数验签或异步通知参数验签
* `wechat.Code2Session()` => 登录凭证校验：获取微信用户OpenId、UnionId、SessionKey
* `wechat.GetAppletAccessToken()` => 获取微信小程序全局唯一后台接口调用凭据
* `wechat.GetAppletPaidUnionId()` => 微信小程序用户支付完成后，获取该用户的 UnionId，无需用户授权
* `wechat.GetPublicUserInfo()` => 微信公众号：获取用户基本信息
* `wechat.GetPublicUserInfoBatch()` => 微信公众号：批量获取用户基本信息
* `wechat.DecryptOpenDataToStruct()` => 加密数据，解密到指定结构体
* `wechat.DecryptOpenDataToBodyMap()` => 加密数据，解密到 BodyMap
* `wechat.GetOpenIdByAuthCode()` => 授权码查询openid
* `wechat.GetOauth2AccessToken()` => 微信第三方登录，code 换取 access_token
* `wechat.RefreshOauth2AccessToken()` => 刷新微信第三方登录后，获取到的 access_token
* `wechat.CheckOauth2AccessToken()` => 检验授权凭证（access_token）是否有效
* `wechat.GetOauth2UserInfo()` => 微信开放平台：获取用户个人信息
* `wechat.DecryptRefundNotifyReqInfo()` => 解密微信退款异步通知的加密数据

---

### QQ支付API

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

### QQ公共API

* `qq.ParseNotifyToBodyMap()` => 解析QQ支付异步通知的结果到BodyMap
* `qq.ParseNotify()` => 解析QQ支付异步通知的参数
* `qq.VerifySign()` => QQ同步返回参数验签或异步通知参数验签

---

### 支付宝支付API

> #### 因支付宝接口太多，如没实现的接口，还请开发者自行调用 `client.PostAliPayAPISelfV2()`方法实现！请参考 `client_test.go` 内的 `TestClient_PostAliPayAPISelfV2()` 方法

> #### 希望有时间的伙伴儿Fork完后，积极提Pull Request，一起完善支付宝各个类别下的接口到相应的go文件中

* 支付宝接口自行实现方法：`client.PostAliPayAPISelfV2()`
* 网页&移动应用 - <font color='#027AFF' size='4'>支付API</font>
    * 统一收单交易支付接口（商家扫用户付款码）：`client.TradePay()`
    * 统一收单线下交易预创建（用户扫商品收款码）：`client.TradePrecreate()`
    * APP支付接口2.0（APP支付）：`client.TradeAppPay()`
    * 手机网站支付接口2.0（手机网站支付）：`client.TradeWapPay()`
    * 统一收单下单并支付页面接口（电脑网站支付）：`client.TradePagePay()`
    * 统一收单交易创建接口（小程序支付）：`client.TradeCreate()`
    * 统一收单线下交易查询: `client.TradeQuery()`
    * 统一收单交易撤销接口: `client.TradeCancel()`
    * 统一收单交易关闭接口: `client.TradeClose()`
    * 统一收单交易退款接口: `client.TradeRefund()`
    * 统一收单退款页面接口: `client.TradePageRefund()`
    * 统一收单交易退款查询: `client.TradeFastPayRefundQuery()`
    * 统一收单交易结算接口: `client.TradeOrderSettle()`
    * 支付宝订单信息同步接口: `client.TradeOrderInfoSync()`
    * ~~聚合支付订单咨询服务: `client.TradeAdvanceConsult()`(失效)~~
    * 花芝轻会员结算申请: `client.PcreditHuabeiAuthSettleApply()`
    * NFC用户卡信息同步: `client.CommerceTransportNfccardSend()`
    * 广告投放数据查询: `client.DataDataserviceAdDataQuery()`
    * 航司电话订票待申请接口: `client.CommerceAirCallcenterTradeApply()`
    * 网商银行全渠道收单业务订单创建: `client.PaymentTradeOrderCreate()`
    * 口碑订单预下单: `client.TradeOrderPrecreate()`
    * 口碑商品交易购买接口: `client.TradeItemorderBuy()`
* 网页&移动应用 - <font color='#027AFF' size='4'>资金API</font>
    * 单笔转账接口：`client.FundTransUniTransfer()`
    * 查询转账订单接口: `client.FundTransOrderQuery()`
    * 支付宝资金账户资产查询接口：`client.FundAccountQuery()`
    * 转账业务单据查询接口：`client.FundTransCommonQuery()`
    * 资金退回接口: `client.FundTransRefund()`
    * 资金授权冻结接口: `client.FundAuthOrderFreeze()`
    * 资金授权发码接口: `client.FundAuthOrderVoucherCreate()`
    * 线上资金授权冻结接口: client:FundAuthOrderAppFreeze()`
    * 资金授权解冻接口: `client.FundAuthOrderUnfreeze()`
    * 资金授权操作查询接口: `client.FundAuthOperationDetailQuery()`
    * 资金授权撤销接口: `client.FundAuthOperationCancel()`
    * 批次下单接口: `client.FundBatchCreate()`
    * 批量转账关单接口: `client.FundBatchClose()`
    * 批量转账明细查询接口: `client.FundBatchDetailQuery()`
    * 现金红包无线支付接口: `client.FundTransAppPay()`
    * 资金收款账号绑定关系查询: `client.FundTransPayeeBindQuery()`
* 网页&移动应用 - <font color='#027AFF' size='4'>会员API</font>
    * 支付宝会员授权信息查询接口（App支付宝登录）：`client.UserInfoShare()`
    * 身份认证初始化服务: `client.UserCertifyOpenInit()`
    * 身份认证开始认证（获取认证链接）: `client.UserCertifyOpenCertify()`
    * 身份认证记录查询: `client.UserCertifyOpenQuery()`
    * 支付宝个人协议页面签约接口: `client.UserAgreementPageSign()`
    * 支付宝个人代扣协议解约接口: `client.UserAgreementPageUnSign()`
    * 支付宝个人代扣协议查询接口: `client.UserAgreementQuery()`
    * 周期性扣款协议执行计划修改接口: `client.UserAgreementExecutionplanModify()`
    * 协议由普通通用代扣协议产品转移到周期扣协议产品: `client.UserAgreementTransfer()`
    * 通用当面付二阶段接口: `client.UserTwostageCommonUse()`
    * 芝麻企业征信基于身份的协议授权: `client.UserAuthZhimaorgIdentityApply()`
    * 查询是否在支付宝公益捐赠的接口: `client.UserCharityRecordexistQuery()`
    * 集分宝发放接口: `client.UserAlipaypointSend()`
    * isv 会员CRM数据回流: `client.MemberDataIsvCreate()`
    * 询家人信息档案(选人授权)组件已选的家人档案信息: `client.UserFamilyArchiveQuery()`
    * 初始化家人信息档案(选人授权)组件: `client.UserFamilyArchiveInitialize()`
    * 实名证件信息比对验证预咨询: `client.UserCertdocCertverifyPreconsult()`
    * 实名证件信息比对验证咨询: `client.UserCertdocCertverifyConsult()`
    * 初始化家庭芝麻GO共享组件: `client.UserFamilyShareZmgoInitialize()`
    * 数字分行银行码明细数据查询: `client.UserDtbankQrcodedataQuery()`
    * 查询集分宝预算库详情: `client.UserAlipaypointBudgetlibQuery()`
* 网页&移动应用 - <font color='#027AFF' size='4'>营销API</font>
    * 小程序生成推广二维码接口：`client.OpenAppQrcodeCreate()`
* 网页&移动应用 - <font color='#027AFF' size='4'>工具类API</font>
    * 用户登陆授权：`client.UserInfoAuth()`
    * 换取授权访问令牌：`client.SystemOauthToken()`
    * 换取应用授权令牌：`client.OpenAuthTokenApp()`
    * 应用支付宝公钥证书下载：`client.PublicCertDownload()`
* 网页&移动应用 - <font color='#027AFF' size='4'>芝麻信用API</font>
    * 芝麻企业信用信用评估初始化: `client.ZhimaCreditEpSceneRatingInitialize()`
    * 信用服务履约同步: `client.ZhimaCreditEpSceneFulfillmentSync()`
    * 加入信用服务: `clinet.ZhimaCreditEpSceneAgreementUse()`
    * 取消信用服务: `client.ZhimaCreditEpSceneAgreementCancel()`
    * 信用服务履约同步(批量): `client.ZhimaCreditEpSceneFulfillmentlistSync()`
    * 芝麻go用户数据回传: `client.ZhimaCreditPeZmgoCumulationSync()`
    * 商家芝麻GO累计数据回传接口: `client.ZhimaMerchantZmgoCumulateSync()`
    * 商家芝麻GO累计数据查询接口: `client.ZhimaMerchantZmgoCumulateQuery()`
    * 芝麻GO签约关单: `client.ZhimaCreditPeZmgoBizoptClose()`
    * 芝麻GO结算退款接口: `client.ZhimaCreditPeZmgoSettleRefund()`
    * 芝麻GO签约预创单: `client.ZhimaCreditPeZmgoPreorderCreate()`
    * 芝麻GO协议解约: `client.ZhimaCreditPeZmgoAgreementUnsign()`
    * 芝麻Go协议查询接口: `client.ZhimaCreditPeZmgoAgreementQuery()`
    * 芝麻Go解冻接口: `client.ZhimaCreditPeZmgoSettleUnfreeze()`
    * 芝麻GO支付下单链路签约申请: `client.ZhimaCreditPeZmgoPaysignApply()`
    * 芝麻GO支付下单链路签约确认: `client.ZhimaCreditPeZmgoPaysignConfirm()`
    * 职得工作证信息匹配度查询: `client.ZhimaCustomerJobworthAdapterQuery()`
    * 职得工作证外部渠道应用数据回流: `client.ZhimaCustomerJobworthSceneUse()`
* 网页&移动应用 - <font color='#027AFF' size='4'>财务API</font>
    * ~~支付宝商家账户当前余额查询：`client.DataBillBalanceQuery()`（失效）~~
    * 查询对账单下载地址：`client.DataBillDownloadUrlQuery()`
* 网页&移动应用 - <font color='#027AFF' size='4'>海关相关API</font>
    * 统一收单报关接口：`client.TradeCustomsDeclare()`
    * 报关接口：`client.AcquireCustoms()`
    * 报关查询接口：`client.AcquireCustomsQuery()`

### 支付宝公共API

* `alipay.GetCertSN()` => 获取证书SN号（app_cert_sn、alipay_cert_sn）
* `alipay.GetRootCertSN()` => 获取证书SN号（alipay_root_cert_sn）
* `alipay.GetRsaSign()` => 获取支付宝参数签名（参数sign值）
* `alipay.SystemOauthToken()` => 换取授权访问令牌（得到access_token，user_id等信息）
* `alipay.FormatPrivateKey()` => 格式化应用私钥
* `alipay.FormatPublicKey()` => 格式化支付宝公钥
* `alipay.FormatURLParam()` => 格式化支付宝请求URL参数
* `alipay.ParseNotifyToBodyMap()` => 解析支付宝支付异步通知的参数到BodyMap
* `alipay.ParseNotifyByURLValues()` => 通过 url.Values 解析支付宝支付异步通知的参数到BodyMap
* `alipay.VerifySign()` => 支付宝异步通知参数验签
* `alipay.VerifySignWithCert()` => 支付宝异步通知参数验签（证书方式）
* `alipay.VerifySyncSign()` => 支付宝同步返回参数验签
* `alipay.DecryptOpenDataToStruct()` => 解密支付宝开放数据到 结构体
* `alipay.DecryptOpenDataToBodyMap()` => 解密支付宝开放数据到 BodyMap
* `alipay.MonitorHeartbeatSyn()` => 验签接口

---

# 二、文档说明

* [GoPay 文档地址](https://pkg.go.dev/github.com/go-pay/gopay)
* QQ支付 使用方法请参考微信的
* 所有方法，如有问题，请仔细查看 wechat/client_test.go、alipay/client_test.go 或 examples
* 有问题请加QQ群（加群验证答案：gopay），微信加好友拉群（微信群有两个，一个活跃群，聊的内容比较杂，一个只聊技术群，平时很少说话，加好友后说明加哪个群，默认全邀请）。在此，非常感谢那些加群后，提出意见和反馈问题的同志们！

QQ群：
<img width="280" height="280" src="https://raw.githubusercontent.com/go-pay/gopay/main/qq_gopay.png"/>
加微信拉群：
<img width="280" height="280" src="https://raw.githubusercontent.com/go-pay/gopay/main/wechat_jerry.png"/>

---

## 1、初始化GoPay客户端并做配置（HTTP请求均默认设置tls.Config{InsecureSkipVerify: true}）

* #### 微信V3（推荐）

> 注意：V3 版本接口持续增加中，并未做沙箱支付，测试请用1分钱测试法

> 注意：`微信平台证书` 和 `微信平台证书序列号`，请自行通过 `wechat.GetPlatformCerts()` 方法维护

> 具体使用介绍，请参考 `gopay/wechat/v3/client_test.go`

```go
import (
    "github.com/go-pay/gopay/wechat/v3"
)

// NewClientV3 初始化微信客户端 V3
//	appid：appid 或者服务商模式的 sp_appid
//	mchid：商户ID 或者服务商模式的 sp_mchid
// 	serialNo：商户证书的证书序列号
//	apiV3Key：apiV3Key，商户平台获取
//	pkContent：私钥 apiclient_key.pem 读取后的内容
client, err = wechat.NewClientV3(Appid, MchId, SerialNo, APIv3Key, PKContent)
if err != nil {
    xlog.Error(err)
    return
}

// 设置微信平台证书和序列号，并启用自动同步返回验签
//	注意：请预先通过 wechat.GetPlatformCerts() 获取并维护微信平台证书和证书序列号
client.SetPlatformCert([]byte(WxPkContent), WxPkSerialNo).AutoVerifySign()

// 打开Debug开关，输出日志
client.DebugSwitch = gopay.DebugOn
```

* #### 微信V2

微信官方文档：[官方文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)

> 注意：微信支付下单等操作可用沙箱环境测试是否成功，但真正支付时，请使用正式环境 `isProd = true`，不然会报错。

> 微信证书二选一：只传 `apiclient_cert.pem` 和 `apiclient_key.pem` 或者只传 `apiclient_cert.p12`

```go
import (
    "github.com/go-pay/gopay/wechat"
)

// 初始化微信客户端
//    appId：应用ID
//    mchId：商户ID
//    apiKey：API秘钥值
//    isProd：是否是正式环境
client := wechat.NewClient("wxdaa2ab9ef87b5497", mchId, apiKey, false)

// 打开Debug开关，输出请求日志，默认关闭
client.DebugSwitch = gopay.DebugOn

// 设置国家：不设置默认 中国国内
//    wechat.China：中国国内
//    wechat.China2：中国国内备用
//    wechat.SoutheastAsia：东南亚
//    wechat.Other：其他国家
client.SetCountry(wechat.China)

// 添加微信pem证书
client.AddCertPemFilePath()
client.AddCertPemFileContent()

// 添加微信pkcs12证书
client.AddCertPkcs12FilePath()
client.AddCertPkcs12FileContent()
```

* #### 支付宝

支付宝官方文档：[官方文档](https://openhome.alipay.com/docCenter/docCenter.htm)

支付宝RSA秘钥生成文档：[生成RSA密钥](https://opendocs.alipay.com/open/291/105971) （推荐使用 RSA2）

技术支持 & 案例 FAQ：[秘钥问题](https://opendocs.alipay.com/support/01rauw)

沙箱环境使用说明：[文档地址](https://opendocs.alipay.com/open/200/105311)

```go
import (
    "github.com/go-pay/gopay/alipay"
)

// 初始化支付宝客户端
//    appId：应用ID
//    privateKey：应用私钥，支持PKCS1和PKCS8
//    isProd：是否是正式环境
client := alipay.NewClient("2016091200494382", privateKey, false)

// 打开Debug开关，输出日志，默认关闭
client.DebugSwitch = gopay.DebugOn

// 设置支付宝请求 公共参数
//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
client.SetLocation().                       // 设置时区，不设置或出错均为默认服务器时间
    SetPrivateKeyType().                    // 设置 支付宝 私钥类型，alipay.PKCS1 或 alipay.PKCS8，默认 PKCS1
    SetAliPayRootCertSN().                  // 设置支付宝根证书SN，通过 alipay.GetRootCertSN() 获取
    SetAppCertSN().                         // 设置应用公钥证书SN，通过 alipay.GetCertSN() 获取
    SetAliPayPublicCertSN().                // 设置支付宝公钥证书SN，通过 alipay.GetCertSN() 获取
    SetCharset("utf-8").                    // 设置字符编码，不设置默认 utf-8
    SetSignType(alipay.RSA2).               // 设置签名类型，不设置默认 RSA2
    SetReturnUrl("https://www.fmm.ink").    // 设置返回URL
    SetNotifyUrl("https://www.fmm.ink").    // 设置异步通知URL
    SetAppAuthToken()                       // 设置第三方应用授权

// 自动同步验签（只支持证书模式）
// 传入 alipayCertPublicKey_RSA2.crt 内容
client.AutoVerifySign("alipayCertPublicKey_RSA2 bytes")

// 证书路径
err := client.SetCertSnByPath("appCertPublicKey.crt", "alipayRootCert.crt", "alipayCertPublicKey_RSA2.crt")
// 证书内容
err := client.SetCertSnByContent("appCertPublicKey bytes", "alipayRootCert bytes", "alipayCertPublicKey_RSA2 bytes")
```

## 2、初始化并赋值BodyMap（client的方法所需的入参）

* #### 微信请求参数
  * 微信V2接口通用参数（mch_id、appid、sign）无需传入，client 请求时会默认处理

具体参数请根据不同接口查看：[微信支付接口文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)

```go
import (
    "github.com/go-pay/gopay/pkg/util"
    "github.com/go-pay/gopay/wechat"
)

// 初始化 BodyMap
bm := make(gopay.BodyMap)
bm.Set("nonce_str", util.GetRandomString(32)).
    Set("body", "H5支付").
    Set("out_trade_no", number).
    Set("total_fee", 1).
    Set("spbill_create_ip", "127.0.0.1").
    Set("notify_url", "https://www.fmm.ink").
    Set("trade_type", TradeType_H5).
    Set("device_info", "WEB").
    Set("sign_type", SignType_MD5).
    SetBodyMap("scene_info", func(bm gopay.BodyMap) {
        bm.SetBodyMap("h5_info", func(bm gopay.BodyMap) {
            bm.Set("type", "Wap")
            bm.Set("wap_url", "https://www.fmm.ink")
            bm.Set("wap_name", "H5测试支付")
        })
    }) /*.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")*/

// 参数 sign ，可单独生成赋值到BodyMap中；也可不传sign参数，client内部会自动获取
// 如需单独赋值 sign 参数，需通过下面方法，最后获取sign值并在最后赋值此参数
sign := wechat.GetParamSign("wxdaa2ab9ef87b5497", mchId, apiKey, body)
// sign, _ := wechat.GetSanBoxParamSign("wxdaa2ab9ef87b5497", mchId, apiKey, body)
bm.Set("sign", sign)
```

* #### 支付宝请求参数

具体参数请根据不同接口查看：[支付宝支付API接口文档](https://opendocs.alipay.com/apis/api_1/alipay.trade.wap.pay)

```go
// 初始化 BodyMap
bm := make(gopay.BodyMap)
bm.Set("subject", "手机网站测试支付").
    Set("out_trade_no", "GZ201909081743431443").
    Set("quit_url", "https://www.fmm.ink").
    Set("total_amount", "100.00").
    Set("product_code", "QUICK_WAP_WAY")
```

## 3、client 方法调用

* #### 微信V3 client（推荐）

```go
// 公共方法
client.SetPlatformCert()
client.AutoVerifySign()
client.V3EncryptText()
client.V3DecryptText()

// 直连商户
wxRsp, err := client.V3TransactionApp(bm)
wxRsp, err := client.V3TransactionJsapi(bm)
wxRsp, err := client.V3TransactionNative(bm)
wxRsp, err := client.V3TransactionH5(bm)
wxRsp, err := client.V3TransactionQueryOrder(bm)
wxRsp, err := client.V3TransactionCloseOrder(bm)

// 服务商
wxRsp, err := client.V3PartnerTransactionApp(bm)
wxRsp, err := client.V3PartnerTransactionJsapi(bm)
wxRsp, err := client.V3PartnerTransactionNative(bm)
wxRsp, err := client.V3PartnerTransactionH5(bm)
wxRsp, err := client.V3PartnerQueryOrder(bm)
wxRsp, err := client.V3PartnerCloseOrder(bm)

// 合单
wxRsp, err := client.V3CombineTransactionApp(bm)
wxRsp, err := client.V3CombineTransactionJsapi(bm)
wxRsp, err := client.V3CombineTransactionNative(bm)
wxRsp, err := client.V3CombineTransactionH5(bm)
wxRsp, err := client.V3CombineQueryOrder(bm)
wxRsp, err := client.V3CombineCloseOrder(bm)
...
```

* #### 微信V2 client

```go
wxRsp, err := client.UnifiedOrder(bm)
wxRsp, err := client.Micropay(bm)
wxRsp, err := client.QueryOrder(bm)
wxRsp, err := client.CloseOrder(bm)
wxRsp, err := client.Reverse(bm)
wxRsp, err := client.Refund(bm)
wxRsp, err := client.QueryRefund(bm)
wxRsp, err := client.DownloadBill(bm)
wxRsp, err := client.DownloadFundFlow(bm)
wxRsp, err := client.BatchQueryComment(bm)
wxRsp, err := client.Transfer(bm)
...
```

* #### 支付宝 client

```go
// 手机网站支付是通过服务端获取支付URL后，然后返回给客户端，请求URL地址即可打开支付页面
payUrl, err := client.TradeWapPay(bm)

// 电脑网站支付是通过服务端获取支付URL后，然后返回给客户端，请求URL地址即可打开支付页面
payUrl, err := client.TradePagePay(bm)

// APP支付是通过服务端获取支付参数后，然后通过Android/iOS客户端的SDK调用支付功能
payParam, err := client.TradeAppPay(bm)

// 商家使用扫码枪等条码识别设备扫描用户支付宝钱包上的条码/二维码，完成收款
aliRsp, err := client.TradePay(bm)

// 支付宝小程序支付时 buyer_id 为必传参数，需要提前获取，获取方法如下两种
//  1、alipay.SystemOauthToken()     返回取值：aliRsp.SystemOauthTokenResponse.UserId
//  2、client.SystemOauthToken()     返回取值：aliRsp.SystemOauthTokenResponse.UserId
aliRsp, err := client.TradeCreate(bm)
aliRsp, err := client.TradeQuery(bm)
aliRsp, err := client.TradeClose(bm)
aliRsp, err := client.TradeCancel(bm)
aliRsp, err := client.TradeRefund(bm)
aliRsp, err := client.TradePageRefund(bm)
aliRsp, err := client.TradeFastPayRefundQuery(bm)
aliRsp, err := client.TradeOrderSettle(bm)
aliRsp, err := client.TradePrecreate(bm)
aliRsp, err := client.FundTransUniTransfer(bm)
aliRsp, err := client.FundTransCommonQuery(bm)
aliRsp, err := client.FundAccountQuery(bm)
aliRsp, err := client.SystemOauthToken(bm)
aliRsp, err := client.OpenAuthTokenApp(bm)
aliRsp, err := client.UserCertifyOpenInit(bm)
aliRsp, err := client.UserCertifyOpenCertify(bm)
aliRsp, err := client.UserCertifyOpenQuery(bm)
...
```

## 4、微信统一下单后，获取微信小程序支付、APP支付、微信内H5支付所需要的 paySign

* #### 微信V3（推荐）

小程序调起支付API：[小程序调起支付API](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_5_4.shtml)

APP调起支付API：[APP调起支付API](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_2_4.shtml)

JSAPI调起支付API：[JSAPI调起支付API](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_4.shtml)

H5调起支付API：[H5调起支付API](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_3_4.shtml)

```go
// jsapi
jsapi, err := client.PaySignOfJSAPI("prepayid")
// app
app, err := client.PaySignOfApp("prepayid")
// 小程序
applet, err := client.PaySignOfApplet("prepayid")
```

* #### 微信V2

微信小程序支付官方文档：[微信小程序支付API](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/payment/wx.requestPayment.html)

APP支付官方文档：[APP端调起支付的参数列表文档](https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_12)

微信内H5支付官方文档：[微信内H5支付文档](https://pay.weixin.qq.com/wiki/doc/api/wxpay/ch/pay/OfficialPayMent/chapter5_5.shtml)

```go
import (
    "github.com/go-pay/gopay/wechat"
)

// ====微信小程序 paySign====
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
packages := "prepay_id=" + wxRsp.PrepayId   // 此处的 wxRsp.PrepayId ,统一下单成功后得到
// 获取微信小程序支付的 paySign
//    appId：AppID
//    nonceStr：随机字符串
//    packages：统一下单成功后拼接得到的值
//    signType：签名方式，务必与统一下单时用的签名方式一致
//    timeStamp：时间
//    apiKey：API秘钥值
paySign := wechat.GetMiniPaySign(AppID, wxRsp.NonceStr, packages, wechat.SignType_MD5, timeStamp, apiKey)

// ====APP支付 paySign====
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
// 获取APP支付的 paySign
// 注意：package 参数因为是固定值，无需开发者再传入
//    appId：AppID
//    partnerid：partnerid
//    nonceStr：随机字符串
//    prepayId：统一下单成功后得到的值
//    signType：签名方式，务必与统一下单时用的签名方式一致
//    timeStamp：时间
//    apiKey：API秘钥值
paySign := wechat.GetAppPaySign(appid, partnerid, wxRsp.NonceStr, wxRsp.PrepayId, wechat.SignType_MD5, timeStamp, apiKey)

// ====微信内H5支付 paySign====
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
packages := "prepay_id=" + wxRsp.PrepayId   // 此处的 wxRsp.PrepayId ,统一下单成功后得到
// 获取微信内H5支付 paySign
//    appId：AppID
//    nonceStr：随机字符串
//    packages：统一下单成功后拼接得到的值
//    signType：签名方式，务必与统一下单时用的签名方式一致
//    timeStamp：时间
//    apiKey：API秘钥值
paySign := wechat.GetH5PaySign(AppID, wxRsp.NonceStr, packages, wechat.SignType_MD5, timeStamp, apiKey)
```

## 5、同步返回参数验签Sign、异步通知参数解析和验签Sign、异步通知返回

异步参数需要先解析，解析出来的结构体或BodyMap再验签

[Gin Web框架](https://github.com/gin-gonic/gin)

[Echo Web框架](https://github.com/labstack/echo)

异步通知处理完后，需回复平台固定数据

* #### 微信V3（推荐）

```go
import (
    "github.com/go-pay/gopay/wechat"
    "github.com/go-pay/gopay/pkg/xlog"
)

// ========同步微信V3支付验签========
// 如已开启自动验签，则无需手动验签操作
wxRsp, err := client.V3TransactionJsapi(bm)
if err != nil {
    xlog.Error(err)
    return
}
err = wechat.V3VerifySign(wxRsp.SignInfo.HeaderTimestamp, wxRsp.SignInfo.HeaderNonce, wxRsp.SignInfo.SignBody, wxRsp.SignInfo.HeaderSignature, WxPkContent)
if err != nil {
    xlog.Error(err)
    return
}

// ========异步通知验签========
notifyReq, err := wechat.V3ParseNotify()
if err != nil {
    xlog.Error(err)
    return
}
// WxPkContent 是通过 wechat.GetPlatformCerts() 接口向微信获取的微信平台公钥证书内容
err = notifyReq.VerifySign(WxPkContent)
if err != nil {
    xlog.Error(err)
    return
}
// ========异步通知解密========
// 普通支付通知解密
result, err := notifyReq.DecryptCipherText(apiV3Key)
// 合单支付通知解密
result, err := notifyReq.DecryptCombineCipherText(apiV3Key)
// 退款通知解密
result, err := notifyReq.DecryptRefundCipherText(apiV3Key)

// ========异步通知应答========
// 退款通知http应答码为200且返回状态码为SUCCESS才会当做商户接收成功，否则会重试。
// 注意：重试过多会导致微信支付端积压过多通知而堵塞，影响其他正常通知。

// 此写法是 gin 框架返回微信的写法
c.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})
// 此写法是 echo 框架返回微信的写法
return c.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})
```

* #### 微信V2

```go
import (
    "github.com/go-pay/gopay"
    "github.com/go-pay/gopay/wechat"
)

// ====同步返回参数验签Sign====
wxRsp, err := client.UnifiedOrder(bm)
// 微信同步返回参数验签或异步通知参数验签
//    apiKey：API秘钥值
//    signType：签名类型（调用API方法时填写的类型）
//    bean：微信同步返回的结构体 wxRsp 或 异步通知解析的结构体 notifyReq
//    返回参数 ok：是否验签通过
//    返回参数 err：错误信息
ok, err := wechat.VerifySign(apiKey, wechat.SignType_MD5, wxRsp)

// ====支付异步通知参数解析和验签Sign====
// 解析支付异步通知的参数
//    req：*http.Request
//    ctx.Request   是 gin 框架的获取 *http.Request
//    ctx.Request() 是 echo 框架的获取 *http.Request
//    返回参数 notifyReq：通知的参数
//    返回参数 err：错误信息
notifyReq, err := wechat.ParseNotifyToBodyMap(ctx.Request)

// 验签操作
ok, err := wechat.VerifySign(apiKey, wechat.SignType_MD5, notifyReq)

// ====退款异步通知参数解析，退款通知无sign，不用验签====
// 
// 解析退款异步通知的参数，解析出来的 req_info 是加密数据，需解密
//    req：*http.Request
//    ctx.Request   是 gin 框架的获取 *http.Request
//    ctx.Request() 是 echo 框架的获取 *http.Request
//    返回参数 notifyReq：通知的参数
//    返回参数 err：错误信息
notifyReq, err := wechat.ParseNotifyToBodyMap(c.Request)
 或
notifyReq, err := wechat.ParseRefundNotify(c.Request)

// ==解密退款异步通知的加密参数 req_info ==
refundNotify, err := wechat.DecryptRefundNotifyReqInfo(notifyReq.ReqInfo, apiKey)

// ==异步通知，返回给微信平台的信息==
rsp := new(wechat.NotifyResponse) // 回复微信的数据
rsp.ReturnCode = gopay.SUCCESS
rsp.ReturnMsg = gopay.OK

// 此写法是 gin 框架返回微信的写法
c.String(http.StatusOK, "%s", rsp.ToXmlString())
// 此写法是 echo 框架返回微信的写法
return c.String(http.StatusOK, rsp.ToXmlString())
```

* #### 支付宝

注意：APP支付、手机网站支付、电脑网站支付 暂不支持同步返回验签

支付宝支付后的同步/异步通知验签文档：[支付结果通知](https://opendocs.alipay.com/open/200/106120)

```go
import (
    "github.com/go-pay/gopay/alipay"
)

// ====同步返回参数验签Sign====
// 如已开启自动验签，则无需手动验签操作
aliRsp, err := client.TradePay(bm)
// 支付宝同步返回验签
//    注意：APP支付，手机网站支付，电脑网站支付 暂不支持同步返回验签
//    aliPayPublicKey：支付宝平台获取的支付宝公钥
//    signData：待验签参数，aliRsp.SignData
//    sign：待验签sign，aliRsp.Sign
//    返回参数ok：是否验签通过
//    返回参数err：错误信息
ok, err := alipay.VerifySyncSign(aliPayPublicKey, aliRsp.SignData, aliRsp.Sign)
//    aliPayPublicKeyCert：支付宝公钥证书存放路径 alipayCertPublicKey_RSA2.crt 或文件内容[]byte
ok, err := alipay.VerifySyncSignWithCert(aliPayPublicKeyCert, aliRsp.SignData, aliRsp.Sign)

// ====异步通知参数解析和验签Sign====
// 解析异步通知的参数
//    req：*http.Request
//    返回参数 notifyReq：通知的参数
//    返回参数 err：错误信息
notifyReq, err = alipay.ParseNotifyToBodyMap(c.Request())     // c.Request()是 echo 框架的获取
 或
notifyReq, err = alipay.ParseNotifyByURLValues()

// 验签操作
ok, err = alipay.VerifySign(aliPayPublicKey, notifyReq)
// 证书验签操作
ok, err = alipay.VerifySignWithCert("alipayCertPublicKey_RSA2.crt", notifyReq)

// ==异步通知，返回支付宝平台的信息==
//    文档：https://opendocs.alipay.com/open/203/105286
//    程序执行完后必须打印输出“success”（不包含引号）。如果商户反馈给支付宝的字符不是success这7个字符，支付宝服务器会不断重发通知，直到超过24小时22分钟。一般情况下，25小时以内完成8次通知（通知的间隔频率一般是：4m,10m,10m,1h,2h,6h,15h）

// 此写法是 gin 框架返回支付宝的写法
c.String(http.StatusOK, "%s", "success")
// 此写法是 echo 框架返回支付宝的写法
return c.String(http.StatusOK, "success")
```

## 6、微信、支付宝 公共API（仅部分说明）

* #### 微信V3 公共API

微信敏感信息加解密、回调接口敏感信息解密

```go
import (
    "github.com/go-pay/gopay/wechat/v3"
)

// 获取微信平台证书和序列号信息
wechat.GetPlatformCerts()
// 请求参数 敏感信息加密
wechat.V3EncryptText() 或 client.V3EncryptText()
// 返回参数 敏感信息解密
wechat.V3DecryptText() 或 client.V3DecryptText()
// 回调通知敏感信息解密
wechat.V3DecryptNotifyCipherText()
wechat.V3DecryptRefundNotifyCipherText()
wechat.V3DecryptCombineNotifyCipherText()
...
```

* #### 微信V2 公共API

官方文档：[code2Session](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html)

button按钮获取手机号码：[button组件文档](https://developers.weixin.qq.com/miniprogram/dev/component/button.html)

微信解密算法文档：[解密算法文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)

```go
import (
    "github.com/go-pay/gopay/wechat"
)

// 获取微信小程序用户的OpenId、SessionKey、UnionId
//    appId：微信小程序的APPID
//    appSecret：微信小程序的AppSecret
//    wxCode：小程序调用wx.login 获取的code
sessionRsp, err := wechat.Code2Session(appId, appSecret, wxCode)

// ====解密微信加密数据到指定结构体====

//小程序获取手机号
data := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
iv := "Cds8j3VYoGvnTp1BrjXdJg=="
session := "lyY4HPQbaOYzZdG+JcYK9w=="
phone := new(wechat.UserPhone)
// 解密开放数据
//    encryptedData：包括敏感数据在内的完整用户信息的加密数据，小程序获取到
//    iv：加密算法的初始向量，小程序获取到
//    sessionKey：会话密钥，通过 wechat.Code2Session() 方法获取到
//    beanPtr：需要解析到的结构体指针，操作完后，声明的结构体会被赋值
err := wechat.DecryptOpenDataToStruct(data, iv, session, phone)
xlog.Debug(*phone)
// 获取微信小程序用户信息
sessionKey := "tiihtNczf5v6AKRyjwEUhQ=="
encryptedData := "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="
iv2 := "r7BXXKkLb8qrSNn05n0qiA=="

// 微信小程序 用户信息
userInfo := new(wechat.AppletUserInfo)
err = wechat.DecryptOpenDataToStruct(encryptedData, iv2, sessionKey, userInfo)
xlog.Debug(*userInfo)

data := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
iv := "Cds8j3VYoGvnTp1BrjXdJg=="
session := "lyY4HPQbaOYzZdG+JcYK9w=="
    
// 解密开放数据到 BodyMap
//    encryptedData:包括敏感数据在内的完整用户信息的加密数据
//    iv:加密算法的初始向量
//    sessionKey:会话密钥
bm, err := wechat.DecryptOpenDataToBodyMap(data, iv, session)
if err != nil {
    xlog.Debug("err:", err)
    return
}
xlog.Debug("WeChatUserPhone:", bm)
```

* #### 支付宝 公共API

支付宝换取授权访问令牌文档：[换取授权访问令牌](https://opendocs.alipay.com/apis/api_9/alipay.system.oauth.token)

获取用户手机号文档：[获取用户手机号](https://opendocs.alipay.com/mini/api/getphonenumber)

支付宝加解密文档：[AES配置文档](https://opendocs.alipay.com/mini/introduce/aes) ，[AES加解密文档](https://opendocs.alipay.com/open/common/104567)
```go
import (
    "github.com/go-pay/gopay/alipay"
)

// 换取授权访问令牌（默认使用utf-8，RSA2）
//    appId：应用ID
//    privateKey：应用私钥，支持PKCS1和PKCS8
//    grantType：值为 authorization_code 时，代表用code换取；值为 refresh_token 时，代表用refresh_token换取，传空默认code换取
//    codeOrToken：支付宝授权码或refresh_token
rsp, err := alipay.SystemOauthToken(appId, privateKey, grantType, codeOrToken)

// 解密支付宝开放数据带到指定结构体
//    以小程序获取手机号为例
phone := new(alipay.UserPhone)
// 解密支付宝开放数据
//    encryptedData:包括敏感数据在内的完整用户信息的加密数据
//    secretKey:AES密钥，支付宝管理平台配置
//    beanPtr:需要解析到的结构体指针
err := alipay.DecryptOpenDataToStruct(encryptedData, secretKey, phone)
xlog.Debug(*phone)
```

## 赞赏多少是您的心意，感谢！

<font color='#07C160' size='4'>微信：</font>
<img width="200" height="200" src="https://raw.githubusercontent.com/go-pay/gopay/main/zanshang_wx.png"/>
<font color='#027AFF' size='4'>支付宝：</font>
<img width="200" height="200" src="https://raw.githubusercontent.com/go-pay/gopay/main/zanshang_zfb.png"/>

## License

```
Copyright 2019 Jerry

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```