package leetcode

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	fmt.Println(preorder[0])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])

	return root
}

func post(root *TreeNode) {
	if root == nil {
		return
	}

	post(root.Left)
	post(root.Right)
	fmt.Println(root.Val)
}
