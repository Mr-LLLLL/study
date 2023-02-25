package leetcode

import (
	"container/heap"
)

func Code_1202(s string, pairs [][]int) string {
	uf := newUnionFind1202(len(s))
	for _, v := range pairs {
		uf.union(v[0], v[1])
	}

	m := make(map[int]*ByteHeap)
	for i := 0; i < len(s); i++ {
		if _, ok := m[uf.find(i)]; !ok {
			m[uf.find(i)] = new(ByteHeap)
		}

		*m[uf.find(i)] = append(*m[uf.find(i)], s[i])
	}

	for _, v := range m {
		heap.Init(v)
	}

	res := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		res = append(res, heap.Pop(m[uf.find(i)]).(byte))
	}

	return string(res)
}

type unionFind1202 struct {
	parent []int
}

func newUnionFind1202(n int) *unionFind1202 {
	obj := new(unionFind1202)
	obj.parent = make([]int, n)

	for i := 0; i < n; i++ {
		obj.parent[i] = i
	}

	return obj
}

func (o *unionFind1202) union(x, y int) {
	fx, fy := o.find(x), o.find(y)
	if fx == fy {
		return
	}

	if fx > fy {
		fx, fy = fy, fx
	}

	o.parent[fy] = fx
}

func (o *unionFind1202) find(i int) int {
	if i != o.parent[i] {
		o.parent[i] = o.find(o.parent[i])
	}

	return o.parent[i]
}

type ByteHeap []byte

func (h ByteHeap) Len() int           { return len(h) }
func (h ByteHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h ByteHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ByteHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(byte))
}

func (h *ByteHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
