package main

import "fmt"

func main() {
	for true {
		fmt.Print("input two number:")
		var x, y int
		fmt.Scan(&x, &y)
		fmt.Printf("%d and %d gcdCN is:%d\n", x, y, greatestCommonDivisorChina(x, y))
	}
}

func greatestCommonDivisorChina(x, y int) int {
	r := 1
	for !((x&1 == 1) || (y&1 == 1)) {
		// x and y is even
		x >>= 1
		y >>= 1
		r <<= 1
	}

	for true {
		for !(x&1 == 1) {
			x >>= 1
		}
		for !(y&1 == 1) {
			y >>= 1
		}
		if x > y {
			x = x - y
		} else {
			y = y - x
		}
		if 0 == x {
			return y * r
		}
		if 0 == y {
			return x * r
		}
	}
	return 0
}
