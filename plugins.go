package main

import (
	"fmt"
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

func NewPlugin(p string) OttoPlugin {
	dir, file := filepath.Split(p)
	name := filepath.Base(file)
	fmt.Printf("input: %q\n\tdir: %q\n\tfile: %q\n", p, dir, filepath.Base(file))
	op := OttoPlugin{
		Path:   p,
		Name:   name,
		Loaded: false,
	}
	ottoPlugins[name] = op
	return op
}

func loadPlugins(r *mux.Router, plugins []string) {
	var p []string
	var err error

	log.Debugln("Loading Plugins")
	if config.Plugdir != "" {
		p, err = filepath.Glob("plugins/**/*.so")
		check(err)

		log.Debugln("Plugins...")
		for _, pl := range p {
			log.Debugln("\t", pl)
			NewPlugin(pl)
		}
	}
}

// ActivatePlugin
func ActivatePlugin(path string, r *mux.Router) {

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
	/*
		ottoPlugins[name] = OttoPlugin{
			Name:   name,
			Loaded: true,
			Path:   path,
		}
	*/
	op := ottoPlugins[name]
	op.Loaded = true
	log.Infoln("  subroutes have been registered ", path)
}
