package leetcode

func exist(board [][]byte, word string) bool {
	type pos struct {
		i int
		j int
	}
	type posAndIndex struct {
		p     pos
		index int
	}

	h := make(map[byte][]pos)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			c := board[i][j]
			h[c] = append(h[c], pos{i: i, j: j})
		}
	}
	isNeighbor := func(p1, p2 pos) bool {
		if (p1.i-p2.i == 1 || p2.i-p1.i == 1) && p1.j == p2.j {
			return true
		} else if (p1.j-p2.j == 1 || p2.j-p1.j == 1) && p1.i == p2.i {
			return true
		}
		return false
	}
	pop := func(stack []posAndIndex) ([]posAndIndex, posAndIndex) {
		return stack[:len(stack)-1], stack[len(stack)-1]
	}

	var (
		stack = make([]posAndIndex, 0)
		p     posAndIndex
		color = make([]posAndIndex, 0)
	)
	for _, v := range h[byte(word[0])] {
		stack = append(stack, posAndIndex{
			p:     v,
			index: 0,
		})
	}
	color = append(color, posAndIndex{index: -1})
	for len(stack) > 0 {
		stack, p = pop(stack)
		if p.index+1 == len(word) {
			return true
		}
		for {
			l := len(color)
			if p.index > color[l-1].index {
				color = append(color, p)
				break
			} else if p.index == color[l-1].index {
				color[len(color)-1] = p
				break
			} else {
				color = color[:l-1]
			}
		}
		nc := byte(word[p.index+1])
	out:
		for _, v := range h[nc] {
			if isNeighbor(p.p, v) {
				for i := 1; i < len(color); i++ {
					if color[i].p == v {
						continue out
					}
				}
				stack = append(stack, posAndIndex{p: v, index: p.index + 1})
			}
		}
	}

	return false
}
