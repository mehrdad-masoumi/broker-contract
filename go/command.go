package notification

import "time"

// ContractVersion is the only supported version of notification.requested.
const ContractVersion = "v1"

// Channel wire values (JSON / RabbitMQ). Keep in sync with proto enum Channel
// and schemas/notification.requested.v1.json.
const (
	ChannelInApp    = "in_app"
	ChannelEmail    = "email"
	ChannelSMS      = "sms"
	ChannelWhatsApp = "whatsapp"
	ChannelPush     = "push"
)

// AllChannels lists every allowed channel in stable order.
var AllChannels = []string{
	ChannelInApp,
	ChannelEmail,
	ChannelSMS,
	ChannelWhatsApp,
	ChannelPush,
}

// Status values returned after accept (not after delivery).
const (
	StatusAccepted  = "accepted"
	StatusScheduled = "scheduled"
	StatusDuplicate = "duplicate"
)

// Routing constants for RabbitMQ ingress of notification.requested.v1.
const (
	CommandsExchange = "notification.commands"
	RequestedRouting = "notification.requested.v1"
	RequestedQueue   = "notification-service.requested.v1"
	DLX              = "notification.dlx"
	RequestedDLQ     = "notification-service.requested.v1.dlq"
	ContentType      = "application/json"
)

// Recipient is the caller-supplied contact snapshot.
// Notification Service must not look up users to fill these fields.
type Recipient struct {
	UserID       string   `json:"user_id,omitempty"`
	Email        string   `json:"email,omitempty"`
	Phone        string   `json:"phone,omitempty"`
	DeviceTokens []string `json:"device_tokens,omitempty"`
	DisplayName  string   `json:"display_name,omitempty"`
}

// NotificationRequested is the public versioned command/event.
// Semantic twin of proto notification.v1.NotificationRequested and
// schemas/notification.requested.v1.json.
type NotificationRequested struct {
	Version        string            `json:"version"`
	MessageID      string            `json:"message_id"`
	IdempotencyKey string            `json:"idempotency_key"`
	SourceService  string            `json:"source_service"`
	TemplateCode   string            `json:"template_code"`
	Locale         string            `json:"locale,omitempty"`
	Recipient      Recipient         `json:"recipient"`
	Channels       []string          `json:"channels"`
	Variables      map[string]any    `json:"variables,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	ScheduledAt    *time.Time        `json:"scheduled_at"`
	RequestedAt    *time.Time        `json:"requested_at,omitempty"`
	OccurredAt     *time.Time        `json:"occurred_at,omitempty"`
	CorrelationID  string            `json:"correlation_id,omitempty"`
	TraceID        string            `json:"trace_id,omitempty"`
}

// EffectiveRequestedAt returns requested_at, falling back to occurred_at.
func (n NotificationRequested) EffectiveRequestedAt() *time.Time {
	if n.RequestedAt != nil {
		return n.RequestedAt
	}
	return n.OccurredAt
}
