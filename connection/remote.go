package connection

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/channels"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
	"github.com/aleperaltabazas/dtp/terminal"
	"net"
)

var ConnectedRemote *dtp.Remote = nil
var AwaitingConnection *net.TCPConn = nil

func WithRemote(fn func(*dtp.Remote), onNil string) {
	if ConnectedRemote == nil {
		fmt.Println(onNil)
	} else {
		fn(ConnectedRemote)
	}
}

func IsConnected() bool {
	return ConnectedRemote != nil
}

func AcceptPending() bool {
	return AwaitingConnection != nil
}

func ShowConsolePrompt() {
	prefix := "> "
	if IsConnected() {
		prefix = fmt.Sprintf("%s@%s> ", ConnectedRemote.Id, ConnectedRemote.Address())
	}

	print(prefix)
}

func Receive() {
	for {
		r := ConnectedRemote

		if r == nil {
			break
		}

		m, err := r.Receive()

		if err != nil {
			fmt.Printf("There was an error receiving message from %s: %e\n", r.Address(), err)
			fmt.Println("Closing connection")

			err = r.Socket.Close()

			if err != nil {
				fmt.Printf("There was an error closing the socket: %e", err)
			}

			ConnectedRemote = nil
			break
		}

		b := false
		switch m.Source {
		//case codes.Fin:
		//	defer r.Socket.Close()
		//	ConnectedRemote = nil
		//	channels.Fin <- *m
		//	return
		case codes.NoSource:
			b = handleNewMessage(r, m)
		default:
			b = channels.Dispatch(m)
		}

		if b {
			break
		}
	}
}

func handleNewMessage(r *dtp.Remote, m *protocol.Message) bool {
	switch m.Code {
	case codes.Ping:
		ping(r, m)
	case codes.Fin:
		fin(r, m)
		return true
	case codes.ListDirectory:
		ls(r, m)
	case codes.ChangeDirectory:
		cd(r, m)
	default:
		fmt.Printf("Unexpected request code %s\n", m.Code)
	}

	return false
}
