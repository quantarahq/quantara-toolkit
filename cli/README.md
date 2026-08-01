# CLI (planned)

A command-line interface for Quantara, written in Go.

## Status

Not started. This directory reserves the shape of the future CLI.

## Planned commands

```
quantara init      scaffold a new Quantara project locally
quantara deploy     deploy a contract via a running quantara-core instance
quantara logs       tail deployment/contract activity for a project
```

## Planned stack

- [Cobra](https://github.com/spf13/cobra) for command structure
- [Viper](https://github.com/spf13/viper) for configuration (config file + env vars)
- Talks to `quantara-core`'s REST API (`POST /api/deploy`, `GET /api/projects`, etc.)
  the same way `quantara-web` does — the CLI is an alternative client, not a
  separate backend.

## Why Go

Matches the rest of the "future tooling" layer (`runtime/`) and is a natural fit for
a single self-contained binary distributed to developer machines.
