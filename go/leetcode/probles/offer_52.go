package probles

import "sort"

func Code_Offer_31(nums []int, target int) int {
	res := 0
	index := sort.SearchInts(nums, target)
	for index < len(nums) {
		if nums[index] == target {
			res++
			index++
		} else {
			break
		}
	}

	return res
}
