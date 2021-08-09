package main

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/console"
	"github.com/aleperaltabazas/dtp/handlers"
	"github.com/aleperaltabazas/dtp/tcp"
	"os"
)

var id string

func main() {
	handlers.Init()
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
