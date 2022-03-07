/*
 * @Author: 馬濤
 * @Date: 2021-01-15 11:12:21
 * @LastEditTime: 2021-01-15 14:54:17
 * @LastEditors: 馬濤
 * @Description:
 * @FilePath: /cf/eventbus/worker.go
 * @MT is your father.
 */

package eventbus

import (
	"runtime"
	"time"
)

type goWorker struct {
	pool *Pool

	task chan *task

	// 回收时间 放回时更新
	recycleTime time.Time
}

func (w *goWorker) run() {
	w.pool.incRunning()
	go func() {
		defer func() {
			w.pool.decRunning()
			w.pool.workerCache.Put(w)
			if p := recover(); p != nil {
				if ph := w.pool.options.PanicHandler; ph != nil {
					ph(p)
				} else {
					w.pool.options.Logger.Printf("worker exits from a panic: %v\n", p)
					var buf [4096]byte
					n := runtime.Stack(buf[:], false)
					w.pool.options.Logger.Printf("worker exits from panic: %s\n", string(buf[:n]))
				}
			}
		}()

		for t := range w.task {
			if t == nil {
				return
			}
			t.f(t.eventHandler, t.args...)
			if ok := w.pool.revertWorker(w); !ok {
				return
			}
		}
	}()
}
