# pay-gateway (Go)

`pay-gateway` is a lightweight HTTP service that wraps `gopay` (this repo) so a Java backend can call payment capabilities remotely, while **payment platform callbacks land in Go** (verify + decrypt), and then are forwarded to Java as trusted events.

This folder is intentionally dependency-light (std + `gopay`, plus optional Redis for idempotency/de-dup). For production, consider durable outbox (DB/MQ) and structured logging/metrics.

Deployment notes: see `doc/pay-gateway-deploy.md`.

## Run

Create a config file:
```bash
cp cmd/pay-gateway/config.example.json /tmp/pay-gateway.json
```

Start:
```bash
go run ./cmd/pay-gateway --config /tmp/pay-gateway.json
```

Health check:
```bash
curl -sS http://127.0.0.1:8088/healthz
```

## Auth (internal APIs)

If `apiAuth.token` is set in config, Java callers must send `X-Pay-Gateway-Token` for all `/v1/**` endpoints.

## Run with Docker Compose (local)

From `cmd/pay-gateway`:
```bash
cp .env.example .env
mkdir -p config secrets
cp config.example.json config/pay-gateway.json
docker compose up -d --build
```

You can keep most values in `.env` and leave `config/pay-gateway.json` focused on merchant configs (keys/certs should be mounted as files under `secrets/`).

## Redis (recommended)

If `PAY_GATEWAY_REDIS_ADDR` (or `redis.addr` in config) is set, pay-gateway will use Redis for:
- Idempotency (create payment/refund)
- Callback de-duplication across instances

Example secret layout:
```text
cmd/pay-gateway/secrets/
  wechat/
    apiclient_key.pem
  alipay/
    app_private_key.pem
    alipay_public_key.pem
```
