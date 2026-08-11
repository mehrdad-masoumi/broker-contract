package event

import "time"

// Ingress event type / routing key constants (KYC domain).
// Producer: kyc-service. Legacy UPPER_SNAKE aliases remain for dual-run consumers.
const (
	TypeKYCSubmittedV1        = "kyc.submitted.v1"
	TypeKYCApprovedV1         = "kyc.approved.v1"
	TypeKYCRejectedV1         = "kyc.rejected.v1"
	TypeKYCMoreInfoRequiredV1 = "kyc.more_info_required.v1"
	TypeKYCStatusChangedV1    = "kyc.status_changed.v1"

	// Legacy dual-run aliases (do not remove until all consumers migrate).
	LegacyKYCSubmitted        = "KYC_SUBMITTED"
	LegacyKYCApproved         = "KYC_APPROVED"
	LegacyKYCRejected         = "KYC_REJECTED"
	LegacyKYCMoreInfoRequired = "KYC_MORE_INFO_REQUIRED"
)

// KYCStatusChangedV1 is the payload for kyc.status_changed.v1 (no document PII).
type KYCStatusChangedV1 struct {
	UserID         int64     `json:"user_id"`
	KYCID          int64     `json:"kyc_id"`
	PreviousStatus string    `json:"previous_status"`
	NewStatus      string    `json:"new_status"`
	ChangedAt      time.Time `json:"changed_at"`
}

// KYCSubmittedV1 is a minimal status event (no document / national-id PII).
type KYCSubmittedV1 struct {
	UserID    int64     `json:"user_id"`
	KYCID     int64     `json:"kyc_id"`
	OccurredAt time.Time `json:"occurred_at"`
}

// KYCApprovedV1 is a minimal status event (no document / national-id PII).
type KYCApprovedV1 struct {
	UserID     int64     `json:"user_id"`
	KYCID      int64     `json:"kyc_id"`
	OccurredAt time.Time `json:"occurred_at"`
}

// KYCRejectedV1 is a minimal status event (no document / national-id PII).
type KYCRejectedV1 struct {
	UserID     int64     `json:"user_id"`
	KYCID      int64     `json:"kyc_id"`
	OccurredAt time.Time `json:"occurred_at"`
}

// KYCMoreInfoRequiredV1 is a minimal status event (no document / national-id PII).
type KYCMoreInfoRequiredV1 struct {
	UserID     int64     `json:"user_id"`
	KYCID      int64     `json:"kyc_id"`
	OccurredAt time.Time `json:"occurred_at"`
}
