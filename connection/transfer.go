package connection

import (
	"bytes"
	"github.com/aleperaltabazas/dtp/channels"
	"github.com/aleperaltabazas/dtp/global"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
	dtp "github.com/aleperaltabazas/dtp/terminal"
	"io"
	"log"
	"os"
)

func SendFileParts(remote *dtp.Remote, fileName string, handler *os.File) {
	// TODO: error handling
	defer handler.Close()

	buf := make([]byte, global.ChunkSize)
	for {
		_, err := handler.Read(buf)

		if err != nil {
			if err == io.EOF {
				remote.Send(codes.FilePart, codes.PartAcknowledged, protocol.FilePart{
					FileName:    fileName,
					Buffer:      nil,
					MoreContent: false,
				})
				_ = <-channels.FilePartAcknowledged
				break
			}
			log.Fatalf("Error reading file: %v", err)
		}

		remote.Send(codes.FilePart, codes.PartAcknowledged, protocol.FilePart{
			FileName:    fileName,
			Buffer:      bytes.Trim(buf, "\x00"),
			MoreContent: true,
		})
		_ = <-channels.FilePartAcknowledged
	}
}

func ReceiveFileParts(r *dtp.Remote, fileHandler *os.File) {
	defer fileHandler.Close()
	for {
		res := <-channels.FilePartReceived
		switch res.Code {
		case codes.FilePart:
			var part protocol.FilePart
			res.Deserialize(&part)

			if part.MoreContent {
				fileHandler.Write(part.Buffer)
				r.Send(codes.Ack, res.Code, nil)
			} else {
				r.Send(codes.Ack, res.Code, nil)
				return
			}
		}
	}
}
