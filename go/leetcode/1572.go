package leetcode

func Code_1572(mat [][]int) int {
	n := len(mat)
	res := 0
	for i := 0; i < n; i++ {
		res += mat[i][i] + mat[i][n-i-1]
	}
	return res - mat[n/2][n/2]*(n&1)
}
