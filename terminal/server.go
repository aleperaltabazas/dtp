package dtp

import (
	"dtp/console"
	"encoding/gob"
	"fmt"
	"net"
)

func Accept(ownId string, tcpAddr *net.TCPAddr, conn *net.TCPConn) (*Remote, error) {
	address := conn.RemoteAddr().String()
	decoder := gob.NewDecoder(conn)
	clientId := new(string)
	err := decoder.Decode(clientId)

	if err != nil {
		return nil, err
	}

	if !console.PromptConfirmation(fmt.Sprintf("%s identifies themselves as %s, are you ok with this? (y/n): ", address, *clientId)) {
		closeError := conn.Close()
		if closeError != nil {
			return nil, closeError
		}
		return nil, nil
	}

	fmt.Printf("New client %s acknowledged correctly, presenting myself...\n", *clientId)

	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(ownId)

	if err != nil {
		return nil, err
	}

	ack := new(string)
	err = decoder.Decode(ack)

	if err != nil {
		closeError := conn.Close()
		if closeError != nil {
			return nil, closeError
		}
		return nil, err
	}

	if *ack != "ack" {
		fmt.Printf("%s rejected the connection\n", *clientId)
		closeError := conn.Close()
		if closeError != nil {
			return nil, closeError
		}
		return nil, nil
	}

	return &Remote{
		Address: tcpAddr,
		Socket:  conn,
		Id:      *clientId,
		encoder: encoder,
		decoder: decoder,
	}, nil
}
