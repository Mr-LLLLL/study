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
		mergeSort(arr, 0, len(arr))
		fmt.Println("the bubble sorted array:", arr)
	}
}

func mergeSort(arr []int, lo, hi int) {
	if (hi - lo) < 2 {
		return
	}
	mi := (lo + hi) >> 1
	mergeSort(arr, lo, mi)
	mergeSort(arr, mi, hi)
	if arr[mi-1] > arr[mi] {
		merge(arr, lo, mi, hi)
	}
}

func merge(arr []int, lo, mi, hi int) {
	tmp := make([]int, mi-lo)
	copy(tmp, arr[lo:mi])

	for i, j, k := lo, 0, mi; j != mi-lo; i++ {
		if k == hi || tmp[j] <= arr[k] {
			arr[i] = tmp[j]
			j++
		} else {
			arr[i] = arr[k]
			k++
		}
	}
}
