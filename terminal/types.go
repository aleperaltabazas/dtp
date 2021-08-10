package dtp

import (
	"fmt"
	"net"
)

type Remote struct {
	Socket    *net.TCPConn
	Id        string
	Pwd       string
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
	Pwd  *string
}

func (r * Remote) FullPath(path string) string {
	return fmt.Sprintf("%s/%s", r.Pwd, path)
}
