package event

import "time"

// Ingress event type / routing key constants (wallet / payments domain).
const (
	TypeDepositCompletedV1      = "deposit.completed.v1"
	TypeWithdrawalCompletedV1   = "withdrawal.completed.v1"
	TypeWalletCreditCompletedV1 = "wallet.credit.completed.v1"
)

// DepositCompletedV1 is the payload for deposit.completed.v1.
// Amount fields are decimal strings (never JSON floats).
type DepositCompletedV1 struct {
	DepositID            int64     `json:"deposit_id"`
	UserID               int64     `json:"user_id"`
	Amount               string    `json:"amount"`
	Currency             string    `json:"currency"`
	AmountInBaseCurrency string    `json:"amount_in_base_currency,omitempty"`
	ExchangeRate         string    `json:"exchange_rate,omitempty"`
	Status               string    `json:"status"`
	CompletedAt          time.Time `json:"completed_at"`
}

// WithdrawalCompletedV1 is the payload for withdrawal.completed.v1.
// Amount fields are decimal strings (never JSON floats).
type WithdrawalCompletedV1 struct {
	WithdrawalID         int64     `json:"withdrawal_id"`
	UserID               int64     `json:"user_id"`
	Amount               string    `json:"amount"`
	Currency             string    `json:"currency"`
	AmountInBaseCurrency string    `json:"amount_in_base_currency,omitempty"`
	ExchangeRate         string    `json:"exchange_rate,omitempty"`
	Status               string    `json:"status"`
	CompletedAt          time.Time `json:"completed_at"`
}

// WalletCreditCompletedV1 is the payload for wallet.credit.completed.v1
// (ack for IB rebate transfer → master wallet credit).
// Amount fields are decimal strings (never JSON floats).
type WalletCreditCompletedV1 struct {
	UserID         int64     `json:"user_id"`
	Amount         string    `json:"amount"`
	Currency       string    `json:"currency"`
	RefType        string    `json:"ref_type"`
	RefID          string    `json:"ref_id"`
	IdempotencyKey string    `json:"idempotency_key"`
	CompletedAt    time.Time `json:"completed_at"`
}
