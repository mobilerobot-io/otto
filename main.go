package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	namecheap "github.com/rustyeddy/go-namecheap"
	"github.com/mobilerobot-io/otto/dom"
	"github.com/mobilerobot-io/otto/service"
	log "github.com/sirupsen/logrus"
)

var (
	domains *dom.DomainManager

	cli    *namecheap.Client
	srv    *service.Service
	config Configuration
)

func main() {
	flag.Parse()

	// If there is anything on the command line we will treat it
	// as a one time command.  Once the command is finished this
	// program will exit
	if len(flag.Args()) > 0 {
		docmd(flag.Args())
		os.Exit(0)
	}

	srv := service.NewService(config.Addr)
	srv.AddService(dom.NewService())
	//srv.AddService(site.NewService())
	srv.Walk()
	srv.Start()
}

func docmd(args []string) {
	// we have at least one argument
	argc := len(args)
	if argc > 1 {

		switch flag.Arg(0) {
		case "domains":
			doms := GetDomains()
			if doms == nil {
				log.Infoln("no domains cached, must fetch")
				os.Exit(-1)
			}
			printDomains(os.Stdout, doms)
		default:
			fmt.Errorf("I do not know what command %s is", args[0])
		}
	}
}

// GetDomains will
func GetDomains() *dom.DomainManager {
	if domains == nil {
		if config.Fetch {
			domains = dom.FetchDomains()
		} else {
			domains = dom.GetDomains()
		}
		if domains == nil {
			return nil
		}

		if config.Cache {
			domains.Save(config.Dir + "/domains.json")
		}

		for _, dom := range domains.Domains() {
			dom.Nameservers()
		}
	}
	return domains
}

// write domains to writer
func printDomains(wr io.Writer, doms *dom.DomainManager) {
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
