**English** | [한국어](README.ko.md)

# relaybox

A generic relay hub: receives any inbound protocol/format, applies CEL/Expr expression-based filter, transform, and routing rules, then forwards messages to outbound channels.

```
Any inbound (HTTP REST / WebSocket / TCP / ...)
        ↓
  Parser pipeline (JSON / Form / XML / Logfmt / Regex)
        ↓
  CEL / Expr expression filter + transform + routing
        ↓
Any outbound (Webhook / ...)
```

## Features

- **Multi-protocol inbound** — HTTP REST, WebSocket, TCP
- **Parser pipeline** — JSON, Form, XML, Logfmt, Regex (with custom pattern) per input
- **Expression-based routing** — per-input CEL/Expr filter, mapping, and routing rules
- **At-least-once delivery** — file-queue backed; messages survive restarts
- **Exponential backoff retry** — per-output `retryCount` / `retryDelayMs`
- **Config hot-reload** — change outputs and rules without restarting
- **Bearer token auth** — per-input independent secret
- **Dot-notation templates** — produce nested JSON output via `parent.child` keys

## Quick Start

### Prerequisites

- Go 1.25+

### Build

```bash
# Clone and build
go build -o relaybox ./cmd/server/

# Copy example config
cp docs/config.example.yaml config.yaml
# Edit config.yaml, then:

# Start server
./relaybox start --config config.yaml
```

### Makefile

```bash
# Build for current platform
make build

# Cross-compile all platforms (output to dist/)
make build-all

# Run tests
make test

# Full release build (clean + build-all + checksums)
make release VERSION=1.0.0
```

## Configuration

`config.yaml` example (see `docs/config.example.yaml` for full reference):

```yaml
server:
  port: 8080
  readTimeout: 30s
  writeTimeout: 30s
  tls:
    enabled: false
    certFile: ""
    keyFile: ""

log:
  level: INFO    # DEBUG, INFO, WARN, ERROR
  format: JSON   # JSON, TEXT

inputs:
  - id: beszel
    type: BESZEL
    engine: CEL          # required — CEL or EXPR
    parser: JSON         # JSON, FORM, XML, LOGFMT, REGEX
    secret: "change-me"
    rules:
      # Rule 1: conditional routing
      - filter: 'data.severity == "HIGH"'
        mapping:
          level: '"CRITICAL"'
        routing:
          - condition: 'data.level == "CRITICAL"'
            outputIds: [ops-webhook]
      # Rule 2: always forward (no filter)
      - outputIds: [notify-bot]

  - id: tcp-input
    type: GENERIC
    engine: CEL
    address: ":9001"
    delimiter: "\n"
    parser: JSON
    rules:
      - outputIds: [ops-webhook]   # simple: no filter, send all

outputs:
  - id: ops-webhook
    type: WEBHOOK
    engine: CEL          # required — CEL or EXPR
    url: "https://hooks.example.com/xyz"
    template:
      text: 'data.input + ": " + data.payload'
    retryCount: 3
    retryDelayMs: 1000
    skipTLSVerify: false

  # Dot-notation keys produce nested JSON
  - id: notify-bot
    type: WEBHOOK
    engine: CEL
    url: "https://example.com/api/v1/bots/1/text"
    secret: "bearer-token"   # sent as Authorization: Bearer <secret>
    template:
      content.type: '"text"'
      content.text: 'data.input + " alert: " + data.payload'
    retryCount: 3
    retryDelayMs: 1000
    timeoutSec: 10

storage:
  type: SQLITE
  path: "./data/relaybox.db"

queue:
  type: FILE
  path: "./data/queue"
  workerCount: 2

worker:
  defaultRetryCount: 3      # fallback when output has no retryCount
  defaultRetryDelay: "1s"   # fallback base retry delay (Go duration)
  pollBackoff: "500ms"      # sleep between empty-queue polls
```

### Expression Variables

All expressions (filter, mapping, routing, template) share the same `data` context:

