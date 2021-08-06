package main

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/cli"
	"github.com/aleperaltabazas/dtp/connection"
	"github.com/aleperaltabazas/dtp/console"
	"strings"
)

func handleCLI() {
	for {
		connected := connection.ConnectedRemote != nil

		prefix := "> "
		if connected {
			prefix = fmt.Sprintf("%s@%s> ", connection.ConnectedRemote.Id, connection.ConnectedRemote.Address())
		}
		input := console.GetLine(prefix)

		words := strings.Split(input, " ")

		switch words[0] {
		case "":
			continue
		case "connect":
			switch len(words) {
			case 1:
				fmt.Println("Missing connection. Usage: connect host:port")
			case 2:
				if !connected {
					cli.Connect(id, words[1])
				} else {
					fmt.Println("Already have an open connection!")
				}
			default:
				fmt.Println("Too many arguments. Usage: connect host:port")
			}
		case "status":
			cli.Status()
		case "disconnect":
			cli.Disconnect()
		case "ping":
			cli.Ping()
		case "exit":
			cli.Exit()
		default:
			fmt.Printf("Unkown input '%s'\n", input)
		}
	}
}
