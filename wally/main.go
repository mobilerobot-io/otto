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
	basedir = flag.String("dir", "/srv/wally", "directory to save crawls in")

	depth = flag.Int("depth", -1, "set the level sub-links are visited")
	pll   = flag.Int("pl", 2, "set the possible parallelism")

	serve   = flag.Bool("serve", false, "Run as a service")
	debug   = flag.Bool("d", false, "Turn it up and debug")
	verbose = flag.Bool("v", false, "Turn on verbose printing to stdout")
	logfile = flag.String("logfile", "wallylog.json", "set the log file")

	tracefile = flag.String("trace", "", "turn trace on and write to file")
	traceout  *os.File

	collector *colly.Collector

	pages sync.Map
	sites sync.Map

	// Queues used to pass data through the pipeline
	urlQ   chan string
	visitQ chan string
)

//                            init
// ====================================================================

func init() {
	urlQ = make(chan string)   // urlQ normalize the url then filter it
	visitQ = make(chan string) // linkQ sends the link off to be visite
}

func main() {
	flag.Parse()

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