| Variable | Description |
|----------|-------------|
| `data.id` | Message ULID |
| `data.input` | Input type (`BESZEL`, `DOZZLE`, `GENERIC`, etc.) |
| `data.payload` | Raw payload string |
| `data.createdAt` | Receive timestamp (RFC3339) |
| `data.<field>` | Fields added by `mapping` expressions |

**Filter** — boolean expression; `false` drops the message:
```yaml
filter: 'data.input == "BESZEL"'
```

**Mapping** — enrich `data` with computed fields:
```yaml
mapping:
  severity: '"HIGH"'
  label: 'data.input + "-alert"'
```

**Routing** — conditional output selection (evaluated after mapping):
```yaml
routing:
  - condition: 'data.severity == "HIGH"'
    outputIds: [ops-webhook]
```

**Template** — render output fields with expressions. Dot-notation keys generate nested JSON:
```yaml
template:
  text: 'data.input + ": " + data.payload'
  content.type: '"text"'
  content.text: 'data.payload'
```

## API

### Ingest Message

```
POST /inputs/{inputId}/messages
Authorization: Bearer <secret>
Content-Type: application/json

{"host": "server1", "status": "down"}
```

Response `201 Created`:
```json
{"id": "01J...", "status": "PENDING"}
```
Header: `Location: /inputs/{inputId}/messages/{messageId}`

### WebSocket Inbound

```
GET /inputs/{inputId}/messages/ws
Authorization: Bearer <secret>
```

Send JSON messages over the connection; handled identically to HTTP POST.

### TCP Inbound

Connect to the configured `address` and send newline-delimited (or custom `delimiter`) messages. No token auth — secure via network policy.

### Health Check

```
GET /healthz
→ 200 OK
```

### API Documentation

```
GET /docs          → Redoc HTML UI
GET /docs/openapi  → OpenAPI spec (JSON)
GET /docs/asyncapi → AsyncAPI spec (JSON)
```

All HTTP responses include an `X-API-Version` header.

## Architecture

Hexagonal Architecture (Ports & Adapters). Dependencies always flow inward toward the domain.

```
domain (0 deps)
  ↑
application/port/{input,output}  ← interface definitions
  ↑
application/service              ← business logic
  ↑
adapter/{input,output}           ← external world
  ↑
cmd/server/main.go               ← DI assembly, cobra CLI
```

| Path | Role |
|------|------|
| `internal/domain/` | Entities (`Message`, `Output`), enums (`InputType`, `MessageStatus`, `OutputType`), sentinel errors |
| `internal/application/port/input/` | `ReceiveMessageUseCase` interface |
| `internal/application/port/output/` | `MessageRepository`, `MessageQueue`, `OutputSender`, `OutputRegistry`, `RuleConfigReader` interfaces |
| `internal/application/service/` | `MessageService` (Receive), `RelayWorker` (Start) |
| `internal/config/` | Viper-based YAML loader, `InMemoryRuleConfigReader`, hot-reload (`Watch`) |
| `internal/adapter/input/http/` | chi router, RFC 7807 errors, `X-API-Version` middleware |
| `internal/adapter/input/websocket/` | gorilla/websocket inbound handler |
| `internal/adapter/output/sqlite/` | sqlc-based SQLite repository |
| `internal/adapter/output/filequeue/` | File-based at-least-once queue |
| `internal/adapter/output/webhook/` | HTTP Webhook sender |
| `cmd/server/` | cobra `start` command, full DI assembly |
| `test/e2e/` | End-to-end flow tests |

## Release

Push a version tag to trigger GitHub Actions — it builds all platform binaries and creates a GitHub Release automatically.

```bash
git tag 1.0.0
git push origin 1.0.0
```

Supported platforms: `linux/amd64`, `linux/arm64`, `darwin/amd64`, `darwin/arm64`, `windows/amd64`, `windows/arm64`

Each release includes SHA256 checksums in `checksums.txt`.

## Development

```bash
# Full test suite (with race detector)
go test -race ./... -timeout 60s

# Static analysis
go vet ./...

# Regenerate sqlc code (after SQL changes)
cd internal/adapter/output/sqlite && sqlc generate

# Build for current platform
make build

# Cross-compile all platforms
make build-all
```
