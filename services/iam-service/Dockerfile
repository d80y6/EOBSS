FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY libs/go-common/ libs/go-common/
COPY services/iam-service/ services/iam-service/

WORKDIR /app/services/iam-service
RUN go mod download
RUN go build -o iam-server ./cmd/server/main.go

FROM alpine:latest
RUN addgroup -S telcoflow && adduser -S telcoflow -G telcoflow
USER telcoflow
WORKDIR /app
COPY --from=builder /app/services/iam-service/iam-server .

EXPOSE 8080
CMD ["./iam-server"]
