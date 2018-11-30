package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

type Configuration struct {
	Addr  string
	Dir   string
	Fetch bool
	Cache bool

	LogLevel  string
	LogOutput string
	LogFormat string
}

func init() {
	config = Configuration{
		Addr: ":3333",
		Dir:  "/srv/invdb/data/",

		Fetch: false,
		Cache: true,

		LogLevel:  "warn",
		LogOutput: "stdout",
		LogFormat: "json",
	}

	flag.StringVar(&config.LogLevel, "level", "warn", "set log level")
	flag.StringVar(&config.LogFormat, "format", "json", "set log format")
	flag.StringVar(&config.LogOutput, "logfile", "stdout", "logfile, stdout or stderr")

	flag.StringVar(&config.Addr, "addr", ":4433", "address and port to listen on")
	flag.StringVar(&config.Dir, "dir", "/srv/invdb/data", "inventory database")

	flag.BoolVar(&config.Fetch, "fetch", false, "fetch from provider = true, read from cache = false")
	flag.BoolVar(&config.Cache, "cache", true, "cache results from API queries")
}

func (c *Configuration) Save(fname string) (err error) {
	var jbuf []byte

	jbuf, err = json.Marshal(c)
	checkError(err)

	err = ioutil.WriteFile(fname, jbuf, 0644)
	checkError(err)

	return
}

func checkError(err error) {
	if err == nil {
		return
	}
	log.Fatalln(err)
}
