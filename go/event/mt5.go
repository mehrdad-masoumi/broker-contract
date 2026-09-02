package event

// Manager API infrastructure event types published to broker.events.
// These are not trade.opened.v1 / trade.closed.v1.
const (
	TypeMt5OrderCreatedV1           = "mt5.order.created.v1"
	TypeMt5OrderUpdatedV1           = "mt5.order.updated.v1"
	TypeMt5OrderDeletedV1           = "mt5.order.deleted.v1"
	TypeMt5DealCreatedV1            = "mt5.deal.created.v1"
	TypeMt5DealUpdatedV1            = "mt5.deal.updated.v1"
	TypeMt5DealDeletedV1            = "mt5.deal.deleted.v1"
	TypeMt5DealPerformedV1          = "mt5.deal.performed.v1"
	TypeMt5PositionCreatedV1        = "mt5.position.created.v1"
	TypeMt5PositionUpdatedV1        = "mt5.position.updated.v1"
	TypeMt5PositionClosedV1         = "mt5.position.closed.v1"
	TypeMt5ReconciliationRequiredV1 = "mt5.manager.reconciliation_required.v1"
)

// Mt5TradePayload is the envelope payload copied from official Manager API objects.
// Volume is native MT5 units (volume + volume_ext). Monetary fields are JSON numbers
// because they come from CIMT* getters (double); consumers should treat them as
// already-lossy float copies of the native API, not as fabricated decimals.
type Mt5TradePayload struct {
	ServerID          string  `json:"server_id"`
	Login             uint64  `json:"login"`
	Ticket            uint64  `json:"ticket"`
	OrderID           uint64  `json:"order_id"`
	DealID            uint64  `json:"deal_id"`
	PositionID        uint64  `json:"position_id"`
	Symbol            string  `json:"symbol"`
	Action            uint32  `json:"action"`
	State             uint32  `json:"state"`
	Entry             uint32  `json:"entry"`
	Volume            uint64  `json:"volume"`
	VolumeExt         uint64  `json:"volume_ext"`
	Price             float64 `json:"price"`
	Profit            float64 `json:"profit"`
	Commission        float64 `json:"commission"`
	Swap              float64 `json:"swap"`
	OpenTime          int64   `json:"open_time"`
	UpdateTime        int64   `json:"update_time"`
	CloseTime         int64   `json:"close_time"`
	TimeMsc           int64   `json:"time_msc"`
	ModificationFlags uint32  `json:"modification_flags"`
	Callback          string  `json:"callback"`
	IdempotencyKey    string  `json:"idempotency_key"`
	Reason            string  `json:"reason,omitempty"`
	ReceivedAt        string  `json:"received_at,omitempty"`
}
