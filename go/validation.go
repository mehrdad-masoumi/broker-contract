package notification

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// FieldError is a single validation failure keyed by JSON field path.
type FieldError struct {
	Field   string
	Message string
}

func (e FieldError) Error() string {
	return e.Field + ": " + e.Message
}

// ValidationError aggregates field errors for a NotificationRequested payload.
type ValidationError struct {
	Fields []FieldError
}

func (e *ValidationError) Error() string {
	if e == nil || len(e.Fields) == 0 {
		return "validation failed"
	}
	parts := make([]string, 0, len(e.Fields))
	for _, f := range e.Fields {
		parts = append(parts, f.Error())
	}
	return "validation failed: " + strings.Join(parts, "; ")
}

func (e *ValidationError) add(field, msg string) {
	e.Fields = append(e.Fields, FieldError{Field: field, Message: msg})
}

func (e *ValidationError) empty() bool {
	return e == nil || len(e.Fields) == 0
}

// ValidateNotificationRequested validates the public command without
// depending on Notification Service internals.
func ValidateNotificationRequested(n NotificationRequested) error {
	ve := &ValidationError{}

	if strings.TrimSpace(n.Version) == "" {
		ve.add("version", "required")
	} else if n.Version != ContractVersion {
		ve.add("version", fmt.Sprintf("must be %q", ContractVersion))
	}
	if strings.TrimSpace(n.MessageID) == "" {
		ve.add("message_id", "required")
	}
	if key := strings.TrimSpace(n.IdempotencyKey); key == "" {
		ve.add("idempotency_key", "required")
	} else if utf8.RuneCountInString(key) > 255 {
		ve.add("idempotency_key", "max length 255")
	}
	if strings.TrimSpace(n.SourceService) == "" {
		ve.add("source_service", "required")
	}
	if strings.TrimSpace(n.TemplateCode) == "" {
		ve.add("template_code", "required")
	}
	if len(n.Channels) == 0 {
		ve.add("channels", "required")
	} else {
		seen := map[string]struct{}{}
		for i, ch := range n.Channels {
			norm := NormalizeChannel(ch)
			if norm == "" {
				ve.add(fmt.Sprintf("channels[%d]", i), "invalid channel")
				continue
			}
			if _, ok := seen[norm]; ok {
				ve.add(fmt.Sprintf("channels[%d]", i), "duplicate channel")
				continue
			}
			seen[norm] = struct{}{}
		}
	}

	hasContact := n.Recipient.UserID != "" ||
		n.Recipient.Email != "" ||
		n.Recipient.Phone != "" ||
		len(n.Recipient.DeviceTokens) > 0
	if !hasContact {
		ve.add("recipient", "at least one of user_id, email, phone, device_tokens is required")
	}

	needsInApp := false
	needsEmail := false
	needsPhone := false
	needsPush := false
	for _, ch := range n.Channels {
		switch NormalizeChannel(ch) {
		case ChannelInApp:
			needsInApp = true
		case ChannelEmail:
			needsEmail = true
		case ChannelSMS, ChannelWhatsApp:
			needsPhone = true
		case ChannelPush:
			needsPush = true
		}
	}
	if needsInApp && strings.TrimSpace(n.Recipient.UserID) == "" {
		ve.add("recipient.user_id", "required when channels include in_app")
	}
	if needsEmail && strings.TrimSpace(n.Recipient.Email) == "" {
		ve.add("recipient.email", "required when channels include email")
	}
	if needsPhone && strings.TrimSpace(n.Recipient.Phone) == "" {
		ve.add("recipient.phone", "required when channels include sms or whatsapp")
	}
	if needsPush && len(n.Recipient.DeviceTokens) == 0 {
		ve.add("recipient.device_tokens", "required when channels include push")
	}

	if ve.empty() {
		return nil
	}
	return ve
}
