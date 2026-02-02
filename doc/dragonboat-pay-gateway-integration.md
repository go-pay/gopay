# Dragonboat × GoPay：支付网关化对接方案（Design）

本文把 `gopay`（Go SDK）网关化为一个独立 Go 服务（Data Plane），并在 `dragonboat-backend` 新增 `ruoyi-pay` 模块作为支付编排与状态机（Control Plane），实现“Java 管理端调用支付能力 + 回调驱动”的闭环。

## Phase 1 范围（验收口径）

- 渠道：**仅 ALIPAY 作为生产级验收**；`WECHAT_V3` 只做代码级兼容/沙箱演示。
- ALIPAY Scene：`PRECREATE`（二维码）+ `WAP`（H5 跳转），统一 `payData` 形态（`qrCode` / `payUrl`）。
- WECHAT_V3 演示建议：`scene=NATIVE`（二维码），复用 `payData.qrCode` 展示二维码即可。
- 币种：网关侧强制 `CNY`；汇率换算与舍入（Round Half Up）由 Java 侧定义并落金额快照。
- 商户：`merchantId` 本期只有一个，但 API/表结构必须保留该字段，避免未来扩展改签名。

## Phase 1 落地配置（建议写入 Nacos）

建议在 `ruoyi-pay.yml` 配置并由 Java 与 Go 共用同一份 secret：
- `pay.gateway.sharedSecret`
- `pay.gateway.sharedSecretPrev`（可选，用于轮换窗口）
- `pay.gateway.clockSkewSeconds=300`
- `pay.gateway.nonceTtlSeconds=300`
- `pay.gateway.defaultMerchantId`（本期单商户时使用）

ruoyi-pay 建议提供内网接口（不经 `ruoyi-gateway` 暴露），并做 HMAC 校验：
- `POST /internal/pay-gateway/events`（pay-gateway → ruoyi-pay）
- `GET /internal/pay-gateway/merchants/snapshot`（pay-gateway 拉取商户快照）

## 1. 总体架构（推荐）

**Go：pay-gateway（新增服务）**
- 对外（公网）：仅接收支付平台回调（微信/支付宝/PayPal…），完成解析、验签/解密、去重，并把“可信事件”投递给 Java（Webhook）。
- 对内（内网）：提供统一支付 API（下单/关单/查询/退款/补偿查询），供 Java 调用。
- 设计目标：**无状态**（便于水平扩展），依赖 Redis 做幂等/去重/Outbox（可选），事件投递链路默认走 HTTP（可按需演进 MQ）。

**Java：`ruoyi-modules/ruoyi-pay`（新增微服务模块）**
- 对外：管理端/业务端 REST（通过 `ruoyi-gateway` 路由 `/pay/**`）。
- 对内：可选 Dubbo API（建议新增 `ruoyi-api/ruoyi-api-payment` 承载契约与 DTO）。
- 职责：支付订单/退款订单的**业务状态机**、补偿任务、商户配置管理（多租户）、权限与审计；并对 pay-gateway 提供“商户配置快照”接口供其拉取（动态加载配置）。

**Redis（建议）**
- Go 幂等/回调去重/Outbox 重试（多实例必配）。
- 可选：用于推送“配置变更通知”（stream），加速 pay-gateway 的配置刷新（仍以轮询兜底）。

**RabbitMQ（可选演进）**
- 若你们更偏好 MQ，也可以把“Go → Java”从 Webhook 演进为 MQ 投递（pay-gateway 需要新增 publisher）。

## 2. L0 能力边界（必须统一抽象）

Java 侧统一抽象 + `ext` 透传，覆盖四大金刚：
- 创建支付（下单）：返回前端拉起所需 `payData`
- 查询支付 / 关单
- 发起退款 / 查询退款
- 回调处理：**仅 Go 接收**（Java 不直面外部回调）

> 金额字段：跨服务传输统一 `Long`（最小货币单位，如分）+ `currency`（ISO 4217），严禁 `float/double`。

## 3. Go pay-gateway：对内统一 API（HTTP/JSON）

建议接口（v1）：
- `POST /v1/payments` 创建支付（幂等）
- `GET /v1/payments/{outTradeNo}` 查询支付
- `POST /v1/payments/{outTradeNo}/close` 关单
- `POST /v1/refunds` 发起退款（幂等）
- `GET /v1/refunds/{outRefundNo}` 查询退款
- `POST /v1/compensations/payments/query` 批量补偿查询（给 Java 定时任务用）

**幂等建议**
- Java 调用时带 `X-Idempotency-Key`（推荐：`tenantId:merchantId:channel:outTradeNo`）。
- Go 使用 Redis 存储“请求 → 响应”结果（JSON）并设置 TTL（如 24h），重复请求直接回放。

**渠道参数透传**
- 请求：`ext`（Map 或 JSON 字符串）原样透传到渠道适配层。
- 响应：`raw`（可选）保留渠道原始响应，便于排障。

