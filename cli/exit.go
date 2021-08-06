package cli

import (
	"dtp/connection"
	"dtp/tcp"
	"fmt"
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
