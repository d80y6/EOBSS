FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY libs/go-common/ libs/go-common/
COPY services/crm-service/ services/crm-service/
WORKDIR /app/services/crm-service
RUN go mod init github.com/telcoflow/telcoflow/services/crm-service || true
RUN go mod tidy
RUN go build -o crm-server ./cmd/server/main.go || echo "main.go missing"

FROM alpine:latest
RUN addgroup -S telcoflow && adduser -S telcoflow -G telcoflow
USER telcoflow
WORKDIR /app
COPY --from=builder /app/services/crm-service/crm-server .
EXPOSE 8080
CMD ["./crm-server"]
