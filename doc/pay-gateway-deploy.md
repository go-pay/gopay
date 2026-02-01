# pay-gateway 部署与配置（Docker/.env）

本仓库的 `cmd/pay-gateway` 是一个 Go HTTP 服务：Java（ruoyi-pay）通过内网调用 `/v1/**`，而支付平台回调只落在 Go（验签/解密），再把可信事件推送到 Java webhook。

## 1) 本地启动（Docker Compose）

在 `cmd/pay-gateway` 目录执行：
```bash
cp .env.example .env
mkdir -p config secrets
cp config.example.json config/pay-gateway.json
docker compose up -d --build
```

## 2) 配置文件与密钥文件

建议把“结构化配置”放在 `config/pay-gateway.json`（商户列表、渠道参数），把“敏感多行密钥/证书”放在 `secrets/` 并通过文件引用：
- 微信：`wechatV3.privateKeyRef`
- 支付宝：`alipay.privateKeyRef` / `alipay.alipayPublicKeyRef`

pay-gateway 会把 `*Ref` 解析为 `${secretsBaseDir}/${ref}`（默认 `secretsBaseDir=/secrets`），并拒绝任何越界路径。

容器挂载（compose 已包含）：
- `./config -> /config:ro`
- `./secrets -> /secrets:ro`

## 3) .env 覆盖（推荐用于容器化）

pay-gateway 支持用环境变量覆盖常用项（见 `cmd/pay-gateway/.env.example`），例如：
- `PAY_GATEWAY_PUBLIC_BASE_URL`（回调域名）
- `PAY_GATEWAY_DEFAULT_TENANT_ID`（多租户关闭时可固定为 `0`）
- `PAY_GATEWAY_SHARED_SECRET`（推荐：Go ↔ Java 共用的 shared secret，用于 HMAC 鉴权与签名）
- `PAY_GATEWAY_SHARED_SECRET_PREV`（可选：用于密钥轮转的旧密钥）
- `PAY_GATEWAY_SHARED_AUTH_REQUIRED`（建议 `true`：只接受 sharedAuth，不再接受 legacy token）
- `PAY_GATEWAY_API_AUTH_TOKEN`（legacy：Java→Go token 鉴权，建议过渡期使用）
- `PAY_GATEWAY_JAVA_WEBHOOK_URL` / `PAY_GATEWAY_JAVA_WEBHOOK_TOKEN`（Go→Java 事件推送）
- `PAY_GATEWAY_MERCHANT_SNAPSHOT_URL`（可选：Go 定时拉取 Java 商户配置快照）
- `PAY_GATEWAY_REDIS_ADDR`（启用 Redis 幂等/去重）

## 4) 鉴权如何结合 ruoyi-auth（推荐做法）

不要让 Go 网关直接校验 Sa-Token 用户态 token（会引入强耦合、额外 RPC/Redis 依赖链路，且难以在 Go 侧复用 Sa-Token 生态）。推荐：

- **用户态权限**：仍由 Java（ruoyi-auth + Sa-Token）处理。
- **服务级鉴权（推荐）**：Go ↔ Java 使用同一个 `sharedAuth.sharedSecret` 做 HMAC 鉴权/签名（Header：`X-Pay-*` 一组），这样只维护 **一份密钥**。
- **密钥管理（Nacos）**：把 shared secret 放到 Nacos 统一下发，并通过 `PAY_GATEWAY_SHARED_SECRET` 注入 pay-gateway；Java 模块同样从 Nacos 读取并校验/签名。

> 仍可保留 `X-Pay-Token`（`apiAuth.token` / `javaWebhook.token`）作为过渡，但建议最终只保留 sharedAuth。

## 5) Redis（上线前必配）

配置 `redis.addr`（或 `PAY_GATEWAY_REDIS_ADDR`）后，网关会启用：
- 下单/退款接口幂等（跨实例）
- 平台回调去重（跨实例）
- （可选）回调事件 Outbox（异步投递 Java webhook）
- （推荐）sharedAuth nonce 防重放（跨实例）

建议：回调去重 key TTL ≥ 7 天；幂等 key TTL ≥ 24 小时（可按业务调整）。

## 6) 回调事件 Outbox（推荐开启）

当 `javaWebhook.async=true` 时，回调处理链路变为：
1) 验签/解密
2) 写入 Redis Stream（默认 stream：`${PAY_GATEWAY_REDIS_KEY_PREFIX}outbox`）
3) 立即对平台返回成功
4) 后台 worker 从 consumer group 消费并重试投递到 Java webhook

多实例部署时请确保：
- 所有实例连接同一 Redis
- `javaWebhook.consumerGroup` 统一（默认 `pay-gateway`）

## 7) 商户配置快照（推荐生产使用）

pay-gateway 支持从 Java 拉取“商户配置快照”，用于：
- 新增商户/改密钥无需重启 pay-gateway
- Go 服务保持无状态（只缓存 client）

启用方式：
- Java 提供一个内网 `GET` 快照接口（返回 `{version, merchants}`）
- 配置 `PAY_GATEWAY_MERCHANT_SNAPSHOT_URL`
- 建议同时开启 `PAY_GATEWAY_SHARED_SECRET`，让 Go 请求快照接口带签名，Java 校验后再返回
