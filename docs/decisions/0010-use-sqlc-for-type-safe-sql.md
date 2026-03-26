---
status: accepted
date: 2026-01-01
---

> [한국어 버전](ko/0010-use-sqlc-for-type-safe-sql.md)

# ADR-0010: Use sqlc for Type-Safe SQL Code Generation

## Context and Problem Statement

Both SQLite and MariaDB storage backends must be supported. Writing raw SQL queries directly provides no type safety, and errors are only discovered at runtime.

## Considered Options

* **sqlc (SQL → Go code generation)** — write `query.sql` and `schema.sql`; sqlc automatically generates type-safe Go code
* **GORM (ORM)** — struct-tag-based ORM that abstracts SQL
* **Raw database/sql** — write SQL directly using the Go standard library

## Decision Outcome

Chosen option: "sqlc", because SQL is validated at compile time, generated Go code is type-safe, and SQL can be controlled directly without an ORM.

## Consequences

* Good, because SQL is validated at compile time, generated Go code is type-safe, and direct SQL control is maintained without an ORM.
* Bad, because `sqlc generate` must be re-run after changes to query.sql/schema.sql, and generated code in the `db/` directory must not be edited directly.
