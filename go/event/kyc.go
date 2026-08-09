package event

import "time"

// Ingress event type / routing key constants (KYC domain).
const (
	TypeKYCStatusChangedV1 = "kyc.status_changed.v1"
)

// KYCStatusChangedV1 is the payload for kyc.status_changed.v1.
type KYCStatusChangedV1 struct {
	UserID         int64     `json:"user_id"`
	KYCID          int64     `json:"kyc_id"`
	PreviousStatus string    `json:"previous_status"`
	NewStatus      string    `json:"new_status"`
	ChangedAt      time.Time `json:"changed_at"`
}
