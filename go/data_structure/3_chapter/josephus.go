package main

import (
	"fmt"
)

func main() {
	var n, k int
	for true {
		fmt.Println("please input n number and k kids")
		fmt.Scan(&n, &k)
		fmt.Printf("winner is :%d\n", josephus(n, k))
	}
}

func josephus(n, k int) int {
	pos := 0
	for i := 2; i <= n; i++ {
		pos = (pos + k) % i
	}
	return pos + 1
}
