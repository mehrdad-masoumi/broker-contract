package event

import "time"

// Egress event type / routing key constants (IB domain).
const (
	TypeIBApplicationSubmittedV1 = "ib.application_submitted.v1"
	TypeIBApplicationApprovedV1  = "ib.application_approved.v1"
	TypeIBApplicationRejectedV1  = "ib.application_rejected.v1"

	TypeIBRankChangedV1 = "ib.rank_changed.v1"
	TypeIBSuspendedV1   = "ib.suspended.v1"
	TypeIBActivatedV1   = "ib.activated.v1"

	TypeIBReferralCreatedV1 = "ib.referral_created.v1"
	TypeIBUserAttributedV1  = "ib.user_attributed.v1"

	TypeIBRebateCreatedV1  = "ib.rebate_created.v1"
	TypeIBRebateReleasedV1 = "ib.rebate_released.v1"
	TypeIBRebateReversedV1 = "ib.rebate_reversed.v1"

	TypeIBRebateTransferRequestedV1 = "ib.rebate_transfer_requested.v1"
	TypeIBRebateTransferCompletedV1 = "ib.rebate_transfer_completed.v1"
	TypeIBRebateTransferFailedV1    = "ib.rebate_transfer_failed.v1"
)

// IBApplicationSubmittedV1 is the payload for ib.application_submitted.v1.
type IBApplicationSubmittedV1 struct {
	ApplicationID int64     `json:"application_id"`
	UserID        int64     `json:"user_id"`
	SubmittedAt   time.Time `json:"submitted_at"`
}

// IBApplicationApprovedV1 is the payload for ib.application_approved.v1.
type IBApplicationApprovedV1 struct {
	ApplicationID int64     `json:"application_id"`
	UserID        int64     `json:"user_id"`
	IBID          int64     `json:"ib_id"`
	RankID        int64     `json:"rank_id"`
	ApprovedBy    int64     `json:"approved_by,omitempty"`
	ApprovedAt    time.Time `json:"approved_at"`
}

// IBApplicationRejectedV1 is the payload for ib.application_rejected.v1.
type IBApplicationRejectedV1 struct {
	ApplicationID int64     `json:"application_id"`
	UserID        int64     `json:"user_id"`
	RejectedBy    int64     `json:"rejected_by,omitempty"`
	Reason        string    `json:"reason,omitempty"`
	RejectedAt    time.Time `json:"rejected_at"`
}

// IBRankChangedV1 is the payload for ib.rank_changed.v1.
type IBRankChangedV1 struct {
	IBID       int64     `json:"ib_id"`
	UserID     int64     `json:"user_id"`
	OldRankID  int64     `json:"old_rank_id"`
	NewRankID  int64     `json:"new_rank_id"`
	ChangedBy  int64     `json:"changed_by,omitempty"`
	ChangedAt  time.Time `json:"changed_at"`
	Reason     string    `json:"reason,omitempty"`
}

// IBSuspendedV1 is the payload for ib.suspended.v1.
type IBSuspendedV1 struct {
	IBID        int64     `json:"ib_id"`
	UserID      int64     `json:"user_id"`
	SuspendedBy int64     `json:"suspended_by,omitempty"`
	Reason      string    `json:"reason,omitempty"`
	SuspendedAt time.Time `json:"suspended_at"`
}

// IBActivatedV1 is the payload for ib.activated.v1.
type IBActivatedV1 struct {
	IBID        int64     `json:"ib_id"`
	UserID      int64     `json:"user_id"`
	ActivatedBy int64     `json:"activated_by,omitempty"`
	ActivatedAt time.Time `json:"activated_at"`
}

// IBReferralCreatedV1 is the payload for ib.referral_created.v1 (campaign created).
type IBReferralCreatedV1 struct {
	CampaignID   int64     `json:"campaign_id"`
	IBID         int64     `json:"ib_id"`
	UserID       int64     `json:"user_id"`
	ReferralCode string    `json:"referral_code"`
	CreatedAt    time.Time `json:"created_at"`
}

// IBUserAttributedV1 is the payload for ib.user_attributed.v1.
type IBUserAttributedV1 struct {
	UserID       int64     `json:"user_id"`
	IBID         int64     `json:"ib_id"`
	RootIBID     int64     `json:"root_ib_id"`
	CampaignID   int64     `json:"campaign_id"`
	ReferralID   int64     `json:"referral_id"`
	ReferralCode string    `json:"referral_code"`
	Depth        int       `json:"depth"`
	AttributedAt time.Time `json:"attributed_at"`
}

// IBRebateCreatedV1 is the payload for ib.rebate_created.v1.
// Amount / lot fields are decimal strings (never JSON floats).
type IBRebateCreatedV1 struct {
	RebateID        int64     `json:"rebate_id"`
	TradeID         int64     `json:"trade_id"`
	UserID          int64     `json:"user_id"`
	BeneficiaryIBID int64     `json:"beneficiary_ib_id"`
	RootIBID        int64     `json:"root_ib_id"`
	ReferralID      int64     `json:"referral_id,omitempty"`
	RebateType      string    `json:"rebate_type"`
	Lot             string    `json:"lot"`
	GrossAmount     string    `json:"gross_amount"`
	Currency        string    `json:"currency"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
}

// IBRebateReleasedV1 is the payload for ib.rebate_released.v1
// (pending → available settlement).
type IBRebateReleasedV1 struct {
	RebateID        int64     `json:"rebate_id"`
	BeneficiaryIBID int64     `json:"beneficiary_ib_id"`
	Amount          string    `json:"amount"`
	Currency        string    `json:"currency"`
	ReleasedAt      time.Time `json:"released_at"`
}

// IBRebateReversedV1 is the payload for ib.rebate_reversed.v1.
type IBRebateReversedV1 struct {
	RebateID        int64     `json:"rebate_id"`
	BeneficiaryIBID int64     `json:"beneficiary_ib_id"`
	Amount          string    `json:"amount"`
	Currency        string    `json:"currency"`
	Reason          string    `json:"reason,omitempty"`
	ReversedAt      time.Time `json:"reversed_at"`
}

// IBRebateTransferRequestedV1 is the payload for ib.rebate_transfer_requested.v1.
type IBRebateTransferRequestedV1 struct {
	TransferID     int64     `json:"transfer_id"`
	IBID           int64     `json:"ib_id"`
	UserID         int64     `json:"user_id"`
	Amount         string    `json:"amount"`
	Currency       string    `json:"currency"`
	IdempotencyKey string    `json:"idempotency_key"`
	RequestedAt    time.Time `json:"requested_at"`
}

// IBRebateTransferCompletedV1 is the payload for ib.rebate_transfer_completed.v1.
type IBRebateTransferCompletedV1 struct {
	TransferID  int64     `json:"transfer_id"`
	IBID        int64     `json:"ib_id"`
	UserID      int64     `json:"user_id"`
	Amount      string    `json:"amount"`
	Currency    string    `json:"currency"`
	CompletedAt time.Time `json:"completed_at"`
}

// IBRebateTransferFailedV1 is the payload for ib.rebate_transfer_failed.v1.
type IBRebateTransferFailedV1 struct {
	TransferID int64     `json:"transfer_id"`
	IBID       int64     `json:"ib_id"`
	UserID     int64     `json:"user_id"`
	Amount     string    `json:"amount"`
	Currency   string    `json:"currency"`
	Reason     string    `json:"reason,omitempty"`
	FailedAt   time.Time `json:"failed_at"`
}
