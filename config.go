package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

type Configuration struct {
	Addrport string
	Dir      string
	Fetch    bool
	Cache    bool
	Plugdir  string // Directory that contains plugin dirs

	ListPlugins bool
	ListRoutes  bool

	LogLevel  string
	LogOutput string
	LogFormat string
}

func init() {
	config = Configuration{
		Addrport:  ":3333",
		LogLevel:  "warn",
		LogOutput: "stdout",
		LogFormat: "json",
	}

	flag.StringVar(&config.LogLevel, "level", "warn", "set log level")
	flag.StringVar(&config.LogFormat, "format", "json", "set log format")
	flag.StringVar(&config.LogOutput, "logfile", "stdout", "logfile, stdout or stderr")
	flag.StringVar(&config.Plugdir, "plugdir", "plugins", "the dir to look for plugins")

	flag.StringVar(&config.Addrport, "addr", ":4433", "address and port to listen on")
	flag.BoolVar(&config.ListRoutes, "list-routes", true, "Walk the routes after they have been added")
	flag.BoolVar(&config.ListPlugins, "list-plugins", true, "List the plugins we are using")
}

// Save we can start using store for this, correct?
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
