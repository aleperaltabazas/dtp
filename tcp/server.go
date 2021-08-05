package tcp

import (
	"dtp/console"
	"dtp/semaphores"
	"encoding/gob"
	"fmt"
	"net"
)

func Accept(ownId string, tcpAddr *net.TCPAddr, conn *net.TCPConn, inputChan chan string) (*DtpRemote, error) {
	address := conn.RemoteAddr().String()
	decoder := gob.NewDecoder(conn)
	clientId := new(string)
	err := decoder.Decode(clientId)

	if err != nil {
		fmt.Printf("Failed to decode client id %s\n", err)
		return nil, err
	}

	fmt.Printf("\n%s identifies themselves as %s, are you ok with this? (y/n): ", address, *clientId)
	semaphores.RedirectionLock.Lock()
	semaphores.RedirectInput = true
	semaphores.RedirectionLock.Unlock()
	res := <- inputChan
	if !console.Confirm(res) {
		fmt.Printf("Closing connection to %s\n", address)
		conn.Close()
		return nil, nil
	}

	fmt.Printf("New client %s acknowledged correctly, presenting myself...\n", *clientId)

	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(ownId)

	if err != nil {
		fmt.Printf("Failed to decode client id %s\n", err)
		return nil, err
	}

	ack := new(string)

	err = decoder.Decode(ack)

	if err != nil {
		fmt.Printf("Failed to receive ACK from %s\n", address)
		conn.Close()
		return nil, err
	}

	if *ack != "ack" {
		fmt.Printf("%s rejected the connection\n", *ack)
		conn.Close()
		return nil, nil
	}

	return &DtpRemote{
		Address: tcpAddr,
		Socket:  conn,
		Id:      *clientId,
		encoder: encoder,
		decoder: decoder,
	}, nil
}
