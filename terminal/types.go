package dtp

import (
	"encoding/gob"
	"net"
)

type Remote struct {
	Address *net.TCPAddr
	Socket  *net.TCPConn
	Id      string
	encoder *gob.Encoder
	decoder *gob.Decoder
}
