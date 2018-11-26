package dom

import "testing"

func TestNC(t *testing.T) {
	doms := ncDomains()
	if doms == nil {
		t.Error("ncDomains expected a few domains got nil")
	}

	if len(doms.Domains()) < 1 {
		t.Error("expected a lot of domains got (0)")
	}
}
