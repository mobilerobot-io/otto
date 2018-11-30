package main

import (
	"flag"
	"log"
	"net/http"
	"plugin"
	"time"

	"github.com/gorilla/mux"
	"github.com/mobilerobot-io/otto/api"
)

// Everything on the command line should be a plugin
func main() {
	flag.Parse()

	r := mux.NewRouter()

	// This will serve files under http://localhost:8000/static/<filename>
	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	srv := &http.Server{
		Handler: r,
		Addr:    ":8777",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	for _, arg := range flag.Args() {

		pl, err := plugin.Open(arg)
		check(err)

		svcsym, err := pl.Lookup("Service")
		check(err)

		svc := *svcsym.(*api.Service)
		path := svc.Path()

		s := r.PathPrefix(path).Subrouter()
		svc.Register(s)

		// pass the new subrouter to our register callback from our plugin
		//err = reg.(func(r *mux.Router) error)(s)
		check(err)
	}

	log.Println("  otto is starting on ", srv.Addr)
	srv.ListenAndServe()

	log.Println("Good bye")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
