package leetcode

func numIslands(grid [][]byte) int {
	type pos struct {
		i, j int
	}
	stack := make([]pos, 0)
	cnt := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '1' {
				continue
			}

			stack = append(stack, pos{
				i: i,
				j: j,
			})

			cnt++
			for len(stack) > 0 {
				pop := stack[len(stack)-1]
				grid[pop.i][pop.j] = '0'
				stack = stack[:len(stack)-1]
				if pop.j-1 >= 0 && grid[pop.i][pop.j-1] == '1' {
					stack = append(stack, pos{
						i: pop.i,
						j: pop.j - 1,
					})
				}
				if pop.j+1 < len(grid[0]) && grid[pop.i][pop.j+1] == '1' {
					stack = append(stack, pos{
						i: pop.i,
						j: pop.j + 1,
					})
				}
				if pop.i-1 >= 0 && grid[pop.i-1][pop.j] == '1' {
					stack = append(stack, pos{
						i: pop.i - 1,
						j: pop.j,
					})
				}
				if pop.i+1 < len(grid) && grid[pop.i+1][pop.j] == '1' {
					stack = append(stack, pos{
						i: pop.i + 1,
						j: pop.j,
					})
				}
			}
		}
	}

	return cnt
}
