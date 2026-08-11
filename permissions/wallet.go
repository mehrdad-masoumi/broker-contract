package permissions

// WalletPermissions returns wallet-service domain permissions.
func WalletPermissions() []Permission {
	return []Permission{
		{Key: WalletRead, Service: "wallet", Module: "wallet", Group: "Wallets", Name: "Read Wallet", Description: "List and view wallets", Route: "GET /list"},
		{Key: WalletTransactionRead, Service: "wallet", Module: "wallet", Group: "Wallets", Name: "Read Transactions", Description: "View wallet transactions", Route: "GET /:wallet_id/transactions"},
		{Key: WalletDepositRead, Service: "wallet", Module: "wallet", Group: "Deposits", Name: "Read Deposits", Description: "View deposit information"},
		{Key: WalletWithdrawRead, Service: "wallet", Module: "wallet", Group: "Withdrawals", Name: "Read Withdrawals", Description: "List and view withdrawals", Route: "GET /withdrawals"},
		{Key: WalletWithdrawApprove, Service: "wallet", Module: "wallet", Group: "Withdrawals", Name: "Approve Withdrawal", Description: "Approve withdrawal requests", Route: "POST /withdrawals/:id/approve"},
		{Key: WalletWithdrawReject, Service: "wallet", Module: "wallet", Group: "Withdrawals", Name: "Reject Withdrawal", Description: "Reject withdrawal requests", Route: "POST /withdrawals/:id/reject"},
		{Key: WalletWithdrawProcess, Service: "wallet", Module: "wallet", Group: "Withdrawals", Name: "Process Withdrawal", Description: "Process, complete, or fail withdrawals", Route: "POST /withdrawals/:id/process"},
		{Key: WalletWithdrawSettings, Service: "wallet", Module: "wallet", Group: "Withdrawals", Name: "Manage Withdrawal Settings", Description: "View and update withdrawal settings", Route: "GET /withdrawal-settings"},
		{Key: WalletWithdrawNetworks, Service: "wallet", Module: "wallet", Group: "Withdrawals", Name: "Manage Withdrawal Networks", Description: "Manage withdrawal networks", Route: "GET /withdrawal-networks"},
	}
}
