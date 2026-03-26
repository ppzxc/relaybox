---
status: accepted
date: 2026-02-20
---

> [한국어 버전](ko/0018-use-input-id-as-routing-key.md)

# ADR-0018: Remove InputType Enum, Use Input ID as Routing Key

## Context and Problem Statement

The initial design used an InputType enum (HTTP, WEBSOCKET, TCP) as the routing key. However, it was impossible to distinguish multiple inputs of the same protocol, and identifying inputs by their config ID was more natural than by protocol. For example, when running two HTTP inputs, InputType alone could not route each one independently.

## Considered Options

* **Input ID (inputs[].id value from config file)** — identify each input with a unique string ID and use it as the routing key
* **Keep InputType enum** — continue using the protocol type (HTTP/WEBSOCKET/TCP) as the routing key
* **Composite key (type + id)** — combine the protocol type and ID to form the routing key

## Decision Outcome

Chosen option: "Input ID only", because multiple inputs of the same protocol can be distinguished, the config file and code remain consistent, and the meaning is clear when accessed as `data.input` in CEL expressions.

## Consequences

* Good, because multiple inputs of the same protocol are distinguishable, config and code are consistent, and `data.input` in CEL expressions is semantically clear.
* Bad, because removing the enum required code changes (refactored in commit ad946e8).
