package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	var err error = nil
	for err == nil {
		fmt.Print("please input a number:")
		_, err = fmt.Scan(&n)
		fmt.Println("the number( " + fmt.Sprintf("%f", n) + " )square root is : " + fmt.Sprintf("%f", sqrt(n)))
	}
}

// return the result of square root
// newton algorithm
func sqrt(n float64) float64 {
	res := n
	for math.Abs(res*res-n) > 0.1 {
		res = (n/res + res) / 2
	}
	return res
}
