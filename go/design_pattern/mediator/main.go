package main

import "fmt"

func main() {
	var p, p1 *int
	i := 0
	p, p1 = &i, &i
	fmt.Println(p == p1)
}

