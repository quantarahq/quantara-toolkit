# quantara-toolkit

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Status: placeholder](https://img.shields.io/badge/status-placeholder-lightgrey.svg)](ROADMAP.md)

Future developer tooling for **Quantara** — an open-source developer infrastructure
platform for the Soroban smart contract ecosystem.

> **Status: placeholder, on purpose.** This repo exists to reserve and document the
> shape of the future Quantara tooling ecosystem — a CLI and a runtime — so
> contributors and grant reviewers can see where the project is headed without the
> working MVP repos being cluttered with unfinished code. **Nothing here is
> implemented.** For the real, working product, see
> [quantarahq/quantara-core](https://github.com/quantarahq/quantara-core) (backend +
> Soroban contract) and [quantarahq/quantara-web](https://github.com/quantarahq/quantara-web)
> (dashboard).

Part of the Quantara project:

| Repo | What it is |
|---|---|
| [quantara-core](https://github.com/quantarahq/quantara-core) | Backend API + Soroban contract — the working MVP |
| [quantara-web](https://github.com/quantarahq/quantara-web) | Next.js dashboard — the working MVP |
| **quantara-toolkit** (this repo) | Placeholder for a future CLI and runtime |

---

## Table of contents

- [Why this repo exists now, empty](#why-this-repo-exists-now-empty)
- [Planned layout](#planned-layout)
- [The CLI](#the-cli-cli)
- [The runtime](#the-runtime-runtime)
- [Relationship to quantara-core](#relationship-to-quantara-core)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [FAQ](#faq)
- [License](#license)

## Why this repo exists now, empty

The Quantara MVP (`quantara-core` + `quantara-web`) demonstrates the core developer
workflow — create a project, deploy, inspect the contract registry — through a web
dashboard talking to a REST API. That's a deliberate, narrow scope: per the
project's own philosophy, *"build a complete but minimal developer workflow instead
of incomplete versions of every possible feature."*

A CLI and an execution runtime are natural next steps once that foundation is
validated by real usage — but building them speculatively, before the API they'd
depend on has proven itself, would be exactly the kind of premature, half-built
feature sprawl the MVP philosophy exists to avoid. This repo is where that work will
land *when it's actually time*, and until then it serves a real purpose: it's a
public, versioned place to have the design conversation, so the eventual
implementation isn't designed from scratch under time pressure.

## Planned layout

```
quantara-toolkit/
├── cli/          Future Go CLI — quantara init / deploy / logs
│   └── README.md
├── runtime/      Future Go execution layer — sandboxes, deployment workers, simulation
│   └── README.md
├── examples/     Future example projects built with the CLI
│   └── README.md
├── ROADMAP.md    Milestone-by-milestone sequencing
├── README.md     you are here
├── CONTRIBUTING.md / CODE_OF_CONDUCT.md / SECURITY.md / CHANGELOG.md
└── LICENSE
```

## The CLI (`cli/`)

**Language:** Go. **Planned libraries:** [Cobra](https://github.com/spf13/cobra) for
command structure, [Viper](https://github.com/spf13/viper) for configuration
(config file + environment variables).

**Planned commands:**

```
quantara init       scaffold a new Quantara project locally
quantara deploy      deploy a contract via a running quantara-core instance
quantara logs        tail deployment/contract activity for a project
```

The CLI is meant to be **an alternative client to the same REST API `quantara-web`
already uses** — not a separate backend, not a shortcut around `quantara-core`. If
you can do it through the dashboard, you'll eventually be able to do it from the
terminal, against the exact same endpoints. See
[`cli/README.md`](cli/README.md) for more detail.

## The runtime (`runtime/`)

**Language:** Go. **Possible responsibilities:**

- **Sandbox environments** — isolated execution for testing deployments before they
  touch a real network.
- **Deployment workers** — background processes that carry out deployments
  asynchronously, replacing `quantara-core`'s current synchronous simulation.
- **Simulation engine** — dry-run a deployment and show what would happen without
  actually deploying.
- **Monitoring agents** — watch deployed contracts and report activity back into
  `quantara-core`'s contract registry.

**Possible stack:** Docker SDK for Go (worker isolation), gRPC (if the runtime ends
up needing to be a separate service rather than invoked in-process). See
[`runtime/README.md`](runtime/README.md).

## Relationship to `quantara-core`

Today, `POST /api/deploy` on `quantara-core` *simulates* a deployment synchronously
and returns immediately — there's no queue, no worker, no real network call. The
runtime described above is specifically where that simulation would eventually be
replaced with something real, **without changing the API contract** the dashboard
and CLI already depend on. That's the whole point of keeping this work in a
separate repo behind a stable interface, rather than growing it directly inside
`quantara-core`.

## Roadmap

See [ROADMAP.md](ROADMAP.md) for the full milestone breakdown. In short:

- **Milestone 0 (today):** `quantara-core` + `quantara-web` are a complete, working
  MVP; this repo is a placeholder.
- **Milestone 1:** a read-only CLI (`quantara logs`, project listing) — pure
  alternative client, no new backend work required.
- **Milestone 2:** a write-capable CLI (`quantara deploy`, `quantara init`), plus
  config file/env var support.
- **Milestone 3:** the runtime — real deployment workers behind the existing API.

Explicitly **not** on this roadmap: multi-chain support, a hosted/managed Quantara,
or anything resembling a general-purpose CI/CD platform. Different products, not
natural extensions of this one.

## Contributing

This repo isn't ready for implementation PRs yet — **the most useful contribution
right now is design discussion**, not code. Start with
[CONTRIBUTING.md](CONTRIBUTING.md). Open a
[design discussion issue](https://github.com/quantarahq/quantara-toolkit/issues/new/choose)
with thoughts on CLI ergonomics, config shape, or runtime architecture. If you want
to write code today, [quantara-core](https://github.com/quantarahq/quantara-core)
and [quantara-web](https://github.com/quantarahq/quantara-web) are the repos with
open `good-first-issue`s.

## FAQ

**Why reserve empty directories instead of just writing this in the README?**
Because `cli/`, `runtime/`, and `examples/` each have their own README describing
what's planned specifically for that piece, and because a real directory structure
makes the intended shape of the eventual monorepo-within-a-repo concrete and
linkable, rather than an abstract paragraph.

**Why Go for the CLI and runtime, when the backend is Java and the frontend is
TypeScript?**
A CLI distributed as a single static binary is a much better experience in Go than
in Java (no JVM startup cost, no bundling a runtime) or Node (no `node_modules` to
ship). The runtime's likely need for lightweight process/container orchestration
also fits Go's ecosystem (Docker SDK for Go, gRPC) better than the alternatives.

**When will this actually start being built?**
No committed timeline — see [ROADMAP.md](ROADMAP.md)'s framing. It starts being
real work once the MVP has real usage to design a CLI *against*, rather than
guessing at ergonomics in a vacuum.

## License

[MIT](LICENSE)
