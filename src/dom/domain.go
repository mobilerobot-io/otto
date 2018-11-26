package dom

import (
	"fmt"
	"net"
	"strings"

	namecheap "github.com/billputer/go-namecheap"
)

// Domain is our struct wrapped around the raw namecheap data
type Domain struct {
	namecheap.DomainGetListResult

	ns []*net.NS
	//hosts []string

	sld, tld string // Used by namecheap queries

	provider string
	err      error
}

// DomainFromNC
func DomainFromNC(d namecheap.DomainGetListResult) (dom Domain, err error) {
	dom = Domain{
		DomainGetListResult: d,
		provider:            "namecheap",
	}

	parts := strings.Split(d.Name, ".")
	if len(parts) < 2 {
		return dom, fmt.Errorf("expected domain (sld.tld) got (%s) ", d.Name)
	}
	dom.tld = parts[len(parts)-1]
	dom.sld = parts[0]
	return dom, nil
}

func (d *Domain) Provider() string {
	return d.provider
}

// DNS Records
// ========================================================================

// DNSHosts will call namecheap.domains.dns.getHosts
func (d *Domain) DNSHosts() (nchosts *namecheap.DomainDNSGetHostsResult, err error) {
	cli := getClient()
	nchosts, err = cli.DomainsDNSGetHosts(d.sld, d.tld)
	return nchosts, err
}

func (d *Domain) String() (s string) {
	nstr := ""
	for _, ns := range d.Nameservers() {
		nstr = nstr + ns.Host
	}
	return fmt.Sprintf("%s NS: [%s]", d.Name, nstr)
}
