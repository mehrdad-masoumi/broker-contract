package schemas

import _ "embed"

// NotificationRequestedV1 is the JSON Schema for routing key
// notification.requested.v1.
//
//go:embed notification.requested.v1.json
var NotificationRequestedV1 []byte

// UserRegisteredV1 is the JSON Schema for the user.registered.v1 event payload.
//
//go:embed user.registered.v1.json
var UserRegisteredV1 []byte

// UserDisabledV1 is the JSON Schema for the user.disabled.v1 event payload.
//
//go:embed user.disabled.v1.json
var UserDisabledV1 []byte

// UserVerifiedV1 is the JSON Schema for the user.verified.v1 event payload.
//
//go:embed user.verified.v1.json
var UserVerifiedV1 []byte

// UserUpdatedV1 is the JSON Schema for the user.updated.v1 event payload.
//
//go:embed user.updated.v1.json
var UserUpdatedV1 []byte

// UserLoggedInV1 is the JSON Schema for the user.logged_in.v1 event payload.
//
//go:embed user.logged_in.v1.json
var UserLoggedInV1 []byte

// TradeClosedV1 is the JSON Schema for the trade.closed.v1 event payload.
//
//go:embed trade.closed.v1.json
var TradeClosedV1 []byte
