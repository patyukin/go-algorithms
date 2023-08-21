package l199

// https://leetcode.com/problems/binary-tree-right-side-view

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var levelValue int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			levelValue = node.Val

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, levelValue)
	}

	return result
}
