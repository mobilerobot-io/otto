package api

import "github.com/gorilla/mux"

type Service interface {
	Name() string
	Register(path string, handler *mux.Router)
}
