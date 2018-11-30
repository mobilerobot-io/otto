package main

import (
	"flag"
	"log"
	"net/http"
	"plugin"
	"time"

	"github.com/gorilla/mux"
)

// Everything on the command line should be a plugin
func main() {
	flag.Parse()

	r := mux.NewRouter()

	// This will serve files under http://localhost:8000/static/<filename>
	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	srv := &http.Server{
		Handler: r,
		Addr:    ":8777",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	for _, arg := range flag.Args() {

		pl, err := plugin.Open(arg)
		check(err)

		n, err := pl.Lookup("Name")
		check(err)

		// Determine the name and path for our new subroute
		name := *n.(*string)
		path := "/" + name

		// Create our new subroutee
		s := r.PathPrefix(path).Subrouter()

		// Get the Register functions symbol from our plugin and register
		regsym, err := pl.Lookup("Register")
		check(err)

		// Now register our plugin
		reg := regsym.(func(*mux.Router))
		reg(s)
	}

	log.Println("  otto is starting on ", srv.Addr)
	srv.ListenAndServe()

	log.Println("Good bye")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
