package notification

import (
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	notificationv1 "github.com/mehrdad-masoumi/broker-contract/gen/go/notification/v1"
)

// FromProto maps a protobuf NotificationRequested into the JSON-oriented command.
func FromProto(p *notificationv1.NotificationRequested) (NotificationRequested, error) {
	if p == nil {
		return NotificationRequested{}, fmt.Errorf("notification is nil")
	}
	out := NotificationRequested{
		Version:          p.GetVersion(),
		MessageID:        p.GetMessageId(),
		IdempotencyKey:   p.GetIdempotencyKey(),
		SourceService:    p.GetSourceService(),
		TemplateCode:     p.GetTemplateCode(),
		Locale:           p.GetLocale(),
		CorrelationID:    p.GetCorrelationId(),
		TraceID:          p.GetTraceId(),
		EventID:          p.GetEventId(),
		EventType:        p.GetEventType(),
		EventVersion:     p.GetEventVersion(),
		Producer:         p.GetProducer(),
		AggregateType:    p.GetAggregateType(),
		AggregateID:      p.GetAggregateId(),
		AggregateVersion: p.GetAggregateVersion(),
		CausationID:      p.GetCausationId(),
		OperationID:      p.GetOperationId(),
	}
	if r := p.GetRecipient(); r != nil {
		out.Recipient = Recipient{
			UserID:       r.GetUserId(),
			Email:        r.GetEmail(),
			Phone:        r.GetPhone(),
			DeviceTokens: append([]string(nil), r.GetDeviceTokens()...),
			DisplayName:  r.GetDisplayName(),
		}
	}
	for _, ch := range p.GetChannels() {
		wire := protoChannelEnumToWire(ch)
		if wire == "" {
			return NotificationRequested{}, fmt.Errorf("invalid channel enum: %v", ch)
		}
		out.Channels = append(out.Channels, wire)
	}
	if vars := p.GetVariables(); len(vars) > 0 {
		out.Variables = make(map[string]any, len(vars))
		for k, v := range vars {
			out.Variables[k] = v
		}
	}
	if meta := p.GetMetadata(); len(meta) > 0 {
		out.Metadata = make(map[string]string, len(meta))
		for k, v := range meta {
			out.Metadata[k] = v
		}
	}
	if ts := p.GetScheduledAt(); ts != nil {
		t := ts.AsTime().UTC()
		out.ScheduledAt = &t
	}
	if ts := p.GetRequestedAt(); ts != nil {
		t := ts.AsTime().UTC()
		out.RequestedAt = &t
	}
	if ts := p.GetOccurredAt(); ts != nil {
		t := ts.AsTime().UTC()
		out.OccurredAt = &t
	}
	return out, nil
}

// ToProto maps the JSON-oriented command into protobuf NotificationRequested.
func ToProto(n NotificationRequested) (*notificationv1.NotificationRequested, error) {
	p := &notificationv1.NotificationRequested{
		Version:          n.Version,
		MessageId:        n.MessageID,
		IdempotencyKey:   n.IdempotencyKey,
		SourceService:    n.SourceService,
		TemplateCode:     n.TemplateCode,
		Locale:           n.Locale,
		CorrelationId:    n.CorrelationID,
		TraceId:          n.TraceID,
		EventId:          n.EventID,
		EventType:        n.EventType,
		EventVersion:     n.EventVersion,
		Producer:         n.Producer,
		AggregateType:    n.AggregateType,
		AggregateId:      n.AggregateID,
		AggregateVersion: n.AggregateVersion,
		CausationId:      n.CausationID,
		OperationId:      n.OperationID,
		Recipient: &notificationv1.Recipient{
			UserId:       n.Recipient.UserID,
			Email:        n.Recipient.Email,
			Phone:        n.Recipient.Phone,
			DeviceTokens: append([]string(nil), n.Recipient.DeviceTokens...),
			DisplayName:  n.Recipient.DisplayName,
		},
	}
	for _, ch := range n.Channels {
		enum, err := wireChannelToProtoEnum(ch)
		if err != nil {
			return nil, err
		}
		p.Channels = append(p.Channels, enum)
	}
	if len(n.Variables) > 0 {
		p.Variables = make(map[string]string, len(n.Variables))
		for k, v := range n.Variables {
			p.Variables[k] = fmt.Sprint(v)
		}
	}
	if len(n.Metadata) > 0 {
		p.Metadata = make(map[string]string, len(n.Metadata))
		for k, v := range n.Metadata {
			p.Metadata[k] = v
		}
	}
	if n.ScheduledAt != nil {
		p.ScheduledAt = timestamppb.New(n.ScheduledAt.UTC())
	}
	if n.RequestedAt != nil {
		p.RequestedAt = timestamppb.New(n.RequestedAt.UTC())
	}
	if n.OccurredAt != nil {
		p.OccurredAt = timestamppb.New(n.OccurredAt.UTC())
	}
	return p, nil
}

func protoChannelEnumToWire(ch notificationv1.Channel) string {
	switch ch {
	case notificationv1.Channel_CHANNEL_IN_APP:
		return ChannelInApp
	case notificationv1.Channel_CHANNEL_EMAIL:
		return ChannelEmail
	case notificationv1.Channel_CHANNEL_SMS:
		return ChannelSMS
	case notificationv1.Channel_CHANNEL_WHATSAPP:
		return ChannelWhatsApp
	case notificationv1.Channel_CHANNEL_PUSH:
		return ChannelPush
	default:
		return ""
	}
}

func wireChannelToProtoEnum(wire string) (notificationv1.Channel, error) {
	switch NormalizeChannel(wire) {
	case ChannelInApp:
		return notificationv1.Channel_CHANNEL_IN_APP, nil
	case ChannelEmail:
		return notificationv1.Channel_CHANNEL_EMAIL, nil
	case ChannelSMS:
		return notificationv1.Channel_CHANNEL_SMS, nil
	case ChannelWhatsApp:
		return notificationv1.Channel_CHANNEL_WHATSAPP, nil
	case ChannelPush:
		return notificationv1.Channel_CHANNEL_PUSH, nil
	default:
		return notificationv1.Channel_CHANNEL_UNSPECIFIED, fmt.Errorf("invalid channel %q", wire)
	}
}

// NormalizeChannels returns trimmed lower-case unique valid channels.
func NormalizeChannels(in []string) []string {
	seen := map[string]struct{}{}
	out := make([]string, 0, len(in))
	for _, ch := range in {
		n := NormalizeChannel(ch)
		if n == "" {
			continue
		}
		if _, ok := seen[n]; ok {
			continue
		}
		seen[n] = struct{}{}
		out = append(out, n)
	}
	return out
}

// EnsureVersion sets version to v1 when empty.
func EnsureVersion(n *NotificationRequested) {
	if n == nil {
		return
	}
	if strings.TrimSpace(n.Version) == "" {
		n.Version = ContractVersion
	}
}

// NowUTC returns current UTC time pointer helper for examples/tests.
func NowUTC() *time.Time {
	t := time.Now().UTC()
	return &t
}
