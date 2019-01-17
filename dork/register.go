package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

var Name = "do"

func Register(name string, s *mux.Router) {
	s.HandleFunc("/", DoHandler).Name("dork")
	//s.HandleFunc("/{str}", Handler).Name("echo")
	log.Infoln("  digital ocean was registered... ")
}

func DoHandler(w http.ResponseWriter, r *http.Request) {
	log.Infoln("I got called by echo service")

	displayAccount(w)
	displayProjects(w)
	// displayVMs(w)
	// displayVolumes(w)
	displayCDN(w)
	displayActions(w)

}
