# Carrier-Grade Operational Runbook

## 1. Monitoring & Alerts
- **High Order Failure Rate**: Check `order-service` logs and Temporal workflow state. Often caused by CRM timeout or Catalog unavailability.
- **Provisioning Saga Failure**: Inspect Southbound adapters. Common issues: Network element unreachable or Radius secret mismatch.
- **Billing Latency**: Check ClickHouse disk space and Kafka partition lag.

## 2. Scaling
- Use `kubectl scale` for manual overrides or HPA for automated scaling based on CPU/Memory.
- ClickHouse: Add new shards and update `distributed` table definitions.

## 3. Incident Management
- Every Critical/Major alarm automatically triggers a TMF621 Trouble Ticket.
- AI Ops Agent provides automated RCA and impact summaries in the incident log.
