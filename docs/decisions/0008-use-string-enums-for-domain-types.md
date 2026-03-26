---
status: accepted
date: 2026-01-01
---

> [한국어 버전](ko/0008-use-string-enums-for-domain-types.md)

# ADR-0008: Use String-Based Domain Enums

## Context and Problem Statement

Domain enums such as `MessageStatus` (PENDING/DELIVERED/FAILED) and `OutputType` (WEBHOOK/SLACK/DISCORD) need to be defined. Values must be human-readable during JSON serialization and meaningful when stored in the database.

## Considered Options

* **type X string + uppercase constants** — declare `type MessageStatus string` and use string constants such as `"PENDING"`
* **type X int + iota** — integer-based enum declaration
* **protobuf-style enum** — code generated from a protobuf definition

## Decision Outcome

Chosen option: "type X string + uppercase constants (\"PENDING\", \"DELIVERED\", etc.)", because no custom `MarshalJSON` implementation is required for JSON/DB storage, and values are immediately readable in logs and the database.

## Consequences

* Good, because no custom MarshalJSON needed for JSON/DB storage, values are immediately readable in logs, and type safety is maintained.
* Bad, because string comparison is negligibly slower than integer comparison (insignificant in practice).
