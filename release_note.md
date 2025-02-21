## 版本号：v1.5.109

* 修改记录：
  * gopay：golang.org/x/crypto 版本升级到 v0.33.0。
  * 支付宝：作废 client.UserAgreementPageSignInApp() 方法，使用 client.UserAgreementPageSignInQRCode() 方法替换。
  * 支付宝：修改 client.FundTransAppPay() 方法返回参数。
  * PayPal：新增 Payment Method Tokens 相关接口。
  * PayPal：新增 client.WithoutAutoRefreshToken() 方法，不自动刷新token。
  * 支付宝V3：新增 商家转账 相关的接口。
    * client.FundAccountQuery()，支付宝资金账户资产查询接口。
    * client.FundQuotaQuery()，转账额度查询接口。
    * client.FundTransUniTransfer()，单笔转账接口。
    * client.DataBillEreceiptApply()，申请电子回单(incubating)。
    * client.DataBillEreceiptQuery()，查询电子回单状态(incubating)。
    * client.FundTransCommonQuery()，转账业务单据查询接口。
    * client.FundTransMultistepTransfer()，多步转账创建并支付。
    * client.FundTransMultistepQuery()，多步转账查询接口。

## 版本号：v1.5.108

* 修改记录：
  * gopay：golang.org/x/crypto 版本升级到 v0.32.0。
  * 微信V3：支持设置代理 Host 地址，client.SetProxyHost()，wechat.SetProxyHost()。
  * 微信V3：已支持新版本商家转账相关接口。
    * client.V3TransferBills()，发起转账。
    * client.V3TransferBillsCancel()，撤销转账。
    * client.V3TransferBillsMerchantQuery()，商户单号查询转账单。
    * client.V3TransferBillsQuery()，微信单号查询转账单。
    * client.V3TransferElecsignMerchant()，商户单号申请电子回单。
    * client.V3TransferElecsign()，微信单号申请电子回单。
    * client.V3TransferElecsignQuery()，微信单号查询电子回单。
    * client.V3TransferElecsignMerchantQuery()，商户单号查询电子回单。
    * notify.DecryptTransferBillsNotifyCipherText()，解密 新版商家转账通知 回调中的加密参数。
  * 支付宝V3：支持设置代理 Host 地址，client.SetProxyHost()。
  * 补充部分支付宝V3接口。
  * 支付宝V3：处理公共参数方法 a.pubParamsHandle() 内逻辑变更，优先判断非body map内的参数赋值。

## 版本号：v1.5.107

* 修改记录：
  * gopay：golang.org/x/crypto 版本升级到 v0.31.0。
  * gopay：xhttp transport 的默认 MaxIdleConnsPerHost 设置为 1000，MaxConnsPerHost 设置为 3000。
  * 微信V3：新增掌纹支付相关接口。
  * 微信V3：补充 支付有礼相关接口。
  * 微信V3：补充 电子发票相关接口。
  * 支付宝V3：补充 支付宝的若干v3接口。

## 版本号：v1.5.106

* 修改记录：
  * 支付宝：支付宝支持V3接口，接口还在完善中，欢迎提PR一起建设。
  * gopay：支付宝、微信接口请求支持设置自定义RequestId的函数，client.SetRequestIdFunc()。
  * gopay：修复部分已知问题。

## 版本号：v1.5.105

* 修改记录：
  * gopay：golang.org/x/crypto v0.26.0 版本升级到 v0.27.0。
  * 支付宝：新增 client.MarketingCampaignCashCreate()，创建现金活动接口。
  * 支付宝：新增 client.MarketingCampaignCashTrigger()，触发现金红包活动接口。
  * 支付宝：新增 client.MarketingCampaignCashStatusModify()，更改现金活动状态接口。
  * 支付宝：新增 client.MarketingCampaignCashListQuery()，现金活动列表查询接口。
  * 支付宝：新增 client.MarketingCampaignCashDetailQuery()，现金活动详情查询接口。
  * 支付宝：新增 client.MerchantQipanCrowdCreate()，上传创建人群接口。
  * 支付宝：新增 client.MerchantQipanCrowdUserAdd()，人群中追加用户接口。
  * 支付宝：新增 client.MerchantQipanCrowdUserDelete()，人群中删除用户接口。
  * 支付宝：新增 client.MarketingQipanTagBaseBatchQuery()，棋盘人群圈选标签基本信息查询接口。
  * 支付宝：新增 client.MarketingQipanTagQuery()，棋盘标签圈选值查询接口。
  * 支付宝：新增 client.MarketingQipanCrowdOperationCreate()，棋盘人群创建接口。
  * 支付宝：新增 client.MarketingQipanCrowdTagQuery()，查询圈选标签列表接口。
  * 支付宝：新增 client.MarketingQipanCrowdWithTagCreate()，标签圈选创建人群接口。
  * 支付宝：新增 client.MarketingQipanCrowdWithTagQuery()，标签圈选预估人群规模接口。
  * 支付宝：新增 client.MarketingQipanCrowdBatchQuery()，查询人群列表接口。
  * 支付宝：新增 client.MarketingQipanCrowdQuery()，查询人群详情接口。
  * 支付宝：新增 client.MarketingQipanCrowdModify()，修改人群接口。
  * 支付宝：新增 client.MarketingQipanBoardQuery()，看板分析接口。
  * 支付宝：新增 client.MarketingQipanInsightQuery()，画像分析接口。
  * 支付宝：新增 client.MarketingQipanBehaviorQuery()，行为分析接口。
  * 支付宝：新增 client.MarketingQipanTrendQuery()，趋势分析接口。
  * 支付宝：新增 client.MarketingQipanInsightCityQuery()，常住省市查询接口。
  * PayPal：新增 client.AddTrackingNumber()，添加物流单号接口。
  * PayPal：优化部分结构体字段。。
  * 支付宝：FundTransCommonQuery 结构体补充字段。
  * Apple：client.SendConsumptionInformation() 方法成功状态码判断修改。
  * gopay：升级内部依赖module库。

## 版本号：Release 1.5.104

* 修改记录：
  * gopay：golang.org/x/crypto v0.24.0 版本升级到 v0.26.0。
  * 支付宝：client.TradeOrderSettle()，response 完善字段。
  * 微信V3：新增 client.V3QQTransactionH5()，QQ小程序H5下单。
  * 微信V3：新增 client.V3CombineQQTransactionH5()，合单QQ小程序下单-H5。
  * 微信V3：新增 wechat.V3DecryptViolationNotifyCipherText()，解密 服务商子商户处置记录 回调中的加密信息。#412
  * 微信V3：新增 V3NotifyReq method req.DecryptViolationCipherText()，解密 服务商子商户处置记录 回调中的加密信息。
  * 微信V3：新增 wechat.V3DecryptTransferBatchNotifyCipherText()，解密 商家转账批次回调通知 回调中的加密信息。
  * 微信V3：新增 V3NotifyReq method req.DecryptTransferBatchCipherText()，解密 商家转账批次回调通知 回调中的加密信息。
  * 支付宝：新增 client.MarketingCardTemplateCreate()，会员卡模板创建接口。
  * 支付宝：新增 client.MarketingCardTemplateModify()，会员卡模板修改接口。
  * 支付宝：新增 client.MarketingCardTemplateQuery()，会员卡模板查询接口。
  * 支付宝：新增 client.MarketingCardUpdate()，会员卡更新接口。
  * 支付宝：新增 client.MarketingCardQuery()，会员卡查询接口。
  * 支付宝：新增 client.MarketingCardDelete()，会员卡删卡接口。
  * 支付宝：新增 client.MarketingCardMessageNotify()，会员卡消息通知接口。
  * 支付宝：新增 client.MarketingCardFormTemplateSet()，会员卡开卡表单模板配置接口。
  * 支付宝：新增 client.OfflineMaterialImageUpload()，上传门店照片和视频接口。

