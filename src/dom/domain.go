package dom

import (
	"fmt"
	"net"

	namecheap "github.com/rustyeddy/go-namecheap"
	log "github.com/sirupsen/logrus"
)

// Domain is our struct wrapped around the raw namecheap data
type Domain struct {
	*namecheap.DomainGetListResult
	ns       []*net.NS
	provider string
	Err      error
}

// DomainFromNC
func DomainFromNC(d namecheap.DomainGetListResult) (dom Domain) {
	dom = Domain{
		DomainGetListResult: &d,

		provider: "namecheap",
	}
	return dom
}

// Nameservers will return the
func (d *Domain) Nameservers() []*net.NS {
	// TODO: an error
	var err error
	d.ns, err = net.LookupNS(d.Name)
	if err != nil {
		log.Warnf("NS lookup error %v", err)
		d.Err = err
		return nil
	}
	return d.ns
}

func (d *Domain) AddNameserver(name string) {
	ns := &net.NS{
		Host: name,
	}
	d.ns = append(d.ns, ns)
}

func (d *Domain) Provider() string {
	return d.provider
}

func (d *Domain) String() (s string) {
	nstr := ""
	for _, ns := range d.Nameservers() {
		nstr = nstr + ns.Host
	}
	return fmt.Sprintf("%s NS: [%s]", d.Name, nstr)
}
