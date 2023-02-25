package leetcode

func Code_interview_01_05(first, second string) bool {
	len1st, len2nd := len(first), len(second)

	dp := make([][]int, len1st+1)
	for i := 0; i <= len1st; i++ {
		dp[i] = make([]int, len2nd+1)
		dp[i][0] = i
	}

	for j := 1; j <= len2nd; j++ {
		dp[0][j] = j
	}

	for i := 0; i < len1st; i++ {
		for j := 0; j < len2nd; j++ {
			if first[i] == second[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j+1], dp[i+1][j], dp[i][j]) + 1
			}
		}
	}

	return dp[len1st][len2nd] <= 1
}
