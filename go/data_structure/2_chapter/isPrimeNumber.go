package main

import (
	"fmt"
)

func main() {
	var n int
	for true {
		fmt.Print("please input number:")
		fmt.Scanln(&n)
		fmt.Println(isPrimeNumber(n))
	}
}

var array []bool

func isPrimeNumber(n int) bool {
	if n >= len(array) {
		eratosthenes(n + 1)
	}
	return !array[n]
}

func eratosthenes(n int) {
	if n < 2 {
		n = 2
	}
	array = make([]bool, n)
	// 0 and 1 is not prime number
	array[0] = true
	array[1] = true
	for i := 2; float64(i) < sqrt(float64(n), 1.0); i++ {
		if !array[i] {
			for j := i * i; j < n; j += i {
				array[j] = true
			}
		}
	}
}

func sqrt(n, precision float64) float64 {
	res := n
	for res*res <= n-precision || res*res >= n+precision {
		res = (n/res + res) / 2
	}
	return res
}
