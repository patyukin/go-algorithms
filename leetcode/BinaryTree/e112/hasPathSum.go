package e112

// 112. Path Sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasSum(node *TreeNode, currentSum, targetSum int) bool {
	if node == nil {
		return false
	}

	currentSum += node.Val
	if node.Right == nil && node.Left == nil && currentSum == targetSum {
		return true
	}

	isLeftHasSum := hasSum(node.Left, currentSum, targetSum)
	isRightHasSum := hasSum(node.Right, currentSum, targetSum)

	return isLeftHasSum || isRightHasSum
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return hasSum(root, 0, targetSum)
}

func hasPathSumSimple(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val
	if root.Right == nil && root.Left == nil && targetSum == 0 {
		return true
	}

	isLeftHasSum := hasPathSumSimple(root.Left, targetSum)
	isRightHasSum := hasPathSumSimple(root.Right, targetSum)

	return isLeftHasSum || isRightHasSum
}
