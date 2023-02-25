package leetcode

import (
	"container/list"
)

type MinStack struct {
	stack, minStack *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	s := new(MinStack)
	s.stack = list.New()
	s.minStack = list.New()

	return *s
}

func (this *MinStack) Push(x int) {
	this.stack.PushBack(x)
	if this.minStack.Len() == 0 {
		this.minStack.PushBack(x)

		return
	}

	if x < this.minStack.Back().Value.(int) {
		this.minStack.PushBack(x)
	} else {
		this.minStack.PushBack(this.minStack.Back().Value)
	}
}

func (this *MinStack) Pop() {
	this.stack.Remove(this.stack.Back())
	this.minStack.Remove(this.minStack.Back())
}

func (this *MinStack) Top() int {
	return this.stack.Back().Value.(int)
}

func (this *MinStack) Min() int {
	return this.minStack.Back().Value.(int)
}
