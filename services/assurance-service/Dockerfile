FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY libs/go-common/ libs/go-common/
COPY services/assurance-service/ services/assurance-service/
WORKDIR /app/services/assurance-service
RUN go mod init github.com/telcoflow/telcoflow/services/assurance-service || true
RUN go mod tidy
RUN go build -o assurance-server ./cmd/server/main.go || echo "main.go missing"

FROM alpine:latest
RUN addgroup -S telcoflow && adduser -S telcoflow -G telcoflow
USER telcoflow
WORKDIR /app
COPY --from=builder /app/services/assurance-service/assurance-server .
EXPOSE 8080
CMD ["./assurance-server"]
