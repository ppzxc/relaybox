---
status: accepted
date: 2026-01-01
---

> [한국어 버전](ko/0009-use-explicit-message-status-machine.md)

# ADR-0009: Use Explicit Message State Machine

## Context and Problem Statement

Messages follow the state transitions PENDING → DELIVERED or PENDING → FAILED → PENDING (requeue). Invalid transitions (e.g., DELIVERED → PENDING) must be prevented. A `CanTransitionTo()` method is added to `domain.MessageStatus` to explicitly define the allowed transitions: PENDING→DELIVERED, PENDING→FAILED, and FAILED→PENDING.

## Considered Options

* **Explicit CanTransitionTo() method** — add a transition-validity check method to the domain type
* **Convention only, no runtime check** — document allowed transitions and rely on developer discipline
* **DB-level constraint** — use CHECK constraints to prevent invalid state values from being stored

## Decision Outcome

Chosen option: "Explicit CanTransitionTo() method", because invalid transitions return `ErrInvalidTransition` for immediate failure, and the allowed transitions are documented directly in the domain code.

## Consequences

* Good, because invalid transitions return ErrInvalidTransition for fast failure, and allowed transitions are documented in domain code.
* Bad, because `CanTransitionTo()` must be updated whenever a new state is added.
