package cli

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/connection"
	"github.com/aleperaltabazas/dtp/global"
	"github.com/aleperaltabazas/dtp/terminal"
)

func Connect(host string) {
	fmt.Printf("Establishing handshake with %s\n", host)
	t, err := dtp.Connect(global.Id, host)
	if err != nil {
		fmt.Printf("Failed to establish connection to %s: %s\n", host, err.Error())
	}

	if t != nil {
		connection.ConnectedRemote = t
		go connection.Receive()
		fmt.Printf("Connected to %s!\n", t.Id)
	}
}
