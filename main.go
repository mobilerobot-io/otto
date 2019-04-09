package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// All global variables here
var (
	config Configuration
	server *http.Server
	router *mux.Router
)

// Everything on the command line should be a plugin
func main() {
	flag.Parse()

	// set the return values to the corresponding globals
	server, router := NewServer(config.Addrport)
	loadPlugins(server, router, flag.Args())
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}