package permissions

// SupportPermissions returns ticket/department/OP permissions owned by support-service.
// Canonical keys are fine-grained (ticket.* / department.* / support.*); colon keys remain as legacy aliases.
func SupportPermissions() []Permission {
	return []Permission{
		// Canonical fine-grained — tickets
		{Key: TicketCreateKey, Service: ServiceSupport, Module: "ticket", Group: "Tickets", Name: "Create Ticket", Description: "Create new support tickets", Route: "POST /ticket"},
		{Key: TicketReadKey, Service: ServiceSupport, Module: "ticket", Group: "Tickets", Name: "Read Ticket", Description: "View support tickets", Route: "GET /ticket/:id"},
		{Key: TicketListKey, Service: ServiceSupport, Module: "ticket", Group: "Tickets", Name: "List Tickets", Description: "List all support tickets", Route: "GET /ticket/list"},
		{Key: TicketReplyKey, Service: ServiceSupport, Module: "ticket", Group: "Tickets", Name: "Reply Ticket", Description: "Answer support tickets", Route: "POST /ticket/:id/answer"},
		{Key: TicketAssignKey, Service: ServiceSupport, Module: "ticket", Group: "Tickets", Name: "Assign Ticket", Description: "Assign support tickets", Route: "PATCH /ticket/:id/assignment"},
		{Key: TicketCloseKey, Service: ServiceSupport, Module: "ticket", Group: "Tickets", Name: "Close Ticket", Description: "Close support tickets", Route: "POST /ticket/:id/close"},
		{Key: TicketReportKey, Service: ServiceSupport, Module: "ticket", Group: "Tickets", Name: "Read Ticket Reports", Description: "View ticket reports and statistics", Route: "GET /ticket/total-report"},

		// Legacy colon ticket aliases
		{Key: TicketCreate, Service: ServiceSupport, Module: "ticket", Group: "Tickets", Name: "Create Ticket (legacy)", Description: "Legacy alias for ticket.create"},
		{Key: TicketRead, Service: ServiceSupport, Module: "ticket", Group: "Tickets", Name: "Read Ticket (legacy)", Description: "Legacy alias for ticket.read"},
		{Key: TicketList, Service: ServiceSupport, Module: "ticket", Group: "Tickets", Name: "List Tickets (legacy)", Description: "Legacy alias for ticket.list"},
		{Key: TicketAnswer, Service: ServiceSupport, Module: "ticket", Group: "Tickets", Name: "Answer Ticket (legacy)", Description: "Legacy alias for ticket.reply"},
		{Key: TicketClose, Service: ServiceSupport, Module: "ticket", Group: "Tickets", Name: "Close Ticket (legacy)", Description: "Legacy alias for ticket.close"},
		{Key: TicketReportRead, Service: ServiceSupport, Module: "ticket", Group: "Tickets", Name: "Read Ticket Reports (legacy)", Description: "Legacy alias for ticket.report.read"},

		// Departments
		{Key: DepartmentRead, Service: ServiceSupport, Module: "department", Group: "Departments", Name: "Read Department", Description: "View support departments", Route: "GET /departments"},
		{Key: DepartmentManage, Service: ServiceSupport, Module: "department", Group: "Departments", Name: "Manage Department", Description: "Create and update departments, members, and roles", Route: "POST /departments"},

		// Operator performance
		{Key: SupportPerformanceRead, Service: ServiceSupport, Module: "operator_performance", Group: "Reports", Name: "Read Operator Performance", Description: "View operator performance statistics and reports", Route: "GET /backoffice/operator-performance/statistics"},
		{Key: OperatorPerformanceRead, Service: ServiceSupport, Module: "operator_performance", Group: "Reports", Name: "Read Operator Performance (legacy)", Description: "Legacy alias for support.performance.read"},
	}
}
