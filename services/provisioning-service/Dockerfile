FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY libs/go-common/ libs/go-common/
COPY services/provisioning-service/ services/provisioning-service/
WORKDIR /app/services/provisioning-service
RUN go mod init github.com/telcoflow/telcoflow/services/provisioning-service || true
RUN go mod tidy
RUN go build -o provisioning-server ./cmd/server/main.go || echo "main.go missing"

FROM alpine:latest
RUN addgroup -S telcoflow && adduser -S telcoflow -G telcoflow
USER telcoflow
WORKDIR /app
COPY --from=builder /app/services/provisioning-service/provisioning-server .
EXPOSE 8080
CMD ["./provisioning-server"]
