package main

// https://leetcode.com/problems/binary-tree-inorder-traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorder
func calc(root *TreeNode) []int {
	var result []int

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		result = append(result, root.Val)
		inorder(root.Right)
	}

	inorder(root)

	return result
}
