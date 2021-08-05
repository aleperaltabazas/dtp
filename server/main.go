package main

import (
	"dtp/tcp"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	port := ":" + arguments[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}
	l, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	fmt.Printf("Listening in port %s\n", port)
	for {
		c, err := l.AcceptTCP()
		host, err := tcp.Accept("falopa", tcpAddr, c)
		if err != nil {
			fmt.Println(err)
			c.Close()
			return
		}
		if host == nil {
			c.Close()
			return
		}
		go handle(host)
	}
}

func handle(host *tcp.DtpRemote) {
	fmt.Printf("Serving %s\n", host.Address)
	for {
		message := new(string)
		err := host.Receive(message)

		if err != nil {
			break
		}

		fmt.Printf("Received message from %s: %s", host.Id, *message)

		if *message == "STOP" {
			break
		} else if *message == "ping" {
			host.Send("pong")
		}
	}

	err := host.Close()
	if err != nil {
		fmt.Printf("Failed to close connection with %s\n", host.Id)
	}
}
