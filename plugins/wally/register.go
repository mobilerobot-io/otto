package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

var Name = "walk"

func Register(name string, s *mux.Router) {
	//s.HandleFunc("/", WalkHandler)
	s.HandleFunc("/{url}", WalkHandler).Name("walk")
	log.Infoln("  digital ocean was registered... ")
}

func WalkHandler(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Walk handler at your service ... ")

	vars := mux.Vars(r)
	url := vars["url"]
	walk(url)
}
