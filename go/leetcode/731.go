package main

import (
	"container/list"
)

type MyCalendarTwo struct {
	*list.List
}

type pair struct {
	v1, v2 int
}

func Constructor731() MyCalendarTwo {
	m := MyCalendarTwo{}
	m.List = list.New()

	return m
}

func (m *MyCalendarTwo) Book(start, end int) bool {
	m.putDelta(start, 1)
	m.putDelta(end, -1)

	active := 0
	node := m.List.Front()
	for node != nil {
		p := node.Value.(*pair)
		active += p.v2
		if active >= 3 {
			m.putDelta(start, -1)
			m.putDelta(end, 1)
			return false
		}

		node = node.Next()
	}

	return true
}

func (m *MyCalendarTwo) putDelta(key int, delta int) {
	node := m.List.Front()
	if node == nil {
		m.List.PushBack(&pair{key, delta})
		return
	}

	for node != nil {
		k := node.Value.(*pair).v1
		if key < k {
			m.List.InsertBefore(&pair{key, delta}, node)
			return
		} else if key == k {
			p := node.Value.(*pair)
			p.v2 += delta
			return
		}
		node = node.Next()
	}

	m.List.PushBack(&pair{key, delta})
}
