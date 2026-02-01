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
- 微信：`wechatV3.privateKeyFile`
- 支付宝：`alipay.privateKeyFile` / `alipay.alipayPublicKeyFile`

容器挂载（compose 已包含）：
- `./config -> /config:ro`
- `./secrets -> /secrets:ro`

## 3) .env 覆盖（推荐用于容器化）

pay-gateway 支持用环境变量覆盖常用项（见 `cmd/pay-gateway/.env.example`），例如：
- `PAY_GATEWAY_PUBLIC_BASE_URL`（回调域名）
- `PAY_GATEWAY_API_AUTH_TOKEN`（Java→Go 内网鉴权）
- `PAY_GATEWAY_JAVA_WEBHOOK_URL` / `PAY_GATEWAY_JAVA_WEBHOOK_TOKEN`（Go→Java 事件推送）
- `PAY_GATEWAY_REDIS_ADDR`（启用 Redis 幂等/去重）

## 4) 鉴权如何结合 ruoyi-auth（推荐做法）

不要让 Go 网关直接校验 Sa-Token 用户态 token（会引入强耦合与额外依赖链路）。推荐：
- **服务级鉴权**：Java 调用 Go 时带 `X-Pay-Gateway-Token`，Go 只做字符串校验（`apiAuth.token`）。
- **统一管密钥**：把 `PAY_GATEWAY_API_AUTH_TOKEN` / `PAY_GATEWAY_JAVA_WEBHOOK_TOKEN` 作为配置项放到 Nacos（或你们的配置中心）统一下发与轮转；避免硬编码与密钥复用。

## 5) Redis（上线前必配）

配置 `redis.addr`（或 `PAY_GATEWAY_REDIS_ADDR`）后，网关会启用：
- 下单/退款接口幂等（跨实例）
- 平台回调去重（跨实例）
- （可选）回调事件 Outbox（异步投递 Java webhook）

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
