package main

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/console"
	"github.com/aleperaltabazas/dtp/global"
	"github.com/aleperaltabazas/dtp/handlers"
	"os"
)

func main() {
	handlers.Init()
	askId()
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Please provide a port number!")
		return
	}

	port := ":" + arguments[1]

	global.Listener = startServer(port)
	console.NewLine()
	handleCLI()

	println("Bye!")
}

func askId() {
	for {
		id := console.Prompt("Please, tell me your id: ")

		if id == console.EOF {
			os.Exit(0)
		}
		if len(id) > 0 {
			global.Id = id
			return
		}
	}
}
