/*
 * @Author: 馬濤
 * @Date: 2021-01-09 11:28:51
 * @LastEditTime: 2021-01-15 15:00:23
 * @LastEditors: 馬濤
 * @Description:
 * @FilePath: /cf/eventbus/spinlock.go
 * @MT is your father.
 */

package eventbus

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type spinLock uint32

func (sl *spinLock) Lock() {
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		runtime.Gosched()
	}
}

func (sl *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}

// NewSpinLock instantiates a spin-lock.
func NewSpinLock() sync.Locker {
	return new(spinLock)
}
