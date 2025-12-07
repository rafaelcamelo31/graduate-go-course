# syntax=docker/dockerfile:1

ARG GO_VERSION=1.25.4

FROM golang:${GO_VERSION} AS base
WORKDIR /app
COPY temperature_service/go.mod temperature_service/go.sum ./temperature_service/
COPY observability/go.mod observability/go.sum ./observability/
WORKDIR /app/temperature_service
RUN go mod download

FROM base AS builder
WORKDIR /app
COPY temperature_service/ ./temperature_service/
COPY observability/ ./observability/
WORKDIR /app/temperature_service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C ./cmd -o temperature

FROM builder AS tester
RUN go test -v ./...

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /app/bin
COPY --from=builder /app/temperature_service/cmd/temperature .
COPY --from=builder /app/temperature_service/.env /app
ENTRYPOINT [ "./temperature" ]
