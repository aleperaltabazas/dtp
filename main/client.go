package main

import (
	"dtp/console"
	"dtp/semaphores"
	"dtp/tcp"
	"fmt"
	"strings"
)

func handleCLI() {
	for {
		connected := remote != nil

		prefix := "> "
		if connected {
			prefix = fmt.Sprintf("%s@%s> ", remote.Id, remote.Address.String())
		}
		input := console.GetLine(prefix)

		semaphores.RedirectionLock.Lock()
		if semaphores.RedirectInput {
			inputChan <- input
			semaphores.RedirectInput = false
			semaphores.RedirectionLock.Unlock()
			continue
		}

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
				remote.Close()
				remote = nil
			}
		case "ping":
			if !connected {
				fmt.Println("You are not connected to any remote!")
			} else {
				remote.Send("ping")
				fmt.Println("Ping")
			}
		case "exit":
			if connected {
				fmt.Println("Disconnecting from remote...")
				remote.Close()
				remote = nil
			}
			fmt.Println("Closing server...")
			listener.Close()
			break
		default:
			fmt.Printf("Unkown input '%s'\n", input)
		}
		semaphores.RedirectionLock.Unlock()
	}
}
