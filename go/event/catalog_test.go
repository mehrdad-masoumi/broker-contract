package event

import "testing"

func TestCatalogNoCoreProducer(t *testing.T) {
	for _, e := range Catalog() {
		if e.Producer == "" {
			t.Fatalf("event %q has empty producer", e.EventType)
		}
		if e.Producer == "core" || e.Producer == "core-service" {
			t.Fatalf("event %q still owned by %q", e.EventType, e.Producer)
		}
	}
}

func TestProducerOfKYCAndSupport(t *testing.T) {
	cases := map[string]string{
		TypeKYCStatusChangedV1: ProducerKYC,
		LegacyKYCApproved:      ProducerKYC,
		TypeTicketCreatedV1:    ProducerSupport,
		LegacyTicketCreated:    ProducerSupport,
		TypeBonusGrantedV1:     ProducerBonus,
		TypeTradeClosedV1:      ProducerBroker,
		TypeMt5DealCreatedV1:   ProducerMt5ManagerBridge,
	}
	for et, want := range cases {
		got, ok := ProducerOf(et)
		if !ok {
			t.Fatalf("missing catalog entry for %q", et)
		}
		if got != want {
			t.Fatalf("%q: producer=%q want %q", et, got, want)
		}
	}
}
