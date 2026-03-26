# Architecture Decision Records

> [한국어 버전](ko/README.md)

This directory records architecture decisions for the relaybox project in [MADR 4.0.0](https://adr.github.io/madr/) format.

## Index

| ADR | Title | Status | Date |
|-----|-------|--------|------|
| [0001](0001-record-architecture-decisions.md) | Record Architecture Decisions as ADRs | accepted | 2026-03-26 |
| [0002](0002-use-hexagonal-architecture.md) | Adopt Hexagonal Architecture | accepted | 2026-01-01 |
| [0003](0003-use-manual-dependency-injection.md) | Use Manual Dependency Injection (No DI Framework) | accepted | 2026-01-01 |
| [0004](0004-use-chi-http-router.md) | Use chi as HTTP Router | accepted | 2026-01-01 |
| [0005](0005-use-ulid-for-message-ids.md) | Use ULID for Message IDs | accepted | 2026-01-01 |
| [0006](0006-use-file-based-message-queue.md) | Use File-Based Message Queue | accepted | 2026-01-01 |
| [0007](0007-use-ack-nack-queue-semantics.md) | Use Ack/Nack Queue Consumption Pattern | accepted | 2026-01-01 |
| [0008](0008-use-string-enums-for-domain-types.md) | Use String-Based Domain Enums | accepted | 2026-01-01 |
| [0009](0009-use-explicit-message-status-machine.md) | Use Explicit Message State Machine | accepted | 2026-01-01 |
| [0010](0010-use-sqlc-for-type-safe-sql.md) | Use sqlc for Type-Safe SQL Code Generation | accepted | 2026-01-01 |
| [0011](0011-use-yaml-config-with-hot-reload.md) | Use Viper YAML Config with Hot Reload | accepted | 2026-01-01 |
| [0012](0012-use-rfc7807-error-responses.md) | Use RFC 7807 Problem Details Error Responses | accepted | 2026-01-01 |
| [0013](0013-use-api-version-header.md) | Use X-API-Version Response Header Instead of URL Versioning | accepted | 2026-01-01 |
| [0014](0014-add-dual-expression-engines.md) | Add CEL + Expr Dual Expression Engines | accepted | 2026-02-01 |
| [0015](0015-add-multi-protocol-input.md) | Add Multi-Protocol Input Support (HTTP, WebSocket, TCP) | accepted | 2026-02-01 |
| [0016](0016-add-parser-pipeline.md) | Add Parser Pipeline with Graceful Degradation | accepted | 2026-02-01 |
| [0017](0017-migrate-to-cgo-free-sqlite.md) | Migrate to CGO-Free SQLite (mattn → modernc) | accepted | 2026-02-15 |
| [0018](0018-use-input-id-as-routing-key.md) | Remove InputType Enum, Use Input ID as Routing Key | accepted | 2026-02-20 |
| [0019](0019-add-mariadb-storage-adapter.md) | Add MariaDB Storage Adapter | accepted | 2026-03-01 |
| [0020](0020-use-dot-notation-template-keys.md) | Use Dot-Notation Template Keys for Nested JSON | accepted | 2026-02-10 |

## Adding a New ADR

When a new architecture decision is made:

1. Create a file with the next number: `NNNN-title-with-dashes.md`
2. Use the template below:

```markdown
---
status: accepted
date: YYYY-MM-DD
---

> [한국어 버전](ko/NNNN-title-with-dashes.md)

# ADR-NNNN: Title

## Context and Problem Statement

2-3 sentences explaining why this decision was needed.

## Considered Options

* **Option A** — one-line description
* **Option B** — one-line description

## Decision Outcome

Chosen option: "Option A", because {key reason}.

## Consequences

* Good, because {positive impact}
* Bad, because {negative impact or trade-off}
```

3. Add an entry to the index table in this README
4. Also create a Korean translation in [`ko/`](ko/README.md)
