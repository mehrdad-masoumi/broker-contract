package permissions

// BrokerPermissions returns trading/MT5 and platform-config permissions owned by broker-service.
func BrokerPermissions() []Permission {
	return []Permission{
		{Key: BrokerAccountCreate, Service: ServiceBroker, Module: "broker", Group: "Accounts", Name: "Create Account", Description: "Create new broker accounts", Route: "POST /account"},
		{Key: BrokerAccountRead, Service: ServiceBroker, Module: "broker", Group: "Accounts", Name: "Read Account", Description: "View broker accounts", Route: "GET /account"},
		{Key: BrokerAccountUpdate, Service: ServiceBroker, Module: "broker", Group: "Accounts", Name: "Update Account", Description: "Update broker account information", Route: "PUT /account/:id"},
		{Key: BrokerAccountDelete, Service: ServiceBroker, Module: "broker", Group: "Accounts", Name: "Delete Account", Description: "Delete broker accounts", Route: "DELETE /account/:id"},
		{Key: BrokerTradeRead, Service: ServiceBroker, Module: "broker", Group: "Trades", Name: "Read Trade", Description: "View trade information", Route: "GET /trades"},
		{Key: BrokerPackageCreate, Service: ServiceBroker, Module: "broker", Group: "Packages", Name: "Create Package", Description: "Create new broker packages", Route: "POST /broker-packages"},
		{Key: BrokerPackageRead, Service: ServiceBroker, Module: "broker", Group: "Packages", Name: "Read Package", Description: "View broker packages", Route: "GET /broker-packages"},
		{Key: BrokerPackageUpdate, Service: ServiceBroker, Module: "broker", Group: "Packages", Name: "Update Package", Description: "Update broker packages", Route: "PUT /broker-packages/:id"},
		{Key: BrokerPackageDelete, Service: ServiceBroker, Module: "broker", Group: "Packages", Name: "Delete Package", Description: "Delete broker packages", Route: "DELETE /broker-packages/:id"},
		{Key: BrokerAccountTypeCreate, Service: ServiceBroker, Module: "broker", Group: "Account Types", Name: "Create Account Type", Description: "Create new broker account types", Route: "POST /account-types"},
		{Key: BrokerAccountTypeRead, Service: ServiceBroker, Module: "broker", Group: "Account Types", Name: "Read Account Type", Description: "View broker account types", Route: "GET /account-types"},
		{Key: BrokerAccountTypeGet, Service: ServiceBroker, Module: "broker", Group: "Account Types", Name: "Get Account Type", Description: "View a broker account type by ID", Route: "GET /account-types/:id"},
		{Key: BrokerAccountTypeUpdate, Service: ServiceBroker, Module: "broker", Group: "Account Types", Name: "Update Account Type", Description: "Update broker account types", Route: "PUT /account-types/:id"},
		{Key: BrokerAccountTypeDelete, Service: ServiceBroker, Module: "broker", Group: "Account Types", Name: "Delete Account Type", Description: "Delete broker account types", Route: "DELETE /account-types/:id"},
		{Key: BrokerMetaGroupsRead, Service: ServiceBroker, Module: "broker", Group: "Account Types", Name: "Read Meta Groups", Description: "List MT5 groups available for account type mapping", Route: "GET /meta/groups"},
		{Key: BrokerAccountCurrencyCreate, Service: ServiceBroker, Module: "broker", Group: "Currencies", Name: "Create Account Currency", Description: "Create trading account currencies", Route: "POST /account-currencies"},
		{Key: BrokerAccountCurrencyRead, Service: ServiceBroker, Module: "broker", Group: "Currencies", Name: "Read Account Currency", Description: "View trading account currencies", Route: "GET /account-currencies"},
		{Key: BrokerAccountCurrencyGet, Service: ServiceBroker, Module: "broker", Group: "Currencies", Name: "Get Account Currency", Description: "View a trading account currency by ID", Route: "GET /account-currencies/:id"},
		{Key: BrokerAccountCurrencyUpdate, Service: ServiceBroker, Module: "broker", Group: "Currencies", Name: "Update Account Currency", Description: "Update trading account currencies", Route: "PUT /account-currencies/:id"},
		{Key: BrokerAccountCurrencyDelete, Service: ServiceBroker, Module: "broker", Group: "Currencies", Name: "Delete Account Currency", Description: "Delete trading account currencies", Route: "DELETE /account-currencies/:id"},
		{Key: BrokerPositionRead, Service: ServiceBroker, Module: "broker", Group: "Positions", Name: "Read Position", Description: "View trading positions", Route: "GET /positions"},
		{Key: BrokerPositionCreate, Service: ServiceBroker, Module: "broker", Group: "Positions", Name: "Create Position", Description: "Open new trading positions", Route: "POST /positions"},
		{Key: BrokerPositionDelete, Service: ServiceBroker, Module: "broker", Group: "Positions", Name: "Delete Position", Description: "Close trading positions", Route: "DELETE /positions/:ticket"},
		{Key: BrokerTradeSync, Service: ServiceBroker, Module: "broker", Group: "Trades", Name: "Sync Trades", Description: "Synchronize trades", Route: "POST /trades/sync"},
		{Key: BrokerPositionSync, Service: ServiceBroker, Module: "broker", Group: "Positions", Name: "Sync Positions", Description: "Synchronize positions", Route: "POST /positions/sync"},
		{Key: BrokerTradeReportRead, Service: ServiceBroker, Module: "broker", Group: "Trades", Name: "Read Trade Report", Description: "View trade reports and statistics", Route: "GET /trades/report"},

		{Key: ConfigRead, Service: ServiceBroker, Module: "config", Group: "Config", Name: "Read Config", Description: "View configuration settings", Route: "GET /config"},
		{Key: ConfigUpdate, Service: ServiceBroker, Module: "config", Group: "Config", Name: "Update Config", Description: "Update configuration settings", Route: "PUT /config"},
	}
}
