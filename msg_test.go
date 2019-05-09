package main

import (
	"fmt"
	"net/url"
	"testing"
)

var (
	messages_out = []string{
		"msg://joy&x=200,y=200",
		"http://rpi:3/but/1/0",
	}
)

func TestMessage(t *testing.T) {
	for _, str := range messages_out {

		u, err := url.Parse(str)
		check(err)

		msg := &Message{
			URL: u,
		}
		/*
			msg.Scheme = "msg"
			msg.Host = "rpi"
			msg.Path = "joy/1"
			msg.RawQuery = "right=100,left=100"
			msg.Fragment = "anchor"
		*/
		msgstr := msg.String()
		fmt.Println(msgstr)
		msg.Opaque = "opaque"
		fmt.Println(msgstr)
	}
}
