package dom

import (
	"net/http"

	"github.com/gorilla/mux"
)

type DomainService struct {
	service.Service
	name string
	path string

	*mux.Router
}

func NewService() *DomainService {
	return &DomainService{
		name: "Domains", // name of service
		path: "domains", // root path after service root
	}
}

func (s *DomainService) Name() string {
	return s.name
}

func (s *DomainService) Path() string {
	return s.path
}

func (s *DomainService) Register(sub *mux.Router) {
	s.Router = sub
	sub.HandleFunc("/domain", s.domListHandler)
}

func (s *DomainService) domListHandler(w http.ResponseWriter, r *http.Request) {
	doms := GetDomains()
	domains := doms.Domains()
	s.WriteJSON(w, domains)
}