## 版本号：Release 1.5.103

* 修改记录：
  * gopay：xlog 库更新，支持接口自定义logger。
  * gopay：go mod 版本升级至 1.21。
  * 支付宝：新增 client.PayAppMarketingConsult()，商户前置内容咨询接口。
  * 支付宝：新增 client.MarketingActivityOrderVoucherCreate()，创建商家券活动接口
  * 支付宝：新增 client.MarketingActivityOrderVoucherCodeDeposit()，同步商家券券码接口
  * 支付宝：新增 client.MarketingActivityOrderVoucherModify()，修改商家券活动基本信息接口
  * 支付宝：新增 client.MarketingActivityOrderVoucherStop()，停止商家券活动接口
  * 支付宝：新增 client.MarketingActivityOrderVoucherAppend()，修改商家券活动发券数量上限接口
  * 支付宝：新增 client.MarketingActivityOrderVoucherUse()，同步券核销状态接口
  * 支付宝：新增 client.MarketingActivityOrderVoucherRefund()，取消券核销状态接口
  * 支付宝：新增 client.MarketingActivityOrderVoucherQuery()，查询商家券活动接口
  * 支付宝：新增 client.MarketingActivityOrderVoucherCodeCount()，统计商家券券码数量接口

## 版本号：Release 1.5.102

* 修改记录：
  * gopay：golang.org/x/crypto v0.23.0 版本升级到 v0.24.0。
  * 微信V3：client.V3Refund() 返回值修复错误结构。
  * 支付宝：新增 client.AntMerchantShopPageQuery()，店铺分页查询接口。
  * 支付宝：新增 client.AntMerchantExpandMccQuery()，商户mcc信息查询接口。
  * 支付宝：新增 client.AntMerchantExpandShopReceiptAccountSave()，店铺增加收单账号接口。
  * 支付宝：新增 client.DataBillEreceiptApply()，申请电子回单(incubating)接口。
  * 支付宝：新增 client.DataBillEreceiptQuery()，查询电子回单状态(incubating)接口。
  * 支付宝：新增 client.FaceVerificationInitialize()，APP人脸核身初始化接口。
  * 支付宝：新增 client.FaceVerificationQuery()，APP人脸核身结果查询接口。
  * 支付宝：新增 client.FaceCertifyInitialize()，H5人脸核身初始化接口。
  * 支付宝：新增 client.FaceCertifyVerify()，H5人脸核身开始认证接口。
  * 支付宝：新增 client.FaceCertifyQuery()，H5人脸核身查询记录接口。
  * 支付宝：新增 client.FaceSourceCertify()，纯服务端人脸核身接口。
  * 支付宝：新增 client.FaceCheckInitialize()，人脸检测初始化接口。
  * 支付宝：新增 client.FaceCheckQuery()，人脸检测结果数据查询接口。
  * 支付宝：新增 client.OcrServerDetect()，服务端OCR接口。
  * 支付宝：新增 client.OcrMobileInitialize()，App端OCR初始化接口。
  * 支付宝：新增 client.OcrCommonDetect()，文字识别OCR接口。

## 版本号：Release 1.5.101

* 修改记录：
  * gopay：golang.org/x/crypto v0.20.0 版本升级到 v0.23.0。
  * gopay：xhttp移除外部依赖，修复文件上传相关接口的Bug。
  * 微信V3：新增 client.V3AbnormalRefund()，发起异常退款。
  * 微信V3：新增 client.V3FavorCallbackUrl()，查询消息通知地址。
  * 微信V3：新增 client.V3ComplaintImage()，商户反馈图片请求接口。
  * 微信V3：修改 wechat.V3DecryptNotifyCipherText() -> wechat.V3DecryptPayNotifyCipherText()，解密 普通支付 回调中的加密信息
  * 微信V3：修改 wechat.DecryptPartnerCipherText() -> wechat.DecryptPartnerPayCipherText()，解密 普通支付 回调中的加密信息
  * 微信V3：新增 wechat.V3DecryptNotifyCipherTextToStruct()，解密 统一数据 到指针结构体对象。
  * 支付宝：新增 client.ZolozAuthenticationSmilepayInit()，刷脸支付初始化接口。
  * 支付宝：新增 client.ZolozAuthenticationCustomerFtokenQuery()，查询刷脸结果信息接口。
  * 支付宝：新增 client.MarketingActivityDeliveryChanged()，推广计划状态变更消息接口。
  * 支付宝：新增 client.MarketingActivityDeliveryCreate()，创建推广计划接口。
  * 支付宝：新增 client.MarketingActivityDeliveryQuery()，查询推广计划接口。
  * 支付宝：新增 client.MarketingActivityDeliveryStop()，停止推广计划接口。
  * 支付宝：新增 client.MarketingMaterialImageUpload()，营销图片资源上传接口。
  * 支付宝：新增 client.AntMerchantExpandIndirectImageUpload()，图片上传接口。

## 版本号：Release 1.5.100

* 修改记录：
  * 支付宝：整理异步通知 NotifyRequest 结构体。
  * 微信V2：新增 client.ProfitSharingOrderAmountQuery()，查询订单待分账金额。
  * 微信V2：新增 client.ProfitSharingMerchantRatioQuery()，查询最大分账比例。
  * gopay：依赖库整理拆分整理，go mod 版本升级到 go1.20。
  * gopay：升级 1.5.100 版本需要伴随部分代码替换修改，谨慎升级。
  * gopay：优化一些已知问题。

## 版本号：Release 1.5.99

* 修改记录：
  * 微信V3：新增微信服务商模式支付分。
  * PayPal：支持代理BaseURL，国内部署以代理转发请求流量。

## 版本号：Release 1.5.98

* 修改记录：
  * 支付宝：PR新增 client.PostFileAliPayAPISelfV2()，文件上传自定义方法（未做验证，不知道是否好用）。
  * Apple：更新部分结构体字段。
  * PayPal：增加Token的自动刷新功能。

## 版本号：Release 1.5.97

* 修改记录：
  * 支付宝：新增 client.TradeOrderOnSettleQuery()，分账剩余金额查询。
  * 支付宝：新增 client.TradeRoyaltyRateQuery()，分账比例查询。
  * 支付宝：更新部分结构体，增加 open_id 相关字段。
  * 扫呗：新增扫呗支付方式。
  * gopay：xhttp库升级改造。

## 版本号：Release 1.5.96

* 修改记录：
  * 拉卡拉：新增拉卡拉支付。
  * interface{} -> any。

## 版本号：Release 1.5.95

* 修改记录：
  * 支付宝：新增 client.DoAliPay()。支付宝http请求方法。
  * Apple：苹果SDK client 化。新增多个 App Store Server API。client 模式为做严格测试（无测试信息key等信息）

