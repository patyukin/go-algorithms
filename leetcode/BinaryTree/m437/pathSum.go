package m437

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	pathSums := make(map[int]int)
	pathSums[0] = 1

	var dfs func(node *TreeNode, sum, targetSum int, pathSums map[int]int) int

	dfs = func(node *TreeNode, sum, targetSum int, pathSums map[int]int) int {
		if node == nil {
			return 0
		}

		sum += node.Val
		count, ok := pathSums[sum-targetSum]
		if !ok {
			count = 0
		}

		pathSums[sum]++
		count += dfs(node.Left, sum, targetSum, pathSums) + dfs(node.Right, sum, targetSum, pathSums)
		pathSums[sum]--

		return count
	}

	return dfs(root, 0, targetSum, pathSums)
}
