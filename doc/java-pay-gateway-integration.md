# Java 对接 pay-gateway（基于 GoPay SDK）

`pay-gateway` 是一个 Go HTTP 服务（见 `cmd/pay-gateway`），把本仓库 `gopay` 的渠道能力（微信 V3 / 支付宝）以统一 API 暴露给 Java，同时 **支付平台回调只落在 Go**（验签/解密），再把“可信事件”推送给 Java。

> 金额字段统一使用 **Long（最小货币单位）** 传输，禁止 `double/float`。

## 1) 部署与配置（Go）

启动命令：
```bash
go run ./cmd/pay-gateway --config /path/to/pay-gateway.json
```

配置示例：`cmd/pay-gateway/config.example.json`
- `publicBaseUrl`：外网可访问的域名，用于生成回调地址（微信/支付宝会回调到 Go）。
- `defaultTenantId`：当 Java 侧关闭多租户时可设置为固定值（默认 `0`）。此时 Java 调用 `/v1/**` 可不传 `tenantId`。
- `sharedAuth.*`：**推荐**，Go ↔ Java 使用同一个 shared secret 做 HMAC 鉴权与签名（只维护一份密钥）。
- `apiAuth.token`：兼容旧模式的 token（Header：`X-Pay-Gateway-Token`）。建议仅作为过渡，最终切到 `sharedAuth`。
- `javaWebhook.url`：Java 内网接收事件的地址（Go → Java）。
- `javaWebhook.token`：兼容旧模式的 token（Go → Java Header：`X-Pay-Gateway-Token`）。若启用 `sharedAuth`，建议留空。
- `javaWebhook.async`：是否启用“回调事件先入 Redis Outbox，再异步投递 Java webhook”（推荐开启，需要 Redis）。
- `javaWebhook.consumerGroup`：Outbox 消费者组（多实例部署时用于协作消费）。
- `tls.caFile`：可选，自定义 CA（用于 pay-gateway 出站 TLS 校验；默认系统 CA）。
- `redis.*`：强烈建议配置（用于 **幂等** 与 **回调去重**）。不配置时网关会退化为“内存幂等/去重”，仅适用于单实例开发环境。
- `merchantSync.*`：可选，Go 定时从 Java 拉取“商户配置快照”，实现“配置变更无需重启网关”（推荐生产使用）。

> 多实例 + `sharedAuth` 建议必配 Redis：网关会用 Redis 做 nonce 防重放（单实例开发环境可退化为内存）。

### 配置覆盖（.env）

pay-gateway 支持使用环境变量覆盖部分配置（便于容器化部署），常用项：
- `PAY_GATEWAY_PUBLIC_BASE_URL`
- `PAY_GATEWAY_DEFAULT_TENANT_ID`
- `PAY_GATEWAY_SHARED_SECRET` / `PAY_GATEWAY_SHARED_SECRET_PREV`
- `PAY_GATEWAY_SHARED_AUTH_REQUIRED`
- `PAY_GATEWAY_API_AUTH_TOKEN`（legacy）
- `PAY_GATEWAY_JAVA_WEBHOOK_URL`
- `PAY_GATEWAY_JAVA_WEBHOOK_TOKEN`
- `PAY_GATEWAY_MERCHANT_SNAPSHOT_URL`
- `PAY_GATEWAY_REDIS_ADDR` / `PAY_GATEWAY_REDIS_PASSWORD` / `PAY_GATEWAY_REDIS_DB` / `PAY_GATEWAY_REDIS_KEY_PREFIX`

## 2) Java → Go：核心 API（L0）

### 2.0 认证（推荐：sharedAuth，只维护一份密钥）

pay-gateway 支持两种鉴权方式（二选一，建议使用 sharedAuth）：

1) **Shared Secret HMAC（推荐）**

Java 调用所有 `/v1/**` 接口携带以下 Header：
- `X-Pay-Gateway-Timestamp`: UNIX 秒
- `X-Pay-Gateway-Nonce`: 随机串（建议 16 bytes）
- `X-Pay-Gateway-Body-SHA256`: `base64(sha256(body))`（GET/空 body 使用空串的 sha）
- `X-Pay-Gateway-Signature`: `base64(hmac_sha256(secret, canonical))`

canonical 串（精确到最后一个换行）：
```text
METHOD + "\n" +
REQUEST_URI + "\n" +
TIMESTAMP + "\n" +
NONCE + "\n" +
BODY_SHA256 + "\n"
```

其中 `REQUEST_URI` 必须包含 query string（例如 `/v1/payments/xxx?merchantId=...&channel=WECHAT_V3`）。

2) **Legacy Token（过渡用）**

当 `apiAuth.token` 配置非空时，Java 调用 `/v1/**` 也可携带：
- `X-Pay-Gateway-Token: <apiAuth.token>`

> 推荐做法：把 `sharedAuth.sharedSecret` 放到 Nacos 统一下发（Go 与 Java 共用），避免维护两套 token。

### 2.1 创建支付

`POST /v1/payments`

Header：
- `X-Pay-Gateway-*`：sharedAuth 鉴权 header（推荐）。
- `X-Pay-Gateway-Token`：legacy token（可选）。
- `X-Idempotency-Key`：幂等键（未传时网关会按 `tenantId/merchantId/channel/outTradeNo` 生成）。网关默认内存幂等；配置 `redis.addr`（或 `PAY_GATEWAY_REDIS_ADDR`）后会自动切换为 Redis 幂等（推荐）。

Body（字段）：
- `channel`：`WECHAT_V3` / `ALIPAY`
- `scene`：
  - 微信：`JSAPI`（需要 `openid`）、`MINIAPP`（需要 `openid`）、`APP`、`H5`、`NATIVE`
  - 支付宝：`APP`、`WAP`、`PAGE`、`PRECREATE`

