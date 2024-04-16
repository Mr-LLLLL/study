package main

import (
	"container/list"
	"fmt"
	_ "net/http/pprof"
)

type Node1 struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node1) Set() {
	n.Val = 1
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func vec_to_() {
	vec := []int{1, 2, 3, 4, 5}
	root := &Node{Val: vec[0]}
	queue := list.New()
	queue.PushBack(root)
	for i := 1; i < len(vec); i++ {
		oldNode := queue.Front().Value.(*Node)
		queue.Remove(queue.Front())

		newNode := &Node{Val: vec[i]}

		func() {

		}()

		a := func() {

		}
		a()

		queue.PushBack(newNode)
		if i&1 == 1 {
			oldNode.Left = newNode
		} else {
			oldNode.Right = newNode
		}
	}
}

func (Node) vec_to_() {
	vec := []int{1, 2, 3, 4, 5}
	root := &Node{Val: vec[0]}
	queue := list.New()
	queue.PushBack(root)
	for i := 1; i < len(vec); i++ {
		oldNode := queue.Front().Value.(*Node)
		queue.Remove(queue.Front())
		newNode := &Node{Val: vec[i]}

		func() {

		}()

		a := func() {

		}
		a()

		queue.PushBack(newNode)
		if i&1 == 1 {
			oldNode.Left = newNode
		} else {
			oldNode.Right = newNode
		}
	}
}

func main() {
	s := "\n"

	for _, c := range s {
		fmt.Println(c)
	}
}
