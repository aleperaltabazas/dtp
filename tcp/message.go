package tcp

import (
	"encoding/gob"
	"net"
)

func Send(c *net.TCPConn, message interface{}) error {
	encoder := gob.NewEncoder(c)
	return encoder.Encode(message)
}

func Receive(c *net.TCPConn, message interface{}) error {
	decoder := gob.NewDecoder(c)
	return decoder.Decode(message)
}
