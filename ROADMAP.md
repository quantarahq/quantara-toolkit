# Roadmap

This complements [README.md](README.md)'s architecture sketch with a rough sequencing
of how this repo would go from placeholder to real tooling. None of this is committed
to a timeline — it's here so contributors and grant reviewers can see the intended
path, and so design discussion issues have something concrete to react to.

## Milestone 0 — today

- `quantara-core` (backend + Soroban contract) and `quantara-web` (dashboard) form a
  complete, working MVP.
- This repo is a placeholder describing where CLI/runtime work will eventually land.

## Milestone 1 — CLI, read-only

- `quantara init` scaffolds a local project directory.
- `quantara logs` / a `quantara projects list`-style command talks to a running
  `quantara-core` instance read-only, using the same REST API `quantara-web` uses.
- No new backend functionality required — this is purely an alternative client.

## Milestone 2 — CLI, write operations

- `quantara deploy` triggers `POST /api/deploy` from the command line.
- Config file + env var support (Viper) for pointing at different `quantara-core`
  instances (local, staging, etc.).

## Milestone 3 — runtime

- Replace `quantara-core`'s synchronous simulated deploy with an actual worker
  process from `runtime/`, without changing the public API contract.
- Sandbox/simulation execution before a real deploy.

## Explicitly out of scope for this roadmap

Multi-chain support, a hosted/managed version of Quantara, and anything resembling a
general-purpose CI/CD platform — those would be different products, not natural
extensions of this one.
