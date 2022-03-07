package main

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const mutexLocked = 1 << iota // mutex is locked

// Mutex is simple sync.Mutex + ability to try to Lock.
type Mutex struct {
	in sync.Mutex
}

// Unlock unlocks m.
// It is a run-time error if m is not locked on entry to Unlock.
//
// A locked Mutex is not associated with a particular goroutine.
// It is allowed for one goroutine to lock a Mutex and then
// arrange for another goroutine to unlock it.
func (m *Mutex) Unlock() {
	m.in.Unlock()
}

// TryLock tries to lock m. It returns true in case of success, false otherwise.
func (m *Mutex) TryLock() bool {
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.in)), 0, mutexLocked)
}
