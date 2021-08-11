package main

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/console"
	"github.com/aleperaltabazas/dtp/global"
	"github.com/aleperaltabazas/dtp/handlers"
	"log"
	"os"
	"os/exec"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Missing: COMMAND\n\nUsage: dtp COMMAND")
		os.Exit(1)
	}

	switch args[1] {
	case "config":
		configCmd.Parse(args[2:])
		home, err := os.UserHomeDir()

		if err != nil {
			log.Fatalf("There was an error getting the user home: %v", err)
		}

		cmd := exec.Command("/bin/nano", fmt.Sprintf("%s/.config/dtp.yaml", home))
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()

		if err != nil {
			fmt.Printf("Failed to open editor: %v\n", err)
			os.Exit(1)
		}
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
		fmt.Printf("Invalid argument '%s'\n\nUsage: dtp COMMAND\n", args[1])
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
