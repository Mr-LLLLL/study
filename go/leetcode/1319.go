package main

type unionFind1319 struct {
	parent   []int
	setCount int // 当前连通分量数目
}

func newUnionFind1319(n int) *unionFind1319 {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &unionFind1319{parent, n}
}

func (uf *unionFind1319) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind1319) union(x, y int) {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return
	}
	uf.parent[fy] = fx
	uf.setCount--
}

func Code_1319(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	uf := newUnionFind1319(n)
	for _, c := range connections {
		uf.union(c[0], c[1])
	}

	return uf.setCount - 1
}
