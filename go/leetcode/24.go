package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	var node *ListNode
	for head != nil && head.Next != nil {
		node = head.Next
		head.Next = node.Next
		node.Next = head
		pre.Next = node
		pre = head
		head = head.Next
	}
	return dummy.Next
}
