package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	*http.Server
	*mux.Router
}

func NewService(addr string) (s *Service) {
	r := mux.NewRouter()
	s = &Service{
		Router: r,
		Server: &http.Server{
			Addr:           addr,
			Handler:        r,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}

	// Now start adding handlers
	r.HandleFunc("/", s.homeHandler)
	return s
}

func (s *Service) AddService(svc Service) {
	sub := s.subrouter(svc.Path())
	svc.Register(sub)
}

func (s *Service) Start() {
	log.Infoln("Starting service at ", s.Addr)
	err := s.ListenAndServe()
	checkError(err)
}

func (s *Service) subrouter(name string) *mux.Router {
	return s.Router.PathPrefix("/" + name).Subrouter()
}

func (s *Service) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home sweet home")
}

func (s *Service) WriteJSON(w http.ResponseWriter, obj interface{}) {
	jb, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), 501)
	}
	w.Write(jb)
}

func (s *Service) Walk() {
	r := s.Router
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

func checkError(err error) error {
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}
