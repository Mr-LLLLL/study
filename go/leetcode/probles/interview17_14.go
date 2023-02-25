package probles

import "math"

func smallestK(arr []int, k int) []int {
	if k == 0 {
		return nil
	}

	binHeap := make([]int, 0, k)
	for i := 0; i < k; i++ {
		binHeap = insert(binHeap, arr[i])
	}
	for i := k; i < len(arr); i++ {
		if arr[i] < binHeap[0] {
			binHeap[0] = arr[i]
			resort(binHeap)
		}
	}

	return binHeap
}

func insert(binHeap []int, n int) []int {
	binHeap = append(binHeap, n)
	i := len(binHeap) - 1
	for {
		p := ((i + 1) >> 1) - 1
		if p < 0 {
			break
		}
		if binHeap[p] > binHeap[i] {
			break
		} else {
			binHeap[p], binHeap[i] = binHeap[i], binHeap[p]
			i = p
		}
	}

	return binHeap
}

func resort(binHeap []int) {
	i := 0
	for {
		l, r := getLeft(i), getRight(i)
		switch compare(value(binHeap, i), value(binHeap, l), value(binHeap, r)) {
		case 0:
			return
		case 1:
			binHeap[i], binHeap[l] = binHeap[l], binHeap[i]
			i = l
		case 2:
			binHeap[i], binHeap[r] = binHeap[r], binHeap[i]
			i = r
		}
	}
}

func value(arr []int, i int) int {
	l := len(arr)
	if i >= l {
		return math.MinInt
	}
	return arr[i]
}

func compare(i, j, k int) int {
	if j > i && j >= k {
		return 1
	}
	if k > i && k > j {
		return 2
	}

	return 0
}

func getLeft(i int) int {
	return 2*i + 1
}

func getRight(i int) int {
	return 2 * (i + 1)
}
