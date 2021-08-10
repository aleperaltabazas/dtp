package cli

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/channels"
	"github.com/aleperaltabazas/dtp/connection"
	"github.com/aleperaltabazas/dtp/filesystem"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
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
			handler, err := os.Open(localFileName)

			if err != nil {
				fmt.Printf("send: failed to open file: %v\n", err.Error())
				return
			}

			r.Send(codes.SendFile, codes.NoSource, protocol.TransferFile{
				FileName: filePath,
			})
			res := <-channels.Send

			switch res.Code {
			case codes.SendAccepted:
				connection.SendFileParts(r, localFileName, handler)
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
