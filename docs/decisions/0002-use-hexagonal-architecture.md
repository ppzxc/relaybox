---
status: accepted
date: 2026-01-01
---

> [한국어 버전](ko/0002-use-hexagonal-architecture.md)

# ADR-0002: Adopt Hexagonal Architecture

## Context and Problem Statement

The webhook relay hub must support interchangeable inputs (HTTP, WebSocket, TCP), outputs (Webhook, Slack, etc.), and storage backends (SQLite, MariaDB). When external dependencies penetrate domain logic, testing and adapter replacement become difficult. Long-term, external system changes should not affect business logic, requiring proper isolation.

## Considered Options

* **Hexagonal Architecture (Ports & Adapters)** — define input/output ports as interfaces around the domain, implement via adapters
* **Layered Architecture** — traditional vertical layers: Controller → Service → Repository
* **No explicit architecture** — split packages as needed without a defined structure

## Decision Outcome

Chosen option: "Hexagonal Architecture (Ports & Adapters)", because domain logic is fully isolated from external dependencies, making adapter swaps (SQLite→MariaDB, file queue→external queue) straightforward, and domain logic can be tested independently via port interfaces.

## Consequences

* Good, because adapters are easily swappable (SQLite→MariaDB, file queue→external queue), domain logic can be tested independently, and business logic is protected from external changes.
* Bad, because interface definitions add overhead and directory structure complexity increases.
