package main

import (
	"dtp/connection"
	"dtp/tcp"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func startServer(port string) *net.TCPListener {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
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

			if err != nil {
				fmt.Println(err)
				closeError := c.Close()
				if closeError != nil {
					fmt.Printf("Failed to close the socket: %s\n", err.Error())
				}
				continue
			}

			if connection.AwaitingConnection == nil {
				connection.AwaitingConnection = c
			} else {
				err := tcp.Send(c, "busy")
				if err != nil {
					fmt.Printf("Failed to reject the client: %s\n", err)
				}
				closeError := c.Close()
				if closeError != nil {
					fmt.Printf("Failed to close the socket: %s\n", err.Error())
				}
			}
		}
	}()

	return l
}
