package leetcode

import (
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	// return quickSelect(nums, k)
	return heapSelect(nums, k)
}

func heapSelect(a []int, k int) int {
	h := new(_heap)
	h.heapify(a[:k])

	for _, v := range a[k:] {
		h.push(v)
	}

	return h.top()
}

type _heap struct {
	arr []int
	n   int
}

func (h *_heap) heapify(a []int) *_heap {
	h.n = len(a)
	for _, v := range a {
		h.arr = append(h.arr, v)

		i := len(h.arr) - 1
		for i != 0 {
			p := (i - 1) >> 1
			if h.arr[i] < h.arr[p] {
				h.arr[i], h.arr[p] = h.arr[p], h.arr[i]
			} else {
				break
			}
			i = p
		}

	}

	return h
}

func (h *_heap) push(n int) {
	if n <= h.top() {
		return
	}

	h.arr[0] = n
	i := 0
	for {
		l := 2*i + 1
		r := 2 * (i + 1)

		min := 0
		if l < h.n && r < h.n {
			if h.arr[l] < h.arr[r] {
				min = l
			} else {
				min = r
			}
		} else if l < h.n {
			min = l
		} else if r < h.n {
			min = r
		} else {
			break
		}

		if h.arr[i] > h.arr[min] {
			h.arr[i], h.arr[min] = h.arr[min], h.arr[i]
			i = min
		} else {
			break
		}
	}
}

func (h _heap) top() int {
	return h.arr[0]
}

func quickSelect(a []int, k int) int {
	index := 0
	l := 0
	r := len(a) - 1
	for r > l {
		index = partition(a, l, r, rand.Intn(r-l)+l)
		if index < k-1 {
			l = index + 1
		} else if index > k-1 {
			r = index - 1
		} else {
			return a[k-1]
		}
	}

	return a[k-1]
}

func partition(a []int, l, r, pilot int) int {
	for l < r {
		for l < r {
			if a[l] > a[pilot] {
				l++
				continue
			}
			a[l], a[pilot] = a[pilot], a[l]
			pilot = l
			break
		}

		for l < r {
			if a[r] <= a[pilot] {
				r--
				continue
			}
			a[r], a[pilot] = a[pilot], a[r]
			pilot = r
			break
		}
	}

	return pilot
}
