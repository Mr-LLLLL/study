package probles

func Code_42(height []int) int {
	ans := 0
	left := 0
	right := len(height) - 1
	leftMax, rightMax := 0, 0

	for left < right {
		leftMax = max(height[left], leftMax)
		rightMax = max(height[right], rightMax)
		if leftMax < rightMax {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}

	return ans
}
