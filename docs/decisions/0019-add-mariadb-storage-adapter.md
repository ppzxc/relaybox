---
status: accepted
date: 2026-03-01
---

> [한국어 버전](ko/0019-add-mariadb-storage-adapter.md)

# ADR-0019: Add MariaDB Storage Adapter

## Context and Problem Statement

SQLite is a single-file database well-suited for small deployments, but high-volume message environments or deployments that already use an external database require a server-based database such as MariaDB. The port/adapter structure of the hexagonal architecture already provided a foundation for swappable storage.

## Considered Options

* **Add MariaDB adapter** — implement a new storage adapter using a MariaDB/MySQL-compatible driver
* **Add PostgreSQL adapter** — implement an adapter using a PostgreSQL driver
* **SQLite only** — keep SQLite as the sole storage and address scaling through other means

## Decision Outcome

Chosen option: "Add MariaDB adapter", because it supports high-volume message environments, allows reuse of existing MariaDB infrastructure, and the storage factory pattern enables runtime switching between SQLite and MariaDB.

## Consequences

* Good, because high-volume environments are supported, existing MariaDB infrastructure can be reused, and the storage factory pattern allows runtime SQLite↔MariaDB switching.
* Bad, because operating a MariaDB server adds operational overhead, and both storage adapters must be kept in sync (schema changes require updates in both).
