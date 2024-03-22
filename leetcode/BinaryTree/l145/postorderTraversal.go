package l145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		postorder(node.Left)
		postorder(node.Right)
		result = append(result, node.Val)
	}

	postorder(root)

	return result
}

func postorderTraversalFast(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	result = append(result, postorderTraversalFast(root.Left)...)
	result = append(result, postorderTraversalFast(root.Right)...)
	result = append(result, root.Val)

	return result
}

func postorderTraversalLoop(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append([]int{current.Val}, result...)
		if current.Left != nil {
			stack = append(stack, current.Left)
		}

		if current.Right != nil {
			stack = append(stack, current.Right)
		}
	}

	return result
}
