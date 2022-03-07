package main

func Code_1568(grid [][]int) int {
	return NewGrid(grid).minDay()
}

type grid struct {
	grid [][]int
	dx   [4]int
	dy   [4]int
	xLen int
	yLen int
}

func NewGrid(arr [][]int) *grid {
	grid := new(grid)
	grid.dx = [4]int{0, 1, 0, -1}
	grid.dy = [4]int{1, 0, -1, 0}
	grid.xLen = len(arr)
	grid.yLen = len(arr[0])
	grid.grid = arr

	return grid
}

func (g *grid) dfs(x, y int) {
	g.grid[x][y] = 2
	for i := 0; i < len(g.dx); i++ {
		tx := g.dx[i] + x
		ty := g.dy[i] + y
		if tx < 0 || tx >= g.xLen || ty < 0 || ty >= g.yLen || g.grid[tx][ty] != 1 {
			continue
		}
		g.dfs(tx, ty)
	}
}

func (g *grid) count() int {
	cnt := 0
	for i := 0; i < g.xLen; i++ {
		for j := 0; j < g.yLen; j++ {
			if g.grid[i][j] == 1 {
				cnt++
				g.dfs(i, j)
			}
		}
	}

	// restore grid
	for i := 0; i < g.xLen; i++ {
		for j := 0; j < g.yLen; j++ {
			if g.grid[i][j] == 2 {
				g.grid[i][j] = 1
			}
		}
	}

	return cnt
}

func (g *grid) minDay() int {
	if g.count() != 1 {
		return 0
	}
	for i := 0; i < g.xLen; i++ {
		for j := 0; j < g.yLen; j++ {
			if g.grid[i][j] == 1 {
				g.grid[i][j] = 0
				if g.count() != 1 {
					return 1
				}
				g.grid[i][j] = 1
			}
		}
	}

	return 2
}
