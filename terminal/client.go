package dtp

import (
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/aleperaltabazas/dtp/auth"
	"net"
)

func (h *Remote) Send(message interface{}) {
	err := h.encoder.Encode(message)
	if err != nil {
		fmt.Printf("Failed to deliver message to %s\n", h.Id)
		return
	}
}

func (h *Remote) Receive(message interface{}) error {
	err := h.decoder.Decode(message)
	if err != nil {
		fmt.Printf("Failed to receive message from %s", h.Id)
		return err
	}
	return nil
}

func (h *Remote) Close() error {
	h.Send("STOP")
	return h.Socket.Close()
}

func Connect(id, address string) (*Remote, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		return nil, err
	}

	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(authenticationRequest{Id: id, Passphrase: auth.Passphrase()})

	if err != nil {
		closeErr := conn.Close()
		if closeErr != nil {
			return nil, closeErr
		}
		return nil, err
	}

	dec := gob.NewDecoder(conn)
	serverId := new(authenticationResponse)
	err = dec.Decode(&serverId)

	if err != nil {
		closeErr := conn.Close()
		if closeErr != nil {
			return nil, closeErr
		}
		return nil, err
	}

	if serverId.Code == -1 {
		closeErr := conn.Close()
		if closeErr != nil {
			return nil, closeErr
		}
		return nil, errors.New("server rejected authentication")
	}

	// TODO: crossed passphrase validation

	err = encoder.Encode("ack")

	if err != nil {
		closeErr := conn.Close()
		if closeErr != nil {
			return nil, closeErr
		}
		return nil, err
	}

	return &Remote{
		Socket:  conn,
		encoder: encoder,
		decoder: dec,
		Id:      *serverId.Id,
	}, nil
}
