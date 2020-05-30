package server

import (
	"github.com/recraft/recraft-server/connection"
	"github.com/recraft/recraft-server/console"
)

//NewServer creates server instance
func NewServer(address string, port int16, logType int8) error {
	server := &connection.Server{Network: &connection.ServerNetwork{Address: address, Port: port, Connected: false}}
	go server.Listen(&console.Console{Loglevel: console.LogLevel(logType)})

	for {
		status := <-server.Info
		if status.Action == "close" {
			server.Console.Info("Closed server")
			return nil
		}
	}
}
