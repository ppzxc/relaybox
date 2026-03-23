[English](README.md) | **한국어**

# relaybox

범용 릴레이 허브: 어떤 인바운드 프로토콜/포맷도 수신하고, CEL/Expr 표현식 기반 필터·변환·라우팅 규칙을 통해 아웃바운드 채널로 전달한다.

```
어떤 인바운드 (HTTP REST / WebSocket / TCP / ...)
        ↓
  파서 파이프라인 (JSON / Form / XML / Logfmt / Regex)
        ↓
  CEL / Expr 표현식 필터 + 변환 + 라우팅
        ↓
어떤 아웃바운드 (Webhook / ...)
```

## 주요 기능

- **멀티 프로토콜 인바운드** — HTTP REST, WebSocket, TCP
- **파서 파이프라인** — 입력별로 JSON, Form, XML, Logfmt, Regex(커스텀 패턴) 지원
- **표현식 기반 라우팅** — 입력별 CEL/Expr 필터, 매핑, 라우팅 규칙
- **at-least-once 전달** — 파일 큐 기반, 재시작 시에도 메시지 보존
- **지수 백오프 재시도** — 아웃풋별 `retryCount` / `retryDelayMs` 설정
- **설정 핫리로드** — 재시작 없이 아웃풋/규칙 변경 가능
- **Bearer 토큰 인증** — 입력별 독립 시크릿
- **Dot-notation 템플릿** — `parent.child` 키로 중첩 JSON 출력 생성

## 빠른 시작

### 사전 요구 사항

- Go 1.25+

### 빌드

```bash
# 클론 후 빌드
go build -o relaybox ./cmd/server/

# 설정 준비
cp docs/config.example.yaml config.yaml
# config.yaml 수정 후:

# 서버 시작
./relaybox start --config config.yaml
```

### Makefile

```bash
# 현재 플랫폼 빌드
make build

# 전체 플랫폼 크로스 컴파일 (dist/ 디렉토리에 출력)
make build-all

# 테스트 실행
make test

# 릴리스 빌드 (clean + build-all + checksums)
make release VERSION=1.0.0
```

## 설정

`config.yaml` 예시 (전체 레퍼런스는 `docs/config.example.yaml` 참고):

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
    engine: CEL          # 필수 — CEL 또는 EXPR
    parser: JSON         # JSON, FORM, XML, LOGFMT, REGEX
    secret: "change-me"
    rules:
      # Rule 1: 조건부 라우팅
      - filter: 'data.severity == "HIGH"'
        mapping:
          level: '"CRITICAL"'
        routing:
          - condition: 'data.level == "CRITICAL"'
            outputIds: [ops-webhook]
      # Rule 2: 필터 없이 전체 전송
      - outputIds: [notify-bot]

  - id: tcp-input
    type: GENERIC
    engine: CEL
    address: ":9001"
    delimiter: "\n"
    parser: JSON
    rules:
      - outputIds: [ops-webhook]   # 단순: 필터 없이 전체 전송

outputs:
  - id: ops-webhook
    type: WEBHOOK
    engine: CEL          # 필수 — CEL 또는 EXPR
    url: "https://hooks.example.com/xyz"
    template:
      text: 'data.input + ": " + data.payload'
    retryCount: 3
    retryDelayMs: 1000
    skipTLSVerify: false

  # Dot-notation 키는 중첩 JSON을 생성한다
  - id: notify-bot
    type: WEBHOOK
    engine: CEL
    url: "https://example.com/api/v1/bots/1/text"
    secret: "bearer-token"   # Authorization: Bearer <secret> 로 전송
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
  defaultRetryCount: 3      # 아웃풋에 retryCount 없을 때 폴백
  defaultRetryDelay: "1s"   # 폴백 기본 재시도 대기 (Go duration)
  pollBackoff: "500ms"      # 빈 큐 폴링 간격
```

### 표현식 변수

모든 표현식(필터, 매핑, 라우팅, 템플릿)은 동일한 `data` 컨텍스트를 공유한다:

| 변수 | 설명 |
|------|------|
| `data.id` | 메시지 ULID |
| `data.input` | 입력 타입 (`BESZEL`, `DOZZLE`, `GENERIC` 등) |
| `data.payload` | 원본 페이로드 문자열 |
| `data.createdAt` | 수신 타임스탬프 (RFC3339) |
| `data.<field>` | `mapping` 표현식으로 추가된 필드 |

**필터** — 불리언 표현식; `false`이면 메시지 드롭:
```yaml
filter: 'data.input == "BESZEL"'
```

**매핑** — 계산된 필드로 `data` 보강:
```yaml
mapping:
  severity: '"HIGH"'
  label: 'data.input + "-alert"'
