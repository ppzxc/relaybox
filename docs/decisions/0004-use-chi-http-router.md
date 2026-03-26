---
status: accepted
date: 2026-01-01
---

> [한국어 버전](ko/0004-use-chi-http-router.md)

# ADR-0004: Use chi as HTTP Router

## Context and Problem Statement

HTTP handler registration, middleware chaining, and URL parameter extraction are required. Go stdlib `net/http` has limited pattern matching and requires extra work for URL parameter extraction. When adopting an external router, stdlib compatibility and maintenance activity must both be considered.

## Considered Options

* **go-chi/chi** — lightweight router fully compatible with stdlib `net/http`, middleware-centric design
* **gorilla/mux** — mature router, but archived in 2022 with no active maintenance
* **stdlib net/http (Go 1.22+)** — zero-dependency approach leveraging improved pattern matching introduced in Go 1.22

## Decision Outcome

Chosen option: "go-chi/chi", because full stdlib `net/http` compatibility makes adapter replacement easy, middleware chaining is clean, and active maintenance ensures long-term stability.

## Consequences

* Good, because full stdlib `net/http` compatibility, clean middleware chaining, and active maintenance ensure long-term stability.
* Bad, because it adds an external dependency (lightweight, but an additional module compared to stdlib).
