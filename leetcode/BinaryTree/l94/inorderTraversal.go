package l94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}

	inorder(root)

	return result
}

func inorderTraversalLoop(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)

		root = root.Right
	}

	return result
}

func inorderTraversalFaster(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	result = append(result, inorderTraversalFaster(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversalFaster(root.Right)...)

	return result
}
