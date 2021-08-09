package global

import (
	"net"
	"sync"
)

var Stop = false
var StopLock = sync.Mutex{}

var Id string
var Listener *net.TCPListener
