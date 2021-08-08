package connection

import (
	"github.com/aleperaltabazas/dtp/filesystem"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
	dtp "github.com/aleperaltabazas/dtp/terminal"
)

func ls(remote *dtp.Remote, m *protocol.Message) {
	var path string

	if len(m.Body) == 0 {
		path = filesystem.GetCurrentDirectory()
	} else {
		m.Deserialize(path)
	}

	files := filesystem.ListDirectory(path)

	remote.Send(codes.Ack, m.Code, files)
}
