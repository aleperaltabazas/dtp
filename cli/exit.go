package cli

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/connection"
	"github.com/aleperaltabazas/dtp/tcp"
)

func Exit() {
	if connection.AcceptPending() {
		Reject()
	}

	if connection.IsConnected() {
		Disconnect()
	}

	err := tcp.Listener.Close()

	if err != nil {
		fmt.Printf("There was an error closing the TCP server: %s\n", err.Error())
	}
}
