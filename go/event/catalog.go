package event

// Producer service names for domain events (must not be generic "core").
const (
	ProducerUser             = "user-service"
	ProducerAuth             = "auth-service"
	ProducerWallet           = "wallet-service"
	ProducerIB               = "ib-service"
	ProducerNotification     = "notification-service"
	ProducerBroker           = "broker-service"
	ProducerKYC              = "kyc-service"
	ProducerSupport          = "support-service"
	ProducerBonus            = "bonus-service"
	ProducerMedia            = "media-service"
	ProducerMt5ManagerBridge = "mt5-manager-bridge"
)

// CatalogEntry describes ownership of a versioned or legacy event type.
type CatalogEntry struct {
	EventType string
	Producer  string
	Legacy    bool // true for UPPER_SNAKE dual-run aliases
}

// Catalog returns the audited event ownership map.
// Legacy aliases are preserved until all consumers are verified migrated.
func Catalog() []CatalogEntry {
	return []CatalogEntry{
		// user-service
		{EventType: TypeUserRegisteredV1, Producer: ProducerUser},
		{EventType: TypeUserDisabledV1, Producer: ProducerUser},
		{EventType: TypeUserVerifiedV1, Producer: ProducerUser},
		{EventType: TypeUserUpdatedV1, Producer: ProducerUser},
		{EventType: TypeUserLoggedInV1, Producer: ProducerUser},

		// wallet-service
		{EventType: TypeDepositCompletedV1, Producer: ProducerWallet},
		{EventType: TypeWithdrawalCompletedV1, Producer: ProducerWallet},
		{EventType: TypeWalletCreditCompletedV1, Producer: ProducerWallet},

		// broker-service (trading)
		{EventType: TypeTradingAccountCreatedV1, Producer: ProducerBroker},
		{EventType: TypeTradeOpenedV1, Producer: ProducerBroker},
		{EventType: TypeTradeClosedV1, Producer: ProducerBroker},
		{EventType: TypeTradeCancelledV1, Producer: ProducerBroker},

		// ib-service
		{EventType: TypeIBApplicationSubmittedV1, Producer: ProducerIB},
		{EventType: TypeIBApplicationApprovedV1, Producer: ProducerIB},
		{EventType: TypeIBApplicationRejectedV1, Producer: ProducerIB},
		{EventType: TypeIBRankChangedV1, Producer: ProducerIB},
		{EventType: TypeIBSuspendedV1, Producer: ProducerIB},
		{EventType: TypeIBActivatedV1, Producer: ProducerIB},
		{EventType: TypeIBReferralCreatedV1, Producer: ProducerIB},
		{EventType: TypeIBUserAttributedV1, Producer: ProducerIB},
		{EventType: TypeIBRebateCreatedV1, Producer: ProducerIB},
		{EventType: TypeIBRebateReleasedV1, Producer: ProducerIB},
		{EventType: TypeIBRebateReversedV1, Producer: ProducerIB},
		{EventType: TypeIBRebateTransferRequestedV1, Producer: ProducerIB},
		{EventType: TypeIBRebateTransferCompletedV1, Producer: ProducerIB},
		{EventType: TypeIBRebateTransferFailedV1, Producer: ProducerIB},

		// kyc-service
		{EventType: TypeKYCSubmittedV1, Producer: ProducerKYC},
		{EventType: TypeKYCApprovedV1, Producer: ProducerKYC},
		{EventType: TypeKYCRejectedV1, Producer: ProducerKYC},
		{EventType: TypeKYCMoreInfoRequiredV1, Producer: ProducerKYC},
		{EventType: TypeKYCStatusChangedV1, Producer: ProducerKYC},
		{EventType: LegacyKYCSubmitted, Producer: ProducerKYC, Legacy: true},
		{EventType: LegacyKYCApproved, Producer: ProducerKYC, Legacy: true},
		{EventType: LegacyKYCRejected, Producer: ProducerKYC, Legacy: true},
		{EventType: LegacyKYCMoreInfoRequired, Producer: ProducerKYC, Legacy: true},

		// support-service
		{EventType: TypeTicketCreatedV1, Producer: ProducerSupport},
		{EventType: TypeTicketAssignedV1, Producer: ProducerSupport},
		{EventType: TypeTicketRepliedV1, Producer: ProducerSupport},
		{EventType: TypeTicketClosedV1, Producer: ProducerSupport},
		{EventType: TypeOperatorPerformanceRecordedV1, Producer: ProducerSupport},
		{EventType: LegacyTicketCreated, Producer: ProducerSupport, Legacy: true},
		{EventType: LegacyTicketAssigned, Producer: ProducerSupport, Legacy: true},
		{EventType: LegacyOperatorPerformanceRecorded, Producer: ProducerSupport, Legacy: true},

		// bonus-service
		{EventType: TypeBonusGrantedV1, Producer: ProducerBonus},
		{EventType: TypeBonusRevokedV1, Producer: ProducerBonus},
		{EventType: LegacyBonusGranted, Producer: ProducerBonus, Legacy: true},
		{EventType: LegacyBonusRevoked, Producer: ProducerBonus, Legacy: true},

		// media-service (reserved)
		{EventType: TypeMediaUploadCompletedV1, Producer: ProducerMedia},
		{EventType: TypeMediaFileDeletedV1, Producer: ProducerMedia},

		// mt5-manager-bridge (Manager API infrastructure; not trade.opened.v1)
		{EventType: TypeMt5OrderCreatedV1, Producer: ProducerMt5ManagerBridge},
		{EventType: TypeMt5OrderUpdatedV1, Producer: ProducerMt5ManagerBridge},
		{EventType: TypeMt5OrderDeletedV1, Producer: ProducerMt5ManagerBridge},
		{EventType: TypeMt5DealCreatedV1, Producer: ProducerMt5ManagerBridge},
		{EventType: TypeMt5DealUpdatedV1, Producer: ProducerMt5ManagerBridge},
		{EventType: TypeMt5DealDeletedV1, Producer: ProducerMt5ManagerBridge},
		{EventType: TypeMt5DealPerformedV1, Producer: ProducerMt5ManagerBridge},
		{EventType: TypeMt5PositionCreatedV1, Producer: ProducerMt5ManagerBridge},
		{EventType: TypeMt5PositionUpdatedV1, Producer: ProducerMt5ManagerBridge},
		{EventType: TypeMt5PositionClosedV1, Producer: ProducerMt5ManagerBridge},
		{EventType: TypeMt5ReconciliationRequiredV1, Producer: ProducerMt5ManagerBridge},
	}
}

// ProducerOf returns the owning producer for an event type, if catalogued.
func ProducerOf(eventType string) (string, bool) {
	for _, e := range Catalog() {
		if e.EventType == eventType {
			return e.Producer, true
		}
	}
	return "", false
}
