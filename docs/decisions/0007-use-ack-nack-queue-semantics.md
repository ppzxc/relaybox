---
status: accepted
date: 2026-01-01
---

> [한국어 버전](ko/0007-use-ack-nack-queue-semantics.md)

# ADR-0007: Use Ack/Nack Queue Consumption Pattern

## Context and Problem Statement

Messages must be returned to the queue for retry when webhook delivery fails. Messages must not be lost if a consumer crashes during processing. The implementation renames `.json` to `.json.processing` on Dequeue; Ack deletes the `.processing` file; Nack renames it back to `.json`; on server startup, any remaining `.processing` files are restored to `.json`.

## Considered Options

* **Ack/Nack callback pattern** — return AckFunc/NackFunc alongside the message on Dequeue; the consumer explicitly reports the processing result
* **Simple Dequeue-Delete pattern** — delete immediately on Dequeue; re-enqueue on failure
* **Transaction-based** — atomic processing via database transactions

## Decision Outcome

Chosen option: "Ack/Nack callback pattern", because the message is preserved until processing completes, enabling automatic recovery after a crash, and the interface is conceptually compatible with message brokers such as RabbitMQ.

## Consequences

* Good, because at-least-once delivery is guaranteed, automatic recovery after consumer crash, and the interface is conceptually compatible with message brokers (RabbitMQ, etc.).
* Bad, because at-least-once semantics require consumers to implement idempotency.
