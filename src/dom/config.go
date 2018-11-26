package dom

import (
	"flag"
)

type Configuration struct {
	Basedir string

	NS bool // Get nameserver records

	Fetch bool // Fetch form namecheap
	Save  bool // Save a local copy

	Format string // text or json

	Verbosity int // 0 = silence, 3 = chatty, log warn &> is always on
}

var (
	config Configuration
)

func init() {
	var err error
	flag.StringVar(&config.Basedir, "dir", "/srv/invdb/", "data directory")

	flag.BoolVar(&config.Fetch, "fetch", false, "fetch from provider")
	flag.BoolVar(&config.Save, "save", false, "save results from run")

	flag.BoolVar(&config.NS, "ns", false, "fetch nameservers")

	flag.IntVar(&config.Verbosity, "verbosity", 3, "verbosity leve 0 = silence, 3 = chatty")
	flag.StringVar(&config.Format, "format", "json", "text, json")
}
