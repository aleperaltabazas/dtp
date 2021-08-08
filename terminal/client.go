package dtp

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"github.com/aleperaltabazas/dtp/auth"
	"github.com/aleperaltabazas/dtp/filesystem"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
	"net"
)

func (r *Remote) Send(requestCode string, source string, body interface{}) error {
	var bs = make([]byte, 0)

	if body != nil {
		j, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bs = j
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
		return nil, err
	}

	return &message, nil
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

	cwd := filesystem.GetCurrentDirectory()

	j, err := json.Marshal(cwd)
	if err != nil {
		return nil, err
	}

	closeErr = encoder.Encode(protocol.Message{
		Code:   codes.Ack,
		Source: codes.NoSource,
		Body:   j,
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
		Pwd:     *serverId.Pwd,
	}, nil
}
