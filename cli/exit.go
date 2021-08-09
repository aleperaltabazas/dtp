package cli

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/connection"
	"github.com/aleperaltabazas/dtp/global"
	"os"
)

func Exit() {
	if connection.AcceptPending() {
		Reject()
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
