package event

import "time"

// Ingress event type / routing key constants (trading domain).
const (
	TypeTradingAccountCreatedV1 = "trading_account.created.v1"
	TypeTradeOpenedV1           = "trade.opened.v1"
	TypeTradeClosedV1           = "trade.closed.v1"
	TypeTradeCancelledV1        = "trade.cancelled.v1"
)

// TradingAccountCreatedV1 is the payload for trading_account.created.v1.
type TradingAccountCreatedV1 struct {
	UserID           int64     `json:"user_id"`
	TradingAccountID int64     `json:"trading_account_id"`
	AccountType      string    `json:"account_type"`
	Platform         string    `json:"platform,omitempty"`
	Currency         string    `json:"currency"`
	CreatedAt        time.Time `json:"created_at"`
}

// TradeOpenedV1 is the payload for trade.opened.v1.
// Monetary and lot fields are decimal strings (never JSON floats).
type TradeOpenedV1 struct {
	TradeID          int64     `json:"trade_id"`
	PositionID       int64     `json:"position_id,omitempty"`
	UserID           int64     `json:"user_id"`
	TradingAccountID int64     `json:"trading_account_id"`
	Symbol           string    `json:"symbol"`
	AccountType      string    `json:"account_type"`
	Side             string    `json:"side"`
	Lot              string    `json:"lot"`
	OpenPrice        string    `json:"open_price,omitempty"`
	OpenTime         time.Time `json:"open_time"`
}

// TradeClosedV1 is the payload for trade.closed.v1.
// Monetary and lot fields are decimal strings (never JSON floats).
type TradeClosedV1 struct {
	TradeID          int64     `json:"trade_id"`
	PositionID       int64     `json:"position_id,omitempty"`
	UserID           int64     `json:"user_id"`
	TradingAccountID int64     `json:"trading_account_id"`
	Symbol           string    `json:"symbol"`
	AccountType      string    `json:"account_type"`
	Side             string    `json:"side"`
	OpenTime         time.Time `json:"open_time"`
	CloseTime        time.Time `json:"close_time"`
	DurationSeconds  int64     `json:"duration_seconds"`
	Lot              string    `json:"lot"`
	OpenPrice        string    `json:"open_price,omitempty"`
	ClosePrice       string    `json:"close_price,omitempty"`
	Profit           string    `json:"profit,omitempty"`
	Commission       string    `json:"commission,omitempty"`
	Swap             string    `json:"swap,omitempty"`
	TradeStatus      string    `json:"trade_status"`
}

// TradeCancelledV1 is the payload for trade.cancelled.v1.
type TradeCancelledV1 struct {
	TradeID          int64     `json:"trade_id"`
	UserID           int64     `json:"user_id"`
	TradingAccountID int64     `json:"trading_account_id,omitempty"`
	CancelledAt      time.Time `json:"cancelled_at"`
	Reason           string    `json:"reason,omitempty"`
}
