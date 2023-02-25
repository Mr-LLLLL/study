package probles

func Code_713(nums []int, k int) int {
	left, ans := 0, 0
	prod := 1
	for right, v := range nums {
		prod *= v
		for prod >= k && left <= right {
			prod /= nums[left]
			left++
		}
		ans += right - left + 1
	}

	return ans
}
