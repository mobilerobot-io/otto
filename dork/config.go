package main

import (
	"flag"

	hd "github.com/mitchellh/go-homedir"
)

type Configuration struct {
	Creds   string // digi
	Basedir string

	// checks and stuff
	Account  bool
	Actions  bool
	CDNs     bool
	Projects bool

	Output string
	Format string
}

func init() {
	hdir, err := hd.Dir()
	panicIfError(err)

	docreds := hdir + "/.config/digitalocean/creds.json"
	flag.StringVar(&config.Creds, "creds", docreds, "point to digital ocean creds file")
	flag.StringVar(&config.Basedir, "dir", "/srv/invdb/data/", "set the invdb directory")

	flag.BoolVar(&config.Account, "account", false, "check account")
	flag.BoolVar(&config.Actions, "actions", false, "check and display actions")
	flag.BoolVar(&config.CDNs, "cdns", false, "display the CDNs we have")
	flag.BoolVar(&config.Projects, "proj", false, "display projects")
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func panicIfNil(obj interface{}) {
	if obj == nil {
		panic(obj)
	}
}

func okCheck(ok bool, msg string) bool {
	if !ok {
		//log.Warnln("Everyhing is NOT OK ", msg)
		panic(msg)
	}
	return ok
}
