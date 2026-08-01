# Runtime (planned)

A future execution layer for Quantara, written in Go.

## Status

Not started. This directory reserves the shape of the future runtime.

## Possible responsibilities

- **Sandbox environments** — isolated execution for testing contract deployments
  before they touch a real network.
- **Deployment workers** — background processes that carry out deployments
  asynchronously, instead of `quantara-core`'s current synchronous simulation.
- **Simulation engine** — dry-run a deployment and surface what would happen
  without actually deploying.
- **Monitoring agents** — watch deployed contracts and report activity back to
  `quantara-core`'s contract registry.

## Possible stack

- Docker SDK for Go, for sandboxing/worker isolation
- gRPC, if the runtime ends up needing to talk to `quantara-core` as a separate
  service rather than being invoked in-process

## Relationship to quantara-core

Today, `quantara-core`'s `POST /api/deploy` *simulates* a deployment synchronously.
The runtime is where that simulation would be replaced with something real, without
changing the API contract the dashboard and CLI depend on.
