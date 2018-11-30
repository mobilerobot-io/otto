package main

import (
	"flag"
	"os"
	"runtime/trace"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

// Hide some noisy, boring but important details getting all the
// configuration options correct, and from all the right places:
// a) config file, b) env variables, c) cmd line options.
type Configuration struct {
	Algorithm string // what algo to run on the coming scrape
	Allow     string // comma separated list of hosts to allow
	Basedir   string // working directory for the applicatio

	Debug       bool
	Depth       int // Crawl Depth (default to 1 to be nice)
	Parallelism int // Level of Parallelism

	Serve  bool // Start and run as a service accepting http requests
	Client bool // Run as a command and print to stdout

	Logfile   string // Name of the logfile
	Tracefile string // Name of trace file, if file empty no trace is run
}

type Globals struct {

	// The trace io.Writer if we are tracing, nil means no trace
	traceout *os.File // tracefile

	// Our collector
	colector *colly.Collector
}

func configed() {
	flag.Parse()

	var err error
	// Create our storage directory if it does not already exist
	if _, err = os.Stat(basedir); err != nil {
		if err = os.MkdirAll(basedir, 0644); err != nil {
			panic("could not create base directory")
		}
	}

	// Open the log file and set it for logging
	o := os.Stdout
	if config.Logfile != "" {
		o, err = os.OpenFile(config.Logfile, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
			//log.Error("failed to write log file ", logfile)
			//io = os.Stdout
		}
	}
	// Setup the logger and other outputs
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	log.SetOutput(o)

	// Set the debug level if we have turned debugging on
	if config.Debug {
		log.SetLevel(log.DebugLevel)
	}

	if config.Tracefile != "" {
		traceout, err = os.Create(config.Tracefile)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := traceout.Close(); err != nil {
				panic(err)
			}
		}()
	}

	if err := trace.Start(traceout); err != nil {
		panic(err)
	}
	defer trace.Stop()
}
