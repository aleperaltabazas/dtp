package dtp

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/aleperaltabazas/dtp/auth"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
	"net"
)

func (r *Remote) Send(requestCode string, source string, body interface{}) error {
	var bs = make([]byte, 0)

	if body != nil {
		var buf *bytes.Buffer
		enc := gob.NewEncoder(buf)
		err := enc.Encode(body)

		if err != nil {
			return err
		}
	}

	request := protocol.Message{
		Code:   requestCode,
		Source: source,
		Body:   bs,
	}

	err := r.encoder.Encode(request)
	if err != nil {
		return err
	}

	return nil
}

func (r *Remote) Receive() (*protocol.Message, error) {
	var message protocol.Message

	err := r.decoder.Decode(&message)
	if err != nil {
		fmt.Printf("Failed to receive message from %s", r.Id)
		return nil, err
	}
	return &message ,nil
}

func (r *Remote) Close() error {
	source := codes.NoSource
	err := r.Send(codes.Fin, source, nil)

	if err != nil {
		return err
	}

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
		Code: codes.Ack,
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
