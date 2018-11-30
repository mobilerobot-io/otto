package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	Name    string = "echo"
	eRouter *mux.Router
)

// AddService is used by otto to add this service.  We expect
// the caller to provide us with a router (subrouter)
func Register(s *mux.Router) {
	s.HandleFunc("/", Handler)
	s.HandleFunc("/{str}", Handler)
	log.Infoln("  echo was registered... ")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Infoln("I got called by echo service")

	var str string
	var ex bool

	vars := mux.Vars(r)
	log.Printf("%+v", vars)
	if str, ex = vars["str"]; ex {
		n, err := w.Write([]byte(str))
		check(err)

		log.Infof("  responded, wrote %d bytes", n)
	} else {
		w.Write([]byte(r.URL.String()))
	}
}

func Path() string {
	return "/" + Name
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
