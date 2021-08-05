package tcp

import (
	"encoding/gob"
	"net"
)

type DtpRemote struct {
	Address *net.TCPAddr
	Socket  *net.TCPConn
	Id      string
	encoder *gob.Encoder
	decoder *gob.Decoder
}
