package main

// Example: build a gRPC SendNotificationRequest from the shared command.
// This is documentation-only; it does not dial a server.

import (
	"fmt"
	"time"

	notificationv1 "github.com/mehrdad-masoumi/broker-contract/gen/go/notification/v1"
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
	if err := notification.ValidateNotificationRequested(cmd); err != nil {
		panic(err)
	}
	pb, err := notification.ToProto(cmd)
	if err != nil {
		panic(err)
	}
	req := &notificationv1.SendNotificationRequest{Notification: pb}
	fmt.Printf("rpc=%s template=%s idempotency=%s\n",
		notificationv1.NotificationService_ServiceDesc.Methods[0].MethodName,
		req.Notification.TemplateCode,
		req.Notification.IdempotencyKey,
	)
}
