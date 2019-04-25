package main

import (
	"net/http"
	"path/filepath"
	"plugin"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Mods is just another name for plugins that will not conflict with
// the standard package plugins
type OttoPlugin struct {
	Name   string
	Path   string
	Loaded bool
}

var (
	ottoPlugins map[string]OttoPlugin
)

func init() {
	ottoPlugins = make(map[string]OttoPlugin)
}

func loadPlugins(s *http.Server, r *mux.Router, plugins []string) {
	var p []string
	var err error
	if config.Plugdir != "" {
		p, err = filepath.Glob(config.Plugdir + "/*/*.so")
		check(err)

		log.Info("Plugins...")
		for _, pl := range plugins {
			log.Infoln("\t", pl)
		}
	}

	if plugins == nil {
		log.Debug("We appear to have no plugins returning ")
		return
	} else {
		p = plugins
	}
	for _, name := range p {
		log.Debugln("Doing plugin ", name)
		doPlugin(name, r)
	}
}

func doPlugin(path string, r *mux.Router) {

	log.Infoln("  New plugin ", path)
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

	ottoPlugins[name] = OttoPlugin{
		Name:   name,
		Loaded: true,
		Path:   path,
	}
	log.Infoln("  subroutes have been registered ", path)
}
