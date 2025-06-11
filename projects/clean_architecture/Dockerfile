FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /app/ordersystem ./cmd/ordersystem


# Run stage
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/ordersystem .
COPY /cmd/ordersystem/.env .

RUN chmod +x /app/ordersystem

CMD ["./ordersystem"]
