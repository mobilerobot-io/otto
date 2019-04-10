package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// NewServer will create the single global http.server and mux.Router, it will
// also create some built in routes we will be using
func NewServer(addr string) (s *http.Server, r *mux.Router) {
	// Create the router and server object to house the router
	r = mux.NewRouter()
	s = &http.Server{
		Handler: r,
		Addr:    config.Addrport,
		// Good practice: enforce timeouts for servers youcreate!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.HandleFunc("/", OttoHandler)
	r.HandleFunc("/routes", routeHandler)
	r.HandleFunc("/plugins", pluginHandler)
	return s, r
}

func OttoHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	log.Debug("Entered Otto Handler")
	w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, "Category: %v\n", vars["category"])
	fmt.Fprintf(w, "Wee, I Go!")
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	log.Debug("Entered Otto Handler")
	w.WriteHeader(http.StatusOK)
	WalkRoutes(router, w, w)
	//fmt.Fprintf(w, "Wee, I Go!")
}

func pluginHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	log.Debug("Entered Plugin Handler")
	w.WriteHeader(http.StatusOK)

	for n, _ := range ottoPlugins {
		fmt.Fprintf(w, n)
	}
	fmt.Fprintf(w, "done")
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
