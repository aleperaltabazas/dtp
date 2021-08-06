package cli

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/connection"
)

func Disconnect() {
	if connection.ConnectedRemote == nil {
		fmt.Println("Nothing to disconnect from")
	} else {
		err := connection.ConnectedRemote.Close()
		if err != nil {
			fmt.Printf("There was an error disconnecting from %s: %s\n", connection.ConnectedRemote.Id, err.Error())
		}
		connection.ConnectedRemote = nil
	}
}
