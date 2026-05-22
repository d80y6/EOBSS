# Sequence Diagram: Order Fulfillment

```mermaid
sequence_diagram
    participant C as Customer (Self-Care)
    participant AP as API Gateway
    participant OM as Order Management
    participant PC as Product Catalog
    participant CRM as CRM Service
    participant T as Temporal Workflow
    participant I as Inventory
    participant P as Provisioning
    participant B as Billing

    C->>AP: Submit Product Order
    AP->>OM: POST /productOrder
    OM->>CRM: Validate Customer Status
    OM->>PC: Validate Product Offering
    OM->>T: Start OrderWorkflow
    T->>I: Reserve Resources (IP, Port)
    T->>P: Activate Service (Radius/Router)
    P-->>T: Activation Success
    T->>I: Update Service Status (Active)
    T->>B: Trigger Billing Activation
    T-->>OM: Workflow Completed
    OM-->>C: Order Completed Notification
```
