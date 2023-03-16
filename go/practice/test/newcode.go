package main

import (
	"fmt"
	"strconv"
)

func TestWithInput() {
	transTest()
	dupAndOsrtTest()
	emptyBottleTest()
}

func transTest() {
	var s string
	fmt.Scanln(&s)
	fmt.Println(trans(s))
}

func trans(s string) int {
	getNum := func(s byte) int {
		if s >= '0' && s <= '9' {
			n, _ := strconv.Atoi(string(s))
			return n
		}

		return int(s-'A') + 10
	}
	sum := 0
	bit := len(s) - 3
	for i, v := range s[2:] {
		tmp := getNum(byte(v))
		for j := 0; j < bit-i; j++ {
			tmp *= 16
		}
		sum += tmp
	}
	return sum
}

func dupAndOsrtTest() {
	a := 0
	cnt := 0
	arr := make([]int, 0)
	for {
		fmt.Scanln(&a)
		cnt++
		arr = append(arr, a)
		if cnt > arr[0] {
			break
		}
	}
	for _, v := range dupAndOsrt(arr[1:]) {
		fmt.Println(v)
	}
}

func dupAndOsrt(arr []int) []int {
	tmp := [500]bool{}
	for _, v := range arr {
		tmp[v-1] = true
	}
	res := make([]int, 0)
	for i, v := range tmp {
		if v {
			res = append(res, i+1)
		}
	}

	return res
}

func emptyBottleTest() {
	a := 0
	arr := make([]int, 0)
	for {
		fmt.Scanln(&a)
		if a == 0 {
			break
		}
		arr = append(arr, a)
	}
	for _, v := range arr {
		fmt.Println(emptyBottle(v))
	}
}

func emptyBottle(n int) int {
	total := 0
	for n >= 3 {
		total += n / 3
		n = n%3 + n/3
	}
	if n == 2 {
		total++
	}
	return total
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类
 * @param k int整型
 * @return ListNode类
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	reverse := func(node *ListNode, k int) (*ListNode, bool) {
		dummy := &ListNode{
			Next: node,
		}
		for i := 0; i < k; i++ {
			if node == nil {
				return dummy.Next, false
			}
			node = node.Next
		}

		tail := dummy.Next
		for i := 1; i < k; i++ {
			node = dummy.Next
			dummy.Next = tail.Next
			tail.Next = tail.Next.Next
			dummy.Next.Next = node
		}
		return dummy.Next, true
	}

	var (
		b     bool
		dummy = &ListNode{
			Next: head,
		}
		node = dummy
	)
	for node != nil {
		node.Next, b = reverse(node.Next, k)
		if !b {
			break
		}

		for i := 0; i < k; i++ {
			node = node.Next
		}
	}

	return dummy.Next
}
