package connection

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/channels"
	"github.com/aleperaltabazas/dtp/console"
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

func Receive(r *dtp.Remote) {
	for {
		m, err := r.Receive()

		if err != nil {
			fmt.Printf("There was an error receiving message from %s: %e\n", r.Address(), err)
			continue
		}

		if m.Source != codes.NoSource {
			channels.Dispatch(m)
		} else {
			handleNewMessage(r, m)
		}
	}
}

func handleNewMessage(r *dtp.Remote ,  m * protocol.Message) {
	switch m.Code {
	case codes.Ping:
		err := r.Send(codes.Ping, codes.Ping, nil)

		if err != nil {
			fmt.Printf("Failed to answer the ping request: %e", err)
		}
		console.NewLine()
		ShowConsolePrompt()
	}
}