## 版本号：Release 1.5.94

* 修改记录：
  * 支付宝：新增 client.TradeSettleConfirm()，统一收单确认结算接口。
  * 微信V3：新增 client.V3MediaDownloadImage()，图片下载
  * 微信V3：修正 查询投诉单协商历史 返回结构提结构

## 版本号：Release 1.5.93

* 修改记录：
  * 微信V2：服务商模式下，付款码支付response新增 sub_openid、sub_is_subscribe 字段。
  * 微信V3：新增 client.V3AsyncApply4SubModifySettlement()，(新)修改结算账户。
  * 微信V3：新增 client.V3Apply4SubMerchantsApplication()，查询结算账户修改申请状态。
  * 支付宝：沙箱模式网关URL替换升级成新版，沙箱调试时请跳转至新版沙箱应用。

## 版本号：Release 1.5.92

* 修改记录：
  * 微信V3：新增 client.V3ComplaintUpdateRefundProgress()，更新退款审批结果。
  * 微信V3：查询投诉单列表/详情，返回新增字段。
  * 微信V3：拆分model文件内结构体至不同文件内。
  * 微信V3：删除 client.V3TransferMerchantDetailQuery()，请更换 client.V3TransferMerchantDetail()
  * 支付宝：拆分model文件内结构体至不同文件内。

## 版本号：Release 1.5.91

* 修改记录：
  * 通联支付：新增通联支付对接，实现基本支付及其他接口。

## 版本号：Release 1.5.90

* 修改记录：
  * 微信V3：新增 client.V3BusinessPointsStatusQuery()，商圈会员待积分状态查询。
  * 微信V3：新增 client.V3BusinessParkingSync()，商圈会员停车状态同步。
  * 微信V3：新增 国密 支持。
  * 微信V3：新增 微信支付分停车服务 相关接口。
  * gopay：go mod 版本升级到 go1.18。

## 版本号：Release 1.5.89

* 修改记录：
  * 微信V2：新增 签约、扣款服务 相关接口。
  * 微信V3：新增 预扣款通知 接口。
  * gopay：debug 模式下，增加签名字符串的输出，更新 mod 包，优化部分方法。
  * alipay：新增 client.OpenAuthTokenAppInviteCreate()，ISV向商户发起应用授权邀约。
  * PayPal：新增 invoice 相关接口。

## 版本号：Release 1.5.88

* 修改记录：
  * PayPal：新增 invoice 相关接口。
  * gopay：更新 mod 包，优化部分方法。
  * alipay：修改过期时间字段类型为 int。

## 版本号：Release 1.5.87

* 修改记录：
  * 微信V3：新增小程序调起支付分所需要的支付参数方法：client.PaySignOfAppletScore()。
  * 微信V3：异步通知解析方法接收BodySize大小调至5MB。
  * 支付宝：统一收单交易退款接口，response补充接收字段。
  * 支付宝：alipay.SystemOauthToken() 接口增加可变参数 AppAuthToken 字段，如需要时可传。

## 版本号：Release 1.5.86

* 修改记录：
  * 微信V3：优化异步验签方法，v3NotifyReq.VerifySignByPKMap()，通过证书Map自动选择相应的证书验签。
  * 微信V3：微信平台证书获取方法变更，client.WxPublicKey() 获取最新微信平台证书；client.WxPublicKeyMap()，获取微信平台证书Map。
  * 支付宝：修复 app_auth_token 的问题。

## 版本号：Release 1.5.85

* 修改记录：
  * 微信V3：修复自动验签的bug

## 版本号：Release 1.5.84

* 修改记录：
  * 微信V3：优化微信平台证书的获取和选择
  * gopay：新增 retry 工具类包

## 版本号：Release 1.5.83

* 修改记录：
  * alipay：新增商家分账相关接口
  * alipay：更新部分注释上的接口文档地址。

## 版本号：Release 1.5.82

* 修改记录：
  * gopay：xhttp库增加BodySize的自定义化设置，各支付渠道 Client 增加 client.SetBodySize() 方法。
  * gopay：xlog 支持自定义实现接口。
  * apple：修改结构体字段 `revocationReason` 类型。

## 版本号：Release 1.5.81

* 修改记录：
  * 微信V3：签名方法，统一判断请求时path后缀是否为 ?，修复当bm为nil时的签名错误问题。
  * QQ：增加了异常结果处理。
  * 支付宝：补充了部分接口的回执参数。

## 版本号：Release 1.5.80

* 修改记录：
  * PayPal：增加payout相关的v1接口
  * PayPal：部分接口返回参数补齐

## 版本号：Release 1.5.79

* 修改记录：
  * PayPal：增加订阅相关的v1接口
  * 微信V3：更新部分接口上注释的文档地址
  * 支付宝：接口请求中，如下公共参数（version、return_url、notify_url、app_auth_token）支持 BodyMap 中此次请求自定义设置

## 版本号：Release 1.5.78

* 修改记录：
  * 微信V3：V3EcommerceBalance() 缺失参数补充
  * 微信V3：新增 client.V3EcommerceDayBalance() 方法，电商平台预约提现
  * 微信V3：修复 银行列表获取相关接口，路由修正
  * 微信V3：修复 client.V3BankSearchBank() 接口，私钥解密出错的问题

## 版本号：Release 1.5.77

* 修改记录：
  * 微信V3：V3EcommerceRefundQueryById()、V3EcommerceRefundQueryByNo()，缺失参数补充
  * 微信V3：新增 client.V3EcommerceWithdraw() 方法，电商平台预约提现
  * 微信V3：新增 client.V3EcommerceWithdrawStatus() 方法，电商平台查询预约提现状态

## 版本号：Release 1.5.76

* 修改记录：
  * gopay：大量优化error处理和返回，统一部分通用错误到 error.go 中
  * 支付宝：新增 alipay.IsBizError()，判断并捕获业务错误
  * 微信、支付宝：优化部分 error 返回格式以及透传，优化参数校验返回

## 版本号：Release 1.5.75

* 修改记录：
  * 微信V3：client.V3Apply4SubModifySettlement()，sub_mchid 问题处理
  * 微信V3：微信分账接收方，model参数补充添加 detail_id
  * PayPal：注释中接口文档地址更新
  * PayPal：新增 client.OrderConfirm()，订单确认
  * PayPal：OrderDetail、Capture、Payer、Name 等结构体，遗漏参数补充
  * gopay：pkg/xtime/parse_format.go，优化 DurationToUnit() 方法，int -> int64

## 版本号：Release 1.5.74

* 修改记录：
  * gopay：一些小改动，util.GetRandomString() -> util.RandomString()
  * gopay：升级 xlog

## 版本号：Release 1.5.73

* 修改记录：
  * Apple：新增内购支付通知V2解析

## 版本号：Release 1.5.72

* 修改记录：
  * Apple：返回参数类型错误修复，pending_renewal_info -> 数组类型
  * QQ：获取 AccessToken 结果 expires_in 类型修复，expires_in -> 字符串类型
  * 微信V3：证书相关代码优化

## 版本号：Release 1.5.71

* 修改记录：
  * 微信V2：去除所有微信小程序、公众号相关接口，请使用 wechat-sdk
  * 支付宝：client.UserCertdocCertverifyConsult() 方法，增加 authToken 参数
  * 微信V3：新增 银行组件（服务商） 相关接口，详情查看v3文档最下方的接口列表

