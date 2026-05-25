# Production-grade Dockerfile for Go services
FROM golang:1.24-alpine AS builder

# Add CA certificates and build dependencies
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy shared libraries
COPY libs/go-common/ libs/go-common/

# ARG to allow building different services
ARG SERVICE_NAME
COPY services/${SERVICE_NAME}/ services/${SERVICE_NAME}/

WORKDIR /app/services/${SERVICE_NAME}
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/server/main.go

# Use Distroless for minimal attack surface
FROM gcr.io/distroless/static-debian12:latest-amd64
WORKDIR /
COPY --from=builder /app/server /server
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/server"]
