package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"plugin"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type Service struct {
	Name string
	Addr string
	Path string
	*http.Server
}

// All global variables here
var (
	config  Configuration
	service *Service
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
	/*
		service := Service{
			Name:   "otto",
			Addr:   ":4422",
			Path:   "/otto",
			Server: srv,
		}
	*/

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
		sub := r.PathPrefix(path).Subrouter()
		log.Infoln("  subrouter created ", path)

		// Get the Register functions symbol from our plugin and register
		regf, err := pl.Lookup("Register")
		check(err)

		// Now register our plugin by passing the newly created
		// subrouter to the new plugin's Register(*mux.Router) function
		regf.(func(string, *mux.Router))(name, sub)
		log.Infoln("  subroutes have been registered ", path)
	}
	WalkRoutes(r, os.Stdout, os.Stderr)

	log.Println("  otto is starting on ", srv.Addr)
	err := srv.ListenAndServe()
	log.Printf("Good bye...%v ", err)
}

func WalkRoutes(r *mux.Router, w io.Writer, e io.Writer) {
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Fprintln(w, "ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Fprintln(w, "Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Fprintln(w, "Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Fprintln(w, "Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Fprintln(w, "Methods:", strings.Join(methods, ","))
		}
		fmt.Fprintln(w)
		return nil
	})
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
