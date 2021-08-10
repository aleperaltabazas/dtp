package dtp

import (
	"encoding/json"
	"errors"
	"github.com/aleperaltabazas/dtp/auth"
	"github.com/aleperaltabazas/dtp/filesystem"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
	"github.com/aleperaltabazas/dtp/tcp"
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

	err := tcp.Send(r.Socket, request)
	if err != nil {
		return err
	}

	return nil
}

func (r *Remote) Receive() (*protocol.Message, error) {
	var m protocol.Message

	err := tcp.Receive(r.Socket, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
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

	closeErr = tcp.Send(conn, authenticationRequest{Id: id, Passphrase: auth.Passphrase()})

	if closeErr != nil {
		closeErr := conn.Close()
		if closeErr != nil {
			return nil, closeErr
		}
		return nil, closeErr
	}

	var serverId authenticationResponse
	closeErr = tcp.Receive(conn, &serverId)

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

	closeErr = tcp.Send(conn, protocol.Message{
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
		Socket: conn,
		Id:     *serverId.Id,
		Pwd:    *serverId.Pwd,
	}, nil
}
