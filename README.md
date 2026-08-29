# Broker Contract

Shared, versioned contracts for Broker microservices.

This repository is the source of truth for Notification ingress (gRPC + RabbitMQ).
It has **no dependency** on Notification Service internals.

Module:

```text
github.com/mehrdad-masoumi/broker-contract
```

Future GitHub URL:

```text
https://github.com/mehrdad-masoumi/broker-contract
```

## Layout

```text
broker-contract/
  proto/notification/v1/notification.proto   # notification gRPC
  proto/user/v1/user.proto                   # user gRPC (+ ApplyKYCIdentity)
  proto/media/v1/media.proto                 # media typed S2S contracts
  gen/go/                                    # generated Go (do not edit)
  go/event/                                  # event envelopes + ownership catalog
  permissions/                               # shared RBAC catalog (no Service: core)
  schemas/                                   # RabbitMQ JSON Schema
  examples/
  buf.yaml
  buf.gen.yaml
  Makefile
```

## Media transport decision

`proto/media/v1/media.proto` defines typed RPCs for CreateUpload, CompleteUpload,
GetFileMetadata, CreateDownloadURL, ValidateOwnership, and DeleteFile.

**Active boundary remains HTTP** `/internal/media/*` (used by kyc-service and
other callers). Do not force a gRPC migration until media-service exposes a gRPC
listener and callers are switched deliberately.

## Versioning

- Additive changes stay in `v1`
- Breaking changes introduce `v2/` (proto + schema + routing key)
- Permission catalog: fine-grained dotted keys are canonical; colon keys are
  legacy aliases kept for JWT/role backward compatibility
- Event catalog: legacy `KYC_*` / `TICKET_*` / `BONUS_*` aliases are preserved
  until all consumers are verified migrated

## Scope boundary

This module contains **wire contracts only**. Ports, repositories, providers, and
delivery workers belong in Notification Service (or other service repos).

## Semantic contract: `notification.requested.v1`

Protobuf `NotificationRequested` and JSON Schema `notification.requested.v1.json`
are semantic twins. Callers must supply full recipient contacts — Notification
Service will **not** call User Service.

Required fields:

| Field | Notes |
|-------|--------|
| `version` | must be `"v1"` |
| `message_id` | unique per message |
| `idempotency_key` | **mandatory** across gRPC and Rabbit |
| `source_service` | caller name |
| `template_code` | DB template code |
| `recipient` | contacts snapshot |
| `channels` | `in_app` \| `email` \| `sms` \| `whatsapp` \| `push` |

Optional for legacy commands: `locale`, `variables`, `metadata`, `scheduled_at`, `requested_at` / `occurred_at`, `correlation_id`, `trace_id`.

Domain-transition commands also carry the canonical metadata set:
`event_id`, `event_type`, `event_version`, `producer`, `aggregate_type`,
`aggregate_id`, positive `aggregate_version`, and optional `causation_id` /
`operation_id`. The core set is atomic: either all core fields are absent for a
legacy command, or all are present and version-consistent. This lets
Notification Service durably fence replay and stale aggregate deliveries.

### Example JSON (RabbitMQ body)

```json
{
  "version": "v1",
  "message_id": "550e8400-e29b-41d4-a716-446655440000",
  "idempotency_key": "withdrawal:123:approved",
  "source_service": "withdrawal-service",
  "template_code": "withdrawal_approved",
  "locale": "fa-IR",
  "recipient": {
    "user_id": "123",
    "email": "user@example.com",
    "phone": "+989121234567",
    "display_name": "Mehrdad"
  },
  "channels": ["in_app", "email", "sms"],
  "variables": {
    "amount": "1000",
    "currency": "USDT"
  },
  "metadata": {
    "withdrawal_id": "123"
  },
  "scheduled_at": null,
  "requested_at": "2026-07-29T10:00:00Z"
}
```

## Choosing transport

### Use **gRPC** when

- You need an immediate **Accepted** response
- You need `notification_id` in the response
- The notification is sensitive (OTP / security)
- You want validation errors synchronously

### Use **RabbitMQ** when

- The flow is asynchronous
- You do not need a direct response
- The producer must not depend on Notification Service availability at call time
- Throughput / buffering matters more than sync ack

### Both transports

- Share the same semantic contract
- Require `idempotency_key`
- Produce the same durable outcome
- Enter the same Application Command Service inside Notification Service

### RabbitMQ topology

| Resource | Name |
|----------|------|
| Exchange | `notification.commands` |
| Routing key | `notification.requested.v1` |
| Queue | `notification-service.requested.v1` |
| DLX | `notification.dlx` |
| DLQ | `notification-service.requested.v1.dlq` |

Content-Type: `application/json` (body must validate against the JSON Schema).

### gRPC

```protobuf
service NotificationService {
  rpc SendNotification(SendNotificationRequest)
      returns (SendNotificationResponse);
}
```

Response includes `notification_id`, `status`, `duplicate`, `accepted_at`, optional `error_code` / `error_message`.
The RPC accepts after durable persist (notification + outbox); it does **not** wait for email/SMS delivery.

## Go usage

```go
import (
    notification "github.com/mehrdad-masoumi/broker-contract/go"
    notificationv1 "github.com/mehrdad-masoumi/broker-contract/gen/go/notification/v1"
)

// JSON / Rabbit
cmd, err := notification.ParseAndValidateJSON(body)

// gRPC request → shared command
cmd, err := notification.FromProto(req.GetNotification())
```

Constants for topology live in package `notification` (`CommandsExchange`, `RequestedRouting`, …).

## Generate

```bash
make deps
make proto
# or: make buf-generate   # if buf is installed
```

Requires `protoc`, `protoc-gen-go`, and `protoc-gen-go-grpc` on `PATH`.
On Windows you may point `PROTOC` at a local binary (see Makefile).
