package connection

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/channels"
	"github.com/aleperaltabazas/dtp/filesystem"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
	dtp "github.com/aleperaltabazas/dtp/terminal"
	"os"
	"time"
)

func receiveFile(r *dtp.Remote, m *protocol.Message) {
	// TODO: error handling
	var send protocol.SendFile
	m.Deserialize(&send)

	if filesystem.DoesFileExist(send.FileName) {
		r.Send(codes.SendRejected, m.Code, fmt.Sprintf("file %s already exists", send.FileName))
	}

	fileHandler, err := os.OpenFile(send.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		r.Send(codes.SendRejected, m.Code, err.Error())
	}

	start := time.Now().UnixNano() / 1000000
	for {
		res := <-channels.FileParteReceived
		switch res.Code {
		case codes.FilePart:
			var part protocol.FilePart
			res.Deserialize(&part)

			if part.MoreContent {
				fileHandler.Write(part.Buffer)
			} else {
				end := time.Now().UnixNano() / 1000000
				fmt.Printf("File transfer completed in %v ms\n", end-start)
				return
			}
		}
	}
}
