package cli

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/connection"
)

func Pwd() {
	r := connection.ConnectedRemote

	if r != nil {
		fmt.Println(r.Pwd)
	} else {
		fmt.Println("You're not connected to anything!")
	}
}
