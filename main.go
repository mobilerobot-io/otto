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
	// the router.  This will cause all of our builtin REST commands to be
	// registered and our Websocket handler will be created.
	server, router = NewServer(config.Addrport)

	// Now we will load up our plugins
	loadPlugins(config.Plugdir)

	// list the plugins is a command line arg requests so
	if config.ListPlugins {
		log.Infoln("Plugins available ")
		for n, _ := range ottoPlugins {
			log.Infoln("\t" + n)
		}
	}

	// Treat the remaining arguments as plugins that must be
	for _, p := range flag.Args() {
		activatePlugin(p, router)
	}

	if config.ListRoutes {
		log.Println("Registered routes: ")
		WalkRoutes(router, os.Stdout, os.Stderr)
	}

	// Now we'll start the server if we have been configured to
	// run in daemonic modez
	if config.NoDaemon {
		os.Exit(0)
	}

	go mqtt_run()

	// Listen for and handler HTTP HTML, REST and Websocket requests
	log.Infoln("  otto is starting on ", server.Addr)
	err := server.ListenAndServe()
	log.Fatal(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
