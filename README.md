# TelcoFlow OSS/BSS Platform

TelcoFlow is a production-grade, cloud-native, AI-ready OSS/BSS platform designed for modern telecom operators (MVNOs, ISPs, 5G/LTE providers). It follows TM Forum ODA (Open Digital Architecture) and Open API standards.

## Project Structure

- `apps/`: Frontend applications (Next.js, React).
  - `admin-portal`: Internal management portal for CSRs, Engineers, and Administrators.
  - `self-care-portal`: Customer-facing portal for account management and shopping.
- `services/`: Microservices (Go, Java/Spring Boot, Python).
  - `iam-service`: Identity and Access Management (Keycloak integration).
  - `crm-service`: Customer Relationship Management (TMF629/632).
  - `catalog-service`: Product Catalog Management (TMF620).
  - `order-service`: Service & Product Order Management (TMF622/641).
  - `billing-service`: Rating, Billing, and Invoicing (TMF678).
  - `provisioning-service`: Service Activation and Network Integration.
  - `inventory-service`: Resource and Service Inventory (TMF639/640).
  - `assurance-service`: Fault and Performance Management (TMF642).
  - `ai-ops-service`: AI/ML for Predictive Maintenance, Fraud, and Churn.
- `infra/`: Infrastructure as Code and Deployment manifests.
  - `terraform`: Cloud infrastructure provisioning.
  - `helm`: Kubernetes package management.
  - `k8s`: Raw Kubernetes manifests for core components.
- `libs/`: Shared libraries and utilities for different runtimes.
- `api/specs`: TM Forum Open API specifications and custom gRPC/GraphQL definitions.

## Core Principles

- **API-First**: Every capability is exposed via a well-defined API (TMF Open APIs).
- **Event-Driven**: Asynchronous communication via Kafka for decoupled workflows.
- **Cloud-Native**: Designed for Kubernetes with horizontal scaling and high availability.
- **Model-Driven**: Logic driven by metadata and product catalog definitions.
- **Zero Trust**: Security at every layer with OIDC/OAuth2.

## Technology Stack

- **Languages**: Go (Performance), Java (Orchestration), Python (AI/Data), TypeScript (Frontend).
- **Database**: PostgreSQL (Relational), Redis (Caching), ClickHouse (Analytics).
- **Messaging**: Apache Kafka.
- **Orchestration**: Temporal (Workflows).
- **Observability**: Prometheus, Grafana, Loki, OpenTelemetry.
- **Infrastructure**: Kubernetes, Helm, Terraform.
