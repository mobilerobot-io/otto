package main

import (
	"flag"

	hd "github.com/mitchellh/go-homedir"
)

type Configuration struct {
	Creds   string // digi
	Basedir string
	Output  string
	Format  string
}

func init() {
	hdir, err := hd.Dir()
	panicIfError(err)

	docreds := hdir + "/.config/digitalocean/creds.json"
	flag.StringVar(&config.Creds, "creds", docreds, "point to digital ocean creds file")
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
