package leetcode

import "math"

func giveGem(gem []int, operations [][]int) int {
	for _, v := range operations {
		opeNum := gem[v[0]] >> 1

		gem[v[0]] -= opeNum
		gem[v[1]] += opeNum
	}

	max := math.MinInt
	min := math.MaxInt
	for _, v := range gem {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}
