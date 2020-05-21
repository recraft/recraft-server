package types

import "net"

type ClientInfo struct {
	Username   string
	Status     string
	Connection *net.Conn
}
