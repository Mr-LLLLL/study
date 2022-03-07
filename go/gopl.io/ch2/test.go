package main

import "fmt"

func main() {
	fmt.Printf("path: %p\n", test())
}

func test() *int {
	p := new(int)
	fmt.Printf("path: %p\n", p)
	return p
}
