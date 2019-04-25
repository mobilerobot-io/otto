package main

import (
	"flag"
	"log"
	"net/http"

	"bitbucket.org/classroomsystems/wsgi"
)

func main() {
	addr := flag.String("a", ":7700", "address on which to serve HTTP")
	flag.Parse()
	var module, app string
	if flag.NArg() == 1 {
		module = flag.Arg(0)
		app = module
	} else if flag.NArg() == 2 {
		module = flag.Arg(0)
		app = flag.Arg(1)
	} else {
		log.Fatalln("No WSGI module specified")
	}
	err := wsgi.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	defer wsgi.Finalize()
	handler, err := wsgi.NewHandler(module, app, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer handler.Close()
	http.Handle("/", handler)
	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
