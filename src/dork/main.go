package main

import (
	"flag"
	"log"

	"github.com/digitalocean/godo"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/multi"
)

var (
	config Configuration // Configuration and flags are handled in config.go

	client  *godo.Client
	account *godo.Account
)

func main() {
	flag.Parse()

	client = doClient()
	g := multi.NewDirectedGraph()

	projs := doProjects()
	var lastNode *graph.Node
	for _, p := range projs {
		n := g.NewNode()
		g.AddNode(n)

		if lastNode != nil {
			g.NewLine(lastNode, n)
			lastNode = n
		}
		p = p
	}
	g.AddNode(g.NewNode())
	log.Fatalf("%+v\n", g)

	/*
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
	*/
}