## 4. Go pay-gateway：回调接入（公网）

建议回调路由：
- `POST /callbacks/wechat/v3`
- `POST /callbacks/wechat/v2`
- `POST /callbacks/alipay`
- （可选）`POST /callbacks/paypal`

处理流程（强约束）：
1) 读取 headers/body → 解析 → 验签/解密（失败直接返回失败回执）
2) **去重**：以平台事件唯一键去重（例如 `transaction_id + event_type` / `notify_id` 等），Redis `SETNX` + TTL（如 7d）
3) 投递事件到 Java（Webhook；可选先写入 Redis Outbox 再异步投递）
4) 返回平台规定回执（微信/支付宝格式不同）

## 5. Webhook 事件契约（Go → Java）

pay-gateway 会向 Java 配置的 `javaWebhook.url` 发起 `POST`（JSON）。建议 Java：
- 校验 shared secret HMAC（`X-Pay-*` 头），只维护一份密钥（放 Nacos 下发）
- 以 `eventId` 或 `idempotencyKey` 做消费幂等（唯一索引/去重表）
- 轻量处理后快速返回 2xx（业务重活异步化）

事件 Envelope（JSON，字段建议）：
- `eventId`：UUID（全局唯一）
- `eventType`：如 `payment.succeeded`
- `eventVersion`：整数（从 1 开始，便于演进）
- `occurredAt`：RFC3339 时间
- `tenantId`
- `merchantId`
- `channel`：`WECHAT_V3`/`ALIPAY`/`PAYPAL`…
- `outTradeNo` / `transactionId` / `outRefundNo` / `refundId`（按事件类型填充）
- `amount`：最小货币单位 long
- `currency`
- `tradeState` / `refundState`（统一枚举）
- `raw`：可选（原始回调 body + 关键 headers 的脱敏快照）
- `signatureVerified`：bool（必须）
- `idempotencyKey`：用于 Java 侧幂等消费（建议：`{tenantId}:{merchantId}:{outTradeNo|outRefundNo}`）

## 6. Java `ruoyi-pay`：模块结构建议（目录级）

建议包结构（对齐现有 DDD 分层风格：controller → service → mapper → domain）：
- `org.dromara.pay`
  - `controller/`：管理端 REST（配置、支付单、退款单、手动补偿）
  - `domain/`
    - `entity/`：`PayMerchantConfig`、`PayPaymentOrder`、`PayRefundOrder`、`PayEventConsumeLog`
    - `bo/`：创建支付/退款 BO
    - `vo/`：返回给前端/调用方的 VO
  - `mapper/`
  - `service/` + `service/impl/`
  - `client/`：Go 网关 client（HTTP）
  - `webhook/`：接收 pay-gateway 事件（`javaWebhook.url` 对应 Controller）
  - `task/`：补偿查询定时任务（仅扫 `PAYING/REFUNDING`）

建议新增的 API 契约模块（可选但推荐）：
- `ruoyi-api/ruoyi-api-payment`：放 Dubbo API + DTO（供 `ruoyi-dragonboat` 等服务调用）

### 6.1 建议数据表（MySQL，按 tenant 隔离）

> 表名按你们项目习惯可加 `cb_` 前缀；以下用 `pay_` 说明语义。金额字段统一“最小货币单位 long”。

1) `pay_merchant_config`（商户/应用配置）
- `id` PK
- `tenant_id`
- `merchant_id`（业务侧商户/租户内唯一）
- `channel`（WECHAT_V3/ALIPAY/PAYPAL…）
- `config_version`（乐观锁/变更号）
- `config_cipher`（加密后的配置 JSON，或存 `secret_ref`）
- `status`（ENABLED/DISABLED）
- `create_by/update_by/create_time/update_time`
- 索引：`uniq(tenant_id, merchant_id, channel)`

2) `pay_payment_order`（支付单，状态机主表）
- `id` PK
- `tenant_id`
- `biz_order_no`（业务订单号，可空但建议填）
- `merchant_id`
- `channel`
- `currency`
- `amount`（最小货币单位）
- `out_trade_no`（你们生成，幂等核心键）
- `channel_trade_no`（平台单号，如 transaction_id/trade_no）
- `status`（INIT/PAYING/SUCCESS/CLOSED/FAILED）
- `pay_data`（下单返回的拉起参数 JSON，可选）
- `ext`（透传 JSON，可选）
- `last_error_code/last_error_msg`（可选）
- `create_time/update_time`
- 索引：`uniq(tenant_id, out_trade_no)`、`idx(tenant_id, status, update_time)`、`idx(tenant_id, biz_order_no)`

