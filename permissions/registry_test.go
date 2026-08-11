package permissions

import "testing"

func TestAllUniqueAndValid(t *testing.T) {
	perms := All()
	if len(perms) == 0 {
		t.Fatal("expected permissions")
	}
	if err := Validate(perms); err != nil {
		t.Fatal(err)
	}
}
