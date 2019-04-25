package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// All global variables here
var (
	config Configuration
	server *http.Server
	router *mux.Router
)

func init() {
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

// Everything on the command line should be a plugin
func main() {

	// flags are set in config.go, parse em and get our command line args ready
	flag.Parse()

	// Create the server along with the router, our plugins will register with
	// the router.
	server, router = NewServer(config.Addrport)

	// Now we will load up our plugins
	loadPlugins(router, flag.Args())
	if config.ListPlugins {
		for n, _ := range ottoPlugins {
			log.Infoln(n)
		}
	}

	if config.ListRoutes {
		log.Println("Registered routes: ")
		WalkRoutes(router, os.Stdout, os.Stderr)
	}

	// Now we'll start the server if we have been configured to
	// run in daemonic modez
	if config.Daemon {
		log.Infoln("  otto is starting on ", server.Addr)
		err := server.ListenAndServe()
		log.Fatal(err)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
