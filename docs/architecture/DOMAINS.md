# TelcoFlow Domain Model

TelcoFlow is partitioned into 10 core domains based on TM Forum ODA Functional Blocks and Domain-Driven Design (DDD) principles.

## 1. IAM & Security (Identity Context)
- **Bounded Context**: Authentication, Authorization, Audit.
- **Key Entities**: User, Role, Permission, Tenant, AuditLog.
- **Standards**: OAuth2, OIDC, RBAC, ABAC.

## 2. CRM (Party & Customer Context)
- **Bounded Context**: Party Management, Customer Management, Account Management.
- **Key Entities**: Individual, Organization, Customer, BillingAccount, SettlementAccount.
- **TMF Alignment**: TMF632 (Party), TMF629 (Customer), TMF666 (Account).

## 3. Product Catalog (Product Context)
- **Bounded Context**: Catalog Management, Offering Management.
- **Key Entities**: ProductOffering, ProductSpecification, ProductPrice, Bundle.
- **TMF Alignment**: TMF620 (Product Catalog).

## 4. Order Management (Ordering Context)
- **Bounded Context**: Order Capture, Order Orchestration.
- **Key Entities**: ProductOrder, ServiceOrder, OrderItem, Milestone.
- **TMF Alignment**: TMF622 (Product Order), TMF641 (Service Order).

## 5. Billing & Charging (Revenue Context)
- **Bounded Context**: Rating, Invoicing, Mediation, Real-time Charging.
- **Key Entities**: CDR (Usage), Invoice, Payment, Balance, Wallet.
- **TMF Alignment**: TMF678 (Customer Bill), TMF676 (Payment).

## 6. Inventory (Resource & Service Context)
- **Bounded Context**: Service Inventory, Resource Inventory.
- **Key Entities**: ServiceInstance, ResourceInstance (Logical/Physical), SIM, IPAddress.
- **TMF Alignment**: TMF638 (Service Inventory), TMF639 (Resource Inventory).

## 7. Provisioning (Production Context)
- **Bounded Context**: Activation, Network Configuration.
- **Key Entities**: ProvisioningTask, Adapter, Command.

## 8. Assurance (Assurance Context)
- **Bounded Context**: Alarm Management, Trouble Ticketing, SLA Management.
- **Key Entities**: Alarm, TroubleTicket, SLAThreshold.
- **TMF Alignment**: TMF642 (Alarm), TMF621 (Trouble Ticket).

## 9. Observability (Insight Context)
- **Bounded Context**: Metrics, Logs, Traces.

## 10. AI Operations (Intelligence Context)
- **Bounded Context**: ML Inference, Agentic Workflow.
