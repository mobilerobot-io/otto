package main

import (
	"encoding/json"
	"io/ioutil"

	homedir "github.com/mitchellh/go-homedir"
	namecheap "github.com/rustyeddy/go-namecheap"
)

type nccreds struct {
	Name, Token, User string
}

// get client will get the creds and create a client
func getClient() (ncli *namecheap.Client) {
	if ncli == nil {
		u, t, n := creds()
		ncli = namecheap.NewClient(u, t, n)
		if ncli == nil {
			panic("no client")
		}
	}
	return ncli
}

// return the Namecheap Token string
func creds() (u, t, v string) {

	fname, err := homedir.Dir()
	panicIfError(err)

	credfile := fname + "/" + ".config/namecheap/creds.json"
	b, err := ioutil.ReadFile(credfile)
	panicIfError(err)

	var creds nccreds
	err = json.Unmarshal(b, &creds)
	panicIfError(err)

	// convert bytes to string
	return creds.Name, creds.Token, creds.User
}
