package dom

import "testing"

// TestFetchAndReadDomains first "fetches" the domains from the provider
// (namecheap), verifies it has more than one.  The second fetch reads the
// domain list from the cache.
func TestFetchAndReadDomains(t *testing.T) {
	var prov, cache *DomainManager

	// Fetch domains from the provider
	if prov = FetchDomains(); prov == nil {
		t.Error("expected (domains) got (nil)")
	}

	var plen, clen int
	if plen = len(prov.Domains()); plen < 1 {
		t.Errorf("expected (> 0) got (%d)", plen)
	}

	// Save the domains (just the domains, not the indexes)
	if err := prov.Save(); err != nil {
		t.Errorf("failed saving domains %v", err)
	}

	// Read the domains from the local cache (the filesystem)
	if cache = GetDomains(); cache == nil {
		t.Error("expected (cached) domains got (nil)")
	}

	if clen = len(cache.Domains()); clen < 1 {
		t.Errorf("expected (> 0) got (%d)", clen)
	}

	if plen != clen {
		t.Errorf("domain count expected (provider == cache) got (%d and %d)", plen, clen)
	}

	// TODO XXX - Compare domain lists to be thourough
}
