package connection

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/filesystem"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
	dtp "github.com/aleperaltabazas/dtp/terminal"
	"os"
)

func sendFile(r *dtp.Remote, m *protocol.Message) {
	// TODO: error handling
	var send protocol.TransferFile
	m.Deserialize(&send)

	if !filesystem.DoesFileExist(send.FileName) {
		r.Send(codes.BringRejected, m.Code, fmt.Sprintf("file %s doesn't exist", send.FileName))
		return
	}

	fileHandler, err := os.OpenFile(send.FileName, os.O_RDONLY, 0644)
	if err != nil {
		r.Send(codes.BringRejected, m.Code, err.Error())
		return
	}

	r.Send(codes.BringAccepted, m.Code, nil)
	go SendFileParts(r, send.FileName, fileHandler)
}
