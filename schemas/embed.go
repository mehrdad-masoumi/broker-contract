package schemas

import _ "embed"

// NotificationRequestedV1 is the JSON Schema for routing key
// notification.requested.v1.
//
//go:embed notification.requested.v1.json
var NotificationRequestedV1 []byte
