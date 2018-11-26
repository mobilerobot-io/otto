package otto

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

type Configuration struct {
	LogLevel  string
	LogOutput string
	LogFormat string

	Basedir string // base directory for saving data
	Fetch   bool
	Cache   bool // cache input from api calls? Default = yes
}

var (
	config Configuration
)

func init() {
	flag.StringVar(&config.LogLevel, "level", "warn", "set log level")
	flag.StringVar(&config.LogFormat, "format", "json", "set log format")
	flag.StringVar(&config.LogOutput, "logfile", "stdout", "logfile, stdout or stderr")

	flag.StringVar(&config.Basedir, "dir", "/srv/invdb", "inventory database")

	flag.BoolVar(&config.Fetch, "fetch", false, "fetch from provider = true, read from cache = false")
	flag.BoolVar(&config.Cache, "cache", true, "cache results from API queries")
}

func GetConfig() *Configuration {
	return &config
}

func (c *Configuration) Save() (err error) {
	var jbuf []byte

	jbuf, err = json.Marshal(c)
	fatalIfError(err)

	err = ioutil.WriteFile(c.Basedir+"/config.json", jbuf, 0644)
	fatalIfError(err)

	return
}

func fatalIfError(err error) {
	if err == nil {
		return
	}
	log.Fatalln(err)
}
