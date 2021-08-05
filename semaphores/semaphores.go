package semaphores

import "sync"

var StdinLock = sync.Mutex{}

var atomicMutex = sync.Mutex{}

func Atomically(fn func()) {
	atomicMutex.Lock()
	fn()
	atomicMutex.Unlock()
}
