package main

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/console"
	"github.com/aleperaltabazas/dtp/global"
	"github.com/aleperaltabazas/dtp/handlers"
	"os"
)

func main() {
	args := os.Args

	switch args[1] {
	case "config":
		configCmd.Parse(args[2:])
	case "start":
		handlers.Init()
		askId()
		startCmd.Parse(args[2:])
		port := fmt.Sprintf(":%v", *portOption)

		global.Listener = startServer(port)
		console.NewLine()
		handleCLI()

		println("Bye!")
	default:
		fmt.Println("Missing: COMMAND\n\nUsage: dtp COMMAND")
		os.Exit(1)
	}

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
