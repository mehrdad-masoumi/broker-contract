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

func TestNoCoreServiceOwnership(t *testing.T) {
	for _, p := range All() {
		if p.Service == "core" {
			t.Fatalf("permission %q still owned by core", p.Key)
		}
	}
}

func TestOwnershipRetarget(t *testing.T) {
	byKey := ByKey()
	cases := map[string]string{
		BonusRead:              ServiceBonus,
		KYCSubmissionListKey:   ServiceKYC,
		KYCSubmissionList:      ServiceKYC,
		TicketReadKey:          ServiceSupport,
		TicketRead:             ServiceSupport,
		DepartmentRead:         ServiceSupport,
		SupportPerformanceRead: ServiceSupport,
		MediaAdminRead:         ServiceMedia,
		BrokerAccountRead:      ServiceBroker,
		ConfigRead:             ServiceBroker,
	}
	for key, want := range cases {
		p, ok := byKey[key]
		if !ok {
			t.Fatalf("missing permission %q", key)
		}
		if p.Service != want {
			t.Fatalf("%q: service=%q want %q", key, p.Service, want)
		}
	}
}

func TestBonusExtraRoutesInByRoute(t *testing.T) {
	m := ByRoute()
	if keys := m["POST /bonuses/evaluate"]; len(keys) == 0 || keys[0] != BonusRead {
		t.Fatalf("evaluate route: %v", keys)
	}
	if keys := m["POST /bonuses/:id/triggers"]; len(keys) == 0 || keys[0] != BonusUpdate {
		t.Fatalf("create trigger route: %v", keys)
	}
}

func TestFineGrainedKeysPresent(t *testing.T) {
	byKey := ByKey()
	for _, key := range []string{
		KYCSubmissionListKey, KYCSubmissionReadKey, KYCDocumentReviewKey, KYCSubmissionReviewKey,
		TicketReadKey, TicketReplyKey, TicketAssignKey, TicketCloseKey,
		DepartmentRead, DepartmentManage, SupportPerformanceRead,
		MediaAdminRead, MediaFileDelete, MediaFileManage,
	} {
		if _, ok := byKey[key]; !ok {
			t.Fatalf("missing fine-grained key %q", key)
		}
	}
}
