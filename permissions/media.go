package permissions

// MediaPermissions returns media-domain permissions owned by media-service.
func MediaPermissions() []Permission {
	return []Permission{
		{Key: MediaAdminRead, Service: ServiceMedia, Module: "media", Group: "Media", Name: "Admin Read Media", Description: "Admin read access to media files (ACL override)"},
		{Key: MediaFileManage, Service: ServiceMedia, Module: "media", Group: "Media", Name: "Manage Media Files", Description: "Create and complete media uploads via S2S"},
		{Key: MediaFileDelete, Service: ServiceMedia, Module: "media", Group: "Media", Name: "Delete Media Files", Description: "Delete media files via S2S"},
	}
}
