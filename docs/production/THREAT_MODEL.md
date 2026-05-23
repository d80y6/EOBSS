# TelcoFlow Threat Model

## 1. Identity & Access (IAM)
- **Threat**: Unauthorized access to Customer PII.
- **Countermeasure**: OAuth2/OIDC via Keycloak, mTLS between services, and Vault-backed secret management.

## 2. Network & Infrastructure
- **Threat**: DDoS on Order/Provisioning endpoints.
- **Countermeasure**: Kubernetes Gateway API with rate limiting, Cloudflare/WAF, and network isolation via NetworkPolicies.

## 3. Data Persistence
- **Threat**: SQL Injection in CRM/Billing.
- **Countermeasure**: Use of GORM with parameterized queries, input validation in the handler layer, and least-privilege DB users.

## 4. Supply Chain
- **Threat**: Vulnerable dependencies in Go/Node.js modules.
- **Countermeasure**: Automated dependency scanning (Snyk/GitHub Actions), container image signing, and SBOM generation.
