package e101

// https://leetcode.com/problems/symmetric-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left == nil && root.Right != nil || root.Left != nil && root.Right == nil {
		return false
	}

	left := []*TreeNode{root.Left}
	right := []*TreeNode{root.Right}

	for len(left) > 0 || len(right) > 0 {
		if len(left) != len(right) {
			return false
		}

		levelSize := len(left)
		for i := 0; i < levelSize; i++ {
			leftNode := left[0]
			rightNode := right[0]

			left = left[1:]
			right = right[1:]

			if leftNode.Val != rightNode.Val {
				return false
			}

			if leftNode.Left == nil && rightNode.Right != nil || leftNode.Left != nil && rightNode.Right == nil {
				return false
			}

			if leftNode.Right == nil && rightNode.Left != nil || leftNode.Right != nil && rightNode.Left == nil {
				return false
			}

			if leftNode.Left != nil && rightNode.Right != nil {
				left = append(left, leftNode.Left)
				right = append(right, rightNode.Right)
			}

			if leftNode.Right != nil && rightNode.Left != nil {
				left = append(left, leftNode.Right)
				right = append(right, rightNode.Left)
			}
		}
	}

	return true
}

func isSymmetricRecursion(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return (left.Val == right.Val) && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

func isSymmetricLoop(root *TreeNode) bool {
	if root == nil {
		return false
	}

	var queue []*TreeNode
	queue = append(queue, root)
	queue = append(queue, root)

	for len(queue) != 0 {
		left, right := queue[0], queue[1]
		queue = queue[2:]

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		queue = append(queue, left.Left)
		queue = append(queue, right.Right)
		queue = append(queue, left.Right)
		queue = append(queue, right.Left)
	}

	return true
}