```

**라우팅** — 조건부 아웃풋 선택 (매핑 이후 평가):
```yaml
routing:
  - condition: 'data.severity == "HIGH"'
    outputIds: [ops-webhook]
```

**템플릿** — 표현식으로 아웃풋 필드를 렌더링. Dot-notation 키는 중첩 JSON을 생성한다:
```yaml
template:
  text: 'data.input + ": " + data.payload'
  content.type: '"text"'
  content.text: 'data.payload'
```

## API

### 메시지 수신

```
POST /inputs/{inputId}/messages
Authorization: Bearer <secret>
Content-Type: application/json

{"host": "server1", "status": "down"}
```

응답 `201 Created`:
```json
{"id": "01J...", "status": "PENDING"}
```
헤더: `Location: /inputs/{inputId}/messages/{messageId}`

### WebSocket 인바운드

```
GET /inputs/{inputId}/messages/ws
Authorization: Bearer <secret>
```

연결 후 JSON 메시지를 전송하면 HTTP POST와 동일하게 처리된다.

### TCP 인바운드

설정한 `address`로 연결 후 개행(또는 커스텀 `delimiter`) 구분 메시지를 전송한다. 토큰 인증 없음 — 네트워크 정책으로 보안 적용.

### 헬스 체크

```
GET /healthz
→ 200 OK
```

### API 문서

```
GET /docs          → Redoc HTML UI
GET /docs/openapi  → OpenAPI 스펙 (JSON)
GET /docs/asyncapi → AsyncAPI 스펙 (JSON)
```

모든 HTTP 응답에는 `X-API-Version` 헤더가 포함된다.

## 아키텍처

헥사고날 아키텍처(Ports & Adapters). 의존성 방향은 항상 도메인을 향해 안쪽으로만 흐른다.

```
domain (0 deps)
  ↑
application/port/{input,output}  ← 인터페이스 정의
  ↑
application/service              ← 비즈니스 로직
  ↑
adapter/{input,output}           ← 외부 세계와 연결
  ↑
cmd/server/main.go               ← DI 조립, cobra CLI
```

| 경로 | 역할 |
|------|------|
| `internal/domain/` | 엔티티(`Message`, `Output`), 열거형(`InputType`, `MessageStatus`, `OutputType`), 센티넬 에러 |
| `internal/application/port/input/` | `ReceiveMessageUseCase` 인터페이스 |
| `internal/application/port/output/` | `MessageRepository`, `MessageQueue`, `OutputSender`, `OutputRegistry`, `RuleConfigReader` 인터페이스 |
| `internal/application/service/` | `MessageService`(Receive), `RelayWorker`(Start) |
| `internal/config/` | Viper 기반 YAML 로더, `InMemoryRuleConfigReader`, hot-reload(`Watch`) |
| `internal/adapter/input/http/` | chi 라우터, RFC 7807 에러, `X-API-Version` 헤더 미들웨어 |
| `internal/adapter/input/websocket/` | gorilla/websocket 인바운드 핸들러 |
| `internal/adapter/output/sqlite/` | sqlc 기반 SQLite 저장소 |
| `internal/adapter/output/filequeue/` | 파일 기반 at-least-once 큐 |
| `internal/adapter/output/webhook/` | HTTP Webhook 송신 |
| `cmd/server/` | cobra `start` 커맨드, 전체 DI 조립 |
| `test/e2e/` | 전체 흐름 E2E 테스트 |

## 릴리스

버전 태그를 푸시하면 GitHub Actions가 자동으로 모든 플랫폼 바이너리를 빌드하고 GitHub Release를 생성한다.

```bash
git tag 1.0.0
git push origin 1.0.0
```

지원 플랫폼: `linux/amd64`, `linux/arm64`, `darwin/amd64`, `darwin/arm64`, `windows/amd64`, `windows/arm64`

각 릴리스에는 `checksums.txt`에 SHA256 체크섬이 포함된다.

## 개발

```bash
# 전체 테스트 (race detector 포함)
go test -race ./... -timeout 60s

# 정적 분석
go vet ./...

# sqlc 코드 재생성 (SQL 변경 후)
cd internal/adapter/output/sqlite && sqlc generate

# 현재 플랫폼 빌드
make build

# 전체 플랫폼 크로스 컴파일
make build-all
```
