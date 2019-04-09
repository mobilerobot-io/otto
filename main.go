package main

import (
	"flag"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/mobilerobot-io/otto"
)

type Service struct {
	Name string
	Addr string
	Path string
	*http.Server
}

// All global variables here
var (
	config  Configuration
	service *Service
)

// Everything on the command line should be a plugin
func main() {
	flag.Parse()

	srv := otto.NewServer(Configuration.Addr)

	loadPlugins(flag.Args)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
