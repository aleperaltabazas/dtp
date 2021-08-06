package main

import (
	"dtp/console"
	"dtp/tcp"
	"fmt"
	"os"
)

var id string

func main() {
	id = console.Prompt("Please, tell me your id: ")
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Please provide a port number!")
		return
	}

	port := ":" + arguments[1]

	tcp.Listener = startServer(port)
	handleCLI()

	println("Bye!")
}
