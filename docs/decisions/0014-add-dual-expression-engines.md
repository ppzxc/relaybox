---
status: accepted
date: 2026-02-01
---

> [한국어 버전](ko/0014-add-dual-expression-engines.md)

# ADR-0014: Add CEL + Expr Dual Expression Engines

## Context and Problem Statement

In v0.2.0, an expression language became necessary for message filtering, mapping, and routing. Forcing a single engine removes user choice, and each engine has different strengths. CEL (Common Expression Language) is suited for strongly-typed complex expressions, while Expr provides a concise Go-native syntax.

## Considered Options

* **CEL only** — handle all expressions with Google Common Expression Language as the sole engine
* **Expr only** — handle all expressions with antonmedv/expr's Go-native syntax as the sole engine
* **Plugin registry (both CEL and Expr)** — manage expression engines via a registry; users choose the engine per input/output

## Decision Outcome

Chosen option: "Plugin registry (both CEL and Expr)", because forcing a single engine creates expressiveness limitations for certain use cases, and the registry pattern is consistent with hexagonal architecture while leaving the door open for future engine additions.

## Consequences

* Good, because users can select the engine per input/output, increasing expression flexibility.
* Good, because CEL is suited for strongly-typed complex expressions while Expr is suited for concise Go-native expressions.
* Good, because the registry pattern makes it easy to add new expression engines in the future.
* Bad, because syntax differences between the two engines may confuse users when both are used in the same rules file.
* Bad, because both engine dependencies are included, increasing binary size.
