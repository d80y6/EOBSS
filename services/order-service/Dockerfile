FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY libs/go-common/ libs/go-common/
COPY services/order-service/ services/order-service/
WORKDIR /app/services/order-service
RUN go mod init github.com/telcoflow/telcoflow/services/order-service || true
RUN go mod tidy
RUN go build -o order-server ./cmd/server/main.go || echo "main.go missing"

FROM alpine:latest
RUN addgroup -S telcoflow && adduser -S telcoflow -G telcoflow
USER telcoflow
WORKDIR /app
COPY --from=builder /app/services/order-service/order-server .
EXPOSE 8080
CMD ["./order-server"]
