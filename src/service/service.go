package service

import (
	"fmt"
	"net/http"
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

func (s *Service) Start() {

	log.Infoln("Starting service at ", s.Addr)
	err := s.ListenAndServe()
	checkError(err)
}

func (s *Service) Subrouter(name string) *mux.Router {
	return s.Router.PathPrefix("/" + name).Subrouter()
}

func (s *Service) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home sweet home")
}

func checkError(err error) error {
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}
