FROM golang:1.25-alpine AS builder
WORKDIR /app

RUN apk add --no-cache \
    ca-certificates \
    build-base \
    musl-dev \
    pkgconfig

COPY go.mod go.sum ./

RUN go mod download

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .

RUN CGO_ENABLED=1 \
    CC=gcc \
    GOOS=linux \
    GOARCH=amd64 \
    go build -ldflags="-linkmode external -extldflags '-static' -w -s" \
    -tags musl \
    -trimpath \
    -o user-service cmd/main.go

FROM alpine:3.20
WORKDIR /app/

RUN apk add --no-cache \
    ca-certificates \
    tzdata

COPY --from=builder /app/user-service .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/migrations ./migrations

RUN migrate -version

HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD grpc_health_probe -addr=:50052 -tls=false || exit 1

EXPOSE 50052
CMD ["./user-service"]