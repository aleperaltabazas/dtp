package cli

import (
	"fmt"
	"github.com/aleperaltabazas/dtp/channels"
	"github.com/aleperaltabazas/dtp/connection"
	"github.com/aleperaltabazas/dtp/filesystem"
	"github.com/aleperaltabazas/dtp/protocol"
	"github.com/aleperaltabazas/dtp/protocol/codes"
	"os"
	"strings"
)

func Bring(args []string) {
	r := connection.ConnectedRemote

	if r != nil {
		switch len(args) {
		case 0:
			fmt.Println("bring: not enough arguments")
		case 1:
			filePath := args[0]
			remotePath := filePath
			if string(filePath[0]) != "/" {
				remotePath = r.FullPath(remotePath)
			}

			localFileName := name(filePath)

			if filesystem.DoesFileExist(localFileName) {
				fmt.Printf("bring: file %s\n already exists", localFileName)
				return
			}

			fileHandler, err := os.OpenFile(localFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Printf("bring: failed to create file: %v\n", err.Error())
				return
			}

			r.Send(codes.BringFile, codes.NoSource, protocol.TransferFile{
				FileName: remotePath,
			})
			res := <-channels.Bring

			switch res.Code {
			case codes.BringAccepted:
				connection.ReceiveFileParts(r, fileHandler)
			case codes.BringRejected:
				fmt.Printf("bring: rejected: %s\n", string(res.Body))
				return
			}
		default:
			fmt.Println("bring: too many arguments")
		}
	} else {
		fmt.Println("You're not connected to anything")
	}
}

func name(fileName string) string {
	s := strings.Split(fileName, "/")
	return s[len(s)-1]
}
