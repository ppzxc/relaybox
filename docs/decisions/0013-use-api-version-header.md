---
status: accepted
date: 2026-01-01
---

> [한국어 버전](ko/0013-use-api-version-header.md)

# ADR-0013: Use X-API-Version Response Header Instead of URL Versioning

## Context and Problem Statement

The API version must be communicated to clients. URL versioning (/v1/..., /v2/...) increases routing complexity and forces existing clients to update their URLs. A way to convey version information while keeping the URL structure simple is needed.

## Considered Options

* **X-API-Version response header** — include an X-API-Version header in all responses to inform clients of the current API version
* **URL path versioning (/v1/...)** — embed the version in the URL path for explicit per-version routing
* **Accept header versioning** — negotiate the version via the request Accept header (e.g., `application/vnd.relaybox.v1+json`)

## Decision Outcome

Chosen option: "X-API-Version response header", because it keeps the URL structure simple while allowing clients to easily check the current server version, and clients do not need to update URLs when the API version increments.

## Consequences

* Good, because the URL structure remains simple, keeping routing complexity low.
* Good, because existing clients do not need to update their URLs when the API version increases.
* Bad, because the version is only in the response header, so clients cannot negotiate a specific version in their request.
