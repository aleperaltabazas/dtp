package cli

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/channels"
	"github.com/aleperaltabazas/dtp/connection"
	"github.com/aleperaltabazas/dtp/protocol/codes"
)

func Ls(args []string) {
	if len(args) > 1 {
		fmt.Println("ls: too many arguments")
		return
	}

	r := connection.ConnectedRemote

	if r == nil {
		fmt.Println("You're not connected to anything!")
	} else {
		path := r.Pwd

		if len(args) == 1 {
			p := args[0]
			if string(p) == "/" {
				path = p
			} else {
				path = fmt.Sprintf("%s/%s", r.Pwd, p)
			}
		}

		err := r.Send(codes.ListDirectory, codes.NoSource, path)

		if err != nil {
			fmt.Printf("Error sending LS request: %e", err)
			return
		}

		m := <-channels.Ls

		switch m.Code {
		case codes.Ack:
			var files []string
			err = m.Deserialize(&files)

			if err != nil {
				fmt.Printf("Error deserializing body into files: %e", err)
				return
			}

			for _, f := range files {
				fmt.Println(f)
			}
		case codes.Error:
			var errorMessage string
			err = m.Deserialize(&errorMessage)

			if err != nil {
				fmt.Printf("Error deserializing body: %e", err)
				return
			}

			fmt.Println(errorMessage)
		}
	}
}
