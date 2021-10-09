### 支付宝支付 API

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
    * 花芝轻会员结算申请: `client.PcreditHuabeiAuthSettleApply()`
    * NFC用户卡信息同步: `client.CommerceTransportNfccardSend()`
    * 广告投放数据查询: `client.DataDataserviceAdDataQuery()`
    * 航司电话订票待申请接口: `client.CommerceAirCallcenterTradeApply()`
    * 网商银行全渠道收单业务订单创建: `client.PaymentTradeOrderCreate()`
    * 聚合支付订单咨询服务: `client.KoubeiTradeOrderAggregateConsult()`
    * 口碑订单预下单: `client.KoubeiTradeOrderPrecreate()`
    * 口碑商品交易购买接口: `client.KoubeiTradeItemorderBuy()`
    * 口碑订单预咨询: `client.KoubeiTradeOrderConsult()`
    * 口碑商品交易退货接口: `client.KoubeiTradeItemorderRefund()`
    * 口碑商品交易查询接口: `client.KoubeiTradeItemorderQuery()`
    * 码商发码成功回调接口: `client.KoubeiTradeTicketTicketcodeSend()`
    * 口碑凭证延期接口: `client.KoubeiTradeTicketTicketcodeDelay()`
    * 口碑凭证码查询: `client.KoubeiTradeTicketTicketcodeQuery()`
    * 口碑凭证码撤销核销: `client.KoubeiTradeTicketTicketcodeCancel()`
    * 修改蚂蚁店铺: `client.AntMerchantShopModify()`
    * 蚂蚁店铺创建: `client.AntMerchantShopCreate()`
    * 蚂蚁店铺创建咨询: `client.AntMerchantShopConsult()`
    * 商户申请单查询: `client.AntMerchantOrderQuery()`
    * 店铺查询接口: `client.AntMerchantShopQuery()`
    * 蚂蚁店铺关闭: `client.AntMerchantShopClose()`
    * 申请权益发放: `client.CommerceBenefitApply()`
    * 权益核销: `client.CommerceBenefitVerify()`
    * 还款账单查询: `client.TradeRepaybillQuery()`
* 网页&移动应用 - <font color='#027AFF' size='4'>资金API</font>
    * 单笔转账接口：`client.FundTransUniTransfer()`
    * 查询转账订单接口: `client.FundTransOrderQuery()`
    * 支付宝资金账户资产查询接口：`client.FundAccountQuery()`
    * 转账业务单据查询接口：`client.FundTransCommonQuery()`
    * 资金退回接口: `client.FundTransRefund()`
    * 资金授权冻结接口: `client.FundAuthOrderFreeze()`
    * 资金授权发码接口: `client.FundAuthOrderVoucherCreate()`
    * 线上资金授权冻结接口: `client:FundAuthOrderAppFreeze()`
    * 资金授权解冻接口: `client.FundAuthOrderUnfreeze()`
    * 资金授权操作查询接口: `client.FundAuthOperationDetailQuery()`
    * 资金授权撤销接口: `client.FundAuthOperationCancel()`
    * 批次下单接口: `client.FundBatchCreate()`
    * 批量转账关单接口: `client.FundBatchClose()`
    * 批量转账明细查询接口: `client.FundBatchDetailQuery()`
    * 现金红包无线支付接口: `client.FundTransAppPay()`
    * 资金收款账号绑定关系查询: `client.FundTransPayeeBindQuery()`
    * 资金转账页面支付接口: `client.FundTransPagePay()`
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

### 支付宝公共 API

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