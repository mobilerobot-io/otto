package main

/* Serve up our services with srv.
   - static file server
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
	flag.StringVar(&addr, "addr", "", "Port and local address to bind defualt :80")
}

func main() {
	flag.Parse()
	r := mux.NewRouter()

	log.Infoln("Serving static sites ")
	for _, site := range getSites() {
		path := dir + site
		s := "/" + site

		log.Infof("\t%-10s %s", site, path)
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

// getSites will return a list of site names depending on the how the
// command was invoked: -dir <dir> will cause the program to scan the
// given directory and serve all subdirectories as a static website.
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
