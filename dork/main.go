package main

import (
	"flag"
	"log"

	"github.com/digitalocean/godo"
)

var (
	config Configuration // Configuration and flags are handled in config.go

	client  *godo.Client
	account *godo.Account
)

func main() {
	var objs []interface{}
	if config.Account {
		acct := doAccount()
		objs = append(objs, acct)
	}

	if config.Actions {
		actions := doActions()
		objs = append(objs, actions)
	}

	if config.CDNs {
		cdns := doCDN()
		objs = append(objs, cdns)
	}

	if config.Projects {
		projs := doProjects()
		objs = append(objs, projs)
	}

	resp := getResponse(config.Output, config.Format)
	for _, o := range objs {
		resp.Respond(o)
	}
}