## 版本号：Release 1.5.69

* 修改记录：
  * 微信V3：修改 client.V3RefundQuery()、增加入参参数，适配 服务商 模式
  * 微信V3：修复 client.V3Apply4SubSubmit()，接口路由修复
  * gopay：BodyMap 新增 Unmarshal() 方法，解析数据到结构体、数组指针

## 版本号：Release 1.5.68

* 修改记录：
  * 微信V3：修复 client.V3ComplaintResponse()、client.V3ComplaintComplete()， complaintId 参数类型错误问题
  * 微信V3：新增 电商收付通（分账）相关接口，详情查看v3文档最下方的接口列表
  * 微信V3：新增 电商收付通（补差）相关接口，详情查看v3文档最下方的接口列表
  * 微信V3：新增 电商收付通（退款）相关接口，详情查看v3文档最下方的接口列表
  * 微信V3：返回参数中字段，ID写法全部改写为Id写法

## 版本号：Release 1.5.67

* 修改记录：
  * 微信V3：配合微信文档修改，拆分服务商 批量转账 相关接口，接口如下：
  * 微信V3：新增 client.V3PartnerTransfer()
  * 微信V3：新增 client.V3PartnerTransferQuery()
  * 微信V3：新增 client.V3PartnerTransferDetail()
  * 微信V3：新增 client.V3PartnerTransferMerchantQuery()
  * 微信V3：新增 client.V3PartnerTransferMerchantDetail()
  * 微信V3：新增 client.V3Withdraw()
  * 微信V3：新增 client.V3WithdrawStatus()
  * 微信V3：新增 client.V3WithdrawDownloadErrBill()
    (10) 微信V3：修改 V3TransferDetailQuery() => V3TransferDetail()
    (11) 微信V3：修改 V3TransferMerchantDetailQuery() => V3TransferMerchantDetail()

## 版本号：Release 1.5.66

* 修改记录：
  * 微信V3：fix bug that `{"code":"PARAM_ERROR","message":"平台证书序列号Wechatpay-Serial错误"}`

## 版本号：Release 1.5.65

* 修改记录：
  * 微信V3：新增 client.V3EcommerceApply()，二级商户进件
  * 微信V3：新增 client.V3EcommerceApplyStatus()，查询申请状态
  * 微信V3：新增 client.V3GoldPlanManage()，点金计划管理
  * 微信V3：新增 client.V3GoldPlanBillManage()，商家小票管理
  * 微信V3：新增 client.V3GoldPlanFilterManage()，同业过滤标签管理
  * 微信V3：新增 client.V3GoldPlanOpenAdShow()，开通广告展示
  * 微信V3：新增 client.V3GoldPlanCloseAdShow()，关闭广告展示
  * 微信V3：公有化 wechat.GetReleaseSign()、wechat.GetSandBoxSign() 方法
  * 微信V3：修改 client.V3PartnerCloseOrder() 入参参数
    (10) GoPay：一些小修改优化

## 版本号：Release 1.5.64

* 修改记录：
  * xhttp：恢复 xhttp

## 版本号：Release 1.5.63

* 修改记录：
  * GoPay：部分代码优化
  * xhttp：xhttp client 优化，支持自定义client，默认还是使用标准 http.Client

## 版本号：Release 1.5.62

* 修改记录：
  * 微信V3：client 内 WxSerialNo、ApiV3Key 公有化
  * 微信V3：client 提供新方法 client.WxPublicKey() 直接获取 微信平台公钥
  * 微信V3：wechat 提供新方法 wechat.V3VerifySignByPK()，不再推荐使用 wechat.V3VerifySign()
  * 微信V3：V3NotifyReq 提供新方法 notify.VerifySignByPK()，不再推荐使用 notify.VerifySign()
  * 微信V3：整理微信v3说明文档

## 版本号：Release 1.5.61

* 修改记录：
  * gopay：更新 xhttp pkg, 方法全部增加 context 传递

## 版本号：Release 1.5.60

* 修改记录：
  * 微信V3：不再推荐使用 client.SetPlatformCert() 方法
  * 微信V3：新增 client.GetAndSelectNewestCert() 方法
  * 微信V3：重构 client.AutoVerifySign() 方法
  * QQ：新增 qq.GetAccessToken() 方法
  * QQ：新增 qq.GetOpenId() 方法
  * QQ：新增 qq.GetUserInfo() 方法

## 版本号：Release 1.5.59

* 修改记录：
  * 微信V3：证书获取方法返回结构体，去除 SignInfo 字段
  * gopay：BodyMap，EncodeURLParams 方法稍作调整
  * PayPal：PayPal支付能力接入（订单、支付）

## 版本号：Release 1.5.58

* 修改记录：
  * 微信V3：新增 client.V3FavorMediaUploadImage() 图片上传(营销专用)
  * 微信V3：新增 client.V3EcommerceIncomeRecord() 特约商户银行来账查询
  * 微信V3：新增 client.V3EcommerceBalance() 查询特约商户账户实时余额
  * 微信V3：新增 client.V3BusiFavorSend() 发放消费卡
  * 微信V3：新增 client.V3PartnershipsBuild() 建立合作关系
  * 微信V3：新增 client.V3PartnershipsTerminate() 终止合作关系
  * 微信V3：新增 client.V3PartnershipsList() 查询合作关系列表
  * 微信V3：修改 client.V3PartnerQueryOrder() 入参参数调整
  * 微信V3：修改 client.V3BillLevel2FundFlowBill() => client.V3BillEcommerceFundFlowBill() 申请特约商户资金账单
    (10) 支付宝：按照支付宝更新后的文档，修改大量接口返回参数结构体字段

## 版本号：Release 1.5.57

* 修改记录：
  * 微信V3：修复一些已知问题
  * 支付宝：一些细小的修复，部分参数类型更正

## 版本号：Release 1.5.56

* 修改记录：
  * 微信V3：修改 client.V3ProfitShareReturnResult() 接口入参，适配服务商模式
  * 微信V3：部分接口参数需要加密，修复 V3EncryptText() 和 V3DecryptText() 方法
  * 支付宝：修改 alipay.NewClient()，增加error返回值，去除Client内部分字段
  * Apple：新增apple pay的 apple.VerifyReceipt() 校验收据API
  * 优化代码中所有有关证书的解析操作

## 版本号：Release 1.5.55

* 修改记录：
  * 微信V3：wechat.NewClientV3()，去掉初始化参数 appid，所以方法中需要 appid 或sp_appid 的，需要自行传参
  * 微信V3：新增 代金券 相关接口
  * 微信V3：新增 商家券 相关接口
  * 微信V2、V3：修复部分接口发现的Bug

## 版本号：Release 1.5.54

* 修改记录：
  * 微信V3：新增微信支付分回调参数解密方法 notifyReq.DecryptScoreCipherText()
  * 微信V3：新增分账接口 client.V3ProfitShareMerchantConfigs()
  * Readme：更新Readme

## 版本号：Release 1.5.53

* 修改记录：
  * 支付宝：补充接口
  * 微信V3：修改支付分相关接口的返回参数字段， out_trade_no 为 out_order_no

## 版本号：Release 1.5.52