3) `pay_refund_order`（退款单）
- `id` PK
- `tenant_id`
- `merchant_id`
- `channel`
- `currency`
- `refund_amount`（最小货币单位）
- `out_trade_no`
- `channel_trade_no`（可冗余）
- `out_refund_no`（你们生成，幂等核心键）
- `channel_refund_no`（平台退款单号）
- `status`（INIT/REFUNDING/REFUNDED/FAILED）
- `reason`（可选）
- `create_time/update_time`
- 索引：`uniq(tenant_id, out_refund_no)`、`idx(tenant_id, out_trade_no)`、`idx(tenant_id, status, update_time)`

4) `pay_event_inbox`（事件消费幂等表）
- `id` PK
- `tenant_id`
- `event_id`（MQ envelope.eventId）
- `event_type`
- `source`（pay-gateway）
- `consumed_at`
- `raw`（可选，落库前先脱敏）
- 索引：`uniq(tenant_id, event_id)`

## 7. Java 状态机与补偿（强制落地）

核心原则：
- 下单同步响应 ≠ 支付成功；以 **回调事件** 为准。
- 必须有补偿：定时任务扫描超时单 → 调 Go `/v1/compensations/...` 批量查询 → 修正状态。

建议统一状态：
- Payment：`INIT` → `PAYING` → `SUCCESS|CLOSED|FAILED`
- Refund：`INIT` → `REFUNDING` → `REFUNDED|FAILED`

## 8. 安全与配置（必须）

### 8.1 Go 侧 TLS 校验（禁止 InsecureSkipVerify）
`gopay` SDK 默认 http client 配置存在 `InsecureSkipVerify` 风险；网关化时必须自定义 `http.Transport`：
- 启用系统/自建 CA 证书池
- 强制校验 TLS（不要跳过）
- 启用连接复用（keep-alive）与合理超时
- 出网域名白名单（仅支付平台域名）

### 8.2 Go ↔ Java 内网鉴权
Go 调用 Java 的“商户配置拉取/刷新”等内部接口，建议：
- **推荐：shared secret HMAC**（Nacos 配置下发，Go 与 Java 共用同一份密钥）：
  - Java → Go：调用 `/v1/**` 时带 `X-Pay-*` 签名 header
  - Go → Java：推送 webhook 事件时带 `X-Pay-*` 签名 header
  - Go → Java：拉取商户快照时带 `X-Pay-*` 签名 header
- 可选：叠加 IP 白名单；后续可演进 mTLS。
- 所有敏感字段日志脱敏，严禁打印私钥/证书全文。

### 8.3 多租户配置
- Java 管理端维护商户配置（tenant scoped）。
- Go 运行时按 `tenantId + merchantId + channel` 缓存 client；配置变更通过“Java 快照拉取 + 轮询兜底”刷新。
- 若 Java 侧已关闭多租户：Go 配置 `defaultTenantId=0`，Java 调用网关可不传 `tenantId`。

## 9. 交付清单（建议里程碑）

M1（可联调）：
- pay-gateway：微信 V3 + 支付宝 L0、回调入站、Webhook 事件推送（可选 Outbox）、Redis 幂等/去重、TLS 加固、sharedAuth
- ruoyi-pay：支付单/退款单表、事件接收（webhook）、状态机、补偿任务、管理端查询接口、商户配置快照接口

M2（上线前加固）：
- 配置管理 UI + 审计
- 失败重试与告警
- 对账/账单下载（可选）

## 10. 契约示例（便于联调）

### 10.1 创建支付（Java → Go）

请求：
```json
{
  "tenantId": "0",
  "merchantId": "mch_001",
  "channel": "WECHAT_V3",
  "scene": "JSAPI",
  "outTradeNo": "P202602010001",
  "bizOrderNo": "O202602010001",
  "currency": "CNY",
  "amount": 100,
  "subject": "订单支付",
  "description": "订单 O202602010001",
  "expireAt": "2026-02-01T13:00:00+08:00",
  "openid": "user-openid",
  "ext": {
    "attach": "{\"biz\":\"dragonboat\"}"
  }
}
```

响应（示例，payData 形态随渠道而不同）：
```json
{
  "code": "OK",
  "outTradeNo": "P202602010001",
  "status": "PAYING",
  "payData": {
    "prepayId": "wx-prepay-id",
    "appId": "wx-appid",
    "nonceStr": "xxx",
    "timeStamp": "xxx",
    "package": "prepay_id=xxx",
    "signType": "RSA",
    "paySign": "xxx"
  }
}
```

### 10.2 支付成功事件（Go → MQ → Java）

```json
{
  "eventId": "c3c9f1a8-6b88-4f39-9f4a-0c9a2c2cc9d0",
  "eventType": "payment.succeeded",
  "eventVersion": 1,
  "occurredAt": "2026-02-01T12:01:02Z",
  "tenantId": "100",
  "merchantId": "mch_001",
  "channel": "WECHAT_V3",
  "outTradeNo": "P202602010001",
  "transactionId": "420000xxxx",
  "amount": 100,
  "currency": "CNY",
  "tradeState": "SUCCESS",
  "signatureVerified": true,
  "idempotencyKey": "0:mch_001:P202602010001"
}
```
