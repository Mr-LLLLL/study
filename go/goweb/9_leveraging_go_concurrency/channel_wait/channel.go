package main

import (
	"fmt"
	"time"
)

func printNumber(w chan struct{}) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	w <- struct{}{}
}

func printLetter(w chan struct{}) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	w <- struct{}{}
}

func main() {
	w1, w2 := make(chan struct{}), make(chan struct{})
	go printNumber(w1)
	go printLetter(w2)
	<-w1
	<-w2
}
