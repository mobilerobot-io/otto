package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	namecheap "github.com/billputer/go-namecheap"
	"github.com/mobilerobot-io/otto"
	"github.com/mobilerobot-io/otto/dom"
	log "github.com/sirupsen/logrus"
)

var (
	domains *dom.Domains
	client  *namecheap.Client
)

func main() {
	flag.Parse()

	var args []string
	wr := os.Stdout

	if len(flag.Args()) < 1 {
		args = []string{"domains"}
	} else {
		args = flag.Args()
	}

	switch args[0] {
	case "domains":
		doms := GetDomains()
		if doms == nil {
			log.Infoln("no domains cached, must fetch")
			os.Exit(-1)
		}
		printDomains(wr, doms)
	case "sites":
		//listSites(wr, sites)
	default:
		fmt.Errorf("I do not know what command %s is", args[0])
	}
}

// GetDomains will
func GetDomains() *dom.Domains {
	if domains == nil {
		config := otto.GetConfig()
		if config.Fetch {
			domains = dom.FetchDomains()
		} else {
			domains = dom.GetDomains()
		}
		if domains == nil {
			return nil
		}

		if config.Cache {
			domains.Save()
		}

		for _, dom := range domains.Domains() {
			dom.Nameservers()
		}
	}
	return domains
}

// write domains to writer
func printDomains(wr io.Writer, doms *dom.Domains) {
	fmt.Fprintln(wr, "Domains ... ")
	for n, dom := range doms.Domains() {
		fmt.Fprintf(wr, "%s -> %+v\n", n, dom)
	}
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func fatalIfNil(obj interface{}) {
	if obj == nil {
		log.Fatalln("object is nil")
	}
}
