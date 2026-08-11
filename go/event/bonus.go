package event

import "time"

// Bonus event types. Producer: bonus-service.
// Legacy UPPER_SNAKE aliases remain for dual-run consumers.
const (
	TypeBonusGrantedV1 = "bonus.granted.v1"
	TypeBonusRevokedV1 = "bonus.revoked.v1"

	LegacyBonusGranted = "BONUS_GRANTED"
	LegacyBonusRevoked = "BONUS_REVOKED"
)

// BonusGrantedV1 is a minimal grant signal (amount as decimal string).
type BonusGrantedV1 struct {
	BonusID    int64     `json:"bonus_id"`
	GrantID    int64     `json:"grant_id"`
	UserID     int64     `json:"user_id"`
	Amount     string    `json:"amount,omitempty"`
	Currency   string    `json:"currency,omitempty"`
	OccurredAt time.Time `json:"occurred_at"`
}

// BonusRevokedV1 is a minimal revoke signal.
type BonusRevokedV1 struct {
	BonusID    int64     `json:"bonus_id"`
	GrantID    int64     `json:"grant_id"`
	UserID     int64     `json:"user_id"`
	OccurredAt time.Time `json:"occurred_at"`
}
