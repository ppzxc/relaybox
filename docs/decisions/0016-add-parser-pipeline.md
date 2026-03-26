---
status: accepted
date: 2026-02-01
---

> [한국어 버전](ko/0016-add-parser-pipeline.md)

# ADR-0016: Add Parser Pipeline with Graceful Degradation

## Context and Problem Statement

Input messages can arrive in a variety of formats: JSON, URL-encoded form, XML, logfmt, and others. CEL/Expr expressions need parsed structures to reference fields. Supporting only a single format makes it impossible to handle webhooks from diverse sources, undermining relaybox's role as a general-purpose relay hub.

## Considered Options

* **Parser pipeline (graceful degradation)** — try multiple parsers (JSON, Form, XML, Logfmt, Regex) in order; fall back to raw payload if all fail
* **JSON only** — implement only JSON parsing to keep things simple
* **Raw payload only** — expose the raw string to CEL expressions without any parsing

## Decision Outcome

Chosen option: "Parser pipeline (graceful degradation)", because it flexibly supports webhook formats from diverse sources (Slack, GitHub, Beszel, etc.) while preventing message loss by falling back to the raw payload on parse failure.

## Consequences

* Good, because JSON/Form/XML/Logfmt/Regex sources are all supported, and message loss is prevented via raw payload fallback.
* Bad, because each parser adapter must be maintained, and parsed data (ParsedData) is not persisted to the database — it is only available during processing.
