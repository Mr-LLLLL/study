package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, k int
	var err error
	for err == nil {
		fmt.Println("please input two number: ")
		_, err = fmt.Scan(&n, &k)
		fmt.Println("the result is :" + strconv.Itoa(G(n, k)))
	}
}

//if k == 0; the function is extract rooting function
func G(n, k int) int {
	if n < 1 {
		return k
	} else {
		return G(n-2*k-1, k+1)
	}
}
