package connection

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/filesystem"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
	dtp "github.com/aleperaltabazas/dtp/terminal"
	"os"
)

func receiveFile(r *dtp.Remote, m *protocol.Message) {
	// TODO: error handling
	var send protocol.TransferFile
	m.Deserialize(&send)

	if filesystem.DoesFileExist(send.FileName) {
		r.Send(codes.SendRejected, m.Code, fmt.Sprintf("file %s already exists", send.FileName))
		return
	}

	fileHandler, err := os.OpenFile(send.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		r.Send(codes.SendRejected, m.Code, err.Error())
		return
	}

	r.Send(codes.SendAccepted, m.Code, nil)
	go ReceiveFileParts(r, fileHandler)
}
