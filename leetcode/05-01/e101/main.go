package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 101. Symmetric Tree
// https://leetcode.com/problems/symmetric-tree
func isSymmetricTree(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	return (t1.Val == t2.Val) && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}

func main() {
	fmt.Println(isSymmetricTree(&TreeNode{}))
}
