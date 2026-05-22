FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY libs/go-common/ libs/go-common/
COPY services/inventory-service/ services/inventory-service/
WORKDIR /app/services/inventory-service
RUN go mod init github.com/telcoflow/telcoflow/services/inventory-service || true
RUN go mod tidy
RUN go build -o inventory-server ./cmd/server/main.go || echo "main.go missing"

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/services/inventory-service/inventory-server .
EXPOSE 8080
CMD ["./inventory-server"]
