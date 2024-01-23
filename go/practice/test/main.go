package main

import (
	"container/list"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	_ "net/http/pprof"
	"os"
)

type Node1 struct {
	Val   int
	Left  *Node
	Right *Node
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
	f, _ := os.Open("/tmp/test/tomcat-util.jar")
	m := md5.New()
	io.Copy(m, f)
	fmt.Println(hex.EncodeToString(m.Sum(nil)))
}
