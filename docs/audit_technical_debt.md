# TelcoFlow Audit: Technical Debt Report

## Core Services
- **Assurance Service**: Empty handlers and services. Skeleton only.
- **Billing Service**: Invoicing service partially implemented; clickhouse repository missing; empty handler.
- **Catalog Service**: Empty handler and service.
- **CRM Service**: Basic CRUD implemented, but uses Mock Repository instead of GORM/PostgreSQL.
- **IAM Service**: RBAC is a mock (always returns true); empty handler.
- **Inventory Service**: Empty handler and service. NetBox client is a skeleton.
- **Order Service**: Empty handler and service. Temporal workflows are skeletons.
- **Provisioning Service**: Empty handler and service. Adapters (Radius, Open5GS, Kamailio) are stubs.

## Libraries
- **go-common**: Contains useful modules (Kafka, Vault, Resilience) but they are not integrated into any microservice.

## Infrastructure
- **K8s/Helm**: Helm charts exist but lack full environment variable mappings and secret management for the current microservice state.
- **Docker**: Most Dockerfiles are present but need review for production hardening.
