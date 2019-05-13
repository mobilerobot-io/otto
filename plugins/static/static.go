package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
)

type Configuration struct {
	Name   string
	Dir    string
	Prefix string
}

var (
	Name   string = "static"
	config Configuration
)

func init() {
	config = Configuration{
		Name:   "static",
		Dir:    ".",
		Prefix: "static",
	}
	flag.StringVar(&config.Dir, "static", ".", "Directory to serve up")
}

func Register(name string, sub *mux.Router) {

	// This will serve files under http://localhost:8000/static/<filename>
	sub.PathPrefix(name).Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(config.Dir))))
}
