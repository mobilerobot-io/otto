package main

import (
	"flag"

	"os"
	"sync"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

//                            Globals
// ====================================================================
var (
	algo    = flag.String("algo", "rusty", "choose the algorigthm to run")
	allowed = flag.String("allowed", "", "add more allowed domains")
	//basedir = flag.String("dir", "/srv/wally", "directory to save crawls in")
	depth = flag.Int("depth", 2, "set the level sub-links are visited")
	pll   = flag.Int("pl", 2, "set the possible parallelism")

	traceout  *os.File
	collector *colly.Collector

	basedir string

	pages sync.Map
	sites sync.Map

	// Queues used to pass data through the pipeline
	urlQ   chan string
	visitQ chan string

	config Configuration
)

//                            init
// ====================================================================

func init() {
	basedir = "/srv/invdb/data"
	urlQ = make(chan string)   // urlQ normalize the url then filter it
	visitQ = make(chan string) // linkQ sends the link off to be visite
}

func walk(url string) {

	// Spin up the function that is going to watch for URLs
	go watchUrlQ(urlQ)
	go watchVisitQ(visitQ)

	// Walk the arguments and visit each of the hosts
	for _, url := range flag.Args() {
		addUrlQ(url)
	}

	log.Infoln("Walking complete, now for output")
	PrintPages(os.Stdout)
}
