package dom

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/rustyeddy/go-namecheap"
	log "github.com/sirupsen/logrus"
)

// Domains represents a list of domains.  We are keeping both the
// list from namecheap for now, along with a list converted into
// our universal format.  Along with an index of ids.
type Domains struct {
	NClist  []namecheap.DomainGetListResult // namecheap domain list
	Domains map[string]*Domain              // domain
	ids     map[int]*Domain                 // id map
}

// GetDomains will get the local cache if it exists,
func GetDomains() (doms Domains) {

	// These are stored in as the origin response from namecheap
	// wrapped in our own structure adding a little meta data to
	// give the object some context
	if config.Fetch {
		doms = ncFetchDomains()
	} else {
		doms = readDomains()
	}
	return doms
}

func (doms *Domains) Index() *Domains {
	if doms.Domains == nil {
		doms.Domains = make(map[string]*Domain)
		doms.ids = make(map[int]*Domain)
	}
	for _, ncdom := range doms.NClist {
		d := new(Domain)
		d.DomainGetListResult = &ncdom
		doms.Domains[ncdom.Name] = d
		doms.ids[ncdom.ID] = d
	}
	return doms
}

// Save will save the cached namecheap response, we can build indexes
// from the original response.
func (doms *Domains) Save() error {
	fname := config.Basedir + "data/domains.json"

	log.Infof("Saving %d domains to %s\n", len(doms.Domains), fname)
	jbuf, err := json.Marshal(doms.Domains)
	if err != nil {
		panic(err)
	}

	if fi, err := os.Create(fname); err != nil {
		panic(err)
	} else {
		if _, err := fi.Write(jbuf); err != nil {
			panic(err)
		}
	}
	log.Infoln("\tdone, domains saved")
	return nil
}

func (doms *Domains) Output(w io.Writer) {
	doms.Text(w)
}

func (doms *Domains) Text(w io.Writer) {
	fmt.Fprintf(w, "Domains...\n")
	for _, domain := range domains.Domains {
		fmt.Fprintf(w, "\t%s\t%d\n", domain.Name, domain.ID)
	}
}

// Fetch and read domains
// ====================================================================

// fetchDomains will grab our domains from the provider,
// namecheap in our case
func ncFetchDomains() (dl Domains) {
	// Get a list of your domains

	client = getClient()
	domains, err := client.DomainsGetList()
	panicIfError(err)

	dl.Domains = make(map[string]*Domain)
	dl.ids = make(map[int]*Domain)
	dl.NClist = domains // save the original namecheap response

	// Wrap NC domains into our domains (we will save both)
	for _, ncdom := range domains {
		dom := DomainFromNC(ncdom)
		dl.Domains[dom.Name] = &dom
		dl.ids[dom.ID] = &dom
	}
	return dl
}

// readDomains gets our domain list from a saved file somewhere
func readDomains() (domains Domains) {
	fname := config.Basedir + "data/domains.json"
	if jbuf, err := ioutil.ReadFile(fname); err != nil {
		panic(err)
	} else {
		if err := json.Unmarshal(jbuf, &domains); err != nil {
			panic(err)
		}
	}
	return domains
}
