package permissions

// UserPermissions returns identity and RBAC-admin permissions owned by user-service.
func UserPermissions() []Permission {
	return []Permission{
		{Key: AuthUserRead, Service: "user", Module: "auth", Group: "Users", Name: "Read User", Description: "View user information", Route: "GET /users"},
		{Key: AuthUserCreate, Service: "user", Module: "auth", Group: "Users", Name: "Create User", Description: "Create new users", Route: "POST /users"},
		{Key: AuthUserUpdate, Service: "user", Module: "auth", Group: "Users", Name: "Update User", Description: "Update user information", Route: "PUT /users/:id"},
		{Key: AuthUserDelete, Service: "user", Module: "auth", Group: "Users", Name: "Delete User", Description: "Delete users", Route: "DELETE /users/:id"},
		{Key: AuthUserAssignRole, Service: "user", Module: "auth", Group: "Users", Name: "Assign Role", Description: "Assign roles to users", Route: "PUT /users/:id/roles"},
		{Key: AuthRoleRead, Service: "user", Module: "auth", Group: "Roles", Name: "Read Role", Description: "View roles", Route: "GET /roles"},
		{Key: AuthRoleCreate, Service: "user", Module: "auth", Group: "Roles", Name: "Create Role", Description: "Create new roles", Route: "POST /roles"},
		{Key: AuthRoleUpdate, Service: "user", Module: "auth", Group: "Roles", Name: "Update Role", Description: "Update role information", Route: "PUT /roles/:id"},
		{Key: AuthRoleDelete, Service: "user", Module: "auth", Group: "Roles", Name: "Delete Role", Description: "Delete roles", Route: "DELETE /roles/:id"},
		{Key: AuthPermissionRead, Service: "user", Module: "auth", Group: "Permissions", Name: "Read Permission", Description: "View permissions", Route: "GET /permissions"},
		{Key: AuthRolePermissionManage, Service: "user", Module: "auth", Group: "Roles", Name: "Manage Role Permissions", Description: "Manage role permissions", Route: "PUT /roles/:id/permissions"},

		{Key: UserProfileRead, Service: "user", Module: "user", Group: "Profile", Name: "Read Profile", Description: "View user profile", Route: "GET /profile"},
		{Key: UserProfileUpdate, Service: "user", Module: "user", Group: "Profile", Name: "Update Profile", Description: "Update user profile", Route: "PUT /profile"},
		{Key: UserSettingsRead, Service: "user", Module: "user", Group: "Settings", Name: "Read Settings", Description: "View user settings", Route: "GET /settings"},
		{Key: UserSettingsUpdate, Service: "user", Module: "user", Group: "Settings", Name: "Update Settings", Description: "Update user settings", Route: "PUT /settings"},
		{Key: UserRegistrationReportRead, Service: "user", Module: "user", Group: "Reports", Name: "Read Registration Report", Description: "View user registration reports", Route: "GET /users/registration-report"},
		{Key: UserPasswordChange, Service: "user", Module: "user", Group: "Users", Name: "Change Password", Description: "Change user password", Route: "POST /users/:id/change-password"},
	}
}
