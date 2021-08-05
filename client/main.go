package main

import (
	"dtp/console"
	"dtp/tcp"
	"fmt"
	"log"
	"os"
	"strings"
)

var id string

func main() {
	id = console.Prompt("Please, tell me your id")
	handleInput()
}

func handleInput() {
	connected := false
	commands := make(chan string)
	for {
		input, err := console.Reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		words := strings.Split(strings.TrimSpace(input), " ")

		switch strings.TrimSpace(words[0]) {
		case "connect":
			if !connected {
				connected = true
				host := tcp.Connect(id, strings.TrimSpace(words[1]))
				go handleConnection(host, commands)
			} else {
				fmt.Println("Already have an open connection!")
			}
		case "stop":
			{
				commands <- "stop"
				connected = false
			}
		case "ping":
			commands <- "ping"
		case "exit":
			commands <- "stop"
			os.Exit(0)
		}
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
