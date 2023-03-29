package leetcode

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	ans := make([]int, n+1)
	ans[1] = 1
	for i, step := 2, 1; i <= n; i++ {
		if i == step<<1 {
			ans[i] = 1
			step = i
		} else {
			ans[i] = ans[i-step] + 1
		}
	}
	return ans
}
