package dtp

import (
	"encoding/gob"
	"fmt"
	"github.com/aleperaltabazas/dtp/console"
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
	err = encoder.Encode(id)

	if err != nil {
		closeErr := conn.Close()
		if closeErr != nil {
			return nil, closeErr
		}
		return nil, err
	}

	dec := gob.NewDecoder(conn)
	serverId := new(string)
	err = dec.Decode(&serverId)

	if err != nil {
		closeErr := conn.Close()
		if closeErr != nil {
			return nil, closeErr
		}
		return nil, err
	}

	if !console.PromptConfirmation(fmt.Sprintf("%s identifies themselves as %s, are you ok with this?", address, *serverId)) {
		fmt.Printf("Closing connection to %s\n", address)
		encodeError := encoder.Encode("nak")

		if encodeError != nil {
			return nil, encodeError
		}

		closeErr := conn.Close()
		if closeErr != nil {
			return nil, closeErr
		}
		return nil, nil
	}

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
		Address: tcpAddr,
		encoder: encoder,
		decoder: dec,
		Id:      *serverId,
	}, nil
}
