package permissions

// CorePermissions returns broker-domain permissions owned by core.
func CorePermissions() []Permission {
	return []Permission{
		// Broker accounts
		{Key: BrokerAccountCreate, Service: "core", Module: "broker", Group: "Accounts", Name: "Create Account", Description: "Create new broker accounts", Route: "POST /account"},
		{Key: BrokerAccountRead, Service: "core", Module: "broker", Group: "Accounts", Name: "Read Account", Description: "View broker accounts", Route: "GET /account"},
		{Key: BrokerAccountUpdate, Service: "core", Module: "broker", Group: "Accounts", Name: "Update Account", Description: "Update broker account information", Route: "PUT /account/:id"},
		{Key: BrokerAccountDelete, Service: "core", Module: "broker", Group: "Accounts", Name: "Delete Account", Description: "Delete broker accounts", Route: "DELETE /account/:id"},
		{Key: BrokerTradeRead, Service: "core", Module: "broker", Group: "Trades", Name: "Read Trade", Description: "View trade information", Route: "GET /trades"},
		{Key: BrokerPackageCreate, Service: "core", Module: "broker", Group: "Packages", Name: "Create Package", Description: "Create new broker packages", Route: "POST /broker-packages"},
		{Key: BrokerPackageRead, Service: "core", Module: "broker", Group: "Packages", Name: "Read Package", Description: "View broker packages", Route: "GET /broker-packages"},
		{Key: BrokerPackageUpdate, Service: "core", Module: "broker", Group: "Packages", Name: "Update Package", Description: "Update broker packages", Route: "PUT /broker-packages/:id"},
		{Key: BrokerPackageDelete, Service: "core", Module: "broker", Group: "Packages", Name: "Delete Package", Description: "Delete broker packages", Route: "DELETE /broker-packages/:id"},
		{Key: BrokerAccountTypeCreate, Service: "core", Module: "broker", Group: "Account Types", Name: "Create Account Type", Description: "Create new broker account types", Route: "POST /account-types"},
		{Key: BrokerAccountTypeRead, Service: "core", Module: "broker", Group: "Account Types", Name: "Read Account Type", Description: "View broker account types", Route: "GET /account-types"},
		{Key: BrokerAccountTypeGet, Service: "core", Module: "broker", Group: "Account Types", Name: "Get Account Type", Description: "View a broker account type by ID", Route: "GET /account-types/:id"},
		{Key: BrokerAccountTypeUpdate, Service: "core", Module: "broker", Group: "Account Types", Name: "Update Account Type", Description: "Update broker account types", Route: "PUT /account-types/:id"},
		{Key: BrokerAccountTypeDelete, Service: "core", Module: "broker", Group: "Account Types", Name: "Delete Account Type", Description: "Delete broker account types", Route: "DELETE /account-types/:id"},
		{Key: BrokerAccountCurrencyCreate, Service: "core", Module: "broker", Group: "Currencies", Name: "Create Account Currency", Description: "Create trading account currencies", Route: "POST /account-currencies"},
		{Key: BrokerAccountCurrencyRead, Service: "core", Module: "broker", Group: "Currencies", Name: "Read Account Currency", Description: "View trading account currencies", Route: "GET /account-currencies"},
		{Key: BrokerAccountCurrencyGet, Service: "core", Module: "broker", Group: "Currencies", Name: "Get Account Currency", Description: "View a trading account currency by ID", Route: "GET /account-currencies/:id"},
		{Key: BrokerAccountCurrencyUpdate, Service: "core", Module: "broker", Group: "Currencies", Name: "Update Account Currency", Description: "Update trading account currencies", Route: "PUT /account-currencies/:id"},
		{Key: BrokerAccountCurrencyDelete, Service: "core", Module: "broker", Group: "Currencies", Name: "Delete Account Currency", Description: "Delete trading account currencies", Route: "DELETE /account-currencies/:id"},
		{Key: BrokerPositionRead, Service: "core", Module: "broker", Group: "Positions", Name: "Read Position", Description: "View trading positions", Route: "GET /positions"},
		{Key: BrokerPositionCreate, Service: "core", Module: "broker", Group: "Positions", Name: "Create Position", Description: "Open new trading positions", Route: "POST /positions"},
		{Key: BrokerPositionDelete, Service: "core", Module: "broker", Group: "Positions", Name: "Delete Position", Description: "Close trading positions", Route: "DELETE /positions/:ticket"},
		{Key: BrokerTradeSync, Service: "core", Module: "broker", Group: "Trades", Name: "Sync Trades", Description: "Synchronize trades", Route: "POST /trades/sync"},
		{Key: BrokerPositionSync, Service: "core", Module: "broker", Group: "Positions", Name: "Sync Positions", Description: "Synchronize positions", Route: "POST /positions/sync"},
		{Key: BrokerTradeReportRead, Service: "core", Module: "broker", Group: "Trades", Name: "Read Trade Report", Description: "View trade reports and statistics", Route: "GET /trades/report"},

		// Tickets
		{Key: TicketCreate, Service: "core", Module: "ticket", Group: "Tickets", Name: "Create Ticket", Description: "Create new support tickets", Route: "POST /ticket"},
		{Key: TicketRead, Service: "core", Module: "ticket", Group: "Tickets", Name: "Read Ticket", Description: "View support tickets", Route: "GET /ticket/:id"},
		{Key: TicketList, Service: "core", Module: "ticket", Group: "Tickets", Name: "List Tickets", Description: "List all support tickets", Route: "GET /ticket/list"},
		{Key: TicketAnswer, Service: "core", Module: "ticket", Group: "Tickets", Name: "Answer Ticket", Description: "Answer support tickets", Route: "POST /ticket/:id/answer"},
		{Key: TicketClose, Service: "core", Module: "ticket", Group: "Tickets", Name: "Close Ticket", Description: "Close support tickets", Route: "POST /ticket/:id/close"},
		{Key: TicketReportRead, Service: "core", Module: "ticket", Group: "Tickets", Name: "Read Ticket Reports", Description: "View ticket reports and statistics", Route: "GET /ticket/total-report"},

		// KYC
		{Key: KYCSubmissionList, Service: "core", Module: "kyc", Group: "KYC", Name: "List KYC Submissions", Description: "List user KYC submissions", Route: "GET /backoffice/kyc/submissions/list"},
		{Key: KYCSubmissionRead, Service: "core", Module: "kyc", Group: "KYC", Name: "Read KYC Submission", Description: "View user KYC submission details", Route: "GET /backoffice/kyc/submissions/:user_id"},
		{Key: KYCDocumentReview, Service: "core", Module: "kyc", Group: "KYC", Name: "Review KYC Document", Description: "Approve or reject KYC documents", Route: "PATCH /backoffice/kyc/documents/:id/review"},
		{Key: KYCSubmissionReview, Service: "core", Module: "kyc", Group: "KYC", Name: "Review KYC Submission", Description: "Approve, reject, or request more information for a KYC case", Route: "POST /backoffice/kyc/submissions/:user_id/review"},

		// Bonus / config / operator
		{Key: BonusCreate, Service: "core", Module: "bonus", Group: "Bonus", Name: "Create Bonus", Description: "Create bonus campaigns", Route: "POST /bonuses"},
		{Key: BonusRead, Service: "core", Module: "bonus", Group: "Bonus", Name: "Read Bonus", Description: "View bonus campaigns", Route: "GET /bonuses"},
		{Key: BonusUpdate, Service: "core", Module: "bonus", Group: "Bonus", Name: "Update Bonus", Description: "Update bonus campaigns and their triggers/rules/limits", Route: "PUT /bonuses/:id"},
		{Key: BonusDelete, Service: "core", Module: "bonus", Group: "Bonus", Name: "Delete Bonus", Description: "Delete bonus campaigns", Route: "DELETE /bonuses/:id"},
		{Key: ConfigRead, Service: "core", Module: "config", Group: "Config", Name: "Read Config", Description: "View configuration settings", Route: "GET /config"},
		{Key: ConfigUpdate, Service: "core", Module: "config", Group: "Config", Name: "Update Config", Description: "Update configuration settings", Route: "PUT /config"},
		{Key: OperatorPerformanceRead, Service: "core", Module: "operator_performance", Group: "Reports", Name: "Read Operator Performance", Description: "View operator performance statistics and reports", Route: "GET /backoffice/operator-performance/statistics"},
	}
}
