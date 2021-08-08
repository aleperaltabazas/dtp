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
		Reject()
	}

	if connection.IsConnected() {
		Disconnect()
	}

	global.StopLock.Lock()
	global.Stop = true
	err := tcp.Listener.Close()
	global.StopLock.Unlock()

	if err != nil {
		fmt.Printf("There was an error closing the TCP server: %s\n", err.Error())
	}

	fmt.Println("Bye!")
	os.Exit(0)
}
