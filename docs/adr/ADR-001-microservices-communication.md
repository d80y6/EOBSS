# ADR-001: Microservices Communication Strategy

## Status
Accepted

## Context
TelcoFlow requires a robust, scalable, and loosely coupled communication mechanism between its various domains (CRM, Catalog, Billing, Inventory, etc.).

## Decision
We will adopt a hybrid approach:
1. **Synchronous (REST/gRPC)**: For UI-to-Service and real-time query operations. gRPC will be preferred for high-performance internal service-to-service communication. All external-facing APIs must comply with TM Forum Open API standards (REST).
2. **Asynchronous (Event-Driven)**: For state changes, cross-domain notifications, and long-running processes using **Apache Kafka**.

## Consequences
- Requires a Schema Registry for Kafka to manage event contracts.
- Increased complexity in handling distributed transactions (solved by Sagas/Temporal).
- Improved scalability and resilience.
