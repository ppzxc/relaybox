---
status: accepted
date: 2026-01-01
---

> [한국어 버전](ko/0011-use-yaml-config-with-hot-reload.md)

# ADR-0011: Use Viper YAML Config with Hot Reload

## Context and Problem Statement

Routing rules and output settings must be updatable without restarting the server. The configuration file format should be easy for humans to read and edit. Minimizing downtime when modifying rules in production is important.

## Considered Options

* **Viper + YAML + fsnotify WatchConfig()** — load YAML via Viper; detect file changes with fsnotify-based WatchConfig() and hot-reload
* **Environment variables only** — inject all configuration via environment variables; changes require a restart
* **Config file + manual restart** — keep a YAML file but reload changes only via server restart

## Decision Outcome

Chosen option: "Viper + YAML + fsnotify WatchConfig()", because rules/outputs may change frequently in production and must be applied without a restart, and Viper simultaneously supports default values, environment variable overrides, and robust unmarshalling.

## Consequences

* Good, because rules/outputs changes require no server restart, resulting in zero downtime.
* Good, because YAML format is highly readable and easy to edit directly.
* Good, because Viper's default values, environment variable overrides, and type validation can all be leveraged.
* Bad, because hot reload only applies to rules/outputs; server binding address, storage path, and other server/storage settings still require a restart.
* Bad, because configuration errors (malformed YAML, invalid CEL expressions, etc.) are detected at runtime and may impact production.
