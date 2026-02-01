# syntax=docker/dockerfile:1.6

FROM golang:1.24-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/pay-gateway ./cmd/pay-gateway


FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -S app \
    && adduser -S -G app app

COPY --from=builder /out/pay-gateway /usr/local/bin/pay-gateway

USER app
EXPOSE 8088

ENTRYPOINT ["pay-gateway"]

