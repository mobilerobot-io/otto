package main

import (
	"encoding/json"
	"net/url"

	log "github.com/sirupsen/logrus"
)

// Message is represented by a URL, but passed around micro services as
// inputs from sensors and outputs for control systems and display.
type Message struct {
	*url.URL
	message string
	err     error
}

// Turn this into an channel
func Incoming(m string) {
	url, err := url.ParseQuery(m)
	if err != nil {
		panic(err)
	}
	log.Infof("incoming message %s ~> %+v", m, url)
}

func toJSON(msg interface{}) string {
	js, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(js)
}
