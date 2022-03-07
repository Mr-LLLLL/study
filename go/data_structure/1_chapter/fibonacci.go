package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var err error = nil
	for err == nil {
		fmt.Print("input the fibonacci number: ")
		_, err = fmt.Scanln(&n)
		fmt.Println("the fib(" + strconv.Itoa(n) + ") = " + strconv.Itoa(fib(n)))
	}
}

func fib(n int) int {
	f, g := 1, 0
	for n--; n > 0; n-- {
		f = f + g
		g = f - g
	}
	return g
}