* 修改记录：
  * 支付宝：补充 支付API 相关接口
  * pkg：xhttp.Client 的 Transport 默认配置：Proxy: http.ProxyFromEnvironment

## 版本号：Release 1.5.51

* 修改记录：
  * 微信：新增 特约商户进件（服务商平台） 相关接口
  * 支付宝：补充完整 芝麻分 相关接口
  * 支付宝：补充 会员API 相关接口

## 版本号：Release 1.5.50

* 修改记录：
  * 支付宝：新增 芝麻分 相关接口
  * 支付宝：当判断 Response 中 code!="10000" 时，不再返回nil，而是返回 aliRsp 结果

## 版本号：Release 1.5.49

* 修改记录：
  * 微信V3：新增 wechat.GetPlatformCerts()，无需初始化V3client，直接获取微信平台证书和序列号等信息
  * gopay：更新 go mod version
  * 微信V2：新增 client.CustomsDeclareOrder()，订单附加信息提交（正式环境）
  * 微信V2：新增 client.CustomsDeclareQuery()，订单附加信息查询（正式环境）
  * 微信V2：新增 client.CustomsReDeclareOrder()，订单附加信息重推（正式环境）
  * 支付宝：新增 client.TradeCustomsDeclare()，统一收单报关接口（正式环境）
  * 支付宝：新增 client.AcquireCustoms()，报关接口（正式环境），未经测试
  * 支付宝：新增 client.AcquireCustomsQuery()，报关查询接口（正式环境），未经测试

## 版本号：Release 1.5.48

* 修改记录：
  * 微信V3：修复 平台证书序列号Wechatpay-Serial错误 问题
  * 微信V3：新增 client.SetPlatformCert()，设置 微信支付平台证书 和 证书序列号 方法
  * 微信V3：新增 client.V3EncryptText()，请求参数 敏感信息 加密方法
  * 微信V3：新增 client.V3DecryptText()，返回参数 敏感信息 解密方法
  * 微信V3：修改 client.AutoVerifySign() 方法无需传参，但需要提前调用 client.SetPlatformCert() 设置 微信支付平台证书 和
    证书序列号

## 版本号：Release 1.5.47

* 修改记录：
  * 微信V3：新增 转账相关 相关接口
  * 微信V3：新增 账户余额查询 相关接口
  * 微信V3：新增 来账识别 相关接口

## 版本号：Release 1.5.46

* 修改记录：
  * 微信V3：新增敏感信息加解密方法，wechat.V3EncryptText() 加密数据，wechat.V3DecryptText() 解密数据
  * 微信V3：新增 微信先享卡 相关接口
  * 微信V3：新增 支付即服务 相关接口
  * 微信V3：新增 智慧商圈 相关接口

## 版本号：Release 1.5.45

* 修改记录：
  * 支付宝：优化现有代码，修复公钥证书模式下，同步验签失败的问题
  * 支付宝：新增 client.AutoVerifySign()，自动同步验签设置（公钥证书模式）
  * 支付宝：新增 client.PublicCertDownload()，应用支付宝公钥证书下载
  * 支付宝：新增 client.FundTransPayeeBindQuery()，资金收款账号绑定关系查询
  * 支付宝：新增 client.OpenAppQrcodeCreate()，小程序生成推广二维码接口
  * 支付宝：新增 client.UserAgreementPageSign()，支付宝个人协议页面签约接口
  * 支付宝：新增 client.UserAgreementPageUnSign()，支付宝个人代扣协议解约接口
  * 支付宝：新增 client.UserAgreementQuery()，支付宝个人代扣协议查询接口
  * 支付宝：新增 client.TradeOrderInfoSync()，支付宝订单信息同步接口
    (10) 支付宝：新增 client.TradeAdvanceConsult()，订单咨询服务

## 版本号：Release 1.5.44

* 修改记录：
  * 微信V3：新增 图片上传 接口
  * 微信V3：新增 视频上传 接口
  * 微信V3：修复消费者投诉接口中的图片上传失败问题

## 版本号：Release 1.5.42

* 修改记录：
  * 迁移新仓库 https://github.com/go-pay/gopay

## 版本号：Release 1.5.41

* 修改记录：
  * 微信V3：新增 消费者投诉2.0 相关接口
  * 微信V3：新增 分账 相关接口

## 版本号：Release 1.5.40

* 修改记录：
  * 微信V3：新增微信支付分（免确认模式）相关接口
  * 微信V3：新增微信支付分（免确认预授权模式）相关接口

## 版本号：Release 1.5.39

* 修改记录：
  * 微信V3：新增微信支付分（公共API）相关接口

## 版本号：Release 1.5.38

* 修改记录：
  * 微信：去掉所以微信 client 方法中需要传证书的参数，请统一在初始化client时，添加证书
  * 微信：使用方法请参考 wechat/client_test.go 下的初始化，以及各个方法使用

## 版本号：Release 1.5.37

* 修改记录：
  * 支付宝：修改 client.FundAuthOrderAppFreeze() 接口返回参数
  * 支付宝：新增 client.GetRequestSignParam()，获取已签名的完整请求参数
  * 微信V3：增加 client.GetPlatformCerts()，获取微信平台证书公钥，增加注释说明
  * 支付宝：拆分 _test.go 文件

## 版本号：Release 1.5.36

* 修改记录：
  * 支付宝：新增 资金API 类别 接口实现
  * 微信V3：修复银行转账接口，银行卡号和收款人的 RSA 加密bug
  * 一些其他小修复调整

## 版本号：Release 1.5.35

* 修改记录：
  * 微信V3：修复 普通支付回调通知 解密后的结构体问题
  * 微信V3：新增 合单支付回调通知
  * 微信V3：修复 退款回调通知 解密后的结构体问题

## 版本号：Release 1.5.34

* 修改记录：
  * 微信V3：修复 client.GetPlatformCerts() 的返回值 code 问题
  * ReadMe：补充 README.md 说明
  * 微信V3：修复 PaySignOfApp() 签名出错的问题
  * 微信V2：部分文件调整

## 版本号：Release 1.5.32

* 修改记录：
  * xhttp：Fix bug about Transport

## 版本号：Release 1.5.31

* 修改记录：
  * 微信：新增服务商支付接口

## 版本号：Release 1.5.30

* 修改记录：
  * BodyMap：恢复 bm.Get() 方法获取的是string类型，增加 bm.GetInterface()
  * 微信：新增V3版本退款查询接口

## 版本号：Release 1.5.29

发布时间：2021/02/27 22:49

* 修改记录：
  * 支付宝：新增 client.PostAliPayAPISelfV2()，比非V2版本更灵活化，具体参考 client_test.go 内的
    TestClient_PostAliPayAPISelfV2() 方法
  * BodyMap：新增 bm.SetFormFile() 的部分方法，修改 bm.Get() 方法，新增bm.GetString() 方法
  * xHttp：更新 httpClient， httpClient.Type() 支持 TypeMultipartFormData 类型
  * go mod 版本改为 1.14

## 版本号：Release 1.5.28

发布时间：2021/02/19 18:48

