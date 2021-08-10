package cli

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/channels"
	"github.com/aleperaltabazas/dtp/connection"
	"github.com/aleperaltabazas/dtp/filesystem"
	"github.com/aleperaltabazas/dtp/global"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
	dtp "github.com/aleperaltabazas/dtp/terminal"
	"io"
	"log"
	"os"
)

func Send(args []string) {
	r := connection.ConnectedRemote

	if r != nil {
		switch len(args) {
		case 0:
			fmt.Println("send: missing source file")
		case 1:
			filePath := args[0]
			localFileName := filesystem.MakeAbsolute(filePath)

			if !filesystem.DoesFileExist(localFileName) {
				fmt.Println("send: no such file or directory")
				return
			}

			r.Send(codes.Send, codes.NoSource, protocol.SendFile{
				FileName: filePath,
			})
			res := <-channels.Send

			switch res.Code {
			case codes.SendAccepted:
				doSend(r, filePath)
			case codes.SendRejected:
				fmt.Printf("send: rejected: %s\n", string(res.Body))
				return
			}
		default:
			fmt.Println("send: too many arguments")
		}
	} else {
		fmt.Println("You're not connected to anything!")
	}
}

func doSend(remote *dtp.Remote, fileName string) {
	// TODO: error handling
	handler, err := os.Open(fileName)
	defer handler.Close()

	if err != nil {
		log.Fatalf("Error to read file %s: %v", fileName, err.Error())
	}

	buf := make([]byte, global.ChunkSize)
	for {
		_, err = handler.Read(buf)

		if err != nil {
			if err == io.EOF {
				remote.Send(codes.FilePart, codes.SendAccepted, protocol.FilePart{
					FileName:    fileName,
					Buffer:      nil,
					MoreContent: false,
				})
				_ = <-channels.FilePartAcknowledged
				break
			}
			log.Fatalf("Error reading file: %v", err)
		}

		remote.Send(codes.FilePart, codes.SendAccepted, protocol.FilePart{
			FileName:    fileName,
			Buffer:      buf,
			MoreContent: true,
		})
		_ = <-channels.FilePartAcknowledged
	}
}
