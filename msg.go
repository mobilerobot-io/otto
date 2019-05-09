package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Message is represented by a URL, but passed around micro services as
// inputs from sensors and outputs for control systems and display.
type Message struct {
	*url.URL
	message string
	err     error
}

func Incoming(m string) {
	url, err := url.ParseQuery(`l=200&r=200`)
	if err != nil {
		panic(err)
	}
	fmt.Println(toJSON(url))
}

func toJSON(msg interface{}) string {
	js, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(js)
}
