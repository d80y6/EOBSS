# Dependency Graph Analysis

- **Upstream**: IAM (Auth), Vault (Secrets).
- **Core**: Order Service -> CRM, Catalog, Provisioning, Billing.
- **Data**: All Services -> PostgreSQL; Billing -> ClickHouse.
- **Event**: All -> Kafka.
