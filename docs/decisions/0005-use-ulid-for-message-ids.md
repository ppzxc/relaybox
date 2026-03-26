---
status: accepted
date: 2026-01-01
---

> [한국어 버전](ko/0005-use-ulid-for-message-ids.md)

# ADR-0005: Use ULID for Message IDs

## Context and Problem Statement

Messages need a unique identifier that is safe to expose in URLs. IDs must be generatable at the application layer without a database round-trip, and must support collision-free distributed generation at high message volumes.

## Considered Options

* **ULID (oklog/ulid)** — 128-bit URL-safe ID combining a 48-bit timestamp and 80-bit random component
* **UUID v4 (random)** — fully random 128-bit ID, the industry standard
* **Auto-increment integer** — sequential integer ID issued by the database

## Decision Outcome

Chosen option: "ULID", because time-ordered sorting improves index efficiency when querying messages, the URL-safe character set allows direct use in URLs without additional encoding, and IDs can be generated at the application layer without a database call, making them suitable for distributed environments.

## Consequences

* Good, because time-sortable (lexicographic), URL-safe, generatable without a database, and with an extremely low collision probability.
* Bad, because ULID is less widely known than UUID and may require discussion when integrating with external systems.
