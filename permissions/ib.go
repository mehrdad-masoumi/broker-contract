package permissions

// IBPermissions returns ib-service domain permissions.
func IBPermissions() []Permission {
	return []Permission{
		{Key: IBRead, Service: "ib", Module: "ib", Group: "IB", Name: "Read IB", Description: "View IB profiles and applications", Route: "GET /ibs"},
		{Key: IBUpdate, Service: "ib", Module: "ib", Group: "IB", Name: "Update IB", Description: "Update IB profile state (rank, suspend, activate)", Route: "PUT /ibs/:id/rank"},
		{Key: IBApplicationRead, Service: "ib", Module: "ib", Group: "Applications", Name: "Read Applications", Description: "List and view IB applications", Route: "GET /applications"},
		{Key: IBApplicationApprove, Service: "ib", Module: "ib", Group: "Applications", Name: "Approve Application", Description: "Approve IB applications", Route: "POST /applications/:id/approve"},
		{Key: IBApplicationReject, Service: "ib", Module: "ib", Group: "Applications", Name: "Reject Application", Description: "Reject IB applications", Route: "POST /applications/:id/reject"},
		{Key: IBRankManage, Service: "ib", Module: "ib", Group: "Ranks", Name: "Manage Ranks", Description: "Create, update, and delete IB ranks", Route: "POST /ranks"},
		{Key: IBCampaignManage, Service: "ib", Module: "ib", Group: "Campaigns", Name: "Manage Campaigns", Description: "Update IB campaigns", Route: "PUT /campaigns/:id"},
		{Key: IBRebateManage, Service: "ib", Module: "ib", Group: "Rebates", Name: "Manage Rebates", Description: "Manage rebate rules and reverse rebates", Route: "POST /rebate-rules"},
		{Key: IBSettingsManage, Service: "ib", Module: "ib", Group: "Settings", Name: "Manage Settings", Description: "View and update IB settings", Route: "PUT /settings"},
		{Key: IBReportRead, Service: "ib", Module: "ib", Group: "Reports", Name: "Read Reports", Description: "View IB dashboards and reports", Route: "GET /dashboard"},
	}
}
