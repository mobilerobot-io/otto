package otto

import (
	"net"
)

type Node interface {
	NodeID
	Interfaces() (ifaces []*net.Interface)
}

// NodeID is the DNS name, IP Address
type NodeID interface {
	Name() string    // DNS or host name
	Domain() string  // Domain for this node
	IPAddr() string  // IPAddress
	MACAddr() string // MACAddr
}

type Domain interface {
	Name() string // name of this domain
	Nameservers() (nslist []*net.NS, err error)
	Provider() (registrar string)
}
