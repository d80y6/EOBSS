FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY libs/go-common/ libs/go-common/
COPY services/billing-service/ services/billing-service/
WORKDIR /app/services/billing-service
RUN go mod init github.com/telcoflow/telcoflow/services/billing-service || true
RUN go mod tidy
RUN go build -o billing-server ./cmd/server/main.go || echo "main.go missing"

FROM alpine:latest
RUN addgroup -S telcoflow && adduser -S telcoflow -G telcoflow
USER telcoflow
WORKDIR /app
COPY --from=builder /app/services/billing-service/billing-server .
EXPOSE 8080
CMD ["./billing-server"]
