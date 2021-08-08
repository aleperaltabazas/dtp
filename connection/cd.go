package connection

import (
	"github.com/aleperaltabazas/dtp/filesystem"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
	dtp "github.com/aleperaltabazas/dtp/terminal"
)

func cd(r *dtp.Remote, m *protocol.Message) {
	var path string
	m.Deserialize(&path)

	if !filesystem.DoesDirectoryExist(path) {
		r.Send(codes.Nak, m.Code, nil)
	} else {
		p := filesystem.MakeAbsolute(path)
		r.Send(codes.Ack, m.Code, p)
	}
}
