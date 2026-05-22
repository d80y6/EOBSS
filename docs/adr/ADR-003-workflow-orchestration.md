# ADR-003: Workflow Orchestration with Temporal

## Status
Accepted

## Context
Telecom processes like Order Fulfillment, Service Provisioning, and Trouble Ticketing involve multiple steps, long-running activities, and complex error handling/retries across distributed systems.

## Decision
We will use **Temporal** as the primary workflow orchestration engine.

## Rationale
- **Resiliency**: Temporal handles retries, timeouts, and state persistence automatically.
- **Visibility**: Provides a clear view of the execution state of every order and provisioning task.
- **Scalability**: Can handle millions of concurrent workflows.
- **Developer Experience**: Workflows are written in standard code (Go/Java).

## Consequences
- Operational overhead of maintaining a Temporal cluster.
- All orchestration logic (e.g., TMF622 Order Management) will be implemented as Temporal Workflows.
