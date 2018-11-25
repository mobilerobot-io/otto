package main

import (
	"github.com/gonum/graph"
)

// Network collects a bunch of stuff
type Network struct {
	graph.Undirect
}

func NewNetwork() (n Network) {
	n = Network{}
	return n
}
