package main

import (
	"fmt"
	"time"
)

func cron() {
	ch := make(chan int)
	go func() {
		timer := time.NewTimer(time.Second * 1) // timer 只能按时触发一次，可通过Reset()重置后继续触发。
		var x int
		for {
			select {
			case <-timer.C:
				x++
				fmt.Printf("%d,%s\n", x, time.Now().Format("2006-01-02 15:04:05"))
				if x < 10 {
					timer.Reset(time.Second * 3)
				} else {
					ch <- x
				}
			}
		}
	}()
	<-ch
}

func cron1() {
	for {
		now := time.Now()
		// 计算下一个零点
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(time.Now()))
		<-t.C
		fmt.Printf("定时结算Boottime表数据，结算完成: %v\n", time.Now())
	}
}
