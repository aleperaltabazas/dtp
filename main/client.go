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

		connection.ShowConsolePrompt()
		input := console.GetLine()

		words := strings.Split(input, " ")

		if len(words) == 0 {
			continue
		}

		command := words[0]
		args := words[1:]

		switch command {
		case "":
			continue
		case ":connect":
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
		case ":status":
			cli.Status()
		case ":disconnect":
			cli.Disconnect()
		case ":ping":
			cli.Ping()
		case ":exit":
			cli.Exit()
		case "ls":
			cli.Ls(args)
		case "pwd":
			cli.Pwd()
		case "cd":
			cli.Cd(args)
		default:
			fmt.Printf("Unkown input '%s'\n", input)
		}
	}
}
