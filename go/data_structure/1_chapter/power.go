package main

import "fmt"
import "strconv"

func main() {
	var n, p int
	var err error = nil
	for err == nil {
		fmt.Print("input base, power: ")
		_, err = fmt.Scan(&n, &p)
		fmt.Println("the " + strconv.Itoa(n) + " power of " + strconv.Itoa(p) + " is: " + strconv.Itoa(power(n, p)))
	}
}

func power(n, p int) int {
	if p < 0 {
		p = -p
		n = 1 / n
	}
	res, temp := 1, n
	for p != 0 {
		if (p & 1) != 0 {
			res *= temp
		}
		temp *= temp
		p >>= 1
	}
	return res
}
