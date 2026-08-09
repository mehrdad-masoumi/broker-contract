package event

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// EventVersionV1 is the current contract version embedded in event type suffixes.
const EventVersionV1 = "v1"

// RabbitMQ topology for platform domain events.
const (
	EventsExchange = "broker.events"
	EventsDLX      = "broker.events.dlx"
	ContentType    = "application/json"
)

// Envelope is the shared wire wrapper for all broker.events messages.
// Routing key equals EventType (e.g. trade.closed.v1).
type Envelope struct {
	EventID       string          `json:"event_id"`
	EventType     string          `json:"event_type"`
	EventVersion  string          `json:"event_version"`
	OccurredAt    time.Time       `json:"occurred_at"`
	Producer      string          `json:"producer"`
	CorrelationID string          `json:"correlation_id,omitempty"`
	Payload       json.RawMessage `json:"payload"`
}

// NewEnvelope builds an Envelope with a fresh event_id, occurred_at=now,
// event_version derived from eventType (suffix after last ".") or EventVersionV1,
// and payload marshaled as JSON.
func NewEnvelope(eventType, producer, correlationID string, payload any) (Envelope, error) {
	raw, err := json.Marshal(payload)
	if err != nil {
		return Envelope{}, fmt.Errorf("marshal payload: %w", err)
	}

	id, err := newEventID()
	if err != nil {
		return Envelope{}, err
	}

	return Envelope{
		EventID:       id,
		EventType:     eventType,
		EventVersion:  versionFromEventType(eventType),
		OccurredAt:    time.Now().UTC(),
		Producer:      producer,
		CorrelationID: correlationID,
		Payload:       raw,
	}, nil
}

// MarshalEnvelope serializes an Envelope to JSON bytes for RabbitMQ publishing.
func MarshalEnvelope(env Envelope) ([]byte, error) {
	return json.Marshal(env)
}

func versionFromEventType(eventType string) string {
	if i := strings.LastIndex(eventType, "."); i >= 0 && i+1 < len(eventType) {
		return eventType[i+1:]
	}
	return EventVersionV1
}

func newEventID() (string, error) {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", fmt.Errorf("generate event_id: %w", err)
	}
	// UUID v4 variant bits.
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		b[0], b[1], b[2], b[3],
		b[4], b[5],
		b[6], b[7],
		b[8], b[9],
		b[10], b[11], b[12], b[13], b[14], b[15]), nil
}
