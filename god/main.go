package main

/*

The GoDaemon (an unfortunant acronym, the name will need to be changed, eventually)

*/
import (
	"flag"
	"io/ioutil"

	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	dir  string
	addr string
)

func init() {
	flag.StringVar(&dir, "dir", "/srv/www/", "Serve up a page or site from directory")
	flag.StringVar(&addr, "addr", ":4343", "Port and local address to bind to")
}

func main() {
	flag.Parse()
	r := mux.NewRouter()

	for _, site := range getSites() {
		path := dir + site
		s := "/" + site

		log.Printf("dir: %s ~ site: %s ~ path %s ~ s %s", dir, site, path, s)
		r.PathPrefix(s).Handler(http.StripPrefix(s, http.FileServer(http.Dir(path))))
	}

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}

}

func getSites() (sites []string) {
	if dir == "" {
		sites = flag.Args()
		return sites
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		sites = append(sites, file.Name())
	}
	return sites
}
