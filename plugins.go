package main

import (
	"flag"
	"net/http"
	"os"
	"path/filepath"
	"plugin"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Mods is just another name for plugins that will not conflict with
// the standard package plugins
type Mod struct {
	Name   string
	Loaded bool
	plugin.Plugin
}

func loadPlugins(s *http.Server, r *mux.Router, plugins []string) {
	if config.Plugins != "" {
		plugins, err := filepath.Glob(config.Plugins)
		check(err)

		log.Info("Plugins...")
		for _, pl := range plugins {
			log.Infoln("\t", pl)
		}
		//os.Exit(0)
	}

	if flag.Args() == 0 {
		return
	}

	for _, name := range flag.Args() {
		doPlugin(name, r)
	}

	if config.Routes {
		WalkRoutes(r, os.Stdout, os.Stderr)
	}

	log.Println("  otto is starting on ", server.Addr)
	err := s.ListenAndServe()
	log.Printf("Good bye...%v ", err)
}

func doPlugin(path string, r *mux.Router) {

	log.Infoln("  new plugin ", path)
	//path := name + "/" + name + ".so"
	pl, err := plugin.Open(path)
	check(err)

	// TODO: Add Flags and Help ...
	n, err := pl.Lookup("Name")
	check(err)

	// Determine the name and path for the new subroute
	name := *n.(*string)
	url := "/" + name
	if name == "static" || name == "clowdops.net" {
		url = "/"
	}

	log.Infof("   name %s path %s url %s ", name, path, url)

	// Create our new subroute
	sub := r.PathPrefix(url).Subrouter()
	log.Infoln("  subrouter created ", url)

	// Get the Register functions symbol from our plugin and register
	regf, err := pl.Lookup("Register")
	check(err)

	// Now register our plugin by passing the newly created
	// subrouter to the new plugin's Register(*mux.Router) function
	regf.(func(string, *mux.Router))(name, sub)
	log.Infoln("  subroutes have been registered ", path)
}
