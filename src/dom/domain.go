package dom

import (
	"net"

	namecheap "github.com/rustyeddy/go-namecheap"
	log "github.com/sirupsen/logrus"
)

// Domain is our struct wrapped around the raw namecheap data
type Domain struct {
	*namecheap.DomainGetListResult
	NS       []*net.NS
	Provider string

	Err error
}

// DomainFromNC
func DomainFromNC(d namecheap.DomainGetListResult) (dom Domain) {
	dom = Domain{
		DomainGetListResult: &d,
		Provider:            "namecheap",
	}
	return dom
}

// Nameservers will return the
func (d *Domain) Nameservers() []*net.NS {
	// TODO: an error
	var err error
	d.NS, err = net.LookupNS(d.Name)
	if err != nil {
		log.Warnf("NS lookup error %v", err)
		d.Err = err
		return nil
	}
	return d.NS
}
