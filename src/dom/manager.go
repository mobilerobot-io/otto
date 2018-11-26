package dom

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	namecheap "github.com/billputer/go-namecheap"
	"github.com/mobilerobot-io/otto"
	log "github.com/sirupsen/logrus"
)

type DomainListReader interface {
	Domains() *DomainManager
}

// DomainList represents a list of domains.  We are keeping both the
// list from namecheap for now, along with a list converted into
// our universal format.  Along with an index of ids.
type DomainManager struct {
	nclist []namecheap.DomainGetListResult // namecheap domain list
	dommap map[string]Domain               // domain
	ids    map[int]*Domain                 // id map
}

// GetDomains will get the local cache if it exists,
func GetDomains() (doms *DomainManager) {
	// These are stored in as the origin response from namecheap
	// wrapped in our own structure adding a little meta data to
	// give the object some context
	if doms = readDomains(); doms == nil {
		doms = FetchDomains()
	}
	return doms
}

func FetchDomains() (doms *DomainManager) {
	return ncDomains()
}

func (doms *DomainManager) Domain(name string) (dom Domain, found bool) {
	fatalIfNil(doms.dommap)
	if dom, ex := doms.dommap[name]; ex {
		return dom, ex
	}
	return dom, true
}

func (doms *DomainManager) Domains() map[string]Domain {
	if doms.dommap == nil {
		doms.dommap = make(map[string]Domain)
		doms.ids = make(map[int]*Domain)
	}
	for _, ncdom := range doms.nclist {
		d := Domain{
			DomainGetListResult: ncdom,
		}
		doms.dommap[ncdom.Name] = d
		doms.ids[ncdom.ID] = &d
	}
	return doms.dommap
}

// Save will save the cached namecheap response, we can build indexes
// from the original response.
func (doms *DomainManager) Save() error {
	config := otto.GetConfig()
	fname := config.Basedir + "/data/domains.json"

	if len(doms.dommap) < 1 {
		log.Infoln("No domains to save returning")
		return nil
	}

	log.Infof("Saving %d domains to %s\n", len(doms.dommap), fname)
	jbuf, err := json.Marshal(doms.dommap)
	fatalIfError(err)

	fi, err := os.Create(fname)
	fatalIfError(err)

	_, err = fi.Write(jbuf)
	fatalIfError(err)

	log.Infoln("\tdone, domains saved")
	return nil
}

func (doms *DomainManager) Output(w io.Writer) {
	doms.Text(w)
}

func (doms *DomainManager) Text(w io.Writer) {
	fmt.Fprintf(w, "Domains...\n")
	for _, domain := range doms.dommap {
		fmt.Fprintf(w, "\t%s\t%d\n", domain.Name, domain.ID)
	}
}

// readDomains gets our domain list from a saved file somewhere
func readDomains() *DomainManager {
	config := otto.GetConfig()

	fname := config.Basedir + "/data/domains.json"

	log.Infof("Reading domains %s ...\n", fname)
	jbuf, err := ioutil.ReadFile(fname)
	if err != nil {
		// it is ok if we do not have a domains file ...
		log.Infoln(err)
		return nil
	}

	var domains DomainManager
	err = json.Unmarshal(jbuf, &domains)
	fatalIfError(err)
	return &domains
}

func fatalIfError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func fatalIfNil(obj interface{}) {
	if obj == nil {
		log.Fatalln("obj nil")
	}
}
