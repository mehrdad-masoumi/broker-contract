package permissions

// BonusPermissions returns bonus-domain permissions owned by bonus-service.
func BonusPermissions() []Permission {
	return []Permission{
		{Key: BonusCreate, Service: ServiceBonus, Module: "bonus", Group: "Bonus", Name: "Create Bonus", Description: "Create bonus campaigns", Route: "POST /bonuses"},
		{Key: BonusRead, Service: ServiceBonus, Module: "bonus", Group: "Bonus", Name: "Read Bonus", Description: "View bonus campaigns", Route: "GET /bonuses"},
		{Key: BonusUpdate, Service: ServiceBonus, Module: "bonus", Group: "Bonus", Name: "Update Bonus", Description: "Update bonus campaigns and their triggers/rules/limits", Route: "PUT /bonuses/:id"},
		{Key: BonusDelete, Service: ServiceBonus, Module: "bonus", Group: "Bonus", Name: "Delete Bonus", Description: "Delete bonus campaigns", Route: "DELETE /bonuses/:id"},
	}
}

// BonusExtraRoutes maps nested/evaluate HTTP routes → permission keys for ByRoute.
// Kept separate so All()/Validate stay unique-keyed.
func BonusExtraRoutes() map[string]string {
	return map[string]string{
		"GET /bonuses/:id":             BonusRead,
		"POST /bonuses/evaluate":       BonusRead,
		"POST /bonuses/evaluate-user":  BonusRead,
		"POST /bonuses/evaluate-ib":    BonusRead,
		"GET /bonuses/:id/triggers":    BonusRead,
		"POST /bonuses/:id/triggers":   BonusUpdate,
		"PUT /bonuses/triggers/:id":    BonusUpdate,
		"DELETE /bonuses/triggers/:id": BonusUpdate,
		"GET /bonuses/:id/rules":       BonusRead,
		"POST /bonuses/:id/rules":      BonusUpdate,
		"PUT /bonuses/rules/:id":       BonusUpdate,
		"DELETE /bonuses/rules/:id":    BonusUpdate,
		"GET /bonuses/:id/limits":      BonusRead,
		"POST /bonuses/:id/limits":     BonusUpdate,
		"PUT /bonuses/limits/:id":      BonusUpdate,
		"DELETE /bonuses/limits/:id":   BonusUpdate,
	}
}
