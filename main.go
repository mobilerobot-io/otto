package main

import (
	"flag"
	"net/http"
	"plugin"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

// All global variables here
var (
	config Configuration
)

// Everything on the command line should be a plugin
func main() {
	flag.Parse()

	r := mux.NewRouter()
	srv := &http.Server{
		Handler: r,
		Addr:    config.Addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	for _, arg := range flag.Args() {

		log.Infoln("  new plugin ", arg)
		pl, err := plugin.Open(arg)
		check(err)

		n, err := pl.Lookup("Name")
		check(err)

		// Determine the name and path for our new subroute
		name := *n.(*string)
		path := "/" + name
		if name == "static" || name == "clowdops.net" {
			path = "/"
		}

		log.Infof("   name %s path %s", name, path)

		// Create our new subroutee
		s := r.PathPrefix(path).Subrouter()

		log.Infoln("  subrouter created ", path)

		// Get the Register functions symbol from our plugin and register
		regf, err := pl.Lookup("Register")
		check(err)

		// Now register our plugin by passing the newly created
		// subrouter to the new plugin's Register(*mux.Router) function
		regf.(func(string, *mux.Router))(name, s)
		log.Infoln("  subroutes have been registered ", path)
	}

	log.Println("  otto is starting on ", srv.Addr)
	err := srv.ListenAndServe()

	log.Println("Good bye... ", err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
