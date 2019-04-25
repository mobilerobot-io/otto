package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
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

	// Read the template
	tmpl, err := ioutil.ReadFile("tmpl/index.html")
	check(err)

	t, err := template.New("index").Parse(string(tmpl))
	check(err)

	data := struct {
		Title   string
		Content string
	}{
		Title:   "OttO",
		Content: "The oTTo Maker",
	}

	err = t.Execute(w, data)
	check(err)
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	log.Debug("Entered Route Handler")
	w.WriteHeader(http.StatusOK)
	WalkRoutes(router, w, w)
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

// WalkRoutes will gather all routes we have registered and write the results
// the io.Writer that has been provided us, we will also write all errors to
// the respective io.Writer.
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
