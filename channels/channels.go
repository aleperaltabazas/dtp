package channels

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
)

var Ping = make(chan protocol.Message)
var Fin = make(chan protocol.Message)
var Ls = make(chan protocol.Message)
var Cd = make(chan protocol.Message)
var Send = make(chan protocol.Message)
var FilePartAcknowledged = make(chan protocol.Message)
var FilePartReceived = make(chan protocol.Message)
var Bring = make(chan protocol.Message)

func Dispatch(m *protocol.Message) bool {
	switch m.Source {
	case codes.Ping:
		Ping <- *m
	case codes.Fin:
		Fin <- *m
		return true
	case codes.ListDirectory:
		Ls <- *m
	case codes.ChangeDirectory:
		Cd <- *m
	case codes.SendFile:
		Send <- *m
	case codes.BringFile:
		Bring <- *m
	case codes.FilePart:
		FilePartAcknowledged <- *m
	case codes.PartAcknowledged:
		FilePartReceived <- *m
	default:
		fmt.Printf("Unexpected source %s\n", m.Source)
	}
	return false
}
