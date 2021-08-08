package global

import "sync"

var Stop = false
var StopLock = sync.Mutex{}
