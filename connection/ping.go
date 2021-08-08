package connection

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/console"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
	dtp "github.com/aleperaltabazas/dtp/terminal"
)

func Ping(r *dtp.Remote, m *protocol.Message) {
	err := r.Send(codes.Ping, codes.Ping, nil)

	if err != nil {
		fmt.Printf("Failed to answer the ping request: %e", err)
		console.NewLine()
		ShowConsolePrompt()
	}
}
