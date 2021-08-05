package main

import (
	"dtp/console"
	"dtp/tcp"
	"fmt"
	"net"
	"os"
)

var remote *tcp.DtpRemote = nil
var id string
var listener *net.TCPListener
var inputChan = make(chan string)

func main() {
	id = console.Prompt("Please, tell me your id")
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Please provide a port number!")
		return
	}

	port := ":" + arguments[1]

	listener = startServer(port)
	handleCLI()

	println("Bye!")
}
