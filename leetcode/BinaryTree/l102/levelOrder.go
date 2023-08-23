package l102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var levels [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var levelValues []int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelValues = append(levelValues, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		levels = append(levels, levelValues)
	}

	return levels
}

func traverse(node *TreeNode, level int, result *[][]int) {
	if node == nil {
		return
	}

	if level == len(*result) {
		*result = append(*result, make([]int, 0))
	}

	(*result)[level] = append((*result)[level], node.Val)
	traverse(node.Left, level+1, result)
	traverse(node.Right, level+1, result)
}

func levelOrderRecursive(root *TreeNode) [][]int {
	var result [][]int

	traverse(root, 0, &result)

	return result
}
