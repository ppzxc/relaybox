---
status: accepted
date: 2026-02-15
---

> [한국어 버전](ko/0017-migrate-to-cgo-free-sqlite.md)

> Supersedes: decision to use mattn/go-sqlite3 (v0.1.0)

# ADR-0017: Migrate to CGO-Free SQLite (mattn → modernc)

## Context and Problem Statement

The project initially used mattn/go-sqlite3, but its CGO dependency made cross-compilation (linux/arm64, darwin, windows) complex. CGO had to be eliminated to produce fully static binaries for distribution. Setting up platform-specific C toolchains in the CI pipeline was also a significant burden.

## Considered Options

* **modernc.org/sqlite (CGO-free pure Go implementation)** — use a pure Go SQLite implementation that runs without CGO
* **Keep mattn/go-sqlite3 + configure CGO cross-compilation** — set up per-platform cross-compilation toolchains in CI
* **Drop SQLite, support PostgreSQL only** — remove SQLite and switch to a single PostgreSQL storage backend

## Decision Outcome

Chosen option: "modernc.org/sqlite", because `CGO_ENABLED=0` allows cross-compilation for 6 platforms with a single Go toolchain, greatly reducing build complexity.

## Consequences

* Good, because `CGO_ENABLED=0` enables cross-compilation for 6 platforms, producing static binaries and simplifying Docker images.
* Bad, because performance may be marginally lower than mattn, and some SQLite extensions may not be supported.
