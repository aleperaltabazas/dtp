package connection

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/console"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
	dtp "github.com/aleperaltabazas/dtp/terminal"
)

func Fin(r *dtp.Remote, m *protocol.Message) {
	err := r.Send(codes.Ack, codes.Fin, nil)

	if err != nil {
		fmt.Printf("There was an error sending FIN-ACK: %e\n", err)
	}

	err = r.Socket.Close()

	if err != nil {
		fmt.Printf("There was an error closing the socket: %e\n", err)
	}

	ConnectedRemote = nil
	console.NewLine()
	fmt.Printf("%s disconnected\n", r.Id)
	ShowConsolePrompt()
}
