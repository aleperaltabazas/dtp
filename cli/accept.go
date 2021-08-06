package cli

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/connection"
	"github.com/aleperaltabazas/dtp/terminal"
	"net"
)

func Accept(ownId string, address *net.TCPAddr, conn *net.TCPConn) {
	t, err := dtp.Accept(ownId, address, conn)

	if err != nil {
		fmt.Printf("There was an error accepting the connection from %s: %s\n", address, err.Error())
	} else {
		connection.ConnectedRemote = t
		connection.AwaitingConnection = nil
		fmt.Printf("Connected to %s!\n", t.Id)
	}
}
