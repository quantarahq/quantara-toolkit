# CLI

A command-line interface for Quantara, written in Go.

## Status

**Milestone 1 (read-only) in progress.** `projects list`, `projects get`, and
`logs` are implemented and talk to a running `quantara-core` instance the same
way `quantara-web` does. `init` scaffolds a project directory locally. `deploy`
(Milestone 2, a write operation) is not implemented yet.

## Build and run

```bash
cd cli
go build -o quantara .
./quantara --help
```

Requires a running `quantara-core` instance (see its
[infra/docker-compose.yml](https://github.com/quantarahq/quantara-core/blob/main/infra/docker-compose.yml)).
By default the CLI talks to `http://localhost:8080`; override with `--api-url`
or the `QUANTARA_API_URL` env var.

## Commands

```
quantara init [directory]        scaffold a local project (writes quantara.json)
quantara projects list           list all projects
quantara projects get <id>       show a single project
quantara logs <project-id>       show deployment + contract activity for a project
```

Planned, not yet implemented:

```
quantara deploy      deploy a contract via a running quantara-core instance (Milestone 2)
```

## Stack

- [Cobra](https://github.com/spf13/cobra) for command structure
- Talks to `quantara-core`'s REST API (`GET /api/projects`, `GET /api/projects/{id}`,
  `GET /api/projects/{id}/deployments`, `GET /api/projects/{id}/contracts`) — the
  CLI is an alternative client, not a separate backend. See
  [`internal/apiclient`](internal/apiclient) for the client implementation.
- Config file + env var support via Viper is planned for Milestone 2, alongside
  `deploy`; for now the API URL is a single `--api-url` flag / `QUANTARA_API_URL`
  env var.

## Why Go

Matches the rest of the "future tooling" layer (`runtime/`) and is a natural fit for
a single self-contained binary distributed to developer machines.
