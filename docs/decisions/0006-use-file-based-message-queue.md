---
status: accepted
date: 2026-01-01
---

> [한국어 버전](ko/0006-use-file-based-message-queue.md)

# ADR-0006: Use File-Based Message Queue

## Context and Problem Statement

Messages received over HTTP must be queued and processed asynchronously rather than forwarded to webhooks immediately. The goal is to operate as a single binary without any external broker (Kafka, RabbitMQ).

## Considered Options

* **Filesystem-based queue (JSON files)** — store messages as JSON files on disk; a worker consumes them sequentially
* **Redis-based queue** — use Redis List as a queue for high-throughput processing
* **In-memory queue (channels)** — in-process queue using Go channels

## Decision Outcome

Chosen option: "Filesystem-based queue (JSON files)", because the system can be deployed as a single binary with no external dependencies, and unprocessed messages can be recovered after a process restart.

## Consequences

* Good, because no external dependencies, unprocessed messages are recoverable after restart (at-least-once delivery), and queue contents can be inspected directly as files.
* Bad, because throughput is limited by file I/O, and race conditions are possible on distributed filesystems such as NFS.
