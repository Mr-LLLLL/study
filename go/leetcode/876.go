package main

func middleNode(head *ListNode) *ListNode {
	first := head
	second := head
	for second.Next != nil && second.Next.Next != nil {
		first = first.Next
		second = second.Next.Next
	}
	if second.Next != nil {
		first = first.Next
	}

	return first
}
