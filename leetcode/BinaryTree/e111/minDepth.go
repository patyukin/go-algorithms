package e111

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}

	if root.Right == nil && root.Left != nil {
		return 1 + minDepth(root.Left)
	}

	left := 1 + minDepth(root.Left)
	right := 1 + minDepth(root.Right)

	if left < right {
		return left
	}

	return right
}
