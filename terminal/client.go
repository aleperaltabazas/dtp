package dtp

import (
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/aleperaltabazas/dtp/auth"
	"github.com/aleperaltabazas/dtp/protocol"
	"net"
)

func (r *Remote) Send(message interface{}) {
	err := r.encoder.Encode(message)
	if err != nil {
		fmt.Printf("Failed to deliver message to %s\n", r.Id)
		return
	}
}

func (r *Remote) Receive(message interface{}) error {
	err := r.decoder.Decode(message)
	if err != nil {
		fmt.Printf("Failed to receive message from %s", r.Id)
		return err
	}
	return nil
}

func (r *Remote) Close() error {
	r.Send("STOP")
	return r.Socket.Close()
}

func Connect(id, address string) (*Remote, error) {
	tcpAddr, closeErr := net.ResolveTCPAddr("tcp4", address)
	if closeErr != nil {
		return nil, closeErr
	}

	conn, closeErr := net.DialTCP("tcp4", nil, tcpAddr)
	if closeErr != nil {
		return nil, closeErr
	}

	encoder := gob.NewEncoder(conn)
	closeErr = encoder.Encode(authenticationRequest{Id: id, Passphrase: auth.Passphrase()})

	if closeErr != nil {
		closeErr := conn.Close()
		if closeErr != nil {
			return nil, closeErr
		}
		return nil, closeErr
	}

	dec := gob.NewDecoder(conn)
	serverId := new(authenticationResponse)
	closeErr = dec.Decode(&serverId)

	if closeErr != nil {
		closeErr := conn.Close()
		if closeErr != nil {
			return nil, closeErr
		}
		return nil, closeErr
	}

	switch serverId.Code {
	case authenticationError:
		closeErr := conn.Close()
		if closeErr != nil {
			return nil, closeErr
		}
		return nil, errors.New("server rejected authentication")
	case busy:
		closeErr = conn.Close()
		if closeErr != nil {
			return nil, closeErr
		}
		return nil, errors.New("server busy")
	}

	// TODO: crossed passphrase validation

	closeErr = encoder.Encode(protocol.Message{
		Code: protocol.Ack,
		Body: nil,
	})

	if closeErr != nil {
		closeErr := conn.Close()
		if closeErr != nil {
			return nil, closeErr
		}
		return nil, closeErr
	}

	return &Remote{
		Socket:  conn,
		encoder: encoder,
		decoder: dec,
		Id:      *serverId.Id,
	}, nil
}
