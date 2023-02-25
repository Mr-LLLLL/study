package probles

func findBall(grid [][]int) []int {
	row := len(grid[0])
	col := len(grid)
	ans := make([]int, row)
out:
	for i := 0; i < row; i++ {
		r := i
		for j := 0; j < col; j++ {
			if grid[j][r] == 1 {
				if r+1 == row || grid[j][r+1] == -1 {
					ans[i] = -1
					continue out
				} else {
					r++
				}
			} else {
				if r-1 == -1 || grid[j][r-1] == 1 {
					ans[i] = -1
					continue out
				} else {
					r--
				}
			}
		}
		ans[i] = r
	}
	return ans
}
