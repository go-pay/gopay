# Repository Guidelines

GoPay is a Go SDK that wraps multiple payment providers (WeChat, Alipay, PayPal, etc.). This guide covers the basics for contributing.

## Project Structure & Module Organization

- Root package (`/`) contains shared types and helpers (for example `BodyMap`, common constants).
- Provider packages live in top-level folders: `alipay/`, `wechat/`, `paypal/`, `qq/`, `apple/`, `allinpay/`, `lakala/`, `saobei/`.
  - Versioned APIs are nested (for example `alipay/v3/`, `wechat/v3/`).
- Shared utilities are in `pkg/` (for example `pkg/xhttp/`, `pkg/jwt/`).
- Provider docs live in `doc/` (one Markdown file per provider).
- Runnable samples live in `examples/`.

## Build, Test, and Development Commands

Run:
```bash
go test ./...
```

Run with verbose output (matches CI):
```bash
go test -v ./...
```

Run lint (matches `.github/workflows/go.yml`):
```bash
golangci-lint run -v --disable=unused ./...
```

Format code:
```bash
gofmt -w .
```

Try an example:
```bash
go run ./examples/wechat/wx_UnifiedOrder.go
```

## Coding Style & Naming Conventions

- Use `gofmt` formatting; don’t hand-align code.
- Keep provider-specific code in its provider folder; avoid leaking provider constants into the root package.
- Exported identifiers use `PascalCase`; files use `snake_case.go` when needed.

## Testing Guidelines

- Tests are colocated as `*_test.go`. Prefer deterministic unit tests for signing, parsing, and serialization.
- If a test needs real credentials or network calls, read config from environment variables and `t.Skip` when missing—never commit production keys/certs.

## Commit & Pull Request Guidelines

- Commit subjects are short and action-oriented (common patterns: “Fix …”, “Update …”, “Support …”), often scoped by provider (example: `wechat/v3: add ...`) and may include `(#123)` when applicable.
- PRs should include: what changed, why, how to test (`go test ./...`), and any doc updates in `doc/`. For releases, also update `constant.go` (`Version`) and `release_note.md`.

