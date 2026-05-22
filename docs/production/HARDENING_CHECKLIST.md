# TelcoFlow Production Hardening Checklist

## 1. Security
- [ ] Enable mTLS between all microservices (Istio/Linkerd).
- [ ] Configure RBAC/ABAC in Keycloak.
- [ ] Rotate all database and Kafka credentials in Vault.
- [ ] Enable Kubernetes NetworkPolicies for namespace isolation.
- [ ] Configure PodSecurityPolicies or Admission Controllers.

## 2. Reliability & Resilience
- [ ] Configure Horizontal Pod Autoscalers (HPA) for all core services.
- [ ] Set up Multi-AZ deployments for PostgreSQL and Kafka.
- [ ] Implement Circuit Breakers for all inter-service gRPC calls.
- [ ] Verify Temporal worker redundancy.

## 3. Observability
- [ ] Verify OpenTelemetry trace propagation across the Saga workflows.
- [ ] Configure Grafana alerting for Telecom KPIs (e.g., Order Failure Rate > 1%).
- [ ] Ensure all logs are structured and ingested by Loki.

## 4. Performance
- [ ] Optimize PostgreSQL indexing for TMF622 order queries.
- [ ] Fine-tune Kafka partition counts for billing usage events.
- [ ] Enable Redis persistence for session management.
