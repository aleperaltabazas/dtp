package dtp

import (
	"encoding/gob"
	"fmt"
	"github.com/aleperaltabazas/dtp/auth"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
	"github.com/aleperaltabazas/dtp/tcp"
	"net"
)

func Accept(ownId string, conn *net.TCPConn) (*Remote, error) {
	decoder := gob.NewDecoder(conn)
	clientId := new(authenticationRequest)
	err := decoder.Decode(clientId)

	if err != nil {
		return nil, err
	}

	if !auth.Authenticate(clientId.Passphrase) {
		fmt.Println("Authentication failed. Closing connection")
		sendErr := tcp.Send(conn, authenticationResponse{
			Code: authenticationError,
			Id:   nil,
		})
		if sendErr != nil {
			panic(sendErr)
		}
		closeErr := conn.Close()
		if closeErr != nil {
			panic(closeErr)
		}
		return nil, nil
	}

	fmt.Printf("\nNew client %s acknowledged correctly, presenting myself...\n", clientId.Id)

	// TODO: crossed passphrase validation
	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(authenticationResponse{Code: authenticationOk, Id: &ownId})

	if err != nil {
		return nil, err
	}

	ack := new(protocol.Message)
	err = decoder.Decode(ack)

	if err != nil {
		closeError := conn.Close()
		if closeError != nil {
			return nil, closeError
		}
		return nil, err
	}

	if ack.Code != codes.Ack {
		fmt.Printf("%s rejected the connection\n", *clientId)
		closeError := conn.Close()
		if closeError != nil {
			return nil, closeError
		}
		return nil, nil
	}

	return &Remote{
		Socket:  conn,
		Id:      clientId.Id,
		encoder: encoder,
		decoder: decoder,
	}, nil
}

func Reject(conn *net.TCPConn) {
	encoder := gob.NewEncoder(conn)
	err := encoder.Encode(authenticationResponse{
		Code: busy,
		Id:   nil,
	})

	if err != nil {
		fmt.Printf("Failed to report rejection: %e\n", err)
	}

	err = conn.Close()

	if err != nil {
		fmt.Printf("Failed to close socket: %e\n", err)
	}
}
