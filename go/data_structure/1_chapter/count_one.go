package main

import (
	"fmt"
	"math"
)

/*
 * count "1" in binary of integer
 */

func main() {
	var n uint32
	var err error = nil
	for err == nil {
		_, err = fmt.Scanln(&n)
		fmt.Println("the binary of integer include " + fmt.Sprint(countOne(n)) + " \"1\" version 1")
		fmt.Println("the binary of integer include " + fmt.Sprint(countOne2(n)) + " \"1\" version 2")
	}
}

func countOne(n uint32) uint32 {
	n = round(n, 0)
	n = round(n, 1)
	n = round(n, 2)
	n = round(n, 3)
	n = round(n, 4)
	return n
}

func power(i uint32) uint32 {
	return 1 << i
}

func mask(i uint32) uint32 {
	return math.MaxUint32 / (power(power(i)) + 1)
}

func round(n, i uint32) uint32 {
	return (n & mask(i)) + (n >> power(i) & mask(i))
}

// this function equal to countOne, but more easy to understand
func countOne2(n uint32) uint32 {
	n = 0x55555555 & n + 0x5555555 & (n >> 1)
	n = 0x33333333 & n + 0x3333333 & (n >> 2)
	n = 0x0f0f0f0f & n + 0x00f0f0f & (n >> 4)
	n = 0x00ff00ff & n + 0x0ff00ff & (n >> 8)
	n = 0x0000ffff & n + 0x000ffff & (n >> 16)
	return n
}
