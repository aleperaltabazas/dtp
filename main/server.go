package main

import (
	"dtp/tcp"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func startServer(port string) *net.TCPListener {
	tcpAddr, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}
	l, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	rand.Seed(time.Now().Unix())

	fmt.Printf("Listening in port %s\n", port)
	go func() {
		for {
			c, err := l.AcceptTCP()
			host, err := tcp.Accept("falopa", tcpAddr, c, inputChan)
			if err != nil {
				fmt.Println(err)
				c.Close()
				break
			}
			if host == nil {
				c.Close()
				break
			}
			go handle(host)
		}
	}()

	return l
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
