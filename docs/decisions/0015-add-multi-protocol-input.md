---
status: accepted
date: 2026-02-01
---

> [한국어 버전](ko/0015-add-multi-protocol-input.md)

# ADR-0015: Add Multi-Protocol Input Support (HTTP, WebSocket, TCP)

## Context and Problem Statement

Messages must be received from a variety of sources (Beszel, Grafana, etc.). Some sources do not support HTTP POST and instead use WebSocket or TCP streams. Supporting only a single protocol leaves some sources unreachable, preventing relaybox from fulfilling its role as a general-purpose webhook relay hub.

## Considered Options

* **HTTP only** — receive all messages via the `POST /inputs/{inputId}/messages` endpoint
* **HTTP + WebSocket** — add a gorilla/websocket-based inbound WebSocket handler alongside HTTP
* **HTTP + WebSocket + TCP** — support all three protocols, covering TCP stream sources as well

## Decision Outcome

Chosen option: "HTTP + WebSocket + TCP (all three protocols)", because monitoring tools such as Beszel only support TCP or WebSocket output, and relaybox, aiming to be a general-purpose relay hub, must not depend on the protocol constraints of its sources.

## Consequences

* Good, because TCP/WebSocket output from monitoring tools like Beszel can be received directly.
* Good, because the hexagonal architecture's driving adapter pattern makes it easy to add new protocol adapters independently.
* Bad, because three separate driving adapters (HTTP, WebSocket, TCP) must each be maintained, increasing operational complexity.
* Bad, because TCP has no frame boundaries, requiring custom delimiter parsing and additional edge-case handling.
