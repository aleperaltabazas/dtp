package cli

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/connection"
	"github.com/aleperaltabazas/dtp/tcp"
)

func Reject() {
	if !connection.AcceptPending() {
		fmt.Println("No pending connection")
	} else {
		err := tcp.Send(connection.AwaitingConnection, "reject")

		if err != nil {
			fmt.Printf("There was an error rejecting the connection: %s\n", err.Error())
		}
		closeError := connection.AwaitingConnection.Close()
		if closeError != nil {
			fmt.Printf("There was an error closing the socket: %s\n", err.Error())
		}
		connection.AwaitingConnection = nil
	}
}
