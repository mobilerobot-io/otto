package main

import (
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
	Dir    string
	Active bool
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
	op := OttoPlugin{
		Path:   p,
		Name:   name,
		Dir:    dir,
		Active: false,
	}
	ottoPlugins[name] = op
	return op
}

func loadPlugins(dir string) {
	var p []string
	var err error

	log.Infoln("Loading Plugins from dir: " + dir)
	if dir == "" {
		dir = "."
	}

	// TODO add a check for the same directory and one below
	// for deployment scenarios
	p, err = filepath.Glob(dir + "/plugins/*/*.so")
	check(err)

	for _, pl := range p {
		log.Infoln("\t", pl)
		NewPlugin(pl)
	}
}

// ActivatePlugin
func activatePlugin(fname string, r *mux.Router) {

	p, ex := ottoPlugins[fname]
	if !ex {
		return
	}
	log.Infoln("  New plugin ", p.Path)
	pl, err := plugin.Open(p.Path)
	check(err)

	// TODO: Add Flags and Help ...
	n, err := pl.Lookup("Name")
	check(err)

	// Determine the name and path for the new subroute
	name := *n.(*string)
	url := "/" + name
	if name == "root" {
		url = "/"
	}
	log.Infof("   name %s path %s url %s ", name, p.Path, url)

	// Create our new subroute
	sub := r.PathPrefix(url).Subrouter()
	log.Infoln("  subrouter created ", url)

	// Get the Register functions symbol from our plugin and register
	regf, err := pl.Lookup("Register")
	check(err)

	// Now register our plugin by passing the newly created
	// subrouter to the new plugin's Register(*mux.Router) function
	regf.(func(string, *mux.Router))(name, sub)
	op, ex := ottoPlugins[name]
	if !ex {
		ottoPlugins[name] = OttoPlugin{
			Name:   name,
			Active: true,
			Path:   fname,
		}
		op = ottoPlugins[name]
	}
	op.Active = true
	log.Infof("\tplug %s activated", name)
}
