package main

import (
	"fmt"
	"log"
)

func main() {
	var n int
	for true {
		fmt.Println("please input arry:")
		var arr []int
		for true {
			var i int
			_, err := fmt.Scan(&i)
			if err != nil {
				break
			}
			arr = append(arr, i)
		}
		fmt.Print("please input shift number:")
		_, err := fmt.Scan(&n)
		if err != nil {
			log.Fatal(err)
		}
		arr1 := make([]int, len(arr))
		copy(arr1, arr)
		fmt.Println("the left shifted arr is:", leftShift(arr, n))
		fmt.Println("the right shifted arr is:", rightshift(arr1, n))
	}
}

func rightshift(arr []int, n int) []int {
	length := len(arr)
	n %= length
	reverse(arr, 0, length-1)
	reverse(arr, 0, n-1)
	reverse(arr, n, length-1)
	return arr
}

func leftShift(arr []int, n int) []int {
	length := len(arr)
	n %= length
	reverse(arr, 0, n-1)
	reverse(arr, n, length-1)
	reverse(arr, 0, length-1)
	return arr
}

func reverse(arr []int, lo, hi int) {
	if lo < hi {
		arr[lo], arr[hi] = arr[hi], arr[lo]
		reverse(arr, lo+1, hi-1)
	}
}