* 修改记录：
  * QQ：新增 client.AddCertFileContent()，解决无证书文件，只有证书内容的问题
  * 支付宝：新增 alipay.VerifySyncSignWithCert()，同步证书验签
  * 支付宝：新增 client.SetCertSnByContent()，通过应用公钥证书内容设置 app_cert_sn、alipay_root_cert_sn、alipay_cert_sn
  * 支付宝：删除废弃接口 client.FundTransToaccountTransfer()
  * fix BodyMap 的部分方法

## 版本号：Release 1.5.27

发布时间：2021/02/03 18:50

* 修改记录：
  * GoPay：去掉对 gotil 的强依赖

## 版本号：Release 1.5.26

发布时间：2021/01/29 19:38

* 修改记录：
  * 微信：重新整理文件分级，商户分账模块增加test方法说明
  * BodyMap: 去除 GetArrayBodyMap()、GetBodyMap() 方法

## 版本号：Release 1.5.25

发布时间：2020/12/31 18:38

* 修改记录：
  * 微信：v3 基础支付接口完成，使用请参考：gopay/wechat/v3/client_test.go

## 版本号：Release 1.5.24

发布时间：2020/12/21 18:58

* 修改记录：
  * 微信：证书支持二选一，只传 apiclient_cert.pem 和 apiclient_key.pem 或者只传 apiclient_cert.p12

## 版本号：Release 1.5.23

发布时间：2020/12/15 17:58

* 修改记录：
  * 微信：新增 client.AddCertFileContent()，解决无证书文件，只有证书内容的问题

## 版本号：Release 1.5.22

发布时间：2020/12/04 02:58

* 修改记录：
  * 更新 gotil，修复xlog导致 go 1.14 以下版本报bug问题
  * 采纳 WenyXu 的意见，优化BodyMap的用法

## 版本号：Release 1.5.20

发布时间：2020/09/30 23:58

* 修改记录：
  * 微信：client 添加 DebugSwitch 开关，默认关闭，不输出 请求参数和返回参数，通过 client.DebugSwitch = gopay.DebugOn 打开
  * 支付宝：client 添加 DebugSwitch 开关，默认关闭，不输出 请求参数和返回参数，通过 client.DebugSwitch = gopay.DebugOn 打开
  * QQ：client 添加 DebugSwitch 开关，默认关闭，不输出 请求参数和返回参数，通过 client.DebugSwitch = gopay.DebugOn 打开
  * 更新 Gotil

## 版本号：Release 1.5.19

发布时间：2020/09/20 23:58

* 修改记录：
  * 微信：修复 client.ProfitSharingQuery() 接口的Bug，https://github.com/iGoogle-ink/gopay/issues/68
  * 微信：优化 client.doProdPost()
  * 支付宝：优化 client.doAliPay()
  * 微信：项目文件区分改动，开放平台接口和微信公众号区分
  * 微信：替换 wechat.GetAppLoginAccessToken() = > wechat.GetOauth2AccessToken()
  * 微信：替换 wechat.RefreshAppLoginAccessToken() = > wechat.RefreshOauth2AccessToken()
  * 微信：替换 wechat.GetUserInfoOpen() = > wechat.GetOauth2UserInfo()
  * 微信：替换 wechat.GetUserInfo() = > wechat.GetPublicUserInfo()
  * 微信：新增 wechat.CheckOauth2AccessToken() 检验授权凭证（access_token）是否有效
    (10) 微信：新增 wechat.GetPublicUserInfoBatch() 批量获取用户基本信息（微信公众号）
    (11) 微信：新增 client.SendCashRed() 发放现金红包
    (12) 微信：新增 client.SendGroupCashRed() 发放现金裂变红包
    (13) 微信：新增 client.SendAppletRed() 发放小程序红包
    (14) 微信：新增 client.QueryRedRecord() 查询红包记录
    (15) QQ：新增 client.SendCashRed() 创建现金红包，（未经测试，有条件的帮忙测一下吧，有问题提PR）
    (16) QQ：新增 client.DownloadRedListFile() 对账单下载，（未经测试，有条件的帮忙测一下吧，有问题提PR）
    (17) QQ：新增 client.QueryRedInfo() 查询红包详情，（未经测试，有条件的帮忙测一下吧，有问题提PR）

## 版本号：Release 1.5.18

发布时间：2020/08/29 18:30

* 修改记录：
  * 微信：修复 client.AddCertFilePath() 无效的Bug
  * QQ：修复 client.AddCertFilePath() 无效的Bug
  * Gotil：升级 gotil 到 v1.0.7-beta2 版本
  * 支付宝：OpenAuthTokenAppResponse 结构体中 ExpiresIn、ReExpiresIn
    字段改为int64（有用户反馈返回的是int类型，但文档写的是string），如果此处有问题，请立马联系改回去。

## 版本号：Release 1.5.17

发布时间：2020/08/23 15:30

* 修改记录：
  * 微信：Response model 增加字段
  * ReadMe：修改部分遗留未更改的文档内容
  * 支付宝：添加证书由只支持证书路径，改为支持证书路径或者这书Byte数组
  * 支付宝：修复SystemOauthToken()方法未添加 AppCertSN 和 AliPayRootCertSN 的问题

## 版本号：Release 1.5.16

发布时间：2020/07/29 18:30

* 修改记录：
  * 微信：新增公共方法：wechat.GetUserInfoOpen()，微信开放平台：获取用户个人信息(UnionID机制)
  * Gotil：升级 gotil 到 v1.0.4 版本
  * 微信：新增ReadMe说明，微信支付下单等操作可用沙箱环境测试是否成功，但真正支付时，请使用正式环境，isProd = true，不然会报错

## 版本号：Release 1.5.15

发布时间：2020/07/09 18:30

* 修改记录：
  * 微信：新增client方法：client.ProfitSharing()，请求单次分账
  * 微信：新增client方法：client.MultiProfitSharing()，请求多次分账
  * 微信：新增client方法：client.ProfitSharingQuery()，查询分账结果
  * 微信：新增client方法：client.ProfitSharingAddReceiver()，添加分账接收方
  * 微信：新增client方法：client.ProfitSharingRemoveReceiver()，删除分账接收方
  * 微信：新增client方法：client.ProfitSharingFinish()，完结分账
  * 微信：新增client方法：client.ProfitSharingReturn()，分账回退
  * 微信：新增client方法：client.ProfitSharingReturnQuery()，分账回退结果查询
  * 微信：新增client方法：client.PayBank()，企业付款到银行卡API
    (10) 微信：新增client方法：client.QueryBank()，查询企业付款到银行卡API
    (11) 微信：新增client方法：client.GetRSAPublicKey()，获取RSA加密公钥API
    (12) 微信：修改client方法名：client.PostRequest() -> client.PostWeChatAPISelf()
    (13) QQ：修改client方法名：client.PostRequest() -> client.PostQQAPISelf()
    (14) 说明：方法未经严格测试，还请开发者在开始使用时确认是否正常使用，有问题请提 issue

## 版本号：Release 1.5.14

发布时间：2020/06/27 3:30

* 修改记录：
  * 引入 github.com/iGoogle-ink/gotil 包
  * 替换 log 输出样式
  * 支付宝：新增client方法：client.PostAliPayAPISelf()，支付宝接口自行实现方法

## 版本号：Release 1.5.12

发布时间：2020/05/20 02:10

