package main

// Example: publish notification.requested.v1 JSON to RabbitMQ topology constants.
// This is documentation-only; it does not connect to a broker.

import (
	"encoding/json"
	"fmt"
	"time"

	notification "github.com/mehrdad-masoumi/broker-contract/go"
)

func main() {
	now := time.Now().UTC()
	cmd := notification.NotificationRequested{
		Version:        notification.ContractVersion,
		MessageID:      "550e8400-e29b-41d4-a716-446655440000",
		IdempotencyKey: "withdrawal:123:approved",
		SourceService:  "withdrawal-service",
		TemplateCode:   "withdrawal_approved",
		Locale:         "fa-IR",
		Recipient: notification.Recipient{
			UserID:      "123",
			Email:       "user@example.com",
			Phone:       "+989121234567",
			DisplayName: "Mehrdad",
		},
		Channels: []string{
			notification.ChannelInApp,
			notification.ChannelEmail,
			notification.ChannelSMS,
		},
		Variables: map[string]any{
			"amount":   "1000",
			"currency": "USDT",
		},
		Metadata: map[string]string{
			"withdrawal_id": "123",
		},
		RequestedAt: &now,
	}
	body, err := json.Marshal(cmd)
	if err != nil {
		panic(err)
	}
	if err := notification.ValidateJSONSchema(body); err != nil {
		panic(err)
	}
	fmt.Printf("exchange=%s routing_key=%s queue=%s\n",
		notification.CommandsExchange,
		notification.RequestedRouting,
		notification.RequestedQueue,
	)
	fmt.Printf("body=%s\n", body)
}
