package event

import "time"

// Ingress event type / routing key constants (user domain).
const (
	TypeUserRegisteredV1 = "user.registered.v1"
	TypeUserDisabledV1   = "user.disabled.v1"
	TypeUserVerifiedV1   = "user.verified.v1"
	TypeUserUpdatedV1    = "user.updated.v1"
	TypeUserLoggedInV1   = "user.logged_in.v1"
)

// UserRegisteredV1 is the payload for user.registered.v1.
type UserRegisteredV1 struct {
	UserID       int64     `json:"user_id"`
	ReferralCode string    `json:"referral_code,omitempty"`
	RegisteredAt time.Time `json:"registered_at"`
	Country      string    `json:"country,omitempty"`
}

// UserDisabledV1 is the payload for user.disabled.v1.
type UserDisabledV1 struct {
	UserID     int64     `json:"user_id"`
	DisabledAt time.Time `json:"disabled_at"`
	Reason     string    `json:"reason,omitempty"`
}

// UserVerifiedV1 is the payload for user.verified.v1.
type UserVerifiedV1 struct {
	UserID     int64     `json:"user_id"`
	Email      string    `json:"email,omitempty"`
	Phone      string    `json:"phone,omitempty"`
	VerifiedAt time.Time `json:"verified_at"`
}

// UserUpdatedV1 is the payload for user.updated.v1.
type UserUpdatedV1 struct {
	UserID    int64     `json:"user_id"`
	Email     string    `json:"email,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Status    string    `json:"status,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserLoggedInV1 is the payload for user.logged_in.v1.
type UserLoggedInV1 struct {
	UserID    int64     `json:"user_id"`
	LoggedInAt time.Time `json:"logged_in_at"`
	IPAddress string    `json:"ip_address,omitempty"`
}
