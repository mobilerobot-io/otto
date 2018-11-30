package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	Name string = "store"
	Path string = "/store"
	Dir  string = "/srv/invdb"
)

// Register our callback handler
func Register(s *mux.Router) {
	s.HandleFunc("/", StoreHandler)
	s.HandleFunc("/{name}", StoreItemHandler).Methods("GET", "POST", "PUT", "DELETE")
	//s.HandleFunc("/{name}", StorePutHandler).Methods("POST", "PUT")
	//s.HandleFunc("/{name}", StorePutHandler).Methods("DELETE")
	log.Infoln("  store was registered ... ")
}

// StoreHandler will return general information about our storage
func StoreHandler(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, Dir)
}

// StoreItemHandler will get, create, modify or delete the named object
func StoreItemHandler(w http.ResponseWriter, r *http.Request) {

	var obj interface{}
	//vars := mux.Vars(r)
	//name := vars["name"]
	switch r.Method {
	case "GET":
	case "PUT", "POST":
	case "DELETE":
	default:
	}
	WriteJSON(w, obj)
}

func ReadJSON(rd io.Reader, obj interface{}) (b []byte, err error) {
	_, err = rd.Read(b)
	check(err)

	err = json.Unmarshal(b, obj)
	check(err)

	return b, err
}

func WriteJSON(wr io.Writer, obj interface{}) (err error) {
	var jbytes []byte
	jbytes, err = json.Marshal(obj)
	check(err)

	nbytes, err := wr.Write(jbytes)
	check(err)

	log.Infoln("saved json ", string(nbytes))
	return
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
