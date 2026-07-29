package notification

import "strings"

// IsValidChannel reports whether s is an allowed channel wire value.
func IsValidChannel(s string) bool {
	switch strings.TrimSpace(s) {
	case ChannelInApp, ChannelEmail, ChannelSMS, ChannelWhatsApp, ChannelPush:
		return true
	default:
		return false
	}
}

// NormalizeChannel lowercases and trims; returns empty if unknown.
func NormalizeChannel(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	if IsValidChannel(s) {
		return s
	}
	return ""
}

// ProtoChannelToWire maps proto enum numeric names / generated string names
// to JSON wire values. Accepts both "CHANNEL_EMAIL" style and "email".
func ProtoChannelToWire(name string) string {
	name = strings.TrimSpace(name)
	switch strings.ToUpper(name) {
	case "CHANNEL_IN_APP", "IN_APP":
		return ChannelInApp
	case "CHANNEL_EMAIL", "EMAIL":
		return ChannelEmail
	case "CHANNEL_SMS", "SMS":
		return ChannelSMS
	case "CHANNEL_WHATSAPP", "WHATSAPP":
		return ChannelWhatsApp
	case "CHANNEL_PUSH", "PUSH":
		return ChannelPush
	}
	if IsValidChannel(strings.ToLower(name)) {
		return strings.ToLower(name)
	}
	return ""
}

// WireChannelToProtoNumber returns the proto enum number for a wire channel.
// 0 means unspecified / unknown.
func WireChannelToProtoNumber(wire string) int32 {
	switch NormalizeChannel(wire) {
	case ChannelInApp:
		return 1
	case ChannelEmail:
		return 2
	case ChannelSMS:
		return 3
	case ChannelWhatsApp:
		return 4
	case ChannelPush:
		return 5
	default:
		return 0
	}
}
