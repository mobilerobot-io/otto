package main

import (
	"flag"
	"fmt"
	"os"

	namecheap "github.com/rustyeddy/go-namecheap"
)

var (
	domains Domains
	client  *namecheap.Client
)

func main() {
	flag.Parse()

	// First try, read domains and write them
	domains = GetDomains()
	if config.Verbosity > 0 {
		fmt.Printf("Got %d domains\n", len(domains.Domains))
	}
	if config.Save {
		domains.Save()
	}

	if config.Verbosity == 0 {
		os.Exit(0) // all as well nothing to output
	}

	for _, dom := range domains.Domains {
		fmt.Printf("%d %s ns: ", dom.ID, dom.Name)
		if ns := dom.Nameservers(); ns != nil && len(ns) > 0 {
			for _, h := range ns {
				fmt.Printf("%s ", h.Host)
			}
		}
		fmt.Println("")
	}
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
