package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	namecheap "github.com/billputer/go-namecheap"
	"github.com/mobilerobot-io/otto"
	"github.com/mobilerobot-io/otto/dom"
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
		listDomains(wr, doms)
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
		fatalIfNil(domains)

		if config.Cache {
			domains.Save()
		}

		for _, dom := range domains.Domains() {
			if ns := dom.Nameservers(); ns != nil && len(ns) > 0 {
				for _, h := range ns {
					fmt.Printf("ns %+v", h)
				}
			}
		}
	}
	return domains
}

// write domains to writer
func listDomains(wr io.Writer, doms *dom.Domains) {
	for _, dom := range doms.Domains() {
		fmt.Fprintln(wr, dom.String()) // dom.String()
	}
}

/*
func GetSites() (s *Sites) {
	return s
}

func ListSites(wr io.Writer, sites []*Sites) {
	for _, site := range sites {
		fmt.Fprintln(wr, site.String())
	}
}
*/
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
