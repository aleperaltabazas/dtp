package cli

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/channels"
	"github.com/aleperaltabazas/dtp/connection"
	"github.com/aleperaltabazas/dtp/protocol/codes"
	"log"
	"os"
)

func Cd(args []string) {
	r := connection.ConnectedRemote

	if r != nil {
		switch len(args) {
		case 0:
		case 1:
			path := args[0]
			first := string(path[0])

			if first != "/" {
				path = fmt.Sprintf("%s/%s", r.Pwd, path)
			}
			err := r.Send(codes.ChangeDirectory, codes.NoSource, path)

			if err != nil {
				fmt.Printf("There was an error sending CD: %e\n", err)
				return
			}

			res := <-channels.Cd

			switch res.Code {
			case codes.Ack:
				err = res.Deserialize(&path)

				if err != nil {
					fmt.Printf("Failed to deserialize absolute path: %e\n", err)
					return
				}

				connection.ConnectedRemote.Pwd = path
			case codes.Nak:
				fmt.Printf("No such directory: %s\n", path)
			default:
				fmt.Printf("unexpected code response: %s\n", res.Code)
			}
		default:
			fmt.Println("Too many arguments. Usage: cd dir")
		}
	} else {
		fmt.Println("You're not connected to anything!")
	}
}

func CdLocal(args []string) {
	switch len(args) {
	case 0:
		home, _ := os.UserHomeDir()
		os.Chdir(home)
	case 1:
		err := os.Chdir(args[0])

		if err != nil {
			log.Fatalf(err.Error())
		}
	default:
		fmt.Println("cd: too many arguments")
	}
}
