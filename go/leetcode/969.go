package main

import (
	"sort"
)

func Code_969(arr []int) []int {
	ans := make([]int, 0)
	n := len(arr)

	b := make([]pair, n)
	for i := 0; i < n; i++ {
		b[i].v1 = i + 1
		b[i].v2 = arr[i]
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i].v2 > b[j].v2
	})

	for _, v := range b {
		for _, f := range ans {
			if v.v1 <= f {
				v.v1 = f + 1 - v.v1
			}
		}
		ans = append(ans, v.v1)
		ans = append(ans, n)
		n--
	}

	return ans
}
