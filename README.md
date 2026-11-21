<div align="center">

# Go Auth Service

**Modular authentication and authorization, written in Go.**

<p>
  <img src="https://img.shields.io/badge/Go-1.26.4+-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go 1.26.4+">
  <img src="https://img.shields.io/badge/status-in%20development-orange?style=flat-square" alt="Status: in development">
  <img src="https://img.shields.io/badge/API-HTTP%20%7C%20JSON-4B5563?style=flat-square" alt="HTTP JSON API">
</p>

</div>

> [!WARNING]
> This project is under active development. APIs, configuration, and storage support may change without notice.

## Overview

Go Auth Service is a modular authentication and authorization backend written in Go. It exposes a JSON HTTP API under `/api/v1` and is designed to keep business rules independent from transport and persistence choices.

## Tech Stack

- **Go 1.26.4+** — standard library `net/http`, no web framework
- **In-memory storage** — current development driver
- **PostgreSQL / SQLite** — planned persistence adapters (`pgx/v5`, `modernc.org/sqlite`)

## Getting Started

### Requirements

- Go 1.26.4 or newer

### Run locally

From the project root:

```bash
cp script/run-api.sh.example script/run-api.sh
chmod +x script/run-api.sh
./script/run-api.sh memory
```

The server listens on `http://localhost:8080`; API routes are under `/api/v1`.

Or start it directly:

```bash
DB_DRIVER=memory go run ./cmd/api
```

> `DB_DRIVER` is required. Without it, the service exits during configuration loading.

## Configuration

| Variable | Default | Description |
| --- | --- | --- |
| `APP_PORT` | `8080` | HTTP listen port |
| `DB_DRIVER` | — | Storage driver (`memory` is currently implemented) |
| `DB_SQLITE_PATH` | — | SQLite database path |
| `DB_POSTGRES_*` | — | PostgreSQL connection settings |

See [`script/run-api.sh.example`](script/run-api.sh.example) for the full environment layout.

> [!NOTE]
> Only the in-memory repository is implemented. SQLite and PostgreSQL support is not ready yet.

## Testing

Run the full suite:

```bash
go test ./...
```

Run a single test:

```bash
go test ./internal/infra/database/memory -run TestMemoryAuthorityRepository_Create -v
```

## Status

Under active development. APIs, configuration, and persistence support may change. See [`docs/project/PROGRESS.md`](docs/project/PROGRESS.md) for the current checklist.

## Contributing

Feedback and suggestions are welcome. This project follows [Conventional Commits](https://www.conventionalcommits.org/); keep commit messages short, imperative, and in English.

---

<p align="center">
  <sub>Built with Go.</sub>
</p>
