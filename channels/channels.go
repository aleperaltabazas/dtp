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
var FileParteReceived = make(chan protocol.Message)

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
	case codes.Send:
		Send <- *m
	case codes.SendAccepted:
		FileParteReceived <- *m
	default:
		fmt.Printf("Unexpected source %s\n", m.Source)
	}
	return false
}
