package notification_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	notificationv1 "github.com/mehrdad-masoumi/broker-contract/gen/go/notification/v1"
	notification "github.com/mehrdad-masoumi/broker-contract/go"
)

func sampleJSON() []byte {
	return []byte(`{
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
}`)
}

func TestValidateJSONSchema_Sample(t *testing.T) {
	require.NoError(t, notification.ValidateJSONSchema(sampleJSON()))
}

func TestParseAndValidateJSON(t *testing.T) {
	cmd, err := notification.ParseAndValidateJSON(sampleJSON())
	require.NoError(t, err)
	require.Equal(t, "v1", cmd.Version)
	require.Equal(t, "withdrawal:123:approved", cmd.IdempotencyKey)
	require.Equal(t, []string{"in_app", "email", "sms"}, cmd.Channels)
}

func TestValidateNotificationRequested_MissingIdempotency(t *testing.T) {
	err := notification.ValidateNotificationRequested(notification.NotificationRequested{
		Version:       "v1",
		MessageID:     "m1",
		SourceService: "svc",
		TemplateCode:  "t",
		Recipient:     notification.Recipient{Email: "a@b.c"},
		Channels:      []string{"email"},
	})
	require.Error(t, err)
}

func TestValidateNotificationRequested_InAppRequiresUserID(t *testing.T) {
	err := notification.ValidateNotificationRequested(notification.NotificationRequested{
		Version:        "v1",
		MessageID:      "m1",
		IdempotencyKey: "k",
		SourceService:  "svc",
		TemplateCode:   "t",
		Recipient:      notification.Recipient{Email: "a@b.c"},
		Channels:       []string{"in_app"},
	})
	require.Error(t, err)
}

func TestProtoRoundTripCompatibility(t *testing.T) {
	cmd, err := notification.ParseAndValidateJSON(sampleJSON())
	require.NoError(t, err)

	pb, err := notification.ToProto(cmd)
	require.NoError(t, err)

	raw, err := proto.Marshal(pb)
	require.NoError(t, err)

	var decoded notificationv1.NotificationRequested
	require.NoError(t, proto.Unmarshal(raw, &decoded))

	back, err := notification.FromProto(&decoded)
	require.NoError(t, err)
	require.Equal(t, cmd.IdempotencyKey, back.IdempotencyKey)
	require.Equal(t, cmd.TemplateCode, back.TemplateCode)
	require.Equal(t, cmd.Channels, back.Channels)
	require.Equal(t, cmd.Recipient.Email, back.Recipient.Email)
}

func TestProtobufGenerationArtifactsExist(t *testing.T) {
	_, file, _, ok := runtime.Caller(0)
	require.True(t, ok)
	root := filepath.Join(filepath.Dir(file), "..")
	pb := filepath.Join(root, "gen", "go", "notification", "v1", "notification.pb.go")
	grpc := filepath.Join(root, "gen", "go", "notification", "v1", "notification_grpc.pb.go")
	_, err := os.Stat(pb)
	require.NoError(t, err)
	_, err = os.Stat(grpc)
	require.NoError(t, err)

	// Ensure service descriptor was generated.
	require.NotEmpty(t, notificationv1.NotificationService_ServiceDesc.ServiceName)
	require.Equal(t, "notification.v1.NotificationService", notificationv1.NotificationService_ServiceDesc.ServiceName)
}

func TestSchemaFileMatchesEmbedded(t *testing.T) {
	_, file, _, ok := runtime.Caller(0)
	require.True(t, ok)
	path := filepath.Join(filepath.Dir(file), "..", "schemas", "notification.requested.v1.json")
	raw, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Contains(t, string(raw), `"title": "notification.requested.v1"`)
	// Sample instance must validate against the on-disk schema.
	require.NoError(t, notification.ValidateJSONSchema(sampleJSON()))

	var doc map[string]any
	require.NoError(t, json.Unmarshal(sampleJSON(), &doc))
	require.Equal(t, "v1", doc["version"])
}

func TestScheduledAtAllowed(t *testing.T) {
	at := time.Date(2026, 8, 1, 12, 0, 0, 0, time.UTC)
	n := notification.NotificationRequested{
		Version:        "v1",
		MessageID:      "m1",
		IdempotencyKey: "k",
		SourceService:  "svc",
		TemplateCode:   "t",
		Recipient:      notification.Recipient{Email: "a@b.c"},
		Channels:       []string{"email"},
		ScheduledAt:    &at,
	}
	require.NoError(t, notification.ValidateNotificationRequested(n))
	pb, err := notification.ToProto(n)
	require.NoError(t, err)
	require.NotNil(t, pb.ScheduledAt)
}
