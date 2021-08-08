package cli

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/channels"
	"github.com/aleperaltabazas/dtp/connection"
	"github.com/aleperaltabazas/dtp/console"
	"github.com/aleperaltabazas/dtp/protocol/codes"
)

func Disconnect() {
	r := connection.ConnectedRemote
	if r == nil {
		console.NewLine()
		fmt.Println("Nothing to disconnect from")
	} else {
		err := r.Send(codes.Fin, codes.NoSource, nil)

		if err != nil {
			fmt.Printf("There was an error sending FIN to %s: %e\n", r.Address(), err)
		}
		_ = <- channels.Fin
		connection.ConnectedRemote = nil

		defer r.Socket.Close()
	}
}
