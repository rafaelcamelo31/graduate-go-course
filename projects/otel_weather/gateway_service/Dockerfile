# syntax=docker/dockerfile:1

ARG GO_VERSION=1.25.4

FROM golang:${GO_VERSION} AS base
WORKDIR /app
COPY gateway_service/go.mod gateway_service/go.sum ./gateway_service/
COPY observability/go.mod observability/go.sum ./observability/
WORKDIR /app/gateway_service
RUN go mod download

FROM base AS builder
WORKDIR /app
COPY gateway_service/ ./gateway_service/
COPY observability/ ./observability/
WORKDIR /app/gateway_service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C ./cmd -o gateway

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /app/bin
COPY --from=builder /app/gateway_service/cmd/gateway .
ENTRYPOINT [ "./gateway" ]
