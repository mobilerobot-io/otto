package srv

// Serve up our services with srv.
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
	flag.StringVar(&www, "dir", "/srv/www/", "Serve up a page or site from directory")
	flag.StringVar(&addr, "addr", ":80", "Port and local address to bind defualt :80")
}

// Start a static file server using the base directory as root file system
func StaticServer(addrport string, basedir string) {
	r := mux.NewRouter()

	log.Infof("Listing on %s ~> serving static sites from ", basedir)

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