示例（微信 JSAPI）：
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
  "openid": "user-openid"
}
```

说明：
- 当 Go 配置了 `defaultTenantId` 且 Java 侧关闭多租户时，`tenantId` 可以不传，网关会自动补齐为 `defaultTenantId`。

### 2.2 查询支付

`GET /v1/payments/{outTradeNo}?merchantId=...&channel=WECHAT_V3[&tenantId=...]`

### 2.3 关单

`POST /v1/payments/{outTradeNo}/close`
```json
{ "tenantId":"100", "merchantId":"mch_001", "channel":"WECHAT_V3" }
```

### 2.4 发起退款

`POST /v1/refunds`

注意：
- **微信 V3 必须传 `totalAmount`**（原订单金额，最小货币单位），以及 `refundAmount`（本次退款金额）。
- 支付宝目前仅支持 `CNY`（网关会把“分 → 元字符串”）。

示例（微信退款）：
```json
{
  "tenantId": "100",
  "merchantId": "mch_001",
  "channel": "WECHAT_V3",
  "outTradeNo": "P202602010001",
  "outRefundNo": "R202602010001",
  "currency": "CNY",
  "totalAmount": 100,
  "refundAmount": 100,
  "reason": "订单取消"
}
```

### 2.5 查询退款

`GET /v1/refunds/{outRefundNo}?merchantId=...&channel=WECHAT_V3[&tenantId=...]`

支付宝退款查询额外需要：
- `outTradeNo` 或 `tradeNo`

示例：
`GET /v1/refunds/R202602010001?tenantId=100&merchantId=mch_001&channel=ALIPAY&outTradeNo=P202602010001`

### 2.6 补偿查询（批量）

`POST /v1/compensations/payments/query`
```json
{
  "tenantId": "0",
  "merchantId": "mch_001",
  "channel": "WECHAT_V3",
  "outTradeNos": ["P202602010001", "P202602010002"]
}
```

建议 Java 侧定时任务：
- 扫描 `PAYING/REFUNDING` 超时单 → 调该接口批量查询 → 修正状态（避免仅靠回调）。

## 3) 支付平台 → Go：回调地址（必须落 Go）

pay-gateway 会按订单自动设置回调地址（你不需要 Java 侧拼接 notify_url）：
- 微信 V3：`POST /callbacks/wechat/v3/{tenantId}/{merchantId}`
- 支付宝：`POST /callbacks/alipay/{tenantId}/{merchantId}`

说明：
- 同步模式（`javaWebhook.async=false`）：Go 会先验签/解密，再同步推送事件给 Java；若推送失败，会对平台返回失败，让平台重试回调（实现“至少一次”）。
- 异步模式（`javaWebhook.async=true`，推荐）：Go 会先验签/解密 → 写入 Redis Outbox → 立即对平台返回成功回执；后台 worker 负责重试投递到 Java webhook（更稳健，避免依赖平台重试节奏）。
- 微信退款：网关在发起退款时会设置 `notify_url` 指向同一个微信回调地址；回调事件会以 `refund.*` 的 `eventType` 推送给 Java。
- 去重策略（建议启用 Redis）：平台回调可能重复投递。网关会对同一事件做去重；若事件已投递成功，会直接对平台返回成功回执。

## 4) Go → Java：事件推送（Webhook）

Go 会向 `javaWebhook.url` 发起 `POST`（JSON），并携带：
- sharedAuth（推荐）：`X-Pay-Gateway-*` 一组签名 header（同 2.0）
- legacy token（可选）：`X-Pay-Gateway-Token: <token>`

事件体（示例）：
```json
{
  "eventId": "WECHAT_V3:0f3d6b9c-7a0a-4b7d-9b3b-2d6b5f6f2caa",
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
  "idempotencyKey": "100:mch_001:P202602010001"
}
```

常见 `eventType`：
- `payment.succeeded` / `payment.closed` / `payment.updated`
- `refund.succeeded` / `refund.closed` / `refund.failed` / `refund.updated`

Java 接收端要求（强制）：
- 校验 sharedAuth 签名（推荐）；或校验 `X-Pay-Gateway-Token`（legacy）
- 以 `eventId`（或 `idempotencyKey`）做消费幂等（唯一索引/去重表）
- 轻量处理后立即返回 2xx（业务重活异步化），避免阻塞支付平台回调链路

## 5) （推荐）Java 提供商户配置快照接口（Go 拉取）

当配置 `merchantSync.snapshotUrl`（或 `PAY_GATEWAY_MERCHANT_SNAPSHOT_URL`）后，pay-gateway 会定时 `GET` 拉取：
- 响应 JSON（示例）：
```json
{
  "version": "2026-02-01T12:00:00Z",
  "merchants": [
    {
      "tenantId": "0",
      "merchantId": "mch_001",
      "wechatV3": { "appId":"...", "mchId":"...", "serialNo":"...", "apiV3Key":"...", "privateKeyFile":"/secrets/wechat/apiclient_key.pem" },
      "alipay": { "isProd": false, "appId":"...", "privateKeyFile":"/secrets/alipay/app_private_key.pem", "alipayPublicKeyFile":"/secrets/alipay/alipay_public_key.pem" }
    }
  ]
}
```

建议：
- Java 使用数据库加密存储商户密钥（或 Nacos 配置中心托管），对外仅暴露“快照”给 pay-gateway。
- 若启用 `sharedAuth.sharedSecret`，该快照接口也应按 2.0 校验签名（Go 会对 GET 请求签名）。
