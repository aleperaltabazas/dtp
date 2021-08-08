package cli

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/channels"
	"github.com/aleperaltabazas/dtp/connection"
	"github.com/aleperaltabazas/dtp/protocol/codes"
	dtp "github.com/aleperaltabazas/dtp/terminal"
	"time"
)

func Ping() {
	connection.WithRemote(func(remote *dtp.Remote) {
		start := time.Now().UnixNano()
		startMillis := start / 1000000
		err := remote.Send(codes.Ping, codes.NoSource, nil)

		if err != nil {
			fmt.Printf("Failed to ping %s: %e", remote.Address(), err)
		}

		fmt.Printf("Awaiting ping response...\n")
		_ = <-channels.Ping
		end := time.Now().UnixNano()
		endMillis := end / 1000000
		fmt.Printf("Ping: %v\n", endMillis-startMillis)
	}, "You are not connected to anything!")
}
