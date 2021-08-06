package cli

import (
	"dtp/connection"
	"fmt"
)

func Status() {
	if connection.IsConnected() {
		fmt.Printf("Connected to %s\n at %s", connection.ConnectedRemote.Id, connection.ConnectedRemote.Address.String())
	} else if connection.AcceptPending() {
		fmt.Printf("Connection request pending approval from %s\n", connection.AwaitingConnection.RemoteAddr().String())
	} else {
		fmt.Printf("Not connected to any DTP.\nNo connection requests open\n")
	}
}
