package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("commencing countdown.	press return to abort.")
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	fmt.Println("launch")
}
