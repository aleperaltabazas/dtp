package tcp

import (
	"dtp/console"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
)

func (h *DtpRemote) Send(message interface{}) {
	err := h.encoder.Encode(message)
	if err != nil {
		fmt.Printf("Failed to deliver message to %s\n", h.Id)
		return
	}
}

func (h *DtpRemote) Receive(message interface{}) error {
	err := h.decoder.Decode(message)
	if err != nil {
		fmt.Printf("Failed to receive message from %s", h.Id)
		return err
	}
	return nil
}

func (h *DtpRemote) Close() error {
	h.Send("STOP")
	return h.Socket.Close()
}

func Connect(id, address string) *DtpRemote {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(id)

	if err != nil {
		println("Write to server failed:", err.Error())
		conn.Close()
		os.Exit(1)
	}

	dec := gob.NewDecoder(conn)
	serverId := new(string)
	err = dec.Decode(&serverId)

	if err != nil {
		fmt.Printf("Failed to receive id from %s\n", address)
		conn.Close()
		os.Exit(1)
	}

	if !console.Confirm(fmt.Sprintf("%s identifies themselves as %s, are you ok with this?", address, serverId)) {
		fmt.Printf("Closing connection to %s\n", address)
		encoder.Encode("nak")
		conn.Close()
		return nil
	}

	err = encoder.Encode("ack")

	if err != nil {
		conn.Close()
		log.Fatal("Failed to write ACK to server")
	}

	return &DtpRemote{
		Socket:  conn,
		Address: tcpAddr,
		encoder: encoder,
		decoder: dec,
		Id:      *serverId,
	}
}
