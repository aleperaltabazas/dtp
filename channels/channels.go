package channels

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
)

var Ping = make(chan protocol.Message)
var Fin = make(chan protocol.Message)

func Dispatch(m *protocol.Message) bool {
	switch m.Source {
	case codes.Ping:
		Ping <- *m
	case codes.Fin:
		Fin <- *m
		return true
	default:
		fmt.Printf("Unexpected source %s\n", m.Source)
	}
	return false
}
