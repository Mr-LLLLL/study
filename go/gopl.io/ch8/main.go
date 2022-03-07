package main

import "fmt"

func main() {
	ch := make(chan int)
	ch2 := ch

	go func() {
		<-ch2
	}()
	ch <- 1

	fmt.Printf("%p\n", ch)
	fmt.Printf("%p\n", ch2)
}
