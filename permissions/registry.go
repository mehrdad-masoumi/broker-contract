package permissions

import "fmt"

// All returns the full permission catalog across all services.
func All() []Permission {
	out := make([]Permission, 0, 128)
	out = append(out, UserPermissions()...)
	out = append(out, CorePermissions()...)
	out = append(out, WalletPermissions()...)
	out = append(out, IBPermissions()...)
	out = append(out, NotificationPermissions()...)
	return out
}

// ByKey returns a map of key → Permission for the full catalog.
func ByKey() map[string]Permission {
	m := make(map[string]Permission, 128)
	for _, p := range All() {
		m[p.Key] = p
	}
	return m
}

// ByRoute builds METHOD+path → permission keys for in-memory route guards.
// Empty Route entries are skipped. Duplicate routes accumulate keys.
func ByRoute() map[string][]string {
	m := make(map[string][]string)
	for _, p := range All() {
		if p.Route == "" {
			continue
		}
		m[p.Route] = append(m[p.Route], p.Key)
	}
	return m
}

// Validate ensures every permission has required fields and unique keys.
func Validate(perms []Permission) error {
	seen := make(map[string]struct{}, len(perms))
	for _, p := range perms {
		if p.Key == "" {
			return fmt.Errorf("permission key cannot be empty")
		}
		if p.Name == "" {
			return fmt.Errorf("permission %q: name cannot be empty", p.Key)
		}
		if p.Service == "" {
			return fmt.Errorf("permission %q: service cannot be empty", p.Key)
		}
		if p.Module == "" {
			return fmt.Errorf("permission %q: module cannot be empty", p.Key)
		}
		if _, ok := seen[p.Key]; ok {
			return fmt.Errorf("duplicate permission key %q", p.Key)
		}
		seen[p.Key] = struct{}{}
	}
	return nil
}