* 修改记录：
  * http_client：增加默认请求的超时时间 60s，增加 SetTimeout() 方法，可自定义超时时间
  * 微信：修改 申请退款、退款查询、订单查询 接口的返回结构体，增加带下标的部分字段
  * 微信：增加 申请退款、退款查询、订单查询 接口的返回值，新增了 resBm BodyMap 类型，方便接收结构体中未定义到的下标字段
  * 微信：新增client方法：client.GetTransferInfo()，查询企业退款，此方法实际暂未测试，请自行测试，有问题提issue

## 版本号：Release 1.5.11

发布时间：2020/05/13 17:15

* 修改记录：
  * 支付宝：修复rsp解析出错的问题 client.SystemOauthToken()
  * 微信：修改部分公共方法Rsp结构体参数问题，同步微信文档

## 版本号：Release 1.5.10

发布时间：2020/05/06 20:15

* 修改记录：
  * 微信：修改部分公共方法返回值结构体字段类型
  * drone fix

## 版本号：Release 1.5.9

发布时间：2020/04/25 15:32

* 修改记录：
  * 支付宝：异步验签，推荐使用 alipay.ParseNotifyToBodyMap()，解析参数后参数在Verify验签时，推荐传入参数BodyMap bm。
  * 支付宝：修改公共方法：alipay.ParseNotifyResultToBodyMap() 为 alipay.ParseNotifyToBodyMap()
  * 支付宝：修改公共方法：alipay.ParseNotifyResultByURLValues() 为 alipay.ParseNotifyByURLValues()
  * 支付宝：废弃公共方法：alipay.ParseNotifyResult()，因为异步通知有参数因为支付接口不同，返回的字段不同，无法使用结构体全部定义好
  * 支付宝：调整了部分接口的文档地址
  * 微信：修改公共方法：wechat.ParseNotifyResultToBodyMap() 为 wechat.ParseNotifyToBodyMap()
  * 微信：修改公共方法：wechat.ParseNotifyResult() 为 wechat.ParseNotify()
  * 微信：修改公告方法：wechat.ParseRefundNotifyResult() 为 wechat.ParseRefundNotify()

## 版本号：Release 1.5.8

发布时间：2020/04/18 21:32

* 修改记录：
  * 微信：新增Client方法：client.PostRequest()，向微信发送Post请求，对于本库未提供的微信API，可自行实现，通过此方法发送请求
  * 微信：微信同步返回结构体类型全部修改为string类型，验签出错的问题
  * 微信：Client方法，需要传证书的接口方法，入参类型统一改为any,无需传证书地址时，由 "" 改为 nil
  * QQ：同微信改动
  * 支付宝：model结构体参数全部修改为string类型

## 版本号：Release 1.5.7

发布时间：2020/03/25 20:32

* 修改记录：
  * 支付宝：修改 client.UserCertifyOpenQuery() 方法的返回值解析类型报错问题，官方文档类型实例有误

## 版本号：Release 1.5.6

发布时间：2020/03/06 17:32

* 修改记录：
  * 支付宝：新增Client方法：client.SetPrivateKeyType()，设置 支付宝 私钥类型，alipay.PKCS1 或 alipay.PKCS8，默认 PKCS1。
  * 支付宝：修改公共方法：alipay.GetRsaSign()，增加了私钥类型参数，并将私钥的格式化操作，移动到该方法内，传入的私钥无需事先格式化。

## 版本号：Release 1.5.5

发布时间：2020/03/05 18:32

* 修改记录：
  * 支付宝：新增Client方法：client.DataBillBalanceQuery()，支付宝商家账户当前余额查询。
  * 支付宝：新增Client方法：client.DataBillDownloadUrlQuery()，查询对账单下载地址。
  * 支付宝：开放公共方法：alipay.GetRsaSign()，获取支付宝参数签名（参数sign值）。
  * 支付宝：开放公共方法：alipay.FormatURLParam()，格式化支付宝请求URL参数。
  * 支付宝：新增公共方法：alipay.ParseNotifyResultByURLValues()，通过 url.Values 解析支付宝支付异步通知的参数到Struct。

## 版本号：Release 1.5.4

发布时间：2020/02/29 14:32

* 修改记录：
  * 支付宝：新增Client方法：client.UserInfoAuth()，用户登陆授权。（方法未测试通过，待有测试条件的同学测试一下吧）
  * 支付宝：新增公共方法：alipay.MonitorHeartbeatSyn()，验签接口。（方法未测试通过，待有测试条件的同学测试一下吧）

## 版本号：Release 1.5.3

发布时间：2020/02/19 11:32

* 修改记录：
  * 支付宝：修改公共方法：SystemOauthToken()，添加参数 signType

## 版本号：Release 1.5.2

发布时间：2020/02/14 13:32

* 修改记录：
  * 支付宝：官方单笔转账接口更新，新增Client方法：client.FundTransUniTransfer()，单笔转账接口
  * 支付宝：新增Client方法：client.FundTransCommonQuery()，转账业务单据查询接口
  * 支付宝：新增Client放大：client.FundAccountQuery()，支付宝资金账户资产查询接口
  * 支付宝：Client的方法，必选参数校验

## 版本号：Release 1.5.1

发布时间：2020/01/03 17:32

* 修改记录：
  * 由于下载包需要 /v2 的问题，替换版本号到 1.x，代码不变，只改变版本号记录

## 版本号：Release 2.0.5

发布时间：2020/01/01 22:55

* 修改记录：
  * 添加一些函数参数判空操作，避免Panic
  * 去掉不用的结构体 ReturnMessage
  * 去掉 go mod v1.4.8版本的依赖

## 版本号：Release 2.0.4

发布时间：2019/12/24 14:29

* 修改记录：
  * 支付宝：新增支付宝公钥文件验证签方法（公钥证书模式）：client.VerifySignWithCert()

## 版本号：Release 2.0.3

发布时间：2019/12/18 19:25

* 修改记录：
  * 微信：新增Client方法：client.AuthCodeToOpenId()，授权码查询OpenId（正式）
  * 微信：新增Client方法：client.Report()，交易保障
  * 微信：新增Client方法：client.EntrustPublic()，公众号纯签约（正式）
  * 微信：新增Client方法：client.EntrustAppPre()，APP纯签约-预签约接口-获取预签约ID（正式）
  * 微信：新增Client方法：client.EntrustH5()，H5纯签约（正式）
  * 微信：新增Client方法：client.EntrustPaying()，支付中签约（正式）

## 版本号：Release 2.0.2

发布时间：2019/12/13 23:25

* 修改记录：
  * 微信：删除Client方法：client.AddCertFileByte()
  * 版本限制 golang 1.13，fmt.Errorf() 使用 %w 格式化 error，1.13新特性

## 版本号：Release 2.0.1

发布时间：2019/12/13 17:20

* 修改记录：
  * 处理 go mod 包，go get github.com/iGoogle-ink/gopay/v2

## 版本号：Release 2.0.0

发布时间：2019/12/12 18:22

* 修改记录：
  * 按支付渠道模块分包大调整
  * 一大推细小改动，不一一描述了
  * 支付宝：修改公共API方法：alipay.GetCertSN()，不再支持支付宝根证书的SN获取
  * 支付宝：新增公共API方法：alipay.GetRootCertSN()，获取root证书序列号SN
  * 支付宝：新增Client方法：alipay.SetLocation()，设置时区，不设置或出错均为默认服务器时间

