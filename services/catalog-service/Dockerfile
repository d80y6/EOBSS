FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY libs/go-common/ libs/go-common/
COPY services/catalog-service/ services/catalog-service/
WORKDIR /app/services/catalog-service
RUN go mod init github.com/telcoflow/telcoflow/services/catalog-service || true
RUN go mod tidy
RUN go build -o catalog-server ./cmd/server/main.go || echo "main.go missing"

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/services/catalog-service/catalog-server .
EXPOSE 8080
CMD ["./catalog-server"]
