package main

import (
	"dtp/console"
	"dtp/semaphores"
	"dtp/tcp"
	"fmt"
	"strings"
)

func handleInput() {
	commands := make(chan string)
	for {
		connected := remote != nil

		prefix := ""
		if remote != nil {
			prefix = fmt.Sprintf("%s@%s", remote.Id, remote.Address.String())
		}

		semaphores.StdinLock.Lock()
		input := console.GetLine(prefix)
		words := strings.Split(input, " ")

		switch words[0] {
		case "":
			continue
		case "connect":
			switch len(words) {
			case 1:
				fmt.Println("Missing remote. Usage: connect host:port")
			case 2:
				if !connected {
					remote = tcp.Connect(id, strings.TrimSpace(words[1]))
					go handleConnection(remote, commands)
				} else {
					fmt.Println("Already have an open connection!")
				}
			default:
				fmt.Println("Too many arguments. Usage: connect host:port")
			}
		case "disconnect":
			if !connected {
				fmt.Println("You are not connected to any remote!")
			} else {
				commands <- "stop"
				connected = false
				commands = make(chan string)
			}
		case "ping":
			if !connected {
				fmt.Println("You are not connected to any remote!")
			} else {
				fmt.Println("Ping")
				commands <- "ping"
			}
		case "exit":
			if connected {
				fmt.Println("Disconnecting from remote...")
				commands <- "stop"
			}
			fmt.Println("Closing server...")
			listener.Close()
			break
		default:
			fmt.Printf("Unkown input '%s'\n", input)
		}
		semaphores.StdinLock.Unlock()
	}
}

func handleConnection(host *tcp.DtpRemote, commands chan string) {
	for {
		command := <-commands
		switch command {
		case "stop":
			{
				host.Close()
				break
			}
		case "ping":
			{
				host.Send("ping")
			}
		default:
			fmt.Printf("Unknown command %s\n", command)
		}
	}
}
