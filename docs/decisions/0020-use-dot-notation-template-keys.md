---
status: accepted
date: 2026-02-10
---

> [한국어 버전](ko/0020-use-dot-notation-template-keys.md)

# ADR-0020: Use Dot-Notation Template Keys for Nested JSON

## Context and Problem Statement

When rendering webhook payloads, services like Slack and Discord require nested JSON structures. The goal is to express nested structures in the config file using the form `"content.text": "{{ .message }}"`. Having users write nested JSON directly makes config files complex and error-prone.

## Considered Options

* **Automatic dot-notation key conversion** — automatically convert keys like `"a.b.c"` into nested structures `{"a":{"b":{"c":val}}}`
* **Users write nested JSON directly** — describe the complete nested JSON in the config file
* **Full JSON via Go text/template** — render the entire payload using Go text/template

## Decision Outcome

Chosen option: "Automatic dot-notation key conversion", because Slack/Discord webhook formats can be expressed concisely, improving config file readability and reducing user errors.

## Consequences

* Good, because Slack/Discord webhook formats are expressible concisely, improving config file readability.
* Bad, because key names containing a literal dot (.) cannot be used, and the conversion logic introduces additional complexity.
