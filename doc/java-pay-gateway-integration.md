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
- `apiAuth.token`：Java 调用 Go 内网 API 的鉴权 token（Header：`X-Pay-Gateway-Token`）。为空则不校验（不建议）。
- `javaWebhook.url`：Java 内网接收事件的地址（Go → Java）。
- `javaWebhook.token`：Go 请求 Java 时携带 `X-Pay-Gateway-Token`，Java 需校验。
- `tls.caFile`：可选，自定义 CA（用于 pay-gateway 出站 TLS 校验；默认系统 CA）。

建议：`apiAuth.token` 与 `javaWebhook.token` 分开配置（两个方向的鉴权密钥不要复用）。

## 2) Java → Go：核心 API（L0）

### 2.0 认证（推荐开启）

当 `apiAuth.token` 配置非空时，Java 调用所有 `/v1/**` 接口必须携带：
- `X-Pay-Gateway-Token: <apiAuth.token>`

### 2.1 创建支付

`POST /v1/payments`

Header：
- `X-Pay-Gateway-Token`：鉴权 token（推荐开启 `apiAuth.token` 后强制传）。
- `X-Idempotency-Key`：幂等键（未传时网关会按 `tenantId/merchantId/channel/outTradeNo` 生成）。当前实现为**内存缓存**，仅用于开发/单实例；生产请替换为 Redis/DB。

Body（字段）：
- `channel`：`WECHAT_V3` / `ALIPAY`
- `scene`：
  - 微信：`JSAPI`（需要 `openid`）、`MINIAPP`（需要 `openid`）、`APP`、`H5`、`NATIVE`
  - 支付宝：`APP`、`WAP`、`PAGE`、`PRECREATE`

示例（微信 JSAPI）：
```json
{
  "tenantId": "100",
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

### 2.2 查询支付

`GET /v1/payments/{outTradeNo}?tenantId=...&merchantId=...&channel=WECHAT_V3`

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

`GET /v1/refunds/{outRefundNo}?tenantId=...&merchantId=...&channel=WECHAT_V3`

支付宝退款查询额外需要：
- `outTradeNo` 或 `tradeNo`

示例：
`GET /v1/refunds/R202602010001?tenantId=100&merchantId=mch_001&channel=ALIPAY&outTradeNo=P202602010001`

### 2.6 补偿查询（批量）

`POST /v1/compensations/payments/query`
```json
{
  "tenantId": "100",
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
- Go 会先验签/解密，再推送事件给 Java；若推送失败，会对平台返回失败，让平台重试回调（实现“至少一次”）。
- 微信退款：网关在发起退款时会设置 `notify_url` 指向同一个微信回调地址；回调事件会以 `refund.*` 的 `eventType` 推送给 Java。

## 4) Go → Java：事件推送（Webhook）

Go 会向 `javaWebhook.url` 发起 `POST`（JSON），并携带：
- Header：`X-Pay-Gateway-Token: <token>`

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
- 校验 `X-Pay-Gateway-Token`
- 以 `eventId`（或 `idempotencyKey`）做消费幂等（唯一索引/去重表）
- 轻量处理后立即返回 2xx（业务重活异步化），避免阻塞支付平台回调链路
