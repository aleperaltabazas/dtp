package dtp

import (
	"encoding/gob"
	"net"
)

type Remote struct {
	Socket  *net.TCPConn
	Id      string
	encoder *gob.Encoder
	decoder *gob.Decoder
}

func (r *Remote) Address() string {
	return r.Socket.LocalAddr().String()
}

type authenticationRequest struct {
	Passphrase []byte
	Id         string
}

const (
	busy                = -100
	authenticationError = -1
	authenticationOk    = 1
)

type authenticationResponse struct {
	Code int
	Id   *string
}
