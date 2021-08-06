package main

import (
	"dtp/cli"
	"dtp/connection"
	"dtp/console"
	"fmt"
	"strings"
)

func handleCLI() {
	for {
		connected := connection.ConnectedRemote != nil

		prefix := "> "
		if connected {
			prefix = fmt.Sprintf("%s@%s> ", connection.ConnectedRemote.Id, connection.ConnectedRemote.Address.String())
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
