# syntax=docker/dockerfile:1

ARG GO_VERSION=1.25.4

FROM golang:${GO_VERSION} AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM base AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C ./cmd -o temperature

FROM builder AS tester
RUN go test -v ./...

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /app/bin
COPY --from=builder /app/cmd/temperature .
COPY --from=builder /app/.env /app
ENTRYPOINT [ "./temperature" ]
