package cli

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/connection"
	"github.com/aleperaltabazas/dtp/global"
	"github.com/aleperaltabazas/dtp/tcp"
	"os"
)

func Exit() {
	if connection.AcceptPending() {
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

	if connection.IsConnected() {
		Disconnect()
	}

	global.StopLock.Lock()
	global.Stop = true

	if global.Listener != nil {
		err := global.Listener.Close()
		if err != nil {
			fmt.Printf("There was an error closing the TCP server: %s\n", err.Error())
		}
	}
	global.StopLock.Unlock()

	fmt.Println("Bye!")
	os.Exit(0)

}
