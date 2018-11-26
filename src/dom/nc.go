package dom

import (
	"encoding/json"
	"io/ioutil"

	namecheap "github.com/billputer/go-namecheap"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
)

type nccreds struct {
	Name, Token, User string
}

// get client will get the creds and create a client
func getClient() (ncli *namecheap.Client) {
	if ncli == nil {
		u, t, n := creds()
		ncli = namecheap.NewClient(u, t, n)
		if ncli == nil {
			log.Fatalln("no client")
		}
	}
	return ncli
}

// return the Namecheap Token string
func creds() (u, t, v string) {

	fname, err := homedir.Dir()
	fatalIfError(err)

	credfile := fname + "/" + ".config/namecheap/creds.json"
	b, err := ioutil.ReadFile(credfile)
	fatalIfError(err)

	var creds nccreds
	err = json.Unmarshal(b, &creds)
	fatalIfError(err)

	// convert bytes to string
	return creds.Name, creds.Token, creds.User
}

// Fetch domains
// ====================================================================

// fetchDomains will grab our domains from the provider,
// namecheap in our case
func ncDomains() *DomainManager {
	// Get a list of your domains
	cli := getClient()
	ncdoms, err := cli.DomainsGetList()
	fatalIfError(err)

	dlist := &DomainManager{
		nclist: ncdoms,
		dommap: make(map[string]Domain),
		ids:    make(map[int]*Domain),
	}

	// Create a list of otto.DomainManager form namecheap.Domains
	for _, ncdom := range ncdoms {
		dom, err := DomainFromNC(ncdom)
		fatalIfError(err)

		dlist.dommap[dom.Name] = dom
		dlist.ids[dom.ID] = &dom
	}
	return dlist
}
