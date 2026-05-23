# TelcoFlow: Order-to-Cash Walkthrough

This guide describes the end-to-end flow of a customer ordering a Fiber Broadband service and being billed for usage.

## 1. Product Discovery (Catalog Domain)
The customer browses the **Self-Care Portal**, which fetches available offers from the **Catalog Service** (TMF620).
- **Offering**: "HyperFiber 1Gbps"
- **Price**: $70/month + $0.01 per GB over 1TB.

## 2. Order Capture (CRM & Order Domains)
The customer submits an order. The **Order Management Service** (TMF622) receives the request.
- **Validation**: Order Service calls **CRM Service** to verify customer KYC and credit status.
- **Persistence**: The order is stored in PostgreSQL with state `Acknowledged`.

## 3. Orchestration (Workflow Domain)
Order Service starts a **Temporal Workflow**.
- **Decomposition**: The Product Order is split into a **Service Order** (Logical activation) and a **Resource Order** (Physical port allocation).
- **Inventory Reservation**: **Inventory Service** (TMF639) marks a specific OLT port and IP address as `Reserved`.

## 4. Provisioning (Production Domain)
The **Provisioning Service** executes a **Saga Workflow**.
- **Network Config**: The **Radius Adapter** configures the AAA server with the new user's credentials.
- **Hardware Config**: The **MikroTik/Router Adapter** pushes the PPPoE profile to the local BRAS.
- **Activation**: Once network elements confirm, Inventory Service updates the resource status to `Active`.

## 5. Usage & Mediation (Revenue Domain)
The customer starts using the internet.
- **CDR Generation**: The network router exports IPFIX/Netflow records.
- **Mediation**: The **Billing Service Usage Mediator** consumes these events from **Kafka**.
- **Rating**: The **Rating Engine** calculates the cost (e.g., if over 1TB, it applies the $0.01/GB tier).
- **Analytics**: Rated records are stored in **ClickHouse** for real-time dashboarding.

## 6. Assurance (Assurance Domain)
If a fiber cut occurs:
- **Alarm**: The OLT sends an SNMP trap to **Assurance Service** (TMF642).
- **Correlation**: The **Correlation Engine** identifies that this OLT serves the customer from Step 1.
- **Proactive AI**: The **AI Ops Agent** summarizes the impact and suggests a reroute or notifies the customer automatically via the Self-Care portal.
