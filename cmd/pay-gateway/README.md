# pay-gateway (Go)

`pay-gateway` is a lightweight HTTP service that wraps `gopay` (this repo) so a Java backend can call payment capabilities remotely, while **payment platform callbacks land in Go** (verify + decrypt), and then are forwarded to Java as trusted events.

This folder is intentionally dependency-light (std + `gopay` only). For production, add durable idempotency (Redis), durable outbox, and structured logging/metrics.

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

