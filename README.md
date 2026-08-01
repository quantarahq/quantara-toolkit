# quantara-toolkit

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Future developer tooling for **Quantara** — an open-source developer infrastructure
platform for the Soroban smart contract ecosystem.

> **Status: placeholder.** This repo exists to reserve the shape of the future Quantara
> tooling ecosystem so contributors and grant reviewers can see where the project is
> headed. Nothing here is implemented yet — see
> [quantarahq/quantara-core](https://github.com/quantarahq/quantara-core) and
> [quantarahq/quantara-web](https://github.com/quantarahq/quantara-web) for the working
> MVP (backend API, Soroban contract, and dashboard).

## Why this repo exists now, empty

The Quantara MVP (`quantara-core` + `quantara-web`) demonstrates the core developer
workflow — create a project, deploy, inspect the contract registry — through a web
dashboard talking to a REST API. A CLI and a runtime are natural next steps once that
foundation is validated, but building them now would be scope creep: per the project's
MVP philosophy, a complete-but-minimal workflow beats incomplete versions of every
possible feature. This repo is the place that work will land.

## Planned layout

```
quantara-toolkit/
├── cli/        Future Go CLI — quantara init / deploy / logs
├── runtime/    Future Go execution layer — sandboxes, deployment workers, simulation
├── examples/   Future example projects built with the CLI
└── README.md
```

See each subdirectory's README for what's planned there, and [ROADMAP.md](ROADMAP.md)
for how this would be sequenced into milestones.

## Roadmap

### CLI (`cli/`)

Language: Go (planned libraries: [Cobra](https://github.com/spf13/cobra) for commands,
[Viper](https://github.com/spf13/viper) for configuration).

Planned commands:

```
quantara init      scaffold a new Quantara project
quantara deploy     deploy a contract via quantara-core
quantara logs       tail deployment/contract activity
```

### Runtime (`runtime/`)

Language: Go.

Planned responsibilities: sandboxed execution environments, deployment worker
processes, a simulation engine for testing deployments before they hit
`quantara-core`, and monitoring agents.

## Contributing

This repo isn't ready for implementation PRs yet — the most useful contribution right
now is discussion. Open an issue on this repo (or on
[quantara-core](https://github.com/quantarahq/quantara-core) if it's about the MVP
itself) if you have thoughts on CLI ergonomics or the runtime design. See
[CONTRIBUTING.md](CONTRIBUTING.md).

## License

[MIT](LICENSE)
