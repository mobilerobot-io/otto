package otto

import (
	"log"
)

type Configuration struct {
	LogLevel  string
	LogOutput string
	LogFormat string

	// Addrport string
	// Basedir  string // base directory for saving data
	// Fetch    bool
	// Cache    bool // cache input from api calls? Default = yes
}

func fatalIfError(err error) {
	if err == nil {
		return
	}
	log.Fatalln(err)
}
