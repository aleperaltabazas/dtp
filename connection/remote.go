package connection

import (
	"dtp/terminal"
	"fmt"
	"net"
)

var ConnectedRemote *dtp.Remote = nil
var AwaitingConnection *net.TCPConn = nil

func WithRemote(fn func(*dtp.Remote), onNil string) {
	if ConnectedRemote == nil {
		fmt.Println(onNil)
	} else {
		fn(ConnectedRemote)
	}
}

func IsConnected() bool {
	return ConnectedRemote != nil
}

func AcceptPending() bool {
	return AwaitingConnection != nil
}
