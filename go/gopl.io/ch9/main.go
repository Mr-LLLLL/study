package main

import (
	"fmt"
)

func main() {
	var in interface{}
	var i int
	in = i
	fmt.Printf("%p, %p\n", &i, &in)
	fmt.Printf("%d, %d\n", i, in)

	var j = test()
	fmt.Printf("%p\n", &j)
}

func test() (i int) {
	fmt.Printf("%p\n", &i)
	return
}
