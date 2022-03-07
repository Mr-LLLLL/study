package main

func Code_1339(root *TreeNode) int {
	Sum4Node(root)
	return getMax(root, root.Val) % 1000000007
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Sum4Node(root *TreeNode) {
	if root.Left != nil {
		Sum4Node(root.Left)
		root.Val += root.Left.Val
	}
	if root.Right != nil {
		Sum4Node(root.Right)
		root.Val += root.Right.Val
	}
}

func getMax(node *TreeNode, rootVal int) int {
	lmax := 0
	rmax := 0

	curr := node.Val * (rootVal - node.Val)

	if node.Left != nil {
		lmax = node.Left.Val * (rootVal - node.Left.Val)
	}
	if node.Right != nil {
		rmax = node.Right.Val * (rootVal - node.Right.Val)
	}

	if curr > lmax && curr > rmax {
		return curr
	}

	if lmax > rmax {
		return getMax(node.Left, rootVal)
	} else {
		return getMax(node.Right, rootVal)
	}
}
