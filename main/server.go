package main

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/connection"
	dtp "github.com/aleperaltabazas/dtp/terminal"
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
			if connection.ConnectedRemote == nil {
				r, err := dtp.Accept(id, c)
				if err != nil {
					fmt.Println(err)
					closeError := c.Close()
					if closeError != nil {
						fmt.Printf("Failed to close the socket: %s\n", err.Error())
					}
					continue
				}
				connection.ConnectedRemote = r
				fmt.Printf("Connected to %s!\n", r.Id)
				fmt.Printf("%s@%s> ", r.Id, r.Address())
			} else {
				dtp.Reject(c)
			}
		}
	}()

	return l
}
