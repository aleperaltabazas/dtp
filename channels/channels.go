package channels

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
)

var Ping = make(chan protocol.Message)

func Dispatch(m *protocol.Message) {
	switch m.Source {
	case codes.Ping:
		Ping <- *m
	default:
		fmt.Printf("Unexpected source %s\n", m.Source)
	}
}
