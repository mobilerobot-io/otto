package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	service api.Service
)

func init() {
	Service = EchoService{
		name: "echo",
	}
}

// AddService is used by otto to add this service.  We expect
// the caller to provide us with a router (subrouter)
func (es *EchoService) Register(r *mux.Router) {
	r.HandleFunc("/echo", es.Handler)
	log.Infoln("  echo was registered... ")
}

func (es *EchoService) Handler(w http.ResponseWriter, r *http.Request) {
	log.Infoln("I got called by echo service")
	w.Write([]byte(r.URL.String()))
}

func (es *EchoService) Name() string {
	return es.name
}

func (es *EchoService) Path() string {
	return "/" + es.name
}
