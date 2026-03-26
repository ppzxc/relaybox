---
status: accepted
date: 2026-01-01
---

> [한국어 버전](ko/0003-use-manual-dependency-injection.md)

# ADR-0003: Use Manual Dependency Injection (No DI Framework)

## Context and Problem Statement

Multiple adapters and services need to be wired together. The Go ecosystem offers DI frameworks such as Wire and Fx, but each comes with trade-offs — compile-time code generation and runtime reflection, respectively. In the early stages of the project the dependency graph is simple, so the cost-benefit of introducing a framework must be evaluated.

## Considered Options

* **Manual DI (cmd/server/main.go)** — directly instantiate and inject all adapters in the `runServer()` function
* **Wire (code-generation based)** — Google Wire automatically generates DI code at compile time
* **Fx (runtime reflection based)** — Uber Fx resolves dependencies at runtime via reflection

## Decision Outcome

Chosen option: "Manual DI (cmd/server/main.go)", because assembly order and dependencies are explicitly visible in `cmd/server/main.go` without external dependencies, and at the current scale the overhead of introducing a framework outweighs the benefits.

## Consequences

* Good, because no external dependencies are added, assembly order and relationships are explicit in code, easing debugging and onboarding.
* Bad, because `main.go` must be modified directly when adding new dependencies, and management burden grows as the dependency graph becomes more complex.