## 版本号：Release 1.4.8

发布时间：2019/12/11 16:40

* 修改记录：
  * 1.几 最后一个版本
  * 修复一些问题
  * 支付宝：修改公共API方法：gopay.GetCertSN()，不再支持支付宝根证书的SN获取
  * 支付宝：新增公共API方法：gopay.GetRootCertSN()，获取root证书序列号SN
  * 微信：修改公共API方法：gopay.GetWeChatSanBoxParamSign()，修复 沙箱验签出错问题

## 版本号：Release 1.4.6

发布时间：2019/12/09 18:37

* 修改记录：
  * 移除第三方http请求库，自己封装了一个请求库使用，解决不会设置 goproxy 无法下载包的问题
  * 使用中如有问题，请微信联系作者处理 或 提issue

## 版本号：Release 1.4.5

发布时间：2019/12/07 21:56

* 修改记录：
  * 支付宝：修复 公钥证书模式 下，同步返回参数接收问题，返回接收结构体增加参数 alipay_cert_sn，同步返回证书模式验签，暂时未解决
  * 支付宝：修改公共API方法：gopay.VerifyAliPaySign()，不再支持同步验签，只做异步通知验签
  * 支付宝：新增公共API方法：gopay.VerifyAliPaySyncSign()，支付宝同步返回验签
  * 支付宝：新增Client方法：client.SetAliPayPublicCertSN()，设置 支付宝公钥证书SN，通过 gopay.GetCertSN() 获取
    alipay_cert_sn
  * 支付宝：新增Client方法：client.SetAppCertSnByPath()，设置 app_cert_sn 通过应用公钥证书路径
  * 支付宝：新增Client方法：client.SetAliPayPublicCertSnByPath()，设置 alipay_cert_sn 通过 支付宝公钥证书文件路径
  * 支付宝：新增Client方法：client.SetAliPayRootCertSnByPath()，设置 alipay_root_cert_sn 通过支付宝CA根证书文件路径

## 版本号：Release 1.4.4

发布时间：2019/11/16 15:56

* 修改记录：
  * 支付宝：新增公共API方法：gopay.ParseAliPayNotifyResultToBodyMap()，解析支付宝支付异步通知的参数到BodyMap
  * 支付宝：修改公共API方法：gopay.VerifyAliPaySign()，支付宝异步验签支持传入 BodyMap
  * 微信：新增Client方法：client.AddCertFileByte()，添加微信证书 Byte 数组
  * 微信：新增Client方法：client.AddCertFilePath()，添加微信证书 Path 路径
  * 微信：微信Client需要证书的方法，如已使用client.AddCertFilePath()或client.AddCertFileByte()
    添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，否则，3证书Path均不可空
  * BodyMap 的Set方法去掉switch判断，直接赋值
  * WeChatClient、AliPayClient 加锁
  * 修改部分小问题和部分样式

## 版本号：Release 1.4.3

发布时间：2019/11/12 01:15

* 修改记录：
  * 微信：新增公共API方法：gopay.ParseWeChatRefundNotifyResult()，解析微信退款异步通知的请求参数
  * 微信：新增公共API方法：gopay.DecryptRefundNotifyReqInfo()，解密微信退款异步通知的加密数据
  * 支付宝：修改Client方法：client.AliPayUserCertifyOpenCertify()，身份认证开始认证（获取认证链接）
  * 修改部分小问题

## 版本号：Release 1.4.2

发布时间：2019/11/11 16:43

* 修改记录：
  * 支付宝：新增Client方法：client.AliPayUserCertifyOpenInit()，身份认证初始化服务
  * 支付宝：新增Client方法：client.AliPayUserCertifyOpenCertify()，身份认证开始认证
  * 支付宝：新增Client方法：client.AliPayUserCertifyOpenQuery()，身份认证记录查询
  * 支付宝：所有方法的返回结构体下的 XxxResponse 字段，统一修改为 Response，如：
    type AliPayTradeCreateResponse struct {
    Response createResponse `json:"alipay_trade_create_response,omitempty"`
    SignData string         `json:"-"`
    Sign string         `json:"sign"`
    }
  * 支付宝：修改一些代码格式问题
  * 支付宝：client.AlipayOpenAuthTokenApp() 修改为 client.AliPayOpenAuthTokenApp()

## 版本号：Release 1.4.1

发布时间：2019/11/04 14:28

* 修改记录：
  * 支付宝：修改公共API方法：GetCertSN()，修复获取支付宝根证书获取 sn 的问题（wziww）
  * 支付宝：新增Client方法：client.SetAppCertSN()，可自行获取公钥 sn 并赋值
  * 支付宝：新增Client方法：client.SetAliPayRootCertSN()，可自行获取根证书 sn 并赋值

## 版本号：Release 1.4.0

发布时间：2019/10/10 13:51

* 修改记录：
  * AliPayNotifyRequest 结构体，新增加两个字段：method、timestamp，修复电脑网站支付，配置 return_url 支付成功后，支付宝请求该
    return_url 返回参数验签失败的问题
  * 去除支付宝老验签方法 VerifyAliPayResultSign()
  * 去除微信老验签方法 VerifyWeChatResultSign()

## 版本号：Release 1.3.9

发布时间：2019/09/30 00:01

* 修改记录：
  * 修复支付宝支付验签出错的问题！

## 版本号：Release 1.3.8

发布时间：2019/09/24 17:51

* 修改记录：
  * 代码风格修改更新

## 版本号：Release 1.3.7

发布时间：2019/09/22 11:41

* 修改记录：
  * README 增加 go mod 安装gopay的方法指导

## 版本号：Release 1.3.6

发布时间：2019/09/09 23:51

* 修改记录：
  * 新增支付宝Client方法：client.AlipayUserInfoShare() => 支付宝会员授权信息查询接口（App支付宝登录）

## 版本号：Release 1.3.6

发布时间：2019/09/05 02:55

* 修改记录：
  * 更改微信公共API方法名称：gopay.GetAccessToken() to gopay.GetWeChatAppletAccessToken() => 获取微信小程序全局唯一后台接口调用凭据
  * 更改微信公共API方法名称：gopay.GetPaidUnionId() to gopay.GetWeChatAppletPaidUnionId() => 微信小程序用户支付完成后，获取该用户的
    UnionId，无需用户授权
  * 新增微信公共API方法：gopay.GetAppWeChatLoginAccessToken() => App应用微信第三方登录，code换取access_token
  * 新增微信公共API方法：gopay.RefreshAppWeChatLoginAccessToken() => 刷新App应用微信第三方登录后，获取的 access_token

## 版本号：Release 1.3.5

发布时间：2019/09/05 02:10

* 修改记录：
  * 支付宝、微信Client 由私有改为公有

## 版本号：Release 1.3.4

发布时间：2019/09/03 19:26

* 修改记录：
  * 新增支付宝公共API方法：gopay.GetCertSN() => 获取证书SN号（app_cert_sn、alipay_root_cert_sn、alipay_cert_sn）
  * 新增支付宝Client方法：client.SetAliPayRootCertSN() => 设置支付宝根证书SN，通过 gopay.GetCertSN() 获取
  * 新增支付宝Client方法：client.SetAppCertSN() => 设置应用公钥证书SN，通过 gopay.GetCertSN() 获取
