/*
Otto is a simple tool to help keep track of the online inventory
we are managing.  Monitors inventory, and launches tools like terraform
and ansible to ensure our infrastructure is operating correctly.
*/
package otto

import "github.com/gorilla/mux"

type Service interface {
	Name() string
	Path() string
	Register(sub *mux.Router)
}
