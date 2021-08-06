package cli

import (
	"dtp/connection"
	"dtp/terminal"
	"fmt"
)

func Connect(ownId, host string) {
	fmt.Printf("Establishing handshake with %s\n", host)
	t, err := dtp.Connect(ownId, host)
	if err != nil {
		fmt.Printf("Failed to establish connection to %s: %s\n", host, err.Error())
	}

	if t != nil {
		connection.ConnectedRemote = t
		fmt.Printf("Connected to %s!\n", t.Id)
	}
}
