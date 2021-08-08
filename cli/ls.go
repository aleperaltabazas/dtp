package cli

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/channels"
	"github.com/aleperaltabazas/dtp/connection"
	"github.com/aleperaltabazas/dtp/protocol/codes"
)

func Ls() {
	r := connection.ConnectedRemote

	if r == nil {
		fmt.Println("You're not connected to anything!")
	} else {
		err := r.Send(codes.ListDirectory, codes.NoSource, nil)

		if err != nil {
			fmt.Printf("Error sending LS request: %e", err)
			return
		}

		m := <-channels.Ls

		var files []string
		err = m.Deserialize(&files)

		if err != nil {
			fmt.Printf("Error deserializing body into files: %e", err)
			return
		}

		for _, f := range files {
			fmt.Println(f)
		}
	}
}
