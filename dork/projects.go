package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/digitalocean/godo"
)

// Projects
// ====================================================================
func doProjects() []godo.Project {
	ctx, opts := listOpts()

	projs, _, err := client.Projects.List(ctx, opts)
	panicIfError(err)
	return projs
}

func displayProjects(w io.Writer) (ok bool) {
	projs := doProjects()
	panicIfNil(projs)

	fmt.Fprintf(w, "projects (%d)...\n", len(projs))
	for _, proj := range projs {

		fmt.Fprintf(w, "\t%s: %s\n", proj.Name, proj.Description)
		ok := displayProjectResources(w, proj.ID)
		okCheck(ok, "display project resources")
	}
	return true
}

func doResources(id string) (res []godo.ProjectResource) {
	var err error

	ctx, opts := listOpts()
	res, _, err = client.Projects.ListResources(ctx, id, opts)
	panicIfError(err)
	return res
}

func displayProjectResources(w io.Writer, id string) (ok bool) {
	res := doResources(id)
	fmt.Fprintf(w, "Project id: %s\n", id)
	for _, v := range res {

		foo := strings.Split(v.URN, ":")
		fmt.Fprintf(w, "%20s: %-30s status %26s\n", foo[1], foo[2], v.Status)
	}
	return true
}
