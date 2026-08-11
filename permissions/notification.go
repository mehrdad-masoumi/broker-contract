package permissions

// NotificationPermissions returns notification-service domain permissions.
func NotificationPermissions() []Permission {
	return []Permission{
		{Key: NotificationRead, Service: "notification", Module: "notification", Group: "Notifications", Name: "Read Notifications", Description: "View notifications and delivery status"},
		{Key: NotificationSend, Service: "notification", Module: "notification", Group: "Notifications", Name: "Send Notification", Description: "Create and send notifications"},
		{Key: NotificationTemplate, Service: "notification", Module: "notification", Group: "Templates", Name: "Manage Templates", Description: "Manage notification templates"},
		{Key: NotificationBatch, Service: "notification", Module: "notification", Group: "Batches", Name: "Manage Batches", Description: "Manage notification batches"},
	}
}
