package main

import "fmt"

func main() {
	for true {
		fmt.Println("please input array:")
		var arr []int
		for true {
			var i int
			_, err := fmt.Scan(&i)
			if err != nil {
				break
			}
			arr = append(arr, i)
		}
		bubbleSort(arr, 0, len(arr))
		fmt.Println("the bubble sorted array:", arr)
	}
}

func bubbleSort(arr []int, lo, hi int) {
	for lo < hi {
		lo = bubbleMin(arr, lo, hi)
		hi = bubbleMax(arr, lo, hi)
	}
}

func bubbleMax(arr []int, lo, hi int) int {
	last := lo
	lo++
	for ; lo < hi; lo++ {
		if arr[lo-1] > arr[lo] {
			last = lo
			arr[lo-1], arr[lo] = arr[lo], arr[lo-1]
		}
	}
	return last
}

func bubbleMin(arr []int, lo, hi int) int {
	first := hi
	hi--
	for ; lo < hi; hi-- {
		if arr[hi] < arr[hi-1] {
			first = hi
			arr[hi], arr[hi-1] = arr[hi-1], arr[hi]
		}
	}
	return first
}
