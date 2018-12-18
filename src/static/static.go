package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	Name string
	Dir  string
)

func init() {
	Name = "clowdops.net"
	Dir = "/srv/www/clowdops.net"
}

func Register(name string, sub *mux.Router) {
	// This will serve files under http://localhost:8000/static/<filename>
	sub.PathPrefix("/clowdops.net/").Handler(http.StripPrefix("/clowdops.net/", http.FileServer(http.Dir(Dir))))
}
