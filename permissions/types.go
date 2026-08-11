// Package permissions is the central shared permission catalog for all broker services.
// user-service syncs these definitions into its permissions table; domain services
// enforce keys from JWT claims via RequirePermission.
package permissions

// Permission is a code-first permission definition.
type Permission struct {
	Key         string // e.g. "kyc:submission:review"
	Service     string // owning service/domain, e.g. "core", "wallet", "user"
	Module      string // feature module within the service, e.g. "kyc", "broker"
	Group       string // Admin Panel grouping label, e.g. "KYC", "Accounts"
	Name        string // short title for Admin Panel
	Description string
	Route       string // optional "METHOD /path" for route-based guards
}

// Key constants — prefer these over raw strings in route middleware.
const (
	// user-service / auth RBAC admin
	AuthUserRead             = "auth:user:read"
	AuthUserCreate           = "auth:user:create"
	AuthUserUpdate           = "auth:user:update"
	AuthUserDelete           = "auth:user:delete"
	AuthUserAssignRole       = "auth:user:assign-role"
	AuthRoleRead             = "auth:role:read"
	AuthRoleCreate           = "auth:role:create"
	AuthRoleUpdate           = "auth:role:update"
	AuthRoleDelete           = "auth:role:delete"
	AuthPermissionRead       = "auth:permission:read"
	AuthRolePermissionManage = "auth:role_permission:manage"

	UserProfileRead             = "user:profile:read"
	UserProfileUpdate           = "user:profile:update"
	UserSettingsRead            = "user:settings:read"
	UserSettingsUpdate          = "user:settings:update"
	UserRegistrationReportRead  = "user:registration-report:read"
	UserPasswordChange          = "user:password:change"

	// core — broker
	BrokerAccountCreate         = "broker:account:create"
	BrokerAccountRead           = "broker:account:read"
	BrokerAccountUpdate         = "broker:account:update"
	BrokerAccountDelete         = "broker:account:delete"
	BrokerTradeRead             = "broker:trade:read"
	BrokerPackageCreate         = "broker:package:create"
	BrokerPackageRead           = "broker:package:read"
	BrokerPackageUpdate         = "broker:package:update"
	BrokerPackageDelete         = "broker:package:delete"
	BrokerAccountTypeCreate     = "broker:account-type:create"
	BrokerAccountTypeRead       = "broker:account-type:read"
	BrokerAccountTypeGet        = "broker:account-type:get"
	BrokerAccountTypeUpdate     = "broker:account-type:update"
	BrokerAccountTypeDelete     = "broker:account-type:delete"
	BrokerAccountCurrencyCreate = "broker:account-currency:create"
	BrokerAccountCurrencyRead   = "broker:account-currency:read"
	BrokerAccountCurrencyGet    = "broker:account-currency:get"
	BrokerAccountCurrencyUpdate = "broker:account-currency:update"
	BrokerAccountCurrencyDelete = "broker:account-currency:delete"
	BrokerPositionRead          = "broker:position:read"
	BrokerPositionCreate        = "broker:position:create"
	BrokerPositionDelete        = "broker:position:delete"
	BrokerTradeSync             = "broker:trade:sync"
	BrokerPositionSync          = "broker:position:sync"
	BrokerTradeReportRead       = "broker:trade:report:read"

	// core — ticket
	TicketCreate     = "ticket:create"
	TicketRead       = "ticket:read"
	TicketList       = "ticket:list"
	TicketAnswer     = "ticket:answer"
	TicketClose      = "ticket:close"
	TicketReportRead = "ticket:report:read"

	// core — kyc
	KYCSubmissionList   = "kyc:submission:list"
	KYCSubmissionRead   = "kyc:submission:read"
	KYCDocumentReview   = "kyc:document:review"
	KYCSubmissionReview = "kyc:submission:review"

	// core — bonus / config / operator
	BonusCreate              = "bonus:create"
	BonusRead                = "bonus:read"
	BonusUpdate              = "bonus:update"
	BonusDelete              = "bonus:delete"
	ConfigRead               = "config:read"
	ConfigUpdate             = "config:update"
	OperatorPerformanceRead  = "operator_performance:read"

	// wallet-service
	WalletRead               = "wallet:read"
	WalletTransactionRead    = "wallet:transaction:read"
	WalletDepositRead        = "wallet:deposit:read"
	WalletWithdrawRead       = "wallet:withdraw:read"
	WalletWithdrawApprove    = "wallet:withdraw:approve"
	WalletWithdrawReject     = "wallet:withdraw:reject"
	WalletWithdrawProcess    = "wallet:withdraw:process"
	WalletWithdrawSettings   = "wallet:withdraw:settings"
	WalletWithdrawNetworks   = "wallet:withdraw:networks"

	// ib-service
	IBRead            = "ib:read"
	IBUpdate          = "ib:update"
	IBApplicationRead = "ib:application:read"
	IBApplicationApprove = "ib:application:approve"
	IBApplicationReject  = "ib:application:reject"
	IBRankManage      = "ib:rank:manage"
	IBCampaignManage  = "ib:campaign:manage"
	IBRebateManage    = "ib:rebate:manage"
	IBSettingsManage  = "ib:settings:manage"
	IBReportRead      = "ib:report:read"

	// notification-service
	NotificationRead     = "notification:read"
	NotificationSend     = "notification:send"
	NotificationTemplate = "notification:template:manage"
	NotificationBatch    = "notification:batch:manage"
)
