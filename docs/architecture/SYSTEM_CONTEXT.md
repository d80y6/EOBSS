# TelcoFlow System Context

## Overview
TelcoFlow is a cloud-native OSS/BSS platform designed to manage the entire lifecycle of telecom services and customers.

## High-Level Architecture (Mermaid)

```mermaid
graph TD
    subgraph "Engagement Layer (Apps)"
        AdminPortal[Admin Portal]
        SelfCarePortal[Self-Care Portal]
        NOCPortal[NOC Portal]
    end

    subgraph "Core Business Domains (Go/Java)"
        IAM[IAM Service]
        CRM[CRM Service]
        Catalog[Product Catalog]
        Order[Order Management]
        Billing[Billing Service]
    end

    subgraph "Production Domains (Go)"
        Inventory[Inventory Service]
        Provisioning[Provisioning Service]
        Assurance[Assurance Service]
    end

    subgraph "Intelligence Layer (Python)"
        AIOps[AI Operations]
    end

    subgraph "Infrastructure & Data"
        Kafka[Kafka - Event Bus]
        Postgres[PostgreSQL - System of Record]
        Redis[Redis - Cache]
        Temporal[Temporal - Workflow Engine]
        ClickHouse[ClickHouse - Analytics]
    end

    %% Connections
    EngagementLayer --> IAM
    EngagementLayer --> CoreBusinessDomains
    CoreBusinessDomains --> Kafka
    CoreBusinessDomains --> Postgres
    CoreBusinessDomains --> Temporal
    Order --> Inventory
    Provisioning --> Inventory
    Assurance --> Inventory
    Kafka --> AIOps
    Kafka --> ClickHouse
```

## Integration Points
- **Northbound**: Open APIs (TMF) for external partner and digital channel integration.
- **Southbound**: Provisioning adapters for network elements (Routers, OLTs, 5G Cores).
- **Security**: Keycloak as the Identity Provider (IdP).
