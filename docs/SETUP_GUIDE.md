# TelcoFlow Platform Setup Guide

## Prerequisites
- Docker & Docker Compose
- Kubernetes Cluster (v1.28+)
- Helm 3
- Go 1.24+
- Python 3.11+
- Node.js 20+

## Local Development
1. Clone the repository.
2. Run `docker-compose up -d` to start infrastructure dependencies (Postgres, Kafka, Temporal, Keycloak).
3. Build services: `make build`.
4. Run individual services using `go run cmd/server/main.go` within each service directory.

## Kubernetes Deployment
1. Install shared infrastructure:
   ```bash
   helm install infra ./infra/helm/telcoflow-infra
   ```
2. Apply Kubernetes Gateway manifests:
   ```bash
   kubectl apply -f ./infra/k8s/gateway/
   ```
3. Deploy microservices using ArgoCD or raw manifests:
   ```bash
   kubectl apply -f ./services/order-service/k8s/ # (Generate these per service)
   ```

## Monitoring
- Access Grafana at `http://localhost:3000` (default port in k8s).
- Import dashboards from `infra/k8s/observability/`.
