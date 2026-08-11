package event

import "time"

// Support / ticket event types. Producer: support-service.
// Legacy UPPER_SNAKE aliases remain for dual-run consumers.
const (
	TypeTicketCreatedV1  = "ticket.created.v1"
	TypeTicketAssignedV1 = "ticket.assigned.v1"
	TypeTicketRepliedV1  = "ticket.replied.v1"
	TypeTicketClosedV1   = "ticket.closed.v1"

	TypeOperatorPerformanceRecordedV1 = "operator_performance.recorded.v1"

	LegacyTicketCreated               = "TICKET_CREATED"
	LegacyTicketAssigned              = "TICKET_ASSIGNED"
	LegacyOperatorPerformanceRecorded = "OPERATOR_PERFORMANCE_RECORDED"
)

// TicketCreatedV1 is a minimal ticket lifecycle payload (no message body PII).
type TicketCreatedV1 struct {
	TicketID     int64     `json:"ticket_id"`
	UserID       int64     `json:"user_id"`
	DepartmentID int64     `json:"department_id,omitempty"`
	OccurredAt   time.Time `json:"occurred_at"`
}

// TicketAssignedV1 is a minimal assignment payload.
type TicketAssignedV1 struct {
	TicketID   int64     `json:"ticket_id"`
	UserID     int64     `json:"user_id"`
	AssigneeID int64     `json:"assignee_id"`
	OccurredAt time.Time `json:"occurred_at"`
}

// TicketRepliedV1 is a minimal reply signal (no answer body).
type TicketRepliedV1 struct {
	TicketID   int64     `json:"ticket_id"`
	UserID     int64     `json:"user_id"`
	ActorID    int64     `json:"actor_id,omitempty"`
	OccurredAt time.Time `json:"occurred_at"`
}

// TicketClosedV1 is a minimal close signal.
type TicketClosedV1 struct {
	TicketID   int64     `json:"ticket_id"`
	UserID     int64     `json:"user_id"`
	OccurredAt time.Time `json:"occurred_at"`
}

// OperatorPerformanceRecordedV1 is a minimal OP signal.
type OperatorPerformanceRecordedV1 struct {
	OperatorID int64     `json:"operator_id"`
	TicketID   int64     `json:"ticket_id,omitempty"`
	OccurredAt time.Time `json:"occurred_at"`
}
