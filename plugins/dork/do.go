package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/digitalocean/godo"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

var (
	config Configuration // Configuration and flags are handled in config.go

	client  *godo.Client
	account *godo.Account
)

// Client
// ====================================================================

// doClient will find the access token, prepare for authentication
// ready to call the server with some important stuff
func doClient() (cli *godo.Client) {
	creds := doCreds()
	tokenSource := &TokenSource{
		AccessToken: creds.Token,
	}
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	cli = godo.NewClient(oauthClient)
	return cli
}

// Credentials
// ====================================================================

type docreds struct {
	Name, Token string
}

// doCreds reads the digital credits from a json file.
func doCreds() (creds docreds) {
	b, err := ioutil.ReadFile(config.Creds)
	panicIfError(err)

	err = json.Unmarshal(b, &creds)
	panicIfError(err)

	return creds
}

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

// Account
// ====================================================================

// doAccount fetch and return the Digital Ocean account
func doAccount() (a *godo.Account) {
	ctx := context.TODO()
	a, _, err := client.Account.Get(ctx)
	panicIfError(err)

	// Check if something may be wrong with the account
	if a.Status != "active" {
		log.Errorf("URGENT! expected account status (active) got (%s)", a.Status)
	}
	return a
}

// check your account for sanity
func displayAccount(w io.Writer) {
	account := doAccount()
	if account.Status != "active" {
		log.Errorf("DO Account issue expected status (active) got (%s) ", account.Status)
		return
	}
	fmt.Fprintf(w, account.String())
}

// Actions
// ====================================================================
func listOpts() (ctx context.Context, opt *godo.ListOptions) {
	opt = &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	ctx = context.TODO()
	return ctx, opt
}

func doActions() []godo.Action {
	ctx, opt := listOpts()
	actions, _, err := client.Actions.List(ctx, opt)
	panicIfError(err)
	return actions
}

func displayActions(w io.Writer) (ok bool) {
	acts := doActions()

	for _, act := range acts {
		fmt.Fprintf(w, "%d ~ %s %s == %s %s\n", act.ID, act.Type, act.Status, act.ResourceType, act.CompletedAt)
	}
	fmt.Fprintf(w, "\n")
	return true
}

// CDN
// ====================================================================
func doCDN() []godo.CDN {
	ctx, opts := listOpts()

	cdns, _, err := client.CDNs.List(ctx, opts)
	panicIfError(err)
	return cdns
}

func displayCDN(w io.Writer) (ok bool) {
	cdns := doCDN()
	panicIfNil(cdns)

	fmt.Fprintf(w, "CDNs ...")
	for _, cdn := range cdns {
		fmt.Fprintf(w, "\t%s => %s\n", cdn.Origin, cdn.Endpoint)
	}
	return true
}

func Do() {
	var objs []interface{}
	acct := doAccount()

	objs = append(objs, acct)

	actions := doActions()
	objs = append(objs, actions)

	cdns := doCDN()
	objs = append(objs, cdns)

	projs := doProjects()
	objs = append(objs, projs)

	resp := getResponse(config.Output, config.Format)
	for _, o := range objs {
		resp.Respond(o)
	}
}
