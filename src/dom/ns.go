package dom

import (
	"net"

	nc "github.com/billputer/go-namecheap"
	log "github.com/sirupsen/logrus"
)

// addNameservers appends a new nameserver to the name server list
func (d *Domain) addNameserver(name string) {
	ns := &net.NS{
		Host: name,
	}
	d.ns = append(d.ns, ns)
}

// DNS Lookups
// ========================================================================

// dnsNameservers will return the name servers via DNS query
func (d *Domain) Nameservers() []*net.NS {
	// TODO: an error
	var err error
	d.ns, err = net.LookupNS(d.Name)
	if err != nil {
		log.Warnf("NS lookup error %v", err)
		d.err = err
		return nil
	}
	return d.ns
}

// Namecheap Lookups
// ========================================================================
func (d *Domain) GetHosts() *nc.DomainDNSGetHostsResult {
	ncli := getClient()

	hosts, err := ncli.DomainsDNSGetHosts(d.sld, d.tld)
	fatalIfNil(err)
	log.Fatalf("%+v\n", hosts)
	return hosts
}
