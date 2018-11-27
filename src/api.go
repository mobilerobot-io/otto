package otto

import "github.com/gorilla/mux"

type Service interface {
	Name() string
	Path() string
	Register(sub *mux.Router)
}
