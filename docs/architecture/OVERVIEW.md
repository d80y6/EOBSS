# TelcoFlow Architecture Overview

## 1. System Context
TelcoFlow sits at the heart of a telecom operator's business and technical operations.

## 2. Logical Architecture (ODA Alignment)
The platform is aligned with TM Forum's Open Digital Architecture (ODA).

### Engagement Layer
- **Admin Portal**: Internal users.
- **Self-Care Portal**: B2C/B2B customers.

### Core Business Logic Layer
- **Party & Customer (CRM)**: TMF629/632.
- **Product Catalog**: TMF620.
- **Order Management**: TMF622.
- **Billing & Revenue**: TMF678.

### Production Layer
- **Service & Resource Inventory**: TMF639/640.
- **Provisioning**: Network activation.
- **Assurance**: TMF642.

## 3. Technology Stack
- Go, Java, Python
- PostgreSQL, Kafka, Redis, ClickHouse
- Kubernetes, Helm, Terraform
