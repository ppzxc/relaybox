---
status: accepted
date: 2026-01-01
---

> [한국어 버전](ko/0012-use-rfc7807-error-responses.md)

# ADR-0012: Use RFC 7807 Problem Details Error Responses

## Context and Problem Statement

A consistent format is needed for communicating errors to clients over the HTTP API. Custom error formats differ across client parsing implementations. Various clients (webhook senders, management tools, etc.) must be able to handle errors in a uniform way.

## Considered Options

* **RFC 7807 Problem Details (application/problem+json)** — IETF standard error response format including type/title/status/detail fields
* **Custom JSON error format** — define a project-specific error struct and serialize it as JSON
* **HTTP status codes only** — express errors solely via HTTP status codes with no response body

## Decision Outcome

Chosen option: "RFC 7807 Problem Details (application/problem+json)", because adhering to the IETF standard lets clients predict the error structure without additional documentation, and the Content-Type header immediately identifies an error response.

## Consequences

* Good, because standard format ensures high client compatibility and reduces the burden of implementing custom error-parsing logic.
* Good, because `Content-Type: application/problem+json` clearly distinguishes error responses from normal responses.
* Good, because standardized type/title/status/detail fields consistently convey error type and detail.
* Bad, because RFC 7807 field names (type, title, etc.) may feel somewhat foreign relative to the domain language.
