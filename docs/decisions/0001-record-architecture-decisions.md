---
status: accepted
date: 2026-03-26
---

> [한국어 버전](ko/0001-record-architecture-decisions.md)

# ADR-0001: Record Architecture Decisions as ADRs

## Context and Problem Statement

Decisions exist only implicitly in code, making it difficult to understand the "why" behind each choice. AI tools like Claude Code cannot determine the rationale behind decisions when reading code without context. There is a risk of repeating the same discussions during future maintenance.

## Considered Options

* **MADR format** — Markdown Any Decision Records, a lightweight format with structured sections and front matter
* **Lightweight format (plain text)** — free-form text with no defined structure
* **No documentation** — leave decisions only in code and commit messages

## Decision Outcome

Chosen option: "MADR format", because the structured sections (Context, Options, Decision, Consequences) are optimized for Claude Code readability, allowing both AI tools and team members to quickly understand the rationale behind decisions.

## Consequences

* Good, because AI tools and team members can quickly grasp decision rationale, reducing repetitive discussions.
* Bad, because writing an ADR adds overhead whenever a new architecture decision is made.
