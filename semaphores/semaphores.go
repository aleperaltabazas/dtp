package semaphores

import "sync"

var StdinLock = sync.Mutex{}

var atomicMutex = sync.Mutex{}

func Atomically(fn func()) {
	atomicMutex.Lock()
	fn()
	atomicMutex.Unlock()
}

var RedirectInput = false
var RedirectionLock = sync.Mutex{}
