package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func NewServer(addr string) *http.Server {

	r := mux.NewRouter()
	srv := &http.Server{
		Handler: r,
		Addr:    config.Addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

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
